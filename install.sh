#!/bin/sh
set -e

# Glean CLI Installer
#
# Usage:
#   curl -fsSL https://raw.githubusercontent.com/scalvert/glean-cli/main/install.sh | sh

LATEST_VERSION=$(curl -s https://api.github.com/repos/scalvert/glean-cli/releases/latest | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Convert architecture names
case "$ARCH" in
  x86_64) ARCH="x86_64" ;;
  amd64) ARCH="x86_64" ;;
  arm64) ARCH="arm64" ;;
  aarch64) ARCH="arm64" ;;
  *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

# Construct download URL
DOWNLOAD_URL="https://github.com/scalvert/glean-cli/releases/download/${LATEST_VERSION}/glean-cli_${OS}_${ARCH}.tar.gz"

# Create temporary directory
TMP_DIR=$(mktemp -d)
cleanup() {
  rm -rf "$TMP_DIR"
}
trap cleanup EXIT

# Download and extract
echo "Downloading Glean CLI ${LATEST_VERSION}..."
echo "Download URL: $DOWNLOAD_URL"
curl -fsSL "$DOWNLOAD_URL" -o "$TMP_DIR/glean.tar.gz"

echo "Extracting archive..."
tar -tvf "$TMP_DIR/glean.tar.gz"
tar -xzf "$TMP_DIR/glean.tar.gz" -C "$TMP_DIR"

echo "Verifying extracted contents..."
ls -la "$TMP_DIR"

# Install binary
INSTALL_DIR="/usr/local/bin"

# Create install directory if it doesn't exist
if [ ! -d "$INSTALL_DIR" ]; then
  echo "Creating $INSTALL_DIR directory..."
  if ! sudo mkdir -p "$INSTALL_DIR"; then
    echo "Failed to create $INSTALL_DIR"
    exit 1
  fi
fi

# Set proper ownership for /usr/local/bin if it was just created
if [ -d "$INSTALL_DIR" ]; then
  sudo chown -R $(whoami) "$INSTALL_DIR" 2>/dev/null || true
fi

# Find the glean binary
GLEAN_BINARY=$(find "$TMP_DIR" -type f -name "glean" -o -name "glean.exe" | head -n 1)

if [ -z "$GLEAN_BINARY" ]; then
  echo "Error: Could not find glean binary in extracted files"
  echo "Contents of temp directory:"
  ls -R "$TMP_DIR"
  exit 1
fi

# Attempt installation
echo "Installing to $INSTALL_DIR..."
echo "Installing $GLEAN_BINARY to $INSTALL_DIR/glean"

if [ -w "$INSTALL_DIR" ]; then
  # We have write permission
  mv "$GLEAN_BINARY" "$INSTALL_DIR/glean"
  chmod +x "$INSTALL_DIR/glean"
else
  # We need sudo
  echo "Elevated permissions required for installation..."
  sudo mv "$GLEAN_BINARY" "$INSTALL_DIR/glean"
  sudo chmod +x "$INSTALL_DIR/glean"
fi

echo "✨ Glean CLI ${LATEST_VERSION} has been installed to $INSTALL_DIR/glean"
echo "Run 'glean --help' to get started"