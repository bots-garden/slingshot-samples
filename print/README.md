# Print

## Build

```bash
tinygo build -scheduler=none --no-debug \
    -o print.wasm \
    -target wasi main.go
```

## Run

```bash
slingshot run \
    --wasm=./print.wasm \
	--handler=callHandler \
	--input="🤓 I'm a geek"
```