# go-grpc-web-starter

Starter code for a project with the following architecture:

![](https://mermaid.ink/img/eyJjb2RlIjoiZmxvd2NoYXJ0IExSXG4gICAgQVtKYXZhU2NyaXB0IGJyb3dzZXIgY2xpZW50XSA8LS0gZ3JwYy13ZWIgLS0-IEJbRW52b3kgcHJveHldIDwtLSBncnBjIC0tPiBDW0dvIHNlcnZpY2VdIiwibWVybWFpZCI6eyJ0aGVtZSI6ImRlZmF1bHQifSwidXBkYXRlRWRpdG9yIjpmYWxzZX0)

## Requirements:

- [node.js](https://nodejs.org/en/)
- [npm](https://docs.npmjs.com/cli/v6/configuring-npm/install)
- [Envoy](https://www.envoyproxy.io/)
- [Go](https://golang.org/dl/)
- [protoc](https://grpc.io/docs/protoc-installation/), [protoc-gen-go](https://grpc.io/docs/languages/go/quickstart/), [protoc-gen-grpc-web](https://grpc.io/docs/platforms/web/basics/)

## Setup

1. Compile and generate protocol buffers.

   ```sh
   ./gen-proto.sh
   ```

2. Download Echo service dependencies and serve locally.

   ```sh
   cd server
   go mod download
   go run server.go
   ```

3. Run Envoy proxy

   ```sh
   cd envoy
   envoy -c envoy.yaml
   ```

4. Download client dependencies, build, and serve locally.

   ```sh
   cd client
   npm install
   npm run build
   dist/index.html # open in browser
   ```
