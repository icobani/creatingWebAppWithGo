/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 28/12/2016    
* Time      : 23:37
* Developer : ibrahimcobani
*
*******/
package ViewModels

type ProfileDropdownMenu struct {
	Id            int32
	HRefCode      string
	MenuIconClass string
	MenuTitle     string
}

func GetProfileDropdownMenu() []ProfileDropdownMenu {
	result := []ProfileDropdownMenu{
		{
			Id:1,
			HRefCode:"/profile",
			MenuIconClass:"pg-settings_small",
			MenuTitle:"Profil Bilgileri"},
		{
			Id:2,
			HRefCode : "/home",
			MenuIconClass:"pg-outdent",
			MenuTitle:"Geribildirim"},
		{
			Id:3,
			HRefCode:"/items",
			MenuIconClass:"pg-signals",
			MenuTitle:"Yardım"},
		{
			Id:4,
			HRefCode:"#",
			MenuIconClass:"pg-save",
			MenuTitle:"Kaydet"},
	}

	return result
}