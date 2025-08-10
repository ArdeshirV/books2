#!/bin/sh
docker run --rm --name mysql-db -e MYSQL_ROOT_PASSWORD=root -p 3306:3306 -d mysql:latest

