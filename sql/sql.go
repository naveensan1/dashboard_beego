package sql

import (
	"dashboard_beego/models"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//var runlist = []map[string]string{}

func blah(phystopo, subtopo, load string) []models.RunDB {
	db, err := sqlx.Open("mysql",
		"readonly:onlyread@tcp(mvdcjenkins.us.alcatel-lucent.com:3306)/regressdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	run := []models.RunDB{}
	db.Select(&run, "SELECT `regressdb_regression`.`id`,`args`,`link` FROM `regressdb_regression` INNER JOIN `regressdb_load` ON ( `regressdb_regression`.`load_id` = `regressdb_load`.`id` ) WHERE (`regressdb_load`.`load` = ?  AND `regressdb_regression`.`args` REGEXP BINARY ? )", load, subtopo)
	//db.Select(&run, "SELECT * FROM `regressdb_regression` INNER JOIN `regressdb_load` ON ( `regressdb_regression`.`load_id` = `regressdb_load`.`id` ) WHERE (`regressdb_load`.`load` = ?  AND `regressdb_regression`.`args` REGEXP BINARY ? )", load, subtopo)
	//db.Select(&run, "SELECT `regressdb_regression`.`args` FROM `regressdb_regression` WHERE `regressdb_regression`.`args` REGEXP BINARY ? )", subtopo)
	db.Select(&run, "SELECT `regressdb_regression`.* FROM `regressdb_regression` INNER JOIN `regressdb_load` ON ( `regressdb_regression`.`load_id` = `regressdb_load`.`id` ) WHERE (`regressdb_load`.`load` = ?  AND `regressdb_regression`.`args` REGEXP BINARY ? )", load, subtopo)
	fmt.Println(run)
	return run

}

//QuerySubtopo the last 10 runs of a given subtopology
func QuerySubtopo(subtopo string) []map[string]string {
	var runlist = []map[string]string{}
	db, err := sqlx.Open("mysql",
		"readonly:onlyread@tcp(mvdcjenkins.us.alcatel-lucent.com:3306)/regressdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM `regressdb_regression` WHERE `regressdb_regression`.`args` REGEXP BINARY ? ORDER BY `regressdb_regression`.`time_start` DESC LIMIT 10", subtopo)
	if err != nil {
		fmt.Println(err.Error())
	}
	cols, err := rows.Columns() // Remember to check err afterwards
	if err != nil {
		fmt.Println(err.Error())
	}
	lenCN := len(cols)
	vals := make([]interface{}, lenCN)
	//run := make(map[string]string, lenCN)

	for i := range cols {
		vals[i] = new(sql.RawBytes)
	}
	for rows.Next() {

		err = rows.Scan(vals...)
		if err != nil {
			fmt.Println("Failed to scan row", err)

		}
		run := make(map[string]string, lenCN)
		for i := 0; i < lenCN; i++ {

			if rb, ok := vals[i].(*sql.RawBytes); ok {

				//fmt.Printf("%s = %s\n", cols[i], string(*rb))
				run[cols[i]] = string(*rb)

				*rb = nil // reset pointer to discard current value to avoid a bug
			} else {
				fmt.Printf("Cannot Print")
			}
		}
		runlist = append(runlist, run)
	}
	return runlist

}

func queryDB(sqlQuery string, build string) []map[string]string {
	//func queryDB(sqbuild string) []map[string]string {
	var runlist = []map[string]string{}
	db, err := sqlx.Open("mysql",
		"readonly:onlyread@tcp(mvdcjenkins.us.alcatel-lucent.com:3306)/regressdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(sqlQuery, build)
	if err != nil {
		fmt.Println(err.Error())
	}
	cols, err := rows.Columns() // Remember to check err afterwards
	if err != nil {
		fmt.Println(err.Error())
	}
	lenCN := len(cols)
	vals := make([]interface{}, lenCN)
	//run := make(map[string]string, lenCN)

	for i := range cols {
		vals[i] = new(sql.RawBytes)
	}
	for rows.Next() {

		err = rows.Scan(vals...)
		if err != nil {
			fmt.Println("Failed to scan row", err)

		}
		run := make(map[string]string, lenCN)
		for i := 0; i < lenCN; i++ {

			if rb, ok := vals[i].(*sql.RawBytes); ok {

				//fmt.Printf("%s = %s\n", cols[i], string(*rb))
				run[cols[i]] = string(*rb)

				*rb = nil // reset pointer to discard current value to avoid a bug
			} else {
				fmt.Printf("Cannot Print")
			}
		}
		runlist = append(runlist, run)
	}
	return runlist

}

func queryRun(phystopo, subtopo, load string) []map[string]string {
	var runlist = []map[string]string{}
	db, err := sqlx.Open("mysql",
		"readonly:onlyread@tcp(mvdcjenkins.us.alcatel-lucent.com:3306)/regressdb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT `regressdb_regression`.* FROM `regressdb_regression` INNER JOIN `regressdb_load` ON ( `regressdb_regression`.`load_id` = `regressdb_load`.`id` ) WHERE (`regressdb_load`.`load` = ?  AND `regressdb_regression`.`args` REGEXP BINARY ? )", load, subtopo)
	if err != nil {
		fmt.Println(err.Error())
	}
	cols, err := rows.Columns() // Remember to check err afterwards
	if err != nil {
		fmt.Println(err.Error())
	}
	lenCN := len(cols)
	vals := make([]interface{}, lenCN)
	run := make(map[string]string, lenCN)

	for i := range cols {
		vals[i] = new(sql.RawBytes)
	}
	for rows.Next() {

		err = rows.Scan(vals...)
		if err != nil {
			fmt.Println("Failed to scan row", err)

		}
		for i := 0; i < lenCN; i++ {

			if rb, ok := vals[i].(*sql.RawBytes); ok {

				//fmt.Printf("%s = %s\n",cols[i],string(*rb))
				run[cols[i]] = string(*rb)

				*rb = nil // reset pointer to discard current value to avoid a bug
			} else {
				fmt.Printf("Cannot Print")
			}
		}
		runlist = append(runlist, run)
	}
	return runlist

}
