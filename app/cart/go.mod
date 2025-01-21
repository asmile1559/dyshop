module github.com/asmile1559/dyshop/app/cart

go 1.23.4

replace github.com/asmile1559/dyshop/pb => ../../pb

replace github.com/asmile1559/dyshop/utils => ../../utils

require github.com/asmile1559/dyshop/pb v0.0.0-00010101000000-000000000000

require (
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241015192408-796eee8c2d53 // indirect
	google.golang.org/grpc v1.69.4 // indirect
	google.golang.org/protobuf v1.36.3 // indirect
)
