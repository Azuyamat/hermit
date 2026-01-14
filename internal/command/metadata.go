package command

type Metadata struct {
	Name        string
	Description string
	Usage       string
	Flags       []FlagDef
	MinArgs     int
	MaxArgs     int // -1 for unlimited
}

type MetadataBuilder struct {
	metadata Metadata
}

func NewMetadataBuilder(
	name string,
	description string,
) *MetadataBuilder {
	return &MetadataBuilder{
		metadata: Metadata{
			Name:        name,
			Description: description,
			Usage:       name,
			Flags:       []FlagDef{},
			MinArgs:     0,
			MaxArgs:     -1,
		},
	}
}

func (b *MetadataBuilder) Usage(usage string) *MetadataBuilder {
	b.metadata.Usage = usage
	return b
}

func (b *MetadataBuilder) Flags(flags ...FlagDef) *MetadataBuilder {
	b.metadata.Flags = flags
	return b
}

func (b *MetadataBuilder) Args(min, max int) *MetadataBuilder {
	b.metadata.MinArgs = min
	b.metadata.MaxArgs = max
	return b
}

func (b *MetadataBuilder) ExactArgs(count int) *MetadataBuilder {
	b.metadata.MinArgs = count
	b.metadata.MaxArgs = count
	return b
}

func (b *MetadataBuilder) MinArgs(min int) *MetadataBuilder {
	b.metadata.MinArgs = min
	return b
}

func (b *MetadataBuilder) MaxArgs(max int) *MetadataBuilder {
	b.metadata.MaxArgs = max
	return b
}

func (b *MetadataBuilder) Build() Metadata {
	return b.metadata
}
