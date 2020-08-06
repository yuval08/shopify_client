package models

import "encoding/xml"

type MediaErrorCode string

const (
	//Video failed validation.
	MediaErrorCodeVideoValidationError MediaErrorCode = "VIDEO_VALIDATION_ERROR"

	//Media could not be created because the external video could not be found.
	MediaErrorCodeExternalVideoNOtFound MediaErrorCode = "EXTERNAL_VIDEO_NOT_FOUND"

	//Media could not be created because the external video is not listed or is private.
	MediaErrorCodeExternalVideoUnlisted MediaErrorCode = "EXTERNAL_VIDEO_UNLISTED"

	//Media could not be processed because the image could not be downloaded.
	MediaErrorCodeImageDownloadFailure MediaErrorCode = "IMAGE_DOWNLOAD_FAILURE"

	//Media could not be processed because the image could not be processed.
	MediaErrorCodeImageProcessingFailure MediaErrorCode = "IMAGE_PROCESSING_FAILURE"

	//Media could not be processed because the signed url was invalid.
	MediaErrorCodeInvalidSignedURL MediaErrorCode = "INVALID_SIGNED_URL"

	//Media timed out because it is currently being modified by another operation.
	MediaErrorCodeMediaTimeoutError MediaErrorCode = "MEDIA_TIMEOUT_ERROR"

	//Model failed validation.
	MediaErrorCodeModel3DValidationError MediaErrorCode = "MODEL3D_VALIDATION_ERROR"

	//Media error has occured for unknown reason.
	MediaErrorCodeUnknown MediaErrorCode = "UNKNOWN"

	//Media could not be created because the external video has an invalid aspect ratio.
	MediaErrorCodeExternalVideoInvalidAspectRatio MediaErrorCode = "externalVIDEO_INVALID_ASPECT_RATIO"
)

type MediaContentType string

const (
	//An externally hosted video.
	MediaContentTypeExternalVideo MediaContentType = "EXTERNAL_VIDEO"

	//A Shopify hosted image.
	MediaContentTypeImage MediaContentType = "IMAGE"

	//A 3d model.
	MediaContentTypeModel3D MediaContentType = "MODEL_3D"

	//A Shopify hosted video.
	MediaContentTypeVideo MediaContentType = "VIDEO"
)

type StagedUploadHttpMethodType string

const (
	StagedUploadHttpMethodTypePOST StagedUploadHttpMethodType = "POST"
	StagedUploadHttpMethodTypePUT  StagedUploadHttpMethodType = "PUT"
)

type StagedUploadTargetGenerateUploadResource string

const (
	///A collection image.
	StagedUploadTargetGenerateUploadResourceCollectionImage StagedUploadTargetGenerateUploadResource = "COLLECTION_IMAGE"

	///Merchandising::Image resource representation.
	StagedUploadTargetGenerateUploadResourceImage StagedUploadTargetGenerateUploadResource = "IMAGE"

	///Merchandising::Model3d resource representation.
	StagedUploadTargetGenerateUploadResourceModel3D StagedUploadTargetGenerateUploadResource = "MODEL_3D"

	///A product image.
	StagedUploadTargetGenerateUploadResourceProductImage StagedUploadTargetGenerateUploadResource = "PRODUCT_IMAGE"

	///A shop image.
	StagedUploadTargetGenerateUploadResourceShopImage StagedUploadTargetGenerateUploadResource = "SHOP_IMAGE"

	///A timeline event.
	StagedUploadTargetGenerateUploadResourceTimeline StagedUploadTargetGenerateUploadResource = "TIMELINE"

	///Merchandising::Video resource representation.
	StagedUploadTargetGenerateUploadResourceVideo StagedUploadTargetGenerateUploadResource = "VIDEO"
)

type FileUploadResponse struct {
	XMLName xml.Name `xml:"PostResponse"`
	//Text     string   `xml:",chardata"`
	Location string `xml:"Location"`
	Bucket   string `xml:"Bucket"`
	Key      string `xml:"Key"`
	ETag     string `xml:"ETag"`
}
type MediaPreviewImageStatus string

const (
	//Preview image processing has failed.
	MediaPreviewImageStatusFailed MediaPreviewImageStatus = "FAILED"

	//Preview image is being processed.
	MediaPreviewImageStatusProcessing MediaPreviewImageStatus = "PROCESSING"

	//Preview image is ready to be displayed.
	MediaPreviewImageStatusReady MediaPreviewImageStatus = "READY"

	//Preview image is uploaded but not yet processed.
	MediaPreviewImageStatusUploaded MediaPreviewImageStatus = "UPLOADED"
)

type MediaStatus string

const (
	//Media processing has failed.
	MediaStatusFailed MediaStatus = "FAILED"

	//Media is being processed.
	MediaStatusProcessing MediaStatus = "PROCESSING"

	//Media is ready to be displayed.
	MediaStatusReady MediaStatus = "READY"

	//Media has been uploaded but not yet processed.
	MediaStatusUploaded MediaStatus = "UPLOADED"
)

type MediaUserErrorCode string

const (
	//Video validation failed.
	MediaUserErrorCodeVideoValidationError MediaUserErrorCode = "VIDEO_VALIDATION_ERROR"

	//Media cannot be modified. It is currently being modified by another operation.
	MediaUserErrorCodeMediaCannotBeModified MediaUserErrorCode = "MEDIA_CANNOT_BE_MODIFIED"

	//Media does not exist.
	MediaUserErrorCodeMediaDoesNotExist MediaUserErrorCode = "MEDIA_DOES_NOT_EXIST"

	//Model3d creation throttle was exceeded.
	MediaUserErrorCodeModel3DThrottleExceeded MediaUserErrorCode = "MODEL3D_THROTTLE_EXCEEDED"

	//Model validation failed.
	MediaUserErrorCodeModel3DValidationError MediaUserErrorCode = "MODEL3D_VALIDATION_ERROR"

	//Product does not exist.
	MediaUserErrorCodeProductDoesNotExist MediaUserErrorCode = "PRODUCT_DOES_NOT_EXIST"

	//Exceeded the limit of media per product.
	MediaUserErrorCodeProductMediaLimitExceeded MediaUserErrorCode = "PRODUCT_MEDIA_LIMIT_EXCEEDED"

	//Exceeded the limit of media per shop.
	MediaUserErrorCodeShopMediaLimitExceeded MediaUserErrorCode = "SHOP_MEDIA_LIMIT_EXCEEDED"

	//Video creation throttle was exceeded.
	MediaUserErrorCodeVideoThrottleExceeded MediaUserErrorCode = "VIDEO_THROTTLE_EXCEEDED"

	//Input value is invalid.
	MediaUserErrorCodeInvalid MediaUserErrorCode = "INVALID"
)
