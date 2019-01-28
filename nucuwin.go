package main

import (
	//"fmt"
	"os"


	"github.com/aarzilli/nucular"
	"github.com/aarzilli/nucular/label"
	nstyle "github.com/aarzilli/nucular/style"
	"github.com/aarzilli/nucular/rect"
	//ncommand "github.com/aarzilli/nucular/command"

	//"golang.org/x/mobile/event/key"
	//"golang.org/x/mobile/event/mouse"
)


type fenestraWindow struct {
	ShowMenu    bool
	Titlebar    bool
	Border      bool
	Resize      bool
	Movable     bool
	NoScrollbar bool
	//Minimizable bool
	Close       bool

	HeaderAlign nstyle.HeaderAlign

	// Menu status
	//Mprog int
	//Mslider int
	//Mcheck  bool
	//Prog    int
	//Slider  int
	//Check bool

	Theme nstyle.Theme
}

func newFenestraWindow() (nw *fenestraWindow) {
	nw = &fenestraWindow{}
	nw.ShowMenu = menubar
	nw.Titlebar = true
	nw.Border = border
	nw.Resize = resize
	nw.Movable = move
	nw.NoScrollbar = scroll
	nw.Close = true

	nw.HeaderAlign = nstyle.HeaderRight
	//nw.Mprog = 60
	//nw.Mslider = 8
	//nw.Mcheck = true

	return nw
}

// Master Window
func (nw *fenestraWindow) masterWindow(w *nucular.Window) {
	//keybindings(w)
	mw := w.Master()

	style := mw.Style()
	style.NormalWindow.Header.Align = nw.HeaderAlign
	if nw.ShowMenu {
		nw.nucularMenubar(w)
	}
	w.Row(30).Dynamic(1)
	w.Label(hea,"CC")
	//w.Spacing(2)
	for _,d := range(dat) {
		if d.title == "" {
			for i,_ := range(d.textl){
				if d.textr[i]==""{
					w.RowScaled(25).Dynamic(1)
					w.Label(d.textl[i],"CC")
				}else if d.textl[i]==""{
					w.RowScaled(20).Dynamic(2)
					w.Label("*)", "RC")
					w.Label(d.textr[i], "LC")
				}else{
					w.RowScaled(20).Dynamic(2)
					w.Label(d.textl[i], "RC")
					w.Label(d.textr[i], "LC")
				}
			}
		}else if w.TreePush(nucular.TreeTab, d.title, true) {
			//w.RowScaled(20).Dynamic(2)
			//w.Row(20).Dynamic(3)
			//w.Row(20).Static(100)
			for i,_ := range(d.textl){
			//	w.Label(d.textl[i], "RC")
			//	w.Label(d.textr[i], "LC")
				//w.Label("", "LC")
				if d.textr[i]==""{
					w.RowScaled(25).Dynamic(1)
					w.Label(d.textl[i],"CC")
				}else if d.textl[i]==""{
					w.RowScaled(20).Dynamic(2)
					w.Label("*)", "RC")
					w.Label(d.textr[i], "LC")
				}else{
					w.RowScaled(20).Dynamic(2)
					w.Label(d.textl[i], "RC")
					w.Label(d.textr[i], "LC")
				}
			}

			w.TreePop()
		}
	}
	w.RowScaled(30).Dynamic(3) //.Static(300, 100, 100)
	w.Label("", "LC")
	if w.ButtonText(exitbut) {
		 os.Exit(0)
	}
}


