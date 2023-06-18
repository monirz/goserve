# goserve
This is a Golang web service boilerplate using Ben Johnson's package layout. The boilerplate includes a simple RESTful API with a user model, user repository, and user service. It also separates the database service, using an adapter that abstracts the underlying database technology. This means that you can easily swicth to different database technology without affecting your domain logic.

This boilerplate is inspired by the article ["Structuring Applications in Go"](https://medium.com/@benbjohnson/structuring-applications-in-go-3b04be4ff091) by Ben Johnson, which introduced a new package structure that has since become popular in the Go community. By separating the domain from the infrastructure, this package structure allows for more modular and maintainable code.

To get started, simply clone this repository and run `go mod tidy` to download the required dependencies. 


## Installation

1. Clone the repository
```
git clone https://github.com/your-username/goserve.git
```
2. Install the dependencies
```
cd goserve
go mod tidy 

```

3. Create .env

```
touch .env
```
And put these or other necessary variables into the .env file

```
PORT=8090
DB_HOST=locahost
DB_PASSWORD=password
DB_USER=dbuser
DB_NAME=dbname
DB_PORT=5432
```

4. Run the test 

```
go test -v ./...

```

5. Run the application:
```
go run cmd/server/main.go
```





