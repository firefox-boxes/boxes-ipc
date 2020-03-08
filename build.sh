#!/bin/bash

mkdir dist
go build -o dist/boxes-ipc $@
go build -o dist/boxes launcher/Launcher.go