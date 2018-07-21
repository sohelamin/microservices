const express = require('express');
const path = require('path');
const bodyParser = require('body-parser');

const app = express();

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: false }));

app.use('/', require('./routes/index'));
app.use('/webhook', require('./routes/webhook'));

app.listen(8083, () => console.log('Listening on port 8083!'))
