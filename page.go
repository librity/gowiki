/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   page.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/25 20:10:35 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/26 00:52:12 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
)

const pagesDir = "pages/"

type Page struct {
	Title    string
	Body     []byte
	HTMLBody template.HTML
}

func (p *Page) save() error {
	filePath := getPagePath(p.Title)
	return ioutil.WriteFile(filePath, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filePath := getPagePath(title)
	body, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	htmlBody := template.HTML(body)

	return &Page{Title: title, Body: body, HTMLBody: htmlBody}, nil
}

func getPagePath(title string) string {
	return pagesDir + title + ".html"
}

func pageDemo() {
	fmt.Println("=== Page demo ===")
	p1 := &Page{Title: "test", Body: []byte("This is a simple page")}
	p1.save()

	p2, _ := loadPage("test")
	fmt.Println(string(p2.Body))
}
