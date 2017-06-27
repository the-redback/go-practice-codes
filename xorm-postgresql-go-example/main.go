package main

import (
	"github.com/go-xorm/xorm"
	"fmt"

	_ "github.com/lib/pq"
	"log"
	"time"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = 5432
	DB_USER     = "maruf"
	DB_PASSWORD = "pass"
	DB_NAME     = "sample-database"
)


type User_info struct{
	Uid int64  `xorm:"pk not null autoincr"`
	Name string
	Dept string
	CreatedAt time.Time `xorm: "created"`
	UpdatedAt time.Time `xorm: "updated"`

}

type Library_info struct {
	Entry_id int64 `xorm:"pk not null autoincr"`
	Book_id int64
	Student_id int64
}

type Book_info struct{
	Book_id int64 `xorm:"pk not null autoincr"`
	Author_name string
}

//Structure required for joining tables

type User_library struct{
	User_info `xorm:"extends"`
	Library_info `xorm:"extends"`
	Book_info `xorm:"extends"`
}

func (User_library) TableName() string{
	return "user_info"
}

//End Here


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
	err=en.Sync(new(Library_info))
	err=en.Sync(new(Book_info))
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
	u.Dept="CSE"
	affected, _ :=en.Insert(u)

	log.Println("Inserted user id:",u.Uid,":: affected:",affected)


	//Another way without using variable
	affcted, _ := en.Insert(&User_info{Uid: 1104051,Name:"Dipta",Dept: "ME"})
	en.Insert(&User_info{Uid: 1104119,Name:"Kaium",Dept: "PME"})

	log.Println("Insertion signle data. affected:",affcted)
}

func insert_multiple_data(en *xorm.Engine){
	users := make([]User_info, 3)
	users[0].Name = "Shohel"
	users[0].Dept = "Archi"
	users[0].Uid = 1104023

	users[1].Name = "Hasan"
	users[1].Dept = "CSE"
	users[1].Uid = 1104009
	users[2].Name = "Uttam"
	users[2].Dept = "ETE"
	users[2].Uid = 1104089

	affected, _ := en.Insert(&users)

	log.Println("affected:",affected)

}

func query_single_data(en *xorm.Engine){
	user:=User_info{Uid: 1104029}
	has,_:=en.Get(&user)

	log.Println(has)
	log.Println(user)

	//Another way
	var user2 User_info
	has2, _ :=en.Id(1104051).Get(&user2)    //Primary key
	log.Println(has2)
	log.Println(user2)
}

func query_mulltiple_data(en *xorm.Engine){

	var users []User_info
	err := en.Find(&users)
	if err!=nil{
		log.Println(err)
	}

	//log.Println(users)
	for _, user:=range  users{
		log.Println(user)
	}

}


func query_conditional_data(en *xorm.Engine){
	var err error

	var tenusers []User_info
	err =en.Cols("uid", "name").Where("Dept = ?","CSE").Limit(10).Find(&tenusers)
	if err !=nil{
		log.Println(err)
	}


	for _,user := range tenusers{
		log.Println(user)
	}

	//en.Sql("select * from User_info").Find(&tenusers)   //Query by sql

	//Using Mapping
/*	users := make(map[int64]User_info)
	err = en.Find(&users)
	if err!=nil{
		log.Println(err)
	}
	log.Println(users)*/


	//IN function
/*	var UsersIN []User_info
	err = en.In("dept", "CSE", "ME", "CE").Find(&UsersIN) //Get All dept in ("CSE","ME", "CE")
	if err !=nil{
		log.Println(err)
	}


	for _,user := range UsersIN{
		log.Println(user)
	}*/


	//Just a slice of ID
/*	var ints []int64
	err = en.Table("user_info").Cols("uid").Find(&ints)
	if err !=nil{
		log.Println(err)
	}

	log.Println(ints)*/
}


func join_User_Library(en *xorm.Engine){
	var users []User_library
	err:= en.Join("INNER","library_info","library_info.student_id = user_info.uid").Join("INNER","book_info","book_info.book_id = library_info.book_id").Find(&users)
	if err!=nil {
		log.Println(err)
	}
	//log.Println(users)

	for _,user :=range users{
		log.Println(user)
	}
}


func main() {

	en:=connect_database()

	sync_tables(en)
	//insert_data(en)
	//insert_multiple_data(en)

	//query_single_data(en)
	query_mulltiple_data(en)

	//query_conditional_data(en)

	//join_User_Library(en)



}
