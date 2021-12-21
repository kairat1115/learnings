// https://go.dev/doc/database/open-handle
package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql" // Must be imported anyways, even if mysql.Config not used
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	// We can write DSN ourself with hardcode string or formatted with smth like sprintf
	_, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(127.0.0.1:3306)/%s",
			os.Getenv("MYSQL_USER"),
			os.Getenv("MYSQL_PASSWORD"),
			os.Getenv("MYSQL_DATABASE"),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created handle. Sprintf")

	// Specify connection properties.
	cfg := mysql.Config{
		User:      os.Getenv("MYSQL_USER"),
		Passwd:    os.Getenv("MYSQL_PASSWORD"),
		DBName:    os.Getenv("MYSQL_DATABASE"),
		Net:       "tcp",
		Addr:      "127.0.0.1:3306",
		Collation: "utf8mb4_0900_ai_ci",
	}

	// We can format DSN from config, prefered way.
	_, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created handle. FormatDSN")

	// Get a driver-specific connector.
	connector, err := mysql.NewConnector(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	// Get a database handle with OpenDB by passing connector.
	db := sql.OpenDB(connector)
	fmt.Println("Created handle. Connector")

	fmt.Println("Ping!")
	// Ping the database to confirm a connection
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pong!")

}
