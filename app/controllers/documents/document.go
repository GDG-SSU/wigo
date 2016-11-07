package controllers

import (
    "github.com/revel/revel"
    "github.com/GDG-SSU/wigo/app"
    "github.com/GDG-SSU/wigo/app/models"
)

type Document struct {
    *revel.Controller
}

type DocumentForm struct {
    ID  int `valid:"required"`
}

func parseDocumentForm(params *revel.Params) (DocumentForm, error) {
    doc := DocumentForm{}
    params.Bind(&(doc.ID), "id")
    return doc, nil
}


// Document 보기
func (c Document) Document() revel.Result {
    if c.Request.Method == "GET" {
        documentForm, err := parseDocumentForm(c.Params)
        if err != nil {
            return c.RenderText("Document Parsing Error")
        }

        document := models.Document{}
        app.DB.Table("documents").Where("id = ?", documentForm.ID).Find(&document)
        c.RenderArgs["document"] = document
        return c.RenderTemplate("Document/document.html")
    }

    return c.Render()
}
