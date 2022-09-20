package core

var ConfigYml = Config{}
var lock = make(chan struct{})

type Config struct {
	Port    int     `yaml:"port" json:"port"`
	Node    int     `yaml:"node" json:"node"`
	Channel Channel `yaml:"channel" json:"channel"`
}

type Channel struct {
	Type  string `yaml:"type" json:"type"`
	Url   string `yaml:"url" json:"url"`
	Topic string `yaml:"topic"`
}

func (Config) Done() {
	close(lock)
}

func (Config) Wait() {
	<-lock
}
