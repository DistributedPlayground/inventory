# Inventory
The Inventory service is a [gRPC](https://grpc.io/) API, with [Redis](https://redis.io/) persistence, written in Go, and containerized with [Docker](https://www.docker.com/). Its purpose is to ensure accurate product inventory data to facilitate the product sale process. It is a fundamental part of our E-commerce platform within the [Distributed Playground](https://github.com/DistributedPlayground) project.

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

# Distributed Playground
The purpose of this repo is to practice the development of distributed systems
## Project 1 - Ecommerce Platform
For my first project, I'll create an ecommerce platform. This platform will allow users to purchase goods that are maintained by us. This is *not* a two sided marketplace with our users being both companies and consumers -- at least for now. To simplify, we'll centrally define the products.

### Architecture

#### Functional Requirements
- Users should be able to make purchases using card payments
- Users should be able to see available products, including their remaining stock
- Users should be able to sort and filter products

#### Quality Attributes
- Should be designed to handle >10M users per day
- Should have 99.9% uptime
- Should allow for concurrent development from a large distributed team

#### Constraints
- The core services must be completed by 1 engineer (me) in < 1 month. 

#### Design
- We will use Golang as the primary backend language.

- We will provide an internal CP db optimized for writes in postgres
    - This db will be exposed to our systems through a REST API
- Implementing a psudo CQRS (Command Query Responsibility Segregation)
    - "Psudo" here because we really just want to segregrate the most read heavy users (customers) to a separate db and api optimized for their needs
    - This customer read db will be updated through a service that reads from a Kafka queue

#### Services
- **API Gateway**: *Not Implemented*
- [Products](https://github.com/DistributedPlayground/products): A REST API with a postgres db. It allows sellers to manage their products and product collections. This service publishes writes to a kafka queue.
- [Product Search](https://github.com/DistributedPlayground/product-search): A GraphQL API with mongodb. It allows customers to query the current products and collections. It reads updates from kafka to update the mongodb database.
- [Inventory](https://github.com/DistributedPlayground/inventory): A gRPC API with redis. It is intended to maintain the most up-to-date state of product inventory. It reads updates to product inventory made by sellers from kafka, and also writes updates to the product inventory made by customers through purchases.
- **Orders**: *Not Implemented* Orchestrator for *Payments*, *Fufillment*, and *Notifications*, state management with kafka
- **Order Recovery**: *Not Implemented*