package controllers

import (
    "github.com/revel/revel"
    "github.com/asaskevich/govalidator"
    "github.com/GDG-SSU/wigo/app/models"
    "github.com/GDG-SSU/wigo/app"
    "log"
    "github.com/GDG-SSU/wigo/app/routes"
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
            log.Fatal(err)
            return c.RenderText("Fail to DB Query")
        }

        var documents []models.Document

        app.DB.Where("Title LIKE ? OR Content LIKE ?", searchDocumentForm.Title, searchDocumentForm.Title).Find(&documents)

        // Check existing document
        if len(documents) == 0 {
            return c.Redirect(routes.Document.Write())
        }
        c.RenderArgs["documents"] = documents
        return c.RenderTemplate("Document/search_results.html")
    }
    return c.RenderText("Post is not supported")
}
