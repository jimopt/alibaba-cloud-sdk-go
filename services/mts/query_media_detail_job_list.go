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

// invoke QueryMediaDetailJobList api with *QueryMediaDetailJobListRequest synchronously
// api document: https://help.aliyun.com/api/mts/querymediadetailjoblist.html
func (client *Client) QueryMediaDetailJobList(request *QueryMediaDetailJobListRequest) (response *QueryMediaDetailJobListResponse, err error) {
	response = CreateQueryMediaDetailJobListResponse()
	err = client.DoAction(request, response)
	return
}

// invoke QueryMediaDetailJobList api with *QueryMediaDetailJobListRequest asynchronously
// api document: https://help.aliyun.com/api/mts/querymediadetailjoblist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryMediaDetailJobListWithChan(request *QueryMediaDetailJobListRequest) (<-chan *QueryMediaDetailJobListResponse, <-chan error) {
	responseChan := make(chan *QueryMediaDetailJobListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryMediaDetailJobList(request)
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

// invoke QueryMediaDetailJobList api with *QueryMediaDetailJobListRequest asynchronously
// api document: https://help.aliyun.com/api/mts/querymediadetailjoblist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryMediaDetailJobListWithCallback(request *QueryMediaDetailJobListRequest, callback func(response *QueryMediaDetailJobListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryMediaDetailJobListResponse
		var err error
		defer close(result)
		response, err = client.QueryMediaDetailJobList(request)
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

type QueryMediaDetailJobListRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JobIds               string           `position:"Query" name:"JobIds"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type QueryMediaDetailJobListResponse struct {
	*responses.BaseResponse
	RequestId   string                               `json:"RequestId" xml:"RequestId"`
	NonExistIds NonExistIdsInQueryMediaDetailJobList `json:"NonExistIds" xml:"NonExistIds"`
	JobList     JobListInQueryMediaDetailJobList     `json:"JobList" xml:"JobList"`
}

// create a request to invoke QueryMediaDetailJobList API
func CreateQueryMediaDetailJobListRequest() (request *QueryMediaDetailJobListRequest) {
	request = &QueryMediaDetailJobListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "QueryMediaDetailJobList", "mts", "openAPI")
	return
}

// create a response to parse from QueryMediaDetailJobList response
func CreateQueryMediaDetailJobListResponse() (response *QueryMediaDetailJobListResponse) {
	response = &QueryMediaDetailJobListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
