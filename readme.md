# Swagger
To access Swagger, go to [/swagger/index.html](http://localhost:9000/swagger/index.html) after starting the application.

## To Generate docs
First, run
```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Second, run following command in the root directory of the project
```bash
swag init
```
This will generate the docs folder with the swagger documentation.

## To run the application
```bash
go run main.go
```

## Coding Pattern
- Each function's **output parameter** should have field name.
- If a function must return an array of model, it should be named as **GetAll{ModelName}** and should have a **plural** name.
- In package **database**, each function's input parameter should have db as the last parameter.
- For each **"CREATE" , "UPDATE" and "DELETE"** operation, some function in package "database" should be **cascaded operation**.
  - Modified the code to send **transaction** variable instead.
```go
package controllers

func doSomethingInController() {
	// Start a transaction
	tx := database.mainDB.Begin()
	
	// Do some operations 
	// ...
	
	// if error occurs, rollback the transaction with 
	tx.Rollback()
	
	// if no error occurs, commit the transaction with
	tx.Commit()	
}
```

```go
package database

func doSomethingInDatabase(db *gorm.DB) {
	tx.savePoint("savepoint_name")
	
	// Do some operations 
	//...
	
	// if error occurs, rollback the transaction with
	tx.rollbackTo("savepoint_name")
	
	// if no error occurs, just do everything as usual
}
```