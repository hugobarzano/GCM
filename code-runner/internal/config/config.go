package config

import (
	"code-runner/internal/constants"
	"encoding/json"
	"github.com/fulldump/goconfig"
	"log"
	"os"
	"sync"
)

var (
	cfg  *Config
	once sync.Once
)

type Config struct {
	Debug              bool   `json:"debug"`
	Version            bool   `usage:"Show version"`
	GithubClientID     string `json:"githubClientID"`
	GithubClientSecret string `json:"githubClientSecret"`
	MongoUri           string `usage:"Standard MongoDB Hostname"`
	DeployAddress      string `json:"deployAddress"`
	ApiAddress         string `json:"apiAddress"`
	ApiPort            string `json:"apiPort"`
	ApiDns             string `json:"apiDns"`
	EnableTls          bool   `json:"enableTls"`
	TlsKeyFile         string `json:"tlsKeyFile"`
	TlsCertFile        string `json:"tlsCertFile"`
}

func GetConfig() *Config {
	once.Do(func() {
		if cfg == nil {
			cfg = read()
		}
	})
	return cfg
}

func read() *Config {

	c := &Config{
		Debug:              true,
		GithubClientID:     "hhh",
		GithubClientSecret: "ggg",
		MongoUri:           "mongodb://localhost:27017",
		DeployAddress:      "localhost",
		ApiAddress:         "0.0.0.0",
		ApiPort:            "80",
		EnableTls:          false,
	}

	goconfig.Read(c)

	if c.Version {
		log.Println(constants.Version)
		os.Exit(0)
	}

	if c.Debug {
		j := json.NewEncoder(os.Stderr)
		j.SetIndent("", "    ")
		if err := j.Encode(c); err != nil {
			log.Fatal(err.Error())
		}
	}
	return c
}
