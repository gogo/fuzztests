#Fuzz Tests for gogoprotobuf

These tests use https://github.com/dvyukov/go-fuzz/

If you would like to fuzz a new message type simply:
  - add it to fuzz2.proto or fuzz3.proto
  - run > make regenerate
  - run > make fuzz

Enums are not supported, since multiple versions of the same protocol buffer messages are generated and the proto libraries don't support multiple enums with the same name.