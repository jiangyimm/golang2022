package main

import (
	"log"
	"os"

	"github.com/zserge/lorca"
)

//go build -ldflags "-H windowsgui"
//start E:\gitlab-synyi\golang2022\lorca\lorca.exe ZY230417001
//taskkill -t -f /im lorca.exe
//taskkill -t -f /im lorca.exe & start E:\gitlab-synyi\golang2022\lorca\lorca.exe ZY230417001
//taskkill -t -f /im lorca.exe & start E:\gitlab-synyi\golang2022\lorca\lorca.exe ZY230215002

func main() {

	args := os.Args

	inpatId := args[1]

	ui, err := lorca.New("http://10.1.72.55:31090/third/control?inpatId="+inpatId, "", -1, -1, "--remote-allow-origins=*")
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

	// ui.Eval(`
	// window.addEventListener('message', (event) => {
	//     console.log(event);
	//     if(event.data){
	//     }
	//   })
	// `)
	// ui.Eval(`
	// window.dispatchEvent(new CustomEvent("message",{data:{fun:"openNewPatient",model:{user_info:{org_code:"0101",user_id:"99999",user_name:"哈哈哈hhh",user_role:"sj"},visit_info:{source_visit_id:"57b0451319064db287796a6acc1a3e87"}}}}))
	// `)

	// Wait for the browser window to be closed
	<-ui.Done()

	// //合建chan
	// c := make(chan os.Signal, 1)
	// //监听所有信号
	// signal.Notify(c)
	// //阻塞直到有信号传入
	// fmt.Println("启动")
	// s := <-c
	// ui.Close()
	// fmt.Println("退出信号", s)

	//os.Exit(run())

}
