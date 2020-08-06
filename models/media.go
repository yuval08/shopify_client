package models

type CreateMediaInput struct {
	Alt              string           `json:"alt,omitempty"`
	MediaContentType MediaContentType `json:"mediaContentType"`
	OriginalSource   string           `json:"originalSource"`
}

type ProductCreateMediaResponse struct {
	//The newly created media.
	Media []*Media `json:"media"`

	//List of errors that occurred executing the mutation.
	MediaUserErrors []MediaUserError `json:"mediaUserErrors"`

	Product Product `json:"product"`
}

type Media struct {
	//A word or phrase to share the nature or contents of a media.
	Alt string `json:"alt,omitempty"`

	//The media content type.
	MediaContentType MediaContentType `json:"mediaContentType"`

	//Any errors which have occurred on the media.
	MediaErrors []MediaError `json:"mediaErrors"`

	//The preview image for the media.
	Preview MediaPreviewImage `json:"preview"`

	//Current status of the media.
	Status MediaStatus `json:"status"`
}

type MediaUserError struct {
	//Error code to uniquely identify the error.
	Code MediaUserErrorCode `json:"code"`

	//Path to the input field which caused the error.
	Field []string `json:"field"`

	//The error message.
	Message string `json:"message"`
}

type MediaError struct {
	//Code representing the type of error.
	Code MediaErrorCode `json:"code"`

	//Additional details regarding the error.
	Details string `json:"details"`

	//Translated error message.
	Message string `json:"message"`
}

type MediaPreviewImage struct {
	//The preview image for the media.
	Image Image `json:"Image"`

	//Current status of the preview image.
	Status MediaPreviewImageStatus `json:"status"`
}
