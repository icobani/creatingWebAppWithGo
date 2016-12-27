/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 26/12/2016    
* Time      : 22:56
* Developer : ibrahimcobani
*
*******/
package HTML_Templates

import (
	"net/http"
	"text/template"
)

func BranchingLogicTemplatesMain() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		tmpl, err := template.New("test").Parse(page3)
		if err == nil {
			tmpl.Execute(w, req.URL.Path)
		}
	})

	http.ListenAndServe(":8000", nil)
}

const page3 = `
<!DOCTYPE html>
<html>
	<head><title>Example Title</title></head>
	<body>
		{{if eq . "/Google"}}
			<h1>Bu Google neler neler yapmış seninle hiç tam konuşma fırsatı bulamadık değilmi.<<h1>
		{{else}}
			<h1>Hello, {{.}}</h1>
		{{end}}
	</body>
</html>
`


