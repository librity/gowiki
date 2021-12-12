/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   templates.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/26 00:58:15 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/12/11 23:45:13 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package server

import (
	"html/template"
	"net/http"
)

const templatesDir = "templates/"

var templates = template.Must(
	template.ParseFiles(
		tmplPath("pages"),
		tmplPath("show_page"),
		tmplPath("new_page"),
		tmplPath("edit_page"),
	))

func tmplPath(title string) string {
	return templatesDir + title + ".html"
}

func renderTemplate(w http.ResponseWriter, page *Page, name string) {
	err := templates.ExecuteTemplate(w, name+".html", page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
