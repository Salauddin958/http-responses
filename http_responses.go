package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"path"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Profile struct {
	Name    string
	Hobbies []string `xml:"Hobbies>Hobby"`
}

func main() {
	http.HandleFunc("/plain", plainTextResponse)
	http.HandleFunc("/json", jsonResponse)
	http.HandleFunc("/template", templateResponse)
	http.HandleFunc("/image", imageResponse)
	http.HandleFunc("/xml", xmlResponse)
	http.ListenAndServe(":8081", nil)
}

func plainTextResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Plain Text Response")
}
func jsonResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := User{
		ID:    1,
		Name:  "John Doe",
		Email: "johndoe@gmail.com",
		Phone: "000099999",
	}
	json.NewEncoder(w).Encode(user)
}

func templateResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charest=UTF-8")
	tpl, err := template.ParseFiles("sample_template.gohtml")
	if err != nil {
		fmt.Fprintf(w, "Unable to load template")
	}
	user := User{
		ID:    10,
		Name:  "John Doe",
		Email: "johndoe@gmail.com",
		Phone: "000099999",
	}
	tpl.Execute(w, user)
}

func imageResponse(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("images", "sample_image.jpg")
	http.ServeFile(w, r, fp)
}
func xmlResponse(w http.ResponseWriter, r *http.Request) {
	profile := Profile{"Sally", []string{"cricket", "swimming"}}
	x, err := xml.MarshalIndent(profile, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/xml")
	w.Write(x)
}
