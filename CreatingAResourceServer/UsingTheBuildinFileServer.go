/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 26/12/2016    
* Time      : 21:03
* Developer : ibrahimcobani
*
*******/
package CreatingAResourceServer

import "net/http"

func UsingTheBuildinFileServerMain()  {
	http.ListenAndServe(":8000",http.FileServer(http.Dir("public")))
}
