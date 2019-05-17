# MomNote-server

### 数据库
利用Docker创建数据库
~~~bash
# 拉取镜像
docker pull heinoc/mysql
~~~

创建数据库
~~~bash
docker run --name mysql-release -v ~/.mysql-release:/var/lib/mysql -p 23306:3306 -e MYSQL_ROOT_PASSWORD=root_password -d heinoc/mysql
~~~
