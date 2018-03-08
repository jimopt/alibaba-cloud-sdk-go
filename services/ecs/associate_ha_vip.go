package ecs

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

// invoke AssociateHaVip api with *AssociateHaVipRequest synchronously
// api document: https://help.aliyun.com/api/ecs/associatehavip.html
func (client *Client) AssociateHaVip(request *AssociateHaVipRequest) (response *AssociateHaVipResponse, err error) {
	response = CreateAssociateHaVipResponse()
	err = client.DoAction(request, response)
	return
}

// invoke AssociateHaVip api with *AssociateHaVipRequest asynchronously
// api document: https://help.aliyun.com/api/ecs/associatehavip.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AssociateHaVipWithChan(request *AssociateHaVipRequest) (<-chan *AssociateHaVipResponse, <-chan error) {
	responseChan := make(chan *AssociateHaVipResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AssociateHaVip(request)
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

// invoke AssociateHaVip api with *AssociateHaVipRequest asynchronously
// api document: https://help.aliyun.com/api/ecs/associatehavip.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AssociateHaVipWithCallback(request *AssociateHaVipRequest, callback func(response *AssociateHaVipResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AssociateHaVipResponse
		var err error
		defer close(result)
		response, err = client.AssociateHaVip(request)
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

type AssociateHaVipRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	HaVipId              string           `position:"Query" name:"HaVipId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
}

type AssociateHaVipResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// create a request to invoke AssociateHaVip API
func CreateAssociateHaVipRequest() (request *AssociateHaVipRequest) {
	request = &AssociateHaVipRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "AssociateHaVip", "ecs", "openAPI")
	return
}

// create a response to parse from AssociateHaVip response
func CreateAssociateHaVipResponse() (response *AssociateHaVipResponse) {
	response = &AssociateHaVipResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
