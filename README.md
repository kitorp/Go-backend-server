# Backend-server
A Backend Server with GO that can do the following things:
* As a platform user, I need to authenticate with an email address and password.
* As a platform user, I need be able to create, list and delete resources.
* As a platform user, I should not be able to access resources owned by other users.
* As a platform user, I should not be able to create a new resource if the quota is
 exceeded.
* As a platform administrator, I should be able to create, list and delete users and their
 resources.
* As a platform administrator, I should be able to set the quota for any user.

Notes
* A resource is represented by a string with a unique identifier.
* Platform administrator is a platform user as well.
* By default, the quota is not set, which means a user can create as many resources as
he wants.
* You should provide responses for any errors that might occur.
### Prerequisites

You will need mysql client running on your PC. You can easily install it
using brew.

If you don't have brew, run the following command to install brew first

```
/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
```

then run this

```
brew install mysql
```



### Installing

To start the server, you need to configure the config.json file in 
project directory.

Then form the project directory run the following command

```
go run main.go
```

This will start the server

## Running the tests

The test are in client_test under test directory in the project folder.
Go to the test directory and Run:
```
go test -run Test
```
These are some simple Automated test to check if the end to end connection
is working properly. 
It also does a very basic test of the APIS.


## Built With

* [GO-LOGGER](https://github.com/apsdehal/go-logger) - A simple go logger for easy logging
* [JWT-GO](https://github.com/dgrijalva/jwt-go) - A Go implementation of  JSON Web Tokens
* [MY-SQL](https://github.com/go-sql-driver/mysql) - A MySQL-Driver for Go's database/sql package.


## Authors

* **Shadman Protik** - *Initial work* - [Backend Server](https://github.com/kitorp/backend-server)


