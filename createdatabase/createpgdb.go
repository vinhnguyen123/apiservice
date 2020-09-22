package main

import (
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

var DBconnect = map[string]string{
	"DB_USER":     "vinhnt",
	"DB_PASSWORD": "postgres",
	"DB_DBNAME":   "questions_api",
	"DB_PORT":     "5432",
	"DB_HOST":     "127.0.0.1",
}

func main() {
	createPgDb()
	createTableUser()
	createTableQuestion()
	createTableAnswer()
	createTableQuestionAnswerInfo()
	createTableTag()
	createTableQuestionTagInfo()
	createTableTokenUser()
}

func String() string {

	port := DBconnect["DB_PORT"]
	portInt, _ := strconv.Atoi(port)

	return strings.Join([]string{
		fmt.Sprintf("user=%s", DBconnect["DB_USER"]),
		fmt.Sprintf("password=%s", DBconnect["DB_PASSWORD"]),
		fmt.Sprintf("dbname=%s", DBconnect["DB_DBNAME"]),
		// fmt.Sprintf("host=%s", common.EnvVariable("DB_HOST")),
		fmt.Sprintf("port=%d", portInt),
		// FIXME: sslmode should be a property
		"sslmode=disable"}, " ")
}

func DBConn() (db *gorm.DB) {
	db, err := gorm.Open(postgres.Open(String()), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return db
}

func createPgDb() {

	cmd := exec.Command("createdb", "-p", DBconnect["DB_PORT"], "-h", DBconnect["DB_HOST"], "-U", DBconnect["DB_USER"], "-e", DBconnect["DB_DBNAME"])
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		log.Printf("Error: %v", err)
	}
}

func createTableUser() {
	dbExc := DBConn()
	query := `
  CREATE TABLE IF NOT EXISTS users (
  	id serial PRIMARY KEY,
  	user_name VARCHAR ( 255 ) UNIQUE NOT NULL,
  	email VARCHAR ( 255 ) UNIQUE NOT NULL,
  	encrypted_password VARCHAR ( 255 ) NOT NULL,
  	salt VARCHAR ( 255 ) NOT NULL,
  	created_date TIMESTAMP NOT NULL
  );
	`
	tx := dbExc.Exec(query)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}
}

func createTableQuestion() {
	dbExc := DBConn()
	query := `
  CREATE TABLE IF NOT EXISTS questions (
  	id serial PRIMARY KEY,
  	question_desc TEXT NOT NULL,
  	up_vote_by TEXT,
  	down_vote_by TEXT,
		count_up_vote INTEGER DEFAULT 0,
		count_down_vote INTEGER DEFAULT 0,
  	created_by VARCHAR ( 255 ) NOT NULL,
  	created_date TIMESTAMP NOT NULL
  );
	`
	tx := dbExc.Exec(query)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}
}

func createTableAnswer() {
	dbExc := DBConn()
	query := `
  CREATE TABLE IF NOT EXISTS answers (
  	id serial PRIMARY KEY,
  	answer_desc TEXT NOT NULL,
  	up_vote_by TEXT,
  	down_vote_by TEXT,
		count_up_vote INTEGER DEFAULT 0,
		count_down_vote INTEGER DEFAULT 0,
  	created_by VARCHAR ( 255 ) NOT NULL,
  	created_date TIMESTAMP NOT NULL
  );
	`
	tx := dbExc.Exec(query)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}
}

func createTableTag() {
	dbExc := DBConn()
	query := `
  CREATE TABLE IF NOT EXISTS tags (
  	id serial PRIMARY KEY,
  	tag_desc TEXT NOT NULL
  );
	`
	tx := dbExc.Exec(query)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}
}

func createTableQuestionAnswerInfo() {
	dbExc := DBConn()
	query := `
  CREATE TABLE IF NOT EXISTS question_answer_infos (
  	id serial PRIMARY KEY,
  	question_id INTEGER NOT NULL,
  	answer_id INTEGER NOT NULL
  );
	`
	tx := dbExc.Exec(query)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}
}

func createTableQuestionTagInfo() {
	dbExc := DBConn()
	query := `
  CREATE TABLE IF NOT EXISTS question_tag_infos (
  	id serial PRIMARY KEY,
  	question_id INTEGER NOT NULL,
  	tag_id INTEGER NOT NULL
  );
	`
	tx := dbExc.Exec(query)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}
}

func createTableTokenUser() {
	dbExc := DBConn()
	query := `
  CREATE TABLE IF NOT EXISTS token_users (
  	id serial PRIMARY KEY,
		user_id INTEGER NOT NULL,
  	access_uid TEXT NOT NULL,
		at_expires TIMESTAMP NOT NULL
  );
	`
	tx := dbExc.Exec(query)
	if tx.Error != nil {
		fmt.Println(tx.Error)
	}
}
