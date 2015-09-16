#Fuzz Tests for gogoprotobuf

These tests use https://github.com/dvyukov/go-fuzz/

If you would like to fuzz a new message type simply:
  - add it to fuzz2.proto or fuzz3.proto
  - run > make regenerate
  - run > make fuzz

Enums and Extensions are not supported, since multiple versions of the same protocol buffer messages are generated and the proto libraries don't support multiple types with the same name. 

Dependencies:
  - github.com/gogo/protobuf
  - github.com/golang/protobuf
  - github.com/dvyukov/go-fuzz and its dependencies
  - protoc