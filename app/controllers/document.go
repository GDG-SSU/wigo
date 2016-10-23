package controllers

import "github.com/revel/revel"

type Document struct {
    *revel.Controller
}

func (c Document) Write() revel.Result {
    return c.Render()
}
