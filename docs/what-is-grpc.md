### What is gRPC ?

advantages of proto over json
- faster to encode and decode 
- smaller over the network
- strongly typed 

protobuf sends a compact binary format , humans cannot read it
- **Protocol Buffers**, or **protobuf**, is the overall technology/serialization format created by Google.

stub -> means client in most of the cases 

protocol buffer -> google's mature open source mechanism for serialising structured data, can be used with other data formats such as JSON.

working with protocol buffers:
- define the structure for the data you want to serialize in a proto file. 
- this is an ordinary text file with .proto extension 
- profo buffer data is **structured as messages** where each message is a small logical record of information containing a series of name-value pairs called fields.

```example.proto
message Person {
 string name = 1;
 int32 id = 2;
 bool has_subscribed = 3;
}
```
- protoc - protocol buffer compiler to generate data access clasess in your preffered languages. 

## flow

.proto file ->  protoc compiler -> generated .pb.go file -> use generated Go structs in your app

## defining gRPC services
- define gRPC services in a ordinary proto files , with RPC method parameters and return types specified as protocol buffer messages:
  
```example.proto
service Greeter {
 rpc SayHello (HelloRequest) returns (HelloReply){}
}

message HelloRequest {
 string name = 1;
}

message HelloReply {
 string message 1;
}
```
