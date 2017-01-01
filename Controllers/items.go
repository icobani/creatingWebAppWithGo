/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 30/12/2016    
* Time      : 16:09
* Developer : ibrahimcobani
*
*******/
package Controllers

import (
	"net/http"
	"text/template"
	"github.com/icobani/CreatingWebApp/ViewModels"
	"github.com/gorilla/mux"
	"strconv"
	"github.com/icobani/CreatingWebApp/Controllers/Util"
)

type ItemsController struct {
	template *template.Template
}

func (this *ItemsController) get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	idRaw := vars["id"]
	id, err := strconv.Atoi(idRaw)
	if err == nil {
		vm := ViewModels.GetHome(id)
		w.Header().Add("Content-Type", "text/html")
		responseWriter := Util.GetResponseWriter(w, req)
		defer responseWriter.Close()
		this.template.Execute(responseWriter, vm)

	} else {
		//w.WriteHeader(404)
		vm := ViewModels.GetHome(0)
		w.Header().Add("Content-Type", "text/html")
		responseWriter := Util.GetResponseWriter(w, req)
		defer responseWriter.Close()
		this.template.Execute(responseWriter, vm)

	}
}