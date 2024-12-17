package models

type Notifications struct {
	NotificationId int    `json:"notification_id"`
	UserId         int    `json:"user_id"`
	Status         string `json:"status"`
}

type CommercialNotifications struct {
	NotificationId int    `json:"notification_id"`
	Message        string `json:"message"`
}

type LateNotifications struct {
	NotificationId int `json:"notification_id"`
	BookId         int `json:"book_id"`
}
