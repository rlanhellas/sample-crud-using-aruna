package shared

type Client struct {
	Id        int        `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name"`
	BirthDate string     `json:"birth_date" gorm:"column:birth_date"`
	Languages []Language `gorm:"many2many:client_languages"`
}

type Language struct {
	Id   int `json:"id" gorm:"primaryKey"`
	Name string
}

func NewClient() any {
	return &Client{}
}

func (c *Client) ID() any {
	return c.Id
}

func (c *Client) Clone() any {
	cnew := *c
	return &cnew
}

func (c *Client) TableName() string {
	return "client_table"
}

func NewLanguage() any {
	return &Client{}
}

func (l *Language) ID() any {
	return l.Id
}

func (l *Language) Clone() any {
	cnew := *l
	return &cnew
}

func (l *Language) TableName() string {
	return "language_table"
}
