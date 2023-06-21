# Upgrade Script V1

A command-line tool written in Go for upgrading the system and Flatpak packages on Arch Linux.

## Usage

The tool can be executed by running the compiled binary file or by running the Go code directly.

The tool takes optional command-line arguments that will be passed to the underlying package manager commands.

The following command-line arguments are recognized:

- `noconfirm`: bypass confirmation prompts.

The tool will first upgrade the system using the `yay` package manager, and then upgrade the Flatpak packages using the `flatpak` command-line tool.

## Example

```bash
# Upgrade the system and Flatpak packages with confirmation prompts
./upgrade

# Upgrade the system and Flatpak packages without confirmation prompts
./upgrade noconfirm
```
