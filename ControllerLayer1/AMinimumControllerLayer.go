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
package ControllerLayer1

import (
	"net/http"
	"os"
	"text/template"
	"github.com/icobani/CreatingWebApp/controllers"
)

func AMinimumControllerLayerMain() {
	templates := populateTemplates4()
	Controllers.Register(templates)
	http.ListenAndServe(":8000", nil)
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