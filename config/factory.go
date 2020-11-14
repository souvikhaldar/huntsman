package config

// Config is the basic structure for holding mongoDB configuration information
type Config struct {
	MongoURI        string `json:"mongo_uri"`
	MongoUsername   string `json:"mongo_username"`
	MongoPassword   string `json:"mongo_password"`
	MongoDatabase   string `json:"mongo_database"`
	MongoCollection string `json:"mongo_collection"`
	JSFilePath      string `json:"js_file_path"`
}
