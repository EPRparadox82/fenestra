package main

import (
	"fmt"
	"os"
	//"io"
	"bufio"
	"strings"
	"time"

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
	interArgs,empty []string
	hea,filename,themestr,exitbut string
	help, border, resize, move, scroll, menubar bool
	countd int = 0
	version string
	compdate string
	//timer time.Timer
)

type data struct {
	title string
	textl []string
	textr []string
}

func init() {
	filename = "example.cheat"
	exitbut = "QUIT"
	getopt.FlagLong(&border, "no-border", 'b', "Remove Borders")
	getopt.FlagLong(&resize, "no-resize", 'r', "Prohibit resizing")
	getopt.FlagLong(&move, "no-translate", 't', "Prohibit window moving")
	getopt.FlagLong(&scroll, "no-scroll", 's', "Prohibit scrollibars")
	getopt.FlagLong(&menubar, "no-menu", 'm', "Dont Show Menu")
	getopt.FlagLong(&scaling, "magnify", 'g', "Magnification level")
	getopt.FlagLong(&countd, "count-down", 'd', "Set Count down in seconds for auto-quit")
	getopt.FlagLong(&themestr, "color-theme", 'c', "Specify Theme")
	getopt.FlagLong(&help, "help", 'h', "Show Program Usage")
	getopt.FlagLong(&filename, "file", 'f', "The textfile to be parsed and displayed")
	getopt.FlagLong(&exitbut, "quit-button", 'q', "Label of the Quit Button")

}

func main() {

	//timer := time.NewTimer(time.Second*time.Duration(countd*1000))
	//defer timer.Stop()
	go func() {
		//<-timer.C
		time.Sleep(time.Duration(countd)*time.Second)
		if countd>0 {
			fmt.Println("Timer of",countd,"s has Ended")
			os.Exit(0)
		}
	}()
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	getopt.Parse()
	interArgs = getopt.Args()
	additions := false

	border = !border
	resize = !resize
	move = !move
	scroll = !scroll
	menubar = !menubar
	//fmt.Println(info)

	if len(interArgs) > 0 {
		//fmt.Println("Unknown Arguments:",interArgs)
		additions = true
	} else {
		/*fmt.Println("No Args")
		fmt.Println("Border",border)
		fmt.Println("Resize",resize)
		fmt.Println("Translate",move)
		fmt.Println("Scroll",scroll)
		fmt.Println("Menubar",menubar)
		fmt.Println("magn",scaling)*/
	}
	if info.Mode()&os.ModeCharDevice == os.ModeCharDevice {
		//fmt.Println("No Piped Content detected.")
		if additions && filename=="example.cheat"{
			empty = append(empty,"")
			var one []string
			txt := strings.TrimPrefix(fmt.Sprintln(interArgs),"[")
			one = append(one,txt[:len(txt)-2])
			dat = append(dat,data{"",one,empty})
		}else{
			hea,dat = parseFile(loadFile(filename))
		}
	}else{
		hea,dat = parseFile(bufio.NewScanner(os.Stdin))
	}
	if help {
		showHelp()
	}
	//hea,dat = loadfile(filename)

	nw := newFenestraWindow()
	nw.Theme = theme

	Wnd = nucular.NewMasterWindow(0,hea, nw.masterWindow)
	Wnd.SetStyle(nstyle.FromTheme(theme, scaling))
	Wnd.Main()

}
func showHelp() {
	fmt.Printf("Usage of %s   Version %s:\n  A simple and Customizable Cheatsheet viewer.\n  When No Text is piped and no File specified the Parameters\n  will be shown as Text.\n\n", os.Args[0], version)

	getopt.Usage()
	os.Exit(0)
}
func loadFile(filename string) (out *bufio.Scanner) {
	f, err := os.Open(filename)
	if err != nil {
		if filename=="example.cheat" {
			fmt.Println("Example File cannot be Found")
			showHelp()
		}else{
		//fmt.Println("Filename",filename,"was not found!")
			fmt.Println(err)
			os.Exit(1)
		}
	}
	out = bufio.NewScanner(f)
	return
}
func parseFile(in *bufio.Scanner) (head string,out []data){
	scanner := in
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

