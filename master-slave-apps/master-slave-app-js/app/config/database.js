const mariadb = require('mariadb');

const masterPool = mariadb.createPool({
 host: process.env.DB_WRITE_HOST,
  user: process.env.DB_WRITE_USER,
  password: process.env.DB_WRITE_PASSWORD,
  database: process.env.DB_NAME,
  port: parseInt(process.env.DB_WRITE_PORT), // Port untuk master
  connectionLimit: parseInt(process.env.DB_WRITE_CONNECTION_LIMIT),
  acquireTimeout: parseInt(process.env.DB_WRITE_TIMEOUT), // Tingkatkan timeout ke 30 detik
  trace: true // Aktifkan trace untuk debugging
});

const slavePool = mariadb.createPool({
host: process.env.DB_READ_HOST,
  user: process.env.DB_READ_USER,
  password: process.env.DB_READ_PASSWORD,
  database: process.env.DB_NAME,
  port: parseInt(process.env.DB_READ_PORT), // Port untuk slave
  connectionLimit: parseInt(process.env.DB_READ_CONNECTION_LIMIT),
  acquireTimeout: parseInt(process.env.DB_READ_TIMEOUT), // Tingkatkan timeout ke 30 detik
  trace: true // Aktifkan trace untuk debugging
});

const getConnection = async (type) => {
  if (type === 'read') {
    return await slavePool.getConnection();
  } else {
    return await masterPool.getConnection();
  }
};


module.exports = { getConnection };