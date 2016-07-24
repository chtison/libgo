package baseconverter

import (
	"html/template"
	"net/http"
	"net/url"

	"github.com/chtison/libgo/baseconverter"
)

func init() {
	http.HandleFunc("/", handler)
}

var appTemplate = template.Must(template.ParseFiles("baseconverter.html"))

func handler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		data := make(map[string]interface{})

		values := r.URL.Query()
		if number := values.Get("number"); number != "" {
			data["Number"] = number
		}
		if inBase := values.Get("inbase"); inBase != "" {
			data["InBase"] = inBase
		}
		if toBase := values.Get("tobase"); toBase != "" {
			data["ToBase"] = toBase
		}
		if errorInBase := values.Get("einbase"); errorInBase != "" {
			data["ErrorInBase"] = errorInBase
		}
		if errorToBase := values.Get("etobase"); errorToBase != "" {
			data["ErrorToBase"] = errorToBase
		}
		if result := values.Get("result"); result != "" {
			data["Result"] = result
		}

		err := appTemplate.Execute(w, data)
		if err != nil {
			http.Error(w, "err := appTemplate.Execute(w, data)", http.StatusInternalServerError)
			return
		}
		return
	}

	if r.Method == http.MethodPost {

		number := r.PostFormValue("number")
		inBase := r.PostFormValue("inbase")
		toBase := r.PostFormValue("tobase")

		values := url.Values{}

		if number != "" {
			values.Add("number", number)
		}
		if inBase != "" {
			values.Add("inbase", inBase)
		}
		if toBase != "" {
			values.Add("tobase", toBase)
		}

		if number != "" && inBase != "" && toBase != "" {

			result, e1, e2 := baseconverter.BaseToBase(number, inBase, toBase)

			if e1 != nil {
				values.Add("einbase", e1.Error())
			} else if e2 != nil {
				values.Add("etobase", e2.Error())
			} else {
				values.Add("result", result)
			}
		}

		URL, err := url.Parse("/")
		if err != nil {
			http.Error(w, `URL, err := url.Parse("/") failed`, http.StatusInternalServerError)
			return
		}
		URL.RawQuery = values.Encode()

		http.Redirect(w, r, URL.String(), http.StatusSeeOther)
		return
	}
}
