/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 31/12/2016    
* Time      : 16:46
* Developer : ibrahimcobani
*
*******/
package Models

import "time"

type Company_Profile struct {
	code                string
	name                string
	address             string
	taxNumber           string
	taxAreaCode         string
	zipCode             string
	city                string
	country             string
	fiscalYearStartDate time.Time
	dateFormat          string
	decimalSeperator    string
	contactName         string
	contactEMail        string
	contactPhone        string
	contactPosition     string
	logo                []byte
}

func (this *Company_Profile) Code() string {
	return this.code
}
func (this *Company_Profile) Name() string {
	return this.name
}
func (this *Company_Profile) Address() string {
	return this.address
}
func (this *Company_Profile) TaxNumber() string {
	return this.taxNumber
}
func (this *Company_Profile) TaxAreaCode() string {
	return this.taxAreaCode
}
func (this *Company_Profile) ZipCode() string {
	return this.zipCode
}
func (this *Company_Profile) City() string {
	return this.city
}
func (this *Company_Profile) Country() string {
	return this.country
}
func (this *Company_Profile) FiscalYearStartDate() time.Time {
	return this.fiscalYearStartDate
}
func (this *Company_Profile) DateFormat() string {
	return this.dateFormat
}
func (this *Company_Profile) DecimalSeperator() string {
	return this.decimalSeperator
}
func (this *Company_Profile) ContactName() string {
	return this.contactName
}
func (this *Company_Profile) ContactEMail() string {
	return this.contactEMail
}
func (this *Company_Profile) ContactPhone() string {
	return this.contactPhone
}
func (this *Company_Profile) ContactPosition() string {
	return this.contactPosition
}
func (this *Company_Profile) Logo() []byte {
	return this.logo
}

func (this *Company_Profile) SetCode(value string) {
	this.code = value
}
func (this *Company_Profile) SetName(value string) {
	this.name = value
}
func (this *Company_Profile) SetAddress(value string) {
	this.address = value
}
func (this *Company_Profile) SetTaxNumber(value string) {
	this.taxNumber = value
}
func (this *Company_Profile) SetTaxAreaCode(value string) {
	this.taxAreaCode  = value
}
func (this *Company_Profile) SetZipCode(value string) {
	this.zipCode = value
}
func (this *Company_Profile) SetCity(value string) {
	this.city = value
}
func (this *Company_Profile) SetCountry(value string) {
	this.country = value
}
func (this *Company_Profile) SetFiscalYearStartDate(value time.Time) {
	this.fiscalYearStartDate = value
}
func (this *Company_Profile) SetDateFormat(value string) {
	this.dateFormat = value
}
func (this *Company_Profile) SetDecimalSeperator(value string) {
	this.decimalSeperator = value
}
func (this *Company_Profile) SetContactName(value string) {
	this.contactName = value
}
func (this *Company_Profile) SetContactEMail(value string) {
	this.contactEMail = value
}
func (this *Company_Profile) SetContactPhone(value string) {
	this.contactPhone = value
}
func (this *Company_Profile) SetContactPosition(value string) {
	this.contactPosition = value
}
func (this *Company_Profile) SetLogo(value []byte) {
	this.logo  = value
}
