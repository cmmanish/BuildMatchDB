https://medium.com/swlh/how-to-connect-to-mysql-docker-from-python-application-on-macos-mojave-32c7834e5afa
Run Docker for Mysql,
    `docker run  --name ms -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password mysql`

Exec into the container
    `docker exec -it ms bash`

Login to mysql: `mysql -u root -ppassword`
Create database :`CricketDb` and 
Create table `create table odi_ball_by_ball_v2;`

Let’s create a new user named newuser with password newpassword.
    `CREATE USER 'user'@'%' IDENTIFIED BY 'password';`

Grant a sort of admin access to newuser to manage the database
    `GRANT ALL PRIVILEGES ON CricketDb.* to 'user'@'%';`

Choose a Db
`use CricketDb;`

Select no of rows in table: 
`select count(*) from odi_ball_by_ball_v2;`

Backup and restore a mysql database from a running Docker mysql container
docker exec CONTAINER /usr/bin/mysqldump -u root --password=root DATABASE > backup.sql
cat backup.sql | docker exec -i CONTAINER /usr/bin/mysql -u root --password=root DATABASE