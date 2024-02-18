.PHONY: build run db_run db_create

# 应用名称
APP_NAME="fuxiaochen-api-with-go"

# MySQL 容器名称
MYSQL_CONTAINER_NAME="mysql8"
# MySQL root用户密码
MYSQL_ROOT_PASSWORD="123456"
# MySQL 初始化时创建的数据库名称
MYSQL_DB_NAME="fuxiaochen"

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${APP_NAME}

run:
	go run main.go

# 用 Docker 启动 MySQL
db_run:
	docker run --name ${MYSQL_CONTAINER_NAME} -p 3306:3306 -e MYSQL_ROOT_PASSWORD=${MYSQL_ROOT_PASSWORD} -v ${APP_NAME}:/var/lib/mysql -d mysql:8.3.0

# 在启动好的 Docker 容器中创建数据库
db_create:
	docker exec -it ${MYSQL_CONTAINER_NAME} mysql -uroot -p${MYSQL_ROOT_PASSWORD} -e "CREATE DATABASE ${MYSQL_DB_NAME};"