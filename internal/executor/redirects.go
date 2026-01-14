package executor

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/azuyamat/hermit/internal/ast"
)

type redirectHandler struct {
	stdin  io.Reader
	stdout io.Writer
	stderr io.Writer
	files  []io.Closer
}

func (e *Executor) setupRedirects(redirects []ast.Redirect) (
	stdin io.Reader,
	stdout io.Writer,
	stderr io.Writer,
	cleanup func(),
	err error,
) {
	handler := &redirectHandler{}
	cleanup = func() { closeAll(handler.files) }

	for _, redirect := range redirects {
		if err := e.applyRedirect(redirect, handler); err != nil {
			cleanup()
			return nil, nil, nil, nil, err
		}
	}

	return handler.stdin, handler.stdout, handler.stderr, cleanup, nil
}

func (e *Executor) applyRedirect(redirect ast.Redirect, handler *redirectHandler) error {
	target, err := e.getRedirectTarget(redirect.Target)
	if err != nil {
		return err
	}

	switch redirect.Type {
	case ast.RedirectStdout:
		return e.redirectStdout(target, handler, false)
	case ast.RedirectAppend:
		return e.redirectStdout(target, handler, true)
	case ast.RedirectStderr:
		return e.redirectStderr(target, handler, false)
	case ast.RedirectStderrAppend:
		return e.redirectStderr(target, handler, true)
	case ast.RedirectStdin:
		return e.redirectStdin(target, handler)
	case ast.RedirectBoth:
		return e.redirectBoth(target, handler)
	default:
		return fmt.Errorf("unknown redirect type: %v", redirect.Type)
	}
}

func (e *Executor) getRedirectTarget(target ast.Argument) (string, error) {
	targets, err := e.evaluateArgument(target)
	if err != nil {
		return "", fmt.Errorf("evaluating redirect target: %w", err)
	}

	if len(targets) != 1 {
		return "", errors.New("redirect target must expand to exactly one value")
	}

	return e.expandPath(targets[0]), nil
}

func (e *Executor) redirectStdout(target string, handler *redirectHandler, appendMode bool) error {
	flags := os.O_CREATE | os.O_WRONLY
	if appendMode {
		flags |= os.O_APPEND
	} else {
		flags |= os.O_TRUNC
	}

	f, err := os.OpenFile(target, flags, 0644)
	if err != nil {
		return fmt.Errorf("opening stdout file: %w", err)
	}

	handler.files = append(handler.files, f)
	handler.stdout = f
	return nil
}

func (e *Executor) redirectStderr(target string, handler *redirectHandler, appendMode bool) error {
	flags := os.O_CREATE | os.O_WRONLY
	if appendMode {
		flags |= os.O_APPEND
	} else {
		flags |= os.O_TRUNC
	}

	f, err := os.OpenFile(target, flags, 0644)
	if err != nil {
		return fmt.Errorf("opening stderr file: %w", err)
	}

	handler.files = append(handler.files, f)
	handler.stderr = f
	return nil
}

func (e *Executor) redirectStdin(target string, handler *redirectHandler) error {
	f, err := os.Open(target)
	if err != nil {
		return fmt.Errorf("opening stdin file: %w", err)
	}

	handler.files = append(handler.files, f)
	handler.stdin = f
	return nil
}

func (e *Executor) redirectBoth(target string, handler *redirectHandler) error {
	f, err := os.Create(target)
	if err != nil {
		return fmt.Errorf("creating output file: %w", err)
	}

	handler.files = append(handler.files, f)
	handler.stdout = f
	handler.stderr = f
	return nil
}

func (e *Executor) expandPath(path string) string {
	if !strings.HasPrefix(path, "~/") {
		return path
	}

	home, ok := e.context.GetEnv("HOME")
	if !ok {
		return path
	}

	return filepath.Join(home, path[2:])
}
