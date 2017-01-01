/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 01/01/2017    
* Time      : 18:15
* Developer : ibrahimcobani
*
*******/
package Util
import (
	"testing"
)

func Test_CheckOnline(t *testing.T) {
	count := 4
	if count == 4 {
		t.Log("Sayıt 4 geldi.")
		t.Fail()
	}
}
