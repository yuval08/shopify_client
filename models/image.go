package models

type ImageInput struct {
	AltText string `json:"altText,omitempty"`
	ID      string `json:"id,omitempty"`
	Src     string `json:"src,omitempty"`
}

type Image struct {
	//A word or phrase to share the nature or contents of an image.
	AltText string `json:"altText,omitempty"`

	//A unique identifier for the image.
	ID string `json:"id"`

	//The location of the original image as a URL.
	OriginalSrc string `json:"originalSrc"`
}
