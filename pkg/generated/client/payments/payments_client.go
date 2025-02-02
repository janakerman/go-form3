// Code generated by go-swagger; DO NOT EDIT.

// :Form3: Testing!

package payments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/form3tech-oss/go-form3/pkg/client"
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new payments API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, defaults client.Defaults) *Client {
	return &Client{transport: transport, formats: formats, Defaults: defaults}
}

/*
Client for payments API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	Defaults  client.Defaults
}

// range of operations

/*
create payment API
*/
func (a *CreatePaymentRequest) Do() (*CreatePaymentCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreatePayment",
		Method:             "POST",
		PathPattern:        "/transaction/payments",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreatePaymentReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreatePaymentCreated), nil

}

func (a *CreatePaymentRequest) MustDo() *CreatePaymentCreated {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
create payment advice API
*/
func (a *CreatePaymentAdviceRequest) Do() (*CreatePaymentAdviceCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreatePaymentAdvice",
		Method:             "POST",
		PathPattern:        "/transaction/payments/{id}/advices",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreatePaymentAdviceReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreatePaymentAdviceCreated), nil

}

func (a *CreatePaymentAdviceRequest) MustDo() *CreatePaymentAdviceCreated {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
create payment advice submission API
*/
func (a *CreatePaymentAdviceSubmissionRequest) Do() (*CreatePaymentAdviceSubmissionCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreatePaymentAdviceSubmission",
		Method:             "POST",
		PathPattern:        "/transaction/payments/{id}/advices/{adviceId}/submissions",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreatePaymentAdviceSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreatePaymentAdviceSubmissionCreated), nil

}

func (a *CreatePaymentAdviceSubmissionRequest) MustDo() *CreatePaymentAdviceSubmissionCreated {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
create payment recall API
*/
func (a *CreatePaymentRecallRequest) Do() (*CreatePaymentRecallCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreatePaymentRecall",
		Method:             "POST",
		PathPattern:        "/transaction/payments/{id}/recalls",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreatePaymentRecallReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreatePaymentRecallCreated), nil

}

func (a *CreatePaymentRecallRequest) MustDo() *CreatePaymentRecallCreated {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
create payment recall decision API
*/
func (a *CreatePaymentRecallDecisionRequest) Do() (*CreatePaymentRecallDecisionCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreatePaymentRecallDecision",
		Method:             "POST",
		PathPattern:        "/transaction/payments/{id}/recalls/{recallId}/decisions",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreatePaymentRecallDecisionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreatePaymentRecallDecisionCreated), nil

}

func (a *CreatePaymentRecallDecisionRequest) MustDo() *CreatePaymentRecallDecisionCreated {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
create payment recall decision submission API
*/
func (a *CreatePaymentRecallDecisionSubmissionRequest) Do() (*CreatePaymentRecallDecisionSubmissionCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreatePaymentRecallDecisionSubmission",
		Method:             "POST",
		PathPattern:        "/transaction/payments/{id}/recalls/{recallId}/decisions/{decisionId}/submissions",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreatePaymentRecallDecisionSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreatePaymentRecallDecisionSubmissionCreated), nil

}

func (a *CreatePaymentRecallDecisionSubmissionRequest) MustDo() *CreatePaymentRecallDecisionSubmissionCreated {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
create payment recall submission API
*/
func (a *CreatePaymentRecallSubmissionRequest) Do() (*CreatePaymentRecallSubmissionCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreatePaymentRecallSubmission",
		Method:             "POST",
		PathPattern:        "/transaction/payments/{id}/recalls/{recallId}/submissions",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreatePaymentRecallSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreatePaymentRecallSubmissionCreated), nil

}

func (a *CreatePaymentRecallSubmissionRequest) MustDo() *CreatePaymentRecallSubmissionCreated {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
create payment return API
*/
func (a *CreatePaymentReturnRequest) Do() (*CreatePaymentReturnCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreatePaymentReturn",
		Method:             "POST",
		PathPattern:        "/transaction/payments/{id}/returns",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreatePaymentReturnReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreatePaymentReturnCreated), nil

}

