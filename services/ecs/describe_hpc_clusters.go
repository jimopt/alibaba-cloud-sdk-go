package ecs

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

// invoke DescribeHpcClusters api with *DescribeHpcClustersRequest synchronously
// api document: https://help.aliyun.com/api/ecs/describehpcclusters.html
func (client *Client) DescribeHpcClusters(request *DescribeHpcClustersRequest) (response *DescribeHpcClustersResponse, err error) {
	response = CreateDescribeHpcClustersResponse()
	err = client.DoAction(request, response)
	return
}

// invoke DescribeHpcClusters api with *DescribeHpcClustersRequest asynchronously
// api document: https://help.aliyun.com/api/ecs/describehpcclusters.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeHpcClustersWithChan(request *DescribeHpcClustersRequest) (<-chan *DescribeHpcClustersResponse, <-chan error) {
	responseChan := make(chan *DescribeHpcClustersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeHpcClusters(request)
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

// invoke DescribeHpcClusters api with *DescribeHpcClustersRequest asynchronously
// api document: https://help.aliyun.com/api/ecs/describehpcclusters.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeHpcClustersWithCallback(request *DescribeHpcClustersRequest, callback func(response *DescribeHpcClustersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeHpcClustersResponse
		var err error
		defer close(result)
		response, err = client.DescribeHpcClusters(request)
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

type DescribeHpcClustersRequest struct {
	*requests.RpcRequest
}

type DescribeHpcClustersResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	TotalCount  int         `json:"TotalCount" xml:"TotalCount"`
	PageNumber  int         `json:"PageNumber" xml:"PageNumber"`
	PageSize    int         `json:"PageSize" xml:"PageSize"`
	HpcClusters HpcClusters `json:"HpcClusters" xml:"HpcClusters"`
}

// create a request to invoke DescribeHpcClusters API
func CreateDescribeHpcClustersRequest() (request *DescribeHpcClustersRequest) {
	request = &DescribeHpcClustersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeHpcClusters", "ecs", "openAPI")
	return
}

// create a response to parse from DescribeHpcClusters response
func CreateDescribeHpcClustersResponse() (response *DescribeHpcClustersResponse) {
	response = &DescribeHpcClustersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
