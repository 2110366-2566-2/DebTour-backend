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
