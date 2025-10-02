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

func (m *Model) GetConnectorById(_ int) {
}

func (m *Model) GetNodeById(_ int) (cell Cell) {
	return
}

func (m *Model) GetNodeIdByConnectorId(_ int) (id int) {
	return
}

func (m *Model) GetEmptyNodeInputsAndConfigValues() (nodes []Cell) {
	cellList := map[string]map[string]bool{}

	// Create cell list
	for _, cell := range m.Cells {
		if cell.Type == "senergy.NodeElement" {
			cellList[cell.Id] = map[string]bool{}
			for _, port := range cell.InPorts {
				cellList[cell.Id][port] = false
			}
			cellList[cell.Id]["_config"] = true
		}
	}

	// Check if cell input ports have links
	for _, cell := range m.Cells {
		if cell.Type == "link" {
			if cell.Target.Port[:3] == "in-" {
				delete(cellList[cell.Target.Id], cell.Target.Port[3:])
			} else {
				delete(cellList[cell.Target.Id], cell.Target.Port)
			}
		}
	}

	// Remove all cells without free input ports
	for cellId := range cellList {
		if len(cellList[cellId]) < 1 {
			delete(cellList, cellId)
		}
	}

	// Create node list and filter assigned ports
	if len(cellList) > 0 {
		for _, cell := range m.Cells {
			if _, ok := cellList[cell.Id]; ok {
				keys := make([]string, 0, len(cellList[cell.Id]))
				for k := range cellList[cell.Id] {
					if k != "_config" {
						keys = append(keys, k)
					}
				}
				cell.InPorts = keys
				nodes = append(nodes, cell)
			}
		}
	}
	return
}

func (m *Model) GetEmptyNodeInputs() (unassignedNodes []Cell) {
	cellList := map[string]map[string]bool{}

	// Create cell list
	for _, cell := range m.Cells {
		if cell.Type == "senergy.NodeElement" {
			cellList[cell.Id] = map[string]bool{}
			for _, port := range cell.InPorts {
				cellList[cell.Id][port] = false
			}
		}
	}

	// Check if cell input ports have links
	for _, cell := range m.Cells {
		if cell.Type == "link" {
			delete(cellList[cell.Target.Id], cell.Target.Port)
		}
	}

	// Remove all cells without free input ports
	for cellId := range cellList {
		if len(cellList[cellId]) < 1 {
			delete(cellList, cellId)
		}
	}

	// Create node list and filter assigned ports
	if len(cellList) > 0 {
		for _, cell := range m.Cells {
			if _, ok := cellList[cell.Id]; ok {
				keys := make([]string, 0, len(cellList[cell.Id]))
				for k := range cellList[cell.Id] {
					keys = append(keys, k)
				}
				cell.InPorts = keys
				unassignedNodes = append(unassignedNodes, cell)
			}
		}
	}
	return
}
