package connections

import (
	"apiservice/common"
	"fmt"
	"strconv"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConnection struct {
	User             string
	Password         string
	DBName           string
	Host             string
	Port             int
	BinaryParameters bool
	DB               *gorm.DB
}

func String() string {

	port := common.EnvVariable("DB_PORT")
	portInt, _ := strconv.Atoi(port)

	return strings.Join([]string{
		fmt.Sprintf("user=%s", common.EnvVariable("DB_USER")),
		fmt.Sprintf("password=%s", common.EnvVariable("DB_PASSWORD")),
		fmt.Sprintf("dbname=%s", common.EnvVariable("DB_DBNAME")),
		// fmt.Sprintf("host=%s", common.EnvVariable("DB_HOST")),
		fmt.Sprintf("port=%d", portInt),
		// FIXME: sslmode should be a property
		"sslmode=disable"}, " ")
}

func DBConn() (db *gorm.DB) {
	db, err := gorm.Open(postgres.Open(String()), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
