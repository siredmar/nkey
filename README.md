# NKEY

[![Keep a Changelog](https://img.shields.io/badge/changelog-Keep%20a%20Changelog-%23E05735)](CHANGELOG.md)
[![GitHub Release](https://img.shields.io/github/v/release/siredmar/nkey)](https://github.com/siredmar/nkey/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/siredmar/nkey.svg)](https://pkg.go.dev/github.com/siredmar/nkey)
[![go.mod](https://img.shields.io/github/go-mod/go-version/siredmar/nkey)](go.mod)
[![LICENSE](https://img.shields.io/github/license/siredmar/nkey)](LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/siredmar/nkey/build.yml?branch=main)](https://github.com/siredmar/nkey/actions?query=workflow%3Abuild+branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/siredmar/nkey)](https://goreportcard.com/report/github.com/siredmar/nkey)
[![Codecov](https://codecov.io/gh/siredmar/nkey/branch/main/graph/badge.svg)](https://codecov.io/gh/siredmar/nkey)

⭐ `Star` this repository if you find it valuable and worth maintaining.

👁 `Watch` this repository to get notified about new releases, issues, etc.

## Description

This is a minimalist tool to convert NATS nkeys to a private/public keypair.

## Usage

```bash
$ nkey -s <seed>
# example
$ nkey -s SAAOW7V3L5YRJGBYY3ZTXB5WHYU4W3HFQ44MBYIBWAECRMD6RMZJTS426E
Private key: PDVX5O27OEKJQOGG6M5YPNR6FHFWZZMHHDAOCANQBAULA7ULGKM4XXDL2O2N6GNXC7SYZDSCXEWR5NN4SDECHPU77HZX6K7JCJJV5EQFFIZQ
Public key: ADOGXU5U34M3OF7FRSHEFOJND223ZEGIEO7J76PTP4V6SESTL2JALU4U
```
## Setup

```bash
go install github.com/siredmar/nkey@latest
```
## Build

### Terminal

- `make` - execute the build pipeline.
- `make help` - print help for the [Make targets](Makefile).

### Visual Studio Code

`F1` → `Tasks: Run Build Task (Ctrl+Shift+B or ⇧⌘B)` to execute the build pipeline.

## Release

The release workflow is triggered each time a tag with `v` prefix is pushed.

_CAUTION_: Make sure to understand the consequences before you bump the major version.
More info: [Go Wiki](https://github.com/golang/go/wiki/Modules#releasing-modules-v2-or-higher),
[Go Blog](https://blog.golang.org/v2-go-modules).

## Maintenance

Notable files:

- [.github/workflows](.github/workflows) - GitHub Actions workflows,
- [.github/dependabot.yml](.github/dependabot.yml) - Dependabot configuration,
- [.vscode](.vscode) - Visual Studio Code configuration files,
- [.golangci.yml](.golangci.yml) - golangci-lint configuration,
- [.goreleaser.yml](.goreleaser.yml) - GoReleaser configuration,
- [Dockerfile](Dockerfile) - Dockerfile used by GoReleaser to create a container image,
- [Makefile](Makefile) - Make targets used for development, [CI build](.github/workflows) and [.vscode/tasks.json](.vscode/tasks.json),
- [go.mod](go.mod) - [Go module definition](https://github.com/golang/go/wiki/Modules#gomod),
- [tools.go](tools.go) - [build tools](https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module).

## FAQ

### Why Visual Studio Code editor configuration

Developers that use Visual Studio Code can take advantage of the editor configuration.
While others do not have to care about it.
Setting configs for each repo is unnecessary time consuming.
VS Code is the most popular Go editor ([survey](https://blog.golang.org/survey2019-results))
and it is officially [supported by the Go team](https://blog.golang.org/vscode-go).

You can always remove the [.vscode](.vscode) directory if it really does not help you.

### Why GitHub Actions, not any other CI server

GitHub Actions is out-of-the-box if you are already using GitHub.
[Here](https://github.com/mvdan/github-actions-golang) you can learn how to use it for Go.

However, changing to any other CI server should be very simple,
because this repository has build logic and tooling installation in [Makefile](Makefile).

### How can I build on Windows

Install [tdm-gcc](https://jmeubank.github.io/tdm-gcc/)
and copy `C:\TDM-GCC-64\bin\mingw32-make.exe`
to `C:\TDM-GCC-64\bin\make.exe`.
Alternatively, you may install [mingw-w64](http://mingw-w64.org/doku.php)
and copy `mingw32-make.exe` accordingly.

Take a look [here](https://github.com/docker-archive/toolbox/issues/673#issuecomment-355275054),
if you have problems using Docker in Git Bash.

You can also use [WSL (Windows Subsystem for Linux)](https://docs.microsoft.com/en-us/windows/wsl/install-win10)
or develop inside a [Remote Container](https://code.visualstudio.com/docs/remote/containers).
However, take into consideration that then you are not going to use "bare-metal" Windows.

Consider using [goyek](https://github.com/goyek/goyek)
for creating cross-platform build pipelines in Go.

### How can I customize the release

Take a look at GoReleaser [docs](https://goreleaser.com/customization/)
as well as [its repo](https://github.com/goreleaser/goreleaser/)
how it is dogfooding its functionality.
You can use it to add deb/rpm/snap packages, Homebrew Tap, Scoop App Manifest etc.

If you are developing a library and you like handcrafted changelog and release notes,
you are free to remove any usage of GoReleaser.

## Contributing

Feel free to create an issue or propose a pull request.

Follow the [Code of Conduct](CODE_OF_CONDUCT.md).
