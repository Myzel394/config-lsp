#!/bin/sh

rm -rf ./result
rm -rf ./vs-code-extension/out

nix build .#"vs-code-extension"
mkdir ./vs-code-extension/out
cp ./result/* ./vs-code-extension/out
chmod 777 ./vs-code-extension/out -R

