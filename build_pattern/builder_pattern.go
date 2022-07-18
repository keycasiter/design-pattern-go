package build_pattern

type ConfigBuilder struct {
	host   string
	port   int
	maxNum int
}

func NewConfigBuilder() *ConfigBuilder {
	return &ConfigBuilder{}
}

func (c *ConfigBuilder) WithHost(host string) {
	c.host = host
}

func (c *ConfigBuilder) WithPort(port int) {
	c.port = port
}

func (c *ConfigBuilder) WithMaxNum(maxNum int) {
	c.maxNum = maxNum
}

func (c *ConfigBuilder) Build() *ConfigBuilder {
	return c
}
