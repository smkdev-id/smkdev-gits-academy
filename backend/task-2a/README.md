# TODO APP

Create, Read, Update, Delete (CRUD) application of todo.

## How to run application

1. you can setup `env.sh` file, get example from `env.example.sh` file.

```sh
#!bin/bash

# DATABASE
export DB_DRIVER=mysql
export DB_HOST=localhost
export DB_PORT=3306
export DB_USER=<your-db-user>
export DB_PASSWORD=<your-db-password>
export DB_NAME=<your-db-name>

# APPLICATION
export APP_PORT=":8080"
```

2. generate migrate create ⤵️ , open your terminal :

```bash
make migrate-crete
```

3. setup your table in `db/migration/<date-exec>_task2.up.sql`.

```sql
CREATE TABLE todo (
    id VARCHAR(55) PRIMARY KEY NOT NULL,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(555) NOT NULL,
    status BOOLEAN NOT NULL DEFAULT true,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
) ENGINE=INNODB;
```

4. if you have written table, you can exec with command :

```bash
make migrate-up
```

- then, if you want down/drop your table generated :

```bash
make migrate-down
```

5. after you can do migration up your table, you can run command :

```bash
go run main.go
```

## How to run using Docker

1. setup your environment variable :

```sh
#!bin/bash

# DB CONTAINER
export DB_DRIVER=mysql
export DB_HOST=localhost
export DB_PORT=3306
export DB_USER=root
export DB_PASSWORD=123
export DB_NAME=task2

# APPLICATION
export APP_PORT=":8080"
```

2. run with command :

```bash
sudo docker compose up -d
```

- if you down :

```bash
sudo docker compose down -v
```

1. you can create table in `mysql-container`, using command :

```bash
sudo docker exec -it mysql-container /bin/bash
```

- if you inside `mysql-container`, then go `mysql`, command :

```bash
home# mysql -u rooot -p
home# password:123
```

4. inside to `mysql` :

```bash
mysql> use task2;
```

5. create yout table :

```bash
mysql> CREATE TABLE todo (
            id VARCHAR(55) PRIMARY KEY NOT NULL,
            title VARCHAR(255) NOT NULL,
            description VARCHAR(555) NOT NULL,
            status BOOLEAN NOT NULL DEFAULT true,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP
        );
```

### CONGRATULATION 👍
