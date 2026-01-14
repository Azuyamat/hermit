# Hermit

Hermit is a lightweight shell designed to be cross-platform and easy to use while providing quality of life features for developers.

## Commands

Hermit aims to provide a cross-platform implementation of common Unix-style commands. Below is the current implementation status:

### Core Shell Builtins
- [X] `cd` - Change directory
- [X] `pwd` - Print working directory
- [X] `exit` - Exit the shell
- [X] `export` - Set environment variables
- [X] `echo` - Display text
- [X] `clear` - Clear the terminal screen
- [ ] `set` - Set shell variables (non-exported)
- [ ] `unset` - Unset variables
- [ ] `alias` - Create command aliases
- [ ] `unalias` - Remove aliases
- [ ] `history` - Command history
- [ ] `source` / `.` - Execute commands from file

### File Operations
- [X] `cat` - Concatenate and display files
- [X] `ls` - List directory contents
- [ ] `cp` - Copy files/directories
- [ ] `mv` - Move/rename files
- [ ] `rm` - Remove files/directories
- [ ] `mkdir` - Create directories
- [ ] `rmdir` - Remove empty directories
- [ ] `touch` - Create empty file or update timestamp
- [ ] `ln` - Create links
- [ ] `chmod` - Change file permissions (Unix-style, no-op on Windows)
- [ ] `chown` - Change file owner (Unix-style, no-op on Windows)

### Text Processing
- [ ] `grep` - Search text patterns
- [ ] `sed` - Stream editor
- [ ] `awk` - Pattern scanning and processing
- [ ] `cut` - Remove sections from lines
- [ ] `sort` - Sort lines
- [ ] `uniq` - Report or omit repeated lines
- [X] `wc` - Word, line, character count
- [ ] `head` - Output first part of files
- [ ] `tail` - Output last part of files
- [ ] `tr` - Translate characters
- [ ] `diff` - Compare files line by line

### Process Management
- [ ] `ps` - Process status
- [ ] `kill` - Send signal to process
- [ ] `jobs` - List background jobs
- [ ] `fg` - Bring job to foreground
- [ ] `bg` - Continue job in background
- [ ] `wait` - Wait for process completion
- [ ] `sleep` - Delay for specified time
- [ ] `time` - Time command execution

### System Information
- [ ] `env` - Display environment variables
- [ ] `printenv` - Print environment variables
- [ ] `which` - Locate command
- [ ] `whereis` - Locate binary/source/man page
- [ ] `whoami` - Print current user
- [ ] `hostname` - Show/set system hostname
- [ ] `uname` - Print system information
- [ ] `date` - Display or set date and time
- [ ] `uptime` - Show system uptime

### Network Utilities
- [ ] `ping` - Test network connectivity
- [ ] `curl` - Transfer data from URLs
- [ ] `wget` - Download files from web
- [ ] `ssh` - Secure shell client
- [ ] `scp` - Secure copy

### Archive & Compression
- [ ] `tar` - Archive utility
- [ ] `zip` - Package and compress files
- [ ] `unzip` - Extract compressed files
- [ ] `gzip` / `gunzip` - Compress/decompress files

### Miscellaneous
- [ ] `test` / `[` - Evaluate conditional expressions
- [ ] `expr` - Evaluate expressions
- [ ] `basename` - Strip directory and suffix from filenames
- [ ] `dirname` - Strip last component from file name
- [ ] `find` - Search for files in directory hierarchy
- [ ] `xargs` - Build and execute command lines
- [ ] `tee` - Read from stdin and write to stdout and files
- [ ] `yes` - Output a string repeatedly

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