#!/bin/bash

string=$(cat << EOF
Hello World
  🌍
  👋
EOF
)

echo $string

slingshot run \
    --wasm=./print.wasm \
	--handler=callHandler \
	--input="${string}"

