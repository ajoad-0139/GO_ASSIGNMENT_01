package utils

import (
	"github.com/beego/beego/v2/server/web"
)

func JsonError(c *web.Controller, status int, err interface{}) {
	c.Data["json"] = map[string]interface{}{
		"errors": err,
	}
	c.Ctx.Output.SetStatus(status)
	c.ServeJSON()
}

func JsonSuccess(c *web.Controller, status int, data interface{}) {
	c.Data["json"] = data
	c.Ctx.Output.SetStatus(status)
	c.ServeJSON()
}
