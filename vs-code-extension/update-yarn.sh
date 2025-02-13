#!/bin/sh

yarn install --no-frozen-lockfile && yarn2nix > yarn.nix

