/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   wiki.go                                            :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/25 18:51:19 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/25 19:03:40 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := makeFilename(p.Title)
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := makeFilename(title)
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil
}

func makeFilename(title string) string {
	return "pages/" + title + ".txt"
}

func main() {
	p1 := &Page{Title: "Test page", Body: []byte("This is a simple page")}
	p1.save()

	p2, _ := loadPage("Test page")
	fmt.Println(string(p2.Body))
}
