#!/bin/bash
slingshot run \
    --wasm=./print.wasm \
	--handler=callHandler \
	--input="🤓 I'm a geek"
