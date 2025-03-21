FROM golang:1.23-alpine AS builder

WORKDIR /app/server

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/server/main ./cmd/server/main.go


WORKDIR /app/cli

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/cli/main ./cmd/cli/main.go


# Build configuration API binary
#
WORKDIR /app/configuration-api

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/configuration-api/sphere-configuration ./cmd/configuration-api/main.go


FROM alpine:latest AS production

ENV ROUTE_SPHERE_DIRECTORY=/etc/route-sphere

RUN mkdir -p /etc/route-sphere/configurations
COPY .docker/configuration/route-sphere.yaml /etc/route-sphere/route-sphere.yaml

# Create route-sphere user and group
#
RUN addgroup -S route-sphere && adduser -S route-sphere -G route-sphere
RUN chown -R route-sphere:route-sphere $ROUTE_SPHERE_DIRECTORY
# Create the application directory
#
RUN mkdir -p /app
RUN chown -R route-sphere:route-sphere /app

# supervisord
#
RUN apk add --no-cache supervisor
COPY .docker/supervisord/supervisord.conf /etc/supervisord.conf
RUN mkdir -p /var/log/supervisord

WORKDIR /app

COPY --from=builder /app/server/main .
COPY --from=builder /app/cli/main /bin/route-sphere
COPY --from=builder /app/configuration-api/sphere-configuration /bin/sphere-configuration

CMD ["/usr/bin/supervisord"]
