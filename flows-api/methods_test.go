/*
 * Copyright 2020 InfAI (CC SES)
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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestModel_GetEmptyNodeInputsAndConfigValues(t *testing.T) {
	jsonFile, err := os.Open("flows-api_testdata/flow.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var flow Flow
	json.Unmarshal(byteValue, &flow)
	expected := []Cell{{
		Id:             "37eb2c6a-3879-4145-86c1-7d38fdd8b814",
		Name:           "adder",
		DeploymentType: "cloud",
		InPorts:        []string{"value", "timestamp"},
		OutPorts:       []string{"sum", " lastTimestamp"},
		Type:           "senergy.NodeElement",
		Source:         Port{},
		Target:         Port{},
		Image:          "image",
		Config:         nil,
		OperatorId:     "5d2da1c0de2c3100015801f3",
	}}
	if !reflect.DeepEqual(expected, flow.Model.GetEmptyNodeInputsAndConfigValues()) {
		fmt.Println(expected)
		fmt.Println(flow.Model.GetEmptyNodeInputsAndConfigValues())
		t.Error("structs do not match")
	}

}
