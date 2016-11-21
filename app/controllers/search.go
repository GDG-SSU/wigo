package controllers

import (
    "log"

    "github.com/revel/revel"
    "github.com/GDG-SSU/wigo/app"
    "github.com/GDG-SSU/wigo/app/models"
    "github.com/asaskevich/govalidator"
)

type Search struct {
    *revel.Controller
}

type SearchDocumentForm struct {
    SearchWord string `valid:"required"`
}

func parseSearchDocumentForm(p *revel.Params) (SearchDocumentForm, error) {
    doc := &SearchDocumentForm{}
    p.Bind(&(doc.SearchWord), "search_word")

    _, err := govalidator.ValidateStruct(doc)
    return *doc, err
}


// Document 검색
func (c Search) Query(page int) revel.Result {
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
        maxNumOfDocument := 10
        maxNumOfPage := 5

        app.DB.Limit(maxNumOfDocument).Table("documents").Select("id, title").Where("Title LIKE ? OR Content LIKE ?", "%" + searchDocumentForm.SearchWord + "%", "%" + searchDocumentForm.SearchWord + "%").Count(&count).Offset(maxNumOfDocument * (page - 1)).Find(&documents)
        c.RenderArgs["searchWord"] = searchDocumentForm.SearchWord      // 검색어

        // Check existing document
        if len(documents) == 0 {
            c.RenderArgs["isDocumentExist"] = false
            return c.RenderTemplate("Document/search_results.html")
        }

        var mod = 0
        if count % maxNumOfDocument > 0 {
            mod = 1
        }
        // Init page number list
        for i := page - (maxNumOfPage >> 1); !(len(pages) >= maxNumOfPage) && i <= count / maxNumOfDocument + mod; i++ {
            if i < 1 {
                continue
            }
            pages = append(pages, i)
        }

        c.RenderArgs["isDocumentExist"] = true      // document 가 존재하는지
        c.RenderArgs["documents"] = documents       // 검색된 document 리스트
        c.RenderArgs["page"] = page                 // 현재 페이지 number        ex) 4
        c.RenderArgs["pages"] = pages               // 표시될 페이지 numbers      ex) 2, 3, 4, 5, 6
        return c.RenderTemplate("Document/search_results.html")
    }
    return c.RenderText("Post is not supported")
}

