package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"log"
)

func main() {
	// connect to a database
	conn, err := sql.Open("pgx", "host=localhost user=postgres dbname=postgres user=postgres password=")
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}

	log.Println("Connected to database")
	// test my connection
	err = conn.Ping()
	defer conn.Close()
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to ping database: %v\n", err))
	}
	log.Println("Ping to database")
	// get rows from table

	err = getAllRows(conn)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to get all rows: %v\n", err))
	}

	// insert a row
	//query := `insert into users1  (first_name, last_name) values ($1, $2)`
	//_, err = conn.Exec(query, "masud", "ali")
	//if err != nil {
	//	log.Fatal(fmt.Sprintf("Unable to insert row: %v\n", err))
	//}
	//fmt.Println("Insert row successfully")

	// get rows from table again
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to get all rows: %v\n", err))
	}
	// update a row

	updateQuery := `update users1 set first_name = $1 where first_name = $2`
	_, err = conn.Exec(updateQuery, "fuad", "fuadul")
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to insert row: %v\n", err))
	}
	fmt.Println("Successfully updated all rows")
	// get rows from table again
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to get all rows: %v\n", err))
	}
	// get one row by id
	// select id,first_name,last_name from users1 where users1.id = 2
	findById := `select id,first_name,last_name from users1 where users1.id = $1`

	var firstName, lastName string
	var id int
	row := conn.QueryRow(findById, 2)

	err = row.Scan(&id, &firstName, &lastName)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to find row: %v\n", err))
	}
	// get rows from table again
	fmt.Println("Successfully find by ID")
	fmt.Printf("id= %d First Name: %s, Last Name: %s\n", id, firstName, lastName)

	// delete a row
	deleteByID := `delete from users1 where id = $1`
	_, err = conn.Exec(deleteByID, id)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to delete row: %v\n", err))
	}
	fmt.Println("Successfully deleted by ID")
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to get all rows: %v\n", err))
	}
	// get rows again

}

func getAllRows(conn *sql.DB) error {
	rows, err := conn.Query("select id,first_name, last_name from users1")
	if err != nil {
		return err
	}
	defer rows.Close()
	var firstName, lastName string
	var id int
	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			return err
		}
		fmt.Println(id, firstName, lastName)
	}
	if err = rows.Err(); err != nil {
		log.Fatal("Error fetching rows: ", err)
	}
	fmt.Println("------------------------------------")
	return nil
}
