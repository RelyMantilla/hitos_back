package models

type Person struct {
	Id       int `gorm:"primaryKey;autoIncrement:true;unique"`
	Name     string
	Picture  string
	FamilyId int
	Score    float32
}

type InPerson struct {
	Name string
}

func GetPerson(name string) (person []Person, err error) {
	DB.Where("name like ?", "%"+name+"%").Find(&person)

	return person, nil
}

func GetPersonId(id int) (person Person, err error) {
	DB.Where("ID = ?", id).Find(&person)

	return person, nil
}

func SetPerson(pilar Pillar) (int, error) {

	result := DB.Create(&pilar)
	if result.Error != nil {
		return 0, result.Error
	}
	return pilar.Id, nil
}
