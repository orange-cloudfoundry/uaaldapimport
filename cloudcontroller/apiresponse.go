package cloudcontroller

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type Metadata struct {
	Guid string
}
type Resource struct {
	Metadata Metadata
}
type ApiResponse struct {
	Resources []Resource
}

func parseResponse(response *http.Response) (guid string, err error) {
	body := response.Body
	defer body.Close()
	apiResp := &ApiResponse{}
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &apiResp)
	if err != nil {
		return
	}
	if len(apiResp.Resources) != 1 {
		err = errors.New("The api response returns more than 1 resource")
		return
	}
	guid = apiResp.Resources[0].Metadata.Guid
	return
}
