package web

/**
 * NewsUpdateRequest struct
 */
type NewsUpdateRequest struct {
	Id      string   `validate:"required" json:"_id" bson:"_id"`
	Topic   string   `validate:"required" json:"topic" bson:"topic"`
	Title   string   `validate:"required" json:"title" bson:"title"`
	Status  string   `validate:"required" json:"status" bson:"status"`
	Tags    []string `validate:"required" json:"tags" bson:"tags"`
	Content string   `validate:"required" json:"content" bson:"content"`
}

/**
 * NewsCreateRequest struct
 */
type NewsCreateRequest struct {
	Topic   string   `validate:"required" json:"topic" bson:"topic"`
	Title   string   `validate:"required" json:"title" bson:"title"`
	Status  string   `validate:"required" json:"status" bson:"status"`
	Tags    []string `validate:"required" json:"tags" bson:"tags"`
	Content string   `validate:"required" json:"content" bson:"content"`
}

/**
 * NewsDeleteRequest struct
 */
type NewsDeleteRequest struct {
	Id string `json:"_id" bson:"_id"`
}

/**
 * NewsFindRequest struct
 */
type NewsFindRequest struct {
	Status string `json:"status"`
	Topic  string `json:"topic"`
}

/**
 * NewsResponse struct
 */
type NewsResponse struct {
	Id      string   `json:"_id" bson:"_id"`
	Topic   string   `json:"topic" bson:"topic"`
	Title   string   `json:"title" bson:"title"`
	Status  string   `json:"status" bson:"status"`
	Tags    []string `json:"tags" bson:"tags"`
	Content string   `json:"content" bson:"content"`
}
