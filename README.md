# loki

#### odin登录服务

#### 创建mysql-docker
`docker run -d -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=123456 mysql`
#### mysql备份参考
`docker run -d -p 3306:3306 --name mysql -v /data:/var/lib/mysql -e MYSQL_PASS="mypass" mysql`

### 简单压测
-n ：总共的请求执行数，缺省是1；

-c： 并发数，缺省是1；

-t：测试所进行的总时间，秒为单位，缺省50000s

-p：POST时的数据文件

-w: 以HTML表的格式输出结果

get

`ab -c 100 -n 300 -w http://0.0.0.0:10900/auth >> ./ab/get.html`

post
`ab -c 100 -n 3000 -p ./ab/post.txt -T application/x-www-form-urlencoded http://0.0.0.0:10900/login`
