package controllers

import (
	"apiservice/authorization"
	"apiservice/connections"
	"apiservice/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func CreateQuestion(c *gin.Context) {

	if authorization.CheckAuth(c.Request) != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	db := connections.DBConn()

	var json models.QuestionCreateParam

	if err := c.ShouldBindJSON(&json); err == nil {

		question := models.Question{QuestionDesc: json.QuestionDesc, CreatedBy: json.CreatedBy}

		result := db.Create(&question)

		if result.Error != nil {
			c.JSON(500, gin.H{
				"messages": result.Error,
			})
			return
		}

		if json.Tag != "" {
			tag := models.Tag{TagDesc: strings.ToUpper(json.Tag)}
			result := db.Where("tag_desc = ?", strings.ToUpper(json.Tag)).FirstOrCreate(&tag)
			if result.Error != nil {
				c.JSON(500, gin.H{
					"messages": result.Error,
				})
				return
			}

			questaginfo := models.QuestionTagInfo{QuestionID: question.ID, TagID: tag.ID}

			resultQTI := db.Create(&questaginfo)

			if resultQTI.Error != nil {
				c.JSON(500, gin.H{
					"messages": resultQTI.Error,
				})
				return
			}
		}

		c.JSON(200, gin.H{
			"messages": "inserted",
		})
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}
}

func UpdateQuestion(c *gin.Context) {

	if authorization.CheckAuth(c.Request) != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	db := connections.DBConn()

	var json models.QuestionUpdateParam

	if err := c.ShouldBindJSON(&json); err == nil {
		question := models.Question{QuestionDesc: json.QuestionDesc}

		result := db.Model(&models.Question{}).Where("id = ?", c.Param("id")).Updates(question)

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

func GetQuestion(c *gin.Context) {

	if authorization.CheckAuth(c.Request) != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	db := connections.DBConn()
	var question models.Question

	result := db.First(&question, c.Param("id"))
	if result.Error != nil {
		c.JSON(500, gin.H{
			"messages": "Question not found",
		})
		return
	}

	data := map[string]interface{}{
		"Question": question,
	}
	c.JSON(200, data)
}

func AddUpVoteQuestion(c *gin.Context) {

	if authorization.CheckAuth(c.Request) != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	db := connections.DBConn()

	questionModel := models.Question{}

	var json models.AddUpVoteQuestionParam
	result := db.First(&questionModel, c.Param("id"))

	if result.Error != nil {
		c.JSON(500, gin.H{
			"messages": "Question not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&json); err == nil {
		if json.UpVoteBy == "" || strings.Contains(questionModel.UpVoteBy, json.UpVoteBy) {
			c.JSON(500, gin.H{
				"messages": "User had vote",
			})
			return
		}

		if questionModel.CountUpVote != 0 {
			questionModel = models.Question{UpVoteBy: questionModel.UpVoteBy + "," + json.UpVoteBy, CountUpVote: questionModel.CountUpVote + 1}
		} else {
			questionModel = models.Question{UpVoteBy: json.UpVoteBy, CountUpVote: 1}
		}

		result := db.Model(&questionModel).Where("id = ?", c.Param("id")).Updates(questionModel)

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

func AddDownVoteQuestion(c *gin.Context) {

	if authorization.CheckAuth(c.Request) != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	db := connections.DBConn()

	questionModel := models.Question{}

	var json models.AddDownVoteQuestionParam
	result := db.First(&questionModel, c.Param("id"))

	if result.Error != nil {
		c.JSON(500, gin.H{
			"messages": "Question not found",
		})
		return
	}

	if err := c.ShouldBindJSON(&json); err == nil {
		if json.DownVoteBy == "" || strings.Contains(questionModel.DownVoteBy, json.DownVoteBy) {
			c.JSON(500, gin.H{
				"messages": "User had vote",
			})
			return
		}

		if questionModel.CountDownVote != 0 {
			questionModel = models.Question{DownVoteBy: questionModel.DownVoteBy + "," + json.DownVoteBy, CountDownVote: questionModel.CountDownVote + 1}
		} else {
			questionModel = models.Question{DownVoteBy: json.DownVoteBy, CountDownVote: 1}
		}

		result := db.Model(&questionModel).Where("id = ?", c.Param("id")).Updates(questionModel)

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

func GetQuestionByTag(c *gin.Context) {

	if authorization.CheckAuth(c.Request) != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	db := connections.DBConn()

	var tagDesc models.TagParam
	var questionModel []models.GetQuestionTag

	if err := c.ShouldBindJSON(&tagDesc); err == nil {
		result := db.Select("questions.*, tags.tag_desc AS tag_desc").
			Joins(`INNER JOIN question_tag_infos ON question_tag_infos.question_id = questions.id
			INNER JOIN tags ON question_tag_infos.tag_id = tags.id`).
			Where("tags.tag_desc = ?", strings.ToUpper(tagDesc.TagDesc)).
			Find(&questionModel)

		if result.Error != nil {
			c.JSON(500, gin.H{
				"messages": result.Error,
			})
			return
		}

		jsReturn := map[string]interface{}{
			"data":  questionModel,
			"total": len(questionModel),
		}

		c.JSON(http.StatusOK, jsReturn)
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}
}

func GetListQuestion(c *gin.Context) {

	if authorization.CheckAuth(c.Request) != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	db := connections.DBConn()

	var questionsParam models.QuestionListParam
	var questionModel []models.Question

	if err := c.ShouldBindJSON(&questionsParam); err == nil {
		result := db.Select("questions.*, tags.tag_desc AS tag_desc").
			Joins(`LEFT JOIN question_tag_infos ON question_tag_infos.question_id = questions.id
			LEFT JOIN tags ON question_tag_infos.tag_id = tags.id`).
			Limit(questionsParam.Limit).
			Offset(questionsParam.Offset).
			Order("id DESC").
			Find(&questionModel)

		if result.Error != nil {
			c.JSON(500, gin.H{
				"messages": result.Error,
			})
			return
		}

		jsReturn := map[string]interface{}{
			"data":  questionModel,
			"total": len(questionModel),
		}

		c.JSON(http.StatusOK, jsReturn)
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}
}

func DeleteQuestion(c *gin.Context) {

	if authorization.CheckAuth(c.Request) != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	db := connections.DBConn()
	var question models.Question

	result := db.Delete(&question, c.Param("id"))

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
