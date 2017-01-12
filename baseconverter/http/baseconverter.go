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

var getHandlerTemplate = func() func() *template.Template {
	var handlerTemplate *template.Template
	return func() *template.Template {
		if handlerTemplate == nil {
			handlerTemplate = template.Must(template.ParseFiles("templates/baseconverter.html"))
		}
		return handlerTemplate
	}
}()

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		handlerGET(w, r)
		return
	}
	if r.Method == http.MethodPost {
		handlerPOST(w, r)
		return
	}
}

func handlerGET(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	values := r.URL.Query()
	if number := values.Get("number"); number != "" {
		data["Number"] = number
	}
	if inBase := values.Get("inBase"); inBase != "" {
		data["InBase"] = inBase
	}
	if toBase := values.Get("toBase"); toBase != "" {
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

	handlerTemplate := getHandlerTemplate()
	err := handlerTemplate.Execute(w, data)
	if err != nil {
		http.Error(w, "err := handlerTemplate.Execute(w, data)", http.StatusInternalServerError)
		return
	}
	return
}

func handlerPOST(w http.ResponseWriter, r *http.Request) {
	number := r.PostFormValue("number")
	inBase := r.PostFormValue("inBase")
	toBase := r.PostFormValue("toBase")

	values := url.Values{}

	if number != "" {
		values.Add("number", number)
	}
	if inBase != "" {
		values.Add("inBase", inBase)
	}
	if toBase != "" {
		values.Add("toBase", toBase)
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
