rm build/bin/*

wails build -platform "windows/amd64" -nsis -o "hatt-windows-x86_64-portable.exe"

mv ./build/bin/hatt-amd64-installer.exe ./build/bin/hatt-windows-x86_64-installer.exe

wails build -platform "linux/amd64" -o "hatt-linux-x86_64"