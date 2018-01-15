package controllers

import (
	"fmt"
	"os"

	"../models/QuestionModel"
	"./BaseController"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
)

// AdminQuestionList override BaseController
type AdminQuestionList struct {
	BaseController.Base
}

// Get handles GET: http://localhost:8080/<APP_ADMIN_HASH>/question.
// Display question list
func (c *AdminQuestionList) Get() mvc.Result {
	if !c.IsLoggedIn() {
		c.SetRedirectPath(fmt.Sprintf("/%s/question", os.Getenv("APP_ADMIN_HASH")))
		return mvc.Response{Path: "/user/login"}
	}

	questionModel := QuestionModel.New()
	questions, err := questionModel.FindAll()
	if err != nil {
		return mvc.Response{Err: err}
	}

	return mvc.View{
		Name: "admin/questionList.html",
		Data: context.Map{
			"Title":     "Question List",
			"Questions": questions,
			"AdminHash": os.Getenv("APP_ADMIN_HASH"),
		},
	}
}