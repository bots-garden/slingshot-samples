
result1=$(slingshot run \
  --wasm=./one/one.wasm \
  --handler=callHandler \
  --input="🏁"
)

result2=$(slingshot run \
  --wasm=./two/two.wasm \
  --handler=callHandler \
  --input="${result1}"
)

result3=$(slingshot run \
  --wasm=./three/three.wasm \
  --handler=callHandler \
  --input="${result2}"
)

echo ${result3}
