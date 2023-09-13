package model

import (
	"control"
	"encoding/json"
	"fmt"
	"os"
)

/*
	{
	    "connections":{
	        "sqlserver":{
	            "databaseName": "SocialMedia",
	            "user": "viet080702",
	            "password": "nguyenanhviet",
	            "serverName": "MSI\\SQLEXPRESS:1433",
	            "otherParameters": "Encrypt=false"
	        },
	        "port": "8080"
	    }
	}
*/
const CONFIG_PATH = "../Config.json"

type Config struct {
	Connections Connections `json:"connections"`
	Port        string      `json:"port"`
}

type Connections struct {
	Sqlserver SqlServer `json:"sqlserver"`
}

type SqlServer struct {
	DatabaseName    string `json:"databaseName"`
	User            string `json:"user"`
	Password        string `json:"password"`
	ServerName      string `json:"serverName"`
	OtherParameters string `json:"otherParameters"`
}

func GetConfig() (Config, error) {
	var config Config
	configFile, err := os.Open(CONFIG_PATH)
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return config, err
	}
	jStr := control.GetFileContent(configFile)
	json.Unmarshal([]byte(jStr), &config)
	//fmt.Printf("%+v\n", config)
	return config, nil
}
