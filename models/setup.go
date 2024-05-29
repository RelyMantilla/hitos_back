package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connectmodels() {
	dsn := "host=localhost user=postgres password=root dbname=hitos port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	models, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to models!")
	}
	models.AutoMigrate(Competence{}, Family{}, Person{}, Pillar{}, Skill{})
	models.AutoMigrate(&User{})

	models.AutoMigrate(&Tag{})
	models.Migrator().CreateConstraint(&Tag{}, "Questions")
	models.AutoMigrate(&Question{}, &Answer{}, &Comment{}, &VoteComment{}, &VoteAnswer{}, &VoteQuestion{}, &QuestionTag{})
	DB = models

}
