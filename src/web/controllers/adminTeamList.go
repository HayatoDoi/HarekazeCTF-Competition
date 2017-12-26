package controllers

import (
	"os"

	"../models/TeamModel"
	"./BaseController"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
)

// AdminTeamList override BaseController
type AdminTeamList struct {
	BaseController.Base
}

// Get handles GET: http://localhost:8080/<APP_ADMIN_HASH>/team.
// Display team list
func (c *AdminTeamList) Get() mvc.Result {
	teamModel := TeamModel.New()
	teams, err := teamModel.All()
	if err != nil {
		return mvc.Response{Err: err, Code: 500}
	}
	return mvc.View{
		Name: "admin/teamList.html",
		Data: context.Map{
			"Title":     "Team List",
			"Teams":     teams,
			"AdminHash": os.Getenv("APP_ADMIN_HASH"),
		},
	}
}
