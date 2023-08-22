package main

import (
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

func index(w http.ResponseWriter, r *http.Request){
	//fmt.Fprintf(w, "My website works! The name of this page is: %s", r.URL.Path[1:])
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	t,err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = t.ExecuteTemplate(w , "index", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

}

func about_page(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/about/" {
		http.NotFound(w, r)
		return
	}
	t,err := template.ParseFiles("templates/about.html", "templates/header.html", "templates/footer.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = t.ExecuteTemplate(w , "about", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/",index)
	http.HandleFunc("/about/",about_page)

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
