/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   server.go                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/25 20:17:01 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/12/11 23:44:35 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var port = os.Getenv("PORT")

func Start() {
	fmt.Println("=== HTTP Wiki server ===")

	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/pages/", makeHandler(showPageHandler))
	http.HandleFunc("/edit/", makeHandler(editPageHandler))
	http.HandleFunc("/save/", makeHandler(savePageHandler))

	fmt.Println("Listening on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
