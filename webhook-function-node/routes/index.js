const router = require('express').Router();

router.get('/', (req, res) => {
  res.send('There is no place like home.');
})

module.exports = router;
