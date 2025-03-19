#!/bin/bash

docker build --network=host -t bromanonld/route-sphere:latest -f .docker/Dockerfile .
#docker push bromanonld/route-sphere:latest
