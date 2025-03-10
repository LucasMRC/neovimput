# Neovimput

...is brought to you by me. I made this simple thing because I kept trying to use vim keybindings in my web browser inputs.
This way I can fire a vim window, write there and then paste it on the inputs.

### Dependencies

This is made to work with MY setup. Which means this won't work unless you have [Alacritty](https://alacritty.org/) installed. The idea is to add other terminals in the future.
It's using the [attoto/clipboard](https://pkg.go.dev/github.com/atotto/clipboard) package under the hood, so in Linux clipboard management requires you to install xclip or xsel.

Of course, if you don't have [Neovim](https://neovim.io/) then I don't know what you are doing here.


### How should I use this?

The idea is simple: running this will open a dumb neovim buffer. Write whatever you want, then `<C-c>` will write the content in your clipboard and close the neovim window.
No command mode, just you and your nvim keybindings. The buffer filetype is markdown for convenience.

The window will have the title `Neovimput`, so that you can target it on your [i3](https://i3wm.org/) (or your window manager of choice) and set width/height/position/other-stuff-you-want.

### Where do we go from here?

Two things I want to tackle next:
  1. Logs. Simple, but something to allow me to follow the trace always.
  2. A small config file for the write-and-quit command and the filetype.

### Other comments

The name is bad. I don't care.
