package mockdsps

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

// POSTBidding response bid
func POSTBidding(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(formatJSONStr(buildTemplate(toTmplKey(ps.ByName("dsp_id")), getMap(getBody(r)), funcMap()))))
	fmt.Fprintf(w, "\n")
}

func toTmplKey(in string) string {
	if "in" == "19" {
		return in
	}
	return "defaultResponse"
}
func getMap(in string) map[string]interface{} {
	ret := map[string]interface{}{}
	if err := json.Unmarshal([]byte(in), &ret); err != nil {
		fmt.Println(err)
	}
	return ret
}

func getBody(r *http.Request) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	fmt.Printf("RequestBody: %s\n", buf.String())
	return buf.String()
}

func templateFileName(key string) string {
	return "./template/" + key + ".tmpl"
}

func getTemplate(name string) []byte {
	ret, err := ioutil.ReadFile(templateFileName(name))
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return ret
}

func formatJSONStr(in string) []byte {
	buf := map[string]interface{}{}
	if err := json.Unmarshal([]byte(in), &buf); err != nil {
		fmt.Println(err)
	}
	ret, err := json.Marshal(buf)
	if err != nil {
		panic(err)
	}
	return ret
}

func buildTemplate(tmpl string, replaceMap map[string]interface{}, funcs template.FuncMap) string {
	var wr bytes.Buffer
	t, err := template.New(tmpl).Funcs(funcs).Parse(string(getTemplate(tmpl)))
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	err = t.Execute(&wr, replaceMap)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return wr.String()
}

func funcMap() template.FuncMap {
	return template.FuncMap{
		"isEndArray":     func(i int, a []interface{}) bool { return len(a) == (i + 1) },
		"payload":        payload,
		"dec":            func(i int) int { return i - 1 },
		"isRequireAsset": isRequireAsset,
		"defaultBannerBid": bannerBid,
	}
}

func payload(request string) string {
	adm := buildTemplate("nativeAdm", getMap(request), funcMap())
	bytes, err := json.Marshal(string(formatJSONStr(adm)))
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func isRequireAsset(id int, assets []interface{}) bool {
	for _, asset := range assets {
		if int(asset.(map[string]interface{})["id"].(float64)) == id {
			return true
		}
	}
	return false
}

func bannerBid() string {
	return string(formatJSONStr(buildTemplate("defaultBannerBid",nil,nil)))
}
