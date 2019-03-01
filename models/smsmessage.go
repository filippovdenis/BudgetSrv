package models

type SMSMessage struct {
	MessageId int64 `gorm:"primary_key";"AUTO_INCREMENT"`
	MessageBody string
	DateStr string
}