func (a *CreatePaymentReturnRequest) MustDo() *CreatePaymentReturnCreated {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
create payment return reversal API
*/
func (a *CreatePaymentReturnReversalRequest) Do() (*CreatePaymentReturnReversalCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreatePaymentReturnReversal",
		Method:             "POST",
		PathPattern:        "/transaction/payments/{id}/returns/{returnId}/reversals",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreatePaymentReturnReversalReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreatePaymentReturnReversalCreated), nil

}

func (a *CreatePaymentReturnReversalRequest) MustDo() *CreatePaymentReturnReversalCreated {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
create payment return submission API
*/
func (a *CreatePaymentReturnSubmissionRequest) Do() (*CreatePaymentReturnSubmissionCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreatePaymentReturnSubmission",
		Method:             "POST",
		PathPattern:        "/transaction/payments/{id}/returns/{returnId}/submissions",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreatePaymentReturnSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreatePaymentReturnSubmissionCreated), nil

}

func (a *CreatePaymentReturnSubmissionRequest) MustDo() *CreatePaymentReturnSubmissionCreated {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
create payment reversal API
*/
func (a *CreatePaymentReversalRequest) Do() (*CreatePaymentReversalCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreatePaymentReversal",
		Method:             "POST",
		PathPattern:        "/transaction/payments/{id}/reversals",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreatePaymentReversalReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreatePaymentReversalCreated), nil

}

func (a *CreatePaymentReversalRequest) MustDo() *CreatePaymentReversalCreated {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
create payment reversal submission API
*/
func (a *CreatePaymentReversalSubmissionRequest) Do() (*CreatePaymentReversalSubmissionCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreatePaymentReversalSubmission",
		Method:             "POST",
		PathPattern:        "/transaction/payments/{id}/reversals/{reversalId}/submissions",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnc.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreatePaymentReversalSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreatePaymentReversalSubmissionCreated), nil

}

func (a *CreatePaymentReversalSubmissionRequest) MustDo() *CreatePaymentReversalSubmissionCreated {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
create payment submission API
*/
func (a *CreatePaymentSubmissionRequest) Do() (*CreatePaymentSubmissionCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreatePaymentSubmission",
		Method:             "POST",
		PathPattern:        "/transaction/payments/{id}/submissions",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &CreatePaymentSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreatePaymentSubmissionCreated), nil

}

