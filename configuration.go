/*
OSHWA API

# Introduction  Welcome to the OSHWA Open Source Hardware Certification API. We hope that you will use it to build on top of the OSHWA certification program. This documentation provides information on how to use the OSHWA REST API to view, create, and search OSHWA certified open source hardware projects. You can learn more about OSHWA [here](https://www.oshwa.org/about/) and more about OSHWA’s free open source hardware certification program [here](https://certification.oshwa.org/).  This API supports both read and write functions. You can use the read functions to pull information about certified hardware from the [directory](https://certification.oshwa.org/list.html) in order to explore and present it in new ways. We encourage you to use it to find new ways to understand and visualize the world of open source hardware!  You can use the write functions to make it easier to submit registration applications to the program. Originally, the only way to submit hardware for registration was through OSHWA’s [application form](https://application.oshwa.org/apply). We hope that the write functionality will make it easier to integrate certification applications into existing workflows, and to connect platforms that already host open source hardware to the certification program.    If you have questions or comments about the API, please email us at info@oshwa.org.  We would also love to know how you use the API! We encourage you to contact us at info@oshwa.org, or let us know on twitter at @OHSummit.  # Tools This API is documented in **OpenAPI format**. It is built with [Swagger](http://swagger.io) and [ReDoc](https://github.com/Redocly/redoc).  # Using the API In order to use the API, you must register for an API key. You can get your own API key [here](https://certificationapi.oshwa.org/). If your token is not included or is invalid, the API will return an error. If you'd like to test the endpoints, check out our [Swagger](/endpoints) implementation, where you will be able to request a key, add the key to your requests, and explore OSHWA certified projects. You will also find code examples in this documentation.  # Pagination Project and Company results are returned in a wrapper object that contains total, limit, and offset values, which are useful for paginating over results.   ``` {   \"total\": 200, // Total number of matching items   \"offset\": 0, // Number of items skipped in request   \"limit\": 100, // Max number of items in request   \"items\": [ {...} ] // List of items } ```  The default `limit` for requests is 100, and the maximum `limit` is 1000. The default `offset` is 0. The above request returns the first 100 items. To get the next 100 items, the `offset` would be changed to 100. These parameters can be used to loop through the api and retrieve all items.  All items are returned in alphabetical order by `projectName`.   # Authentication  All OSHWA API endpoints require Bearer Token Authentication.  You can get an API key [here](https://certificationapi.oshwa.org/).  # Errors  The OSHWA API uses HTTP response status codes to indicate whether the response was successful. Detailed information about the response and how any errors might be resolved can be found in the body of the error message.  ``` {     \"error\": {         \"statusCode\": 422,         \"errorCode\": \"Unprocessable Entity\",         \"message\": \"Validation Error: Input validation failed\",         \"details\": [             {                 \"msg\": \"Responsible Party Type is required\",                 \"param\": \"responsiblePartyType\",                 \"location\": \"body\"             },         ]     } } ``` 

API version: 1.0.0
Contact: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken = contextKey("accesstoken")

	// ContextServerIndex uses a server configuration from the index.
	ContextServerIndex = contextKey("serverIndex")

	// ContextOperationServerIndices uses a server configuration from the index mapping.
	ContextOperationServerIndices = contextKey("serverOperationIndices")

	// ContextServerVariables overrides a server configuration variables.
	ContextServerVariables = contextKey("serverVariables")

	// ContextOperationServerVariables overrides a server configuration variables using operation specific values.
	ContextOperationServerVariables = contextKey("serverOperationVariables")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

// ServerVariable stores the information about a server variable
type ServerVariable struct {
	Description  string
	DefaultValue string
	EnumValues   []string
}

// ServerConfiguration stores the information about a server
type ServerConfiguration struct {
	URL string
	Description string
	Variables map[string]ServerVariable
}

// ServerConfigurations stores multiple ServerConfiguration items
type ServerConfigurations []ServerConfiguration

// Configuration stores the configuration of the API client
type Configuration struct {
	Host             string            `json:"host,omitempty"`
	Scheme           string            `json:"scheme,omitempty"`
	DefaultHeader    map[string]string `json:"defaultHeader,omitempty"`
	UserAgent        string            `json:"userAgent,omitempty"`
	Debug            bool              `json:"debug,omitempty"`
	Servers          ServerConfigurations
	OperationServers map[string]ServerConfigurations
	HTTPClient       *http.Client
}

// NewConfiguration returns a new Configuration object
func NewConfiguration() *Configuration {
	cfg := &Configuration{
		DefaultHeader:    make(map[string]string),
		UserAgent:        "OpenAPI-Generator/1.0.0/go",
		Debug:            false,
		Servers:          ServerConfigurations{
			{
				URL: "https://certificationapi.oshwa.org",
				Description: "No description provided",
			},
		},
		OperationServers: map[string]ServerConfigurations{
		},
	}
	return cfg
}

// AddDefaultHeader adds a new HTTP header to the default header in the request
func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

// URL formats template on a index using given variables
func (sc ServerConfigurations) URL(index int, variables map[string]string) (string, error) {
	if index < 0 || len(sc) <= index {
		return "", fmt.Errorf("index %v out of range %v", index, len(sc)-1)
	}
	server := sc[index]
	url := server.URL

	// go through variables and replace placeholders
	for name, variable := range server.Variables {
		if value, ok := variables[name]; ok {
			found := bool(len(variable.EnumValues) == 0)
			for _, enumValue := range variable.EnumValues {
				if value == enumValue {
					found = true
				}
			}
			if !found {
				return "", fmt.Errorf("the variable %s in the server URL has invalid value %v. Must be %v", name, value, variable.EnumValues)
			}
			url = strings.Replace(url, "{"+name+"}", value, -1)
		} else {
			url = strings.Replace(url, "{"+name+"}", variable.DefaultValue, -1)
		}
	}
	return url, nil
}

// ServerURL returns URL based on server settings
func (c *Configuration) ServerURL(index int, variables map[string]string) (string, error) {
	return c.Servers.URL(index, variables)
}

func getServerIndex(ctx context.Context) (int, error) {
	si := ctx.Value(ContextServerIndex)
	if si != nil {
		if index, ok := si.(int); ok {
			return index, nil
		}
		return 0, reportError("Invalid type %T should be int", si)
	}
	return 0, nil
}

func getServerOperationIndex(ctx context.Context, endpoint string) (int, error) {
	osi := ctx.Value(ContextOperationServerIndices)
	if osi != nil {
		if operationIndices, ok := osi.(map[string]int); !ok {
			return 0, reportError("Invalid type %T should be map[string]int", osi)
		} else {
			index, ok := operationIndices[endpoint]
			if ok {
				return index, nil
			}
		}
	}
	return getServerIndex(ctx)
}

func getServerVariables(ctx context.Context) (map[string]string, error) {
	sv := ctx.Value(ContextServerVariables)
	if sv != nil {
		if variables, ok := sv.(map[string]string); ok {
			return variables, nil
		}
		return nil, reportError("ctx value of ContextServerVariables has invalid type %T should be map[string]string", sv)
	}
	return nil, nil
}

func getServerOperationVariables(ctx context.Context, endpoint string) (map[string]string, error) {
	osv := ctx.Value(ContextOperationServerVariables)
	if osv != nil {
		if operationVariables, ok := osv.(map[string]map[string]string); !ok {
			return nil, reportError("ctx value of ContextOperationServerVariables has invalid type %T should be map[string]map[string]string", osv)
		} else {
			variables, ok := operationVariables[endpoint]
			if ok {
				return variables, nil
			}
		}
	}
	return getServerVariables(ctx)
}

// ServerURLWithContext returns a new server URL given an endpoint
func (c *Configuration) ServerURLWithContext(ctx context.Context, endpoint string) (string, error) {
	sc, ok := c.OperationServers[endpoint]
	if !ok {
		sc = c.Servers
	}

	if ctx == nil {
		return sc.URL(0, nil)
	}

	index, err := getServerOperationIndex(ctx, endpoint)
	if err != nil {
		return "", err
	}

	variables, err := getServerOperationVariables(ctx, endpoint)
	if err != nil {
		return "", err
	}

	return sc.URL(index, variables)
}
