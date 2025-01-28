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
curl -fsSL "$DOWNLOAD_URL" -o "$TMP_DIR/glean.tar.gz"
tar -xzf "$TMP_DIR/glean.tar.gz" -C "$TMP_DIR"

# Install binary
INSTALL_DIR="/usr/local/bin"
if [ ! -w "$INSTALL_DIR" ]; then
  echo "Installing Glean CLI requires sudo access to $INSTALL_DIR"
  sudo mv "$TMP_DIR/glean" "$INSTALL_DIR/"
  sudo chmod +x "$INSTALL_DIR/glean"
else
  mv "$TMP_DIR/glean" "$INSTALL_DIR/"
  chmod +x "$INSTALL_DIR/glean"
fi

echo "Glean CLI has been installed to $INSTALL_DIR/glean"
echo "Run 'glean --help' to get started"