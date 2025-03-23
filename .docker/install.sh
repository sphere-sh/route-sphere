#!/bin/bash


# Create `/etc/route-sphere` directory
#
sudo mkdir -p /etc/route-sphere
sudo chown -R $USER:$USER /etc/route-sphere

# Place `configuration/route-sphere.yaml` in `/etc/route-sphere`
#
sudo cp .docker/configuration/route-sphere.yaml /etc/route-sphere

# Build the CLI application and place it in `/bin`
#
go build -o route-sphere cmd/cli/main.go
sudo mv route-sphere /bin
sudo chown $USER:$USER /bin/route-sphere