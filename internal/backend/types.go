/*
 *  Copyright (c) 2022, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package backend

type Backend struct {
	Name string `json:"name"`
	Spec Spec   `json:"spec"`
}

type Spec struct {
	CertificateName string          `json:"certificateName"`
	Http2Enabled    bool            `json:"http2Enabled"`
	Timeout         int             `json:"timeout"`
	Credentials     Credentials     `json:"credentials"`
	RetryConfig     RetryConfig     `json:"retryConfig"`
	CircuitBreakers CircuitBreakers `json:"circuitBreakers"`
}

type Credentials struct {
	Type   string
	Secret string
}

type RetryConfig struct {
	Count      int
	StatusCode int
}

type CircuitBreakers struct {
	MaxConnections     int
	MaxRequests        int
	MaxPendingRequests int
	MaxRetries         int
	MaxConnectionPools int
}
