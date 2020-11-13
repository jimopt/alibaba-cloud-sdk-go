package dataworks_public

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DeleteQualityRule invokes the dataworks_public.DeleteQualityRule API synchronously
func (client *Client) DeleteQualityRule(request *DeleteQualityRuleRequest) (response *DeleteQualityRuleResponse, err error) {
	response = CreateDeleteQualityRuleResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteQualityRuleWithChan invokes the dataworks_public.DeleteQualityRule API asynchronously
func (client *Client) DeleteQualityRuleWithChan(request *DeleteQualityRuleRequest) (<-chan *DeleteQualityRuleResponse, <-chan error) {
	responseChan := make(chan *DeleteQualityRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteQualityRule(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DeleteQualityRuleWithCallback invokes the dataworks_public.DeleteQualityRule API asynchronously
func (client *Client) DeleteQualityRuleWithCallback(request *DeleteQualityRuleRequest, callback func(response *DeleteQualityRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteQualityRuleResponse
		var err error
		defer close(result)
		response, err = client.DeleteQualityRule(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DeleteQualityRuleRequest is the request struct for api DeleteQualityRule
type DeleteQualityRuleRequest struct {
	*requests.RpcRequest
	ProjectName string           `position:"Body" name:"ProjectName"`
	RuleId      requests.Integer `position:"Body" name:"RuleId"`
}

// DeleteQualityRuleResponse is the response struct for api DeleteQualityRule
type DeleteQualityRuleResponse struct {
	*responses.BaseResponse
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Data           bool   `json:"Data" xml:"Data"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success        bool   `json:"Success" xml:"Success"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteQualityRuleRequest creates a request to invoke DeleteQualityRule API
func CreateDeleteQualityRuleRequest() (request *DeleteQualityRuleRequest) {
	request = &DeleteQualityRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "DeleteQualityRule", "", "")
	request.Method = requests.POST
	return
}

// CreateDeleteQualityRuleResponse creates a response to parse from DeleteQualityRule response
func CreateDeleteQualityRuleResponse() (response *DeleteQualityRuleResponse) {
	response = &DeleteQualityRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}