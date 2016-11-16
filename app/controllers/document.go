package controllers

import (
    "github.com/revel/revel"
    "github.com/GDG-SSU/wigo/app"
    "github.com/GDG-SSU/wigo/app/models"
)

type Document struct {
    *revel.Controller
}

// Document 보기
func (c Document) View(docId int) revel.Result {
    document := models.Document{}
    app.DB.Table("documents").Where("id = ?", docId).Find(&document)
    // TODO: document not found. 404
    c.RenderArgs["document"] = document
    return c.RenderTemplate("Document/document.html")
}
