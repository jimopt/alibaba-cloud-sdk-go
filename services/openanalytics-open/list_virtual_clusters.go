package openanalytics_open

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

// ListVirtualClusters invokes the openanalytics_open.ListVirtualClusters API synchronously
func (client *Client) ListVirtualClusters(request *ListVirtualClustersRequest) (response *ListVirtualClustersResponse, err error) {
	response = CreateListVirtualClustersResponse()
	err = client.DoAction(request, response)
	return
}

// ListVirtualClustersWithChan invokes the openanalytics_open.ListVirtualClusters API asynchronously
func (client *Client) ListVirtualClustersWithChan(request *ListVirtualClustersRequest) (<-chan *ListVirtualClustersResponse, <-chan error) {
	responseChan := make(chan *ListVirtualClustersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListVirtualClusters(request)
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

// ListVirtualClustersWithCallback invokes the openanalytics_open.ListVirtualClusters API asynchronously
func (client *Client) ListVirtualClustersWithCallback(request *ListVirtualClustersRequest, callback func(response *ListVirtualClustersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListVirtualClustersResponse
		var err error
		defer close(result)
		response, err = client.ListVirtualClusters(request)
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

// ListVirtualClustersRequest is the request struct for api ListVirtualClusters
type ListVirtualClustersRequest struct {
	*requests.RpcRequest
	Type string `position:"Body" name:"Type"`
}

// ListVirtualClustersResponse is the response struct for api ListVirtualClusters
type ListVirtualClustersResponse struct {
	*responses.BaseResponse
	RequestId string                          `json:"RequestId" xml:"RequestId"`
	Data      []DataItemInListVirtualClusters `json:"Data" xml:"Data"`
}

// CreateListVirtualClustersRequest creates a request to invoke ListVirtualClusters API
func CreateListVirtualClustersRequest() (request *ListVirtualClustersRequest) {
	request = &ListVirtualClustersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("openanalytics-open", "2018-06-19", "ListVirtualClusters", "openanalytics", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListVirtualClustersResponse creates a response to parse from ListVirtualClusters response
func CreateListVirtualClustersResponse() (response *ListVirtualClustersResponse) {
	response = &ListVirtualClustersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
