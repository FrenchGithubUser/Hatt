# update the next version in the variables file (variable CURRENT_VERSION)

echo -n "Did you update the version in the variables file and the PKGBUILD ? " 
read updated_version

if [ $updated_version == 'y' ]; then

    rm build/bin/*

    wails build -platform "windows/amd64" -nsis -o "hatt-windows-amd64.exe"

    mv ./build/bin/hatt-amd64-installer.exe ./build/bin/hatt-windows-amd64-installer.exe

    wails build -platform "linux/amd64" -o "hatt-linux-amd64"
fi