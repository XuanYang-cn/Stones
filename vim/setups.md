# Ubuntu&MacOS setup

## 1. Setup ubuntu

### User with root auth

## 2. Setup shell

### 2.1 Terminal

#### Install Terminator for Ubuntu

[Terminator](https://github.com/gnome-terminator/terminator) is only avaliable on Linux.

One can download it from official source, surpported source and package name can be found [here](https://repology.org/project/terminator/versions)

```shell
# for ubuntu
apt install terminator
```

#### Install iTerm2 for MacOS
[iTerm2](https://iterm2.com/) is a terminal emulator for macOS that does amazing things.

### 2.2 Install Zsh
Install [Zsh](https://zsh.org) with the package manager of your choice

```shell
# for ubuntu
apt install zsh

# for macOS
brew install zsh
```
Change default shell:

```shell
chsh -s $(which zsh)
```

Test with:
```shell
$SHELL --version
```

### 2.3 Install oh-my-zsh
Install [Oh-My-Zsh](https://github.com/robbyrussell/oh-my-zsh)

```shell
sh -C "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
```

Oh-my-zsh will generate a new `.zshrc` and backup the old one.

So remember to copy my own `.zshrc` AFTER downloading oh-my-zsh.

### 2.4 Setup oh-my-zsh

#### install plugins

Oh-my-zsh has so many [themes](https://github.com/ohmyzsh/ohmyzsh/wiki/Themes) and [plugins](https://github.com/ohmyzsh/ohmyzsh/wiki/Plugins).

**Themes:** [powerlevel10k](https://github.com/romkatv/powerlevel10k#oh-my-zsh)

**Plugins:** [zsh-syntax-highlighting](https://github.com/zsh-users/zsh-syntax-highlighting),[history-search-multi-word](https://github.com/zdharma/history-search-multi-word)

Forget about how they should be installed on their readme, just clone these repositories into `~/.oh-my-zsh/plugins`,
and add their repository name into `~/.zshrc` `plugin`

```
plugins=(
    git
    zsh-syntax-highlighting
    history-search-multi-word
)
```

### 2.5 Setup environment

#### Setup python virtualenv
#### Setup go env
#### Download cmake

## 3. Setup vim
### 3.1 Download vim
### 3.2 configure plugs for vim
#### Install plug
- [Download](https://github.com/junegunn/vim-plug) `plug.vim` and put it in the `autoload` directory
```shell
# Unix
curl -fLo ~/.vim/autoload/plug.vim --create-dirs \
https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
```
#### :PlugInstall
#### Compile YouCompleteMe
#### Download gtags and ripgrep

