package generator

type Config interface {
	Validate() error
	Configure() error
	RootPath() string
	TemplateConfig() interface{}
	FileSets() []string
	CommandSets() []string
}
