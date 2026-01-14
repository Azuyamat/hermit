package executor

import (
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sync"

	"github.com/azuyamat/hermit/internal/ast"
)

func (e *Executor) executePipeline(pipeline *ast.Pipeline) error {
	if len(pipeline.Commands) == 0 {
		return errors.New("empty pipeline")
	}
	if len(pipeline.Commands) == 1 {
		return e.executeCommand(pipeline.Commands[0])
	}
	return e.executeMultiCommandPipeline(pipeline.Commands)
}

func (e *Executor) executeMultiCommandPipeline(commands []*ast.Command) error {
	var wg sync.WaitGroup
	errChan := make(chan error, len(commands))

	var prevReader io.ReadCloser

	for i := 0; i < len(commands); i++ {
		var stdin io.Reader = e.context.Stdin
		var stdout io.WriteCloser

		if i > 0 {
			stdin = prevReader
		}

		if i < len(commands)-1 {
			r, w := io.Pipe()
			prevReader = r
			stdout = w
		} else {
			stdout = nopWriteCloser{e.context.Stdout}
		}

		wg.Add(1)
		go func(cmd *ast.Command, idx int, stdin io.Reader, stdout io.WriteCloser) {
			defer wg.Done()
			defer stdout.Close()
			defer drainReader(stdin)

			args, err := e.evaluateArguments(cmd.Args)
			if err != nil {
				errChan <- fmt.Errorf("command %d: %w", idx, err)
				return
			}

			customStdin, customStdout, stderr, cleanup, err := e.setupRedirects(cmd.Redirects)
			if err != nil {
				errChan <- fmt.Errorf("command %d redirects: %w", idx, err)
				return
			}
			defer cleanup()

			stdio := &stdio{
				stdin:  coalescReader(customStdin, stdin),
				stdout: coalescWriter(customStdout, stdout),
				stderr: coalescWriter(stderr, e.context.Stderr),
			}

			err = e.runCommand(cmd.Name, args, stdio)
			if err != nil {
				errChan <- fmt.Errorf("command %d (%q): %w", idx, cmd.Name, err)
			}
		}(commands[i], i, stdin, stdout)
	}

	wg.Wait()
	close(errChan)

	var firstErr error
	for err := range errChan {
		if err != nil && firstErr == nil {
			firstErr = err
		}
	}

	if firstErr != nil {
		var exitErr *exec.ExitError
		if errors.As(firstErr, &exitErr) {
			e.context.LastExitCode = exitErr.ExitCode()
		} else {
			e.context.LastExitCode = 1
		}
		return firstErr
	}

	e.context.LastExitCode = 0
	return nil
}

func drainReader(r io.Reader) {
	if r != nil && r != os.Stdin {
		io.Copy(io.Discard, r)
	}
}

type nopWriteCloser struct {
	io.Writer
}

func (nopWriteCloser) Close() error { return nil }
