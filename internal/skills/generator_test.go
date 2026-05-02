package skills

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gleanwork/glean-cli/internal/schema"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Register minimal test schemas so these tests don't depend on cmd/ init().
// init() in a _test.go file only runs when compiling tests in this package.
func init() {
	schema.Register(schema.CommandSchema{
		Command:     "testsearch",
		Description: "Test search description.",
		Flags: map[string]schema.FlagSchema{
			"--query": {Type: "string", Description: "Search query", Required: true},
			"--page":  {Type: "integer", Description: "Page number", Default: 1},
		},
		Example: "glean testsearch --query hello",
	})
	schema.Register(schema.CommandSchema{
		Command:     "testpins",
		Description: "Test pins description.",
		Flags:       map[string]schema.FlagSchema{},
		Example:     "glean testpins list",
	})
}

func runGenerator(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	require.NoError(t, Generate(dir))
	return dir
}

// registeredCommands returns the list of command names that Generate would
// process (registry minus skipCommands).
func registeredCommands() []string {
	var out []string
	for _, name := range schema.List() {
		if skipCommands[name] {
			continue
		}
		out = append(out, name)
	}
	return out
}

func TestGenerate_WritesRootSkill(t *testing.T) {
	dir := runGenerator(t)

	p := filepath.Join(dir, rootSkillName, "SKILL.md")
	info, err := os.Stat(p)
	require.NoError(t, err, "root SKILL.md must exist")
	assert.Greater(t, info.Size(), int64(0), "root SKILL.md must be non-empty")
}

func TestGenerate_WritesReferenceForEveryCommand(t *testing.T) {
	dir := runGenerator(t)

	for _, name := range registeredCommands() {
		p := filepath.Join(dir, rootSkillName, "reference", name+".md")
		_, err := os.Stat(p)
		assert.NoError(t, err, "expected reference file for command %q at %s", name, p)
	}
}

func TestGenerate_NoLegacySkillDirs(t *testing.T) {
	dir := runGenerator(t)

	entries, err := os.ReadDir(dir)
	require.NoError(t, err)
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		if e.Name() == rootSkillName {
			continue
		}
		assert.NotContains(t, e.Name(), skillPrefix,
			"unexpected legacy skill dir after generation: %s", e.Name())
	}
}

func TestGenerate_Idempotent(t *testing.T) {
	dir := t.TempDir()
	require.NoError(t, Generate(dir))
	first := snapshotTree(t, dir)

	require.NoError(t, Generate(dir))
	second := snapshotTree(t, dir)

	assert.Equal(t, first, second, "Generate must be idempotent")
}

func TestCleanStaleSkillDirs_RemovesLegacy(t *testing.T) {
	dir := t.TempDir()
	for _, cmd := range []string{"search", "chat"} {
		stale := filepath.Join(dir, skillPrefix+cmd)
		require.NoError(t, os.MkdirAll(stale, 0o755))
		require.NoError(t, os.WriteFile(filepath.Join(stale, "SKILL.md"), []byte("stale"), 0o644))
	}

	require.NoError(t, cleanStaleSkillDirs(dir))

	for _, cmd := range []string{"search", "chat"} {
		stale := filepath.Join(dir, skillPrefix+cmd)
		_, err := os.Stat(stale)
		assert.True(t, os.IsNotExist(err), "legacy dir %s should be removed", stale)
	}
}

func TestCleanStaleSkillDirs_PreservesRoot(t *testing.T) {
	dir := t.TempDir()
	root := filepath.Join(dir, rootSkillName)
	require.NoError(t, os.MkdirAll(root, 0o755))
	require.NoError(t, os.WriteFile(filepath.Join(root, "SKILL.md"), []byte("# root"), 0o644))

	require.NoError(t, cleanStaleSkillDirs(dir))

	info, err := os.Stat(root)
	require.NoError(t, err, "root dir must survive cleanup")
	assert.True(t, info.IsDir())
}

func TestCleanStaleSkillDirs_PreservesUnrelatedDirs(t *testing.T) {
	dir := t.TempDir()
	unrelated := filepath.Join(dir, "some-users-custom-dir")
	require.NoError(t, os.MkdirAll(unrelated, 0o755))
	require.NoError(t, os.WriteFile(filepath.Join(unrelated, "file.md"), []byte("keep me"), 0o644))

	require.NoError(t, cleanStaleSkillDirs(dir))

	_, err := os.Stat(unrelated)
	assert.NoError(t, err, "unrelated dir must not be touched by cleanup")
}

func TestReferenceFile_NoFrontmatter(t *testing.T) {
	dir := runGenerator(t)

	cmds := registeredCommands()
	require.NotEmpty(t, cmds, "at least one registered command required for this test")

	p := filepath.Join(dir, rootSkillName, "reference", cmds[0]+".md")
	body, err := os.ReadFile(p)
	require.NoError(t, err)
	content := string(body)

	assert.False(t, strings.HasPrefix(content, "---"),
		"reference file must not start with YAML frontmatter: %s", p)
	assert.NotContains(t, content, "PREREQUISITE",
		"reference file must not carry the legacy prereq line")
}

func TestRootSkill_LinksPointToReferenceFiles(t *testing.T) {
	dir := runGenerator(t)

	body, err := os.ReadFile(filepath.Join(dir, rootSkillName, "SKILL.md"))
	require.NoError(t, err)
	content := string(body)

	for _, name := range registeredCommands() {
		expected := "(reference/" + name + ".md)"
		assert.Contains(t, content, expected,
			"root SKILL.md must link to %s", expected)
	}

	assert.NotContains(t, content, "../"+skillPrefix,
		"root SKILL.md must not contain legacy ../glean-cli-<cmd>/ links")
}

func TestRootSkill_IncludesMigrationNote(t *testing.T) {
	dir := runGenerator(t)

	body, err := os.ReadFile(filepath.Join(dir, rootSkillName, "SKILL.md"))
	require.NoError(t, err)
	assert.Contains(t, string(body), "npx -y skills remove",
		"root SKILL.md must carry the per-command cleanup one-liner")
}

// snapshotTree returns a map of relative-path -> file contents for every file
// under root. Used by idempotency tests.
func snapshotTree(t *testing.T, root string) map[string]string {
	t.Helper()
	out := map[string]string{}
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		body, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		out[rel] = string(body)
		return nil
	})
	require.NoError(t, err)
	return out
}
