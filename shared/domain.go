package shared

type Client struct {
	Id        int    `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	BirthDate string `json:"birth_date" gorm:"column:birth_date"`
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
