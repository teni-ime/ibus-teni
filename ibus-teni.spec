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
%define version      1.5.3


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
BuildRequires:  go, libX11-devel
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
* Sat Nov 24 2018 Nguyen Cong Hoang <hoangnc.jp@gmail.com> - 1.5.3
- Build cho Ubuntu 20.04
* Sat Nov 24 2018 Nguyen Cong Hoang <hoangnc.jp@gmail.com> - 1.5.2
- Sửa kiểu gõ tự do không kiểm tra chính tả
* Fri Nov 23 2018 Nguyen Cong Hoang <hoangnc.jp@gmail.com> - 1.5.1
- Sửa kiểu gõ tự do không kiểm tra chính tả
* Sun Nov 4 2018 Nguyen Cong Hoang <hoangnc.jp@gmail.com> - 1.5.0
- Thêm kiểu gõ [Telex] (kiểu gõ này cho phép dùng phím [])
- Sửa lỗi không xóa hết chữ trên Telegram
- Sửa lỗi con trỏ chuột nhảy về đầu dòng trên Facebook chat
* Mon Oct 22 2018 Nguyen Cong Hoang <hoangnc.jp@gmail.com> - 1.4.2
- Sửa lỗi mất chữ khi đang gõ
- Sửa lỗi con trỏ chuột nhảy về đầu dòng trên Facebook chat
* Sun Oct 21 2018 Nguyen Cong Hoang <hoangnc.jp@gmail.com> - 1.4.1
- Cập nhật từ điển
- Sửa lỗi mất chữ khi đang gõ
* Sun Oct 7 2018 Nguyen Cong Hoang <hoangnc.jp@gmail.com> - 1.4.0
- Thêm lựa chọn "Giữ nhiều chữ", cho phép sửa dấu những chữ đã gõ xong
- Thêm lựa chọn "Đúng chính tả", cho phép bỏ qua kiểm tra chính tả
* Sun Sep 23 2018 Nguyen Cong Hoang <hoangnc.jp@gmail.com> - 1.3.3
- Tối ưu chức năng "Loại trừ ứng dụng"
- Sửa lỗi không gõ được sau khi click chuột đi chỗ khác
- Thêm xử lý xóa pre-edit khi click chuột
* Sat Sep 15 2018 Nguyen Cong Hoang <hoangnc.jp@gmail.com> - 1.3.2
- Tối ưu chức năng "Loại trừ ứng dụng"
- Sửa lỗi không gõ được trên FreeOffice và Wine
* Sun Sep 9 2018 Nguyen Cong Hoang <hoangnc.jp@gmail.com> - 1.3.1
- Sửa lỗi chức năng "Loại trừ ứng dụng"
* Tue Sep 4 2018 Nguyen Cong Hoang <hoangnc.jp@gmail.com> - 1.3.0
- Thêm chức năng "Loại trừ ứng dụng"
- Cập nhật từ điển (bổ sung ~700 từ)
* Sun Aug 26 2018 Nguyen Cong Hoang <hoangnc.jp@gmail.com> - 1.2.2
- Thay đổi xử lý commit: forward tất cả các phím khi commit
- Sửa lỗi khôi phục phím khi nhấn phím dấu 2 lần
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