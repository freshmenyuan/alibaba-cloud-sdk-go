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

package rtc

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// GetMPUTaskStatus invokes the rtc.GetMPUTaskStatus API synchronously
// api document: https://help.aliyun.com/api/rtc/getmputaskstatus.html
func (client *Client) GetMPUTaskStatus(request *GetMPUTaskStatusRequest) (response *GetMPUTaskStatusResponse, err error) {
	response = CreateGetMPUTaskStatusResponse()
	err = client.DoAction(request, response)
	return
}

// GetMPUTaskStatusWithChan invokes the rtc.GetMPUTaskStatus API asynchronously
// api document: https://help.aliyun.com/api/rtc/getmputaskstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetMPUTaskStatusWithChan(request *GetMPUTaskStatusRequest) (<-chan *GetMPUTaskStatusResponse, <-chan error) {
	responseChan := make(chan *GetMPUTaskStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetMPUTaskStatus(request)
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

// GetMPUTaskStatusWithCallback invokes the rtc.GetMPUTaskStatus API asynchronously
// api document: https://help.aliyun.com/api/rtc/getmputaskstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetMPUTaskStatusWithCallback(request *GetMPUTaskStatusRequest, callback func(response *GetMPUTaskStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetMPUTaskStatusResponse
		var err error
		defer close(result)
		response, err = client.GetMPUTaskStatus(request)
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

// GetMPUTaskStatusRequest is the request struct for api GetMPUTaskStatus
type GetMPUTaskStatusRequest struct {
	*requests.RpcRequest
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
	AppId   string           `position:"Query" name:"AppId"`
	TaskId  string           `position:"Query" name:"TaskId"`
}

// GetMPUTaskStatusResponse is the response struct for api GetMPUTaskStatus
type GetMPUTaskStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Status    int    `json:"Status" xml:"Status"`
}

// CreateGetMPUTaskStatusRequest creates a request to invoke GetMPUTaskStatus API
func CreateGetMPUTaskStatusRequest() (request *GetMPUTaskStatusRequest) {
	request = &GetMPUTaskStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "GetMPUTaskStatus", "rtc", "openAPI")
	return
}

// CreateGetMPUTaskStatusResponse creates a response to parse from GetMPUTaskStatus response
func CreateGetMPUTaskStatusResponse() (response *GetMPUTaskStatusResponse) {
	response = &GetMPUTaskStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
