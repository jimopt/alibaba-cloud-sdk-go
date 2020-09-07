package edas

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

// ListClusterMembers invokes the edas.ListClusterMembers API synchronously
// api document: https://help.aliyun.com/api/edas/listclustermembers.html
func (client *Client) ListClusterMembers(request *ListClusterMembersRequest) (response *ListClusterMembersResponse, err error) {
	response = CreateListClusterMembersResponse()
	err = client.DoAction(request, response)
	return
}

// ListClusterMembersWithChan invokes the edas.ListClusterMembers API asynchronously
// api document: https://help.aliyun.com/api/edas/listclustermembers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListClusterMembersWithChan(request *ListClusterMembersRequest) (<-chan *ListClusterMembersResponse, <-chan error) {
	responseChan := make(chan *ListClusterMembersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListClusterMembers(request)
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

// ListClusterMembersWithCallback invokes the edas.ListClusterMembers API asynchronously
// api document: https://help.aliyun.com/api/edas/listclustermembers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListClusterMembersWithCallback(request *ListClusterMembersRequest, callback func(response *ListClusterMembersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListClusterMembersResponse
		var err error
		defer close(result)
		response, err = client.ListClusterMembers(request)
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

// ListClusterMembersRequest is the request struct for api ListClusterMembers
type ListClusterMembersRequest struct {
	*requests.RoaRequest
	PageSize    requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage requests.Integer `position:"Query" name:"CurrentPage"`
	ClusterId   string           `position:"Query" name:"ClusterId"`
	EcsList     string           `position:"Query" name:"EcsList"`
}

// ListClusterMembersResponse is the response struct for api ListClusterMembers
type ListClusterMembersResponse struct {
	*responses.BaseResponse
	Code              int               `json:"Code" xml:"Code"`
	Message           string            `json:"Message" xml:"Message"`
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	ClusterMemberPage ClusterMemberPage `json:"ClusterMemberPage" xml:"ClusterMemberPage"`
}

// CreateListClusterMembersRequest creates a request to invoke ListClusterMembers API
func CreateListClusterMembersRequest() (request *ListClusterMembersRequest) {
	request = &ListClusterMembersRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "ListClusterMembers", "/pop/v5/resource/cluster_member_list", "Edas", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListClusterMembersResponse creates a response to parse from ListClusterMembers response
func CreateListClusterMembersResponse() (response *ListClusterMembersResponse) {
	response = &ListClusterMembersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
