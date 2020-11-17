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
	"analytics-parser/flows-api"
	"analytics-parser/lib"
	"github.com/pkg/errors"
)

type FlowParser struct {
	flowApi lib.FlowApiService
}

func NewFlowParser(flowApi lib.FlowApiService) *FlowParser {

	return &FlowParser{flowApi}
}

func (f FlowParser) ParseFlow(id string, userId string, authorization string) (pipeline Pipeline, err error) {

	// Get flow to execute
	flow, err := f.flowApi.GetFlowData(id, userId, authorization)

	if err != nil {
		return
	}

	pipeline = f.CreatePipelineList(flow)

	return
}

func (f FlowParser) CreatePipelineList(flow flows_api.Flow) Pipeline {
	pipeline := Pipeline{FlowId: flow.Id, Image: flow.Image, Operators: make(map[string]Operator)}

	// Create basic operator list
	for _, cell := range flow.Model.Cells {
		if cell.Type == "senergy.NodeElement" {
			inputTopics := getInputTopics(flow, cell.Id)
			var operator = Operator{cell.Id, cell.Name, cell.OperatorId, cell.DeploymentType, cell.Image, inputTopics}
			pipeline.Operators[cell.Id] = operator
		}
	}
	return pipeline
}

func (f FlowParser) GetInputsAndConfig(id string, userId string, authorization string) ([]flows_api.Cell, error) {
	flow, err := f.flowApi.GetFlowData(id, userId, authorization)
	if err != nil {
		return nil, err
	}
	return flow.Model.GetEmptyNodeInputsAndConfigValues(), err
}

func getInputTopics(flow flows_api.Flow, cellId string) (inputTopics []InputTopic) {
	for _, link := range flow.Model.Cells {
		if link.Type == "link" && link.Target.Id == cellId {
			// create mapping
			var mapping Mapping
			if link.Source.Port[:4] == "out-" {
				mapping = Mapping{link.Source.Port[4:], link.Target.Port[3:]}
			} else {
				mapping = Mapping{link.Source.Port, link.Target.Port}
			}
			sourceNode, _ := getNodeById(flow.Model, link.Source.Id)
			targetNode, _ := getNodeById(flow.Model, link.Target.Id)
			topic := InputTopic{}
			if !checkInputTopicExists(inputTopics, link.Source.Id) {
				local := false
				name := sourceNode.Name
				if targetNode.DeploymentType == "local" {
					local = true
					name = sourceNode.Name + "/" + link.Source.Id
				}
				topic.TopicName = getOperatorOutputTopic(name, local)
				topic.FilterType = "OperatorId"
				topic.FilterValue = link.Source.Id
				topic.Mappings = append(topic.Mappings, mapping)
				inputTopics = append(inputTopics, topic)
			} else {
				for key, existingTopic := range inputTopics {
					if existingTopic.FilterValue == link.Source.Id {
						existingTopic.Mappings = append(existingTopic.Mappings, mapping)
						inputTopics[key] = existingTopic
					}
				}
			}
		}
	}
	return
}

func checkInputTopicExists(topics []InputTopic, topicId string) bool {
	if len(topics) == 0 {
		return false
	}
	for _, existingTopic := range topics {
		if existingTopic.FilterValue == topicId {
			return true
		}
	}
	return false
}

func getOperatorOutputTopic(name string, local bool) (opName string) {
	opName = "analytics-" + name
	if local {
		opName = "fog/analytics/" + name
	}
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
