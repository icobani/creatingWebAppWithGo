/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 25/12/2016    
* Time      : 23:08
* Developer : ibrahimcobani
*
*******/
package main

import (
	"net/http"
)


func index(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Babaley nittin."))
}

func main() {
	http.HandleFunc("/pisi", index)
	http.ListenAndServe(":8000", nil)
}
