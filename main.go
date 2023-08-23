package main

import "github.com/webview/webview"

func main() {
	w := webview.New(false)
	defer w.Destroy()
	w.SetTitle("BananoVault")
	w.SetSize(1900, 900, webview.HintNone)
	w.Navigate("https://vault.banano.cc")
	w.Run()
}