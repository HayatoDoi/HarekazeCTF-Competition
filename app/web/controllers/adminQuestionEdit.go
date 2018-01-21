package controllers

import (
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/HayatoDoi/HarekazeCTF-Competition/app/datamodels/QuestionModel"
	"github.com/HayatoDoi/HarekazeCTF-Competition/app/web/controllers/BaseController"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
)

// AdminQuestionEdit override BaseController
type AdminQuestionEdit struct {
	BaseController.Base
}

// GetBy handles GET: http://localhost:8080/<APP_ADMIN_HASH>/question/edit/<question id>.
func (c *AdminQuestionEdit) GetBy(questionId int) mvc.Result {
	if !c.IsLoggedIn() {
		c.SetRedirectPath(fmt.Sprintf("/%s/question/edit/%d", os.Getenv("APP_ADMIN_HASH"), questionId))
		return mvc.Response{Path: "/user/login"}
	}

	questionModel := QuestionModel.New()
	question, err := questionModel.FindId(questionId)
	if err != nil {
		return mvc.Response{Err: err, Code: 500}
	}
	return mvc.View{
		Name: "admin/questionEditForm.html",
		Data: context.Map{
			"Title":     "Question Edit",
			"AdminHash": os.Getenv("APP_ADMIN_HASH"),
			"Question":  question,
			"Token":     c.MakeToken(fmt.Sprintf("/%s/question/edit/%d", os.Getenv("APP_ADMIN_HASH"), questionId)),
		},
	}
}

// PostBy handles GET: http://localhost:8080/<APP_ADMIN_HASH>/question/edit/<question id>.
func (c *AdminQuestionEdit) PostBy(questionId int) mvc.Result {
	if !c.IsLoggedIn() {
		c.SetRedirectPath(fmt.Sprintf("/%s/question/edit/%d", os.Getenv("APP_ADMIN_HASH"), questionId))
		return mvc.Response{Path: "/user/login"}
	}

	var (
		name               = c.Ctx.FormValue("name")
		flag               = c.Ctx.FormValue("flag")
		score              = c.Ctx.FormValue("score")
		genre              = c.Ctx.FormValue("genre")
		publish_start_time = c.Ctx.FormValue("publish_start_time")
		publish_now        = c.Ctx.FormValue("publish_now")
		sentence           = c.Ctx.FormValue("sentence")
		token              = c.Ctx.FormValue("csrf_token")
	)
	if !c.CheckTaken(token, fmt.Sprintf("/%s/question/edit/%d", os.Getenv("APP_ADMIN_HASH"), questionId)) {
		err := errors.New("token error!!")
		return mvc.Response{Err: err, Code: 400}
	}
	if publish_now == "off" && !regexp.MustCompile(`/^(\d{4})-(\d{2})-(\d{2})\s(\d{2}):(\d{2}):(\d{2})$/`).MatchString(publish_start_time) {
		err := errors.New("publish_start_time is yyyy-MM-dd HH:mm:ss!!")
		return mvc.Response{Err: err, Code: 500}
	} else if !regexp.MustCompile(`^[0-9]+$`).MatchString(score) {
		err := errors.New("Score is number!!")
		return mvc.Response{Err: err, Code: 500}
	}
	questionModel := QuestionModel.New()
	err := questionModel.Update(questionId, map[string]string{
		"name":               name,
		"flag":               flag,
		"score":              score,
		"genre":              genre,
		"publish_now":        publish_now,
		"publish_start_time": publish_start_time,
		"sentence":           sentence,
	})
	if err != nil {
		return mvc.Response{Err: err, Code: 500}
	}
	return mvc.Response{
		Path: "/" + os.Getenv("APP_ADMIN_HASH"),
	}
}