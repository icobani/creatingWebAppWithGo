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

func SubTemplateMain() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "text/html")

		templates := template.New("template")
		templates.New("test").Parse(page5)
		templates.New("HeaderSubPage1").Parse(HeaderSubPage1)
		templates.New("FooterSubPage1").Parse(FooterSubPage1)

		context := Context3{
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
				"Kiraz",
				"Çilek",
				"Yeni Dünya"},
			"Şen MANAV",
			"Şen MANAV Meyve Listemiz",
		}
		templates.Lookup("test").Execute(w, context)
	})

	http.ListenAndServe(":8000", nil)
}

const page5 = `
	{{template "HeaderSubPage1" .Title}}
	<body>

		<h1>{{.ListTitle}}</h1>
		<ul>
			{{range .Fruilts}}
			<li>{{.}}</li>
			{{end}}
		</ul>
	</body>
{{template "FooterSubPage1"}}
`

const HeaderSubPage1 = `
<!DOCTYPE html>
<html>
	<head><title>{{.}} biii</title></head>

`

const FooterSubPage1 = `
</html>
`

type Context3 struct {
	Fruilts   [13]string
	Title     string
	ListTitle string
}

