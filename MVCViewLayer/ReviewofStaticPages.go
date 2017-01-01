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
	"fmt"
)

func ReviewofStaticPagesMain() {
	templates := populateTemplates()

	http.HandleFunc("/",
		func(w http.ResponseWriter, req *http.Request) {
			requestedFile := req.URL.Path[1:]
			template :=
				templates.Lookup(requestedFile + ".html")

			if template != nil {
				template.Execute(w, nil)
			} else {
				fmt.Println("404")
				w.WriteHeader(404)
			}
		})

	http.ListenAndServe(":8000", nil)

}

func populateTemplates() *template.Template {
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