package properties

import (
	"github.com/bbcloudGroup/gothic/config"
	"gopkg.in/yaml.v3"
)

type Properties struct {
	Author		string	`yaml:"Author"`
	JwtSecret 	string 	`yaml:"JwtSecret"`
}

func (p *Properties) Register(properties *map[string]interface{}) {
	(*properties)["Author"] = p.Author
	(*properties)["JwtSecret"] = p.JwtSecret
}


func BootStrap(env string) map[string]interface{} {

	properties := make(map[string]interface{})

	c := struct {
		Properties Properties `yaml:"Properties"`
	}{}

	data, err := config.ReadYAML(env)
	if err == nil {
		if err := yaml.Unmarshal(data, &c); err == nil {
			c.Properties.Register(&properties)
		}
	}
	return properties
}