package main

import (
	//"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Page struct{
	Title string
	Body []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage (title string) (*Page, error) {
	filename := title + ".txt"
	body,err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func home_page(w http.ResponseWriter, r *http.Request){
	//fmt.Fprintf(w, "My website works! The name of this page is: %s", r.URL.Path[1:])
	title := r.URL.Path
	p, err := loadPage(title)
	if err != nil {
		p = &Page{ Title: title}
	}
	t,err := template.ParseFiles("templates/index.html")
	if err != nil {
		p = &Page{ Title: title}
	}
	t.Execute(w, p)

}

func about_page(w http.ResponseWriter, r *http.Request){
	title := r.URL.Path[1:len(r.URL.Path)-1]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{ Title: title}
	}
	t,err := template.ParseFiles("templates/about.html")
	if err != nil {
		p = &Page{ Title: title}
	}
	t.Execute(w, p)

}

func main() {
	http.HandleFunc("/about/",about_page)
	http.HandleFunc("/",home_page)

	log.Fatal(http.ListenAndServe(":8080",nil))


}




	// p1 := &Page{Title: "TestPage", Body: []byte("Это работает! Вау")}
	// err := p1.save()
	// if err != nil {
	// 	panic(err)
	// }
	// p2,err := loadPage("TestPage")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(p2.Body))
