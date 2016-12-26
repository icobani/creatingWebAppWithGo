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
	"io/ioutil"
)

func index(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Babaley nittin."))
}

func main() {
	//http.HandleFunc("/pisi", index)

	http.Handle("/", new(MyHandler))

	http.ListenAndServe(":8000", nil)
}

type MyHandler struct {
	http.Handler
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := "./public/" + req.URL.Path
	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		var contentType string

		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	}

}