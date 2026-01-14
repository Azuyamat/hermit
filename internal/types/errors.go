package types

import "fmt"

type ErrExitCode struct {
	ExitCode int
}

func (e *ErrExitCode) Error() string {
	return fmt.Sprintf("process exited with code %d", e.ExitCode)
}

func IsErrExitCode(err error) bool {
	_, ok := err.(*ErrExitCode)
	return ok
}

func NewErrExitCode(exitCode int) *ErrExitCode {
	return &ErrExitCode{ExitCode: exitCode}
}
