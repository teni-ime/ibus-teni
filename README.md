[IBus Teni - a Vietnamese Input Method Editor for IBus](https://github.com/teni-ime/ibus-teni)
===================================
[![Build Status](https://travis-ci.org/teni-ime/ibus-teni.svg?branch=master)](https://travis-ci.org/teni-ime/ibus-teni)
[![GitHub release](https://img.shields.io/github/release/teni-ime/ibus-teni.svg)](https://github.com/teni-ime/ibus-teni/releases/latest)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

Copyright 2018, Nguyen Cong Hoang <<hoangnc.jp@gmail.com>>.

IBus Teni is a Vietnamese Input Method Editor (IME) for IBus.

IBus Teni là một bộ gõ tiếng Việt cho IBus.


Teni là gì ?
------------
* Teni là kết hợp **Te**lex và V**ni** - 2 kiểu gõ tiếng Việt phổ biến nhất.
* Teni cũng là kiểu gõ mặc định của bộ gõ này, vừa gõ được Telex, vừa gõ được Vni.

### Sơ lượt tính năng
* Chỉ bảng mã Unicode
* 2 kiểu gõ: 
  * **Kiểu gõ Teni**
  * **Kiểu gõ Vni**
* 2 kiểu đánh dấu thanh:
  * **Dấu thanh chuẩn**
  * **Dấu thanh kiểu mới**
* Gõ dấu tự do, đánh dấu thanh bằng từ điển

Cài đặt và cấu hình
------------------
*Bên dưới là hướng dẫn cho Ubuntu 18 (các distro khác: [wiki](https://github.com/teni-ime/ibus-teni/wiki))*

### Cài đặt

1. Tải file package: [ibus-teni-version.deb]()

2. Double click vào file để cài hoặc chạy command:

   `sudo dpkg -i ibus-teni-version.deb`

3. Restart IBus:

   `ibus restart`
    
### Cấu hình
1. [Keyboard input method system: IBus](https://github.com/teni-ime/ibus-teni/wiki/H%C6%B0%E1%BB%9Bng-d%E1%BA%ABn-c%E1%BA%A5u-h%C3%ACnh#1-keyboard-input-method-system-ibus)
2. [Add an input source: Vietnamese(Teni)](https://github.com/teni-ime/ibus-teni/wiki/H%C6%B0%E1%BB%9Bng-d%E1%BA%ABn-c%E1%BA%A5u-h%C3%ACnh#2-add-an-input-source-vietnameseteni)
    
### Gỡ bỏ
```
sudo apt remove ibus-teni -y
ibus restart
```

Sử dụng
-------------
* Dùng phím tắt mặc định của IBus (thường là ⊞Win+Space) để chuyển đổi giữa các bộ gõ
* IBus-Teni dùng pre-edit để xử lý phím gõ, mặc định sẽ có gạch chân chữ khi đang gõ
* **Khi pre-edit chưa kết thúc mà nhấn chuột vào chỗ khác thì có 3 khả năng xảy ra tùy chương trình**
    * **Chữ đang gõ bị mất**
    * **Chữ đang gõ được commit vào vị trí mới con trỏ**
    * **Chữ đang gõ được commit vào vị trí cũ**
* **Vì vậy đừng quên commit: khi gõ chỉ một chữ, hoặc chữ cuối câu, hoặc sửa chữ giữa câu: nhấn Ctrl hoặc phím mũi tên (↑→↓←) để commit.**
         

Các phiên bản
------------
* Phiên bản thử nghiệm không công khai hoàn thành vào cuối tháng 5/2018
* Phiên bản thử nghiệm công khai dự kiến phát hành vào đầu tháng 7/2018
* Phiên bản chính thức dự kiến phát hành sau thử nghiệm công khai 2 tuần

Xem trang [release](https://github.com/teni-ime/ibus-teni/releases) để biết chi tiết các phiên bản đã phát hành.

Giấy phép
-------
Toàn bộ code IBus Teni được viết bởi Nguyen Cong Hoang và những người đóng góp được phát hành dưới giấy phép 
[GNU General Public License version 3](https://opensource.org/licenses/GPL-3.0).

Code trong thư mục [src/ibus-teni/vendor](src/third_party) là của các bên thứ 3,
xem các thông báo bản quyền trong từng thư mục con.

* godbus: xem [src/ibus-teni/vendor/github.com/godbus/dbus/README.markdown](src/ibus-teni/vendor/github.com/godbus/dbus/README.markdown)
* goibus: xem [src/ibus-teni/vendor/github.com/sarim/goibus/README.md](src/ibus-teni/vendor/github.com/sarim/goibus/README.md)


Dữ liệu từ điển trong thư mục [dict](dict): xem [dict/LICENSE](dict/LICENSE)
* [Dữ liệu từ điển tiếng Việt của Ho Ngoc Duc](http://www.informatik.uni-leipzig.de/~duc/Dict/)
* [Danh sách viết tắt trong tiếng Việt của QUOC-HUNG NGO](https://sites.google.com/site/ngo2uochung/research/dsviettat-tieng-viet)
