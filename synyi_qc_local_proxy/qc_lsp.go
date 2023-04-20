package main

//go build -ldflags "-H windowsgui"   打包成后台应用

//http://localhost:3000/inpat?inpatId=ZY230417002 打开质控医生端

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/zserge/lorca"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("this is local proxy serve for synyi quality control system"))
	})
	r.Get("/inpat", Inpat)
	http.ListenAndServe(":3000", r)
}

var (
	ui  lorca.UI
	err error
)

func Inpat(rw http.ResponseWriter, r *http.Request) {

	inpatId := r.URL.Query().Get("inpatId")

	if ui != nil {
		err = ui.Load("http://10.1.72.55:31090/third/control?inpatId=" + inpatId)
		if err != nil {
			go newChrome(inpatId)
		} else {
			//TODO ui置顶显示
		}
	} else {
		go newChrome(inpatId)
	}

}

func newChrome(inpatId string) {
	ui, err = lorca.New("http://10.1.72.55:31090/third/control?inpatId="+inpatId, "", -1, -1, "--remote-allow-origins=*")
	if err != nil {
		log.Fatal(err)
	}

	defer ui.Close()

	ui.Eval(`
	localStorage.setItem('user', '{"orgCode":"0101"}')
	`)

	ui.SetBounds(lorca.Bounds{
		Width:  400,
		Height: 900,
		Top:    50,   //-- set margin top
		Left:   1500, //-- set margin left
	})

	<-ui.Done()
}
