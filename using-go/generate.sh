protoc -I src/ --go_out=src/ src/example/simple.proto
protoc -I src/ --go_out=src/ src/enum_example/enum_example.proto
protoc -I src/ --go_out=src/ src/complex/complex.proto