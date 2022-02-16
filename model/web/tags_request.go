package web

/**
 * TagsDeleteRequest struct
 */
type TagsDeleteRequest struct {
	Id string `validate:"required" json:"_id" bson:"_id"`
}

/**
 * TagsRequest struct
 */
type TagsRequest struct {
	Name string `validate:"required" json:"name" bson:"name"`
}

/**
 * TagsResponse struct
 */
type TagsResponse struct {
	Id   string `json:"_id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}

/**
 * TagsUpdateRequest struct
 */
type TagsUpdateRequest struct {
	Id   string `validate:"required" json:"_id" bson:"_id"`
	Name string `json:"name"`
}
