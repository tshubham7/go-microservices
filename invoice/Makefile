.PHONY: protos


protos: 
	protoc -I protos/ protos/invoice.proto --go_out=plugins=grpc:protos/invoice


grpcurl-test:
	grpcurl --plaintext localhost:9002 describe
	grpcurl --plaintext localhost:9002 describe invoice.Invoice.Create
	grpcurl --plaintext localhost:9002 describe invoice.CreateRequest

grpcurl-createInvoice:
	grpcurl --plaintext -d '{"Action": "create", "UserID": "1"}' localhost:9002 invoice.Invoice.Create