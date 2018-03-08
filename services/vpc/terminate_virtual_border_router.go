package vpc

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

// invoke TerminateVirtualBorderRouter api with *TerminateVirtualBorderRouterRequest synchronously
// api document: https://help.aliyun.com/api/vpc/terminatevirtualborderrouter.html
func (client *Client) TerminateVirtualBorderRouter(request *TerminateVirtualBorderRouterRequest) (response *TerminateVirtualBorderRouterResponse, err error) {
	response = CreateTerminateVirtualBorderRouterResponse()
	err = client.DoAction(request, response)
	return
}

// invoke TerminateVirtualBorderRouter api with *TerminateVirtualBorderRouterRequest asynchronously
// api document: https://help.aliyun.com/api/vpc/terminatevirtualborderrouter.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TerminateVirtualBorderRouterWithChan(request *TerminateVirtualBorderRouterRequest) (<-chan *TerminateVirtualBorderRouterResponse, <-chan error) {
	responseChan := make(chan *TerminateVirtualBorderRouterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TerminateVirtualBorderRouter(request)
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

// invoke TerminateVirtualBorderRouter api with *TerminateVirtualBorderRouterRequest asynchronously
// api document: https://help.aliyun.com/api/vpc/terminatevirtualborderrouter.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) TerminateVirtualBorderRouterWithCallback(request *TerminateVirtualBorderRouterRequest, callback func(response *TerminateVirtualBorderRouterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TerminateVirtualBorderRouterResponse
		var err error
		defer close(result)
		response, err = client.TerminateVirtualBorderRouter(request)
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

type TerminateVirtualBorderRouterRequest struct {
	*requests.RpcRequest
	VbrId                string           `position:"Query" name:"VbrId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type TerminateVirtualBorderRouterResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// create a request to invoke TerminateVirtualBorderRouter API
func CreateTerminateVirtualBorderRouterRequest() (request *TerminateVirtualBorderRouterRequest) {
	request = &TerminateVirtualBorderRouterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "TerminateVirtualBorderRouter", "vpc", "openAPI")
	return
}

// create a response to parse from TerminateVirtualBorderRouter response
func CreateTerminateVirtualBorderRouterResponse() (response *TerminateVirtualBorderRouterResponse) {
	response = &TerminateVirtualBorderRouterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
