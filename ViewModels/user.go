/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 30/12/2016    
* Time      : 21:37
* Developer : ibrahimcobani
*
*******/
package ViewModels

type Profile struct {
	Title  string
	Active string
	User   User
}

type User struct {
	Id        int
	Email     string
	FirstName string
	LastName  string
	Stand     Stand
}

type Stand struct {
	Address string
	City    string
	State   string
	Zip     string
}

func GetProfile() Profile {
	result := Profile{
		Title : "Bidircik",
	}
	return result
}

