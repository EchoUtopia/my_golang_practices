package configs

import (
	"github.com/EchoUtopia/zerror"
	"log"
)

type Config struct {
	Name        string
	Category    string
	Switch      string // on or off
	Description string
	ContentType string // default json
	Content     string
}

type Configurer interface {
	CheckAndSet(ctx interface{}, input *Config) error
	UnMarshalContent(content interface{}, input *Config) error
	GetName() string
}

var configsMap = map[string]Configurer{}

func RegisterConfigurer(config Configurer) {
	name := config.GetName()
	if _, ok := configsMap[name]; ok {
		log.Panicf(`config: %s already exists`, name)
	}
	configsMap[config.GetName()] = config
}

func getConfigurer(name string) (Configurer, bool) {
	config, ok := configsMap[name]
	return config, ok
}

func CheckAndSet(ctx interface{}, config *Config) error {
	configurer, ok := getConfigurer(config.Name)
	if !ok {
		return zerror.BadRequest.Errorf(`config: %s not exists`, config.Name)
	}
	return configurer.CheckAndSet(ctx, config)
}
