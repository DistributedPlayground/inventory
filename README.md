# To Create GRPC Code:
`protoc --go_out=api --go_opt=paths=source_relative --go-grpc_out=api --go-grpc_opt=paths=source_relative inventory.proto`

# To Test GRPC:
### List Functions:
`grpcurl -plaintext localhost:5555 list Inventory`
### Hit the Get Endpoint:
Note that this requires first setting an entry in Redis for id:fd97eef4-a9da-4854-99bf-5c1641b37669
`echo '{"id":"fd97eef4-a9da-4854-99bf-5c1641b37669"}' | grpcurl -plaintext -d @ localhost:5555 Inventory.Get`