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

type Home struct {
	Title       string
	Active      string
	ProfileMenu []ProfileDropdownMenu
	Profile     Profile
}

func GetHome(Id int) Home {

	result := Home{
		Title : "aExpence Masraf Yönetimi",
		Active : "Home",
		ProfileMenu : GetProfileDropdownMenu(),
		Profile:GetProfile(),
	}

	switch Id {
	case 3:
		result.Title = "Ürünler"
	case 1:
		result.Title = "Profil Bilgileri"
	}

	return result
}

type Login struct {
	Title  string
	Active string
	ProfileMenu []ProfileDropdownMenu
	Profile     Profile
}

func GetLogin() Login {
	result := Login{
		Title:"aExpence Masraf Yönetimi",
		Active:"",
		ProfileMenu : GetProfileDropdownMenu(),
		Profile:GetProfile(),
	}
	return result
}