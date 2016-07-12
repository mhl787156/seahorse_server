package conf

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

const siteRoot string = "../../../../../assignments/apex_name_subject_to_change"
const port int = 80
const dbipaddress string = "localhost:27017"

type ServerConfig struct {
	Port        int
	HtmlDir     string
	DBIPAddress string
}

func GetConfig() ServerConfig {
	var conf ServerConfig
	if _, err := toml.DecodeFile("./conf.toml", &conf); err != nil {
		fmt.Println("Config file could not be read, continuing with defaults...", err)
		conf.Port = port
		conf.HtmlDir = siteRoot
		conf.DBIPAddress = dbipaddress
	}
	return conf
}
