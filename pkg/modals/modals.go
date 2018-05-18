package modals

//go:generate protoc -I . -I ../../vendor --gogo_out=. librarian.proto
//go:generate easyjson -all librarian.pb.go
//go:generate protoc -I . -I ../../vendor --gogo_out=. attendance.proto
//go:generate easyjson -all attendance.pb.go
