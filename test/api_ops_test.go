/*
openapi for sibyl2 server

Testing OPSApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"testing"

	openapiclient "github.com/opensibyl/sibyl-go-client"
	"github.com/stretchr/testify/require"
)

func Test_openapi_OPSApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OPSApiService OpsMonitorUploadGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, err := apiClient.OPSApi.OpsMonitorUploadGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)

	})

	t.Run("Test OPSApiService OpsPingGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, err := apiClient.OPSApi.OpsPingGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)

	})

}
