package controllers

import (
    "github.com/revel/revel"
    "log"
    "github.com/GDG-SSU/wigo/app/models"
    "github.com/GDG-SSU/wigo/app/routes"
    "github.com/asaskevich/govalidator"
    "github.com/GDG-SSU/wigo/app"
    "fmt"
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


// Document 검색
func (c Document) Search(page int) revel.Result {
    // When submit button is clicked
    if c.Request.Method == "GET" {
        searchDocumentForm, err := parseSearchDocumentForm(c.Params)
        if err != nil {
            log.Fatal(err)
            return c.RenderText("Fail to DB Query")
        }

        var documents []models.Document
        var pages []int
        count := 0
        maxNumOfDocument := 1
        maxNumOfPage := 5

        app.DB.Limit(maxNumOfDocument).Table("documents").Select("id, title").Where("Title LIKE ? OR Content LIKE ?", "%" + searchDocumentForm.Title + "%", "%" + searchDocumentForm.Title + "%").Count(&count).Offset(maxNumOfDocument * (page - 1)).Find(&documents)

        // Check existing document
        if len(documents) == 0 {
            return c.Redirect(routes.Document.Write())
        }

        // Init page list
        for i := page - (maxNumOfPage >> 1); !(len(pages) >= maxNumOfPage) && i <= count / maxNumOfDocument; i++ {
            if i < 1 {
                continue
            }
            pages = append(pages, i)
        }

        c.RenderArgs["documents"] = documents
        c.RenderArgs["searchTitle"] = searchDocumentForm.Title
        c.RenderArgs["page"] = page
        c.RenderArgs["pages"] = pages
        return c.RenderTemplate("Document/search_results.html")
    }
    return c.RenderText("Post is not supported")
}

