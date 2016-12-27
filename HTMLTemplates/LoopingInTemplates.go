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

func LoopingInTemplatesMain() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		tmpl, err := template.New("test").Parse(page4)
		if err == nil {
			context := Context2{
				[13]string{
					"Armut",
					"Muz",
					"Şeftali",
					"Elma",
					"Üzüm",
					"Karpuz",
					"Kavun",
					"Kayısı",
					"Kivi",
					"Trabzon Hurması",
					"Erik",
					"Çilek",
					"Yeni Dünya"},
				"Şen MANAV",
				"Şen MANAV Meyve Listemiz",
			}
			tmpl.Execute(w, context)
		}
	})

	http.ListenAndServe(":8000", nil)
}

const page4 = `
<!DOCTYPE html>
<html>
	<head><title>{{.Title}}</title></head>
	<body>

		<h1>{{.ListTitle}}</h1>
		<ul>
			{{range .Fruilts}}
			<li>{{.}}</li>
			{{end}}
		</ul>
	</body>
</html>
`

type Context2 struct {
	Fruilts [13]string
	Title   string
	ListTitle string
}

