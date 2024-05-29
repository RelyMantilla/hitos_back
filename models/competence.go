package models

type Competence struct {
	Id          int `gorm:"primaryKey;autoIncrement:true;unique"`
	Name        string
	Description string
	PillarID    int
}

func GetCompetence(id int) (competence []Competence, err error) {

	DB.Where("pillar_id = ?", id).Find(&competence)

	return competence, nil
}

func SetCompetence(competence Competence) (id int, err error) {

	DB.Create(&competence)

	return competence.Id, nil
}