func (nw *fenestraWindow) nucularMenubar(w *nucular.Window) {
	w.MenubarBegin()
	w.Row(25).Static(45, 70, 45, 70, 70)
	if w := w.Menu(label.TA("Menu", "CC"), 120, nil); w != nil {
		w.Row(25).Dynamic(1)
		/*if w.MenuItem(label.TA("Hide", "LC")) {
			w.Label("You Wish","LC")
			//nw.ShowMenu = false
		}*/
		if w.MenuItem(label.TA("About", "LC")) {
			nw.showAppAbout(w.Master())
		}
		if w.MenuItem(label.TA("Quit", "LC")) {
			nw.showQuestion(w.Master()) //,"Do You want to Quit?") {
				//os.Exit(0)
			//}
		}
	}
	if w := w.Menu(label.TA("Theme", "CC"), 180, nil); w != nil {
		w.Row(25).Dynamic(1)
		newtheme := nw.Theme
		if w.OptionText("Default Theme", newtheme == nstyle.DefaultTheme) {
			newtheme = nstyle.DefaultTheme
		}
		if w.OptionText("White Theme", newtheme == nstyle.WhiteTheme) {
			newtheme = nstyle.WhiteTheme
		}
		if w.OptionText("Red Theme", newtheme == nstyle.RedTheme) {
			newtheme = nstyle.RedTheme
		}
		if w.OptionText("Dark Theme", newtheme == nstyle.DarkTheme) {
			newtheme = nstyle.DarkTheme
		}
		if newtheme != nw.Theme {
			nw.Theme = newtheme
			w.Master().SetStyle(nstyle.FromTheme(nw.Theme, w.Master().Style().Scaling))
			w.Close()
		}
	}
	w.MenubarEnd()
}

func (nw *fenestraWindow) errorPopup(w *nucular.Window) {
	w.Row(25).Dynamic(1)
	w.Label("A terrible error has occured", "LC")
	w.Row(25).Dynamic(2)
	if w.Button(label.T("OK"), false) {
		w.Close()
	}
	if w.Button(label.T("Cancel"), false) {
		w.Close()
	}
}

func (nw *fenestraWindow) questionPopup(w *nucular.Window) {
	w.Row(25).Dynamic(1)
	w.Label("Are You Sure?", "LC")
	w.Row(25).Dynamic(2)
	if w.Button(label.T("OK"), false) {
		os.Exit(0)
	}
	if w.Button(label.T("Cancel"), false) {
		w.Close()
	}
}

func (nw *fenestraWindow) aboutPopup(w *nucular.Window) {
	w.Row(20).Dynamic(1)
	w.Label("Fenestra", "LC")
	w.Row(40).Dynamic(1)
	w.LabelWrap("A viewer for my text based Cheatsheets.")
	w.Row(15).Dynamic(1)
	w.Label("Usind Nucular by Alessandro Arzilli", "LC")
	w.Label("based on Nuklear by Micha Mettke", "LC")

	if w.Button(label.T("OK"), false) {
		w.Close()
	}
}

func (nw *fenestraWindow) showAppAbout(mw nucular.MasterWindow) {
	var wf nucular.WindowFlags

	if nw.Border {
		wf |= nucular.WindowBorder
	}
	if nw.Resize {
		wf |= nucular.WindowScalable
	}
	if nw.Movable {
		wf |= nucular.WindowMovable
	}
	if nw.NoScrollbar {
		wf |= nucular.WindowNoScrollbar
	}
	if nw.Close {
		wf |= nucular.WindowClosable
	}
	if nw.Titlebar {
		wf |= nucular.WindowTitle
	}
	mw.PopupOpen("About", wf, rect.Rect{20, 100, 300, 190}, true, nw.aboutPopup)
}

func (nw *fenestraWindow) showQuestion(mw nucular.MasterWindow) {
	var wf nucular.WindowFlags

	if nw.Border {
		wf |= nucular.WindowBorder
	}
	if nw.Resize {
		wf |= nucular.WindowScalable
	}
	if nw.Movable {
		wf |= nucular.WindowMovable
	}
	if nw.NoScrollbar {
		wf |= nucular.WindowNoScrollbar
	}
	if nw.Close {
		wf |= nucular.WindowClosable
	}
	if nw.Titlebar {
		wf |= nucular.WindowTitle
	}
	mw.PopupOpen("QUIT", wf, rect.Rect{20, 100, 300, 190}, true, nw.questionPopup)
}
