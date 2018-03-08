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

// invoke SetHttpErrorPageConfig api with *SetHttpErrorPageConfigRequest synchronously
// api document: https://help.aliyun.com/api/cdn/sethttperrorpageconfig.html
func (client *Client) SetHttpErrorPageConfig(request *SetHttpErrorPageConfigRequest) (response *SetHttpErrorPageConfigResponse, err error) {
	response = CreateSetHttpErrorPageConfigResponse()
	err = client.DoAction(request, response)
	return
}

// invoke SetHttpErrorPageConfig api with *SetHttpErrorPageConfigRequest asynchronously
// api document: https://help.aliyun.com/api/cdn/sethttperrorpageconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetHttpErrorPageConfigWithChan(request *SetHttpErrorPageConfigRequest) (<-chan *SetHttpErrorPageConfigResponse, <-chan error) {
	responseChan := make(chan *SetHttpErrorPageConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SetHttpErrorPageConfig(request)
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

// invoke SetHttpErrorPageConfig api with *SetHttpErrorPageConfigRequest asynchronously
// api document: https://help.aliyun.com/api/cdn/sethttperrorpageconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SetHttpErrorPageConfigWithCallback(request *SetHttpErrorPageConfigRequest, callback func(response *SetHttpErrorPageConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SetHttpErrorPageConfigResponse
		var err error
		defer close(result)
		response, err = client.SetHttpErrorPageConfig(request)
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

type SetHttpErrorPageConfigRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	DomainName    string           `position:"Query" name:"DomainName"`
	PageUrl       string           `position:"Query" name:"PageUrl"`
	ErrorCode     string           `position:"Query" name:"ErrorCode"`
}

type SetHttpErrorPageConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// create a request to invoke SetHttpErrorPageConfig API
func CreateSetHttpErrorPageConfigRequest() (request *SetHttpErrorPageConfigRequest) {
	request = &SetHttpErrorPageConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "SetHttpErrorPageConfig", "", "")
	return
}

// create a response to parse from SetHttpErrorPageConfig response
func CreateSetHttpErrorPageConfigResponse() (response *SetHttpErrorPageConfigResponse) {
	response = &SetHttpErrorPageConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