func (a *CreatePaymentSubmissionRequest) MustDo() *CreatePaymentSubmissionCreated {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
get payment API
*/
func (a *GetPaymentRequest) Do() (*GetPaymentOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPayment",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentOK), nil

}

func (a *GetPaymentRequest) MustDo() *GetPaymentOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
get payment admissions API
*/
func (a *GetPaymentAdmissionsRequest) Do() (*GetPaymentAdmissionsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentAdmissions",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/admissions/{admissionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentAdmissionsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentAdmissionsOK), nil

}

func (a *GetPaymentAdmissionsRequest) MustDo() *GetPaymentAdmissionsOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
get payment advice submission API
*/
func (a *GetPaymentAdviceSubmissionRequest) Do() (*GetPaymentAdviceSubmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentAdviceSubmission",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/advices/{adviceId}/submissions/{submissionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentAdviceSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentAdviceSubmissionOK), nil

}

func (a *GetPaymentAdviceSubmissionRequest) MustDo() *GetPaymentAdviceSubmissionOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
get payment recall API
*/
func (a *GetPaymentRecallRequest) Do() (*GetPaymentRecallOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentRecall",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/recalls/{recallId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentRecallReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentRecallOK), nil

}

func (a *GetPaymentRecallRequest) MustDo() *GetPaymentRecallOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
get payment recall admission API
*/
func (a *GetPaymentRecallAdmissionRequest) Do() (*GetPaymentRecallAdmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentRecallAdmission",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/recalls/{recallId}/admissions/{admissionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentRecallAdmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentRecallAdmissionOK), nil

}

func (a *GetPaymentRecallAdmissionRequest) MustDo() *GetPaymentRecallAdmissionOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
get payment recall decision API
*/
func (a *GetPaymentRecallDecisionRequest) Do() (*GetPaymentRecallDecisionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentRecallDecision",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/recalls/{recallId}/decisions/{decisionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentRecallDecisionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentRecallDecisionOK), nil

}

func (a *GetPaymentRecallDecisionRequest) MustDo() *GetPaymentRecallDecisionOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
get payment recall decision admission API
*/
func (a *GetPaymentRecallDecisionAdmissionRequest) Do() (*GetPaymentRecallDecisionAdmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentRecallDecisionAdmission",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/recalls/{recallId}/decisions/{decisionId}/admissions/{admissionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentRecallDecisionAdmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentRecallDecisionAdmissionOK), nil

}

func (a *GetPaymentRecallDecisionAdmissionRequest) MustDo() *GetPaymentRecallDecisionAdmissionOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
get payment recall decision submission API
*/
func (a *GetPaymentRecallDecisionSubmissionRequest) Do() (*GetPaymentRecallDecisionSubmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentRecallDecisionSubmission",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/recalls/{recallId}/decisions/{decisionId}/submissions/{submissionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentRecallDecisionSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentRecallDecisionSubmissionOK), nil

}

func (a *GetPaymentRecallDecisionSubmissionRequest) MustDo() *GetPaymentRecallDecisionSubmissionOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
get payment recall reversal API
*/
func (a *GetPaymentRecallReversalRequest) Do() (*GetPaymentRecallReversalOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentRecallReversal",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/recalls/{recallId}/reversals/{reversalId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentRecallReversalReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentRecallReversalOK), nil

}

func (a *GetPaymentRecallReversalRequest) MustDo() *GetPaymentRecallReversalOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
get payment recall reversal admission API
*/
func (a *GetPaymentRecallReversalAdmissionRequest) Do() (*GetPaymentRecallReversalAdmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentRecallReversalAdmission",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/recalls/{recallId}/reversals/{reversalId}/admissions/{admissionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentRecallReversalAdmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentRecallReversalAdmissionOK), nil

}

func (a *GetPaymentRecallReversalAdmissionRequest) MustDo() *GetPaymentRecallReversalAdmissionOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
get payment recall submission API
*/
func (a *GetPaymentRecallSubmissionRequest) Do() (*GetPaymentRecallSubmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentRecallSubmission",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/recalls/{recallId}/submissions/{submissionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentRecallSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentRecallSubmissionOK), nil

}

func (a *GetPaymentRecallSubmissionRequest) MustDo() *GetPaymentRecallSubmissionOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
get payment return API
*/
func (a *GetPaymentReturnRequest) Do() (*GetPaymentReturnOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentReturn",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/returns/{returnId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentReturnReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentReturnOK), nil

}

func (a *GetPaymentReturnRequest) MustDo() *GetPaymentReturnOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
get payment return admission API
*/
func (a *GetPaymentReturnAdmissionRequest) Do() (*GetPaymentReturnAdmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentReturnAdmission",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/returns/{returnId}/admissions/{admissionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentReturnAdmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentReturnAdmissionOK), nil

}

func (a *GetPaymentReturnAdmissionRequest) MustDo() *GetPaymentReturnAdmissionOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
get payment return reversal API
*/
func (a *GetPaymentReturnReversalRequest) Do() (*GetPaymentReturnReversalOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentReturnReversal",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/returns/{returnId}/reversals/{reversalId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentReturnReversalReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentReturnReversalOK), nil

}

func (a *GetPaymentReturnReversalRequest) MustDo() *GetPaymentReturnReversalOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
get payment return reversal admission API
*/
func (a *GetPaymentReturnReversalAdmissionRequest) Do() (*GetPaymentReturnReversalAdmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentReturnReversalAdmission",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/returns/{returnId}/reversals/{reversalId}/admissions/{admissionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentReturnReversalAdmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentReturnReversalAdmissionOK), nil

}

func (a *GetPaymentReturnReversalAdmissionRequest) MustDo() *GetPaymentReturnReversalAdmissionOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
get payment return submission API
*/
func (a *GetPaymentReturnSubmissionRequest) Do() (*GetPaymentReturnSubmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentReturnSubmission",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/returns/{returnId}/submissions/{submissionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentReturnSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentReturnSubmissionOK), nil

}

func (a *GetPaymentReturnSubmissionRequest) MustDo() *GetPaymentReturnSubmissionOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
get payment reversal API
*/
func (a *GetPaymentReversalRequest) Do() (*GetPaymentReversalOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentReversal",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/reversals/{reversalId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentReversalReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentReversalOK), nil

}

func (a *GetPaymentReversalRequest) MustDo() *GetPaymentReversalOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
get payment reversal admission API
*/
func (a *GetPaymentReversalAdmissionRequest) Do() (*GetPaymentReversalAdmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentReversalAdmission",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/reversals/{reversalId}/admissions/{admissionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentReversalAdmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentReversalAdmissionOK), nil

}

func (a *GetPaymentReversalAdmissionRequest) MustDo() *GetPaymentReversalAdmissionOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
get payment reversal submission API
*/
func (a *GetPaymentReversalSubmissionRequest) Do() (*GetPaymentReversalSubmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentReversalSubmission",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/reversals/{reversalId}/submissions/{submissionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentReversalSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentReversalSubmissionOK), nil

}

func (a *GetPaymentReversalSubmissionRequest) MustDo() *GetPaymentReversalSubmissionOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
get payment submission API
*/
func (a *GetPaymentSubmissionRequest) Do() (*GetPaymentSubmissionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentSubmission",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/submissions/{submissionId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentSubmissionReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentSubmissionOK), nil

}

func (a *GetPaymentSubmissionRequest) MustDo() *GetPaymentSubmissionOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
get payments health API
*/
func (a *GetPaymentsHealthRequest) Do() (*GetPaymentsHealthOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPaymentsHealth",
		Method:             "GET",
		PathPattern:        "/transaction/payments/health",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPaymentsHealthReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPaymentsHealthOK), nil

}

