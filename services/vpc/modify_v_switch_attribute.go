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

// invoke ModifyVSwitchAttribute api with *ModifyVSwitchAttributeRequest synchronously
// api document: https://help.aliyun.com/api/vpc/modifyvswitchattribute.html
func (client *Client) ModifyVSwitchAttribute(request *ModifyVSwitchAttributeRequest) (response *ModifyVSwitchAttributeResponse, err error) {
	response = CreateModifyVSwitchAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// invoke ModifyVSwitchAttribute api with *ModifyVSwitchAttributeRequest asynchronously
// api document: https://help.aliyun.com/api/vpc/modifyvswitchattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyVSwitchAttributeWithChan(request *ModifyVSwitchAttributeRequest) (<-chan *ModifyVSwitchAttributeResponse, <-chan error) {
	responseChan := make(chan *ModifyVSwitchAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyVSwitchAttribute(request)
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

// invoke ModifyVSwitchAttribute api with *ModifyVSwitchAttributeRequest asynchronously
// api document: https://help.aliyun.com/api/vpc/modifyvswitchattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyVSwitchAttributeWithCallback(request *ModifyVSwitchAttributeRequest, callback func(response *ModifyVSwitchAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyVSwitchAttributeResponse
		var err error
		defer close(result)
		response, err = client.ModifyVSwitchAttribute(request)
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

type ModifyVSwitchAttributeRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	VSwitchId            string           `position:"Query" name:"VSwitchId"`
	VSwitchName          string           `position:"Query" name:"VSwitchName"`
	Description          string           `position:"Query" name:"Description"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type ModifyVSwitchAttributeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// create a request to invoke ModifyVSwitchAttribute API
func CreateModifyVSwitchAttributeRequest() (request *ModifyVSwitchAttributeRequest) {
	request = &ModifyVSwitchAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "ModifyVSwitchAttribute", "vpc", "openAPI")
	return
}

// create a response to parse from ModifyVSwitchAttribute response
func CreateModifyVSwitchAttributeResponse() (response *ModifyVSwitchAttributeResponse) {
	response = &ModifyVSwitchAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
