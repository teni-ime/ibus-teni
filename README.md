[IBus Teni - a Vietnamese Input Method Editor for IBus](https://github.com/teni-ime/ibus-teni)
===================================
[![Build Status](https://travis-ci.org/teni-ime/ibus-teni.svg?branch=master)](https://travis-ci.org/teni-ime/ibus-teni)
[![GitHub release](https://img.shields.io/github/release/teni-ime/ibus-teni.svg)](https://github.com/teni-ime/ibus-teni/releases/latest)
[![License: GPL v3](https://img.shields.io/badge/License-GPL%20v3-blue.svg)](https://opensource.org/licenses/GPL-3.0)
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/teni-ime/ibus-teni/wiki/H%C6%B0%E1%BB%9Bng-d%E1%BA%ABn-g%C3%B3p-%C3%BD%2C-b%C3%A1o-l%E1%BB%97i)

Copyright 2018, Nguyen Cong Hoang <<hoangnc.jp@gmail.com>>.

IBus Teni is a Vietnamese Input Method Editor (IME) for IBus.

IBus Teni là một bộ gõ tiếng Việt cho IBus.


Teni là gì ?
------------
* Teni là kết hợp **Te**lex và V**ni** - 2 kiểu gõ tiếng Việt phổ biến nhất.
* Teni cũng là kiểu gõ mặc định của bộ gõ này, vừa gõ được Telex, vừa gõ được Vni.


### Sơ lược tính năng
* Chỉ bảng mã Unicode
* 3 kiểu gõ: 
  * **Kiểu gõ Teni** (Telex + Vni, không cho phép gõ nhanh ư, ơ bằng w, [, ])
  * **Kiểu gõ Vni**
  * **Kiểu gõ Telex** (cho phép gõ nhanh ư, ơ bằng w, [, ])
* 2 kiểu đánh dấu thanh:
  * **Dấu thanh chuẩn**
  * **Dấu thanh kiểu mới**
* Gõ dấu tự do, đánh dấu thanh bằng từ điển
* Có danh sách loại trừ ứng dụng không dùng bộ gõ

Cài đặt và cấu hình
------------------

### Cài đặt (Ubuntu)

```sh
sudo add-apt-repository ppa:teni-ime/ibus-teni
sudo apt-get update
sudo apt-get install ibus-teni
ibus restart
```

**Lệnh bên dưới cho phép đọc event chuột, không bắt buộc nhưng cần để ibus-teni hoạt động tốt**
```sh
sudo usermod -a -G input $USER
```


*Cài đặt cho các bản Linux khác và hướng dẫn cài đặt từ mã nguồn: [wiki](https://github.com/teni-ime/ibus-teni/wiki/H%C6%B0%E1%BB%9Bng-d%E1%BA%ABn-c%C3%A0i-%C4%91%E1%BA%B7t)*
    
### Cấu hình
1. [Keyboard input method system: IBus](https://github.com/teni-ime/ibus-teni/wiki/H%C6%B0%E1%BB%9Bng-d%E1%BA%ABn-c%E1%BA%A5u-h%C3%ACnh#1-keyboard-input-method-system-ibus)
2. [Add an input source: Vietnamese(Teni)](https://github.com/teni-ime/ibus-teni/wiki/H%C6%B0%E1%BB%9Bng-d%E1%BA%ABn-c%E1%BA%A5u-h%C3%ACnh#2-add-an-input-source-vietnameseteni)
    
### Gỡ bỏ
```
sudo apt remove ibus-teni
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
* **Vì vậy đừng quên commit: khi gõ chỉ một chữ, hoặc chữ cuối câu, hoặc sửa chữ giữa câu: nhấn phím *Ctrl* để commit.**
         

Các phiên bản
------------
* Phiên bản thử nghiệm không công khai hoàn thành vào cuối tháng 5/2018
* Phiên bản thử nghiệm công khai phát hành vào đầu tháng 7/2018
* Phiên bản chính thức phát hành vào ngày 29/7/2018

Xem trang [release](https://github.com/teni-ime/ibus-teni/releases) để biết chi tiết các phiên bản đã phát hành.

Góp ý và báo lỗi
--------------
Xem [hướng dẫn](https://github.com/teni-ime/ibus-teni/wiki/H%C6%B0%E1%BB%9Bng-d%E1%BA%ABn-g%C3%B3p-%C3%BD%2C-b%C3%A1o-l%E1%BB%97i)

Giấy phép
-------
Toàn bộ code IBus Teni được viết bởi Nguyen Cong Hoang và những người đóng góp được phát hành dưới giấy phép 
[GNU General Public License version 3](https://opensource.org/licenses/GPL-3.0).

Code trong thư mục [src/ibus-teni/vendor](src/ibus-teni/vendor) là của các bên thứ 3,
xem các thông báo bản quyền trong từng thư mục con.

* godbus: xem [src/ibus-teni/vendor/github.com/godbus/dbus/README.markdown](src/ibus-teni/vendor/github.com/godbus/dbus/README.markdown)
* goibus: xem [src/ibus-teni/vendor/github.com/sarim/goibus/README.md](src/ibus-teni/vendor/github.com/sarim/goibus/README.md)

Dữ liệu từ điển trong thư mục [dict](dict): xem [dict/LICENSE](dict/LICENSE)
* [Dữ liệu từ điển tiếng Việt của Ho Ngoc Duc](http://www.informatik.uni-leipzig.de/~duc/Dict/)
* [Wiktionary tiếng Việt](https://vi.wiktionary.org/wiki/Trang_Chính)
* [Danh sách viết tắt trong tiếng Việt của QUOC-HUNG NGO](https://sites.google.com/site/ngo2uochung/research/dsviettat-tieng-viet)
