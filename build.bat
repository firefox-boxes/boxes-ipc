mkdir dist
go build -o dist/boxes-ipc.exe %*
go build -o dist/boxes.exe launcher/Launcher.go