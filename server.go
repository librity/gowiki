/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   server.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/25 20:17:01 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/26 00:57:17 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const url = "localhost:8080"
const templatesDir = "templates/"

func getTemplatePath(title string) string {
	return templatesDir + title + ".html"
}

func renderTemplate(w http.ResponseWriter, page *Page, name string) {
	templatePath := getTemplatePath(name)
	templ, err := template.ParseFiles(templatePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = templ.Execute(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func pagesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling pages request", r.URL.Path)

	title := r.URL.Path[len("/pages/"):]
	page, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplate(w, page, "page")
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling edit request", r.URL.Path)

	title := r.URL.Path[len("/edit/"):]
	page, err := loadPage(title)
	if err != nil {
		page = &Page{Title: title}
	}

	renderTemplate(w, page, "edit")
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling save request", r.URL.Path)

	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	page := &Page{Title: title, Body: []byte(body)}

	err := page.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/pages/"+title, http.StatusFound)
}

func sartServer() {
	fmt.Println("=== HTTP Wiki server ===")

	http.HandleFunc("/pages/", pagesHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

	fmt.Println("Listenin on http://" + url)
	log.Fatal(http.ListenAndServe(url, nil))
}
