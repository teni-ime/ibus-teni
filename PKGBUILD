#
# Teni-IME - A Vietnamese Input method editor
# Copyright (C) 2018 Nguyen Cong Hoang <hoangnc.jp@gmail.com>
# This file is part of Teni-IME.
#
# Teni-IME is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# Teni-IME is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with Teni-IME.  If not, see <http://www.gnu.org/licenses/>.
#

# Maintainer: Nguyen Cong Hoang <hoangnc.jp@gmail.com>
pkgname=ibus-teni
pkgver=1.5.3
pkgrel=1
pkgdesc='A Vietnamese IME for IBus'
arch=(any)
license=(GPL3)
url="https://github.com/teni-ime/ibus-teni"
depends=(ibus)
makedepends=('go' 'libx11')
source=($pkgname-$pkgver.tar.gz)
md5sums=('SKIP')
options=('!strip')

build() {
  cd "$pkgname-$pkgver"

  make
}


package() {
  cd "$pkgname-$pkgver"

  make DESTDIR="$pkgdir/" install
}
