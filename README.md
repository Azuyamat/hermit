# Hermit

A lightweight, cross-platform shell that bridges the gap between Unix and Windows environments.

## Why Hermit?

Write once, run anywhere. Hermit provides a unified shell experience across Windows, macOS, and Linux by implementing familiar Unix-style commands as native builtins. Whether you're a developer juggling multiple operating systems or a team maintaining cross-platform scripts, Hermit eliminates the friction of platform-specific command variations.

### Built for Speed

Traditional shells spawn external processes for even simple commands like `ls` or `cat`. Every invocation means process creation overhead, disk I/O to locate the binary, and context switching. Hermit's builtin commands execute directly within the shell process, eliminating this overhead entirely. The result? Commands that run instantly, scripts that fly, and a responsive terminal experience that feels snappier than traditional shells.

This performance advantage compounds dramatically in scripts that call commands repeatedly. A loop processing hundreds of files with external `cat` and `grep` commands might take seconds or minutes; the same operations with Hermit's builtins complete in milliseconds.

### The "rc" Mystery

You might notice Hermit uses `~/.hermitrc` for configuration. Ever wonder why shell config files end in "rc"? It stands for "run commands"—a naming convention inherited from the 1965 CTSS operating system's `runcom` facility. When shells like Unix's `/bin/sh` adopted this pattern, they abbreviated it to "rc". So when you edit `.bashrc`, `.zshrc`, or `.hermitrc`, you're literally editing a list of commands that run at shell startup, carrying forward a 60-year-old tradition.

## Features

- **Cross-platform builtins**: Common Unix commands work identically on Windows, macOS, and Linux
- **Fast execution**: Native builtin commands with zero process spawning overhead
- **Shell essentials**: Pipes, I/O redirection, logical operators (`&&`, `||`), and command chaining
- **Developer-friendly**: Built with scripting and automation in mind

## Getting Started

See [BUILTINS.md](BUILTINS.md) for the complete list of implemented and planned commands.

## Roadmap

This project is currently in its early stages and is maintained by a single developer. The roadmap is subject to change based on user feedback and development progress. Planned features include:

- [X] Basic shell functionality (command execution, piping, redirection)
- [ ] Environment variable persistence
- [ ] Command history with navigation
- [ ] Tab completion
- [ ] Script execution (.sh files)
- [ ] Configuration file (~/.hermitrc)
- [ ] Custom prompt configuration
- [ ] Job control (background processes)
- [ ] Signal handling

## License

[MIT License](LICENSE)
