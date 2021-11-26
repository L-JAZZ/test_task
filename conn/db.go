package someshit

import (
	"fmt"
	"test_task/configs"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

// DB - Singleton Database connection
func DB() *sqlx.DB {

	if db == nil {
		cString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
			configs.Config.DB.DbServer,
			configs.Config.DB.DbPort,
			configs.Config.DB.DbUser,
			configs.Config.DB.DbName,
			configs.Config.DB.DbPasswd)

		newDb, err := sqlx.Connect(configs.Config.DB.Driver, cString)

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
