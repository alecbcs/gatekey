package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/alecbcs/gatekey/config"
	"github.com/alecbcs/gatekey/web"
)

func main() {
	http.HandleFunc("/create/", web.Create)
	http.HandleFunc("/report/", web.Report)
	http.HandleFunc("/", web.NotFound)
	fmt.Println("Server Started!")
	http.ListenAndServe(":"+strconv.FormatInt(int64(config.Global.General.Port), 10), nil)
}
