/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   server.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/25 20:17:01 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/25 23:13:03 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"log"
	"net/http"
)

const url = "localhost:8080"

func pagesHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/pages/"):]
	page, _ := loadPage(title)

	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", page.Title, page.Body)
}

func sartServer() {
	fmt.Println("=== HTTP Wiki server ===")

	http.HandleFunc("/", pagesHandler)

	fmt.Println("Listenin on http://" + url)
	log.Fatal(http.ListenAndServe(url, nil))
}
