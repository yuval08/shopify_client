package models

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

/* ---------------------- */
/*           ID           */
/* ---------------------- */

var shopifyGIDFormat = "gid://shopify/%s/%d"

// ID retrieve id
func ID(gid string) int {
	// Split gid with /
	substrings := strings.Split(gid, "/")
	// Get id string
	idStr := substrings[len(substrings)-1]
	if i := strings.Index(idStr, "?"); i > 0 {
		idStr = idStr[0:i]
	}
	// Convert str id to int
	id, _ := strconv.Atoi(idStr)

	return id
}

type ResourceType string

const (
	ResourceTypeLocation ResourceType = "Location"
	ResourceTypeProduct  ResourceType = "Product"
	ResourceTypeVariant  ResourceType = "Variant"
)

// GID retrieve id
func GID(resource ResourceType, id int) string {
	return fmt.Sprintf(shopifyGIDFormat, resource, id)
}

// EncodeGID base64 encode gid
func EncodeGID(gid string) string {
	return base64.URLEncoding.EncodeToString([]byte(gid))
}

// EncodedGID base64 encoded gid
func EncodedGID(resource ResourceType, id int) string {
	return EncodeGID(GID(resource, id))
}

/* ------------------------ */
/*           Type           */
/* ------------------------ */

// Bool bool
func Bool(v bool) *bool {
	return &v
}

// BoolValue bool value
func BoolValue(v *bool) bool {
	if v != nil {
		return *v
	}
	return false
}

// Float32 decimal
func Float32(v float32) *float32 {
	return &v
}

// Float32Value decimal value
func Float32Value(v *float32) float32 {
	if v != nil {
		return *v
	}
	return 0.0
}
