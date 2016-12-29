/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 28/12/2016    
* Time      : 14:20
* Developer : ibrahimcobani
*
*******/
package ViewModels

import "fmt"

type Home struct {
	Title       string
	Active      string
	ProfileMenu []ProfileDropdownMenu
}

func GetHome() Home {
	result := Home{
		Title : "Performanse HR Programı",
		Active : "Home",
		ProfileMenu : GetProfileDropdownMenu(),
	}
	fmt.Println(result)
	return result
}
