/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   server.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/25 20:17:01 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/12/11 23:23:53 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var port = os.Getenv("PORT")

func sartServer() {
	fmt.Println("=== HTTP Wiki server ===")

	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/pages/", makeHandler(pagesHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	fmt.Println("Listening on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
