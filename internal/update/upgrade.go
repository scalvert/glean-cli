package update

import (
	"archive/tar"
	"archive/zip"
	"bytes"
	"compress/gzip"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/gleanwork/glean-cli/internal/httputil"
	"github.com/minio/selfupdate"
)

const (
	releaseDownloadURL = "https://github.com/gleanwork/glean-cli/releases/download"
	checksumFile       = "checksums.txt"
	binaryName         = "glean"
)

// Upgrade checks for a newer release and installs it.
// If the binary was installed via Homebrew it delegates to `brew upgrade`.
// Otherwise it downloads the appropriate archive from GitHub Releases,
// verifies its SHA-256 checksum, and atomically replaces the running binary.
func Upgrade(currentVersion string) error {
	if currentVersion == devVersion || currentVersion == "" {
		return fmt.Errorf("cannot update a dev build — build from source instead")
	}

	// Homebrew-managed install: let brew handle the upgrade.
	if isBrewInstall() {
		return brewUpgrade()
	}

	fmt.Fprintln(os.Stderr, "Checking for updates...")

	latest, err := fetchLatestTag()
	if err != nil {
		return fmt.Errorf("could not fetch latest release: %w", err)
	}

	if !isNewer(latest, currentVersion) {
		fmt.Printf("Already up to date (%s)\n", currentVersion)
		return nil
	}

	fmt.Fprintf(os.Stderr, "Updating %s → %s\n", currentVersion, latest)

	assetName := assetFilename()
	assetURL := fmt.Sprintf("%s/%s/%s", releaseDownloadURL, latest, assetName)
	checksumURL := fmt.Sprintf("%s/%s/%s", releaseDownloadURL, latest, checksumFile)

	archive, err := download(assetURL)
	if err != nil {
		return fmt.Errorf("download failed: %w", err)
	}

	if err := verifyChecksum(archive, assetName, checksumURL); err != nil {
		return fmt.Errorf("checksum verification failed: %w", err)
	}

	binary, err := extractBinary(assetName, archive)
	if err != nil {
		return fmt.Errorf("could not extract binary: %w", err)
	}

	if err := selfupdate.Apply(bytes.NewReader(binary), selfupdate.Options{}); err != nil {
		return fmt.Errorf("could not apply update: %w", err)
	}

	fmt.Printf("Updated to %s\n", latest)
	return nil
}

// isBrewInstall reports whether the running binary lives inside a Homebrew Cellar.
func isBrewInstall() bool {
	exe, err := os.Executable()
	if err != nil {
		return false
	}
	return strings.Contains(exe, "/Cellar/") || strings.Contains(exe, "/homebrew/")
}

// brewUpgrade runs `brew upgrade gleanwork/tap/glean-cli`.
func brewUpgrade() error {
	fmt.Fprintln(os.Stderr, "Detected Homebrew install — running brew upgrade...")
	cmd := exec.Command("brew", "upgrade", "gleanwork/tap/glean-cli")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// assetFilename returns the expected archive name for the current platform.
// Matches the GoReleaser name_template in .goreleaser.yaml:
// glean-cli_{OS}_{arch}.tar.gz  (e.g. glean-cli_Darwin_arm64.tar.gz)
func assetFilename() string {
	goos := runtime.GOOS
	goarch := runtime.GOARCH

	// GoReleaser uses the `title` filter: "darwin" → "Darwin"
	osName := strings.ToUpper(goos[:1]) + goos[1:]
	archName := goarch
	if goarch == "amd64" {
		archName = "x86_64"
	}

	ext := "tar.gz"
	if goos == "windows" {
		ext = "zip"
	}

	return fmt.Sprintf("glean-cli_%s_%s.%s", osName, archName, ext)
}

// download fetches a URL and returns the body bytes.
func download(url string) ([]byte, error) {
	client := httputil.NewHTTPClient(120 * time.Second)
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP %d from %s", resp.StatusCode, url)
	}
	return io.ReadAll(resp.Body)
}

// verifyChecksum fetches checksums.txt and confirms the archive matches.
func verifyChecksum(archive []byte, assetName, checksumURL string) error {
	raw, err := download(checksumURL)
	if err != nil {
		return fmt.Errorf("could not fetch checksums: %w", err)
	}

	want := ""
	for line := range strings.SplitSeq(string(raw), "\n") {
		fields := strings.Fields(line)
		if len(fields) == 2 && fields[1] == assetName {
			want = fields[0]
			break
		}
	}
	if want == "" {
		return fmt.Errorf("no checksum entry found for %s", assetName)
	}

	sum := sha256.Sum256(archive)
	got := hex.EncodeToString(sum[:])
	if got != want {
		return fmt.Errorf("checksum mismatch: got %s want %s", got, want)
	}
	return nil
}

// extractBinary pulls the `glean` (or `glean.exe`) binary out of the archive.
func extractBinary(assetName string, data []byte) ([]byte, error) {
	if strings.HasSuffix(assetName, ".zip") {
		return extractFromZip(data)
	}
	return extractFromTarGz(data)
}

func extractFromTarGz(data []byte) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer gz.Close()

	tr := tar.NewReader(gz)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if hdr.Name == binaryName || strings.HasSuffix(hdr.Name, "/"+binaryName) {
			return io.ReadAll(tr)
		}
	}
	return nil, fmt.Errorf("glean binary not found in archive")
}

func extractFromZip(data []byte) ([]byte, error) {
	r, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return nil, err
	}
	for _, f := range r.File {
		if f.Name == binaryName+".exe" || strings.HasSuffix(f.Name, "/"+binaryName+".exe") ||
			f.Name == binaryName || strings.HasSuffix(f.Name, "/"+binaryName) {
			rc, err := f.Open()
			if err != nil {
				return nil, err
			}
			defer rc.Close()
			return io.ReadAll(rc)
		}
	}
	return nil, fmt.Errorf("glean binary not found in archive")
}
