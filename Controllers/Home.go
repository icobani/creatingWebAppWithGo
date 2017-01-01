/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 29/12/2016    
* Time      : 23:13
* Developer : ibrahimcobani
*
*******/
package Controllers

import (
	"text/template"
	"net/http"
	"github.com/icobani/CreatingWebApp/ViewModels"
	"bitbucket.org/aexpense/pdf/app/Godeps/_workspace/src/github.com/gorilla/mux"
	"strconv"
	"github.com/icobani/CreatingWebApp/Controllers/Util"
	"github.com/icobani/CreatingWebApp/Models"
)

type homeController struct {
	template      *template.Template
	loginTemplate *template.Template
}

func (this *homeController)login(w http.ResponseWriter, req *http.Request) {
	responseWriter := Util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	responseWriter.Header().Add("Content-Type", "text/html")

	if req.Method == "POST" {
		email := req.FormValue("email")
		password := req.FormValue("password")

		member, err := Models.Getmember(email, password)
	}

	vm := ViewModels.GetLogin()
	this.loginTemplate.Execute(responseWriter, vm)
}

func (this *homeController)get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	idRav := vars["id"]
	id, err := strconv.Atoi(idRav)
	if err == nil {
		vm := ViewModels.GetHome(id)
		w.Header().Add("Content-Type", "text/html")
		responseWriter := Util.GetResponseWriter(w, req)
		defer responseWriter.Close()
		this.template.Execute(responseWriter, vm)
	} else {
		vm := ViewModels.GetHome(0)
		w.Header().Add("Content-Type", "text/html")
		responseWriter := Util.GetResponseWriter(w, req)
		defer responseWriter.Close()
		this.template.Execute(responseWriter, vm)
	}
}