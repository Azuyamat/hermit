package builtins

import (
	"github.com/azuyamat/hermit/internal/builtins/core"
	"github.com/azuyamat/hermit/internal/builtins/file"
	"github.com/azuyamat/hermit/internal/command"
)

func RegisterCoreBuiltins(manager *command.Manager) {
	manager.RegisterAll(
		&core.Cd{},
		&core.Clear{},
		&core.Echo{},
		&core.Exit{},
		&core.Export{},
		&core.Pwd{},
		&True{},
		&False{},
		&file.Cat{},
		&file.Wc{},
		&file.Ls{},
	)
}
