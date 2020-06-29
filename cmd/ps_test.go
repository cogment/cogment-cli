package cmd

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"gitlab.com/cogment/cogment/api"
	"net/http"
	"testing"
)

func TestPsCommandWithStatusCode(t *testing.T) {

	const name = "my application"
	const id = "my-app-731841"
	const createdAt = 1570558016

	client := resty.New()
	//client.SetDebug(false)

	httpmock.ActivateNonDefault(client.GetClient())
	defer httpmock.DeactivateAndReset()

	var tests = []struct {
		statusCode           int
		expectedApplications []*api.Application
		hasErr               bool
	}{
		{200, []*api.Application{
			{Id: id, Name: name, CreatedAt: createdAt},
		}, false},
		{400, nil, true},
		{401, nil, true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.statusCode), func(t *testing.T) {

			httpmock.Reset()
			httpmock.RegisterResponder("GET", `/applications`,
				func(req *http.Request) (*http.Response, error) {
					return httpmock.NewJsonResponse(tt.statusCode, []map[string]interface{}{
						{
							"id":         id,
							"name":       name,
							"created_at": createdAt,
						},
					})
				},
			)

			applications, err := runPsCmd(client)

			assert.Equal(t, 1, httpmock.GetTotalCallCount())
			assert.Equal(t, tt.expectedApplications, applications)

			if tt.hasErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

		})
	}
}
