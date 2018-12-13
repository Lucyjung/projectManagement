
# Project Management With Go 

An application written in Golang for Project Management 
Mainly display gantt chart and query data from timesheet(MySQL) 
Additionally, it should connect to REDMINE system (https://www.redmine.org/)

## Contents

- [Installation](#installation)
- [Prerequisite](#prerequisite)
- [Quick start](#quick-start)
- [Structure](#structure)
## Installation
you need to install Go (https://golang.org/dl/) 
then 
```sh
$ go get github.com/go-sql-driver/mysql
```

## Prerequisite
Requires Go 1.6 or later and Go 1.7 will be required soon.


## Quick start
go to project workspace then run 
```sh
$go run app.go
```

## Structure

- config 
    Configuration Structure and API for read 
- controllers
    function handler for API 
- dao
    Data Access Object -> API for model
- models
    Schema for database 
- view 
    Frontend code -> javascript & html 
- app.go : Main function for this app 