package mockdsps

import(
	//"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"bytes"
	"text/template"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/julienschmidt/httprouter"
)

// Bidding response bid
func Bidding(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, replaceMacro(string(getResponses(19))))
	fmt.Printf("Todo show %s", ps.ByName("dsp_id"))
}

func getResponses(dspid int) []byte {
	raw, err := ioutil.ReadFile("./conf.json")
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
	json := simplejson.New()
	err = json.UnmarshalJSON(raw)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
	ret , _ := json.Get("responses").Get("19").Get("response").Encode()
	return ret
}

func replaceMacro(replaced string) string {
	var wr bytes.Buffer
	t, err := template.New("response").Parse(replaced)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
	err = t.Execute(&wr, map[string]string{
		"impid": "hogehoge",
	})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return wr.String()
}
