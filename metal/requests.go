package metal

// CreateDatastoreRequest is the request body for POST /metal/datastore.
type CreateDatastoreRequest struct {
	Name   string `json:"name"`
	Scheme string `json:"scheme" enums:"triple_replication,ec4+2,ec8+3,ec8+4"`
}
