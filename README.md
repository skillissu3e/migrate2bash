# ~/.bashrc

# If not running interactively, don't do anything
[[ $- != *i* ]] && return

# Prompt
PS1='[\u@\h \W]\$ '

# Default programs
export EDITOR="nvim"
export VISUAL="nvim"
export VIDEOPLAYER="mpv"
export TERMINAL="ghostty"
export IMAGEVIEWER="swayimg"
export BROWSER="brave"

# Follow XDG base dir specification
export XDG_CONFIG_HOME="$HOME/.config"
export XDG_DATA_HOME="$HOME/.local/share"
export XDG_CACHE_HOME="$HOME/.cache"

# History files
export LESSHISTFILE="$XDG_CACHE_HOME/less_history"
export PYTHON_HISTORY="$XDG_DATA_HOME/python/history"

# Moving other files and some other vars
export CARGO_HOME="$XDG_DATA_HOME/cargo"
export RUSTUP_HOME="$XDG_DATA_HOME/rustup"
export GOPATH="$XDG_DATA_HOME/go"
export GOBIN="$GOPATH/bin"
export GOMODCACHE="$XDG_CACHE_HOME/go/mod"
export NPM_CONFIG_PREFIX="$XDG_DATA_HOME/npm"
export NPM_CONFIG_CACHE="$XDG_CACHE_HOME/npm"
export DATE=$(date "+%-d %B %Y, %A, %T")
export MANPAGER="nvim +Man!"

export FZF_DEFAULT_OPTS="--style full --color 16 --layout=reverse --height 30% --preview='bat -p --color=always {}'"
export FZF_CTRL_R_OPTS="--style full --color 16 --layout=reverse --height 30% --no-preview"

# Important
export XDG_CURRENT_DESKTOP=sway
export XDG_SESSION_DESKTOP=sway
export XDG_SESSION_TYPE=wayland
export ELECTRON_OZONE_PLATFORM_HINT=wayland
export SDL_VIDEODRIVER=wayland
export QT_QPA_PLATFORM=wayland
export QT_QPA_PLATFORMTHEME=qt6ct
export QT_WAYLAND_DISABLE_WINDOWDECORATION=1
export QT_SCALE_FACTOR=1
export QT_AUTO_SCREEN_SCALE_FACTOR=1
export GDK_SCALE=1
export GDK_BACKEND="wayland,x11,*"
export GDK_USE_PORTAL=1
export WLR_NO_HARDWARE_CURSORS=1

# Aliases

alias grep='grep --color=auto'
alias v='nvim'

alias cat='bat'
alias vid='mpv'
alias top='btop'
alias cp='cp -r'
alias img='swayimg'
alias rm='rm -rf --'
alias mkdir='mkdir -p'
alias reboot='sudo reboot'
alias shutdown='sudo shutdown -h now'

alias ls='lsd --color=always --icon=always'
alias sl='lsd --color=always --icon=always'
alias tree='lsd --almost-all --icon=always --sort=extension --tree'
alias ll='lsd --long --almost-all --git --color=always --icon=always'

# Git
alias ga='git add .'
alias gs='git status'
alias gcm='git commit -m $1'
alias gb='git branch -M master'
alias gp='git push -u origin master'


# Functions
# Usage: '.. 3' takes you 3 folders up
..() { cd "$(eval printf '../%.0s' {1..$1})" || return 1; }


