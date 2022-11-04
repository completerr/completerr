#!/usr/bin/env bash
if [ ! -f /config/config.yaml ]; then
  cp ./config.yaml.example /config/config.yaml
fi

./server