package main

import (
	"fmt"
	"database/sql"

	_ "github.com/lib/pq"

)

type postgresDB struct {
	db *sql.DB
}

const (
	DB_USER     = "maruf"
	DB_PASSWORD = "pass"
	DB_NAME     = "sample-database"
)

func (post *postgresDB)connect_database() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	var err error
	post.db, err = sql.Open("postgres", dbinfo)
	if err!=nil{
		panic(err)
	}
	//defer post.db.Close()


	err =post.db.Ping()
	if err !=nil{
		panic(err)
	}

	fmt.Println("Successfully connected")
}



func (post *postgresDB)insert_query() {
	var err error

	sqlStatement := `INSERT INTO public.user_info(
	name, dept)
	VALUES ($1, $2);`
	id := 1
	err = post.db.QueryRow(sqlStatement, "name","dept456").Scan(&id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("New record ID is:", id)
}



func main() {
	db:=new(postgresDB)

	db.connect_database()
	db.insert_query()

	defer db.db.Close()

}
