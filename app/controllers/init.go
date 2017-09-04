package controllers

import (
	"github.com/revel/revel"
)

func init() {
	revel.InterceptMethod((*Base).After, revel.AFTER)
	revel.InterceptMethod((*Base).Before, revel.BEFORE)
	revel.InterceptMethod((*GrantBase).Before, revel.BEFORE)
}
