#!/bin/bash

set -e

VERSION="1.0.0"
OUTPUT_DIR="./build"
SRC_DIR="/Users/ewenhalgand/Desktop/Fichier-Dev/DevPerso/Golang/passtok-generator"

echo "🚀 Starting cross-platform build of ptk version $VERSION..."

mkdir -p "$OUTPUT_DIR"

# Cross-platform compilation
echo "🖥️ Building for Mac Intel (amd64)..."
GOOS=darwin GOARCH=amd64 go build -o "$OUTPUT_DIR/ptk-mac-intel" "$SRC_DIR"

echo "🍏 Building for Mac ARM64 (Apple Silicon)..."
GOOS=darwin GOARCH=arm64 go build -o "$OUTPUT_DIR/ptk-mac-arm64" "$SRC_DIR"

echo "🐧 Building for Linux amd64..."
GOOS=linux GOARCH=amd64 go build -o "$OUTPUT_DIR/ptk-linux" "$SRC_DIR"

echo "🪟 Building for Windows amd64..."
GOOS=windows GOARCH=amd64 go build -o "$OUTPUT_DIR/ptk-windows.exe" "$SRC_DIR"

echo "✅ Build complete."

# Basic packaging (zip archives)
echo "📦 Packaging binaries..."

zip -j "$OUTPUT_DIR/ptk-mac-intel-$VERSION.zip" "$OUTPUT_DIR/ptk-mac-intel"
zip -j "$OUTPUT_DIR/ptk-mac-arm64-$VERSION.zip" "$OUTPUT_DIR/ptk-mac-arm64"
zip -j "$OUTPUT_DIR/ptk-linux-$VERSION.zip" "$OUTPUT_DIR/ptk-linux"
zip -j "$OUTPUT_DIR/ptk-windows-$VERSION.zip" "$OUTPUT_DIR/ptk-windows.exe"

echo "✅ Packaging complete. Files saved in $OUTPUT_DIR"

# Simple install step for macOS and Linux
OS_TYPE=$(uname -s | tr '[:upper:]' '[:lower:]')

if [[ "$OS_TYPE" == "darwin" || "$OS_TYPE" == "linux-gnu" ]]; then
  echo "🚀 Installing ptk binaries to /usr/local/bin/ (requires write permissions)..."

  if [[ "$OS_TYPE" == "darwin" ]]; then
    cp "$OUTPUT_DIR/ptk-mac-intel" /usr/local/bin/ptk || echo "⚠️ Could not copy Mac Intel binary (permission denied?)"
    cp "$OUTPUT_DIR/ptk-mac-arm64" /usr/local/bin/ptk-arm64 || echo "⚠️ Could not copy Mac ARM64 binary (permission denied?)"
  elif [[ "$OS_TYPE" == "linux-gnu" ]]; then
    cp "$OUTPUT_DIR/ptk-linux" /usr/local/bin/ptk || echo "⚠️ Could not copy Linux binary (permission denied?)"
  fi

  echo "✅ Installation complete. You might need to rerun this script with sudo if permission was denied."
else
  echo "⚠️ Windows detected: please create a proper installer (e.g., Inno Setup) to install ptk-windows.exe."
fi

echo "🎉 All done!"
