const { getConnection } = require('../config/database.js');
const { formatDate } = require('../helpers/date_format.js')

const createItem = async (title, description, qty, price) => {
  const conn = await getConnection('write');
  const nowTime = new Date();
  try {
    const result = await conn.query(
      'INSERT INTO items (title, description, qty, price, created_at, updated_at) VALUES (?, ?, ?, ?, ? ,?)',
      [title, description, qty, price, formatDate(nowTime), formatDate(nowTime)]
    );
    return { id: Number(result.insertId), title, description, qty: Number(qty),price:Number( price) , created_at: formatDate(nowTime)};
  } finally {
    conn.release();
  }
};

// row.tanggal ? new Date(row.tanggal).toISOString() : null
const getAllItems = async () => {
  const conn = await getConnection('read');
  try {
    const rows = await conn.query('SELECT id, title, description, qty, price, created_at, updated_at FROM items');
    if (rows.length === 0) {
      console.log(`Items dengan tidak tersedia`);
      return null;
    }
    return rows.map(row => ({
      id: Number(row.id),
      title: row.title,
      description: row.description,
      qty: Number(row.qty),
      price: Number(row.price),
      created_at: row.created_at ? formatDate(row.created_at) : null,
      updated_at: row.updated_at ? formatDate(row.updated_at) : null
    }));
  } finally {
    conn.release();
  }
};

const getItemById = async (id) => {
  const conn = await getConnection('read');
  try {
    const rows = await conn.query('SELECT id, title, description, qty, price, created_at, updated_at FROM items WHERE id = ?', [id]);

    if (rows.length === 0) {
      console.log(`Item dengan ID: ${id} tidak ditemukan`);
      return null;
    }

    return {
      id: Number(rows[0].id),
      title: rows[0].title,
      description: rows[0].description,
      qty: Number(rows[0].qty),
      price: Number(rows[0].price),
      created_at: rows[0].created_at ? formatDate(rows[0].created_at) : null,
      updated_at: rows[0].updated_at ? formatDate(rows[0].updated_at) : null
    };
  } finally {
    conn.release();
  }
};

const updateItem = async (id, title, description, qty, price) => {
  const conn = await getConnection('write');
   const nowTime = new Date();
  try {
    const result = await conn.query(
      'UPDATE items SET title = ?, description = ?, qty = ?, price = ?, updated_at = ? WHERE id = ?',
      [title, description, qty, price, formatDate(nowTime), id]
    );
    if (result.affectedRows === 0) {
      return null;
    }

    // Ambil data terbaru setelah update
    const updatedRows = await conn.query('SELECT * FROM items WHERE id = ?', [id]);
    if (updatedRows.length === 0) return null;
    
    return {
      id: Number(updatedRows[0].id),
      title: updatedRows[0].title,
      description: updatedRows[0].description,
      qty: Number(updatedRows[0].qty),
      price: Number(updatedRows[0].price),
      created_at: formatDate(updatedRows[0].created_at),
      updated_at: formatDate(updatedRows[0].updated_at)
    };

  } finally {
    conn.release();
  }
};

const deleteItem = async (id) => {
  const conn = await getConnection('write');
  try {
   // console.log(`Mencoba menghapus item dengan ID: ${id}`);
    const result = await conn.query('DELETE FROM items WHERE id = ?', [id]);
    if (result.affectedRows > 0) {
     // console.log(`Berhasil menghapus item dengan ID: ${id}`);
      return true;
    } else {
      //console.log(`Item dengan ID: ${id} tidak ditemukan`);
      return false;
    }
  } finally {
    conn.release();
  }
};

module.exports = { createItem, getAllItems, getItemById, updateItem, deleteItem };