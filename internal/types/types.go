package types

import (
	"context"
	"io"
	"os"
	"strings"
)

type Builtin interface {
	Name() string
	Execute(args []string, context *ExecutionContext, stdin io.Reader, stdout io.Writer, stderr io.Writer) error
}

type ExecutionContext struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer

	Env          map[string]string
	Context      context.Context
	LastExitCode int
	Variables    map[string]string
}

func NewContext() *ExecutionContext {
	env := make(map[string]string)
	for _, e := range os.Environ() {
		parts := strings.SplitN(e, "=", 2)
		if len(parts) == 2 {
			env[parts[0]] = parts[1]
		}
	}

	return &ExecutionContext{
		Stdin:        os.Stdin,
		Stdout:       os.Stdout,
		Stderr:       os.Stderr,
		Env:          env,
		Context:      context.Background(),
		LastExitCode: 0,
		Variables:    make(map[string]string),
	}
}

func (ec *ExecutionContext) SetEnv(key, value string) {
	ec.Env[key] = value
}

func (ec *ExecutionContext) GetEnv(key string) string {
	return ec.Env[key]
}

func (ec *ExecutionContext) EnvSlice() []string {
	env := make([]string, 0, len(ec.Env))
	for k, v := range ec.Env {
		env = append(env, k+"="+v)
	}
	return env
}

func (ec *ExecutionContext) GetWorkingDir() (string, error) {
	return os.Getwd()
}

func (ec *ExecutionContext) SetWorkingDir(dir string) error {
	return os.Chdir(dir)
}
