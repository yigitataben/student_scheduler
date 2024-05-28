package initializers

import "github.com/yigitataben/student_scheduler/models"

func SyncDB() {
	err := DB.AutoMigrate(&models.User{}, &models.Plan{}, &models.Lecture{})
	if err != nil {
		return
	}
}
