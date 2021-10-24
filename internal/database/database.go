package database

import (
	"github.com/sirupsen/logrus"
	_ "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"jwt-go/util"
)

var DB *gorm.DB

func Connect(cfg util.Config) *gorm.DB {
	var databaseConfig = "" +
		"host=" + cfg.DB_DB_HOST +
		" user=" + cfg.DB_USER +
		" password=" + cfg.DB_PASSWORD +
		" dbname=" + cfg.DB_DBNAME +
		" port=" + cfg.DB_PORT +
		" sslmode=" + cfg.DB_SSLMODE

	DB, err := gorm.Open(postgres.Open(databaseConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		logrus.Fatalln("database down")
	}

	logrus.Println("database", cfg.DB_DBNAME, "successfully connected")

	//err = Drop(DB)
	//if err != nil{
	//	logrus.Fatalln("table drops failed")
	//}

	err = Migrate(DB)
	if err != nil {
		logrus.Fatalln("auto migrations failed")
	}

	logrus.Println("migrations successfully migrated")

	return DB
}
