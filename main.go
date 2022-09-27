package main
import (
    "assignment-2/routes"
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "root"
    dbname   = "orders_by"
)

var (
    db *sql.DB
    err error
)

func main() {
    var PORT = ":8080"

    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
    db, err = sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        panic(err)
    }
    fmt.Println("Connected!")
    routers.StartServer().Run(PORT)
}