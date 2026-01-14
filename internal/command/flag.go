package command

type FlagType int

const (
	Bool FlagType = iota
	Str
	Int
)

type FlagDef struct {
	Name        string
	Short       string
	Description string
	Type        FlagType
	Default     any
}

var (
	RecursiveFlag = FlagDef{"recursive", "r", "Recursively apply the command", Bool, false}
	ForceFlag     = FlagDef{"force", "f", "Force the command to run", Bool, false}
	VerboseFlag   = FlagDef{"verbose", "v", "Enable verbose output", Bool, false}
	AllFlag       = FlagDef{"all", "a", "Apply the command to all items", Bool, false}
	NoFlags       = []FlagDef{}
)

type FlagDefBuilder struct {
	flag FlagDef
}

func NewFlagDef(
	name string,
	short string,
	description string,
) *FlagDefBuilder {
	return &FlagDefBuilder{
		flag: FlagDef{
			Name:        name,
			Short:       short,
			Description: description,
		},
	}
}

func (b *FlagDefBuilder) Type(flagType FlagType) *FlagDefBuilder {
	b.flag.Type = flagType
	return b
}

func (b *FlagDefBuilder) Default(defaultValue any) *FlagDefBuilder {
	b.flag.Default = defaultValue
	return b
}

func (b *FlagDefBuilder) Build() FlagDef {
	return b.flag
}

func NewBoolFlag(
	name string,
	short string,
	description string,
) *FlagDefBuilder {
	return NewFlagDef(name, short, description).
		Type(Bool).
		Default(false)
}

func NewStrFlag(
	name string,
	short string,
	description string,
	defaultValue string,
) *FlagDefBuilder {
	return NewFlagDef(name, short, description).
		Type(Str).
		Default(defaultValue)
}

func NewIntFlag(
	name string,
	short string,
	description string,
	defaultValue int,
) *FlagDefBuilder {
	return NewFlagDef(name, short, description).
		Type(Int).
		Default(defaultValue)
}
