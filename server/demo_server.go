/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   demo_server.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/25 20:17:01 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/12/11 23:37:32 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package server

import (
	"fmt"
	"log"
	"net/http"
)

const demoUrl = "localhost:8080"

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func Demo() {
	fmt.Println("=== HTTP demo server ===")

	http.HandleFunc("/", homeHandler)

	fmt.Println("Listenin on http://" + demoUrl)
	log.Fatal(http.ListenAndServe(demoUrl, nil))
}
