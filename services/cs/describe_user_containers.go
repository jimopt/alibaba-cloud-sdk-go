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

// invoke DescribeUserContainers api with *DescribeUserContainersRequest synchronously
// api document: https://help.aliyun.com/api/cs/describeusercontainers.html
func (client *Client) DescribeUserContainers(request *DescribeUserContainersRequest) (response *DescribeUserContainersResponse, err error) {
	response = CreateDescribeUserContainersResponse()
	err = client.DoAction(request, response)
	return
}

// invoke DescribeUserContainers api with *DescribeUserContainersRequest asynchronously
// api document: https://help.aliyun.com/api/cs/describeusercontainers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUserContainersWithChan(request *DescribeUserContainersRequest) (<-chan *DescribeUserContainersResponse, <-chan error) {
	responseChan := make(chan *DescribeUserContainersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUserContainers(request)
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

// invoke DescribeUserContainers api with *DescribeUserContainersRequest asynchronously
// api document: https://help.aliyun.com/api/cs/describeusercontainers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeUserContainersWithCallback(request *DescribeUserContainersRequest, callback func(response *DescribeUserContainersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUserContainersResponse
		var err error
		defer close(result)
		response, err = client.DescribeUserContainers(request)
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

type DescribeUserContainersRequest struct {
	*requests.RoaRequest
	ServiceId string `position:"Query" name:"ServiceId"`
}

type DescribeUserContainersResponse struct {
	*responses.BaseResponse
}

// create a request to invoke DescribeUserContainers API
func CreateDescribeUserContainersRequest() (request *DescribeUserContainersRequest) {
	request = &DescribeUserContainersRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("CS", "2015-12-15", "DescribeUserContainers", "/region/[RegionId]/containers", "", "")
	request.Method = requests.GET
	return
}

// create a response to parse from DescribeUserContainers response
func CreateDescribeUserContainersResponse() (response *DescribeUserContainersResponse) {
	response = &DescribeUserContainersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
