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
type WriteForm struct {
    Title   string  `valid:"required"`
    Content string  `valid:"required"`
}

func parseForm(p *revel.Params) (WriteForm, error) {
    doc := &WriteForm{}
    p.Bind(&(doc.Title), "title")
    p.Bind(&(doc.Content), "content")

    _, err := govalidator.ValidateStruct(doc)
    return *doc, err
}

// Document 보기
func (c Document) View(docId int) revel.Result {
    document := models.Document{}
    app.DB.Table("documents").Where("id = ?", docId).Find(&document)
    // TODO: document not found. 404
    c.RenderArgs["document"] = document
    return c.RenderTemplate("Document/document.html")
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
            // TODO: Handle the submit or database error
            log.Fatal(err)
        }
    }

    return c.Render()
}
