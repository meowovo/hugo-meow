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

自己的vimrc:

```go
set nocompatible "关闭兼容模式
set backspace=indent,eol,start " 防止backspace不生效
set nu "设置行号
set cursorline "突出显示当前行
set cursorcolumn
set showmatch "显示括号匹配

set tabstop=4 "Tab长度为4空格
set shiftwidth=4 "自动缩进长度为4空格
set autoindent "继承前一行的缩进方式

set incsearch "实时搜索
set ignorecase "搜索时大小写不敏感

syntax enable
syntax on
filetype plugin indent on

au InsertLeave *.go,*.sh,*.php write "退出插入模式指定类型的文件自动保存

set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()
Plugin 'VundleVim/Vundle.vim'
Plugin 'junegunn/vim-easy-align'
Plugin 'scrooloose/nerdtree'
Plugin 'jistr/vim-nerdtree-tabs'
Plugin 'Xuyuanp/nerdtree-git-plugin'
Plugin 'majutsushi/tagbar'
Plugin 'airblade/vim-gitgutter'
Plugin 'vim-scripts/SuperTab'
Plugin 'Shougo/neocomplete.vim'
Plugin 'crusoexia/vim-monokai'
Plugin 'fatih/vim-go', { 'tag': '*' } "用于Go的插件，:GoInstallBinaries
Plugin 'dgryski/vim-godef' "输入gd跳转到当前代码的定义处
Plugin 'Valloric/YouCompleteMe' "需在YouComplete目录下用Python3 install一下，选择对应的语言 "例如Go: python3 install.py --go-completer
Plugin 'vim-airline/vim-airline'
Plugin 'SirVer/ultisnips'
Plugin 'honza/vim-snippets'
Plugin 'jiangmiao/auto-pairs'
call vundle#end()

" 开启24bit的颜色，开启这个颜色会更漂亮一些
set termguicolors
" 配色方案, 可以从上面插件安装中的选择一个使用
colorscheme monokai " 主题
set background=dark " 主题背景 dark-深色; light-浅色

autocmd FileType go nnoremap <buffer> gd :call GodefUnderCursor()<cr>

let g:go_fmt_command = "goimports" " 格式化将默认的 gofmt 替换
let g:go_list_type = "quickfix"

let g:go_version_warning = 1
let g:go_highlight_types = 1
let g:go_highlight_fields = 1
let g:go_highlight_functions = 1
let g:go_highlight_function_calls = 1
let g:go_highlight_operators = 1
let g:go_highlight_extra_types = 1
let g:go_highlight_methods = 1
let g:go_highlight_generate_tags = 1

let g:godef_split=2

"==============================================================================
" NERDTree 插件
"==============================================================================

" 打开和关闭NERDTree快捷键
map zz :NERDTreeToggle<CR>
" 显示行号
let NERDTreeShowLineNumbers=1
" 打开文件时是否显示目录
let NERDTreeAutoCenter=1
" 是否显示隐藏文件
let NERDTreeShowHidden=1
" 设置宽度
" let NERDTreeWinSize=31
" 忽略一下文件的显示
let NERDTreeIgnore=['\.pyc','\~$','\.swp']
" 打开 vim 文件及显示书签列表
let NERDTreeShowBookmarks=2

let g:nerdtree_tabs_open_on_console_startup=0

"==============================================================================
" majutsushi/tagbar 插件
"==============================================================================

" majutsushi/tagbar 插件打开关闭快捷键
nmap rr :TagbarToggle<CR>

let g:tagbar_type_go = {
\ 'ctagstype' : 'go',
\ 'kinds' : [
\ 'p:package',
\ 'i:imports:1',
\ 'c:constants',
\ 'v:variables',
\ 't:types',
\ 'n:interfaces',
\ 'w:fields',
\ 'e:embedded',
\ 'm:methods',
\ 'r:constructor',
\ 'f:functions'
\ ],
\ 'sro' : '.',
\ 'kind2scope' : {
\ 't' : 'ctype',
\ 'n' : 'ntype'
\ },
\ 'scope2kind' : {
\ 'ctype' : 't',
\ 'ntype' : 'n'
\ },
\ 'ctagsbin' : 'gotags',
\ 'ctagsargs' : '-sort -silent'
\ }

"==============================================================================
" nerdtree-git-plugin 插件
"==============================================================================
let g:NERDTreeIndicatorMapCustom = {
\ "Modified" : "✹",
\ "Staged" : "✚",
\ "Untracked" : "✭",
\ "Renamed" : "➜",
\ "Unmerged" : "═",
\ "Deleted" : "✖",
\ "Dirty" : "✗",
\ "Clean" : "✔︎",
\ 'Ignored' : '☒',
\ "Unknown" : "?"
\ }

let g:NERDTreeShowIgnoredStatus = 1

" tab 标签页切换快捷键
:nn <Leader>1 1gt
:nn <Leader>2 2gt
:nn <Leader>3 3gt
:nn <Leader>4 4gt
:nn <Leader>5 5gt
:nn <Leader>6 6gt
:nn <Leader>7 7gt
:nn <Leader>8 8gt
:nn <Leader>9 8gt
:nn <Leader>0 :tablast<CR>

"==============================================================================
" snippets 插件
"==============================================================================

let g:ycm_key_list_select_completion = ['<C-n>', '<space>']
let g:ycm_key_list_previous_completion = ['<C-p>', '<Up>']
let g:SuperTabDefaultCompletionType = '<C-n>'

let g:UltiSnipsExpandTrigger="<tab>"
let g:UltiSnipsJumpForwardTrigger="<c-b>"
let g:UltiSnipsJumpBackwardTrigger="<c-z>"
```