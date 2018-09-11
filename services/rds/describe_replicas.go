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

// DescribeReplicas invokes the rds.DescribeReplicas API synchronously
// api document: https://help.aliyun.com/api/rds/describereplicas.html
func (client *Client) DescribeReplicas(request *DescribeReplicasRequest) (response *DescribeReplicasResponse, err error) {
	response = CreateDescribeReplicasResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeReplicasWithChan invokes the rds.DescribeReplicas API asynchronously
// api document: https://help.aliyun.com/api/rds/describereplicas.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeReplicasWithChan(request *DescribeReplicasRequest) (<-chan *DescribeReplicasResponse, <-chan error) {
	responseChan := make(chan *DescribeReplicasResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeReplicas(request)
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

// DescribeReplicasWithCallback invokes the rds.DescribeReplicas API asynchronously
// api document: https://help.aliyun.com/api/rds/describereplicas.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeReplicasWithCallback(request *DescribeReplicasRequest, callback func(response *DescribeReplicasResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeReplicasResponse
		var err error
		defer close(result)
		response, err = client.DescribeReplicas(request)
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

// DescribeReplicasRequest is the request struct for api DescribeReplicas
type DescribeReplicasRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	SecurityToken        string           `position:"Query" name:"SecurityToken"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	AttachDbInstanceData requests.Boolean `position:"Query" name:"AttachDbInstanceData"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	ReplicaId            string           `position:"Query" name:"ReplicaId"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
}

// DescribeReplicasResponse is the response struct for api DescribeReplicas
type DescribeReplicasResponse struct {
	*responses.BaseResponse
	RequestId        string  `json:"RequestId" xml:"RequestId"`
	PageNumber       string  `json:"PageNumber" xml:"PageNumber"`
	TotalRecordCount int     `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageRecordCount  int     `json:"PageRecordCount" xml:"PageRecordCount"`
	Replicas         []Items `json:"Replicas" xml:"Replicas"`
}

// CreateDescribeReplicasRequest creates a request to invoke DescribeReplicas API
func CreateDescribeReplicasRequest() (request *DescribeReplicasRequest) {
	request = &DescribeReplicasRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeReplicas", "rds", "openAPI")
	return
}

// CreateDescribeReplicasResponse creates a response to parse from DescribeReplicas response
func CreateDescribeReplicasResponse() (response *DescribeReplicasResponse) {
	response = &DescribeReplicasResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
