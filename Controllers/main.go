/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 29/12/2016    
* Time      : 22:15
* Developer : ibrahimcobani
*
*******/
package Controllers

import (
	"net/http"
	"os"
	"text/template"
	"strings"
	"bufio"
	"github.com/gorilla/mux"
	"fmt"
)

func Register(templates *template.Template) {
	http.HandleFunc("/assest/css/", serveResource40)
	http.HandleFunc("/assets/img/", serveResource40)
	http.HandleFunc("/assets/js/", serveResource40)
	http.HandleFunc("/assets/less/", serveResource40)
	http.HandleFunc("/assets/plugins/", serveResource40)

	http.HandleFunc("/pages/css/", serveResource40)
	http.HandleFunc("/pages/fonts/", serveResource40)
	http.HandleFunc("/pages/ico/", serveResource40)
	http.HandleFunc("/pages/img/", serveResource40)
	http.HandleFunc("/pages/js/", serveResource40)
	http.HandleFunc("/pages/less/", serveResource40)
	http.HandleFunc("/pages/scss/", serveResource40)
	http.HandleFunc("/pages/", serveResource40)


	router := mux.NewRouter()

	ProfileController := new(profileController)
	ProfileController.template = templates.Lookup("Profile.html")
	router.HandleFunc("/profile/{id}", ProfileController.handle)


	hc := new(homeController)
	hc.template = templates.Lookup("Home.html")
	hc.loginTemplate = templates.Lookup("login.html")
	router.HandleFunc("/home", hc.get)
	router.HandleFunc("/login", hc.login)

	itemscontroller := new(ItemsController)
	itemscontroller.template = templates.Lookup("Items.html")
	router.HandleFunc("/Items", itemscontroller.get)

	itemController := new(ItemsController)
	itemController.template = templates.Lookup("Items.html")
	router.HandleFunc("/items/{id}", itemController.get)




	http.Handle("/", router)
}

func serveResource40(w http.ResponseWriter, req *http.Request) {
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
		fmt.Println("404")
		w.WriteHeader(404)
	}
}