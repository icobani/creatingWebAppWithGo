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
	"time"
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

type Session struct {
	id        int
	memberId  int
	sessionId string
}

func (this *Session) Id() int {
	return this.id
}

func (this *Session) MemberId() int {
	return this.memberId
}

func (this *Session) SessionId() string {
	return this.sessionId
}

func (this *Session) SetId(value int) {
	this.id = value
}

func (this *Session) SetMemberId(value int) {
	this.memberId = value
}

func (this *Session) SetSessionId(value string) {
	this.sessionId = value
}

func GetMember(email string, password string) (Member, error) {

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
			fmt.Println(result)
			return result, nil
		} else {
			fmt.Println("Kullanıcı bulunamadı", email, hex.EncodeToString(pwd[:]), err.Error())
			return result, errors.New("Kullanıcı adı bulunamadı. " + email)
		}

	} else {
		return Member{}, errors.New("Database bağlantısı kurulumadı. " + email)
	}

}

func CreateSession(member Member) (Session, error) {
	result := Session{}
	result.memberId = member.Id()
	sessionId := sha256.Sum256([]byte(member.Email() + time.Now().Format("12:00:00")))
	result.sessionId = hex.EncodeToString(sessionId[:])

	db, err := getDBConnection()
	if err == nil {
		defer db.Close()

		err := db.QueryRow(`
			INSERT INTO public."Session"(
			session_id, member_id)
			values
				($1, $2)
			returning id`, result.sessionId, member.Id()).Scan(&result.id)

		if err == nil {
			return result, nil
		} else {
			fmt.Println(err.Error())
			return Session{}, errors.New("Session durumu database'e yazılamadı.")
		}
	} else {
		return Session{}, errors.New("Veritabanı bağlantısı kurulamadı.")
	}
}