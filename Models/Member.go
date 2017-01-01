/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 01/01/2017    
* Time      : 21:14
* Developer : ibrahimcobani
*
*******/
package Models

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
)

type Member struct {
	email     string
	id        int
	password  string
	firstname string
}

func (this *Member) Email() string {
	return this.email
}

func (this *Member) Id() int {
	return this.id
}

func (this *Member) Password() string {
	return this.password
}

func (this *Member) FirstName() string {
	return this.firstname
}

func (this *Member) SetEmail(value string) {
	this.email = value
}

func (this *Member) SetId(value int) {
	this.id = value
}

func (this *Member) SetPassword(value string) {
	this.password = value
}

func (this *Member) SetFirstName(value string) {
	this.firstname = value
}

func (this *Member)GetMember(email string, password string) (Member, error) {


	db, err := getDBConnection()
	if err == nil {
		defer db.Close()
		pwd := sha256.Sum256([]byte(password))
		fmt.Println(hex.EncodeToString(pwd[:]))

		row := db.QueryRow(`
			select
				id,
				email,
				first_name
			from
				member
			where
				email = $1
				and password = $2`,
			email,
			hex.EncodeToString(pwd[:]))

		result := Member{}
		err := row.Scan(&result.id, &result.email, &result.firstname)
		if err == nil {
			return result, nil
		} else {
			return result, errors.New("Kullanıcı adı bulunamadı. " + email)
		}

	} else {
		return Member{}, errors.New("Database bağlantısı kurulumadı. " + email)
	}

}