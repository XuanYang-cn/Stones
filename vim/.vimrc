" Center view on the search result
noremap n nzz
noremap N Nzz

set hlsearch
noremap <F8> :nohl<CR>
inoremap <F8> :nohl<CR>

set ignorecase
set smartcase
set backspace=2

" Plugs

call plug#begin('~/.vim/plugged')

"auto complete codes
Plug 'ycm-core/YouCompleteMe'

" audo complete pairs
Plug 'jiangmiao/auto-pairs'

" Grammer
Plug 'dense-analysis/ale'

" Markdown grammer
Plug 'iamcco/mathjax-support-for-mkdp'

" Markdown preview
Plug 'iamcco/markdown-preview.vim'

Plug 'morhetz/gruvbox'

" schema color
Plug 'preservim/nerdtree'

" simplefold
Plug 'tmhedberg/SimpylFold'

" Jump
Plug 'kien/ctrlp.vim'

Plug 'vim-airline/vim-airline'

" Auto comment
Plug 'scrooloose/nerdcommenter'

call plug#end()

nmap <F2> :NERDTreeToggle<cr>

" ale
let g:ale_fix_on_save = 0
let g:ale_lint_on_insert_leave = 0
let g:ale_lint_on_enter = 0
let g:ale_echo_cursor = 1
let g:ale_completion_enabled = 0
let g:ale_sign_column_always = 1
let g:airline#extensions#ale#enabled = 1
let g:ale_echo_msg_error_str = 'E'
let g:ale_echo_msg_warning_str = 'W'
let g:ale_echo_msg_format = '[%linter%] %s [%severity%]'
let g:ale_linters = {'python': ['flake8'],
            \ 'go': ['gofmt', 'golint', 'go vet'],
            \        'zsh':['shell'],
            \        'cpp':['cpplint']}

nmap <silent> <leader>k <Plug>(ale_previous_wrap)
nmap <silent> <leader>j <Plug>(ale_next_wrap)

" gruvbox
colorscheme gruvbox
set background=dark

" simplefold
let g:SimpylFold_docstring_preview = 0

" YouCompleteMe
let g:ycm_autoclose_preview_window_after_completion=1
let g:ycm_global_ycm_extra_conf = '~/.ycm_extra_conf.py'
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
# fold and unfold
nnoremap <F9> za
vnoremap <F9> zf

" set encoding
set colorcolumn=110
highlight ColorColumn ctermbg=darkgray
set encoding=utf-8


" split navigation
nnoremap <C-J> <C-W><C-J>
nnoremap <C-K> <C-W><C-K>
nnoremap <C-L> <C-W><C-L>
nnoremap <C-H> <C-W><C-H>

noremap <Leader>y "*y
noremap <Leader>p "*p
noremap <Leader>Y "+y
noremap <Leader>P "+p

"
" NerdComment
" Add spaces after comment delimiters by default
let g:NERDSpaceDelims = 1

" Use compact syntax for prettified multi-line comments
let g:NERDCompactSexyComs = 1

" Align line-wise comment delimiters flush left instead of following code
" indentation
let g:NERDDefaultAlign = 'left'

" Add your own custom formats or override the defaults
" let g:NERDCustomDelimiters = { 'c': { 'left': '/**','right': '*/'  }  }

" Allow commenting and inverting empty lines (useful when commenting a region)
let g:NERDCommentEmptyLines = 1

" Enable trimming of trailing whitespace when uncommenting
let g:NERDTrimTrailingWhitespace = 1

" Enable NERDCommenterToggle to check all selected lines is commented or not
let g:NERDToggleCheckAllLines = 1

vnoremap // y/\V<C-R>=escape(@",'/\')<CR><CR>
