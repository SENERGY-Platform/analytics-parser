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
		Name:           flow.Model.Cells[0].Name,
		DeploymentType: flow.Model.Cells[0].DeploymentType,
		Position:       flow.Model.Cells[0].Position,
		InPorts:        []string{"value", "timestamp"},
		OutPorts:       []string{"sum", " lastTimestamp"},
		Type:           "senergy.NodeElement",
		Source:         flow.Model.Cells[0].Source,
		Target:         flow.Model.Cells[0].Target,
		Image:          flow.Model.Cells[0].Image,
		Config:         nil,
		OperatorId:     flow.Model.Cells[0].OperatorId,
	},
		{
			Id:             "22a28f5b-54d8-4e46-9ba9-c36dc6bd3da8",
			Name:           flow.Model.Cells[1].Name,
			DeploymentType: flow.Model.Cells[1].DeploymentType,
			Position:       flow.Model.Cells[1].Position,
			InPorts:        []string{},
			OutPorts:       []string{"sum", " lastTimestamp"},
			Type:           "senergy.NodeElement",
			Source:         flow.Model.Cells[1].Source,
			Target:         flow.Model.Cells[1].Target,
			Image:          flow.Model.Cells[1].Image,
			Config:         nil,
			OperatorId:     flow.Model.Cells[1].OperatorId,
		},
	}

	if !reflect.DeepEqual(expected, flow.Model.GetEmptyNodeInputsAndConfigValues()) {
		fmt.Printf("%#v \n", expected)
		fmt.Printf("%#v \n", flow.Model.GetEmptyNodeInputsAndConfigValues())
		t.Error("structs do not match")
	}

}
