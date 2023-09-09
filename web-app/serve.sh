#!/bin/bash
slingshot listen \
	--wasm=./webapp.wasm \
	--handler=callHandler \
	--http-port=7070
