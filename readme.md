# pre-requisites
- gin
- mysql
- gorm

## gin
```shell
go get -u github.com/gin-gonic/gin
```

## mysql
```shell
go get -u github.com/go-sql-driver/mysql
```

## gorm, gorm mysql
```shell
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

# mysql run on docker
```shell
docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=1234 -e MYSQL_DATABASE=todolist -e MYSQL_USER=potato -e MYSQL_PASSWORD=1234 --name mysql mysql:5.7
```