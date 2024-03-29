/*
 * Copyright 2018 InfAI (CC SES)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package api

import (
	"fmt"
	"github.com/SENERGY-Platform/analytics-parser/flows-api"
	"github.com/SENERGY-Platform/analytics-parser/lib"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func CreateServer() {
	f := flows_api.NewFlowApi(
		lib.GetEnv("FLOW_API_ENDPOINT", ""),
	)
	port := lib.GetEnv("API_PORT", "8000")
	fmt.Print("Starting Server at port " + port + "\n")
	router := mux.NewRouter()

	e := NewEndpoint(f)
	router.HandleFunc("/", e.getRootEndpoint).Methods("GET")
	router.HandleFunc("/flow/{id}", e.getParseFlow).Methods("GET")
	router.HandleFunc("/flow/getinputs/{id}", e.getGetInputs).Methods("GET")
	handler := cors.New(cors.Options{
		AllowedHeaders: []string{"Content-Type", "Authorization"},
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
	}).Handler(router)
	logger := lib.NewLogger(handler, "CALL")
	log.Fatal(http.ListenAndServe(":"+port, logger))
}
