package model

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"

	_ "vendor/github.com/lib/pq"
	yaml "vendor/gopkg.in/yaml.v2"
)

type IModelDao interface {
	dbSelect(*sql.DB) error
	dbInsert(*sql.DB) error
	dbUpdate(*sql.DB) error
	dbDelete(*sql.DB) error
}

type config struct {
	Host     string `yaml:"Host"`
	DBName   string `yaml:"DBName"`
	DBPort   int    `yaml:"DBPort"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
}

type DBCenter struct {
	c  *config
	db *sql.DB
	op IModelDao
}

func ParseConfig() *config {
	c := new(config)
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile Get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	log.Println(c)

	return c
}

func InitDB() *DBCenter {
	bc := new(DBCenter)
	bc.c = ParseConfig()

	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
		bc.c.Host, bc.c.DBPort, bc.c.User, bc.c.Password, bc.c.DBName)

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	bc.db = db
	log.Println("Successfully connected!")

	return bc
}

func (bc *DBCenter) SetModelDao(curOp IModelDao) {
	bc.op = curOp
	bc.op.dbInsert(bc.db)
	bc.op.dbSelect(bc.db)
	//bc.op.dbUpdate(bc.db)
	bc.op.dbDelete(bc.db)

}
