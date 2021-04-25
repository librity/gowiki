/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   demo_server.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: lpaulo-m <lpaulo-m@student.42sp.org.br>    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2021/04/25 20:17:01 by lpaulo-m          #+#    #+#             */
/*   Updated: 2021/04/25 20:22:59 by lpaulo-m         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"log"
	"net/http"
)

const url = "localhost:8080"

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func sartDemoServer() {
	fmt.Println("=== HTTP demo server ===")

	http.HandleFunc("/", homeHandler)

	fmt.Println("Listenin on http://" + url)
	log.Fatal(http.ListenAndServe(url, nil))
}
