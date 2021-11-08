package models

type Soul struct {
	Id       int
	Content  string
	ImageUrl string `orm:"type(text)"`
}

func GetContentRand() {
}
