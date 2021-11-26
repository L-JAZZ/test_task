package configs

import (
	"encoding/json"
	"fmt"
	"log"

	"os"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type ConfigStruct struct {
	Url        string        `json:"url"`
	ListenIP   string        `json:"listenIP"`
	ListenPort string        `json:"listenPort"`
	DB         DatabaseModel `json:"db"`
}

type DatabaseModel struct {
	Driver   string `json:"driver"`
	DbServer string `json:"db_server"`
	DbPort   int    `json:"db_port"`
	DbName   string `json:"db_name"`
	DbUser   string `json:"db_user"`
	DbPasswd string `json:"db_passwd"`
}

var Config ConfigStruct

func LoadConfiguration(file string) {

	configFile, err := os.Open(file)
	defer func(configFile *os.File) {
		err := configFile.Close()
		if err != nil {
			log.Println("[LoadConfiguration]: .Close() fail")
		}
	}(configFile)

	if err != nil {
		fmt.Println("[LoadConfiguration]: fail")
		panic(fmt.Sprintf("Ошибка чтения конфигурации: %+v", err))
	}
	//fmt.Println("[LoadConfiguration]: success")

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&Config)
	if err != nil {
		return
	}
}

var db *sqlx.DB

// DB - Singleton Database connection
func DB() *sqlx.DB {

	if db == nil {
		cString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
			Config.DB.DbServer,
			Config.DB.DbPort,
			Config.DB.DbUser,
			Config.DB.DbName,
			Config.DB.DbPasswd)

		newDb, err := sqlx.Connect(Config.DB.Driver, cString)

		if err != nil {
			fmt.Printf("%+v", err)
		}
		// else {
		// 	fmt.Printf("Successfully connected with %s!", configs.Config.DB.DbName)
		// }

		db = newDb
	}
	return db
}
