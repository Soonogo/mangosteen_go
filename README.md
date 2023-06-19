# 启动本地数据库

## psql

```bash
docker run -d --name pg-for-go-mangosteen -e POSTGRES_USER=mangosteen -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=mangosteen_dev -e PGDATA=/var/lib/postgresql/data/pgdata -v pg-go-mangosteen-data:/var/lib/postgresql/data --network=network1 postgres:14
```

## mysql

```bash
docker run -d --network=network1 --name mysql-for-go-mangosteen -e MYSQL_DATABASE=mangosteen_dev -e MYSQL_USER=mangosteen -e MYSQL_PASSWORD=123456 -e MYSQL_ROOT_PASSWORD=123456 -v mysql-go-mangosteen-data:/var/lib/mysql mysql:8 --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
```
# 数据库迁移

## 安装工具

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

## 创建迁移文件

```bash
go build . && ./mangosteen db create:migration create_users_table
# 或者
migrate create -ext sql -dir config/migrations -seq create_users_table
```
## 运行迁移文件

```bash
go build . && ./mangosteen db migrate
# 或者
migrate -database "postgres://mangosteen:123456@pg-for-go-mangosteen:5432/mangosteen_dev?sslmode=disable" -source "file://$(pwd)/config/migrations" up
```

## 回滚迁移文件

```bash
go build . && ./mangosteen db migrate:down
# 或者
migrate -database "postgres://mangosteen:123456@pg-for-go-mangosteen:5432/mangosteen_dev?sslmode=disable" -source "file://$(pwd)/config/migrations" down 1
```


## 测试覆盖率
```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverarage.html
```