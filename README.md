# loki

#### odin登录服务

#### 创建mysql-docker
`docker run -d -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=123456 mysql`
#### mysql备份参考
`docker run -d -p 3306:3306 --name mysql -v /data:/var/lib/mysql -e MYSQL_PASS="mypass" mysql`