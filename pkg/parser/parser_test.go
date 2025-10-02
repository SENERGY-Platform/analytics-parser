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

package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	flowsapi "github.com/SENERGY-Platform/analytics-parser/pkg/flows-api"
)

func TestFlowParser_CreatePipelineList(t *testing.T) {
	jsonFile, err := os.Open("parser_testdata/flow.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var flow flowsapi.Flow
	json.Unmarshal(byteValue, &flow)
	parser := NewFlowParser(flowsapi.NewFlowApi(
		"",
	))
	expected := Pipeline{
		FlowId: "5ee0a2831b576d2534f04099",
		Operators: map[string]Operator{
			"22a28f5b-54d8-4e46-9ba9-c36dc6bd3da8": {
				Id:             "22a28f5b-54d8-4e46-9ba9-c36dc6bd3da8",
				Name:           "adder",
				OperatorId:     "5d2da1c0de2c3100015801f3",
				DeploymentType: "cloud",
				ImageId:        "image",
				InputTopics: []InputTopic{
					{
						TopicName:   "analytics-adder",
						FilterType:  "OperatorId",
						FilterValue: "37eb2c6a-3879-4145-86c1-7d38fdd8b814",
						Mappings: []Mapping{
							{"sum", "value"},
							{" lastTimestamp", "timestamp"}},
					},
				},
			},
			"37eb2c6a-3879-4145-86c1-7d38fdd8b814": {
				Id:             "37eb2c6a-3879-4145-86c1-7d38fdd8b814",
				Name:           "adder",
				OperatorId:     "5d2da1c0de2c3100015801f3",
				DeploymentType: "cloud",
				ImageId:        "image",
				InputTopics:    nil,
			},
		},
	}
	list := parser.CreatePipelineList(flow)
	if !reflect.DeepEqual(expected, list) {
		fmt.Println(expected)
		fmt.Println(list)
		file, _ := json.MarshalIndent(list, "", " ")
		_ = ioutil.WriteFile("./parser_testdata/test.json", file, 0644)
		t.Error("structs do not match")
	}
}

func TestFlowParser_CreatePipelineList2(t *testing.T) {
	jsonFile, err := os.Open("parser_testdata/flow2.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var flow flowsapi.Flow
	json.Unmarshal(byteValue, &flow)
	parser := NewFlowParser(flowsapi.NewFlowApi(
		"",
	))
	jsonFileResult, err := os.Open("parser_testdata/flow2-result.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFileResult.Close()
	byteValue, _ = ioutil.ReadAll(jsonFileResult)
	var expected Pipeline
	json.Unmarshal(byteValue, &expected)
	if !reflect.DeepEqual(expected, parser.CreatePipelineList(flow)) {
		fmt.Println("Expected:")
		fmt.Printf("%+v\n", expected)
		fmt.Println("Actual:")
		fmt.Printf("%+v\n", parser.CreatePipelineList(flow))
		file, _ := json.MarshalIndent(parser.CreatePipelineList(flow), "", " ")
		_ = ioutil.WriteFile("./parser_testdata/test.json", file, 0644)
		t.Error("structs do not match")
	}
}
