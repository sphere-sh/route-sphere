FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/main ./cmd/server/main.go

FROM alpine:latest AS production

ENV ROUTE_SPHERE_STATIC_CONFIGURATION_PATH=/etc/sphere/route/main.yaml

# Create `/etc/sphere/route/dynamic`
#
RUN mkdir -p /etc/sphere/route/dynamic

# Put `.docker/main.yaml` in `/etc/sphere/route/`
#
COPY .docker/configuration/main.yaml /etc/sphere/route/main.yaml



WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]