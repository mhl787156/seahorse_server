package conf

import (
  "github.com/BurntSushi/toml"
  "fmt"
)

const siteRoot string = "../../../../../assignments/apex_name_subject_to_change"
const port int = 80

type ServerConfig struct {
  Port int
  HtmlDir string
}


func GetConfig() ServerConfig {
  var conf ServerConfig
  if _, err := toml.DecodeFile("./conf.toml", &conf); err != nil {
    fmt.Println("Config file could not be read, continuing with defaults...", err)
    conf.Port = port
    conf.HtmlDir = siteRoot
  }
  return conf
}
