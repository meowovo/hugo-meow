---
title: "Vim配置"
date: 2020-08-12T11:51:41+08:00
draft: false
---

编译vim前的配置

```go
./configure --with-features=huge \es=huge \ 
    --enable-multibyte \
    --enable-rubyinterp=yes \
    --enable-python3interp=yes \
    --enable-perlinterp=yes \
		--enable-luainterp=yes
    --prefix=/usr/local/vim \
```

make && make install 

需自行建prefix处的文件夹，可直接chown为当前用户

解释：    --enable-python3interp=yes \   开启Python3支持

          --enable-luainterp=yes              开启lua支持，用于 'Shougo/neocomplete.vim' 插件

[.vimrc](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/ba60ee88-881b-4069-94e2-b3b6ef772015/.vimrc)

[.vimrc 预览](https://www.notion.so/vimrc-aa7d25ee8e214e70a43db97f78d93691)