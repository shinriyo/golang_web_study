package main

import (
//    "fmt"
    "net/http"
    "io/ioutil"
//    "os"
//    "template"
//    "old/template"
    "text/template"
)


type Page struct {
    Title string
    Body []byte
}

//func (p *Page) save() os.Error {
func (p *Page) save() error {
  filename := p.Title + ".txt"
  return ioutil.WriteFile(filename, p.Body, 0600)
}

/*
func loadPage(title string) *Page {
    filename := title + ".txt"
    body, _ := ioutil.ReadFile(filename)
    return &Page{Title: title, Body: body}
}
*/

func loadPage(title string) (*Page, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    return &Page{Title: title, Body:body}, nil
}

const lenPath = len("/view/")

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[lenPath:]
    p, _ := loadPage(title)
    //fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
    t, _ := template.ParseFiles("view.html")
    t.Execute(w, p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[lenPath:]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    /*
    fmt.Fprintf(w, "<h1>Editing</h1>"+
        "<textarea name=\"body\">" +
        "<input type=\"submit\" value=\"Save\">" +
        p.Title, p.Title, p.Body)
    */
    t, _ := template.ParseFiles("view.html")
    t.Execute(w, p)
}

func main() {
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/edit/", editHandler)
    //http.HandleFunc("/save/", saveHandler)
    http.ListenAndServe(":8080", nil)
}
