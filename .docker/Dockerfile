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


FROM golang:1.23-alpine AS configuration_api_builder

ENV ROUTE_SPHERE_CONFIGURATION_CONFIG_PATH=/etc/route-sphere/route-sphere.yaml

# Build configuration API binary
#
WORKDIR /app/configuration-api

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/configuration-api/sphere-configuration ./cmd/configuration-api/main.go


FROM alpine:latest AS production

#   ---------------------------------------------------------------------------------
#
#   Install OpenRC with proper container settings
#
#   ---------------------------------------------------------------------------------
RUN apk add --no-cache openrc mdevd-openrc
RUN mkdir -p /run/openrc && \
    touch /run/openrc/softlevel && \
    sed -i 's/^\(tty\d\:\:\)/#\1/g' /etc/inittab && \
    sed -i 's/^\(rc_cgroup_mode=.*\)/#\1/g' /etc/rc.conf && \
    echo 'rc_cgroup_mode="hybrid"' >> /etc/rc.conf && \
    echo 'rc_sys="docker"' >> /etc/rc.conf && \
    sed -i '/getty/d' /etc/inittab


#   ---------------------------------------------------------------------------------
#
#   Configuration API setup
#
#   ---------------------------------------------------------------------------------
COPY --from=configuration_api_builder /app/configuration-api/sphere-configuration /bin/sphere-configuration

RUN addgroup -S sphere-configuration && adduser -S sphere-configuration -G sphere-configuration

# OpenRC service setup
#
COPY .docker/openrc/conf.d /etc/conf.d
COPY .docker/openrc/init.d/sphere-configuration /etc/init.d/sphere-configuration
RUN chmod +x /etc/init.d/sphere-configuration
RUN rc-update add sphere-configuration default

RUN touch /var/log/sphere-configuration.log
RUN chown sphere-configuration:sphere-configuration /var/log/sphere-configuration.log

RUN touch /var/log/sphere-configuration.err
RUN chown sphere-configuration:sphere-configuration /var/log/sphere-configuration.err

#   ---------------------------------------------------------------------------------
#
#   Route Sphere Server setup
#
#   ---------------------------------------------------------------------------------
COPY --from=builder /app/server/main /bin/sphere-server

RUN addgroup -S sphere-server && adduser -S sphere-server -G sphere-server

# OpenRC service setup
#
COPY .docker/openrc/init.d/sphere-server /etc/init.d/sphere-server
RUN chmod +x /etc/init.d/sphere-server
RUN rc-update add sphere-server default

RUN touch /var/log/sphere-server.log
RUN chown sphere-server:sphere-server /var/log/sphere-server.log

RUN touch /var/log/sphere-server.err
RUN chown sphere-server:sphere-server /var/log/sphere-server.err


#   ---------------------------------------------------------------------------------
#
#   CLI Application setup
#
#   ---------------------------------------------------------------------------------
COPY --from=builder /app/cli/main /bin/route-sphere



#   ---------------------------------------------------------------------------------
#
#   Add default configurations to the container.
#
#   ---------------------------------------------------------------------------------
RUN mkdir -p /etc/route-sphere/configuration-api

COPY .docker/configuration/route-sphere.yaml /etc/route-sphere/route-sphere.yaml


#   ---------------------------------------------------------------------------------
#
#   Container setup
#
#   ---------------------------------------------------------------------------------
EXPOSE 80 5166 443
CMD ["/sbin/init"]