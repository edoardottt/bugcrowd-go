package api

type FileAttachment struct {
	ID    string `json:"id,omitempty"`
	Type  string `json:"type,omitempty"`
	Links struct {
		Self string `json:"self,omitempty"`
	} `json:"links,omitempty"`
	Attributes struct {
		FileName    string `json:"file_name,omitempty"`
		FileSize    int    `json:"file_size,omitempty"`
		FileType    string `json:"file_type,omitempty"`
		DownloadURL string `json:"download_url,omitempty"`
		S3SignedURL string `json:"s3_signed_url,omitempty"`
	} `json:"attributes,omitempty"`
	Relationships struct {
		Parent struct {
			Links struct {
				Related struct {
					Href string `json:"href,omitempty"`
				} `json:"related,omitempty"`
			} `json:"links,omitempty"`
			Data struct {
				ID   string `json:"id,omitempty"`
				Type string `json:"type,omitempty"`
			} `json:"data,omitempty"`
		} `json:"parent,omitempty"`
	} `json:"relationships,omitempty"`
}
