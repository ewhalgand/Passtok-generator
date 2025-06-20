# ptk - Password & Token Generator

**ptk** is a simple and secure command-line tool to generate passwords and tokens, built with Go and Cobra.

## ✨ Features

- Generate strong, random passwords or API tokens
- Control the output length and type
- Optional clipboard copy
- Cross-platform support (macOS, Linux, Windows)
- Lightweight and fast

## 📦 Installation

### 🔽 Download precompiled binaries

Download the latest `.zip` archive for your platform from the [Releases](https://github.com/ewhalgand/Passtok-generator/releases) page:

- [⬇️ macOS Intel](https://github.com/ewhalgand/Passtok-generator/releases/download/v1.0.0/ptk-mac-intel-1.0.0.zip)
- [⬇️ macOS Apple Silicon](https://github.com/ewhalgand/Passtok-generator/releases/download/v1.0.0/ptk-mac-arm64-1.0.0.zip)
- [⬇️ Linux](https://github.com/ewhalgand/Passtok-generator/releases/download/v1.0.0/ptk-linux-1.0.0.zip)
- [⬇️ Windows](https://github.com/ewhalgand/Passtok-generator/releases/download/v1.0.0/ptk-windows.exe-1.0.0.zip)

Unzip the archive and place the binary in your `$PATH`.

## ⚙️ Usage

Here are the main commands available:

```bash
# Generate a password (default length, default type)
ptk gen

# Specify the type of string to generate (password or token)
ptk gen --type=password
ptk gen --type=token

# Specify the length of the generated string (e.g., 32 characters)
ptk gen --length=32

# Copy the generated result to the clipboard
ptk gen --copy

# Combine all options
ptk gen --type=token --length=64 --copy
```
