package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var (
	debug         = flag.Bool("debug", true, "enable debugging")
	password      = flag.String("password", "_g0encryptSQL", "the sql login database user password")
	port     *int = flag.Int("port", 1433, "the database port")
	server        = flag.String("server", "HFleitas.fleitasarts.com", "the database server")
	user          = flag.String("user", "mssqlgo", "the sql login database user")
	encrypt       = flag.String("Encrypt", "true", "false by default")
	trust		      = flag.String("TrustServerCertificate", "true", "false by default")
)

func main() {
	flag.Parse()

	if *debug {
		fmt.Printf(" password:%s\n", *password)
		fmt.Printf(" port:%d\n", *port)
		fmt.Printf(" server:%s\n", *server)
		fmt.Printf(" user:%s\n", *user)
		fmt.Printf(" Encrypt:%s\n", *encrypt)
		fmt.Printf(" TrustServerCertificate:%s\n", *trust)
	}

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;Encrypt=%s;TrustServerCertificate=%s", *server, *user, *password, *port, *encrypt, *trust)
	if *debug {
		fmt.Printf(" connString:%s\n", connString)
	}
	conn, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()

	stmt, err := conn.Prepare("select 1, 'abc'")
	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt.Close()

	row := stmt.QueryRow()
	var somenumber int64
	var somechars string
	err = row.Scan(&somenumber, &somechars)
	if err != nil {
		log.Fatal("Scan failed:", err.Error())
	}
	fmt.Printf("somenumber:%d\n", somenumber)
	fmt.Printf("somechars:%s\n", somechars)

	fmt.Printf("bye\n")
}
