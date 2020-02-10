module github.com/openbank/openbank

go 1.12

require (
	github.com/golang/protobuf v1.3.2
	github.com/grpc-ecosystem/grpc-gateway v1.11.3
	github.com/gunk/gunk v0.2.1-0.20200204073540-74a4de10b3fe
	github.com/gunk/opt v0.0.0-20190514110406-385321f21939
	golang.org/x/tools v0.0.0-20190716021316-fefcef05abb1 // indirect
	google.golang.org/genproto v0.0.0-20190708153700-3bdd9d9f5532
	google.golang.org/grpc v1.21.1
)

replace github.com/gunk/gunk => github.com/favadi/gunk v0.1.1-0.20200210044944-4770c0fc340b
