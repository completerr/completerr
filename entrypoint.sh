#!/usr/bin/env sh
if [ ! -f /config/config.yaml ]; then
  cp ./config.example.yaml /config/config.yaml
fi

./server
