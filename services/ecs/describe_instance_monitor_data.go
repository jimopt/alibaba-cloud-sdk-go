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

// invoke DescribeInstanceMonitorData api with *DescribeInstanceMonitorDataRequest synchronously
// api document: https://help.aliyun.com/api/ecs/describeinstancemonitordata.html
func (client *Client) DescribeInstanceMonitorData(request *DescribeInstanceMonitorDataRequest) (response *DescribeInstanceMonitorDataResponse, err error) {
	response = CreateDescribeInstanceMonitorDataResponse()
	err = client.DoAction(request, response)
	return
}

// invoke DescribeInstanceMonitorData api with *DescribeInstanceMonitorDataRequest asynchronously
// api document: https://help.aliyun.com/api/ecs/describeinstancemonitordata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceMonitorDataWithChan(request *DescribeInstanceMonitorDataRequest) (<-chan *DescribeInstanceMonitorDataResponse, <-chan error) {
	responseChan := make(chan *DescribeInstanceMonitorDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeInstanceMonitorData(request)
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

// invoke DescribeInstanceMonitorData api with *DescribeInstanceMonitorDataRequest asynchronously
// api document: https://help.aliyun.com/api/ecs/describeinstancemonitordata.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeInstanceMonitorDataWithCallback(request *DescribeInstanceMonitorDataRequest, callback func(response *DescribeInstanceMonitorDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeInstanceMonitorDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeInstanceMonitorData(request)
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

type DescribeInstanceMonitorDataRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	EndTime              string           `position:"Query" name:"EndTime"`
	Period               requests.Integer `position:"Query" name:"Period"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type DescribeInstanceMonitorDataResponse struct {
	*responses.BaseResponse
	RequestId   string                                   `json:"RequestId" xml:"RequestId"`
	MonitorData MonitorDataInDescribeInstanceMonitorData `json:"MonitorData" xml:"MonitorData"`
}

// create a request to invoke DescribeInstanceMonitorData API
func CreateDescribeInstanceMonitorDataRequest() (request *DescribeInstanceMonitorDataRequest) {
	request = &DescribeInstanceMonitorDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ecs", "2014-05-26", "DescribeInstanceMonitorData", "ecs", "openAPI")
	return
}

// create a response to parse from DescribeInstanceMonitorData response
func CreateDescribeInstanceMonitorDataResponse() (response *DescribeInstanceMonitorDataResponse) {
	response = &DescribeInstanceMonitorDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
