package executor

import "github.com/azuyamat/hermit/internal/types"

type BuiltinRegistry struct {
	builtins map[string]types.Builtin
}

func NewBuiltinRegistry() *BuiltinRegistry {
	return &BuiltinRegistry{
		builtins: make(map[string]types.Builtin),
	}
}

func (r *BuiltinRegistry) Register(b types.Builtin) {
	r.builtins[b.Name()] = b
}

func (r *BuiltinRegistry) Get(name string) (types.Builtin, bool) {
	b, ok := r.builtins[name]
	return b, ok
}

func (r *BuiltinRegistry) IsBuiltin(name string) bool {
	_, ok := r.builtins[name]
	return ok
}
