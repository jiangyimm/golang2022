package main

import (
	"net/url"

	"github.com/webview/webview"
)

func main() {
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Go webview app")
	w.SetSize(600, 400, webview.HintNone)
	w.Navigate("data:text/html," + url.PathEscape(`
  <html>
    <body>
      <h1 style="padding-top: 40vh; text-align: center;">Hello, Go webview!</h1>
    </body>
  </html>
  `))
	w.Run()
}
