package ddospro

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

// invoke DescribeWebLogDownloadFilePage api with *DescribeWebLogDownloadFilePageRequest synchronously
// api document: https://help.aliyun.com/api/ddospro/describeweblogdownloadfilepage.html
func (client *Client) DescribeWebLogDownloadFilePage(request *DescribeWebLogDownloadFilePageRequest) (response *DescribeWebLogDownloadFilePageResponse, err error) {
	response = CreateDescribeWebLogDownloadFilePageResponse()
	err = client.DoAction(request, response)
	return
}

// invoke DescribeWebLogDownloadFilePage api with *DescribeWebLogDownloadFilePageRequest asynchronously
// api document: https://help.aliyun.com/api/ddospro/describeweblogdownloadfilepage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeWebLogDownloadFilePageWithChan(request *DescribeWebLogDownloadFilePageRequest) (<-chan *DescribeWebLogDownloadFilePageResponse, <-chan error) {
	responseChan := make(chan *DescribeWebLogDownloadFilePageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeWebLogDownloadFilePage(request)
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

// invoke DescribeWebLogDownloadFilePage api with *DescribeWebLogDownloadFilePageRequest asynchronously
// api document: https://help.aliyun.com/api/ddospro/describeweblogdownloadfilepage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeWebLogDownloadFilePageWithCallback(request *DescribeWebLogDownloadFilePageRequest, callback func(response *DescribeWebLogDownloadFilePageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeWebLogDownloadFilePageResponse
		var err error
		defer close(result)
		response, err = client.DescribeWebLogDownloadFilePage(request)
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

type DescribeWebLogDownloadFilePageRequest struct {
	*requests.RpcRequest
	SourceIp        string           `position:"Query" name:"SourceIp"`
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage     requests.Integer `position:"Query" name:"CurrentPage"`
	TaskId          requests.Integer `position:"Query" name:"TaskId"`
}

type DescribeWebLogDownloadFilePageResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	FileList   []File `json:"FileList" xml:"FileList"`
}

// create a request to invoke DescribeWebLogDownloadFilePage API
func CreateDescribeWebLogDownloadFilePageRequest() (request *DescribeWebLogDownloadFilePageRequest) {
	request = &DescribeWebLogDownloadFilePageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DDoSPro", "2017-07-25", "DescribeWebLogDownloadFilePage", "", "")
	return
}

// create a response to parse from DescribeWebLogDownloadFilePage response
func CreateDescribeWebLogDownloadFilePageResponse() (response *DescribeWebLogDownloadFilePageResponse) {
	response = &DescribeWebLogDownloadFilePageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
