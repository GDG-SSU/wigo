package controllers

import (
    "github.com/revel/revel"
    "fmt"
)

func (c App) Document(id int) revel.Result {
    fmt.Print("document")
    if c.Request.Method == "GET" {
        fmt.Print(c.Params)
    }

    return c.Render()
}
