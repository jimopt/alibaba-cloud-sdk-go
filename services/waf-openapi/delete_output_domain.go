package waf_openapi

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

// DeleteOutputDomain invokes the waf_openapi.DeleteOutputDomain API synchronously
func (client *Client) DeleteOutputDomain(request *DeleteOutputDomainRequest) (response *DeleteOutputDomainResponse, err error) {
	response = CreateDeleteOutputDomainResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteOutputDomainWithChan invokes the waf_openapi.DeleteOutputDomain API asynchronously
func (client *Client) DeleteOutputDomainWithChan(request *DeleteOutputDomainRequest) (<-chan *DeleteOutputDomainResponse, <-chan error) {
	responseChan := make(chan *DeleteOutputDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteOutputDomain(request)
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

// DeleteOutputDomainWithCallback invokes the waf_openapi.DeleteOutputDomain API asynchronously
func (client *Client) DeleteOutputDomainWithCallback(request *DeleteOutputDomainRequest, callback func(response *DeleteOutputDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteOutputDomainResponse
		var err error
		defer close(result)
		response, err = client.DeleteOutputDomain(request)
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

// DeleteOutputDomainRequest is the request struct for api DeleteOutputDomain
type DeleteOutputDomainRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	SourceIp   string `position:"Query" name:"SourceIp"`
	Domain     string `position:"Query" name:"Domain"`
	Lang       string `position:"Query" name:"Lang"`
	Region     string `position:"Query" name:"Region"`
}

// DeleteOutputDomainResponse is the response struct for api DeleteOutputDomain
type DeleteOutputDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Result    Result `json:"Result" xml:"Result"`
}

// CreateDeleteOutputDomainRequest creates a request to invoke DeleteOutputDomain API
func CreateDeleteOutputDomainRequest() (request *DeleteOutputDomainRequest) {
	request = &DeleteOutputDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("waf-openapi", "2019-09-10", "DeleteOutputDomain", "waf", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteOutputDomainResponse creates a response to parse from DeleteOutputDomain response
func CreateDeleteOutputDomainResponse() (response *DeleteOutputDomainResponse) {
	response = &DeleteOutputDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
