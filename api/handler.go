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
	"slices"
	"strconv"
	"strings"

	"github.com/SENERGY-Platform/analytics-parser/flows-api"
	"github.com/SENERGY-Platform/analytics-parser/lib"
	"github.com/SENERGY-Platform/analytics-parser/parser"
	"github.com/SENERGY-Platform/service-commons/pkg/jwt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CreateServer() {
	f := flows_api.NewFlowApi(
		lib.GetEnv("FLOW_API_ENDPOINT", ""),
	)
	serv := parser.NewFlowParser(f)

	port := lib.GetEnv("SERVER_API_PORT", "8000")

	lib.GetLogger().Info("starting server on port " + port)
	DEBUG, err := strconv.ParseBool(lib.GetEnv("DEBUG", "false"))
	if err != nil {
		lib.GetLogger().Error("Error loading debug value", "error", err)
		DEBUG = false
	}
	if !DEBUG {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "DELETE", "OPTIONS", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	prefix := r.Group(lib.GetEnv("ROUTE_PREFIX", ""))

	prefix.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	prefix.GET("/flow/:id", func(c *gin.Context) {
		id := c.Param("id")
		ret, err := serv.ParseFlow(id, getUserId(c), c.GetHeader("Authorization"))
		if err != nil {
			lib.GetLogger().Error("error parsing flow", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
			return
		}
		c.JSON(http.StatusOK, ret)
	})

	prefix.GET("/flow/getinputs/:id", func(c *gin.Context) {
		id := c.Param("id")
		ret, err := serv.GetInputsAndConfig(id, getUserId(c), c.GetHeader("Authorization"))
		if err != nil {
			lib.GetLogger().Error("error getting inputs", "error", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
			return
		}
		c.JSON(http.StatusOK, ret)
	})

	if !DEBUG {
		err = r.Run(":" + port)
	} else {
		err = r.Run("127.0.0.1:" + port)
	}
	if err == nil {
		lib.GetLogger().Error("could not start api server", "error", err)
	}
}

func getUserId(c *gin.Context) (userId string) {
	forUser := c.Query("for_user")
	if forUser != "" {

		roles := strings.Split(c.GetHeader("X-User-Roles"), ", ")
		if slices.Contains[[]string](roles, "admin") {
			return forUser
		}
	}

	userId = c.GetHeader("X-UserId")
	if userId == "" {
		if c.GetHeader("Authorization") != "" {
			claims, err := jwt.Parse(c.GetHeader("Authorization"))
			if err != nil {
				err = errors.New("Error parsing token: " + err.Error())
				return
			}
			userId = claims.Sub
			if userId == "" {
				userId = "dummy"
			}
		}
	}
	return
}
