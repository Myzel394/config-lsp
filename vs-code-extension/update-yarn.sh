#!/bin/sh

cd vs-code-extension && yarn install --no-frozen-lockfile && yarn2nix > yarn.nix

