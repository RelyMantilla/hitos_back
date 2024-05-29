package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hitos_back/models"
	"hitos_back/utils"
	"net/http"
	"strconv"
)

func GetTags(c *gin.Context) {
	tags, err := models.GetTags()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tags)
}

func GetTag(c *gin.Context) {
	tagid := c.Query("id")
	id, _ := strconv.Atoi(tagid)
	competence, err := models.GetTag(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, competence)
}

func SetTag(c *gin.Context) {
	var inJson models.Tag

	if err := c.ShouldBindJSON(&inJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := models.CreateTag(inJson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, id)
}

func SetQuestion(c *gin.Context) {
	userId, _ := utils.ExtractTokenID(c)
	var inJson models.Question
	inJson.UserID = userId

	if err := c.ShouldBindJSON(&inJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(inJson)

	id, err := models.CreateQuestion(inJson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, tag := range inJson.Tags {
		models.CreateQuestionTag(models.QuestionTag{QuestionID: id, TagID: tag.ID})
	}

	c.JSON(http.StatusOK, id)
}

func SetAnswer(c *gin.Context) {
	var inJson models.Answer

	if err := c.ShouldBindJSON(&inJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := models.CreateAnswer(inJson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, id)
}

func SetComment(c *gin.Context) {
	var inJson models.Comment

	if err := c.ShouldBindJSON(&inJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, err := models.CreateComment(inJson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, id)
}

func GetQuestions(c *gin.Context) {
	parmId := c.Query("tagID")
	id, _ := strconv.Atoi(parmId)
	var questions []models.Question
	var err error
	if id != 0 {
		questions, err = models.GetQuestionsByTag(id)
	} else {
		questions, err = models.GetQuestions()
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, question := range questions {
		question.TotalAns, _ = models.GetTotalAnswers(question.ID)

		question.Tags, _ = models.GetQuestionTags(question.ID)
		questions[i] = question
	}
	c.JSON(http.StatusOK, questions)
}

func GetQuestion(c *gin.Context) {
	parmId := c.Query("ID")
	id, _ := strconv.Atoi(parmId)
	question, err := models.GetQuestion(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	question.View += 1
	_, err = models.UpdateQuestionView(question)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	question.TotalAns, _ = models.GetTotalAnswers(question.ID)
	question.Tags, _ = models.GetQuestionTags(question.ID)
	userId, _ := utils.ExtractTokenID(c)
	vote, _ := models.GetUserVoteQuestion(question.ID, userId)
	if vote.ID != 0 {
		question.Voted = true
		question.Like = vote.Like
	} else {
		question.Voted = false
	}
	c.JSON(http.StatusOK, question)
}
func QuestionSolved(c *gin.Context) {

	var solved models.Solved

	if err := c.ShouldBindJSON(&solved); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	question, _ := models.GetQuestion(solved.QuestionID)
	question.AnswerID = solved.AnswerID
	id, err := models.UpdateQuestion(question)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, id)
}
func GetAnswers(c *gin.Context) {
	parmId := c.Query("questionID")
	id, _ := strconv.Atoi(parmId)
	competence, err := models.GetAnswers(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, competence)
}

func GetAnswer(c *gin.Context) {
	parmId := c.Query("ID")
	id, _ := strconv.Atoi(parmId)
	answer, err := models.GetAnswer(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId, _ := utils.ExtractTokenID(c)
	vote, _ := models.GetUserVoteAnswer(answer.ID, userId)
	if vote.ID != 0 {
		answer.Voted = true
		answer.Like = vote.Like
	} else {
		answer.Voted = false
	}
	c.JSON(http.StatusOK, answer)
}

func GetComments(c *gin.Context) {
	parmId := c.Query("questionID")
	questionID, _ := strconv.Atoi(parmId)
	parm2Id := c.Query("answerID")
	answerID, _ := strconv.Atoi(parm2Id)
	competence, err := models.GetComments(questionID, answerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, competence)
}

func GetComment(c *gin.Context) {
	parmId := c.Query("ID")
	commnetID, _ := strconv.Atoi(parmId)

	comment, err := models.GetComment(commnetID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId, _ := utils.ExtractTokenID(c)
	vote, _ := models.GetUserVoteComment(comment.ID, userId)
	if vote.ID != 0 {
		comment.Voted = true
		comment.Like = vote.Like
	} else {
		comment.Voted = false
	}
	c.JSON(http.StatusOK, comment)
}

func UpdateQuestion(c *gin.Context) {

}

func UpdateAnswers(c *gin.Context) {

}

func UpdateComments(c *gin.Context) {

}

func Like(c *gin.Context) {
	var likes models.Likes

	if err := c.ShouldBindJSON(&likes); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId, _ := utils.ExtractTokenID(c)
	if likes.CommentID != 0 {
		err := likeComment(likes.Like, likes.CommentID, userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	if likes.AnswerID != 0 {
		err := likeAnswer(likes.Like, likes.AnswerID, userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	if likes.QuestionID != 0 {
		err := likeQuestion(likes.Like, likes.QuestionID, userId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, "asd:ok")
}

func likeComment(like bool, commentID uint, userId uint) error {
	vote, err := models.GetUserVoteComment(commentID, userId)

	if vote.ID != 0 {
		if vote.Like == like {
			like = !like
		}
		models.DeleteVoteComment(vote)
	} else {
		models.SetVoteComment(models.VoteComment{CommentID: commentID, UserID: userId, Like: like})
	}

	comment, _ := models.GetVoteComment(commentID)
	if like {
		comment.Vote += 1
	} else {
		comment.Vote -= 1
	}

	_, err = models.UpdateComment(comment)
	if err != nil {
		return err
	}
	return nil
}

func likeAnswer(like bool, answerID uint, userId uint) error {
	vote, err := models.GetUserVoteAnswer(answerID, userId)
	fmt.Println(vote, like)
	if vote.ID != 0 {
		if vote.Like == like {
			like = !like
		}
		models.DeleteVoteAnswer(vote)
	} else {
		models.SetVoteAnswer(models.VoteAnswer{AnswerID: answerID, UserID: userId, Like: like})
	}

	Answer, _ := models.GetVoteAnswer(answerID)
	if like {
		Answer.Vote += 1
	} else {
		Answer.Vote -= 1
	}

	_, err = models.UpdateAnswer(Answer)
	if err != nil {
		return err
	}
	return nil
}
func likeQuestion(like bool, questionID uint, userId uint) error {
	vote, err := models.GetUserVoteQuestion(questionID, userId)
	if vote.ID != 0 {
		if vote.Like == like {
			like = !like
		}
		models.DeleteVoteQuestion(vote)
	} else {
		models.SetVoteQuestion(models.VoteQuestion{QuestionID: questionID, UserID: userId, Like: like})
	}

	Question, _ := models.GetVoteQuestion(questionID)
	if like {
		Question.Vote += 1
	} else {
		Question.Vote -= 1
	}
	_, err = models.UpdateQuestion(Question)
	if err != nil {
		return err
	}
	return nil
}
