package models

import (
	"time"
)

type Question struct {
	ID        uint `gorm:"primaryKey;autoIncrement:true;unique"`
	UserID    uint
	Title     string
	Body      string
	View      int
	AnswerID  uint
	Vote      int
	Voted     bool  `gorm:"-"`
	Like      bool  `gorm:"-"`
	TotalAns  int64 `gorm:"-"`
	Tags      []Tag `gorm:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Tag struct {
	ID  uint `gorm:"primaryKey;autoIncrement:true;unique"`
	Tag string
}
type QuestionTag struct {
	QuestionID uint
	TagID      uint
}

type Answer struct {
	ID         uint `gorm:"primaryKey;autoIncrement:true;unique"`
	UserID     uint
	QuestionID int
	Answer     string
	Vote       int
	Voted      bool `gorm:"-"`
	Like       bool `gorm:"-"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Comment struct {
	ID         uint `gorm:"primaryKey;autoIncrement:true;unique"`
	UserID     uint
	AnswerID   uint
	QuestionID uint
	Comment    string
	Vote       int
	Voted      bool `gorm:"-"`
	Like       bool `gorm:"-"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type VoteComment struct {
	ID        uint `gorm:"primaryKey;autoIncrement:true;unique"`
	UserID    uint
	CommentID uint
	Like      bool
}

type VoteAnswer struct {
	ID       uint `gorm:"primaryKey;autoIncrement:true;unique"`
	UserID   uint
	AnswerID uint
	Like     bool
}

type VoteQuestion struct {
	ID         uint `gorm:"primaryKey;autoIncrement:true;unique"`
	UserID     uint
	QuestionID uint
	Like       bool
}

type Likes struct {
	CommentID  uint
	AnswerID   uint
	QuestionID uint
	Like       bool
}
type Solved struct {
	QuestionID uint
	AnswerID   uint
}

func GetTags() (tags []Tag, err error) {
	DB.Find(&tags)
	return tags, nil
}

func GetTag(tagId uint) (tag Tag, err error) {
	DB.Where("id = ?", tagId).Find(&tag)
	return tag, nil
}

func CreateTag(tag Tag) (id uint, err error) {
	result := DB.Create(&tag)
	if result.Error != nil {
		return 0, result.Error
	}
	return tag.ID, nil
}

func CreateQuestion(question Question) (id uint, err error) {
	//title, body string, tags []string
	result := DB.Create(&question)
	if result.Error != nil {
		return 0, result.Error
	}
	return question.ID, nil
}

func CreateAnswer(answer Answer) (id uint, err error) {
	result := DB.Create(&answer)
	if result.Error != nil {
		return 0, result.Error
	}
	return answer.ID, nil
}

func CreateComment(comment Comment) (id uint, err error) {
	result := DB.Create(&comment)
	if result.Error != nil {
		return 0, result.Error
	}
	return comment.ID, nil
}

func GetQuestions() (questions []Question, err error) {
	DB.Find(&questions)
	return questions, nil
}

func GetQuestionsByTag(tagId int) (questions []Question, err error) {
	var questionTags []QuestionTag
	DB.Where("tag_id = ?", tagId).Find(&questionTags)
	var ids []uint
	for _, questionTag := range questionTags {
		ids = append(ids, questionTag.QuestionID)
	}

	DB.Where("id = ?", ids).Find(&questions)
	return questions, nil

}

func GetQuestion(id uint) (question Question, err error) {
	DB.Where("id = ?", id).Find(&question)
	return question, nil
}

func UpdateQuestionView(question Question) (view int, err error) {
	if question.ID != 0 {
		DB.Save(&question)
	}
	return question.View, nil
}

func GetAnswers(questionId int) (answers []Answer, err error) {
	DB.Where("question_id = ?", questionId).Find(&answers)
	return answers, nil
}

func GetAnswer(answerID int) (answer Answer, err error) {
	DB.Where("id = ?", answerID).Find(&answer)
	return answer, nil
}
func GetTotalAnswers(questionId uint) (total int64, err error) {
	var answer Answer
	//DB.Where("question_id = ?", questionId).Find(&answers)
	DB.Model(&answer).Where("question_id = ?", questionId).Count(&total)
	return total, nil
}

func GetComments(questionId int, answerId int) (questions []Comment, err error) {
	DB.Where("question_id = ? and answer_id = ?", questionId, answerId).Find(&questions)
	return questions, nil
}

func GetComment(commentId int) (comment Comment, err error) {
	DB.Where("id = ?", commentId).Find(&comment)
	return comment, nil

}

func GetVoteQuestion(questionid uint) (question Question, err error) {
	DB.Where("id = ?", questionid).Find(&question)
	return question, nil
}

func GetVoteAnswer(answerId uint) (answer Answer, err error) {
	DB.Where("id = ?", answerId).Find(&answer)
	return answer, nil
}

func GetVoteComment(commentId uint) (comment Comment, err error) {
	DB.Where("id = ?", commentId).Find(&comment)
	return comment, nil
}

func GetUserVoteQuestion(questionID uint, userID uint) (question VoteQuestion, err error) {
	DB.Where("question_id = ? and user_id = ?", questionID, userID).Find(&question)

	return question, nil
}

func GetUserVoteAnswer(answerId uint, userID uint) (answer VoteAnswer, err error) {
	DB.Where("answer_id = ? and user_id = ?", answerId, userID).Find(&answer)
	return answer, nil
}

func GetUserVoteComment(commentId uint, userID uint) (comment VoteComment, err error) {
	DB.Where("comment_id = ? and user_id = ?", commentId, userID).Find(&comment)
	return comment, nil
}

func UpdateQuestion(question Question) (id uint, err error) {
	if question.ID != 0 {
		DB.Save(&question)
	}

	return question.ID, nil
}

func UpdateAnswer(answer Answer) (vote int, err error) {
	if answer.ID != 0 {
		DB.Save(&answer)
	}
	return answer.Vote, nil
}

func UpdateComment(comment Comment) (vote int, err error) {
	if comment.ID != 0 {
		DB.Save(&comment)
	}
	return comment.Vote, nil
}

func UpdateVoteQuestion(question VoteQuestion) (id uint, err error) {
	if question.ID != 0 {
		DB.Save(&question)
	}

	return question.ID, nil
}

func UpdateVoteAnswer(answer VoteAnswer) (id uint, err error) {
	if answer.ID != 0 {
		DB.Save(&answer)
	}
	return answer.ID, nil
}

func UpdateVoteComment(comment VoteComment) (id uint, err error) {
	if comment.ID != 0 {
		DB.Save(&comment)
	}
	return comment.ID, nil
}

func DeleteVoteQuestion(question VoteQuestion) (id uint, err error) {
	if question.ID != 0 {
		DB.Delete(&question)
	}

	return question.ID, nil
}

func DeleteVoteAnswer(answer VoteAnswer) (id uint, err error) {
	if answer.ID != 0 {
		DB.Delete(&answer)
	}
	return answer.ID, nil
}

func DeleteVoteComment(comment VoteComment) (id uint, err error) {
	if comment.ID != 0 {
		DB.Delete(&comment)
	}
	return comment.ID, nil
}

func SetVoteQuestion(voteQuestion VoteQuestion) (id uint, err error) {
	result := DB.Create(&voteQuestion)
	if result.Error != nil {
		return 0, result.Error
	}
	return voteQuestion.ID, nil
}

func SetVoteAnswer(voteAnswer VoteAnswer) (id uint, err error) {
	result := DB.Create(&voteAnswer)
	if result.Error != nil {
		return 0, result.Error
	}
	return voteAnswer.ID, nil
}

func SetVoteComment(voteComment VoteComment) (id uint, err error) {
	result := DB.Create(&voteComment)
	if result.Error != nil {
		return 0, result.Error
	}
	return voteComment.ID, nil
}

func CreateQuestionTag(questionTag QuestionTag) error {
	result := DB.Create(&questionTag)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetQuestionTags(questionID uint) (tags []Tag, err error) {
	var questionTags []QuestionTag
	result := DB.Where("question_id = ?", questionID).Find(&questionTags)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, questionTag := range questionTags {
		tag, _ := GetTag(questionTag.TagID)
		tags = append(tags, tag)

	}
	return tags, nil
}
