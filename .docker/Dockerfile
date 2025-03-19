FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/main ./cmd/server/main.go

FROM alpine:latest AS production

ENV ROUTE_SPHERE_DIRECTORY=/etc/route-sphere

RUN mkdir -p $ROUTE_SPHERE_DIRECTORY

# Copy the standard configuration file
# into the container
#
COPY .docker/configuration/route-sphere.yaml $ROUTE_SPHERE_DIRECTORY/route-sphere.yaml

# Create route-sphere user and group
#
RUN addgroup -S route-sphere && adduser -S route-sphere -G route-sphere
RUN chown -R route-sphere:route-sphere $ROUTE_SPHERE_DIRECTORY

# Create the application directory
#
RUN mkdir -p /app
RUN chown -R route-sphere:route-sphere /app


USER route-sphere

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]