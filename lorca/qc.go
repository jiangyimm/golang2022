package main

import (
	"log"
	"os"

	"github.com/zserge/lorca"
)

//go build -ldflags "-H windowsgui"

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
}
