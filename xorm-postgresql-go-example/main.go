package main

import (
	"github.com/go-xorm/xorm"
	"fmt"

	_ "github.com/lib/pq"
	"log"
	"time"
)



type User_info struct{
	Uid int64  `xorm:"id pk not null autoincr"`
	Name string
	Dept string
	CreatedAt time.Time `xorm: "created"`
	UpdatedAt time.Time `xorm: "updated"`

}

const (
	DB_HOST     = "localhost"
	DB_PORT     = 5432
	DB_USER     = "maruf"
	DB_PASSWORD = "pass"
	DB_NAME     = "sample-database"
)


func connect_database() *xorm.Engine{
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST,DB_PORT,DB_USER, DB_PASSWORD, DB_NAME)
	var err error
	en, err := xorm.NewEngine("postgres", dbinfo)
	if err!=nil{
		log.Println("engine creation failed", err)
	}

	//defer post.db.Close()


	err =en.Ping()
	if err !=nil{
		panic(err)
	}

	log.Println("Successfully connected")
	return en
}

func sync_tables(en *xorm.Engine){
	err:=en.Sync(new(User_info))
	if err!=nil{
		log.Println("creation error",err)
		return
	}
	log.Println("Successfully synced")
}

func insert_data(en *xorm.Engine){
	u:=new(User_info)
	u.Uid=1104029
	u.Name="Maruf"
	u.Dept="cse"
	affected, _ :=en.Insert(u)

	log.Println("Inserted user id:",u.Uid,":: affected:",affected)


	//Another way without using variable
	affcted, _ := en.Insert(&User_info{Uid: 1104051,Name:"Dipta",Dept: "ME"})

	log.Println("Insertion signle data. affected:",affcted)
}

func insert_multiple_data(en *xorm.Engine){
	users := make([]User_info, 3)
	users[0].Name = "shohel"
	users[1].Name = "Hasan"
	users[2].Name = "Uttam"
	affected, _ := en.Insert(&users)

	log.Println("affected:",affected)

	//another way
	users2 := make([]*User_info, 2)
	users2[0]=new(User_info)
	users2[0].Name = "Sanjidul Haque"
	users2[0].Uid=904200
	users2[0].Dept="CSE"

	users2[1]=new(User_info)
	users2[1].Name = "Sauman Biswas"

	affd, _ := en.Insert(&users2)


	log.Println("affected:",affd)

}





func main() {

	en:=connect_database()

	sync_tables(en)
	insert_data(en)
	insert_multiple_data(en)




}
