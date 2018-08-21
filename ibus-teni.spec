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

#release info -----------------------------------------------------------------


%define engine_name  teni
%define package_name ibus-%{engine_name}
%define version      1.2.1


#install directories ----------------------------------------------------------
%define engine_dir   /usr/share/%{package_name}
%define ibus_dir     /usr/share/ibus
%define ibus_cpn_dir /usr/share/ibus/component
%define usr_lib_dir  /usr/lib


#package info -----------------------------------------------------------------
Name:           ibus-%{engine_name}
Version:        %{version}
Release:        1
Summary:        A Vietnamese IME for IBus
License:        GPL-3.0
Group:          System/Localization
URL:            https://github.com/teni-ime/ibus-teni
Packager:       Nguyen Cong Hoang <hoangnc.jp@gmail.com>
BuildRequires:  go
Requires:       ibus
Provides:       locale(ibus:vi)
Source0:        %{package_name}-%{version}.tar.gz

%description
A Vietnamese IME for IBus using Teni-IME
Bộ gõ tiếng Việt cho IBus sử dụng Teni-IME

%global debug_package %{nil}
%prep
%setup


%build
make build


%install
make DESTDIR=%{buildroot} install


%files
%defattr(-,root,root)
%doc README.md LICENSE MAINTAINERS
%dir %{ibus_dir}
%dir %{ibus_cpn_dir}
%dir %{engine_dir}
%{engine_dir}/*
%{ibus_dir}/component/%{engine_name}.xml
%{usr_lib_dir}/ibus-engine-%{engine_name}


%clean
cd ..
rm -rf %{package_name}-%{version}
rm -rf %{buildroot}


%changelog
* Tue Aug 21 2018 Nguyen Cong Hoang <hoangnc.jp@gmail.com> - 1.2.1
- Sửa lỗi khôi phục phím w, [, ] trên kiểu gõ Telex
- Bổ sung danh sách từ tiếng Việt: xịn
* Fri Aug 17 2018 Nguyen Cong Hoang <hoangnc.jp@gmail.com> - 1.2.0
- Thêm kiểu gõ Telex (cho phép gõ nhanh bằng w, [,])
- Sửa lỗi mất gợi ý khi gõ trên thanh địa chỉ của Chrome
* Thu Aug 9 2018 Nguyen Cong Hoang <hoangnc.jp@gmail.com> - 1.1.0
- Thêm xử lý nhanh chóng khôi phục phím khi gõ từ không có trong tiếng Việt
- Sửa lỗi mất space khi gõ trên Dropbox Paper
* Sun Jul 29 2018 Nguyen Cong Hoang <hoangnc.jp@gmail.com> - 1.0.0
- Phiên bản hoàn thiện chính thức: ibus-teni v1.0.0
* Fri Jun 22 2018 Nguyen Cong Hoang <hoangnc.jp@gmail.com> - 0.1
- Phiên bản chính thức đầu tiên: ibus-teni v0.1