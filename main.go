package main

import (
  "bytes"
  "database/sql"
  "fmt"
  "net/http"

  "github.com/gin-gonic/gin"
  _ "github.com/go-sql-driver/mysql"
)

func main() {
  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/go_sample")
  if err != nil {
    fmt.Print(err.Error())
  }
  defer db.Close()
  // make sure connection is available
  err = db.Ping()
  if err != nil {
    fmt.Print(err.Error())
  }
  type Users struct {
    Id         int    `json:"id"`
    First_Name string `json:"first_name"`
    Last_Name  string `json:"last_name"`
    Email      string `json:"email"`
  }
  router := gin.Default()

  // GET a user detail
  router.GET("/users/:id", func(c *gin.Context) {
    var (
      user Users
      result gin.H
    )
    id := c.Param("id")
    row := db.QueryRow("select id, first_name, last_name, email from users where id = ?;", id)
    err = row.Scan(&user.Id, &user.First_Name, &user.Last_Name, &user.Email)
    if err != nil {
      // If no results send null
      result = gin.H{
        "result": nil,
        "count":  0,
      }
    } else {
      result = gin.H{
        "result": user,
      }
    }
    c.JSON(http.StatusOK, result)
  })

  // GET all users
  router.GET("/users", func(c *gin.Context) {
    var (
      user  Users
      users []Users
    )
    rows, err := db.Query("select id, first_name, last_name, email from users;")
    if err != nil {
      fmt.Print(err.Error())
    }
    for rows.Next() {
      err = rows.Scan(&user.Id, &user.First_Name, &user.Last_Name, &user.Email)
      users = append(users, user)
      if err != nil {
        fmt.Print(err.Error())
      }
    }
    defer rows.Close()
    c.JSON(http.StatusOK, gin.H{
      "result": users,
      "count":  len(users),
    })
  })

  // POST new user details
  router.POST("/users", func(c *gin.Context) {
    var buffer bytes.Buffer
    first_name := c.PostForm("first_name")
    last_name := c.PostForm("last_name")
    email := c.PostForm("email")
    password := c.PostForm("password")

    stmt, err := db.Prepare("insert into users (first_name, last_name, email, password) values(?,?,?,?);")
    if err != nil {
      fmt.Print(err.Error())
    }
    _, err = stmt.Exec(first_name, last_name, email, password)

    if err != nil {
      fmt.Print(err.Error())
    }

    // Fastest way to append strings
    buffer.WriteString(first_name)
    buffer.WriteString(" ")
    buffer.WriteString(last_name)
    defer stmt.Close()
    name := buffer.String()
    c.JSON(http.StatusOK, gin.H{
      "message": fmt.Sprintf(" %s successfully created", name),
    })
  })

  // PUT - update a user details
  router.PUT("/users/:id", func(c *gin.Context) {
    var buffer bytes.Buffer
    id := c.Param("id")
    first_name := c.PostForm("first_name")
    last_name := c.PostForm("last_name")
    password := c.PostForm("password")
    stmt, err := db.Prepare("update users set first_name= ?, last_name= ?, password= ? where id= ?;")
    if err != nil {
      fmt.Print(err.Error())
    }
    _, err = stmt.Exec(first_name, last_name, password, id)
    if err != nil {
      fmt.Print(err.Error())
    }

    // Fastest way to append strings
    buffer.WriteString(first_name)
    buffer.WriteString(" ")
    buffer.WriteString(last_name)
    defer stmt.Close()
    name := buffer.String()
    c.JSON(http.StatusOK, gin.H{
      "message": fmt.Sprintf("Successfully updated to %s", name),
    })
  })

  // Delete resources
  router.DELETE("/users/:id", func(c *gin.Context) {
    id := c.Param("id")
    stmt, err := db.Prepare("delete from users where id= ?;")
    if err != nil {
      fmt.Print(err.Error())
    }
    _, err = stmt.Exec(id)
    if err != nil {
      fmt.Print(err.Error())
    }
    c.JSON(http.StatusOK, gin.H{
      "message": fmt.Sprintf("Successfully deleted users: %s", id),
    })
  })

  // POST login a user
  router.POST("/login", func(c *gin.Context) {
    var (
      user Users
      result gin.H
    )
    email    := c.PostForm("email")
    password := c.PostForm("password")
    row      := db.QueryRow("SELECT id, first_name, last_name, email FROM users WHERE email = ? AND password= ? ORDER BY ID ASC LIMIT 1;", email, password)
    err       = row.Scan(&user.Id, &user.First_Name, &user.Last_Name, &user.Email)
    status   := http.StatusOK

    if err != nil {
      status = http.StatusUnprocessableEntity
      result = gin.H{
        "errors": "Invalid email or password!",
      }
    } else {
      result = gin.H{
        "user": user,
        "login": true,
      }
    }

    c.JSON(status, result)
  })


  // listen
  router.Run(":3000")
}
