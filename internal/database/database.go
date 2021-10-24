package database

import (
	"github.com/sirupsen/logrus"
	_ "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const HOST = "localhost"
const USER = "postgres"
const PASSWORD = "root"
const DBNAME = "authorization"
const PORT = "5432"
const SSLMODE = "disable"

const databaseConfig = "host=" + HOST + " user=" + USER + " password=" + PASSWORD + " dbname=" + DBNAME + " port=" + PORT + " sslmode=" + SSLMODE

var DB *gorm.DB

func Connect() *gorm.DB {
	DB, err := gorm.Open(postgres.Open(databaseConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: false,
	})
	if err != nil {
		logrus.Fatalln("database down")
	}

	logrus.Println("database", DBNAME, "successfully connected")

	err = Migrate(DB)
	if err != nil {
		logrus.Fatalln("auto migrations failed")
	}

	//err = Drop(DB)
	//if err != nil{
	//	logrus.Fatalln("table drops failed")
	//}

	logrus.Println("migrations successfully migrated")

	return DB
}
