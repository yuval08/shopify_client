package models

type Product struct {
	ID         string            `json:"id"`
	MetaFields ProductMetaFields `json:"metafields"`
}

type ProductMetaFields struct {
	Edges []struct {
		Node struct {
			ID  string `json:"id"`
			Key string `json:"key"`
		} `json:"node"`
	} `json:"edges"`
	fieldsMap map[string]string
}

func (f ProductMetaFields) GetID(key string) string {
	if f.fieldsMap == nil {
		f.fieldsMap = make(map[string]string)
		for _, edge := range f.Edges {
			f.fieldsMap[edge.Node.Key] = edge.Node.ID
		}
	}
	if val, ok := f.fieldsMap[key]; ok {
		return val
	}
	return ""
}

type ProductResponse struct {
	Product Product `json:"product"`
}

type ProductInput struct {
	CollectionsToJoin      []string               `json:"collectionsToJoin,omitempty"`
	CollectionsToLeave     []string               `json:"collectionsToLeave,omitempty"`
	DescriptionHTML        *string                `json:"bodyHtml,omitempty"`
	GiftCardTemplateSuffix string                 `json:"giftCardTemplateSuffix,omitempty"`
	Handle                 string                 `json:"handle,omitempty"`
	ID                     string                 `json:"id,omitempty"`
	Images                 []ImageInput           `json:"images,omitempty"`
	Metafields             []MetafieldInput       `json:"metafields,omitempty"`
	Options                []string               `json:"options,omitempty"`
	Published              *bool                  `json:"published,omitempty""`
	ProductType            *string                `json:"productType,omitempty"`
	RedirectNewHandle      *bool                  `json:"redirectNewHandle,omitempty"`
	SEO                    *SEOInput              `json:"seo,omitempty"`
	Tags                   []string               `json:"tags,omitempty"`
	TemplateSuffix         string                 `json:"templateSuffix,omitempty"`
	Title                  *string                `json:"title,omitempty"`
	Variants               []*ProductVariantInput `json:"variants,omitempty"`
	Vendor                 *string                `json:"vendor,omitempty"`
}

// ProductDeleteInput product delete input struct
type ProductDeleteInput struct {
	ID string `json:"id,omitempty"`
}

type ProductCreateResponse struct {
	ProductCreate struct {
		Product    Product     `json:"product"`
		UserErrors []UserError `json:"userErrors"`
	} `json:"productCreate"`
}
type ProductUpdateResponse struct {
	ProductUpdate struct {
		Product    Product     `json:"product"`
		UserErrors []UserError `json:"userErrors"`
	} `json:"productUpdate"`
}
