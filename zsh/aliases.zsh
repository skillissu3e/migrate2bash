alias v='nvim'
alias t='tmux'

alias cat='bat'
alias vid='mpv'
alias top='btop'
alias cp='cp -r'
alias img='swayimg'
alias rm='rm -rf --'
alias mkdir='mkdir -p'
alias reboot='sudo reboot'
alias shutdown='sudo shutdown -h now'

alias ls='eza --icons --color=always'
alias sl='eza --icons --color=always'
alias tree='eza --all --icons --sort=type -T'
alias ll='eza -laghm@ --all --icons --git --color=always'

# Git

alias ga='git add .'
alias gs='git status'
alias gcm='git commit -m $1'
alias gb='git branch -M master'
alias gp='git push -u origin master'

# Directories

alias config='cd $HOME/.config'
alias nvimcfg='cd $HOME/.config/nvim'
alias plugincfg='cd $HOME/.config/nvim/lua/plugins'

# Open files

alias zshrc='nvim $HOME/.config/zsh/.zshrc'
alias swaycfg='nvim $HOME/.config/sway/config'
alias aliases='nvim $HOME/.config/zsh/aliases.zsh'

# Functions
