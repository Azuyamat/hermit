package ast

import "github.com/azuyamat/hermit/internal/token"

type RedirectType int

const (
	RedirectStdout RedirectType = iota
	RedirectStderr
	RedirectAppend
	RedirectStderrAppend
	RedirectStdin
	RedirectBoth
)

type Redirect struct {
	Token  token.Token
	Type   RedirectType
	Target Argument
}

func (r *Redirect) String() string {
	var op string
	switch r.Type {
	case RedirectStdout:
		op = ">"
	case RedirectAppend:
		op = ">>"
	case RedirectStderr:
		op = "2>"
	case RedirectStderrAppend:
		op = "2>>"
	case RedirectStdin:
		op = "<"
	case RedirectBoth:
		op = "&>"
	}
	return op + " " + r.Target.String()
}
