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
	"log"

	deploymentLocationLib "github.com/SENERGY-Platform/analytics-fog-lib/lib/location"
	operatorLib "github.com/SENERGY-Platform/analytics-fog-lib/lib/operator"
	"github.com/SENERGY-Platform/analytics-parser/flows-api"
	"github.com/SENERGY-Platform/analytics-parser/lib"
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

func getLinksToTargetNode(flow flows_api.Flow, cellId string) (links []flows_api.Cell) {
	for _, link := range flow.Model.Cells {
		if link.Type == "link" && link.Target.Id == cellId {
			links = append(links, link)
		}
	}
	return 
} 

func getLinksFromSourceNode(cells []flows_api.Cell, cellId string) (links []flows_api.Cell) {
	for _, link := range cells {
		if link.Type == "link" && link.Source.Id == cellId {
			links = append(links, link)
		}
	}
	return 
} 

func (f FlowParser) DecideDeploymentPlatform(cells []flows_api.Cell) (newCells []flows_api.Cell) {
	// atm run operator on cloud when it supports both
	for _, cell := range cells {
		if cell.Type == "senergy.NodeElement" {
			deploymentType := cell.DeploymentType
			if deploymentType == "" || deploymentType == "both" {
				cell.DeploymentType = deploymentLocationLib.Cloud
			}
		}
		newCells = append(newCells, cell)
	}
	return
}

func (f FlowParser) CreatePipelineList(flow flows_api.Flow) Pipeline {
	pipeline := Pipeline{FlowId: flow.Id, Image: flow.Image, Operators: make(map[string]Operator)}

	// Create basic operator list
	cells := f.DecideDeploymentPlatform(flow.Model.Cells)
	for _, cell := range cells {
		if cell.Type == "senergy.NodeElement" {
			inputTopics := getInputTopics(flow, cell)

			deploymentType := cell.DeploymentType

			var upstreamConfig UpstreamConfig
			var downstreamConfig DownstreamConfig
			if deploymentType == deploymentLocationLib.Local {
				log.Println("Check if local operator output of " + cell.Id + " shall be forwarded to cloud")
				upstreamConfig.Enabled = checkIfLocalOutputForwardedToPlatform(cells, cell.Id)
			} else if deploymentType == deploymentLocationLib.Cloud {
				log.Println("Check if cloud operator output of " + cell.Id + " shall be forwarded to fog")
				downstreamConfig.Enabled = checkIfCloudOutputForwardedToFog(cells, cell.Id)
			}

			var operator = Operator{
				cell.Id, 
				cell.Name, 
				cell.OperatorId, 
				deploymentType, 
				cell.Image, 
				inputTopics, 
				cell.Cost,
				upstreamConfig,
				downstreamConfig,
			}
			log.Printf("%s will be deployed to %s - Cloud2Fog: %t Fog2Cloud: %t\n", operator.Id, operator.DeploymentType, downstreamConfig.Enabled, upstreamConfig.Enabled)
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

func checkIfLocalOutputForwardedToPlatform(cells []flows_api.Cell, cellId string) bool {
	// Check whether there exists at least one operator after this that is deployed on the cloud.
	// Then the output of this operator will be forwarded to the platform where it can be accessed by the next operator.
	linksFromNode := getLinksFromSourceNode(cells, cellId)
	for _, link := range linksFromNode {
		targetNode, _ := getNodeById(cells, link.Target.Id)
		if targetNode.DeploymentType == deploymentLocationLib.Cloud {
			return true
		}
	}
	return false
}

func checkIfCloudOutputForwardedToFog(cells []flows_api.Cell, cellId string) bool {
	// Check whether there exists at least one operator after this that is deployed on the fog.
	// Then the output of this operator will be forwarded to fog where it can be accessed by the next operator.

	linksFromNode := getLinksFromSourceNode(cells, cellId)
	for _, link := range linksFromNode {
		targetNode, _ := getNodeById(cells, link.Target.Id)
		if targetNode.DeploymentType == deploymentLocationLib.Local {
			return true
		}
	}
	return false
}

func getInputTopics(flow flows_api.Flow, cell flows_api.Cell) (inputTopics []InputTopic) {
	// Generate input topic names. This will add the required topic prefixes depending whether the operator
	// is deployed on the cloud or local

	linksToTarget := getLinksToTargetNode(flow, cell.Id)

	for _, link := range linksToTarget {
		// create mapping
		var mapping Mapping
		if link.Source.Port[:4] == "out-" {
			mapping = Mapping{link.Source.Port[4:], link.Target.Port[3:]}
		} else {
			mapping = Mapping{link.Source.Port, link.Target.Port}
		}
		sourceNode, _ := getNodeById(flow.Model.Cells, link.Source.Id)
		topic := InputTopic{}
		if !checkInputTopicExists(inputTopics, link.Source.Id) {
			// TODO error handling
			var topicName string
			if cell.DeploymentType == deploymentLocationLib.Local {
				topicName = operatorLib.GenerateFogOperatorTopic(sourceNode.Name, sourceNode.Id, "")
			} else if cell.DeploymentType == deploymentLocationLib.Cloud {
				topicName = operatorLib.GenerateCloudOperatorTopic(sourceNode.Name)
			}
			topic.TopicName = topicName
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

func getNodeById(cells []flows_api.Cell, id string) (flows_api.Cell, error) {
	for _, cell := range cells {
		if id == cell.Id {
			return cell, nil
		}
	}
	return flows_api.Cell{}, errors.New("Could not find any cell")
}
