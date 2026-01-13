package builtins

import (
	"github.com/azuyamat/hermit/internal/builtins/core"
	"github.com/azuyamat/hermit/internal/builtins/file"
	"github.com/azuyamat/hermit/internal/types"
)

func NewCd() types.Builtin {
	return &core.Cd{}
}

func NewClear() types.Builtin {
	return &core.Clear{}
}

func NewEcho() types.Builtin {
	return &core.Echo{}
}

func NewExit() types.Builtin {
	return &core.Exit{}
}

func NewExport() types.Builtin {
	return &core.Export{}
}

func NewPwd() types.Builtin {
	return &core.Pwd{}
}

func NewCat() types.Builtin {
	return &file.Cat{}
}
