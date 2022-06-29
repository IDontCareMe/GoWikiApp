//https://golang-blog.blogspot.com/p/go-web-app.html
package main

import (
	"fmt"
  "io/ioutil"
  //"os"
  "net/http"
  "log"
  "html/template"
)

// Pages
type Page struct {
  Title string
  Body []byte
}
func (p *Page) Save() error {
  //TODO to lower case
  filename := "pages/" + p.Title + ".txt"
  return ioutil.WriteFile(filename, p.Body, 0600)
  //return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
  //TODO to lower case
  filename := "pages/" + title + ".txt"
  //body, _ := os.ReadFile(filename)
  body, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }
  return &Page{Title: title, Body: body}, nil
}

// Templates
var templates = template.Must(template.ParseFiles("templates/view.html", "templates/edit.html"))

func main() {
	http.HandleFunc("/view/", viewHandler)
  http.HandleFunc("/edit/", editHandler)
  http.HandleFunc("/save/", saveHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

// This function handles URL with prefix "/view/"
func viewHandler(w http.ResponseWriter, r *http.Request) {
  //TODO to lower case
  title := r.URL.Path[len("/view/"):]
  p, err := loadPage(title)
  if err != nil {
    http.Redirect(w, r, "/edit/" + title, http.StatusFound)
    return
  }
  renderTemplate(w, "view.html", p)
}

// This function allows to edit Pages
func editHandler( w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/edit/"):]
  p, err := loadPage(title)
  if err != nil {
    p = &Page{Title: title}
  }
  renderTemplate(w, "edit.html", p)
}

// This function saves the pages
func saveHandler( w http.ResponseWriter, r *http.Request) {
  title := r.URL.Path[len("/save/"):]
  body := r.FormValue("body")
  p := &Page{ Title: title, Body: []byte(body) }
  err := p.Save()
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  http.Redirect(w, r, "/view/" + title, http.StatusFound)
}

// This function read and parse template
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
  /*t,err := template.ParseFiles(tmpl + ".html")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  t.Execute(w, p)*/
  err := templates.ExecuteTemplate(w, tmpl, p)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}