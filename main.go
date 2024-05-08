package main

import (
	"github.com/webview/webview"
)

func main() {
	debug := false
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Free P2P Chat")
	w.SetSize(1024, 768, webview.HintNone)
	w.Navigate("https://chat.51pwn.com:2083/")

	w.Run()
}
