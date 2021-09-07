package main

import ("ggiz.statistic.consumer.agent/cartaz" 
        "ggiz.statistic.consumer.agent/consuos"
		"ggiz.statistic.consumer.agent/server"
)  

func main() {
	var rsys *consuos.Status = &consuos.Status{}
	var cartaz cartaz.Cartaz = cartaz.Cartaz{
		Refsys: rsys,
	}
	var rServer *server.GgizConsumer = &server.GgizConsumer{
		RefSys: rsys,
   		RefCartaz: cartaz,
	}

	rsys.Init()
	cartaz.ShowCartazIni()
	cartaz.ShowDetailIni()
	cartaz.ReportIni()
	rServer.InitServer()
}