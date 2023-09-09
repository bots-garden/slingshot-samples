# Web-App

> inspired by [https://www.wasm.builders/k33g_org/host-and-serve-a-lit-spa-with-webassembly-and-capsule-52ak
](https://www.wasm.builders/k33g_org/host-and-serve-a-lit-spa-with-webassembly-and-capsule-52ak)

## Build the webapp

```bash
cd resources
npm install
npm run build
cd ..
tinygo build -scheduler=none --no-debug \
    -o webapp.wasm \
    -target wasi main.go
```

## Serve the webapp

```bash
slingshot listen \
	--wasm=./webapp.wasm \
	--handler=callHandler \
	--http-port=7070
```
