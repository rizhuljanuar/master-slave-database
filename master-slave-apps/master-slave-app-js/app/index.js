const express = require('express');
const bodyParser = require('body-parser');
require('dotenv').config();
const app = express();


const itemRoutes = require('./routes/item');

app.use(express.json());
app.use('/', itemRoutes);
// Middleware untuk menangani route yang tidak ditemukan
app.use((req, res) => {
  console.log(`Route tidak ditemukan: ${req.method} ${req.originalUrl}`);
  res.status(404).json({
    code: 404,
    message: 'Rute tidak ditemukan',
    data: null
  });
});


let  PORT = process.env.APP_PORT;
app.listen(PORT, () => {
  console.log(`Server running on http://localhost:${PORT}`);
});


