# Inventory
The Inventory service is a [gRPC](https://grpc.io/) API, with [Redis](https://redis.io/) persistence, written in Go, and containerized with [Docker](https://www.docker.com/). Its purpose is to ensure accurate product inventory data to facilitate the product sale process. It is a fundamental part of our E-commerce platform within the [Distributed Playground](https://github.com/DistributedPlayground) project. See the [project description](https://github.com/DistributedPlayground/project-description) for more details.

- [Service Architecture](#service-architecture)
- [Endpoint Description](#endpoint-description)
- [Running the Service](#running-the-service)
- [Testing the Service](#testing-the-service)
- [Distributed Playground Description](#distributed-playground)

## Service Architecture
I chose a gRPC API for this service because gRPC is well suited for service-to-service communication given its:
* **Ease of Developer Experience:** It allows the endpoint to be called in a syntax similar to calling a locally defined method.
* **High Performance:** Thanks to the smaller header sizes and multiplexing offered by HTTP/2.
* **Safety:** The consistency of a strongly typed schema can reduce potential errors.

The Inventory service is intended to be consumed by the Orders service *(Not Implemented)*.

## Endpoint Description
The Inventory service exposes a single endpoint:

| RPC Method | Request Type      | Response Type      | Description                 |
|------------|-------------------|--------------------|-----------------------------|
| `Get`      | `InventoryRequest`| `InventoryResponse`| Retrieves inventory count for a specified product ID.|

### Data Structures:
* **InventoryRequest**:
```protobuf
message InventoryRequest {
    string id = 1; // The product ID
}
```

* **InventoryResponse**:
```protobuf
message InventoryResponse {
    string count = 1; // The inventory count
}
```

## Running the Service
1. Ensure that you have cloned each repository within the DistributedPlayground organization. Each repository should be cloned into the same directory so that they are parallel, forming the complete project structure.
2. Follow the instructions detailed in the [infra](https://github.com/DistributedPlayground/infra) repository to set up and run the service.

## Testing the Service:
### List Functions:
`grpcurl -plaintext localhost:<PORT> list Inventory`
### Hit the Get Endpoint:
Note that this requires first setting an entry in Redis.
`echo '{"id":"<REDIS-PRODUCT-ID>"}' | grpcurl -plaintext -d @ localhost:<PORT> Inventory.Get`
