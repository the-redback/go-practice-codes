package main

import (
	"database/sql"
	"log"
	"github.com/rubenv/sql-migrate"
	_ "github.com/lib/pq"

)

func main() {
	postDB, err := sql.Open("postgres","host=localhost port=5432 user=maruf password=pass dbname=sample-database sslmode=disable")
	if err!=nil{
		panic(err)
	}
	log.Println("Successfully connected.")

	mig :=&migrate.FileMigrationSource{
		Dir: "migrations/postgres",
	}

	upAffected, err :=migrate.Exec(postDB,"postgres",mig,migrate.Up)
	if err!=nil {
		log.Println(err)
	}

	log.Println("UP Applied migrations:",upAffected)


/*	downAffected, err :=migrate.Exec(postDB,"postgres",mig,migrate.Down)
	if err!=nil {
		log.Println(err)
	}
	log.Println("Down Applied migrations:",downAffected)*/

}
