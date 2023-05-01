package helpers

import (
	"fmt"
	"log"
	// "os"
	// "strconv"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"github.com/joho/godotenv"
)

var SqlSession *gorm.DB

func NewSqlSession() *gorm.DB{
	return SqlSession
}

func InitMySql()(err error , db *gorm.DB)  {
	envErr := godotenv.Load(".env")
    if envErr != nil {
        log.Fatal("Error loading .env file")
    }
	url := GetEnvStr("db.url")
  	userName := GetEnvStr("db.username")
	passWord := GetEnvStr("db.password")
	dbName := GetEnvStr("db.name")
	port := GetEnvStr("db.port")
	maxIdleConn,_ := GetEnvInt("db.maxIdleConn")
	maxPoolConn,_ := GetEnvInt("db.maxPoolConn")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		userName,
		passWord,
		url,
		port,
		dbName,
	)

	SqlSession,err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err !=nil{
		panic(err.Error())
	}
	sqlDB, err := SqlSession.DB()
	if err != nil {
		panic(err.Error())
	}
	sqlDB.SetMaxIdleConns(maxIdleConn)
	sqlDB.SetMaxOpenConns(maxPoolConn)
	return err,SqlSession
}
