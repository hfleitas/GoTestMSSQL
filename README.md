# GoTestMSSQL

GO-lang test connecting TDS encryption with mssql driver.

### References

MS SQL Server (pure go): [go-mssqldb](https://github.com/denisenkom/go-mssqldb)

See SQL database drivers: [Wiki](https://github.com/golang/go/wiki/SQLDrivers)

How to write Go code: [Document](https://golang.org/doc/code.html)

## Setup

Requires Go 1.8 or above.

Downald & Install Go-Lang [https://golang.org/dl/](https://golang.org/dl/)

Install the go-mssql driver with cmd, which we're going to use to test TDS Encryption.
```
go get github.com/denisenkom/go-mssqldb
```

### Basic Configure & Test Go-lang Installation. 
1. Setup Windows Test Env via cmd.
```
cd %USERPROFILE%
mkdir go
cd go
mkdir src\hello
cd src\hello
```
### Test your installation
Rerefence actual Go [doc](https://golang.org/doc/install#testing)

2. Create a file named hello.go that looks like this:
```
package main

import "fmt"

func main() {
fmt.Printf("hello, world\n")
}
```
5. From that directory, build the hello.exe with the go tool and run it:
```
go build
hello.exe
```
If you see the "hello, world" message then your Go installation is working.

## The Fun Part
6. Create another folder under go\src named encrypt.
```
cd %USERPROFILE%\go\src
mkdir encrypt
cd encrypt
```
7. Create a file named encrypt.go that looks like this. **Edit values** (server - use fqdn **Case-Sensitive**, username, password):
```
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
	trust	      = flag.String("TrustServerCertificate", "true", "false by default")
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
```
8. Build encrypt.exe using the go tool. Note, go help build shows detail help. See also: go install, go get, go clean.
```
go help build
go clean
go build
```

### Run the test 
9. Run encrypt.exe via cmd.
```
cd %USERPROFILE%\go\src\encrypt
encrypt.exe
```
10. Successful output will look like this:
```
C:\Users\hfleitas\go\src\encrypt>encrypt.exe
 password:_g0encryptSQL
 port:1433
 server:HFleitas.fleitasarts.com
 user:mssqlgo
 Encrypt:true
 TrustServerCertificate:true
 connString:server=HFleitas.fleitasarts.com;user id=mssqlgo;password=_g0encryptSQL;port=1433;Encrypt=true;TrustServerCertificate=true
somenumber:1
somechars:abc
bye
```
11. Unsuccessful output will look like this:
```
C:\Users\hfleitas\go\src\encrypt>encrypt.exe
 password:_g0encryptSQL
 port:1433
 server:hfleitas
 user:mssqlgo
 Encrypt:true
 TrustServerCertificate:true
 connString:server=hfleitas;user id=mssqlgo;password=_g0encryptSQL;port=1433;encrypt=true;trust=true
2018/11/20 13:37:41 Prepare failed:TLS Handshake failed: x509: certificate is valid for HFleitas.fleitasarts.com, not hfleitas
```

## Authors

* **Hiram Fleitas** - *This Repo - GoTestMSSQL* - [Hiram Fleitas](https://github.com/hfleitas)

See also the list of [contributors](https://github.com/hfleitas/GoTestMSSQL/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* [denisenkom](https://github.com/denisenkom)

