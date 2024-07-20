#!bin/bash

# DATABASE
export DB_DRIVER=mysql
export DB_HOST=localhost
export DB_PORT=3306
export DB_USER=<your-db-user>
export DB_PASSWORD=<your-db-password>
export DB_NAME=<your-db-name>

# DB CONTAINER (inside container) for running with docker-compose
export DB_DRIVER=mysql
export DB_HOST=localhost
export DB_PORT=3306
export DB_USER=root
export DB_PASSWORD=123
export DB_NAME=task2

# APPLICATION
export APP_PORT=":8080"