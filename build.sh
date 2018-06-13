#!/bin/bash
# Get the name of the app
APP="${PWD##*/}"

# Get this tag as the version
VERSION=$(git describe --abbrev=0 --tags)

exec xgo \
  -go 1.10.x \
  -buildmode=plugin  \
  --targets=darwin/amd64 \
  --ldflags "-X plugin.version=${VERSION}" \
  ../$APP
