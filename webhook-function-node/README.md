# Webhook Function
> A webhook function service written in Node.js

### Installation
1. Go through the project directory then install the dependencies
    ```
    npm install
    ```
2. Run the program
    ```
    node app.js
    ```

### Docker
You can deploy the service using docker
```
docker build -t webhook-function-node .
docker run -p 8083:8083 webhook-function-node
# or
docker-compose up -d
```

## Usage
1. Generate a webhook url
    ```bash
    curl -XPOST -H 'Content-Type: application/json' http://localhost:8083/webhook -d '
    {
        "code": "fetch(\"https:\/\/api.github.com\/users\/sohelamin\")\r\n  .then(res => res.json())\r\n  .then(json => callback(null, json))\r\n  .catch(callback);"
    }
    '
    ```
    Escape your codes properly then put in the JSON

2. Make a POST request to your generated webhook url
    ```bash
    curl -XPOST -H 'Content-Type: application/json' http://localhost:8083/webhook/<WEBHOOK_ID> -d '
    {
        "name": "Sohel Amin"
    }
    '
    ```