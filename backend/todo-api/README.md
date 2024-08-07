
# Golang Simple TODO API

Ini adalah api todolist simple menggunakan go echo, dan gorm

# Contoh Penggunaan / List Endpoint

```
GET :/
Fetch all todolist. example output:
{
    "data": [
        {
            "id": 1,
            "todo": "Membeli kopi kapal api ke warung",
            "finished": false,
            "created_at": "2024-07-21T14:28:55.8079253+07:00",
            "finished_at": "0001-01-01T00:00:00Z"
        },
        {
            "id": 2,
            "todo": "Membeli beras ke toko abadi",
            "finished": true,
            "created_at": "2024-07-21T14:29:06.5568964+07:00",
            "finished_at": "2024-07-21T14:30:28.8959745+07:00"
        }
    ],
    "message": "success"
}
```

```
POST :/
Create new todo. example usage:
{
    "todo": "Membeli kopi kapal api ke warung"
}
```

```
PUT :/{id}
Update todo. example usage:
{
    "finished": true
}
```

```
DELETE :/{id}
Delete todo with id.
```

