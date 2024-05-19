package requests

import (
	"fmt"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type SubscribeRequest struct {
	Email string
}

func NewSubscribeRequest(r *http.Request) (*SubscribeRequest, error) {
	var request SubscribeRequest

	err := r.ParseForm()
	if err != nil {
		return nil, fmt.Errorf("failed to parse form: %w", err)
	}

	request.Email = r.PostForm.Get("email")

	return &request, request.validate()
}

func (r SubscribeRequest) validate() error {
	return validation.Errors{
		"email": validation.Validate(&r.Email, validation.Required, is.Email),
	}.Filter()
}
