const router = require('express').Router();
const webhookController = require('./../controllers/webhook');

router.post('/', webhookController.generate);
router.post('/:id', webhookController.fire);

module.exports = router;
