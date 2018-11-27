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

package flows_api

import (
	"github.com/parnurzeal/gorequest"
	"encoding/json"
)

type FlowApi struct {
	url string
}

func NewFlowApi(url string) *FlowApi {
	return &FlowApi{url}
}

func (f FlowApi) GetFlowData(id string, userId string) (flow Flow, err error) {
	request := gorequest.New()
	_, body, _ := request.Get(f.url+"flow/" + id + "/").Set("X-UserId", userId).End()
	err = json.Unmarshal([]byte(body), &flow)
	return
}