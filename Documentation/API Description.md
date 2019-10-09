# Backend-server

### How to send and receive Request
#### Request
The requests are sends in bytes. first 4 Byte of the packet will contain
the API Request information. The rest will contain the data. You can see the 
dataTypes under utilities directory.
#### Response
The response contain only the appropriate response struct in byte.
### API request code

The first 4 byte of the packet will have one of the following code in
BigEndian format.

```
Login          uint32 = 1
CreateResource uint32 = 2
ListResource   uint32 = 3
DeleteResource uint32 = 4
SetQuota       uint32 = 5
CreateUser     uint32 = 6
ListUser       uint32 = 7
DeleteUser     uint32 = 8
```

### API Descrpiton

A short description of the APIs is given bellow:

#### Login
Request type:

```
type LoginRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}
```

Response type: 

```
type LoginResponse struct {
	Success bool   `json:"Success"`
	Token   string `json:"Token"`
	UserID  int    `json:"UserID"`
	Error   string `json:"Error"`
}
```

* Used bcrypt password hashing mechanism because it is one way hash. And
we can add random salt.
* Token are JSON Web Tokens. It has authentication information, expiration time.
and It can work with multiple server. 


#### Create Resource
Request type:

```
type CreateResourceRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
	UserID   int    `json:"UserID"`
	Resource string `json:"Resource"`
}
```

Response type: 

```
type CommonResponse struct {
	Success bool   `json:"Success"`
	Error   string `json:"Error"`
}
```

* Takes an unique string as resource
 
 #### List Resource
 Request type:
 
 ```
 type ListResourceRequest struct {
 	Email    string `json:"Email"`
 	Password string `json:"Password"`
 	Token    string `json:"Token"`
 	UserID   int    `json:"UserID"`
 }
 ```
 
 Response type: 
 
 ```
type ListResourceResponse struct {
	Resource []string `json:"Resource"`
	Error    string   `json:"Error"`
}
 ```
 
 
 #### Delete Resource
 Request type:
 
 ```
type DeleteResourceRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
	UserID   int    `json:"UserID"`
	Resource string `json:"Resource"`
}

 ```
 
 Response type: 
 
 ```
 type CommonResponse struct {
 	Success bool   `json:"Success"`
 	Error   string `json:"Error"`
 }
 ```
#### Set Quota
Request type:
  
  ```
type SetQuotaRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Token    string `json:"Token"`
	UserID   int    `json:"UserID"`
	Quota    int    `json:"Quota"`
}
 
  ```
  Response type: 
  
  ```
  type CommonResponse struct {
  	Success bool   `json:"Success"`
  	Error   string `json:"Error"`
  }
  ```
 
 #### Create User
 Request type:
 
 ```
type CreateUserRequest struct {
	Email        string `json:"Email"`
	Password     string `json:"Password"`
	Token        string `json:"Token"`
	UserEmail    string `json:"UserEmail"`
	UserPassword string `json:"UserPassword"`
	UserType     int    `json:"UserType"`
}
 ```
 
 Response type: 
 
 ```
type CommonResponse struct {
  	Success bool   `json:"Success"`
  	Error   string `json:"Error"`
  }
 ```
 
  #### List User
  Request type:
  
  ```
  type ListUserRequest struct {
  	Email    string `json:"Email"`
  	Password string `json:"Password"`
  	Token    string `json:"Token"`
  	Limit    int    `json:"Limit"`
  	Offset   int    `json:"Offset"`
  }
  ```
  
  Response type: 
  
  ```
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
  ```
  
   #### Delete User
   Request type:
   
   ```
  type DeleteUserRequest struct {
  	Email    string `json:"Email"`
  	Password string `json:"Password"`
  	Token    string `json:"Token"`
  	UserID   int    `json:"UserID"`
  }
   ```
   
   Response type: 
   
   ```
  type CommonResponse struct {
    	Success bool   `json:"Success"`
    	Error   string `json:"Error"`
    }
   ```
   
   Note: Without login API, other APIs calls can be made with token that we 
   can get from login, or by the email and password.
   