//
// MinIO Object Storage (c) 2021 MinIO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package madmin

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

// GetSubnetLicense - Connect to a minio server and call Get Subnet License API
// to fetch the subnet license if present in the config
func (adm *AdminClient) GetSubnetLicense(ctx context.Context) (string, error) {
	resp, err := adm.executeMethod(ctx,
		http.MethodGet,
		requestData{relPath: adminAPIPrefix + "/get-subnet-license"},
	)
	defer closeResponse(resp)
	if err != nil {
		return "", err
	}

	// Check response http status code
	if resp.StatusCode != http.StatusOK {
		return "", httpRespToErrorResponse(resp)
	}

	lic, err := DecryptData(adm.getSecretKey(), resp.Body)
	if err != nil {
		return "", err
	}

	parts := strings.Split(string(lic), "subnet license=")
	if len(parts) == 2 {
		return parts[1], nil
	}

	return "", errors.New("Unexpected response format")
}
