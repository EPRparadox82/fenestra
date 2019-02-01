# Fenestra

A simple, beautiful and customizable Viewer for my Cheatsheets.
Its written in go and uses nucular.

## Install

A Makefile is included to download all the important libaries and compile with the propper flags.
Just run 

```
make fetch_all
make build
```



## Usage

You can load from file or pipe text into Fenestra as you would with xmessage. Also a self destruct timer can be set to end the cheat peek automatically. The theme features of Nucular can be triggered also from the commandline.

```
cat example.cheat | fenestra 
```
will open the cheatsheet
```
fenestra -f example.cheat -d 10 --color-theme Red
```
will do the same in Red Theme for 10s and close

## Dependencies

[Nucular](https://github.com/aarzilli/nucular) - The awsome Go Port of Nuklear by Alessandro Arzilli

[GetOpt](https://github.com/pborman/getopt) - An versatile Flagparser for CLIs by Paul Borman


## Motivation 

I was configuring a new installation of Xmonad and went fairly creative on Keybindings. As the Xmonad Help only really is one long string piped into an xmessage, I decided to document all my precious keys in a cheatsheet. I often write some Cheatsheets im my home folder to keep clever one-liners or specific program sequences. My usual method of opening them with a texteditor seemed a little uncouth because I will not edit my Xmonad cheatsheet in some time and xmassage seemed a bit too retro for my sweet new setup.
I love the simplicity and the looks of the Nucular GUI libary and decided to use it for a new simple viewer that could display my sheets nicely and give them some order. This is the result. 
