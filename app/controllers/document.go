package controllers

import (
    "github.com/revel/revel"
    "github.com/GDG-SSU/wigo/app"
    "github.com/GDG-SSU/wigo/app/models"
)

type Document struct {
    *revel.Controller
}

func (c Document) Write() revel.Result {

    // When the submit button is clicked
    if c.Request.Method == "POST" {
        var title string
        var content string
        c.Params.Bind(&title, "title")
        c.Params.Bind(&content, "content")

        // Save to database
        app.DB.Create(&models.Document{Title: title, Content: content})
    }

    return c.Render()
}
