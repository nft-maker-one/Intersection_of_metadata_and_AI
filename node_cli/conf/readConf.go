package conf

import (
	"os"

	"gopkg.in/yaml.v3"
)

type ApiConf struct {
	Platform string `yaml:"platform" json:"platform"`
	ApiKey   string `yaml:"apiKey" json:"apiKey"`
	Version  string `yaml:"version" json:"version"`
	ProxyUrl string `yaml:"proxyUrl" json:"proxyUrl"`
}

func InitConf(path string) *ApiConf {
	apiConf := &ApiConf{}
	buf, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(buf, apiConf)
	if err != nil {
		panic(err)
	}
	return apiConf
}
