/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   handlers.go                                        :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/26 00:59:07 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/12/11 23:44:58 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package server

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile("^/(pages|edit|save)/([a-zA-Z0-9]+)$")

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	match := validPath.FindStringSubmatch(r.URL.Path)
	if match == nil {
		http.NotFound(w, r)
		return "", errors.New("invalid Page Title")
	}

	return match[2], nil
}

func makeHandler(callback func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		title, err := getTitle(w, r)
		if err != nil {
			return
		}

		callback(w, r, title)
	}
}

func showPageHandler(w http.ResponseWriter, r *http.Request, title string) {
	fmt.Println("Handling pages request", r.URL.Path)

	page, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}

	renderTemplate(w, page, "show_page")
}

func editPageHandler(w http.ResponseWriter, r *http.Request, title string) {
	fmt.Println("Handling edit request", r.URL.Path)

	page, err := loadPage(title)
	if err != nil {
		page = &Page{Title: title}
	}

	renderTemplate(w, page, "edit_page")
}

func savePageHandler(w http.ResponseWriter, r *http.Request, title string) {
	fmt.Println("Handling save request", r.URL.Path)

	body := r.FormValue("body")
	page := &Page{Title: title, Body: []byte(body)}

	err := page.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/pages/"+title, http.StatusFound)
}
