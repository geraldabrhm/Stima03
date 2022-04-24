const { Client } = require('pg')

const client = new Client({
    host: "localhost",
    user: "postgres",
    port: "5432",
    password: "root",
    database: "dna_tracker"
})

client.connect()

// client.query(`SELECT * FROM disease`, (err, res)=> {
//     if(!err) {
//         console.log(res.rows)
//     } else {
//         console.log(err.message)
//     }
//     client.end
// })