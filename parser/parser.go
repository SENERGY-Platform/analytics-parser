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

package parser

import (
	"analytics-parser/lib"
	"analytics-parser/flows-api"
	"github.com/pkg/errors"
)

type FlowParser struct {
	flowApi lib.FlowApiService
}

func NewFlowParser (flowApi lib.FlowApiService) * FlowParser {

	return &FlowParser{flowApi}
}

func (f FlowParser) ParseFlow (id string, userId  string) Pipeline {
	var pipeline =  make(Pipeline)
	// Get flow to execute
	flow, _ := f.flowApi.GetFlowData(id, userId)

	// Create basic operator list
	for _, cell := range flow.Model.Cells {
		if cell.Type == "senergy.NodeElement" {
			var operator = Operator{cell.Id, cell.Name, cell.Image, make(map [string] InputTopic)}
			pipeline[cell.Id] = operator
		}
	}

	// Append input topics to operators
	for _, link := range flow.Model.Cells {
		if link.Type == "link" {
			node, _ := getNodeById(flow.Model, link.Source.Id)
			var outputTopic = getOperatorOutputTopic(node.Name)
			var topic = InputTopic {}
			var mapping = Mapping{link.Source.Port, link.Target.Port}

			if len(pipeline[link.Target.Id].InputTopics[outputTopic].Mappings) < 1 {
				topic.FilterType = "OperatorId"
				topic.FilterValue = link.Source.Id
			} else {
				topic = pipeline[link.Target.Id].InputTopics[outputTopic]
			}

			topic.Mappings = append(topic.Mappings, mapping)
			pipeline[link.Target.Id].InputTopics[outputTopic] = topic
		}
	}

	return pipeline
}

func (f FlowParser) GetInputsAndConfig (id string, userId string) ([] flows_api.Cell) {
	flow, _ := f.flowApi.GetFlowData(id, userId)
	return flow.Model.GetEmptyNodeInputsAndConfigValues()
}

func getOperatorOutputTopic (name string) (opName string) {
	opName = "analytics-" + name
	return
}

func getNodeById(model flows_api.Model, id string) (flows_api.Cell, error) {
	for _, cell := range model.Cells {
		if id == cell.Id {
			return cell, nil
		}
	}
	return flows_api.Cell{}, errors.New("Could not find any cell")
}
