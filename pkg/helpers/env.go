package helpers

import (
	"os"
	"strconv"
    "github.com/joho/godotenv"
)

func InitEnvSetting()(err error){
    error := godotenv.Load(".env")
    return error
}

func InitTestEnvSetting(envPath string)(err error){
    error := godotenv.Load(envPath)
    return error
}

func GetEnvStr(key string) (string) {
    v := os.Getenv(key)
    return v
}

func GetEnvInt(key string) (int int , err error) {
    s:= GetEnvStr(key)
    v, err := strconv.Atoi(s)
    return v ,err
}

func GetEnvBool(key string) (bool, error) {
    s := GetEnvStr(key)
    v, err := strconv.ParseBool(s)
    return v, err
}