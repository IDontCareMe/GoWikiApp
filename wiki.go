package main

import (
	"fmt"
  "io/ioutil"
  //"os"
)

type Page struct {
  Title string
  Body []byte
}
func (p *Page) Save() error {
  filename := p.Title + ".txt"
  return ioutil.WriteFile(filename, p.Body, 0600)
  //return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) *Page {
  filename := title + ".txt"
  //TODO errors
  //body, _ := os.ReadFile(filename)
  body, _ := ioutil.ReadFile(filename)
  return &Page{Title: title, Body: body}
}

func main() {
	fmt.Println("Hello, World!")
}