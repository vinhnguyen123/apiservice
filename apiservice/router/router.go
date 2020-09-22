package router

import (
	"apiservice/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// r.Static("/public", "./public")
	user := r.Group("/user")
	{
		user.GET("/:id", controllers.GetUser)
		user.POST("/create", controllers.CreateUser)
		user.PATCH("/update/:id", controllers.UpdateUser)
		user.DELETE("/:id", controllers.DeleteUser)
	}

	question := r.Group("/question")
	{
		question.POST("/create", controllers.CreateQuestion)
		question.PATCH("/addupvote/:id", controllers.AddUpVoteQuestion)
		question.PATCH("/adddownvote/:id", controllers.AddDownVoteQuestion)
		question.DELETE("/:id", controllers.DeleteQuestion)
		question.POST("/getquestionbytag", controllers.GetQuestionByTag)
		question.POST("/getlistquestion", controllers.GetListQuestion)
	}

	answer := r.Group("/answer")
	{
		answer.POST("/create", controllers.CreateAnswer)
		answer.PATCH("/update/id", controllers.UpdateAnswer)
		answer.PATCH("/addupvote/:id", controllers.AddUpVoteAnswer)
		answer.PATCH("/adddownvote/:id", controllers.AddDownVoteAnswer)
		answer.DELETE("/:id", controllers.DeleteAnswer)
		answer.GET("/getanswerbyquestion/:question_id", controllers.GetAnswerByQuestion)
	}

	r.POST("/login", controllers.Login)
	r.POST("/logout", controllers.Logout)

	return r
}
