package cmd

import (
	"bytes"
	"testing"

	"github.com/gleanwork/glean-cli/internal/testutils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEntitiesListDryRun(t *testing.T) {
	_, cleanup := testutils.SetupTestWithResponse(t, []byte(`{}`))
	defer cleanup()

	b := bytes.NewBufferString("")
	cmd := NewCmdEntities()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"list", "--dry-run", "--json", `{}`})

	err := cmd.Execute()
	require.NoError(t, err)
}

func TestEntitiesHelp(t *testing.T) {
	b := bytes.NewBufferString("")
	cmd := NewCmdEntities()
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--help"})

	err := cmd.Execute()
	assert.NoError(t, err)
	assert.Contains(t, b.String(), "Usage")
}

func TestEntitiesListBadEnumShowsValidValues(t *testing.T) {
	cmd := NewCmdEntities()
	errB := bytes.NewBufferString("")
	cmd.SetErr(errB)
	cmd.SetArgs([]string{"list", "--json", `{"entityType":"BADVALUE"}`})

	err := cmd.Execute()
	assert.Error(t, err)
	errMsg := err.Error()
	assert.NotContains(t, errMsg, "ListEntitiesRequestEntityType")
	assert.Contains(t, errMsg, "PEOPLE")
}
