package models

type User struct {
	Id int `json:"id"`;
	Name string `json:"name"`;
	Username string `json:"username"`;
	Email string `json:"email"`;
	Password string `json:"password"`;
	Goaltime int `json:"goaltime"`
	RGB []int64 `json:"rgb"`;
	Avatar int `json:avatar;`
}

