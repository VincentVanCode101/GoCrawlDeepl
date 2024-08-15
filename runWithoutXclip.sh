#!/usr/bin/env bash

ENV_FILE="$(dirname "$0")/.env"

args=("$@")
args_string="${args[*]}"

set -a
source $ENV_FILE
set +a


docker run -e TO_LANGUAGE=${TO_LANGUAGE} -e FROM_LANGUAGE=${FROM_LANGUAGE} -e BOT_TOKEN=${BOT_TOKEN} -e CHAT_ID=${CHAT_ID} crawl-deepl:app "$args_string"