package main

import (
    "log"

    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
)

func main() {
  //connect to a PostgreSQL database
  // Replace the connection details (user, dbname, password, host) with your own
    db, err := sqlx.Connect("postgres", "user=postgres dbname=yourdatabase sslmode=disable password=password host=localhost")
    if err != nil {
        log.Fatalln(err)
    }
  
    defer db.Close()

    // Test the connection to the database
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    } else {
        log.Println("Successfully Connected")
    }

}

type User struct {
	Name  string `db:"username"`
	Email string `db:"email"`
   }


   place := User{} // Initialize a User struct to hold retrieved data

   // Execute a SQL query to select "username" and "email" columns from the "users" table
   rows, _ := db.Queryx("SELECT username, email FROM users")
   
   for rows.Next() {
	   err := rows.StructScan(&place) // Scan the current row into the "place" variable
	   if err != nil {
		   log.Fatalln(err)
	   }
	   log.Printf("%#v\n", place) // Log the content of the "place" struct for each row
   }


//

// import (
//  "log"

//  "github.com/jmoiron/sqlx"
//  _ "github.com/lib/pq"
// )

// type User struct {
//    Name  string `db:"username"`
//    Email string `db:"email"`
// }

// func main() {
//      db, err := sqlx.Connect("postgres", "user=postgres dbname=yourdatabase sslmode=disable password=yourpassword host=localhost")
//      if err != nil {
//         log.Fatalln(err)
//      }
//      defer db.Close()

//      if err := db.Ping(); err != nil {
//         log.Fatal(err)
//      } else {
//         log.Println("Successfuly connected")
//      }

//      place := User{}
//      rows, _ := db.Queryx("SELECT username, email FROM users")
//      for rows.Next() {
//         err := rows.StructScan(&place)
//         if err != nil {
//            log.Fatalln(err)
//     }
//     log.Printf("%#v\n", place)
//    }
// }
//