package controllers

import "github.com/revel/revel"

type Document struct {
    *revel.Controller
}

func (c Document) Write() revel.Result {
    var title string
    var content string
    c.Params.Bind(&title, "title")
    c.Params.Bind(&content, "content")
    return c.Render()
}
