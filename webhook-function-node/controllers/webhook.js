const vm = require('vm');
const fetch = require('node-fetch');
const uuidv1 = require('uuid/v1');
const Webhook = require('./../models/webhook');

const callback = (err, output = null) => {
  if (err) {
    console.log(err);
  } else {
    console.log(output);
  }
};

const generate = async (req, res) => {
  const code = req.body.code;
  const uuid = uuidv1();
  await new Webhook({ uuid, code }).save();

  const webhookUrl = 'http://localhost:8083/webhook/' + uuid;
  res.status(201).send({'webhook_url': webhookUrl});
};

const fire = async (req, res) => {
  const webhookId = req.params.id;
  const inputs = req.body;

  const sandbox = { fetch, callback, inputs };
  vm.createContext(sandbox);

  try {
    const model = await new Webhook({ uuid: webhookId }).fetch();
    await vm.runInContext(model.get('code'), sandbox);
  } catch(err) {
    console.log(err);
  }

  res.send({});
};

module.exports = { generate, fire };
