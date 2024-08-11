#!/usr/bin/env bash

set -a
source .env
set +a

CLIPBOARD_CONTENT=$(xclip -selection clipboard -o)

docker run -e BOT_TOKEN=${BOT_TOKEN} -e CHAT_ID=${CHAT_ID} crawl-deepl:app "$CLIPBOARD_CONTENT"