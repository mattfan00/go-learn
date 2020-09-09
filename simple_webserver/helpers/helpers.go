package helpers

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

// parse all of the files into templates at the beginning
// template.Must used to catch any errors
var templates = template.Must(template.ParseFiles("./tmpl/edit.html", "./tmpl/view.html"))

// make sure that path is valid
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func (page *Page) save() error {
	filename := page.Title + ".txt"
	return ioutil.WriteFile(filename, page.Body, 0600)
}

func load(title string) (*Page, error) {
	filename := "./data/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	return &Page{title, body}, err
}

func ViewHandler(w http.ResponseWriter, r *http.Request, title string) {
	loadedPage, err := load(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
	}
	renderTemplate(w, "view", loadedPage)
}

func EditHandler(w http.ResponseWriter, r *http.Request, title string) {
	loadedPage, err := load(title)
	if err != nil {
		loadedPage = &Page{title, nil}
	}
	renderTemplate(w, "edit", loadedPage)
}

func SaveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := []byte(r.FormValue("body"))

	newPage := &Page{title, body}
	if err := newPage.save(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		match := validPath.FindStringSubmatch(r.URL.Path)
		if match == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, match[2])
	}

}

func renderTemplate(w http.ResponseWriter, templateName string, loadedPage *Page) {
	err := templates.ExecuteTemplate(w, templateName+".html", loadedPage)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
