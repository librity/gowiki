/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   handlers.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/26 00:59:07 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/26 01:26:37 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"net/http"
)

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
