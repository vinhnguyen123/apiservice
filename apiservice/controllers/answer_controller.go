package controllers

import (
	"apiservice/authorization"
	"apiservice/connections"
	"apiservice/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateAnswer(c *gin.Context) {

	if authorization.CheckAuth(c.Request) != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	db := connections.DBConn()

	var json models.AnswerCreateParam
	var question models.Question

	if err := c.ShouldBindJSON(&json); err == nil {

		result := db.First(&question, json.QuestionID)

		if result.Error != nil {
			c.JSON(500, gin.H{
				"messages": "Question not found",
			})
			return
		}

		answer := models.Answer{AnswerDesc: json.AnswerDesc, CreatedBy: json.CreatedBy}

		result = db.Create(&answer)

		if result.Error != nil {
			c.JSON(500, gin.H{
				"messages": result.Error,
			})
			return
		}

		result = db.Create(&models.QuestionAnswerInfo{QuestionID: json.QuestionID, AnswerID: answer.ID})

		if result.Error != nil {
			c.JSON(500, gin.H{
				"messages": result.Error,
			})
			return
		}

		c.JSON(200, gin.H{
			"messages": "inserted",
		})
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}
}

func UpdateAnswer(c *gin.Context) {

	if authorization.CheckAuth(c.Request) != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	db := connections.DBConn()

	var json models.AnswerUpdateParam

	if err := c.ShouldBindJSON(&json); err == nil {
		answer := models.Answer{AnswerDesc: json.AnswerDesc}

		result := db.Model(&models.Answer{}).Where("id = ?", c.Param("id")).Updates(answer)

		if result.Error != nil {
			c.JSON(500, gin.H{
				"messages": result.Error,
			})
			return
		}

		c.JSON(200, gin.H{
			"messages": "updated",
		})
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}
}

func GetAnswer(c *gin.Context) {

	if authorization.CheckAuth(c.Request) != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	db := connections.DBConn()
	var answer models.Answer

	result := db.First(&answer, c.Param("id"))
	if result.Error != nil {
		c.JSON(500, gin.H{
			"messages": "Answer not found",
		})
		return
	}

	data := map[string]interface{}{
		"Answer": answer,
	}
	c.JSON(200, data)
}

func AddUpVoteAnswer(c *gin.Context) {

	if authorization.CheckAuth(c.Request) != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	db := connections.DBConn()

	answerModel := models.Answer{}

	var json models.AddUpVoteAnswerParam
	result := db.First(&answerModel, c.Param("id"))

	if result.Error != nil {
		c.JSON(500, gin.H{
			"messages": "Answer not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&json); err == nil {
		if json.UpVoteBy == "" || strings.Contains(answerModel.UpVoteBy, json.UpVoteBy) {
			c.JSON(500, gin.H{
				"messages": "User had vote",
			})
			return
		}

		if answerModel.CountUpVote != 0 {
			answerModel = models.Answer{UpVoteBy: answerModel.UpVoteBy + "," + json.UpVoteBy, CountUpVote: answerModel.CountUpVote + 1}
		} else {
			answerModel = models.Answer{UpVoteBy: json.UpVoteBy, CountUpVote: 1}
		}

		result := db.Model(&answerModel).Where("id = ?", c.Param("id")).Updates(answerModel)

		if result.Error != nil {
			c.JSON(500, gin.H{
				"messages": result.Error,
			})
			return
		}

		c.JSON(200, gin.H{
			"messages": "Add up vote success",
		})
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}
}

func AddDownVoteAnswer(c *gin.Context) {

	if authorization.CheckAuth(c.Request) != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	db := connections.DBConn()

	answerModel := models.Answer{}

	var json models.AddDownVoteAnswerParam
	result := db.First(&answerModel, c.Param("id"))

	if result.Error != nil {
		c.JSON(500, gin.H{
			"messages": "Answer not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&json); err == nil {
		if json.DownVoteBy == "" || strings.Contains(answerModel.DownVoteBy, json.DownVoteBy) {
			c.JSON(500, gin.H{
				"messages": "User had vote",
			})
			return
		}

		if answerModel.CountDownVote != 0 {
			answerModel = models.Answer{DownVoteBy: answerModel.DownVoteBy + "," + json.DownVoteBy, CountDownVote: answerModel.CountDownVote + 1}
		} else {
			answerModel = models.Answer{DownVoteBy: json.DownVoteBy, CountDownVote: 1}
		}

		result := db.Model(&answerModel).Where("id = ?", c.Param("id")).Updates(answerModel)

		if result.Error != nil {
			c.JSON(500, gin.H{
				"messages": result.Error,
			})
			return
		}

		c.JSON(200, gin.H{
			"messages": "Add down vote success",
		})
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}
}

func GetAnswerByQuestion(c *gin.Context) {

	if authorization.CheckAuth(c.Request) != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	db := connections.DBConn()

	var answerModel []models.Answer

	result := db.Select("answers.*, ").
		Joins(`INNER JOIN question_answer_infos ON question_answer_infos.answer_id = answers.id
			INNER JOIN questions ON questions.id = question_answer_infos.question_id`).
		Where("questions.id = ?", c.Param("question_id")).
		Find(&answerModel)

	if result.Error != nil {
		c.JSON(500, gin.H{
			"messages": result.Error,
		})
		return
	}

	jsReturn := map[string]interface{}{
		"data":  answerModel,
		"total": len(answerModel),
	}

	c.JSON(http.StatusOK, jsReturn)
}

func DeleteAnswer(c *gin.Context) {

	if authorization.CheckAuth(c.Request) != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	db := connections.DBConn()
	var answer models.Answer

	result := db.Delete(&answer, c.Param("id"))

	if result.Error != nil {
		c.JSON(500, gin.H{
			"messages": result.Error,
		})
		return
	}

	c.JSON(200, gin.H{
		"messages": "deleted",
	})
}
