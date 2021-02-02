const {EchoRequest, EchoResponse} = require('../echo/echo_pb.js');
const {EchoServiceClient} = require('../echo/echo_grpc_web_pb.js');

var echoService = new EchoServiceClient('http://localhost:8080');

var request = new EchoRequest();
request.setMessage('Hello World!');

echoService.echo(request, {}, function(err, response) {
  console.log(response)
  document.getElementById("hello-world").innerHTML = response.array[0]
});