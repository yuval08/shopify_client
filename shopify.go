package shopify_client

import (
	"fmt"
	smodels "github.com/yuval08/shopify_client/models"
	"github.com/yuval08/shopify_graphql"
	"log"
)

type Shopify struct {
	client       *shopify_graphql.Client
	Locations    *smodels.Locations
	MainLocation int
	Token        string
	URL          string
}

func (s *Shopify) Initialize() {
	s.client = shopify_graphql.NewClient(fmt.Sprintf("https://%s/admin/api/2020-07/graphql.json", s.URL))

	var err error
	if s.Locations, err = s.getLocations(); err != nil {
		log.Fatalf("error loading locations: %s", err)
	}
	s.MainLocation = smodels.ID(s.Locations.Edges[0].Node.ID)
}
