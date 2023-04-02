package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type Person struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}

var people []Person

func main() {
    r := gin.Default()

    r.GET("/people", func(c *gin.Context) {
        c.JSON(http.StatusOK, people)
    })

    r.POST("/people", func(c *gin.Context) {
        var person Person
        if err := c.BindJSON(&person); err != nil {
            c.AbortWithError(http.StatusBadRequest, err)
            return
        }
        people = append(people, person)
        c.JSON(http.StatusOK, person)
    })

    r.Run(":8080")
}
