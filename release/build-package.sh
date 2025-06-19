#!/bin/bash

set -e

VERSION="1.0.0"
OUTPUT_DIR="./build"
SRC_DIR="/Users/ewenhalgand/Desktop/Fichier-Dev/DevPerso/Golang/passtok-generator"

echo "🚀 Starting cross-platform compilation of ptk version $VERSION..."

mkdir -p "$OUTPUT_DIR"

# Compilation multi-plateforme
echo "🖥️ Building for Mac Intel (amd64)..."
GOOS=darwin GOARCH=amd64 go build -o "$OUTPUT_DIR/ptk-mac-intel" "$SRC_DIR"

echo "🍏 Building for Mac ARM64 (Apple Silicon)..."
GOOS=darwin GOARCH=arm64 go build -o "$OUTPUT_DIR/ptk-mac-arm64" "$SRC_DIR"

echo "🐧 Building for Linux amd64..."
GOOS=linux GOARCH=amd64 go build -o "$OUTPUT_DIR/ptk-linux" "$SRC_DIR"

echo "🪟 Building for Windows amd64..."
GOOS=windows GOARCH=amd64 go build -o "$OUTPUT_DIR/ptk-windows.exe" "$SRC_DIR"

echo "✅ Compilation complete."

# Packaging basique (zip archives)
echo "📦 Packaging binaries..."

zip -j "$OUTPUT_DIR/ptk-mac-intel-$VERSION.zip" "$OUTPUT_DIR/ptk-mac-intel"
zip -j "$OUTPUT_DIR/ptk-mac-arm64-$VERSION.zip" "$OUTPUT_DIR/ptk-mac-arm64"
zip -j "$OUTPUT_DIR/ptk-linux-$VERSION.zip" "$OUTPUT_DIR/ptk-linux"
zip -j "$OUTPUT_DIR/ptk-windows-$VERSION.zip" "$OUTPUT_DIR/ptk-windows.exe"

echo "✅ Packaging complete. Files saved in $OUTPUT_DIR"
