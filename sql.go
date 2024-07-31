package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type studentSQL struct {
	id    string
	name  string
	age   int
	grade int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/db_novgo")
	if err != nil {
		return nil, err
	}

	return db, nil
}

// using SQL Query
func sqlQuery() {

	// connect to db
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// don't forget to close connection using defer
	defer db.Close()

	age := 27
	rows, err := db.Query("select id, name, grade from tb_student where age= ?", age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// don't forget to close connection using defer
	defer rows.Close()

	var result []studentSQL

	for rows.Next() {
		var each = studentSQL{}
		var err = rows.Scan(&each.id, &each.name, &each.grade)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.name)
	}
}

// using SQL Query Prepare()
func sqlQueryPrepare() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// don't forget to close connection using defer
	defer db.Close()

	query, err := db.Prepare("select name, grade from tb_student where id= ?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	result1 := studentSQL{}
	query.QueryRow("G001").Scan(&result1.name, &result1.grade)
	fmt.Printf("name: %s\ngrade: %d\n", result1.name, result1.grade)

	result2 := studentSQL{}
	query.QueryRow("E001").Scan(&result2.name, &result2.grade)
	fmt.Printf("name: %s\ngrade: %d\n", result2.name, result2.grade)

	result3 := studentSQL{}
	query.QueryRow("W001").Scan(&result3.name, &result3.grade)
	fmt.Printf("name: %s\ngrade: %d\n", result3.name, result3.grade)

}

// SQL Exec for Insert, Update, & Delete
func sqlExec() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("insert into tb_student values(?, ?, ?, ?)", "G001", "Gaal Gadoot", 29, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("insert to db success")

	_, err = db.Exec("update tb_student set age= ? where id= ?", 30, "G001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("update to db success")

	_, err = db.Exec("delete from tb_student where id=  ?", "B001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("delete data from tb_student success")
}

func main() {
	fmt.Println("SQL on Go!")
	sqlQuery()
	sqlQueryPrepare()
	// sqlExec()
}
