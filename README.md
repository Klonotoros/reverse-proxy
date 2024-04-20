# Go-powered reverse proxy server
Reverse Proxy Server is a lightweight HTTP reverse proxy server implemented in Go. It enables users to proxy incoming HTTP requests to specified backend servers while recording details of the requests made. It provides functionalities for downloading the recorded requests in CSV format and retrieving the total number of recorded requests.

## Features
- Reverse Proxy: Proxies incoming HTTP requests to specified backend servers.
- Record Requests: Records details of the requests made, including timestamp, URL, method, and client IP.
- Download Records: Download recorded requests in CSV format.
- Retrieve Total Records: Get the total number of recorded requests.

## Project Structure
- Service Layer: Implements business logic and interacts with repositories.
- Repository Layer: Handles data access and manipulation.
- Controller Layer: Handles HTTP request handling and routing using the Gin framework.
- Model Layer: Defines the data structures used in the application.

## Endpoints
### PORT:3101
- [ANY] /url: This endpoint acts as a versatile entry point for all HTTP methods (GET, POST, PUT, DELETE, etc.) and all URL paths. It serves as the primary entry point for the reverse proxy functionality. Incoming requests to this endpoint are proxied to the specified backend server whose URL is provided as part of the path. For example, a request to /example.com/api/data will be proxied to http://example.com/api/data. This flexible endpoint enables the reverse proxy server to dynamically handle requests for various backend services without the need to define specific routes for each one.
### PORT:3102
- [GET] /download  (Downloads recorded requests in CSV format)
- [GET] /info (Retrieves the total number of recorded requests)

#
