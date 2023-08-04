#!/bin/bash
docker run -d --restart always -p 13306:3306 -e MYSQL_ROOT_PASSWORD=123456 --name MySQL mysql:5.7
