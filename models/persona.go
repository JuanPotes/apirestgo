package models

type Persona struct {
	ID        int64  `json:"id" gorm:"primary_key;auto_increment"`
	TypeID    string `json:"typeid"`
	Doc       string `json:"doc"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Country   string `json:"country"`
}
