# Grooter

## Goal

I want to create a router that is receiving requests and forwarding them to the desired recipients based on a set of rules defined in a infra-as-code way.

## tasks

### HTTP Server Setup:

- [X] Create a basic HTTP server using the standard net/http package in Go.
- [X] Set up routes and handlers to handle incoming HTTP requests.
- [X] Implement a basic request forwarding mechanism to handle routing.

### Dynamic Configuration:

- [] Design a configuration mechanism that allows users to define routing rules dynamically.
- [] Decide on a configuration format (e.g., YAML, JSON) and create a parser to read and interpret the configuration.
- [] Implement a way to update the configuration without restarting the proxy.


### Reverse Proxying:

- [] Implement a reverse proxy functionality to forward incoming requests to appropriate backend services.
- [] Handle URL path rewriting during proxying.
- [] Explore how to handle request and response headers during proxying.

### Load Balancing:

- [] Implement load balancing algorithms such as round-robin, least connections, etc.
- [] Decide on a mechanism to track backend server health and availability.
- [] Handle failed requests and retries.

### Middleware:

- [] Create a middleware architecture to allow users to modify requests and responses.
- [] Implement middleware modules for common use cases like authentication, rate limiting, logging, etc.

### Dynamic Service Discovery:

- [] Integrate a service discovery mechanism to automatically detect backend services.
- [] Implement a way to add or remove services dynamically.

### TLS Termination:

- [] Implement TLS termination to handle HTTPS requests.
- [] Explore options for managing SSL certificates.

### Metrics and Monitoring:

- [] Integrate a metrics system to collect and monitor proxy-related data.
- [] Decide on metrics to collect, such as request rates, latency, errors, etc.

### Logging:

- [] Set up logging to capture relevant information about requests and proxy behavior.
- [] Implement different log levels and log output formats.

### Command-Line Interface (CLI):

- [] Develop a CLI tool to manage and configure the proxy.
- [] Provide commands to start, stop, reload, and manage proxy instances.

### Documentation:

- [] Create comprehensive documentation explaining how to use and configure your proxy.
- [] Include examples and best practices for setting up routing, load balancing, and middleware.

### Testing and Quality Assurance:

- [] Write unit tests for critical components of the proxy.
- [] Conduct integration tests to ensure all features work together seamlessly.
- [] Perform benchmarking to evaluate the proxy's performance and scalability.
