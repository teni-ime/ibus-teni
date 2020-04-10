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

engine_name=teni
ibus_e_name=ibus-engine-$(engine_name)
pkg_name=ibus-$(engine_name)
version=1.5.3

engine_dir=/usr/share/$(pkg_name)
ibus_dir=/usr/share/ibus

rpm_src_dir=~/rpmbuild/SOURCES
tar_file=$(pkg_name)-$(version).tar.gz
rpm_src_tar=$(rpm_src_dir)/$(tar_file)
tar_options_src=--transform "s/^\./$(pkg_name)-$(version)/" --exclude={"*.tar.gz",".git",".idea"} .

test:
	GOPATH=$(CURDIR) GOCACHE=/tmp go test teni


cover:
	GOPATH=$(CURDIR) GOCACHE=/tmp go test --cover -c -o test_teni_linux teni
	./test_teni_linux -test.coverprofile=teni_cover.out
	GOPATH=$(CURDIR) GOCACHE=/tmp go tool cover -html=teni_cover.out -o teni_cover.html
	rm -f test_teni_linux teni_cover.out


build:
	GOPATH=$(CURDIR) GOCACHE=/tmp go build -ldflags "-w -s" -o $(ibus_e_name) ibus-$(engine_name)


dict-gen:
	cd src/dict-gen && dep ensure -update
	GOPATH=$(CURDIR) GOCACHE=/tmp go build -o dict_gen_linux dict-gen
	./dict_gen_linux
	rm -f dict_gen_linux


tdata-gen:
	go run test-data/test-data-gen.go
	rm test-data/vietnamese.new.dict.telexw.tdata
	rm test-data/vietnamese.sp.dict.telex1.tdata
	rm test-data/vietnamese.sp.dict.telex2.tdata
	rm test-data/vietnamese.sp.dict.telex3.tdata
	rm test-data/vietnamese.sp.dict.telexw.tdata
	rm test-data/vietnamese.std.dict.telexw.tdata


clean:
	rm -f ibus-engine-* *_linux *_cover.html go_test_* go_build_* test *.gz test
	rm -f debian/files
	rm -rf debian/debhelper*
	rm -rf debian/.debhelper
	rm -rf debian/ibus-teni*


install: build
	mkdir -p $(DESTDIR)$(engine_dir)
	mkdir -p $(DESTDIR)/usr/lib/
	mkdir -p $(DESTDIR)$(ibus_dir)/component/

	cp -R -f except.tmpl.txt icon.png wm.bash dict $(DESTDIR)$(engine_dir)
	cp -f $(ibus_e_name) $(DESTDIR)/usr/lib/
	cp -f $(engine_name).xml $(DESTDIR)$(ibus_dir)/component/


uninstall:
	sudo rm -rf $(DESTDIR)$(engine_dir)
	sudo rm -f $(DESTDIR)/usr/lib/$(ibus_e_name)
	sudo rm -f $(DESTDIR)$(ibus_dir)/component/$(engine_name).xml


src: clean
	tar -zcf $(DESTDIR)/$(tar_file) $(tar_options_src)
	cp -f $(pkg_name).spec $(DESTDIR)/
	cp -f $(pkg_name).dsc $(DESTDIR)/
	cp -f debian/changelog $(DESTDIR)/debian.changelog
	cp -f debian/control $(DESTDIR)/debian.control
	cp -f debian/rules $(DESTDIR)/debian.rules
	cp -f PKGBUILD $(DESTDIR)/PKGBUILD


rpm: clean
	tar -zcf $(rpm_src_tar) $(tar_options_src)
	rpmbuild $(pkg_name).spec -ba


#start ubuntu docker:   docker  run  -v `pwd`:`pwd` -w `pwd` -i -t  ubuntu bash
#install buildpackages: apt update && apt install dh-make golang libx11-dev -y
deb: clean
	dpkg-buildpackage


.PHONY: test build clean build install uninstall src rpm deb
