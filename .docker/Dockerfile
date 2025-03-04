FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/main ./cmd/server/main.go

FROM alpine:latest AS production

ENV ROUTE_SPHERE_ENVIRONMENTS_PATH=/etc/route-sphere/environments.yaml

# Create `/etc/sphere/route/dynamic`
#
RUN mkdir -p /etc/route-sphere/dynamic

# Put `.docker/main.yaml` in `/etc/sphere/route/`
#
COPY .docker/configuration/environments.yaml /etc/route-sphere/environments.yaml


WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]