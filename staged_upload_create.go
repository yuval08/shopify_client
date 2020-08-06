package shopify_client

import (
	"context"
	"errors"
	smodels "github.com/yuval08/shopify_client/models"
	"github.com/yuval08/shopify_graphql"
)

func (s *Shopify) StagedUploadCreate(input *smodels.StageUploadInput) ([]smodels.StagedTarget, error) {
	request := shopify_graphql.NewRequest(`
mutation stagedUploadsCreate($input: [StagedUploadInput!]!) {
  stagedUploadsCreate(input: $input) {
    stagedTargets {
      resourceUrl
      url
      parameters {
        name
        value
      }
    }
    userErrors {
      field
      message
    }
  }
}


	   `)
	request.Header.Set("X-Shopify-Access-Token", s.Token)
	request.Var("input", input)
	response := &smodels.StagedUploadCreateResponse{}
	err, _ := s.client.Run(context.Background(), request, response)
	if err != nil {
		return nil, err
	}

	if len(response.StagedUploadCreate.UserErrors) == 0 {
		return response.StagedUploadCreate.Parameters, nil
	}
	return nil, errors.New(response.StagedUploadCreate.UserErrors[0].Message)
}
