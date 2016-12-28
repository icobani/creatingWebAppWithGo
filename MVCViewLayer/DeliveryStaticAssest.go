/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 27/12/2016    
* Time      : 23:04
* Developer : ibrahimcobani
*
*******/
package MVCViewLayer

import (
	"net/http"
	"os"
	"text/template"
	"strings"
	"bufio"
)

func DeliveryStaticAssestMain() {
	templates := populateTemplates2()

	http.HandleFunc("/",
		func(w http.ResponseWriter, req *http.Request) {
			requestedFile := req.URL.Path[1:]
			template :=
				templates.Lookup(requestedFile + ".html")

			if template != nil {
				template.Execute(w, nil)
			} else {
				w.WriteHeader(404)
			}
		})

	http.HandleFunc("/assest/css/", serveResource)
	http.HandleFunc("/assest/img/", serveResource)
	http.HandleFunc("/assest/js/", serveResource)
	http.HandleFunc("/assest/less/", serveResource)
	http.HandleFunc("/assest/plugins/", serveResource)

	http.HandleFunc("/pages/css/", serveResource)
	http.HandleFunc("/pages/fonts/", serveResource)
	http.HandleFunc("/pages/ico/", serveResource)
	http.HandleFunc("/pages/img/", serveResource)
	http.HandleFunc("/pages/js/", serveResource)
	http.HandleFunc("/pages/less/", serveResource)
	http.HandleFunc("/pages/scss/", serveResource)
	http.HandleFunc("/pages/", serveResource)

	http.ListenAndServe(":8000", nil)

}

func serveResource(w http.ResponseWriter, req *http.Request) {
	//erişilmek istenen adresin yada dosyanın content type'ı belirleniyor.
	path := "public" + req.URL.Path
	var contentType string
	if strings.HasSuffix(path, ".css") {
		contentType = "text/css"
	} else if strings.HasSuffix(path, ".png") {
		contentType = "image/png"
	} else {
		contentType = "text/plain"
	}
	f, err := os.Open(path)
	//eğer bir hata yok ise
	if err == nil {
		defer f.Close()
		w.Header().Add("Content-Type", contentType)
		br := bufio.NewReader(f)
		br.WriteTo(w)
	} else {
		w.WriteHeader(404)
	}
}

func populateTemplates2() *template.Template {
	result := template.New("templates")

	basePath := "templates"
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()

	templatePathRaw, _ := templateFolder.Readdir(-1)

	templatePaths := new([]string)
	for _, pathinfo := range templatePathRaw {
		if !pathinfo.IsDir() {
			*templatePaths = append(*templatePaths,
				basePath + "/" + pathinfo.Name())
		}
	}

	result.ParseFiles(*templatePaths...)
	return result

}