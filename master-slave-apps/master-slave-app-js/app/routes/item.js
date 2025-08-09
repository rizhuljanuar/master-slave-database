const express = require('express');
const router = express.Router();
const {
  createItemController,
  getAllItemsController,
  getItemByIdController,
  updateItemController,
  deleteItemController
} = require('../controllers/item');

router.post('/items', createItemController);
router.get('/items', getAllItemsController);
router.get('/items/:id', getItemByIdController);
router.put('/items/:id', updateItemController);
router.delete('/items/:id', deleteItemController);

module.exports = router;