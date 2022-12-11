package model

type Movie struct {
	BaseModel
	Type        string  `gorm:"type:varchar(250);column:type" json:"type"`
	Cast        string  `gorm:"type:varchar(250);cast" json:"cast"`
	Name        string  `gorm:"type:varchar(250); name;index" json:"name"`
	Description string  `gorm:"type:varchar;description" json:"description"`
	Duration    float64 `gorm:"duration" json:"duration"`
	ReleaseDate string  `gorm:"type:varchar(100);release_date" json:"release_date"`
	Country     string  `gorm:"type:varchar(100);country" json:"country"`
	Language    string  `gorm:"type:varchar(250);language" json:"language"`
	Rated       string  `gorm:"type:varchar(250);rated" json:"rated"`
	Director    string  `gorm:"type:varchar(250);director" json:"director"`
	Status      string  `gorm:"type:varchar(100);status" json:"status"`
	Poster      string  `gorm:"poster" json:"poster"`
	Trailer     string  `gorm:"trailer" json:"trailer"`
}

func (Movie) TableName() string {
	return "movie"
}
