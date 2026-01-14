package command

import (
	"errors"
	"fmt"
	"io"

	"github.com/azuyamat/hermit/internal/types"
)

type Manager struct {
	commands map[string]Command
}

func NewManager() *Manager {
	return &Manager{
		commands: make(map[string]Command),
	}
}

func (m *Manager) Register(cmd Command) {
	meta := cmd.Metadata()
	m.commands[meta.Name] = cmd
}

func (m *Manager) RegisterAll(cmds ...Command) {
	for _, cmd := range cmds {
		m.Register(cmd)
	}
}

func (m *Manager) Get(name string) (Command, bool) {
	cmd, ok := m.commands[name]
	return cmd, ok
}

func (m *Manager) List() []Metadata {
	metas := make([]Metadata, 0, len(m.commands))
	for _, cmd := range m.commands {
		metas = append(metas, cmd.Metadata())
	}
	return metas
}

func (m *Manager) Execute(name string, args []string, stdout, sterr io.Writer, stdin io.Reader, shell *types.ExecutionContext) error {
	cmd, ok := m.Get(name)
	if !ok {
		return fmt.Errorf("command not found: %s", name)
	}

	ctx, err := Parse(cmd.Metadata(), args, stdout, sterr, stdin)
	if errors.Is(err, ErrShowHelp) {
		PrintUsage(cmd, sterr)
		return nil
	}
	if err != nil {
		return err
	}

	return cmd.Execute(ctx, shell)
}
