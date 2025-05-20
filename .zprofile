#!/bin/sh

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

# Bootstrap .zshrc to $HOME/.config/zsh/.zshrc,
# any other config files can also reside here
export ZDOTDIR="$XDG_CONFIG_HOME/zsh"

# History files
export LESSHISTFILE="$XDG_CACHE_HOME/less_history"
export PYTHON_HISTORY="$XDG_DATA_HOME/python/history"

# Moving other files and some other vars
export CARGO_HOME="$XDG_DATA_HOME/cargo"
export RUSTUP_HOME="$XDG_DATA_HOME/rustup"
export GOPATH="$XDG_DATA_HOME/go"
export GOBIN="$GOPATH/bin"
export PATH=$PATH:$(go env GOPATH)/bin
export GOMODCACHE="$XDG_CACHE_HOME/go/mod"
export NPM_CONFIG_PREFIX="$XDG_DATA_HOME/npm"
export NPM_CONFIG_CACHE="$XDG_CACHE_HOME/npm"
export DATE=$(date "+%-d %B %Y, %A, %H:%M:%S %Z")
export MANPAGER="nvim +Man!"

export FZF_DEFAULT_OPTS="--style full --color 16 --layout=reverse --height 30% --preview='bat -p --color=always {}'"
export FZF_CTRL_R_OPTS="--style full --color 16 --layout=reverse --height 30% --no-preview"

# Important
export ELECTRON_OZONE_PLATFORM_HINT=wayland
export QT_QPA_PLATFORM=wayland
export QT_WAYLAND_DISABLE_WINDOWDECORATION="1"
