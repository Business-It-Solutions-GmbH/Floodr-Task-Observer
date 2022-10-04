package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type job struct {
	id       int    `json:"id"`
	json     string `json:"json"`
	status   string `json:"status"`
	attackid int    `json:"attackid"`
}

type attack struct {
	id       int    `json:"id"`
	ip       string `json:"ip"`
	name     string `json:"name"`
	runnerid int    `json:"runner_id"`
	userid   int    `json:"user_id"`
}

func main() {
	fmt.Println("Floodr-Task-Observer v.0.0.1 - By BITS")
	dbconnectionString := os.Getenv("DBCONNECTIONSTRING")

	if dbconnectionString == "" {
		fmt.Println("No DB Connection String provided")
		os.Exit(1)
	}
	db, err := sql.Open("mysql", dbconnectionString)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	results, err := db.Query("SELECT * FROM job WHERE status = 'pending'")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var job job
		err = results.Scan(&job.id, &job.json, &job.status, &job.attackid)
		if err != nil {
			panic(err.Error())
		}
		//get attack
		var attack attack
		err := db.QueryRow("SELECT * FROM attack WHERE id = ?", job.attackid).Scan(&attack.id, &attack.ip, &attack.name, &attack.runnerid, &attack.userid)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(job)
		fmt.Println(attack)
	}
	fmt.Println()
}
