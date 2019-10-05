package library

type LoginRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type LoginResponse struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
	UserId   int    `json:"UserId"`
	Error    string `json:"Error"`
}

type CreateResourceRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
	Resource string `json:"Resource"`
}

type CreateResourceResponse struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
	Resource string `json:"Resource"`
	Success  bool   `json:"Success"`
	Error    string `json:"DebugMsg"`
}

type ListResourceRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
}

type ListResourceResponse struct {
	Email    string   `json:"Email"`
	Password string   `json:"Password"`
	Token    string   `json:"Token"`
	Resource []string `json:"Resource"`
	Error    string   `json:"DebugMsg"`
}

type DeleteResourceRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
	Resource string `json:"Resource"`
}

type DeleteResourceResponse struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
	Resource string `json:"Resource"`
	Success  bool   `json:"Success"`
	Error    string `json:"DebugMsg"`
}

type SetQuotaRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
	UserID   int    `json:"UserID"`
	Quota    int    `json:"Quota"`
}

type SetQuotaResponse struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
	UserID   int    `json:"UserID"`
	Quota    int    `json:"Quota"`
	Success  bool   `json:"Success"`
	Error    string `json:"DebugMsg"`
}

type CreateUserRequest struct {
	Email        string `json:"Email"`
	Password     string `json:"Password"`
	Token        string `json:"Token"`
	UserEmail    string `json:"UserEmail"`
	UserPassword string `json:"UserPassword"`
}

type CreateUserResponse struct {
	Email        string `json:"Email"`
	Password     string `json:"Password"`
	Token        string `json:"Token"`
	UserEmail    string `json:"UserEmail"`
	UserPassword string `json:"UserPassword"`
	UserID       int    `json:"UserID"`
	Success      bool   `json:"Success"`
	Error        string `json:"DebugMsg"`
}

type ListUserRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
}

type User struct {
	UserID   int      `json:"UserID"`
	Email    string   `json:"Email"`
	Quota    int      `json:"Quota"`
	Resource []string `json:"Resource"`
}

type ListUserResponse struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
	Users    []User `json:"User"`
	Error    string `json:"DebugMsg"`
}

type DeleteUserRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
	UserID   int    `json:"UserID"`
}

type DeleteUserResponse struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
	UserID   int    `json:"UserID"`
	Success  bool   `json:"Success"`
	Error    string `json:"DebugMsg"`
}
