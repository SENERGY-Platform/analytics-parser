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
	"encoding/json"
	"net/http"

	"github.com/SENERGY-Platform/analytics-parser/lib"
	"github.com/SENERGY-Platform/analytics-parser/parser"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

type Endpoint struct {
	flowApiService lib.FlowApiService
	flowParser     *parser.FlowParser
}

func NewEndpoint(flowApiService lib.FlowApiService) *Endpoint {
	ret := parser.NewFlowParser(flowApiService)
	return &Endpoint{flowApiService, ret}
}

func (e *Endpoint) getRootEndpoint(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	_ = json.NewEncoder(w).Encode(lib.Response{Message: "OK"})
}

func (e *Endpoint) getParseFlow(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	ret, err := e.flowParser.ParseFlow(vars["id"], e.getUserId(req), req.Header.Get("Authorization"))
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		lib.GetLogger().Error("error parsing flow", "error", err)
		w.WriteHeader(500)
		_ = json.NewEncoder(w).Encode(lib.Response{Message: err.Error()})
	} else {
		w.WriteHeader(200)
		_ = json.NewEncoder(w).Encode(ret)
	}
	w.Header().Set("Content-Type", "application/json")
}

func (e *Endpoint) getGetInputs(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	ret, err := e.flowParser.GetInputsAndConfig(vars["id"], e.getUserId(req), req.Header.Get("Authorization"))
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		lib.GetLogger().Error("error getting inputs", "error", err)
		w.WriteHeader(500)
		_ = json.NewEncoder(w).Encode(lib.Response{Message: err.Error()})
	} else {
		w.WriteHeader(200)
		_ = json.NewEncoder(w).Encode(ret)
	}
}

func (e *Endpoint) getUserId(req *http.Request) (userId string) {
	userId = req.Header.Get("X-UserId")
	if userId == "" {
		if req.Header.Get("Authorization") != "" {
			_, claims := parseJWTToken(req.Header.Get("Authorization")[7:])
			userId = claims.Sub
			if userId == "" {
				userId = "dummy"
			}
		}
	}
	return
}

func parseJWTToken(encodedToken string) (token *jwt.Token, claims lib.Claims) {
	token, _ = jwt.ParseWithClaims(encodedToken, &claims, nil)
	return
}
