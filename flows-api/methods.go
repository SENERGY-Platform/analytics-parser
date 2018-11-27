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

func (m *Model) GetConnectorById(id int) (connector Connector) {
	for _, node := range m.Nodes {
		for _, con := range node.Connectors {
			if id == con.Id {
				connector = con
			}
		}
	}
	return
}

func (m *Model) GetNodeById(id int) (node Node) {
	for _, n := range m.Nodes {
		if id == n.Id {
			node = n
		}
	}
	return
}

func (m *Model) GetNodeIdByConnectorId(connectorId int) (id int) {
	for _, node := range m.Nodes {
		for _, connector := range node.Connectors {
			if connectorId == connector.Id {
				id = node.Id
			}
		}
	}
	return
}

func (m *Model) GetEmptyNodeInputs() (unassignedNodes []Node) {
	assignedEdges := map[int]bool{}
	for _, edge := range m.Edges {
		assignedEdges[edge.Source] = true
		assignedEdges[edge.Destination] = true
	}
	for _, node := range m.Nodes {
		var unassignedConnectors []Connector
		for _, connector := range node.Connectors {
			if _, ok := assignedEdges[connector.Id]; !ok && connector.Type == "topConnector" {
				unassignedConnectors = append(unassignedConnectors, connector)
			}
		}
		if len(unassignedConnectors) > 0 {
			node.Connectors = unassignedConnectors
			unassignedNodes = append(unassignedNodes, node)
		}
	}
	return
}
