package models

type Skill struct {
	Id           int `gorm:"primaryKey;autoIncrement:true;unique"`
	Logro        string
	Description  string
	CompetenceId int
	PersonId     int
}

func GetSkill() (Skill []Skill, err error) {
	DB.Find(&Skill)

	return Skill, nil
}

func SetSkill(skill Skill) (id int, err error) {
	DB.Create(&skill)
	return skill.Id, nil
}
