import express from "express"
import mysql from "mysql"
import cors from "cors"

const app = express()

const db = mysql.createConnection({
    host: "localhost",
    user: "root",
    password: "",
    database: "test"
});

// biar bisa nambahin data lewat body
app.use(express.json())
// biar bisa diambil datanya di frontend, ini dikasih setelah install cors
app.use(cors())

// testing connect ke backend
app.get("/", (req, res)=> {
    res.json("hello this is the backend")
})


// mengambil semua data
app.get("/books", (req, res)=> {
    const q = "SELECT * FROM books"
    db.query(q, (err, data)=> {
        if(err) return res.json(err)
        return res.json(data)
    })
})

// post book
app.post("/books", (req, res)=> {
    const q = "INSERT INTO books (`title`, `desc`, `price`, `cover`) VALUES (?)"
    const values = [
        // "title from backend",
        // "desc from backend",
        // "1234"
        // "cover from backend"
        req.body.title,
        req.body.desc,
        req.body.price,
        req.body.cover,
    ];

    db.query(q, [values], (err, data)=> {
        if(err) return res.json(err)
        return res.json("Book has been created successfully.")
    });
});

// delete book
app.delete("/books/:id", (req, res)=> {
    const bookID = req.params.id;
    const q = "DELETE FROM books WHERE id =?"

    db.query(q, [bookID], (err, data)=> {
        if(err) return res.json(err)
        return res.json("Book has been deleted successfully.")
    });
});

// update book, copy dari delete, values copy dari add
app.put("/books/:id", (req, res)=> {
    const bookID = req.params.id;
    const q = "UPDATE books SET `title`=?, `desc`=?, `price`=?, `cover`=? WHERE id=?";

    const values = [
        req.body.title,
        req.body.desc,
        req.body.price,
        req.body.cover,
    ];
    db.query(q, [...values,bookID], (err, data)=> {
        if(err) return res.json(err)
        return res.json("Book has been updated successfully.")
    });
});

app.listen(8800, ()=>{
    console.log("Server is running on port 8800")
})