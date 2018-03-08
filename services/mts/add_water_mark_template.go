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

// invoke AddWaterMarkTemplate api with *AddWaterMarkTemplateRequest synchronously
// api document: https://help.aliyun.com/api/mts/addwatermarktemplate.html
func (client *Client) AddWaterMarkTemplate(request *AddWaterMarkTemplateRequest) (response *AddWaterMarkTemplateResponse, err error) {
	response = CreateAddWaterMarkTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// invoke AddWaterMarkTemplate api with *AddWaterMarkTemplateRequest asynchronously
// api document: https://help.aliyun.com/api/mts/addwatermarktemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddWaterMarkTemplateWithChan(request *AddWaterMarkTemplateRequest) (<-chan *AddWaterMarkTemplateResponse, <-chan error) {
	responseChan := make(chan *AddWaterMarkTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddWaterMarkTemplate(request)
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

// invoke AddWaterMarkTemplate api with *AddWaterMarkTemplateRequest asynchronously
// api document: https://help.aliyun.com/api/mts/addwatermarktemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AddWaterMarkTemplateWithCallback(request *AddWaterMarkTemplateRequest, callback func(response *AddWaterMarkTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddWaterMarkTemplateResponse
		var err error
		defer close(result)
		response, err = client.AddWaterMarkTemplate(request)
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

type AddWaterMarkTemplateRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Name                 string           `position:"Query" name:"Name"`
	Config               string           `position:"Query" name:"Config"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type AddWaterMarkTemplateResponse struct {
	*responses.BaseResponse
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	WaterMarkTemplate WaterMarkTemplate `json:"WaterMarkTemplate" xml:"WaterMarkTemplate"`
}

// create a request to invoke AddWaterMarkTemplate API
func CreateAddWaterMarkTemplateRequest() (request *AddWaterMarkTemplateRequest) {
	request = &AddWaterMarkTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "AddWaterMarkTemplate", "mts", "openAPI")
	return
}

// create a response to parse from AddWaterMarkTemplate response
func CreateAddWaterMarkTemplateResponse() (response *AddWaterMarkTemplateResponse) {
	response = &AddWaterMarkTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
