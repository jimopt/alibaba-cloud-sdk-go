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

// invoke DegradeDBInstanceSpec api with *DegradeDBInstanceSpecRequest synchronously
// api document: https://help.aliyun.com/api/rds/degradedbinstancespec.html
func (client *Client) DegradeDBInstanceSpec(request *DegradeDBInstanceSpecRequest) (response *DegradeDBInstanceSpecResponse, err error) {
	response = CreateDegradeDBInstanceSpecResponse()
	err = client.DoAction(request, response)
	return
}

// invoke DegradeDBInstanceSpec api with *DegradeDBInstanceSpecRequest asynchronously
// api document: https://help.aliyun.com/api/rds/degradedbinstancespec.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DegradeDBInstanceSpecWithChan(request *DegradeDBInstanceSpecRequest) (<-chan *DegradeDBInstanceSpecResponse, <-chan error) {
	responseChan := make(chan *DegradeDBInstanceSpecResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DegradeDBInstanceSpec(request)
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

// invoke DegradeDBInstanceSpec api with *DegradeDBInstanceSpecRequest asynchronously
// api document: https://help.aliyun.com/api/rds/degradedbinstancespec.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DegradeDBInstanceSpecWithCallback(request *DegradeDBInstanceSpecRequest, callback func(response *DegradeDBInstanceSpecResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DegradeDBInstanceSpecResponse
		var err error
		defer close(result)
		response, err = client.DegradeDBInstanceSpec(request)
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

type DegradeDBInstanceSpecRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	DBInstanceClass      string           `position:"Query" name:"DBInstanceClass"`
	DBInstanceStorage    requests.Integer `position:"Query" name:"DBInstanceStorage"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
}

type DegradeDBInstanceSpecResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// create a request to invoke DegradeDBInstanceSpec API
func CreateDegradeDBInstanceSpecRequest() (request *DegradeDBInstanceSpecRequest) {
	request = &DegradeDBInstanceSpecRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DegradeDBInstanceSpec", "rds", "openAPI")
	return
}

// create a response to parse from DegradeDBInstanceSpec response
func CreateDegradeDBInstanceSpecResponse() (response *DegradeDBInstanceSpecResponse) {
	response = &DegradeDBInstanceSpecResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
