# pip install mysql-connector-python
import mysql.connector


DATABASE_NAME = "tiktok"


def CreateDatabase():
    # 连接数据库
    connection = mysql.connector.connect(
        host="localhost",
        user="root",
        password="123456"
    )

    # 创建数据库
    create_database_query = "CREATE DATABASE IF NOT EXISTS " + DATABASE_NAME
    cursor = connection.cursor()
    cursor.execute(create_database_query)

    # 关闭游标和连接
    cursor.close()
    connection.close()


def Create():
    # 连接数据库
    connection = mysql.connector.connect(
        host="localhost",
        user="root",
        password="123456",
        database=DATABASE_NAME
    )

    # 创建游标
    cursor = connection.cursor()

    # 读取 SQL 文件
    sql_file = "tiktok.sql"  # SQL 文件名

    with open(sql_file, "r") as file:
        sql_statements = file.read().split(';')

    for statement in sql_statements:
        if statement.strip():
            cursor.execute(statement)

    connection.commit()
    cursor.close()
    connection.close()

CreateDatabase()
Create()