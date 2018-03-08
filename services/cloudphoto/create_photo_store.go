package cloudphoto

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

// invoke CreatePhotoStore api with *CreatePhotoStoreRequest synchronously
// api document: https://help.aliyun.com/api/cloudphoto/createphotostore.html
func (client *Client) CreatePhotoStore(request *CreatePhotoStoreRequest) (response *CreatePhotoStoreResponse, err error) {
	response = CreateCreatePhotoStoreResponse()
	err = client.DoAction(request, response)
	return
}

// invoke CreatePhotoStore api with *CreatePhotoStoreRequest asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/createphotostore.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreatePhotoStoreWithChan(request *CreatePhotoStoreRequest) (<-chan *CreatePhotoStoreResponse, <-chan error) {
	responseChan := make(chan *CreatePhotoStoreResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreatePhotoStore(request)
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

// invoke CreatePhotoStore api with *CreatePhotoStoreRequest asynchronously
// api document: https://help.aliyun.com/api/cloudphoto/createphotostore.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreatePhotoStoreWithCallback(request *CreatePhotoStoreRequest, callback func(response *CreatePhotoStoreResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreatePhotoStoreResponse
		var err error
		defer close(result)
		response, err = client.CreatePhotoStore(request)
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

type CreatePhotoStoreRequest struct {
	*requests.RpcRequest
	StoreName    string           `position:"Query" name:"StoreName"`
	DefaultQuota requests.Integer `position:"Query" name:"DefaultQuota"`
	BucketName   string           `position:"Query" name:"BucketName"`
	Remark       string           `position:"Query" name:"Remark"`
}

type CreatePhotoStoreResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Action    string `json:"Action" xml:"Action"`
}

// create a request to invoke CreatePhotoStore API
func CreateCreatePhotoStoreRequest() (request *CreatePhotoStoreRequest) {
	request = &CreatePhotoStoreRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudPhoto", "2017-07-11", "CreatePhotoStore", "cloudphoto", "openAPI")
	return
}

// create a response to parse from CreatePhotoStore response
func CreateCreatePhotoStoreResponse() (response *CreatePhotoStoreResponse) {
	response = &CreatePhotoStoreResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
