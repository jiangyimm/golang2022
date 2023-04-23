package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/jiangyimm/go-webview2"
	"github.com/joho/godotenv"
)

var (
	web_path string

	ui       webview2.WebView
	is_debug bool

	qc_address         string
	qc_localproxy_port string
	qc_width           int
	qc_height          int
	qc_top             int
	qc_right           int

	emplCode string
	hospCode string
)

func main() {

	var cur_apppath, _ = os.Getwd()
	fmt.Println(cur_apppath)

	web_path = cur_apppath + "\\webview2"

	//sy_log = cur_apppath + "\\resources\\Synyi-Logo.ico"

	if e := godotenv.Load(); e != nil {
		log.Fatal(e)
	}

	is_debug, _ = strconv.ParseBool(os.Getenv("is_debug"))

	qc_address = os.Getenv("qc_address")
	qc_localproxy_port = os.Getenv("qc_localproxy_port")

	qc_width, _ = strconv.Atoi(os.Getenv("qc_width"))
	qc_height, _ = strconv.Atoi(os.Getenv("qc_height"))
	qc_top, _ = strconv.Atoi(os.Getenv("qc_top"))
	qc_right, _ = strconv.Atoi(os.Getenv("qc_right"))

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("this is local proxy serve for synyi quality control system"))
	})
	r.Get("/inpat", Inpat)
	http.ListenAndServe(":"+qc_localproxy_port, r)
}

func Inpat(rw http.ResponseWriter, r *http.Request) {

	inpatId := r.URL.Query().Get("inpatId")
	url_emplCode := r.URL.Query().Get("emplCode")
	url_hospCode := r.URL.Query().Get("hospCode")

	emplCode = url_emplCode
	hospCode = url_hospCode

	go newUI(inpatId, emplCode, hospCode)

	// //医生变化时，新开窗口
	// if url_emplCode != emplCode || url_hospCode != hospCode {
	// 	emplCode = url_emplCode
	// 	hospCode = url_hospCode
	// 	if ui != nil {
	// 		ui.Destroy()
	// 	}
	// 	go newUI(inpatId, emplCode, hospCode)
	// 	return
	// }

	// emplCode = url_emplCode
	// hospCode = url_hospCode

	// if ui != nil {
	// 	ui.Navigate(qc_address + "?inpatId=" + inpatId)
	// 	//ui.
	// 	// if err != nil {
	// 	// 	go newUI(inpatId, emplCode, hospCode)
	// 	// } else {
	// 	// 	//重置bounds 确保 ui置顶显示
	// 	// 	setBounds()
	// 	// }
	// } else {
	// 	go newUI(inpatId, emplCode, hospCode)
	// }
}

func newUI(inpatId string, emplCode string, hospCode string) {

	if ui != nil {
		ui.Destroy()
	}

	ui = webview2.NewWithOptions(webview2.WebViewOptions{
		Debug:     is_debug,
		AutoFocus: true,
		WindowOptions: webview2.WindowOptions{
			WebPath: web_path,
			Title:   "森亿病历内涵质控",
			Width:   uint(qc_width),
			Height:  uint(qc_height),
			IconId:  0, // icon resource id
			Center:  false,
			Right:   uint(qc_right),
			Top:     uint(qc_top),
		},
	})
	if ui == nil {
		log.Fatalln("Failed to load webview.")
	}
	defer ui.Destroy()
	ui.SetSize(qc_width, qc_height, qc_top, qc_right, webview2.HintFixed)

	var user string = `{"orgCode":"` + hospCode + `","emplCode":"` + emplCode + `"}`

	var setUserJS string = `localStorage.setItem('user','` + user + `')`
	ui.Init(setUserJS)

	ui.Navigate(qc_address + "?inpatId=" + inpatId)

	ui.Run()
}
