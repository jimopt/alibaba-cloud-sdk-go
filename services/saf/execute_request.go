package saf

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

// ExecuteRequest invokes the saf.ExecuteRequest API synchronously
// api document: https://help.aliyun.com/api/saf/executerequest.html
func (client *Client) ExecuteRequest(request *ExecuteRequestRequest) (response *ExecuteRequestResponse, err error) {
	response = CreateExecuteRequestResponse()
	err = client.DoAction(request, response)
	return
}

// ExecuteRequestWithChan invokes the saf.ExecuteRequest API asynchronously
// api document: https://help.aliyun.com/api/saf/executerequest.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ExecuteRequestWithChan(request *ExecuteRequestRequest) (<-chan *ExecuteRequestResponse, <-chan error) {
	responseChan := make(chan *ExecuteRequestResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExecuteRequest(request)
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

// ExecuteRequestWithCallback invokes the saf.ExecuteRequest API asynchronously
// api document: https://help.aliyun.com/api/saf/executerequest.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ExecuteRequestWithCallback(request *ExecuteRequestRequest, callback func(response *ExecuteRequestResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExecuteRequestResponse
		var err error
		defer close(result)
		response, err = client.ExecuteRequest(request)
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

// ExecuteRequestRequest is the request struct for api ExecuteRequest
type ExecuteRequestRequest struct {
	*requests.RpcRequest
	ServiceParameters string `position:"Query" name:"ServiceParameters"`
	Service           string `position:"Query" name:"Service"`
}

// ExecuteRequestResponse is the response struct for api ExecuteRequest
type ExecuteRequestResponse struct {
	*responses.BaseResponse
	RequestId string                 `json:"RequestId" xml:"RequestId"`
	Code      int                    `json:"Code" xml:"Code"`
	Message   string                 `json:"Message" xml:"Message"`
	Data      map[string]interface{} `json:"Data" xml:"Data"`
}

// CreateExecuteRequestRequest creates a request to invoke ExecuteRequest API
func CreateExecuteRequestRequest() (request *ExecuteRequestRequest) {
	request = &ExecuteRequestRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("saf", "2019-05-21", "ExecuteRequest", "saf", "openAPI")
	return
}

// CreateExecuteRequestResponse creates a response to parse from ExecuteRequest response
func CreateExecuteRequestResponse() (response *ExecuteRequestResponse) {
	response = &ExecuteRequestResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
