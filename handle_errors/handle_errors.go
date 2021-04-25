package main

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	db_user   = "USER"
	db_passwd = "PASSWORD"
	db_addr   = "IP/ENDPOINT"
	db_db     = "DATABASE"
)

type Person struct {
	Id   int
	Name string
	Age  int
}

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", db_user, db_passwd, db_addr, db_db))
	HandleError(err)
	defer db.Close()

	people, err := queryData(db)
	HandleError(err)
	fmt.Println(people)
}

func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}

func queryData(db *sql.DB) (people []Person, err error) {
	resp, err := db.Query("SELECT * FROM `person`")
	defer resp.Close()

	if err != nil {
		return people, err
	}

	for resp.Next() {
		var p Person
		err = resp.Scan(&p.Id, &p.Name, &p.Age)
		if err != nil {
			// 根据具体的业务逻辑，如果列表为空，可以返回空列表，直接将 error 在这里处理掉
			// 也可以将这个处理逻辑放掉调用层，那么调用层就需要做相应的处理
			if err == sql.ErrNoRows {
				return people, nil
			}
			return people, err
		}
		people = append(people, p)
	}
	return people, nil
}
