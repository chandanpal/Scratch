package config
import (
        "os"
	"encoding/json"
	"log"
)
var CONF_FILE = "conf.json"
var CONF Config

type Config struct{
    Database struct{
	Dialect string `json:"dialect"`
        Host string `json:"host"`
        Port int `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name string `json:"name"`
    } `json:"database"`
    GITREPO string `json:"gitrepo"`
    Port int `json:"Port"`
    Logger struct{
	EnableConsole bool `json:"enableConsole"`
	ConsoleLevel string `json:"consoleLevel"`
	ConsoleJSONFormat bool `json:"consoleJSONFormat"`
	EnableFile bool `json:"enableFile"`
	FileLevel string `json:"fileLevel"`
	FileJSONFormat bool `json:"fileJSONFormat"`
	FileLocation string `json:"fileLocation"`
	AccessLog string `json:"accessLog"`
	MaxSize int `json:"maxsize"`
	Compress bool `json:"compress"`
	MaxAge int `json:"maxAge"`
	} `json:"logger"`
	HelmRepo struct {
		Type string `json:"type"`
		Url string `json:"url"`
	}`json:"helmRepo"`

}

func LoadConf() {
    confFile, err := os.Open(CONF_FILE)
    defer confFile.Close()
    if err != nil{
        log.Fatalf("Error in opening file ", CONF_FILE)
    }
    jsonParser := json.NewDecoder(confFile)
    err = jsonParser.Decode(&CONF)
    if err != nil{
	log.Fatalf("Loading conf failed", err)
    }
}
