package library

type LoginRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type LoginResponse struct {
	Token  string `json:"Token"`
	UserID int    `json:"UserID"`
	Error  string `json:"Error"`
}

type CreateResourceRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
	UserID   int    `json:"UserID"`
	Resource string `json:"Resource"`
}

type CreateResourceResponse struct {
	Success bool   `json:"Success"`
	Error   string `json:"DebugMsg"`
}

type ListResourceRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
	UserID   int    `json:"UserID"`
}

type ListResourceResponse struct {
	Resource []string `json:"Resource"`
	Error    string   `json:"DebugMsg"`
}

type DeleteResourceRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
	UserID   int    `json:"UserID"`
	Resource string `json:"Resource"`
}

type DeleteResourceResponse struct {
	Success bool   `json:"Success"`
	Error   string `json:"DebugMsg"`
}

type SetQuotaRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
	UserID   int    `json:"UserID"`
	Quota    int    `json:"Quota"`
}

type SetQuotaResponse struct {
	Success bool   `json:"Success"`
	Error   string `json:"DebugMsg"`
}

type CreateUserRequest struct {
	Email        string `json:"Email"`
	Password     string `json:"Password"`
	Token        string `json:"Token"`
	UserEmail    string `json:"UserEmail"`
	UserPassword string `json:"UserPassword"`
}

type CreateUserResponse struct {
	Success bool   `json:"Success"`
	Error   string `json:"DebugMsg"`
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
	Success bool   `json:"Success"`
	Users   []User `json:"User"`
	Error   string `json:"DebugMsg"`
}

type DeleteUserRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
	UserID   int    `json:"UserID"`
}

type DeleteUserResponse struct {
	Success bool   `json:"Success"`
	Error   string `json:"DebugMsg"`
}
