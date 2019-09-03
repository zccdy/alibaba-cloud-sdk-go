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

// ModifyCollationTimeZone invokes the rds.ModifyCollationTimeZone API synchronously
// api document: https://help.aliyun.com/api/rds/modifycollationtimezone.html
func (client *Client) ModifyCollationTimeZone(request *ModifyCollationTimeZoneRequest) (response *ModifyCollationTimeZoneResponse, err error) {
	response = CreateModifyCollationTimeZoneResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyCollationTimeZoneWithChan invokes the rds.ModifyCollationTimeZone API asynchronously
// api document: https://help.aliyun.com/api/rds/modifycollationtimezone.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyCollationTimeZoneWithChan(request *ModifyCollationTimeZoneRequest) (<-chan *ModifyCollationTimeZoneResponse, <-chan error) {
	responseChan := make(chan *ModifyCollationTimeZoneResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyCollationTimeZone(request)
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

// ModifyCollationTimeZoneWithCallback invokes the rds.ModifyCollationTimeZone API asynchronously
// api document: https://help.aliyun.com/api/rds/modifycollationtimezone.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyCollationTimeZoneWithCallback(request *ModifyCollationTimeZoneRequest, callback func(response *ModifyCollationTimeZoneResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyCollationTimeZoneResponse
		var err error
		defer close(result)
		response, err = client.ModifyCollationTimeZone(request)
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

// ModifyCollationTimeZoneRequest is the request struct for api ModifyCollationTimeZone
type ModifyCollationTimeZoneRequest struct {
	*requests.RpcRequest
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	Collation            string           `position:"Query" name:"Collation"`
	Timezone             string           `position:"Query" name:"Timezone"`
}

// ModifyCollationTimeZoneResponse is the response struct for api ModifyCollationTimeZone
type ModifyCollationTimeZoneResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	DBInstanceId string `json:"DBInstanceId" xml:"DBInstanceId"`
	TaskId       string `json:"TaskId" xml:"TaskId"`
	Timezone     string `json:"Timezone" xml:"Timezone"`
	Collation    string `json:"Collation" xml:"Collation"`
}

// CreateModifyCollationTimeZoneRequest creates a request to invoke ModifyCollationTimeZone API
func CreateModifyCollationTimeZoneRequest() (request *ModifyCollationTimeZoneRequest) {
	request = &ModifyCollationTimeZoneRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ModifyCollationTimeZone", "rds", "openAPI")
	return
}

// CreateModifyCollationTimeZoneResponse creates a response to parse from ModifyCollationTimeZone response
func CreateModifyCollationTimeZoneResponse() (response *ModifyCollationTimeZoneResponse) {
	response = &ModifyCollationTimeZoneResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
