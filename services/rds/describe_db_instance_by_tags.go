package rds

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

// invoke DescribeDBInstanceByTags api with *DescribeDBInstanceByTagsRequest synchronously
// api document: https://help.aliyun.com/api/rds/describedbinstancebytags.html
func (client *Client) DescribeDBInstanceByTags(request *DescribeDBInstanceByTagsRequest) (response *DescribeDBInstanceByTagsResponse, err error) {
	response = CreateDescribeDBInstanceByTagsResponse()
	err = client.DoAction(request, response)
	return
}

// invoke DescribeDBInstanceByTags api with *DescribeDBInstanceByTagsRequest asynchronously
// api document: https://help.aliyun.com/api/rds/describedbinstancebytags.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBInstanceByTagsWithChan(request *DescribeDBInstanceByTagsRequest) (<-chan *DescribeDBInstanceByTagsResponse, <-chan error) {
	responseChan := make(chan *DescribeDBInstanceByTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDBInstanceByTags(request)
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

// invoke DescribeDBInstanceByTags api with *DescribeDBInstanceByTagsRequest asynchronously
// api document: https://help.aliyun.com/api/rds/describedbinstancebytags.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDBInstanceByTagsWithCallback(request *DescribeDBInstanceByTagsRequest, callback func(response *DescribeDBInstanceByTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDBInstanceByTagsResponse
		var err error
		defer close(result)
		response, err = client.DescribeDBInstanceByTags(request)
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

type DescribeDBInstanceByTagsRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	ProxyId              string           `position:"Query" name:"proxyId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type DescribeDBInstanceByTagsResponse struct {
	*responses.BaseResponse
	RequestId        string                          `json:"RequestId" xml:"RequestId"`
	PageNumber       int                             `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  int                             `json:"PageRecordCount" xml:"PageRecordCount"`
	TotalRecordCount int                             `json:"TotalRecordCount" xml:"TotalRecordCount"`
	Items            ItemsInDescribeDBInstanceByTags `json:"Items" xml:"Items"`
}

// create a request to invoke DescribeDBInstanceByTags API
func CreateDescribeDBInstanceByTagsRequest() (request *DescribeDBInstanceByTagsRequest) {
	request = &DescribeDBInstanceByTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDBInstanceByTags", "rds", "openAPI")
	return
}

// create a response to parse from DescribeDBInstanceByTags response
func CreateDescribeDBInstanceByTagsResponse() (response *DescribeDBInstanceByTagsResponse) {
	response = &DescribeDBInstanceByTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
