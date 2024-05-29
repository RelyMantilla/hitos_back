package models

type Pillar struct {
	Id          int `gorm:"primaryKey;autoIncrement:true;unique"`
	Pillar      string
	Description string
}

func GetPillar() (pillar []Pillar, err error) {
	DB.Find(&pillar)
	return pillar, nil
}

func SetPillar(pillar Pillar) (id int, err error) {
	result := DB.Create(&pillar)
	if result.Error != nil {
		return 0, result.Error
	}
	return pillar.Id, nil
}
