# Setup ubuntu shell and vim

## Steps for configuring vim
1. Download `vim`
2. `.vimrc`
3. install plug
4. `YouCompleteMe`, TabNine, `ale`

## Steps for configuring shell
1. `Terminator`: thr robot future of terminals
2. Z-shell `zsh`
    - change default shell
3. `Oh-My-Zsh`
4. `powerline` and `powerline fonts`
5. `.zshrc`

---

### 1. Install Terminator

**For Ubuntu**

[Install Terminator](https://gnometerminator.blogspot.com/p/introduction.html)

**For MacOS**

Install iTerm2

### 2. Vim

- [Build vim from source](https://www.vim.org/git.php)
- Get `.vimrc`
- [Download](https://github.com/junegunn/vim-plug) `plug.vim` and put it in the `autoload` directory

```shell
# Unix
curl -fLo ~/.vim/autoload/plug.vim --create-dirs \
https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim 
```
- **:PlugInstall**
- Fix Bugs of `YouCompleteMe`, `ale`
- Set default editor as Vim

### 3. ZSH

- [Install Zsh](https://github.com/robbyrussell/oh-my-zsh/wiki/Installing-ZSH)
- [Install Oh-My-Zsh](https://github.com/robbyrussell/oh-my-zsh)
- Get `.zshrc` file
- Change default shell

```shell
chsh -s /bin/zsh
```
- [Install powerline and powerline fonts](https://github.com/powerline/fonts)
