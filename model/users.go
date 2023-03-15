package model

type Users struct {
	UserId       int    `json:"user_id" gorm:"column:userId;primaryKey;type:int;autoIncrement"`
	UserName     string `json:"username" gorm:"column:userName;type:varchar(256)"`
	UserEmail    string `json:"user_email" gorm:"column:userEmail"`
	UserPassword string `json:"user_password" gorm:"column:userPassword"`
	UserCountry  string `json:"user_country" gorm:"column:userCountry"`
	UserType     int    `json:"user_type" gorm:"column:userType"`
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    Users  `json:"data"`
}
