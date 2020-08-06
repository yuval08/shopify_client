package shopify_client

import (
	"context"
	"errors"
	smodels "github.com/yuval08/shopify_client/models"
	"github.com/yuval08/shopify_graphql"
)

func (s *Shopify) QueryProduct(id string) (*smodels.Product, error) {
	request := shopify_graphql.NewRequest(`
	query($id: ID!) {
	   product(id:$id) {
			id
			metafields(first: 100) {
			  edges {
				node {
				  id
				  key
				}
			  }
			}
		  }
	}
`)
	request.Header.Set("X-Shopify-Access-Token", s.Token)
	request.Var("id", id)
	response := &smodels.ProductResponse{}
	err, _ := s.client.Run(context.Background(), request, response)
	if err != nil {
		return nil, err
	}

	return &response.Product, nil
}

func (s *Shopify) CreateProduct(input *smodels.ProductInput) (*smodels.Product, error) {
	request := shopify_graphql.NewRequest(`
	mutation($input: ProductInput!) {
	 productCreate(input: $input) {
		product {
			id
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
	response := &smodels.ProductCreateResponse{}
	err, _ := s.client.Run(context.Background(), request, response)
	if err != nil {
		return nil, err
	}

	if len(response.ProductCreate.UserErrors) == 0 {
		return &response.ProductCreate.Product, nil
	}
	return nil, errors.New(response.ProductCreate.UserErrors[0].DisplayError())
}

func (s *Shopify) UpdateProduct(input *smodels.ProductInput) (*smodels.Product, error) {
	request := shopify_graphql.NewRequest(`
	mutation($input: ProductInput!) {
	 productUpdate(input: $input) {
		product {
			id
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
	response := smodels.ProductUpdateResponse{}
	err, _ := s.client.Run(context.Background(), request, &response)
	if err != nil {
		return nil, err
	}

	if len(response.ProductUpdate.UserErrors) == 0 {
		return &response.ProductUpdate.Product, nil
	}
	return nil, errors.New(response.ProductUpdate.UserErrors[0].Message)
}
