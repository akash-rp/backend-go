package models

type PhpIni struct {
	MaxExecutionTime      int    `json:"MaxExecutionTime"`
	MaxFileUploads        int    `json:"MaxFileUploads"`
	MaxInputTime          int    `json:"MaxInputTime"`
	MaxInputVars          int    `json:"MaxInputVars"`
	MemoryLimit           int    `json:"MemoryLimit"`
	PostMaxSize           int    `json:"PostMaxSize"`
	SessionCookieLifetime int    `json:"SessionCookieLifetime"`
	SessionGcMaxlifetime  int    `json:"SessionGcMaxlifetime"`
	ShortOpenTag          string `json:"ShortOpenTag"`
	UploadMaxFilesize     int    `json:"UploadMaxFilesize"`
	Timezone              string `json:"Timezone"`
	OpenBaseDir           string `json:"OpenBaseDir"`
}

type UpdatePhpIni struct {
	MaxExecutionTime      int    `json:"MaxExecutionTime"`
	MaxFileUploads        int    `json:"MaxFileUploads"`
	MaxInputTime          int    `json:"MaxInputTime"`
	MaxInputVars          int    `json:"MaxInputVars"`
	MemoryLimit           int    `json:"MemoryLimit"`
	PostMaxSize           int    `json:"PostMaxSize"`
	SessionCookieLifetime int    `json:"SessionCookieLifetime"`
	SessionGcMaxlifetime  int    `json:"SessionGcMaxlifetime"`
	ShortOpenTag          string `json:"ShortOpenTag"`
	UploadMaxFilesize     int    `json:"UploadMaxFilesize"`
	Timezone              string `json:"Timezone"`
	OpenBaseDir           string `json:"OpenBaseDir"`
}

type PhpSettings struct {
	PHPLSAPICHILDREN       int `json:"PHP_LSAPI_CHILDREN"`
	PHPLSAPIMAXIDLE        int `json:"PHP_LSAPI_MAX_IDLE"`
	PHPLSAPIMAXPROCESSTIME int `json:"PHP_LSAPI_MAX_PROCESS_TIME"`
	PHPLSAPIMAXREQUESTS    int `json:"PHP_LSAPI_MAX_REQUESTS"`
	PHPLSAPISLOWREQMSECS   int `json:"PHP_LSAPI_SLOW_REQ_MSECS"`
	InitTimeout            int `json:"initTimeout"`
	Instances              int `json:"instances"`
	MaxConns               int `json:"maxConns"`
	RetryTimeout           int `json:"retryTimeout"`
}

type UpdatePhpSettings struct {
	User     string      `json:"user"`
	Settings PhpSettings `json:"settings"`
}
