package options

type TransferOption struct {
	Listen string     `yaml:"listen"`
	Pipe   PipeOption `yaml:"pipe"`
}
