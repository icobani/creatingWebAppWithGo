/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 01/01/2017    
* Time      : 23:04
* Developer : ibrahimcobani
*
*******/
package Models

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func getDBConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "dbname=aexpense user=postgres password=azura777 sslmode=disable")
	return db, err
}
