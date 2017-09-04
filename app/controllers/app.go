package controllers

import "github.com/revel/revel"

type App struct {
	GrantBase
}

func (c App) Index() revel.Result {
	return c.Render()
}
