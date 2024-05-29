package models

type Family struct {
	Id   int `gorm:"primaryKey;autoIncrement:true;unique"`
	Name string
}

func GetFamily() (family []Family, err error) {
	DB.Find(&family)
	return family, nil
}

func SetFamily(family Family) (int, error) {

	result := DB.Create(&family)
	if result.Error != nil {
		return 0, result.Error
	}
	return family.Id, nil
}
