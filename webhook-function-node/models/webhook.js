const bookshelf = require('./../config/bookshelf');

const Webhook = bookshelf.Model.extend({
    tableName: 'webhooks'
});

module.exports = Webhook;
