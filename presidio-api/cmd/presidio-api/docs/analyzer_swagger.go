package docs

import types "github.com/Microsoft/presidio-genproto/golang"

// swagger:route GET /analyzer/recognizers/{id} analyzer getRecognizer
//
// Get an existing recognizer
//
// responses:
//   200: recognizersGetResponse

// A response includes a list of recognized fields in the text to be anonymized.
// swagger:response recognizersGetResponse
type regognizerResponseWrapper struct {
	// in:body
	Body types.RecognizersGetResponse
}

// swagger:parameters getRecognizer
type regognizerParamsWrapper struct {
	// Recognizer name
	//
	// in: path
	// required: true
	Id string `json:"id"`
}
