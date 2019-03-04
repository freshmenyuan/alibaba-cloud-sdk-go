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

// CreateChannelToken invokes the rtc.CreateChannelToken API synchronously
// api document: https://help.aliyun.com/api/rtc/createchanneltoken.html
func (client *Client) CreateChannelToken(request *CreateChannelTokenRequest) (response *CreateChannelTokenResponse, err error) {
	response = CreateCreateChannelTokenResponse()
	err = client.DoAction(request, response)
	return
}

// CreateChannelTokenWithChan invokes the rtc.CreateChannelToken API asynchronously
// api document: https://help.aliyun.com/api/rtc/createchanneltoken.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateChannelTokenWithChan(request *CreateChannelTokenRequest) (<-chan *CreateChannelTokenResponse, <-chan error) {
	responseChan := make(chan *CreateChannelTokenResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateChannelToken(request)
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

// CreateChannelTokenWithCallback invokes the rtc.CreateChannelToken API asynchronously
// api document: https://help.aliyun.com/api/rtc/createchanneltoken.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateChannelTokenWithCallback(request *CreateChannelTokenRequest, callback func(response *CreateChannelTokenResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateChannelTokenResponse
		var err error
		defer close(result)
		response, err = client.CreateChannelToken(request)
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

// CreateChannelTokenRequest is the request struct for api CreateChannelToken
type CreateChannelTokenRequest struct {
	*requests.RpcRequest
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
	AppId     string           `position:"Query" name:"AppId"`
	ChannelId string           `position:"Query" name:"ChannelId"`
	SessionId string           `position:"Query" name:"SessionId"`
	UId       string           `position:"Query" name:"UId"`
	Nonce     string           `position:"Query" name:"Nonce"`
}

// CreateChannelTokenResponse is the response struct for api CreateChannelToken
type CreateChannelTokenResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ChannelToken string `json:"ChannelToken" xml:"ChannelToken"`
}

// CreateCreateChannelTokenRequest creates a request to invoke CreateChannelToken API
func CreateCreateChannelTokenRequest() (request *CreateChannelTokenRequest) {
	request = &CreateChannelTokenRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "CreateChannelToken", "rtc", "openAPI")
	return
}

// CreateCreateChannelTokenResponse creates a response to parse from CreateChannelToken response
func CreateCreateChannelTokenResponse() (response *CreateChannelTokenResponse) {
	response = &CreateChannelTokenResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
