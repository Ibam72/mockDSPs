package mockdsps

import(
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"os"
	"bytes"
	"text/template"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/julienschmidt/httprouter"
)

// POSTBidding response bid
func POSTBidding(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	dspID := ps.ByName("dsp_id")
	body := getBody(r)
	reqMap := map[string]interface{}{}
	if err := json.Unmarshal([]byte(body), &reqMap); err != nil {
		fmt.Println(err)
    }
	fmt.Fprintf(w, replaceMacro(dspID, reqMap))
	fmt.Fprintf(w, "\n")
	fmt.Printf("RequestBody: %s\n", body)
	fmt.Fprintf(w, getBody(r))
}

func getJSON(strJSON []byte) *simplejson.Json{
	ret := simplejson.New()
	err := ret.UnmarshalJSON(strJSON)
	    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
	return ret
}

func getConfJSON() *simplejson.Json {
	raw, err := ioutil.ReadFile("./conf.json")
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
	return getJSON(raw)
}

func getBody(r *http.Request) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	return buf.String()
}

func templateFileName(dspid string) string {
	return "./responses/" + dspid + ".template"
}
func getResponseTemplate(dspid string) []byte {
	ret, err := ioutil.ReadFile(templateFileName(dspid))
    if err != nil {
        fmt.Println(err.Error())
		return nil
    }
	return ret
}

func replaceMacro(dspID string,reqMap map[string]interface{}) string {
	var wr bytes.Buffer
	t, err := template.New("response").Funcs(funcMap()).Parse(string(getResponseTemplate(dspID)))
    if err != nil {
        fmt.Println(err.Error())
        return ""
    }
	fmt.Println(funcMap())
	err = t.Execute(&wr, reqMap)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return wr.String()
}

func funcMap() template.FuncMap{
	return template.FuncMap{
		"isEndArray": isEndArray,
		"payload": payload,
		"dec": func(i int) int {return i - 1;},
	}
}

func payload(request string) string {
	return request	
}

func isEndArray(index int,array []interface{}) bool {
	return len(array) == (index + 1)
}
