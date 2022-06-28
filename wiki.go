package main

import (
	//"fmt"
  "io/ioutil"
  //"os"
  "net/http"
  "log"
  "html/template"
)

type Page struct {
  Title string
  Body []byte
}
func (p *Page) Save() error {
  //TODO to lower case
  filename := p.Title + ".txt"
  return ioutil.WriteFile(filename, p.Body, 0600)
  //return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
  //TODO to lower case
  filename := title + ".txt"
  //body, _ := os.ReadFile(filename)
  body, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }
  return &Page{Title: title, Body: body}, nil
}

func main() {
	http.HandleFunc("/view/", viewHandler)
  http.HandleFunc("/edit/", editHandler)
  //http.HandleFunc("/save/", saveHandler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}

// This function handles URL with prefix "/view/"
func viewHandler(w http.ResponseWriter, r *http.Request) {
  //TODO to lower case
  title := r.URL.Path[len("/view/"):]
  p, _ := loadPage(title)
  t, _ := template.ParseFiles("templates/view.html")
  t.Execute(w, p)
}

// This function allows to edit Pages
func editHandler( w http.ResponseWriter, r *http.Request){
  title := r.URL.Path[len("/edit/"):]
  p, err := loadPage(title)
  if err != nil {
    p = &Page{Title: title}
  }
  t,_ := template.ParseFiles("templates/edit.html")
  t.Execute(w, p)
}