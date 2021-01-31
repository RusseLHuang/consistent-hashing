package database

import (
	"log"
	"os"
	"time"

	"github.com/RusseLHuang/consistent-hashing/user/entity"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DBConn *gorm.DB

var DBConnMap map[string]*gorm.DB

func openConnection(
	url string,
) *gorm.DB {
	mysqlConfig := viper.Get("mysql").(map[string]interface{})
	user := mysqlConfig["user"].(string)
	password := mysqlConfig["password"].(string)
	database := mysqlConfig["database"].(string)
	charset := mysqlConfig["charset"].(string)

	dsn := user + ":" + password + "@tcp(" + url + ")/" + database + "?charset=" + charset + "&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&entity.User{})

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetMaxOpenConns(1)
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	return db
}

func InitDB() map[string]*gorm.DB {
	if DBConnMap == nil {
		DBConnMap = make(map[string]*gorm.DB)
	}

	dbKeyList := FetchDBKeyList()

	for i := 0; i < len(dbKeyList); i++ {
		dbKey := dbKeyList[i]
		dbConn := FetchDBConnectionInfo(dbKey)
		DBConnMap[dbKey] = openConnection(dbConn)
	}

	return DBConnMap
}
