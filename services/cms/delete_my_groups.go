package cms

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

// invoke DeleteMyGroups api with *DeleteMyGroupsRequest synchronously
// api document: https://help.aliyun.com/api/cms/deletemygroups.html
func (client *Client) DeleteMyGroups(request *DeleteMyGroupsRequest) (response *DeleteMyGroupsResponse, err error) {
	response = CreateDeleteMyGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// invoke DeleteMyGroups api with *DeleteMyGroupsRequest asynchronously
// api document: https://help.aliyun.com/api/cms/deletemygroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteMyGroupsWithChan(request *DeleteMyGroupsRequest) (<-chan *DeleteMyGroupsResponse, <-chan error) {
	responseChan := make(chan *DeleteMyGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteMyGroups(request)
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

// invoke DeleteMyGroups api with *DeleteMyGroupsRequest asynchronously
// api document: https://help.aliyun.com/api/cms/deletemygroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteMyGroupsWithCallback(request *DeleteMyGroupsRequest, callback func(response *DeleteMyGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteMyGroupsResponse
		var err error
		defer close(result)
		response, err = client.DeleteMyGroups(request)
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

type DeleteMyGroupsRequest struct {
	*requests.RpcRequest
	GroupId requests.Integer `position:"Query" name:"GroupId"`
}

type DeleteMyGroupsResponse struct {
	*responses.BaseResponse
	RequestId    string                `json:"RequestId" xml:"RequestId"`
	Success      bool                  `json:"Success" xml:"Success"`
	ErrorCode    int                   `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string                `json:"ErrorMessage" xml:"ErrorMessage"`
	Group        GroupInDeleteMyGroups `json:"Group" xml:"Group"`
}

// create a request to invoke DeleteMyGroups API
func CreateDeleteMyGroupsRequest() (request *DeleteMyGroupsRequest) {
	request = &DeleteMyGroupsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2017-03-01", "DeleteMyGroups", "cms", "openAPI")
	return
}

// create a response to parse from DeleteMyGroups response
func CreateDeleteMyGroupsResponse() (response *DeleteMyGroupsResponse) {
	response = &DeleteMyGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
