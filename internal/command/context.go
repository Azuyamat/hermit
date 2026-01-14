package command

import (
	"fmt"
	"io"
	"strconv"
)

type Context struct {
	args   []string
	flags  map[string]any
	stdout io.Writer
	stderr io.Writer
	stdin  io.Reader
}

func (c *Context) Args() []string {
	return c.args
}

func (c *Context) ArgCount() int {
	return len(c.args)
}

func (c *Context) Bool(name string) bool {
	val, ok := c.flags[name].(bool)
	return ok && val
}

func (c *Context) BoolOr(name string, fallback bool) bool {
	val, ok := c.flags[name].(bool)
	if !ok {
		return fallback
	}
	return val
}

func (c *Context) String(name string) string {
	val, ok := c.flags[name].(string)
	if !ok {
		return ""
	}
	return val
}

func (c *Context) Int(name string) int {
	val, ok := c.flags[name].(int)
	if !ok {
		return 0
	}
	return val
}

func (c *Context) Arg(index int) (string, bool) {
	if index < len(c.args) {
		return c.args[index], true
	}
	return "", false
}

func (c *Context) ArgOr(index int, fallback string) string {
	if index < len(c.args) {
		return c.args[index]
	}
	return fallback
}

func (c *Context) ArgIntOr(index int, fallback int) int {
	if index < len(c.args) {
		if val, err := strconv.Atoi(c.args[index]); err == nil {
			return val
		}
	}
	return fallback
}

func (c *Context) ArgBoolOr(index int, fallback bool) bool {
	if index < len(c.args) {
		if val, err := strconv.ParseBool(c.args[index]); err == nil {
			return val
		}
	}
	return fallback
}

func (c *Context) Stdout() io.Writer {
	return c.stdout
}

func (c *Context) Stderr() io.Writer {
	return c.stderr
}

func (c *Context) Stdin() io.Reader {
	return c.stdin
}

func (c *Context) Print(a ...interface{}) (n int, err error) {
	return fmt.Fprint(c.stdout, a...)
}

func (c *Context) Println(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(c.stdout, a...)
}

func (c *Context) Printf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(c.stdout, format, a...)
}

func (c *Context) Error(a ...interface{}) (n int, err error) {
	return fmt.Fprint(c.stderr, a...)
}

func (c *Context) Errorln(a ...interface{}) (n int, err error) {
	return fmt.Fprintln(c.stderr, a...)
}

func (c *Context) Errorf(format string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(c.stderr, format, a...)
}
