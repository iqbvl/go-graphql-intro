package models

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
)

type Config struct {
	Database []struct {
		Type     string `json:"type"`
		Host     string `json:"host"`
		Instance string `json:"instance"`
		User     string `json:"user"`
		Password string `json:"password"`
		Port     string `json:"port"`
		DBName   string `json:"dbName"`
	} `json:"database"`
	ServerPort string
	Credential struct {
		JWTSecret    string
		JWTExpiresIn string
	}
	AppName        string
	LogFile        string `json:"logFile"`
	EnableGraphiQL bool   `json:"enableGraphiQL"`
	TimeoutSecond  int    `json:"timeout_second"`
}

var configuration Config

func LoadConfiguration() Config {
	var config Config
	configFile, err := os.Open("config.json")
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

func BuildConnString() string {
	config := LoadConfiguration()
	query := url.Values{}
	query.Add("database", config.Database[0].DBName)
	query.Add("packet size", "4096")

	u := &url.URL{
		Scheme: "sqlserver",
		User:   url.UserPassword(config.Database[0].User, config.Database[0].Password),
		Host:   fmt.Sprintf("%s:%s", config.Database[0].Host, config.Database[0].Port),
		//Host:   config.Database[0].Host,
		// Path:  instance, // if connecting to an instance instead of a port
		RawQuery: query.Encode(),
	}
	return u.String()
}

func Timeout() int {
	config := LoadConfiguration()
	return config.TimeoutSecond
}
