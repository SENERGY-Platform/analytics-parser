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
	"net/http"
	"encoding/json"
	"analytics-parser/lib"
	"github.com/gorilla/mux"
	"analytics-parser/parser"
	"fmt"
)

type Endpoint struct {
	flowApiService lib.FlowApiService
	flowParser *parser.FlowParser
}

func NewEndpoint(flowApiService lib.FlowApiService, operator_api lib.OperatorApiService) *Endpoint{
	ret := parser.NewFlowParser(flowApiService, operator_api)
	return &Endpoint{flowApiService ,ret}
}

func (e *Endpoint) getRootEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(lib.Response{"OK"})
}

func (e *Endpoint) getParseFlow(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	ret := e.flowParser.ParseFlow(vars["id"], e.getUserId(req))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(ret)
}

func (e *Endpoint) getGetInputs(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	ret := e.flowParser.GetInputs(vars["id"], e.getUserId(req))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(ret)
}

func (e *Endpoint) getUserId(req *http.Request) (userId string){
	userId = req.Header.Get("X-UserId")
	if userId == "" {
		userId = "admin"
	}
	fmt.Println("UserID: " + userId)
	return
}