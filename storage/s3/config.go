package s3

// Config S3 storage configuration
type Config struct {
	Local           string   `json:"local" yaml:"local" default:"/var/lib/audit"`
	AccessKey       string   `json:"accessKey" yaml:"accessKey"`
	SecretKey       string   `json:"secretKey" yaml:"secretKey"`
	Bucket          string   `json:"bucket" yaml:"bucket"`
	Region          string   `json:"region" yaml:"region"`
	Endpoint        string   `json:"endpoint" yaml:"endpoint"`
	CaCert          string   `json:"cacert" yaml:"cacert"`
	ACL             string   `json:"acl" yaml:"acl"`
	PathStyleAccess bool     `json:"pathStyleAccess" yaml:"pathStyleAccess"`
	UploadPartSize  uint     `json:"uploadPartSize" yaml:"uploadPartSize"`
	ParallelUploads uint     `json:"parallelUploads" yaml:"parallelUploads"`
	Metadata        Metadata `json:"metadata" yaml:"metadata"`
}

// Metadata Metadata configuration for the S3 storage
type Metadata struct {
	IP       bool `json:"ip" yaml:"ip"`
	Username bool `json:"username" yaml:"username"`
}
