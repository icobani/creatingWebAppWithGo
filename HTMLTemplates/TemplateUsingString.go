/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 26/12/2016    
* Time      : 22:31
* Developer : ibrahimcobani
*
*******/
package HTML_Templates

import (
	"net/http"
	"text/template"
)

func TemplateUsingStringMain()  {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		tmpl, err := template.New("test").Parse(page1)
		if err == nil{
			tmpl.Execute(w,req.URL.Path)

		}
	})

	http.ListenAndServe(":8000",nil)
}

const page1 = `
<!DOCTYPE html>
<html>
	<head><title>Example Title</title></head>
	<body>
		<h1>Hello from {{.}}</h1>
	</body>
</html>
`