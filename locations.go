package shopify_client

import (
	"context"
	"fmt"
	smodels "github.com/yuval08/shopify_client/models"
	"github.com/yuval08/shopify_graphql"
)

func (s *Shopify) getLocations() (*smodels.Locations, error) {
	request := shopify_graphql.NewRequest(fmt.Sprintf(`
	query { 
  locations(first:20) {
    edges {
      node {
        id
      }
    }
  }
}
	   `))
	request.Header.Set("X-Shopify-Access-Token", s.Token)
	response := &smodels.LocationsResponse{}
	err, _ := s.client.Run(context.Background(), request, &response)
	if err != nil {
		return nil, err
	}

	return &response.Locations, nil
}
