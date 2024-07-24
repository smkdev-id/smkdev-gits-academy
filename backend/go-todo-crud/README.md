# Todo CRUD Golang

Simple CRUD Api with Golang, SQLite

## Endpoint

1. **Create (POST /):**

   - Metode: `POST`
   - URL: `http://localhost:8000/`
   - Body (raw JSON):
     ```json
     {
       "title": "Meet GITS Academy x SMKDEV",
       "status": false
     }
     ```

2. **Read (GET /):**

   - Metode: `GET`
   - URL: `http://localhost:8000/`

3. **Update (PUT /{id}):**

   - Metode: `PUT`
   - URL: `http://localhost:8080/{id}` (gantilah `{id}` dengan ID todo yang ingin Anda perbarui)
   - Body (raw JSON):
     ```json
     {
       "title": "Meet GITS Academy x SMKDEV",
       "completed": true
     }
     ```

4. **Delete (DELETE /{id}):**
   - Metode: `DELETE`
   - URL: `http://localhost:8080/{id}` (gantilah `{id}` dengan ID todo yang ingin Anda hapus)
