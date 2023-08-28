# Checkout Case

The aim of this project is to complete the case study. Here we made a shopping
cart application. You can find more detailed information [here](/docs/case.md). 
You can also find the notes I took while working [here](/docs/case.png). 
The project receives input in two ways. As Rest and File.
I will be explaining in more detail in the following config structure.


If you need it for testing  postman collection [here](/docs/case.json) and commands collection [here](/docs/case.md).

## How to Config Application

`./config/default.go` you can make the settings you want on this file. `server.servertype: 'rest'` 
If you do, the application will receive rest requests. `server.servertype: 'file'`
the application will read from file and write to file. `server.restserver.pprofenable: 0` 
if you change this setting to 1 you will open pprof. Apart from these, rest address, 
file paths (I recommend you to be careful when changing them). 
Logger config and MongoDB config are available.

## How to Start App

After config settings;

```shell
    docker-compose up -d
```

if you want to see the application logs;

```shell
docker logs -f checkout-app
```
this command and the application will stand up. And you can send input according to your config settings.

## How to Test

Project Test information and run commands;

| *        | Unit test | Integration test |
|----------|-----------|------------------|
| Count    | 197       | 10               |
| Coverage | %84.5     | %78.5            |

* Regenerate mock

```shell
go generate ./...
```

* Unit test

```shell
go test ./...
```

* Integration test

```shell
go test -tags=integration ./...
```

## Tech Stack

* Go v1.19
* Rest server: [echo](https://echo.labstack.com/)
* Logger: [zap](https://github.com/uber-go/zap)
* Config: [viper](https://github.com/spf13/viper)
* Mock: [mockgen](https://github.com/golang/mock)
* Database: [MongoDB](https://www.mongodb.com/)
* Test container: [dockertest](https://github.com/ory/dockertest)
