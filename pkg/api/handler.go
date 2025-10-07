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

package api

import (
	"errors"
	"net/http"
	"os"

	"github.com/SENERGY-Platform/analytics-parser/pkg/parser"
	"github.com/SENERGY-Platform/analytics-parser/pkg/util"
	"github.com/gin-gonic/gin"
)

// getFlow returns a gin.HandlerFunc which handles GET requests to /flow/:id.
// This function calls the ParseFlow method of the given parser.FlowParser with the id from the request path,
// the user id from the request header, and the authorization token from the request header.
// If ParseFlow returns an error, it logs the error and returns an error response to the client.
// Otherwise, it returns a JSON response with the parsed flow data.
// @Summary Get flow info parsed
// @Description	Gets the parsed flow data
// @Tags Flow
// @Produce json
// @Param id path string true "Flow ID"
// @Success	200 {object} parser.Pipeline
// @Failure	500 {string} str
// @Router /flow/{id} [get]
func getFlow(flowParser parser.FlowParser) (string, string, gin.HandlerFunc) {
	return http.MethodGet, FlowIdPath, func(c *gin.Context) {
		id := c.Param("id")
		ret, err := flowParser.ParseFlow(id, c.GetString(UserIdKey), c.GetHeader("Authorization"))
		if err != nil {
			util.Logger.Error("error parsing flow", "error", err, "method", "GET", "path", FlowIdPath)
			_ = c.Error(errors.New(MessageSomethingWrong))
			return
		}
		c.JSON(http.StatusOK, ret)
	}
}

// getFlowInputs returns a gin.HandlerFunc which handles GET requests to /flow/{id}/inputs.
// This function calls the GetInputsAndConfig method of the given parser.FlowParser with the id from the request path,
// the user id from the request header, and the authorization token from the request header.
// If GetInputsAndConfig returns an error, it logs the error and returns an error response to the client.
// Otherwise, it returns a JSON response with the inputs and configuration of the flow.
// @Summary Get flow inputs
// @Description	Gets the inputs and configuration of the given flow
// @Tags Flow
// @Produce json
// @Param id path string true "Flow ID"
// @Success	200 {array} flows_api.Cell
// @Failure	500 {string} str
// @Router /flow/getinputs/{id} [get]
func getFlowInputs(flowParser parser.FlowParser) (string, string, gin.HandlerFunc) {
	return http.MethodGet, FlowInputsPath, func(c *gin.Context) {
		id := c.Param("id")
		ret, err := flowParser.GetInputsAndConfig(id, c.GetString(UserIdKey), c.GetHeader("Authorization"))
		if err != nil {
			util.Logger.Error("error getting inputs for flow "+id, "error", err, "method", "GET", "path", FlowInputsPath)
			_ = c.Error(errors.New(MessageSomethingWrong))
			return
		}
		c.JSON(http.StatusOK, ret)
	}
}

func getHealthCheckH(_ parser.FlowParser) (string, string, gin.HandlerFunc) {
	return http.MethodGet, HealthCheckPath, func(c *gin.Context) {
		c.Status(http.StatusOK)
	}
}

func getSwaggerDocH(_ parser.FlowParser) (string, string, gin.HandlerFunc) {
	return http.MethodGet, "/doc", func(gc *gin.Context) {
		if _, err := os.Stat("docs/swagger.json"); err != nil {
			_ = gc.Error(err)
			return
		}
		gc.Header("Content-Type", gin.MIMEJSON)
		gc.File("docs/swagger.json")
	}
}
