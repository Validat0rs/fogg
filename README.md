# Fogg

Hide those attributes returned by your Cosmos RPC and API endpoints that could identify your node.

## Overview

Several Cosmos API and RPC endpoints return attributes that make it trivial to identify your node on the network, so `fogg` acts as a man in the middle in order to hide that information. Please see the following table for what is currently hidden with `fogg`:

|Endpoint|Service|Hidden Attributes|
|--------|-------|-----------------|
|`/node_info`|API|`ID` and `ListenAddr`|
|`/status`|RPC|`ID`, `ListenAddr` and `ValidatorInfo`.|

## Setup

### Golang

If you are new to Golang, please follow the setup instructions [here](https://golang.org/doc/install).

### Environment

Before running `fogg`, please ensure that you have the following environment variables set:

|Var|Description|
|---|-----------|
|`FOGG_PORT`|The port that the service should run on (e.g.: `3000`)|
|`API_HOST`|The API host, including the protocol and port (e.g.: `http://localhost:1317`).|
|`RPC_HOST`|The RPC host complete with protocol and port (e.g.: `http://localhost:26657`).|

### Nginx

If using Nginx, you can use the following configuration examples as a guide (please remember to set the port to whatever your `fogg` instance is listening on). 

The examples should also demonstrate how you can support both `fogg` and proxying everything else to the API and RPC services directly.

For your API service:

```
location / {
    proxy_set_header X-Forwarded-Host $host:$server_port;
    proxy_set_header X-Forwarded-Server $host;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        
    location ~ /node_info {
        proxy_pass http://localhost:3000/api/node_info;
    }

    proxy_pass http://localhost:1317;
}
```

and similarly for your RPC service:

```
location / {
    proxy_set_header X-Forwarded-Host $host:$server_port;
    proxy_set_header X-Forwarded-Server $host;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        
    location ~ /status {
        proxy_pass http://localhost:3000/rpc/status;
    }

    proxy_pass http://localhost:26657;
}
```

## Run

### Install

To install the binary, run:

```console
make clean install
```

### Start

To then to start the service:

```console
make run-fogg-service
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