const path = require('path');
const knex = require('knex')({
  client: 'sqlite3',
  connection: {
    filename: path.resolve(__dirname, './../data/database.sqlite3')
  },
  useNullAsDefault: true
});

module.exports = require('bookshelf')(knex);
