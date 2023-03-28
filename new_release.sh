# update the next version in the variables file (variable CURRENT_VERSION)

rm build/bin/*

wails build -platform "windows/amd64" -nsis -o "hatt-windows-amd64.exe"

mv ./build/bin/hatt-amd64-installer.exe ./build/bin/hatt-windows-amd64-installer.exe

wails build -platform "linux/amd64" -o "hatt-linux-amd64"