#!/usr/bin/env bash

USER_ID=$(id -u)
GROUP_ID=$(id -g)
USERNAME=$(whoami)

export USER_ID
export GROUP_ID
export USERNAME

xhost +SI:localuser:${USERNAME}

docker compose -f docker-compose.dev.yml build --build-arg USER_ID=${USER_ID} --build-arg GROUP_ID=${GROUP_ID} --build-arg USERNAME=${USERNAME}
docker compose -f docker-compose.dev.yml up -d
