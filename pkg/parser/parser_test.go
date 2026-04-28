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
	"os"
	"testing"

	"github.com/SENERGY-Platform/analytics-parser/lib"
	flowsapi "github.com/SENERGY-Platform/analytics-parser/pkg/flows-api"
	"github.com/SENERGY-Platform/analytics-parser/pkg/util"
	"github.com/google/go-cmp/cmp"
)

func TestFlowParser_CreatePipelineList(t *testing.T) {
	util.InitStructLogger("debug")

	flowData, err := os.ReadFile("parser_testdata/flow.json")
	if err != nil {
		t.Fatalf("failed to read flow file: %v", err)
	}
	var flow flowsapi.Flow
	if err := json.Unmarshal(flowData, &flow); err != nil {
		t.Fatalf("failed to unmarshal flow: %v", err)
	}

	expected := lib.Pipeline{
		FlowId: "5ee0a2831b576d2534f04099",
		Image:  "image",
		Operators: map[string]lib.Operator{
			"22a28f5b-54d8-4e46-9ba9-c36dc6bd3da8": {
				Id:             "22a28f5b-54d8-4e46-9ba9-c36dc6bd3da8",
				Name:           "adder",
				OperatorId:     "5d2da1c0de2c3100015801f3",
				DeploymentType: "cloud",
				ImageId:        "image",
				InputTopics: []lib.InputTopic{
					{
						TopicName:   "analytics-adder",
						FilterType:  "OperatorId",
						FilterValue: "37eb2c6a-3879-4145-86c1-7d38fdd8b814",
						Mappings: []lib.Mapping{
							{"sum", "value"},
							{" lastTimestamp", "timestamp"},
						},
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

	parser := NewFlowParser(flowsapi.NewFlowApi(""))
	actual := parser.CreatePipelineList(flow)

	if diff := cmp.Diff(expected, actual); diff != "" {
		if out, err := json.MarshalIndent(actual, "", "  "); err == nil {
			_ = os.WriteFile("./parser_testdata/test.json", out, 0644)
		}
		t.Errorf("CreatePipelineList() mismatch (-want +got):\n%s", diff)
	}
}

func TestFlowParser_CreatePipelineList2(t *testing.T) {
	util.InitStructLogger("debug")

	flowData, err := os.ReadFile("parser_testdata/flow2.json")
	if err != nil {
		t.Fatalf("failed to open flow file: %v", err)
	}
	var flow flowsapi.Flow
	if err := json.Unmarshal(flowData, &flow); err != nil {
		t.Fatalf("failed to unmarshal flow: %v", err)
	}

	resultData, err := os.ReadFile("parser_testdata/flow2-result.json")
	if err != nil {
		t.Fatalf("failed to open result file: %v", err)
	}
	var expected lib.Pipeline
	if err := json.Unmarshal(resultData, &expected); err != nil {
		t.Fatalf("failed to unmarshal expected result: %v", err)
	}

	parser := NewFlowParser(flowsapi.NewFlowApi(""))
	actual := parser.CreatePipelineList(flow)

	if diff := cmp.Diff(expected, actual); diff != "" {
		if out, err := json.MarshalIndent(actual, "", "  "); err == nil {
			_ = os.WriteFile("./parser_testdata/test.json", out, 0644)
		}
		t.Errorf("CreatePipelineList() mismatch (-want +got):\n%s", diff)
	}
}

func TestFlowParser_CreatePipelineListLocal(t *testing.T) {
	util.InitStructLogger("debug")

	flowData, err := os.ReadFile("parser_testdata/flow_local.json")
	if err != nil {
		t.Fatalf("failed to open flow file: %v", err)
	}
	var flow flowsapi.Flow
	if err := json.Unmarshal(flowData, &flow); err != nil {
		t.Fatalf("failed to unmarshal flow: %v", err)
	}

	resultData, err := os.ReadFile("parser_testdata/flow_local_result.json")
	if err != nil {
		t.Fatalf("failed to open result file: %v", err)
	}
	var expected lib.Pipeline
	if err := json.Unmarshal(resultData, &expected); err != nil {
		t.Fatalf("failed to unmarshal expected result: %v", err)
	}

	parser := NewFlowParser(flowsapi.NewFlowApi(""))
	actual := parser.CreatePipelineList(flow)

	if diff := cmp.Diff(expected, actual); diff != "" {
		if out, err := json.MarshalIndent(actual, "", "  "); err == nil {
			_ = os.WriteFile("./parser_testdata/test.json", out, 0644)
		}
		t.Errorf("CreatePipelineList() mismatch (-want +got):\n%s", diff)
	}
}
