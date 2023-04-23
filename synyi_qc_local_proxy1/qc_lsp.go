package main

//go build -ldflags "-H windowsgui"   打包成后台应用

//http://localhost:3000/inpat?inpatId=ZY230417002&emplCode=99999&hospCode=0101 打开质控医生端

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-vgo/robotgo"
	"github.com/joho/godotenv"
	"github.com/zserge/lorca"
)

var (
	ui  lorca.UI
	err error

	qc_address         string
	qc_localproxy_port string
	qc_width           int
	qc_height          int
	qc_top             int
	qc_left            int

	emplCode string
	hospCode string
)

func main() {

	// 获取屏幕大小
	width, _ := robotgo.GetScreenSize()

	if e := godotenv.Load(); e != nil {
		log.Fatal(e)
	}

	qc_address = os.Getenv("qc_address")
	qc_localproxy_port = os.Getenv("qc_localproxy_port")

	qc_width, _ = strconv.Atoi(os.Getenv("qc_width"))
	qc_height, _ = strconv.Atoi(os.Getenv("qc_height"))
	qc_top, _ = strconv.Atoi(os.Getenv("qc_top"))
	var right, _ = strconv.Atoi(os.Getenv("qc_right"))
	qc_left = width - right - qc_width

	r := chi.NewRouter()
	r.Use(middleware.Logger)
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

	//医生变化时，新开窗口
	if url_emplCode != emplCode || url_hospCode != hospCode {
		emplCode = url_emplCode
		hospCode = url_hospCode
		if ui != nil {
			ui.Close()
		}
		go newChrome(inpatId, emplCode, hospCode)
		return
	}

	emplCode = url_emplCode
	hospCode = url_hospCode

	if ui != nil {
		err = ui.Load(qc_address + "?inpatId=" + inpatId)
		if err != nil {
			go newChrome(inpatId, emplCode, hospCode)
		} else {
			//重置bounds 确保 ui置顶显示
			setBounds()
		}
	} else {
		go newChrome(inpatId, emplCode, hospCode)
	}
}

func newChrome(inpatId string, emplCode string, hospCode string) {
	ui, err = lorca.New("", "", -1, -1, "--remote-allow-origins=*")
	if err != nil {
		log.Fatal(err)
	}

	defer ui.Close()

	ui.Load(qc_address + "?inpatId=" + inpatId)

	var user string = `{"orgCode":"` + hospCode + `","emplCode":"` + emplCode + `"}`

	var setUserJS string = `localStorage.setItem('user','` + user + `')`
	ui.Eval(setUserJS)

	setBounds()

	<-ui.Done()
}

func setBounds() {
	ui.SetBounds(lorca.Bounds{
		Width:  qc_width,
		Height: qc_height,
		Top:    qc_top,  //-- set margin top
		Left:   qc_left, //-- set margin left
	})
}
