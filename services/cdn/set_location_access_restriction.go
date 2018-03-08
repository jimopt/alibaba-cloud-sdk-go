package cdn

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

// invoke SetLocationAccessRestriction api with *SetLocationAccessRestrictionRequest synchronously
// api document: https://help.aliyun.com/api/cdn/setlocationaccessrestriction.html
func (client *Client) SetLocationAccessRestriction(request *SetLocationAccessRestrictionRequest) (response *SetLocationAccessRestrictionResponse, err error) {
	response = CreateSetLocationAccessRestrictionResponse()
	err = client.DoAction(request, response)
	return
}

// invoke SetLocationAccessRestriction api with *SetLocationAccessRestrictionRequest asynchronously
// api document: https://help.aliyun.com/api/cdn/setlocationaccessrestriction.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetLocationAccessRestrictionWithChan(request *SetLocationAccessRestrictionRequest) (<-chan *SetLocationAccessRestrictionResponse, <-chan error) {
	responseChan := make(chan *SetLocationAccessRestrictionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetLocationAccessRestriction(request)
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

// invoke SetLocationAccessRestriction api with *SetLocationAccessRestrictionRequest asynchronously
// api document: https://help.aliyun.com/api/cdn/setlocationaccessrestriction.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetLocationAccessRestrictionWithCallback(request *SetLocationAccessRestrictionRequest, callback func(response *SetLocationAccessRestrictionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetLocationAccessRestrictionResponse
		var err error
		defer close(result)
		response, err = client.SetLocationAccessRestriction(request)
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

type SetLocationAccessRestrictionRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
	Location      string           `position:"Query" name:"Location"`
	Type          string           `position:"Query" name:"Type"`
}

type SetLocationAccessRestrictionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// create a request to invoke SetLocationAccessRestriction API
func CreateSetLocationAccessRestrictionRequest() (request *SetLocationAccessRestrictionRequest) {
	request = &SetLocationAccessRestrictionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "SetLocationAccessRestriction", "", "")
	return
}

// create a response to parse from SetLocationAccessRestriction response
func CreateSetLocationAccessRestrictionResponse() (response *SetLocationAccessRestrictionResponse) {
	response = &SetLocationAccessRestrictionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
