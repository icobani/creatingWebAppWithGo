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
	HRefCode      string
	MenuIconClass string
	MenuTitle     string
}

func GetProfileDropdownMenu() []ProfileDropdownMenu {
	result := []ProfileDropdownMenu{
		{
			HRefCode:"#",
			MenuIconClass:"pg-settings_small",
			MenuTitle:"Sistem Ayarları"},
		{
			HRefCode : "#",
			MenuIconClass:"pg-outdent",
			MenuTitle:"Geribildirim"},
		{
			HRefCode:"#",
			MenuIconClass:"pg-signals",
			MenuTitle:"Yardım"},
		{
			HRefCode:"#",
			MenuIconClass:"pg-save",
			MenuTitle:"Kaydet"},
	}

	return result
}