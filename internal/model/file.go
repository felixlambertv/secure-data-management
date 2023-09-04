package model

type Metadata struct {
	Name     string `bson:"name"`
	KeyName  string `bson:"keyName"`
	Size     int    `bson:"size"`
	Type     string `bson:"type"`
	UploadAt int    `bson:"uploadAt"`
}

type File struct {
	ID         string   `bson:"_id,omitempty"`
	Metadata   Metadata `bson:"metadata"`
	Permission []string `bson:"permission"`
}
