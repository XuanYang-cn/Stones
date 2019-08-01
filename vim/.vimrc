let mapleader=" "

" Center view on the search result
noremap n nzz
noremap N Nzz

set hlsearch
noremap <F8> :nohl<CR>
inoremap <F8> :nohl<CR>a

set ignorecase
set smartcase

" Plugs

call plug#begin('~/.vim/plugged')

"auto complete codes
Plug 'Valloric/YouCompleteMe'

" audo complete pairs
Plug 'jiangmiao/auto-pairs'

" Grammer
Plug 'w0rp/ale'

" Markdown grammer
Plug 'iamcco/mathjax-support-for-mkdp'

" Markdown preview
" :MarkdownPreview
Plug 'iamcco/markdown-preview.vim'

" schema color
Plug 'morhetz/gruvbox'

" simplefold
Plug 'tmhedberg/SimpylFold'

" Jump
Plug 'kien/ctrlp.vim'

Plug 'vim-airline/vim-airline'

call plug#end()

" ale
let g:ale_fix_on_save = 1
let g:ale_completion_enabled = 1
let g:ale_sign_column_always = 1
let g:airline#extensions#ale#enabled = 1
let g:ale_echo_msg_error_str = 'E'
let g:ale_echo_msg_warning_str = 'W'
let g:ale_echo_msg_format = '[%linter%] %s [%severity%]'
let g:ale_linters = {'python': ['flake8']}

" gruvbox
colorscheme gruvbox
set background=dark

" simplefold
let g:SimpylFold_docstring_preview = 1

" YouCompleteMe
let g:ycm_autoclose_preview_window_after_completion=1
map <leader>g  :YcmCompleter GoToDefinitionElseDeclaration<CR>
map <leader>mp :MarkdownPreview<CR>

filetype plugin indent on
syntax on
set nu
set nocompatible

set expandtab
set tabstop=4
set shiftwidth=4

" shutdown beep in vim
set noeb vb t_vb=

" enable fold
set foldenable
set foldmethod=indent
set foldlevel=99

" set encoding
set encoding=utf-8
"press space to fold/unfold code
nnoremap <F9> za
vnoremap <F9> zf

" split navigation
nnoremap <C-J> <C-W><C-J>
nnoremap <C-K> <C-W><C-K>
nnoremap <C-L> <C-W><C-L>
nnoremap <C-H> <C-W><C-H>

