/******
* B1 Yönetim Sistemleri Yazılım ve Danışmanlık Limited Şirketi
* B1 Digitial
* http://www.b1.com.tr
*
*
*
* Date      : 26/12/2016    
* Time      : 20:57
* Developer : ibrahimcobani
*
*******/
package CreatingAResourceServer
import (
"net/http"

"strings"
"os"
"bufio"
	"fmt"
)

func index(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Babaley nittin."))
}

func BufferingaOutputStreamMain() {
	//http.HandleFunc("/pisi", index)

	http.Handle("/", new(MyHandler))

	http.ListenAndServe(":8000", nil)
}

type MyHandler struct {
	http.Handler
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := "./public/" + req.URL.Path
	//data, err := ioutil.ReadFile(string(path))
	f, err := os.Open(path)

	if err == nil {
		var contentType string
		// ilgili dosyanın tipine göre content type belirlenip header'a ekleniyor.
		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else {
			contentType = "text/plain"
		}
		w.Header().Add("Content-Type", contentType)

		bufferedReader := bufio.NewReader(f)
		bufferedReader.WriteTo(w)
	} else {
		fmt.Println("404")
		w.WriteHeader(404)
		w.Write([]byte("404 - " + http.StatusText(404)))
	}

}
