package main

import (
"fmt"
"html/template"
"net/http"
)

func main(){
	http.HandleFunc("/changeuser", createUser)
	http.HandleFunc("/", serveHTML)
	http.ListenAndServe(":8080", nil)
}

func createUser(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	firstname := r.FormValue("firstname")
	surname := r.FormValue("surname")
	email := r.FormValue("email")

	if surname == ""{
		errorMessage := "Error: Surname cannot be empty"
		displayErrorPage(w, errorMessage)
		return
	}
	//inserted a code to create the user using the provided data
	fmt.Printf("User created - Frist Name: %s, Surname: %s, Email: %s\n",firstname, surname, email)

	//success message back to the browser
	successMessage := "User created succesfully."
	displaySuccessPage(w, successMessage)
}

func serveHTML(w, http.ResponseWriter, r *http.Request){
	tpl, err := template.ParseFiles("userChange.html")
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, nil)
}

func displayErrorPage(w http.ResponseWriter, message string){
	tpl, err := template.New("error").Parse("<h1>Success</h1><p>{{.}}</p>")
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, message)
}

func displaySuccessPage(w http.ResponseWriter, message string){
	tpl, err := template.New("success").Parse("<h1>Success</h1><p>{{.}}</p>")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, message)
}