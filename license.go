//
// Copyright (c) 2015-2024 MinIO, Inc.
//
// This file is part of MinIO Object Storage stack
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.
//

package madmin

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/miniohq/license/go/license"
)

// GetLicenseInfo - returns the license info
func (adm *AdminClient) GetLicenseInfo(ctx context.Context) (*license.License, error) {
	// Execute GET on /minio/admin/v3/licenseinfo to get license info.
	resp, err := adm.executeMethod(ctx,
		http.MethodGet,
		requestData{
			relPath: adminAPIPrefix + "/license-info",
		})
	defer closeResponse(resp)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, httpRespToErrorResponse(resp)
	}

	l := license.License{}
	err = json.NewDecoder(resp.Body).Decode(&l)
	if err != nil {
		return nil, err
	}
	return &l, nil
}
