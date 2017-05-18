package main

import (
	"github.com/go-xorm/xorm"
	"fmt"

	_ "github.com/lib/pq"
)

type xormEngine struct {
	engine *xorm.Engine
}

const (
	DB_HOST     = "localhost"
	DB_PORT        = 5432
	DB_USER     = "maruf"
	DB_PASSWORD = "pass"
	DB_NAME     = "sample-database"
)

func (en *xormEngine)connect_database() {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+ "password=%s dbname=%s sslmode=disable",
		DB_HOST,DB_PORT,DB_USER, DB_PASSWORD, DB_NAME)
	var err error
	en.engine, err = xorm.NewEngine("postgres", dbinfo)
	if err!=nil{
		panic(err)
	}

	//defer post.db.Close()


	err =en.engine.Ping()
	if err !=nil{
		panic(err)
	}

	fmt.Println("Successfully connected")
}




func main() {
	en:=new(xormEngine)

	en.connect_database()

}
