package cs

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

// invoke DescribeTaskInfo api with *DescribeTaskInfoRequest synchronously
// api document: https://help.aliyun.com/api/cs/describetaskinfo.html
func (client *Client) DescribeTaskInfo(request *DescribeTaskInfoRequest) (response *DescribeTaskInfoResponse, err error) {
	response = CreateDescribeTaskInfoResponse()
	err = client.DoAction(request, response)
	return
}

// invoke DescribeTaskInfo api with *DescribeTaskInfoRequest asynchronously
// api document: https://help.aliyun.com/api/cs/describetaskinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTaskInfoWithChan(request *DescribeTaskInfoRequest) (<-chan *DescribeTaskInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeTaskInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTaskInfo(request)
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

// invoke DescribeTaskInfo api with *DescribeTaskInfoRequest asynchronously
// api document: https://help.aliyun.com/api/cs/describetaskinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTaskInfoWithCallback(request *DescribeTaskInfoRequest, callback func(response *DescribeTaskInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTaskInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeTaskInfo(request)
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

type DescribeTaskInfoRequest struct {
	*requests.RoaRequest
	TaskId string `position:"Path" name:"TaskId"`
}

type DescribeTaskInfoResponse struct {
	*responses.BaseResponse
}

// create a request to invoke DescribeTaskInfo API
func CreateDescribeTaskInfoRequest() (request *DescribeTaskInfoRequest) {
	request = &DescribeTaskInfoRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "DescribeTaskInfo", "/tasks/[TaskId]", "", "")
	request.Method = requests.GET
	return
}

// create a response to parse from DescribeTaskInfo response
func CreateDescribeTaskInfoResponse() (response *DescribeTaskInfoResponse) {
	response = &DescribeTaskInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
