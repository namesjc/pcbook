# gRPC

gRPC is a high-performance open-source feature-rich RPC framework

RPC stands Remote Procedure Calls

- It is a protocol that allows a program to:

  Execute a procedure of another program located in other computer

  without the developer explicity coding details for the remote interaction

- In the client code, it looks like we're just calling a function fo the server code directly

- The client and server codes can be written in different languages

|                         | HTTP2      | HTTP1.1    |
| ----------------------- | ---------- | ---------- |
| TRANSFER PROTOCOL       | Binary     | Text       |
| HEADERS                 | Compressed | Plain text |
| MULTIPLEXING            | Yes        | No         |
| REQUEST PER CONNECTIONS | Multiple   | 1          |
| SERVER PUSH             | Yes        | No         |
| RELEASE YEAR            | 2015       | 1997       |

## gRPC type

UNARY

CLIENT STREAMING

SERVER STREAMING

BIDEIRECTIONAL STREAMING

![截屏2022-06-11 20.30.13](/Users/echo/Library/Application Support/typora-user-images/截屏2022-06-11 20.30.13.png)

## gRPC vs. REST

| FEATURE         | GRPC                       | REST                           |
| --------------- | -------------------------- | ------------------------------ |
| Protocol        | HTTP/2(fast)               | HTTP/1.1(slow)                 |
| Payload         | Protobuf (binary, small)   | JSON (text, large)             |
| API contract    | Strict, required (.proto)  | Loose, optional (OpenAPI)      |
| Code generation | Built-in (protoc)          | Third-part tools (Swagger)     |
| Security        | TLS/SSL                    | TLS/SSL                        |
| Streaming       | Bidirectional streaming    | Client --> server request only |
| Browser support | Limited (require gRPC-web) | Yes                            |

## Scenario:

- Microservices

  - Low latency and high throuthput communication

  - Strong API contract

- Polyglot environments
  - Code generation out of the box for many languages

- Point-to-point realtime communication
  - Excellent support bidirectional streaming

- Network constrained environment
  - Lightweight message format

## How to define a protocol message

- Name of message: UpperCamelCase

- Name of the field: lower_snake_case

  - Some scalar-value data types:

  - string, bool, bytes

  - Float, double

  - Int32, int64, uint32, uint64, sint32, sint64, etc.

- Data types can be user-defined enums or other messages
- Tags are more important than field names
  - Is an arbitrary integer
    - From 1 to 536,870,911 (or 2^29-1)
    - Except from 19,000 to 19,999 (reserved)
  - From 1 to 15 take 1 bytes
  - From 16 to 2047 take 2 bytes
  - Don't need to be in-order or sequential
  - Must be unique for same-level fields

## Serialize protobuf message

