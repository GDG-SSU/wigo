package controllers

import (
    "github.com/revel/revel"
    "github.com/asaskevich/govalidator"
    "github.com/GDG-SSU/wigo/app/models"
    "github.com/GDG-SSU/wigo/app"
    "log"
)

type SearchDocumentForm struct {
    Title string `valid:"required"`
}

func parseSearchDocumentForm(p *revel.Params) (SearchDocumentForm, error) {
    doc := &SearchDocumentForm{}
    p.Bind(&(doc.Title), "title")

    _, err := govalidator.ValidateStruct(doc)
    return *doc, err
}

func (c Document) Search() revel.Result {
    // When submit button is clicked
    if c.Request.Method == "GET" {
        searchDocumentForm, err := parseSearchDocumentForm(c.Params)
        if err != nil {
            // TODO: present another page
            log.Fatal(err)
            return c.RenderHtml("Fail")
        }

        var document models.Document
        app.DB.First(&document, "Title = ?", searchDocumentForm.Title)
        if document.ID == 0 {
            return c.RenderHtml("Row not Found")
        }
        return c.RenderHtml("Success")
    }
    return c.RenderHtml("Post is not supported")
}
