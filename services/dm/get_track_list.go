package dm

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

// invoke GetTrackList api with *GetTrackListRequest synchronously
// api document: https://help.aliyun.com/api/dm/gettracklist.html
func (client *Client) GetTrackList(request *GetTrackListRequest) (response *GetTrackListResponse, err error) {
	response = CreateGetTrackListResponse()
	err = client.DoAction(request, response)
	return
}

// invoke GetTrackList api with *GetTrackListRequest asynchronously
// api document: https://help.aliyun.com/api/dm/gettracklist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetTrackListWithChan(request *GetTrackListRequest) (<-chan *GetTrackListResponse, <-chan error) {
	responseChan := make(chan *GetTrackListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTrackList(request)
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

// invoke GetTrackList api with *GetTrackListRequest asynchronously
// api document: https://help.aliyun.com/api/dm/gettracklist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetTrackListWithCallback(request *GetTrackListRequest, callback func(response *GetTrackListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTrackListResponse
		var err error
		defer close(result)
		response, err = client.GetTrackList(request)
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

type GetTrackListRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	EndTime              string           `position:"Query" name:"EndTime"`
	Total                string           `position:"Query" name:"Total"`
	Offset               string           `position:"Query" name:"Offset"`
	PageSize             string           `position:"Query" name:"PageSize"`
	OffsetCreateTime     string           `position:"Query" name:"OffsetCreateTime"`
	OffsetCreateTimeDesc string           `position:"Query" name:"OffsetCreateTimeDesc"`
	PageNumber           string           `position:"Query" name:"PageNumber"`
}

type GetTrackListResponse struct {
	*responses.BaseResponse
	RequestId            string             `json:"RequestId" xml:"RequestId"`
	Total                int                `json:"Total" xml:"Total"`
	PageNo               int                `json:"PageNo" xml:"PageNo"`
	PageSize             int                `json:"PageSize" xml:"PageSize"`
	OffsetCreateTime     string             `json:"OffsetCreateTime" xml:"OffsetCreateTime"`
	OffsetCreateTimeDesc string             `json:"OffsetCreateTimeDesc" xml:"OffsetCreateTimeDesc"`
	Data                 DataInGetTrackList `json:"data" xml:"data"`
}

// create a request to invoke GetTrackList API
func CreateGetTrackListRequest() (request *GetTrackListRequest) {
	request = &GetTrackListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dm", "2015-11-23", "GetTrackList", "", "")
	return
}

// create a response to parse from GetTrackList response
func CreateGetTrackListResponse() (response *GetTrackListResponse) {
	response = &GetTrackListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
