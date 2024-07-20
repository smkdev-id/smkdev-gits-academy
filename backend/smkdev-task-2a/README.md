# REST API CRUD Todo-List

These examples are taken from a Todo application built with [Golang](https://go.dev/) using the [Echo framework](github.com/labstack/echo/v4) and [SQLite](https://www.sqlite.org/). The JSON responses are designed in a way that is easy to interpret and use.

Where full URLs are provided in responses they will be rendered as if service
is running on 'http://localhost:8080/'.

## Tech Stack

**Language:** Golang \
**Server:** Echo \
**Database:** SQLite


## Run Locally

Clone the project

```bash
  git clone https://github.com/smkdev-id/smkdev-gits-academy.git
```

Go to the project directory

```bash
  cd smkdev-gits-academy/backend/smkdev-task-2a/
```

Start the server

```bash
  ./main (running main.exe)
```
or

```bash
  go run main.go (running main.go)
```

## Todo-list Endpoints

These endpoints are related to the Todo CRUD operations:

### Get All Todos 

URL : `/` \
Method : `GET`
#### Successful Responses 

Code : `200 OK` \
Content examples 

```json
{
    "status_code": 200,
    "message": "successfully get all todos",
    "data": [
        {
            "id": "dq2",
            "createdAt": "2024-07-20T01:44:59Z",
            "title": "sangkuriang",
            "description": "asdas",
            "status": false
        },
        {
            "id": "d55f7b49-8a6d-4888-8989-f0d8558efdfa",
            "createdAt": "2024-07-20T15:12:35.4954943+07:00",
            "title": "2024 Menjadi Billioner Muda",
            "description": "menjadi kaya, harus kerja keras pantang menyerah, serta tawakal rajin ibadah",
            "status": true
        }
    ]
}
```
#### Error Responses 
Code : `500 Internal Server Error`



### Create Todo

URL : `/` \
Method : `POST`
#### Request
This is a client request to create a todo
```json
{
    "title": "Menyelesaikan soal Matematika",
    "description" : "Selesaikan soal matematika 1 soal per 1 jam",
    "status": true
}
```
#### Successful Responses 

Code : `201 Created` \
Content examples 

```json
{
    "status_code": 201,
    "message": "Create successfully",
    "data": {
        "id": "de0735b7-aa25-42ac-96d3-5b492eb8b541",
        "createdAt": "2024-07-20T23:24:57.4396065+07:00",
        "title": "Menyelesaikan soal Matematika",
        "description": "Selesaikan soal matematika 1 soal per 1 jam",
        "status": true
    }
}
```
#### Error Responses 
Code : `500 Internal Server Error` or `400 Bad request`

### Get A Todo

URL : `/:id` \
Method : `GET`

#### Successful Responses 

Code : `200 OK` \
Content examples 

```json
{
    "status_code": 201,
    "message": "Create successfully",
    "data": {
        "id": "de0735b7-aa25-42ac-96d3-5b492eb8b541",
        "createdAt": "2024-07-20T23:24:57.4396065+07:00",
        "title": "Menyelesaikan soal Matematika",
        "description": "Selesaikan soal matematika 1 soal per 1 jam",
        "status": true
    }
}
```
#### Error Responses 
Code : `500 Internal Server Error` or `404 Not found`

### Update A Todo

URL : `/:id` \
Method : `PUT`
#### Request
This is a client request to update a todo
```json
{
    "title": "2024 Menjadi Billioner Muda",
    "description" : "menjadi kaya, harus kerja keras pantang menyerah, serta tawakal rajin ibadah",
    "status": true
}
```
#### Successful Responses 

Code : `200 OK` \
Content examples 

```json
{
    "status_code": 200,
    "message": "Update successfully",
    "data": {
        "id": "d55f7b49-8a6d-4888-8989-f0d8558efdfa",
        "createdAt": "0001-01-01T00:00:00Z",
        "title": "2024 Menjadi Billioner Muda",
        "description": "menjadi kaya, harus kerja keras pantang menyerah, serta tawakal rajin ibadah",
        "status": true
    }
}
```
#### Error Responses 
Code : `500 Internal Server Error` or `400 Bad request`

### Delete A Todo

URL : `/:id` \
Method : `DELETE`

#### Successful Responses 

Code : `200 OK` \
Content examples 

```json
{
    "status_code": 200,
    "message": "Delete successfully",
    "data": null
}
```
#### Error Response 
Code : `500 Internal Server Error`





* Update A Todo : `PUT /:id`
* Delete A Todo : `DELETE /:id`
## 

This README provides a comprehensive guide to understanding, setting up, and interacting with the Todo-List CRUD REST API built with Golang, Echo, and SQLite3.