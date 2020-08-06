package models

type StageUploadInput struct {
	///The size of the file to upload, in bytes.
	FileSize string `json:"fileSize,omitempty"`
	///The filename of the asset being uploaded.
	FileName string `json:"filename"`
	///The HTTP method to be used by the staged upload.
	HttpMethod StagedUploadHttpMethodType `json:"httpMethod,omitempty"`
	///The MIME type of the asset being uploaded.
	MimeType string                                   `json:"mimeType"`
	Resource StagedUploadTargetGenerateUploadResource `json:"resource"`
}

type StagedUploadCreateResponse struct {
	StagedUploadCreate StagedUploadCreate `json:"stagedUploadsCreate"`
}

type StagedUploadCreate struct {
	Parameters []StagedTarget `json:"stagedTargets"`
	UserErrors []UserError    `json:"userErrors"`
}

type StagedTarget struct {
	Parameters  []StagedUploadParameter `json:"parameters"`
	ResourceURL string                  `json:"resourceUrl"`
	URL         string                  `json:"url"`
}

type StagedUploadParameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
