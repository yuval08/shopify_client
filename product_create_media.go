package shopify_client

import (
	"context"
	"errors"
	"github.com/yuval08/shopify_graphql"
	smodels "github.com/yuval08/shopify_client/models"
)

func (s *Shopify) ProductCreateMedia(productID int, input []smodels.CreateMediaInput) (bool, error) {
	request := shopify_graphql.NewRequest(`
mutation productCreateMedia($productId: ID!, $media: [CreateMediaInput!]!) {
  productCreateMedia(productId: $productId, media: $media) {
    media {
      alt
    }
    mediaUserErrors {
      code
      field
      message
    }
    product {
      id
    }
  }
}
	   `)
	request.Header.Set("X-Shopify-Access-Token", s.Token)
	request.Var("productId", smodels.GID(smodels.ResourceTypeProduct, productID))
	request.Var("media", input)
	response := &smodels.ProductCreateMediaResponse{}
	if err, _ := s.client.Run(context.Background(), request, response); err != nil {
		return false, err
	}

	if len(response.MediaUserErrors) == 0 {
		return true, nil
	}
	return false, errors.New(response.MediaUserErrors[0].Message)
}
