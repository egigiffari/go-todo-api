package gogin

type Config struct {
	Port  string `yaml:"port"`
	Debug bool   `yaml:"debug"`
	Env   string `yaml:"env"`
}
