package library

type LoginRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type LoginResponse struct {
	Success bool   `json:"Success"`
	Token   string `json:"Token"`
	UserID  int    `json:"UserID"`
	Error   string `json:"Error"`
}

type CreateResourceRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
	UserID   int    `json:"UserID"`
	Resource string `json:"Resource"`
}

type CommonResponse struct {
	Success bool   `json:"Success"`
	Error   string `json:"Error"`
}

type ListResourceRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
	UserID   int    `json:"UserID"`
}

type ListResourceResponse struct {
	Resource []string `json:"Resource"`
	Error    string   `json:"Error"`
}

type DeleteResourceRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
	UserID   int    `json:"UserID"`
	Resource string `json:"Resource"`
}

type SetQuotaRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
	UserID   int    `json:"UserID"`
	Quota    int    `json:"Quota"`
}

type CreateUserRequest struct {
	Email        string `json:"Email"`
	Password     string `json:"Password"`
	Token        string `json:"Token"`
	UserEmail    string `json:"UserEmail"`
	UserPassword string `json:"UserPassword"`
}

type ListUserRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
	Limit    int    `json:"Limit"`
	Offset   int    `json:"Offset"`
}

type User struct {
	UserID        int      `json:"UserID"`
	Email         string   `json:"Email"`
	UserType      int      `json:"UserType"`
	Deleted       int      `json:"Deleted"`
	Resource      []string `json:"Resource"`
	ResourceCount int      `json:"ResourceCount"`
	Quota         int      `json:"Quota"`
}

type ListUserResponse struct {
	Users   []User `json:"User"`
	Success bool   `json:"Success"`
	Error   string `json:"Error"`
}

type DeleteUserRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
	UserID   int    `json:"UserID"`
}
