package controllers

import (
	"fmt"
	"os"

	"github.com/TeamHarekaze/HarekazeCTF2018-server/datamodels/TeamModel"
	"github.com/TeamHarekaze/HarekazeCTF2018-server/web/controllers/BaseController"
	"github.com/kataras/iris/mvc"
)

// AdminTeamEnable override BaseController
type AdminTeamEnable struct {
	BaseController.Base
}

// GetBy handles GET: http://localhost:8080/<APP_ADMIN_HASH>/team/enable/<team id>.
func (c *AdminTeamEnable) GetBy(teamId int) mvc.Result {
	if !c.IsLoggedIn() {
		c.SetRedirectPath(fmt.Sprintf("/%s/team/enable/%d", os.Getenv("APP_ADMIN_HASH"), teamId))
		return mvc.Response{Path: "/user/login"}
	}

	teamModel := TeamModel.New()
	err := teamModel.Enable(teamId)
	if err != nil {
		return c.Error(err)
	}
	return mvc.Response{
		Path: "/" + os.Getenv("APP_ADMIN_HASH") + "/team",
	}
}
