package madmin

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	// SupportMetricsVersion1 is version 1
	SupportMetricsVersion1 = "1"
	// HealthInfoVersion is current health info version.
	SupportMetricsVersion = SupportMetricsVersion1
)

// SupportMetricsVersionStruct - struct for health info version
type SupportMetricsVersionStruct struct {
	Version string `json:"version,omitempty"`
	Error   string `json:"error,omitempty"`
}

func (adm *AdminClient) SupportMetrics(ctx context.Context, client *AdminClient) ([]byte, string, error) {
	resp, err := adm.executeMethod(
		ctx, "GET", requestData{
			relPath: adminAPIPrefix + "/supportmetrics",
			// queryValues: v,
		},
	)
	if err != nil {
		closeResponse(resp)
		return nil, "", err
	}

	if resp.StatusCode != http.StatusOK {
		closeResponse(resp)
		return nil, "", httpRespToErrorResponse(resp)
	}

	d, e := io.ReadAll(resp.Body)
	if e != nil {
		closeResponse(resp)
		return nil, "", err
	}

	var version SupportMetricsVersionStruct
	e = json.Unmarshal(d, &version)
	if e != nil {
		closeResponse(resp)
		return nil, "", err
	}

	if version.Error != "" {
		fmt.Println("Closing response because:", version.Error)
		closeResponse(resp)
		return nil, "", errors.New(version.Error)
	}

	switch version.Version {
	case SupportMetricsVersion:
		return d, version.Version, err
	default:
		closeResponse(resp)
		return nil, "", errors.New("Upgrade Minio Client to support metrics version " + version.Version)
	}
}
