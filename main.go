package main

import (
	//"fmt"
	"os"
	"bufio"
	"strings"

	"github.com/aarzilli/nucular"
	_"github.com/aarzilli/nucular/label"
	nstyle "github.com/aarzilli/nucular/style"
)

var (
	scaling = 1.1
	Wnd nucular.MasterWindow
	theme nstyle.Theme = nstyle.DarkTheme
	dat []data
	hea string
)

type data struct {
	title string
	textl []string
	textr []string
}


func main() {

	hea,dat = loadfile("data.txt")

	nw := newNucularWindow()
	nw.Theme = theme

	Wnd = nucular.NewMasterWindow(0,hea, nw.nucularWindow)
	Wnd.SetStyle(nstyle.FromTheme(theme, scaling))
	Wnd.Main()

}
func loadfile(filename string) (head string,out []data){
	f, _ := os.Open(filename)
	scanner := bufio.NewScanner(f)
	title := ""
	keys := []string{}
	info := []string{}
	for scanner.Scan() {
		li := scanner.Text()
		line := strings.Replace(li,"\t","",-1)

		if strings.HasPrefix(line,"//") {
			continue
		}
		if strings.HasPrefix(line,"***") {
			head = line[3:]
			continue
		}
		if strings.HasPrefix(line,"##") {
			if title != "" {
				out = append(out,data{title,keys,info})
			}
			title = line[2:]
			keys = nil
			info = nil
		}else{
			if line != "" {
				da := strings.Split(line,"--")
				keys = append(keys,da[0])
				info = append(info,da[1])
			}
		}
	}

	out = append(out,data{title,keys,info})
	return
}

