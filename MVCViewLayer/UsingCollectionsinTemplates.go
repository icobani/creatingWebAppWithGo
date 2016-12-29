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
	"github.com/icobani/CreatingWebApp/ViewModels"
	"fmt"
)

func UsingCollectionsinTemplatesMain() {
	templates := populateTemplates4()

	http.HandleFunc("/",
		func(w http.ResponseWriter, req *http.Request) {
			requestedFile := req.URL.Path[1:]
			template :=
				templates.Lookup(requestedFile + ".html")

			var context interface{}
			switch strings.ToUpper(requestedFile) {
			case "HOME":
				context = ViewModels.GetHome()
			}

			if template != nil {
				fmt.Println(context)
				template.Execute(w, context)
			} else {
				w.WriteHeader(404)
			}
		})

	http.HandleFunc("/assest/css/", serveResource3)
	http.HandleFunc("/assest/img/", serveResource3)
	http.HandleFunc("/assest/js/", serveResource3)
	http.HandleFunc("/assest/less/", serveResource3)
	http.HandleFunc("/assest/plugins/", serveResource3)

	http.HandleFunc("/pages/css/", serveResource3)
	http.HandleFunc("/pages/fonts/", serveResource3)
	http.HandleFunc("/pages/ico/", serveResource3)
	http.HandleFunc("/pages/img/", serveResource3)
	http.HandleFunc("/pages/js/", serveResource3)
	http.HandleFunc("/pages/less/", serveResource3)
	http.HandleFunc("/pages/scss/", serveResource3)
	http.HandleFunc("/pages/", serveResource3)

	http.ListenAndServe(":8000", nil)

}

func serveResource3(w http.ResponseWriter, req *http.Request) {
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

func populateTemplates4() *template.Template {
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