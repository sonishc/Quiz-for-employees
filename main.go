package main

import (
  "fmt"
  "html/template"
  "log"
  "net/http"
)

var tpl *template.Template

type pageData struct {
  Title     string
  FirstName string
}

func init() {
  tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}
func main() {
  http.HandleFunc("/", idx)
  http.HandleFunc("/about", abot)
  http.HandleFunc("/contact", cntct)
  http.HandleFunc("/apply", aply)
  http.ListenAndServe(":8080", nil)
}

func idx(w http.ResponseWriter, req *http.Request) {
  pd := pageData{
    Title: "Index page",
  }

  err := tpl.ExecuteTemplate(w, "index.gohtml", pd)

  if err != nil {
    log.Println("LOGGED", err)
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    return
  }

  fmt.Println(req.URL.Path)

}
func abot(w http.ResponseWriter, req *http.Request) {
  pd := pageData{
    Title: "About page",
  }

  err := tpl.ExecuteTemplate(w, "about.gohtml", pd)

  if err != nil {
    log.Println(err)
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    return
  }
}
func cntct(w http.ResponseWriter, req *http.Request) {
  pd := pageData{
    Title: "Contact page",
  }
  err := tpl.ExecuteTemplate(w, "contact.gohtml", pd)

  if err != nil {
    log.Println(err)
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    return
  }

}
func aply(w http.ResponseWriter, req *http.Request) {
  pd := pageData{
    Title: "Apply page",
  }

  var first string
  if req.Method == http.MethodPost {
    first = req.FormValue("fname")
    pd.FirstName = first
  }

  err := tpl.ExecuteTemplate(w, "apply.gohtml", pd)

  if err != nil {
    log.Println(err)
    http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    return
  }
}
