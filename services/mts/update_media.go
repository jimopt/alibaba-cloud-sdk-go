package mts

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

// invoke UpdateMedia api with *UpdateMediaRequest synchronously
// api document: https://help.aliyun.com/api/mts/updatemedia.html
func (client *Client) UpdateMedia(request *UpdateMediaRequest) (response *UpdateMediaResponse, err error) {
	response = CreateUpdateMediaResponse()
	err = client.DoAction(request, response)
	return
}

// invoke UpdateMedia api with *UpdateMediaRequest asynchronously
// api document: https://help.aliyun.com/api/mts/updatemedia.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateMediaWithChan(request *UpdateMediaRequest) (<-chan *UpdateMediaResponse, <-chan error) {
	responseChan := make(chan *UpdateMediaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateMedia(request)
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

// invoke UpdateMedia api with *UpdateMediaRequest asynchronously
// api document: https://help.aliyun.com/api/mts/updatemedia.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateMediaWithCallback(request *UpdateMediaRequest, callback func(response *UpdateMediaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateMediaResponse
		var err error
		defer close(result)
		response, err = client.UpdateMedia(request)
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

type UpdateMediaRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	MediaId              string           `position:"Query" name:"MediaId"`
	Title                string           `position:"Query" name:"Title"`
	Description          string           `position:"Query" name:"Description"`
	CoverURL             string           `position:"Query" name:"CoverURL"`
	CateId               requests.Integer `position:"Query" name:"CateId"`
	Tags                 string           `position:"Query" name:"Tags"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type UpdateMediaResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Media     Media  `json:"Media" xml:"Media"`
}

// create a request to invoke UpdateMedia API
func CreateUpdateMediaRequest() (request *UpdateMediaRequest) {
	request = &UpdateMediaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "UpdateMedia", "mts", "openAPI")
	return
}

// create a response to parse from UpdateMedia response
func CreateUpdateMediaResponse() (response *UpdateMediaResponse) {
	response = &UpdateMediaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
