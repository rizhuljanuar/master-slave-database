const { createItem, getAllItems, getItemById, updateItem, deleteItem } = require('../models/item');

const createItemController = async (req, res) => {
  const { title, description, qty, price } = req.body;
  try {
    const item = await createItem(title, description, qty, price);
    res.status(200).json({
      code: 200,
      message: 'success',
      data: item
    });
  } catch (err) {
    res.status(500).json({
      code: 500,
      message: "Gagal membuat item",
      data: null
    });
  }
};

const getAllItemsController = async (req, res) => {
  try {
    const items = await getAllItems();
    if (!items) {
      return res.status(404).json({
      code: 404,
      message: 'Item tidak tersedia',
      data: items
    });
    }
    res.status(200).json({
      code: 200,
      message: 'success',
      data: items
    });
  } catch (err) {
    res.status(500).json({
      code: 500,
      message: 'Gagal mengambil data',
      data: null
    });
  }
};

const getItemByIdController = async (req, res) => {
  try {
    const item = await getItemById(req.params.id);

    if (!item) {
      return res.status(404).json({
        code: 404,
        message: 'Item tidak ditemukan',
        data: null
      });
    }
    res.status(200).json({
      code: 200,
      message: 'success',
      data: item
    });

  } catch (err) {
    res.status(500).json({
      code: 500,
      message: 'Gagal mengambil data',
      data: null
    });
  }
};

const updateItemController = async (req, res) => {
  const { title, description, qty, price } = req.body;
  try {
    const item = await updateItem(req.params.id, title, description, qty, price);
    if (!item) {
      return res.status(404).json({
        code: 404,
        message: 'Item tidak ditemukan',
        data: null
      });
    }
    
    res.status(200).json({
      code: 200,
      message: 'Item berhasil di update',
      data: item
    });
    
  } catch (err) {
    res.status(500).json({
      code: 500,
      message: 'Gagal update item',
      data: null
    });
  }
};

const deleteItemController = async (req, res) => {
  try {
    console.log(`Menerima permintaan DELETE untuk ID: ${req.params.id}`);
    const success = await deleteItem(req.params.id);
    if (!success) {
      return res.status(404).json({
        code: 404,
        message: 'Item tidak ditemukan',
        data: null
      });
    }
    res.status(200).json({
      code: 200,
      message: 'Item berhasil dihapus',
      data: null
    });
  } catch (err) {
    console.error(`Error saat DELETE ID: ${req.params.id}: ${err.message}`);
    res.status(500).json({
      code: 500,
      message: 'Gagal menghapus item',
      data: null
    });
  }
};

module.exports = {
  createItemController,
  getAllItemsController,
  getItemByIdController,
  updateItemController,
  deleteItemController
};