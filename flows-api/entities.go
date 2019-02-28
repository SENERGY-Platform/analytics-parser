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

type Flow struct {
	Id    string `json:"_id,omitempty"`
	Name  string `json:"name,omitempty"`
	Model Model  `json:"model,omitempty"`
}

type Model struct {
	Cells []Cell `json:"cells,omitempty"`
}

type Cell struct {
	Id       string    `json:"id,omitempty"`
	Name  string    `json:"name,omitempty"`
	InPorts  [] string `json:"inPorts,omitempty"`
	OutPorts [] string `json:"outPorts,omitempty"`
	Type     string    `json:"type,omitempty"`
	Source   Port   `json:"source,omitempty"`
	Target   Port   `json:"target,omitempty"`
	Image string `json:"image,omitempty"`
	Config [] ConfigValue `json:"config,omitempty"`
}

type Port struct {
	Id   string `json:"id,omitempty"`
	Port string `json:"port,omitempty"`
}

type ConfigValue struct {
	Name   string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}
