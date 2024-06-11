package models

import "time"

type Daily struct {
	Id     int `gorm:"primaryKey;autoIncrement:true;unique"`
	UserID uint
	Daily  string
	Fecha  time.Time

	// CreatedAt time.Time
	// UpdatedAt time.Time
}

func GetDailys() (daily []Daily, err error) {
	DB.Find(&daily)
	// DB.Where("UserId = ? and Fecha = ?", userId, fecha).Find(&daily)

	return daily, nil
}

func GetDaily(fecha time.Time, userId uint) (daily []Daily, err error) {
	// DB.Find(&daily)
	DB.Where("UserId = ? and Fecha = ?", userId, fecha).Find(&daily)

	return daily, nil
}

// func GetTag(tagId uint) (tag Tag, err error) {
// 	DB.Where("id = ?", tagId).Find(&tag)
// 	return tag, nil
// }

func SetDaily(daily Daily) (int, error) {

	result := DB.Create(&daily)
	if result.Error != nil {
		return 0, result.Error
	}
	return daily.Id, nil
}
