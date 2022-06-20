package main


import (
   // "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
	"log"
)

/*
 * Tag... - a very simple struct
 */
 type Tag struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}


func main() {
    // Open up our database connection.
    db, err := sql.Open("mysql", "springuser:TECHinterrupt123@tcp(127.0.0.1:3306)/fanta")

    // if there is an error opening the connection, handle it
    if err != nil {
        log.Print(err.Error())
    }
    defer db.Close()

    // Execute the query
    results, err := db.Query("SELECT id, name FROM test")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    for results.Next() {
        var tag Tag
        // for each row, scan the result into our tag composite object
        err = results.Scan(&tag.ID, &tag.Name)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
                // and then print out the tag's Name attribute
        log.Printf(tag.Name)
    }

}