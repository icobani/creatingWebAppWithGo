/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 01/01/2017    
* Time      : 18:08
* Developer : ibrahimcobani
*
*******/
package CreatingAResourceServer
import (
	"testing"
)

func Test_CheckOnline(t *testing.T) {
	count := 3
	if count == 4 {
		t.Fail()
	}
}