func (a *GetPaymentsHealthRequest) MustDo() *GetPaymentsHealthOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
get positions API
*/
func (a *GetPositionsRequest) Do() (*GetPositionsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetPositions",
		Method:             "GET",
		PathPattern:        "/organisation/positions",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &GetPositionsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPositionsOK), nil

}

func (a *GetPositionsRequest) MustDo() *GetPositionsOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
list payment advices API
*/
func (a *ListPaymentAdvicesRequest) Do() (*ListPaymentAdvicesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListPaymentAdvices",
		Method:             "GET",
		PathPattern:        "/transaction/payments/{id}/advices/{adviceId}",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &ListPaymentAdvicesReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListPaymentAdvicesOK), nil

}

func (a *ListPaymentAdvicesRequest) MustDo() *ListPaymentAdvicesOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/*
list payments API
*/
func (a *ListPaymentsRequest) Do() (*ListPaymentsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ListPayments",
		Method:             "GET",
		PathPattern:        "/transaction/payments",
		ProducesMediaTypes: []string{"application/json", "application/vnd.api+json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             a,
		Reader:             &ListPaymentsReader{formats: a.formats},
		//AuthInfo: authInfo,
		Context: a.Context,
		Client:  a.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListPaymentsOK), nil

}

func (a *ListPaymentsRequest) MustDo() *ListPaymentsOK {
	r, err := a.Do()
	if err != nil {
		panic(err)
	}
	return r
}

/////////

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
