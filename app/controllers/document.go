package controllers

import (
    "log"

    "github.com/revel/revel"
    "github.com/GDG-SSU/wigo/app"
    "github.com/GDG-SSU/wigo/app/models"
    "github.com/asaskevich/govalidator"
)

type Document struct {
    *revel.Controller
}

/**
 * DocumentForm - A struct representing POST form payload
 */
type DocumentForm struct {
    Title   string  `valid:"required"`
    Content string  `valid:"required"`
}

func parseForm(p *revel.Params) (DocumentForm, error) {
    doc := &DocumentForm{}
    p.Bind(&(doc.Title), "title")
    p.Bind(&(doc.Content), "content")

    _, err := govalidator.ValidateStruct(doc)
    return *doc, err
}

func (c Document) Write() revel.Result {
    // When the submit button is clicked
    if c.Request.Method == "POST" {
        df, err := parseForm(c.Params)
        if (err == nil) {
            // Save to database
            app.DB.Create(&models.Document{Title: df.Title, Content: df.Content})
        } else {
            // TEMP: kill the application if there is an form-parsing error
            log.Fatal(err)
        }
    }

    return c.Render()
}
