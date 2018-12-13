package dao

import (
	"database/sql"
	"fmt"

	config "../config"
	mdl "../models"
	_ "github.com/go-sql-driver/mysql" // for mysql
)

// DB Global Database object
var DB *sql.DB

// Connect : Establish a connection to database
func Connect() {
	c := config.Read()

	const (
		COLLECTION = "dev_timesheet"
	)

	connectionString := fmt.Sprintf("%s:%s@tcp(10.164.232.200:3306)/%s", c.DBUser, c.DBPass, COLLECTION)

	var err error
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	err = db.Ping()
	if err != nil {
		fmt.Print(err.Error())
	}
	DB = db

}

// QueryProject Test
func QueryProject(id int) mdl.Project {

	var prj mdl.Project
	err := DB.QueryRow("SELECT project_id,project_name FROM project WHERE project_id=?", id).Scan(&prj.ProjectID, &prj.ProjectName)

	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	//log.Println(prj.ProjectID)
	//log.Println(prj.ProjectName)
	return prj
}
