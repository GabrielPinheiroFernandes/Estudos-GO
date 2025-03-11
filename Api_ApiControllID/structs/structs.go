package structs

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Pass       string `json:"pass"`
	Img        string `json:"img"` 
	Card_value string `json:"card_value"`
}
