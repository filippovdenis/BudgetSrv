package main

type Entity interface {

}

type Category struct {
	CategoryId int64 `gorm:"primary_key";"AUTO_INCREMENT"`
	CategoryName string
}
