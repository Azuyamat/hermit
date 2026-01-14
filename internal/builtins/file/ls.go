package file

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/azuyamat/hermit/internal/command"
	"github.com/azuyamat/hermit/internal/types"
)

type Ls struct{}

func (l *Ls) Metadata() command.Metadata {
	return command.NewMetadataBuilder("ls", "List directory contents").
		Usage("ls [options] [path ...]").
		Flags(
			command.NewBoolFlag("all", "a", "Do not ignore entries starting with .").Build(),
			command.NewBoolFlag("long", "l", "Use a long listing format").Build(),
			command.NewBoolFlag("human-readable", "h", "Print sizes in human readable format").Build(),
			command.NewBoolFlag("recursive", "R", "List subdirectories recursively").Build(),
			command.NewBoolFlag("directory", "d", "List directories themselves, not their contents").Build(),
			command.NewBoolFlag("almost-all", "A", "Do not list implied . and ..").Build(),
		).
		MinArgs(0).
		MaxArgs(-1).
		Build()
}

func (l *Ls) Execute(ctx *command.Context, shell *types.ExecutionContext) error {
	paths := ctx.Args()
	if len(paths) == 0 {
		paths = []string{"."}
	}

	if ctx.Bool("directory") {
		return l.listPaths(ctx, paths)
	}

	var files, dirs []string
	for _, path := range paths {
		if info, err := os.Stat(path); err != nil {
			fmt.Fprintf(ctx.Stderr(), "ls: cannot access '%s': %v\n", path, err)
		} else if info.IsDir() {
			dirs = append(dirs, path)
		} else {
			files = append(files, path)
		}
	}

	if len(files) > 0 {
		l.listPaths(ctx, files)
	}

	for i, dir := range dirs {
		if len(paths) > 1 || len(files) > 0 {
			if i > 0 || len(files) > 0 {
				fmt.Fprintln(ctx.Stdout())
			}
			fmt.Fprintf(ctx.Stdout(), "%s:\n", dir)
		}

		if ctx.Bool("recursive") {
			l.listRecursive(ctx, dir)
		} else {
			l.listDirectory(ctx, dir)
		}
	}

	return nil
}

func (l *Ls) listPaths(ctx *command.Context, paths []string) error {
	if ctx.Bool("long") {
		for _, path := range paths {
			if info, err := os.Stat(path); err == nil {
				l.printLong(ctx, path, info)
			}
		}
	} else {
		for _, path := range paths {
			fmt.Fprintln(ctx.Stdout(), filepath.Base(path))
		}
	}
	return nil
}

func (l *Ls) listDirectory(ctx *command.Context, dir string) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}

	filtered := l.filter(ctx, entries)
	sort.Slice(filtered, func(i, j int) bool { return filtered[i].Name() < filtered[j].Name() })

	if ctx.Bool("long") {
		for _, entry := range filtered {
			if info, err := entry.Info(); err == nil {
				l.printLong(ctx, filepath.Join(dir, entry.Name()), info)
			}
		}
	} else {
		for _, entry := range filtered {
			fmt.Fprintln(ctx.Stdout(), entry.Name())
		}
	}

	return nil
}

func (l *Ls) listRecursive(ctx *command.Context, root string) error {
	l.listDirectory(ctx, root)

	entries, err := os.ReadDir(root)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if !entry.IsDir() || !l.shouldShow(ctx, entry.Name()) || entry.Name() == "." || entry.Name() == ".." {
			continue
		}

		subdir := filepath.Join(root, entry.Name())
		fmt.Fprintf(ctx.Stdout(), "\n%s:\n", subdir)
		l.listRecursive(ctx, subdir)
	}

	return nil
}

func (l *Ls) filter(ctx *command.Context, entries []fs.DirEntry) []fs.DirEntry {
	var filtered []fs.DirEntry
	for _, entry := range entries {
		if l.shouldShow(ctx, entry.Name()) {
			filtered = append(filtered, entry)
		}
	}
	return filtered
}

func (l *Ls) shouldShow(ctx *command.Context, name string) bool {
	if ctx.Bool("all") {
		return true
	}
	if !strings.HasPrefix(name, ".") {
		return true
	}
	return ctx.Bool("almost-all") && name != "." && name != ".."
}

func (l *Ls) printLong(ctx *command.Context, path string, info fs.FileInfo) {
	size := fmt.Sprintf("%d", info.Size())
	if ctx.Bool("human-readable") {
		size = l.humanSize(info.Size())
	}

	fmt.Fprintf(ctx.Stdout(), "%s %8s %s %s\n",
		l.perms(info.Mode()), size, l.formatTime(info.ModTime()), filepath.Base(path))
}

func (l *Ls) perms(mode fs.FileMode) string {
	p := [10]rune{'-', '-', '-', '-', '-', '-', '-', '-', '-', '-'}
	if mode.IsDir() {
		p[0] = 'd'
	} else if mode&fs.ModeSymlink != 0 {
		p[0] = 'l'
	}

	masks := []fs.FileMode{0400, 0200, 0100, 0040, 0020, 0010, 0004, 0002, 0001}
	chars := []rune{'r', 'w', 'x', 'r', 'w', 'x', 'r', 'w', 'x'}

	for i, mask := range masks {
		if mode&mask != 0 {
			p[i+1] = chars[i]
		}
	}

	return string(p[:])
}

func (l *Ls) formatTime(t time.Time) string {
	if t.After(time.Now().AddDate(0, -6, 0)) && t.Before(time.Now().AddDate(0, 0, 1)) {
		return t.Format("Jan _2 15:04")
	}
	return t.Format("Jan _2  2006")
}

func (l *Ls) humanSize(size int64) string {
	if size < 1024 {
		return fmt.Sprintf("%dB", size)
	}

	div, exp := int64(1024), 0
	for n := size / 1024; n >= 1024; n /= 1024 {
		div *= 1024
		exp++
	}

	return fmt.Sprintf("%.1f%s", float64(size)/float64(div), []string{"K", "M", "G", "T", "P", "E"}[exp])
}
