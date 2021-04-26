/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   server.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/25 20:17:01 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/26 00:15:18 by lpaulo-m         ###   ########.fr       */
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

func makeTemplatePath(title string) string {
	return templatesDir + title + ".html"
}

func renderTemplate(w http.ResponseWriter, page *Page, name string) {
	templatePath := makeTemplatePath(name)
	templ, _ := template.ParseFiles(templatePath)
	templ.Execute(w, page)
}

func pagesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling pages request", r.URL.Path)

	title := r.URL.Path[len("/pages/"):]
	page, _ := loadPage(title)

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

func sartServer() {
	fmt.Println("=== HTTP Wiki server ===")

	http.HandleFunc("/pages/", pagesHandler)
	http.HandleFunc("/edit/", editHandler)

	fmt.Println("Listenin on http://" + url)
	log.Fatal(http.ListenAndServe(url, nil))
}
