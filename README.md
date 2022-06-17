# Fogg

Obfuscate details returned by your Cosmos RPC and API endpoints for additional privacy. Currently `fogg` will hide your unique Node ID as well as the listen address of your API and RPC services. 

## Setup

While we're assuming that you're using a reverse proxy (such as Nginx) in front of your API and RPC services, you should be able to use `fogg` in a way that best suites your particular deployment. Simply forward the traffic that's hitting the API endpoint `/node_info`, and the RPC endpoint `/status`, to your running `fogg` instance. 

### Golang

If you are new to Golang, please follow the setup instructions [here](https://golang.org/doc/install).

### Environment

Before running the `fogg` service, please ensure that you have the following environment variables set:

|Var|Description|
|---|-----------|
|`FOGG_PORT`|The port that the service should run on (e.g.: `3000`)|
|`API_HOST`|The API host, including the protocol and port (e.g.: `http://localhost:1317`).|
|`RPC_HOST`|The RPC host complete with protocol and port (e.g.: `http://localhost:26657`).|

### Nginx

If using Nginx, you can use the following configuration examples as a guide (please remember to set the port to whatever your `fogg` instance is listening on).

For your API service:

```
location /node_info {
    proxy_set_header X-Forwarded-Host $host:$server_port;
    proxy_set_header X-Forwarded-Server $host;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_pass http://localhost:3000/node_info;
}
```

and similarly for your RPC service:

```
location /status {
    proxy_set_header X-Forwarded-Host $host:$server_port;
    proxy_set_header X-Forwarded-Server $host;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_pass http://localhost:3000/status;
}
```

## Development

### Linter

To run the linter:

```console
make lint
```

### Tests

To run the tests and see test coverage:

```console
make tests
```