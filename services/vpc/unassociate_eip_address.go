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

// invoke UnassociateEipAddress api with *UnassociateEipAddressRequest synchronously
// api document: https://help.aliyun.com/api/vpc/unassociateeipaddress.html
func (client *Client) UnassociateEipAddress(request *UnassociateEipAddressRequest) (response *UnassociateEipAddressResponse, err error) {
	response = CreateUnassociateEipAddressResponse()
	err = client.DoAction(request, response)
	return
}

// invoke UnassociateEipAddress api with *UnassociateEipAddressRequest asynchronously
// api document: https://help.aliyun.com/api/vpc/unassociateeipaddress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnassociateEipAddressWithChan(request *UnassociateEipAddressRequest) (<-chan *UnassociateEipAddressResponse, <-chan error) {
	responseChan := make(chan *UnassociateEipAddressResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnassociateEipAddress(request)
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

// invoke UnassociateEipAddress api with *UnassociateEipAddressRequest asynchronously
// api document: https://help.aliyun.com/api/vpc/unassociateeipaddress.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UnassociateEipAddressWithCallback(request *UnassociateEipAddressRequest, callback func(response *UnassociateEipAddressResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnassociateEipAddressResponse
		var err error
		defer close(result)
		response, err = client.UnassociateEipAddress(request)
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

type UnassociateEipAddressRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	AllocationId         string           `position:"Query" name:"AllocationId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	InstanceType         string           `position:"Query" name:"InstanceType"`
}

type UnassociateEipAddressResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// create a request to invoke UnassociateEipAddress API
func CreateUnassociateEipAddressRequest() (request *UnassociateEipAddressRequest) {
	request = &UnassociateEipAddressRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "UnassociateEipAddress", "vpc", "openAPI")
	return
}

// create a response to parse from UnassociateEipAddress response
func CreateUnassociateEipAddressResponse() (response *UnassociateEipAddressResponse) {
	response = &UnassociateEipAddressResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
