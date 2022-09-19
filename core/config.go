package core

type Config struct {
	Port    int     `yaml:"port"`
	Node    int     `yaml:"node"`
	Channel Channel `yaml:"channel"`
}

type Channel struct {
	Type string `yaml:"type"`
	Url  string `yaml:"url"`
}
