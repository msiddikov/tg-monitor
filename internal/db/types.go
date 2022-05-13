package db

type (
	Monitor struct {
		ChatId   int64
		Id       int
		Name     string
		Url      string
		Method   string
		Body     string
		Interval string
	}

	User struct {
		ChatId int64
		Name   string
		Number string
	}
)
