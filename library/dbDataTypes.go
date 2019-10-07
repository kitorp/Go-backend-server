package library

type DBData struct {
	UserId   int    `json:"Email"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
	UserType int
	deleted  int
	Token    string `json:"Token"`
	Resource []byte
	Quota    int
}
