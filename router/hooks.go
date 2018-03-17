package router

import (
	"net/http"

	"github.com/IdanLoo/gopack/util/project"

	"github.com/IdanLoo/gopack/model"
	"github.com/gin-gonic/gin"
)

// HooksGroup is a router group
var HooksGroup = Router.Group("/hooks")

func push(ctx *gin.Context) {
	body := &model.PushBody{}

	if err := ctx.BindJSON(body); err != nil {
		ctx.JSON(http.StatusOK, err)
		return
	}

	err := project.Clone(body.Repository.Name, body.Branch(), body.Repository.URL)
	if err != nil {
		ctx.JSON(http.StatusOK, err)
		return
	}

	ctx.JSON(http.StatusOK, body)
}

func init() {
	HooksGroup.
		POST("/push", push)
}
