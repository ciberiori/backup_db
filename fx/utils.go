package fx

import (
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
	BACKUP_SRC  string
}

func GetConfig(params ...string) Configuration {
	configuration := Configuration{}
	filename := "config.json"

	gonfig.GetConf(filename, &configuration)
	return configuration

}
