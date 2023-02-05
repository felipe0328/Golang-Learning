# Project # 1 Gin + SQL: [GinSQL](/ginSql/)
Creation of a simple web server using gin, an sql db and swagger to document all the endpoints.

Frameworks used:
* [GIN](https://github.com/gin-gonic/gin)
* [Swaggo](https://github.com/swaggo/swag)
* [Testify](https://github.com/stretchr/testify)


## Project Description
This project is a web server imitating a task managment service, it will containt routes to get the employees, the projects and the tasks, the swagger documentation can be found here

[Swagger Documentation](/ginSql/docs/)

it will run in localhost:8080/docs/index.html

[Swagger Web Documentation](http://localhost:8080/docs/index.html)

This project is using an [uiBackery](https://uibakery.io/sql-playground) Database, to avoid the need of create a database and its migrations, making it easy for the user to test this project just running the main.go file

To use the project just run

    go run main.go

If you want to run the tests, just run
    
    go test ./...