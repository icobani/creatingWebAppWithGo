/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 30/12/2016    
* Time      : 21:44
* Developer : ibrahimcobani
*
*******/
package Controllers

import (
	"github.com/icobani/CreatingWebApp/Controllers/Util"
	"net/http"
	"text/template"
	"github.com/icobani/CreatingWebApp/ViewModels"
	"fmt"
)

type profileController struct {
	template *template.Template
}

func (this *profileController)handle(w http.ResponseWriter, req *http.Request) {
	responseWriter := Util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	vm := ViewModels.GetHome(1)
	vm.Profile.User.Id = 4
	fmt.Println(vm)
	if (req.Method == "POST") {
		vm.Profile.User.FirstName = req.FormValue("FirstName")
		vm.Profile.User.LastName = req.FormValue("LastName")
		vm.Profile.User.Email = req.FormValue("Email")
		vm.Profile.User.Stand.Address = req.FormValue("Address")
		vm.Profile.User.Stand.City = req.FormValue("City")
		vm.Profile.User.Stand.State = req.FormValue("State")
		vm.Profile.User.Stand.Zip = req.FormValue("Zip")
	}
	fmt.Println("****************")
	fmt.Println(vm)
	responseWriter.Header().Add("Content-Type", "text/html")
	this.template.Execute(responseWriter, vm)
}