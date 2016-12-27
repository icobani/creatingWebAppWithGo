/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 26/12/2016    
* Time      : 22:40
* Developer : ibrahimcobani
*
*******/
package HTML_Templates

import (
	"net/http"
	"text/template"
)
func DynamicTemplateUsingObjectMain()  {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		tmpl, err := template.New("test").Parse(page2)
		if err == nil{
			context := Context{Message:"Ibrahim ÇOBANİ"}
			tmpl.Execute(w,context)
		}
	})

	http.ListenAndServe(":8000",nil)
}

const page2 = `
<!DOCTYPE html>
<html>
	<head><title>Example Title</title></head>
	<body>
		<h1>Hello from {{.Message}}</h1>
	</body>
</html>
`

type Context struct {
	Message string
}