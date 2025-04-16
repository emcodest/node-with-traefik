var express = require('express');
var router = express.Router();

/* GET home page. */
router.get('/', function(req, res, next) {
  // nice
  res.render('index', { title: 'Emcode App' });
});

module.exports = router;
