package cdn

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

// invoke DescribeDomainsBySource api with *DescribeDomainsBySourceRequest synchronously
// api document: https://help.aliyun.com/api/cdn/describedomainsbysource.html
func (client *Client) DescribeDomainsBySource(request *DescribeDomainsBySourceRequest) (response *DescribeDomainsBySourceResponse, err error) {
	response = CreateDescribeDomainsBySourceResponse()
	err = client.DoAction(request, response)
	return
}

// invoke DescribeDomainsBySource api with *DescribeDomainsBySourceRequest asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainsbysource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainsBySourceWithChan(request *DescribeDomainsBySourceRequest) (<-chan *DescribeDomainsBySourceResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainsBySourceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainsBySource(request)
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

// invoke DescribeDomainsBySource api with *DescribeDomainsBySourceRequest asynchronously
// api document: https://help.aliyun.com/api/cdn/describedomainsbysource.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDomainsBySourceWithCallback(request *DescribeDomainsBySourceRequest, callback func(response *DescribeDomainsBySourceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainsBySourceResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainsBySource(request)
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

type DescribeDomainsBySourceRequest struct {
	*requests.RpcRequest
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
	Sources       string           `position:"Query" name:"Sources"`
}

type DescribeDomainsBySourceResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	Sources     string      `json:"Sources" xml:"Sources"`
	DomainsList DomainsList `json:"DomainsList" xml:"DomainsList"`
}

// create a request to invoke DescribeDomainsBySource API
func CreateDescribeDomainsBySourceRequest() (request *DescribeDomainsBySourceRequest) {
	request = &DescribeDomainsBySourceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainsBySource", "", "")
	return
}

// create a response to parse from DescribeDomainsBySource response
func CreateDescribeDomainsBySourceResponse() (response *DescribeDomainsBySourceResponse) {
	response = &DescribeDomainsBySourceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
