# Route Sphere

Route Sphere is a lightweight, domain-based reverse proxy written in Go that efficiently handles incoming HTTP/HTTPS
requests and routes them to the appropriate backend services based on domain names.

## Features

- **TLS Support**: Serve multiple domains with different TLS certificates
- **Domain-based Routing**: Route requests to different backends based on domain name
- **Dynamic Configuration**: Update configuration without restarting the service
- **Easy to Use**: Simple YAML-based configuration

## Starting with Route Sphere

### 0. Creating an Route Sphere cloud account
Start by [registering](https://route.sphere.sh/register) a new account on Route Sphere cloud.

#### 0.1. Obtaining Client ID and Secret

#### 0.2. Create a new connection

#### 0.3. Downloading certificates for secure communication

### 1. Running Route Sphere in Docker

The easiest way to get started with Route Sphere is to use the official Docker image.

```yaml
services:
  route_sphere:
    image: bromanonld/route-sphere:latest
    ports:
      - "80:80"
      - "443:443" 
```

### 2. Your First Configuration Provider

### 3. Your First Domain Configuration

### 4. Listening to incoming requests

### 5. Adding domains

### 6. Connecting services