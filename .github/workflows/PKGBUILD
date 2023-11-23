# Maintainer: FrenchGithubUser < nomail >
pkgname='hatt-bin'
_pkgname="hatt"
pkgver=0.3.5
pkgrel=1
pkgdesc="Tool to search for files through multiple websites "
arch=('x86_64')
url="https://github.com/FrenchGithubUser/Hatt"
_assets_url="https://raw.githubusercontent.com/FrenchGithubUser/Hatt/main"
license=('GPL')
depends=("webkit2gtk")
#makedepends=("")
#optdepends=("")
provides=("$_pkgname")
conflicts=("$_pkgname")
source=("$_assets_url/frontend/public/images/hatt-logo.png" "$_assets_url/.pkg/hatt-bin.desktop")
source_x86_64=("${url}/releases/download/${pkgver}/hatt-linux-amd64")
sha256sums_x86_64=('SKIP')
sha256sums=('SKIP' 'SKIP')

package() {
        install -Dm755 "$_pkgname-linux-amd64" "$pkgdir/usr/bin/$_pkgname"
        install -Dm644 $srcdir/$pkgname.desktop $pkgdir/usr/share/applications/$_pkgname.desktop
        install -Dm644 $srcdir/$_pkgname-logo.png $pkgdir/usr/share/pixmaps/$_pkgname.png
}
