package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"

	"github.com/pborman/getopt/v2"

	"github.com/aarzilli/nucular"
	_"github.com/aarzilli/nucular/label"
	nstyle "github.com/aarzilli/nucular/style"
)

var (
	scaling = 1.1
	Wnd nucular.MasterWindow
	theme nstyle.Theme = nstyle.DarkTheme
	dat []data
	interArgs []string
	hea,filename,themestr string
	help, border, resize, move, scroll, menubar bool

	version string
	compdate string
)

type data struct {
	title string
	textl []string
	textr []string
}

func init() {
	filename = "none"
	getopt.FlagLong(&border, "no-border", 'b', "Remove Borders")
	getopt.FlagLong(&resize, "no-resize", 'r', "Prohibit resizing")
	getopt.FlagLong(&move, "no-translate", 't', "Prohibit window moving")
	getopt.FlagLong(&scroll, "no-scroll", 's', "Prohibit scrollibars")
	getopt.FlagLong(&menubar, "no-menu", 'm', "Dont Show Menu")
	getopt.FlagLong(&scaling, "magnify", 'g', "Magnification level")
	getopt.FlagLong(&themestr, "color-theme", 'c', "Specify Theme")
	getopt.FlagLong(&help, "help", 'h', "Show Program Usage")
	getopt.FlagLong(&filename, "file", 'f', "The textfile to be parsed and displayed")

}

func main() {
	getopt.Parse()
	interArgs = getopt.Args()

	border = !border
	resize = !resize
	move = !move
	scroll = !scroll
	menubar = !menubar

	if len(interArgs) > 0 {
		fmt.Println("Args:",interArgs)
	} else {
		fmt.Println("No Args")
		fmt.Println("Border",border)
		fmt.Println("Resize",resize)
		fmt.Println("Translate",move)
		fmt.Println("Scroll",scroll)
		fmt.Println("Menubar",menubar)
		fmt.Println("magn",scaling)
	}
	if help {
		fmt.Printf("Usage of %s   Version %s:\n  A simple and Customizable Cheatsheet viewer.\n\n", os.Args[0], version)

		getopt.Usage()
		os.Exit(0)
	}
	hea,dat = loadfile("data.txt")

	nw := newFenestraWindow()
	nw.Theme = theme

	Wnd = nucular.NewMasterWindow(0,hea, nw.masterWindow)
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
			//if title != "" {
				out = append(out,data{title,keys,info})
			//}
			title = line[2:]
			keys = nil
			info = nil
		}else{
			if line != "" {
				if strings.Contains(line, "--"){
					da := strings.Split(line,"--")
					keys = append(keys,da[0])
					info = append(info,da[1])
				}else{
					keys = append(keys,line)
					info = append(info,"")
				}
			}
		}
	}

	out = append(out,data{title,keys,info})
	return
}

