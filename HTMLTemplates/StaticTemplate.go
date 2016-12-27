/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 26/12/2016    
* Time      : 22:03
* Developer : ibrahimcobani
*
*******/
package HTML_Templates

import (
	"net/http"
	"text/template"
)

func StaticTemplateMain() {

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		tmpl, err := template.New("test").Parse(doc)
		if err == nil{
			tmpl.Execute(w,nil)

		}
	})

	http.ListenAndServe(":8000",nil)
}

const doc = `
<!DOCTYPE html>
<html>
	<head><title>Example Title</title></head>
	<body>
		<h1>Hello from Template</h1>
	</body>
</html>
`