package command

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/atotto/clipboard"
)

const (
	FILE_NAME      = "neovimput"
	BUFFER_OPTIONS = "-c 'setlocal cmdheight=0' -c 'setlocal laststatus=0' -c 'setlocal filetype=markdown' -c 'noremap <silent> : <Nop>' -c 'startinsert' -c 'setlocal shortmess+=c'"
	SAVE_CMD       = "-c 'noremap <C-q> <CMD>silent wq<CR>' -c 'inoremap <C-q> <CMD>silent wq<CR>'"
)

func Run() {
	// I'll have this here for the future logs
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	temp_dir := os.TempDir()
	file_path := fmt.Sprintf("%s/%s", temp_dir, FILE_NAME)
	cmd_string := fmt.Sprintf("nvim %s %s %s", file_path, BUFFER_OPTIONS, SAVE_CMD)

	cmd := exec.Command("alacritty", "-T", "Neovimput", "-e", "sh", "-c", cmd_string)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := os.Stat(file_path); err != nil {
		log.Fatalln(err)
	}

	dat, err := os.ReadFile(file_path)
	if err != nil {
		log.Fatalln(err)
	}
	if len(dat) > 0 {
		// remove new line character
		dat = dat[:len(dat)-1]
	}

	clipboard.WriteAll(string(dat))

	if err = os.Remove(file_path); err != nil {
		log.Printf(`
Something failed when clearing the buffer, but fear not! the content was successfully written to the clipboard. You will find this content on the buffer the next time you open the neovimput, though.
The error found was:
%s
`,
			err.Error())
	}
}
