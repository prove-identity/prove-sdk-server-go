// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provesdkservergo

import (
	"bytes"
	"context"
	"fmt"
	"github.com/prove-identity/prove-sdk-server-go/internal/hooks"
	"github.com/prove-identity/prove-sdk-server-go/internal/utils"
	"github.com/prove-identity/prove-sdk-server-go/models/components"
	"github.com/prove-identity/prove-sdk-server-go/models/operations"
	"github.com/prove-identity/prove-sdk-server-go/models/sdkerrors"
	"io"
	"net/http"
	"net/url"
)

type V3 struct {
	sdkConfiguration sdkConfiguration
}

func newV3(sdkConfig sdkConfiguration) *V3 {
	return &V3{
		sdkConfiguration: sdkConfig,
	}
}

// V3TokenRequest - Request OAuth token.
// Send this request to request the OAuth token.
func (s *V3) V3TokenRequest(ctx context.Context, request *components.V3TokenRequest) (*operations.V3TokenRequestResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "V3TokenRequest",
		OAuth2Scopes:   []string{},
		SecuritySource: s.sdkConfiguration.Security,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := url.JoinPath(baseURL, "/token")
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, true, "Request", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	req.Header.Set("Content-Type", reqContentType)

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := s.sdkConfiguration.Client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"400", "4XX", "500", "5XX"}, httpRes.StatusCode) {
		_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		} else if _httpRes != nil {
			httpRes = _httpRes
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	res := &operations.V3TokenRequestResponse{
		HTTPMeta: components.HTTPMetadata{
			Request:  req,
			Response: httpRes,
		},
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out components.V3TokenResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.V3TokenResponse = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 500:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out sdkerrors.Error
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			return nil, &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}

// V3ChallengeRequest - Submit challenge.
// Send this request to submit challenge information. Either a DOB or last 4 of SSN needs to be submitted if neither was submitted to the /start endpoint. It will return a correlation ID, user information, and the next step to call in the flow.
func (s *V3) V3ChallengeRequest(ctx context.Context, request *components.V3ChallengeRequest) (*operations.V3ChallengeRequestResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "V3ChallengeRequest",
		OAuth2Scopes:   []string{},
		SecuritySource: s.sdkConfiguration.Security,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := url.JoinPath(baseURL, "/v3/challenge")
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, true, "Request", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	req.Header.Set("Content-Type", reqContentType)

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := s.sdkConfiguration.Client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"400", "4XX", "500", "5XX"}, httpRes.StatusCode) {
		_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		} else if _httpRes != nil {
			httpRes = _httpRes
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	res := &operations.V3ChallengeRequestResponse{
		HTTPMeta: components.HTTPMetadata{
			Request:  req,
			Response: httpRes,
		},
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out components.V3ChallengeResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.V3ChallengeResponse = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 500:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out sdkerrors.Error
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			return nil, &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}

// V3CompleteRequest - Complete flow.
// Send this request to verify the user and complete the flow. It will return a correlation ID, user information, and the next step to call in the flow.
func (s *V3) V3CompleteRequest(ctx context.Context, request *components.V3CompleteRequest) (*operations.V3CompleteRequestResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "V3CompleteRequest",
		OAuth2Scopes:   []string{},
		SecuritySource: s.sdkConfiguration.Security,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := url.JoinPath(baseURL, "/v3/complete")
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, true, "Request", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	req.Header.Set("Content-Type", reqContentType)

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := s.sdkConfiguration.Client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"400", "4XX", "500", "5XX"}, httpRes.StatusCode) {
		_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		} else if _httpRes != nil {
			httpRes = _httpRes
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	res := &operations.V3CompleteRequestResponse{
		HTTPMeta: components.HTTPMetadata{
			Request:  req,
			Response: httpRes,
		},
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out components.V3CompleteResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.V3CompleteResponse = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 500:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out sdkerrors.Error
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			return nil, &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}

// V3StartRequest - Start flow.
// Send this request to start a Prove flow. It will return a correlation ID and an authToken for the client SDK.
func (s *V3) V3StartRequest(ctx context.Context, request *components.V3StartRequest) (*operations.V3StartRequestResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "V3StartRequest",
		OAuth2Scopes:   []string{},
		SecuritySource: s.sdkConfiguration.Security,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := url.JoinPath(baseURL, "/v3/start")
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, true, "Request", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	req.Header.Set("Content-Type", reqContentType)

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := s.sdkConfiguration.Client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"400", "4XX", "500", "5XX"}, httpRes.StatusCode) {
		_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		} else if _httpRes != nil {
			httpRes = _httpRes
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	res := &operations.V3StartRequestResponse{
		HTTPMeta: components.HTTPMetadata{
			Request:  req,
			Response: httpRes,
		},
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out components.V3StartResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.V3StartResponse = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 500:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out sdkerrors.Error
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			return nil, &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}

// V3ValidateRequest - Validate phone number.
// Send this request to check the phone number entered/discovered earlier in the flow is validated. It will return a correlation ID and the next step.
func (s *V3) V3ValidateRequest(ctx context.Context, request *components.V3ValidateRequest) (*operations.V3ValidateRequestResponse, error) {
	hookCtx := hooks.HookContext{
		Context:        ctx,
		OperationID:    "V3ValidateRequest",
		OAuth2Scopes:   []string{},
		SecuritySource: s.sdkConfiguration.Security,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := url.JoinPath(baseURL, "/v3/validate")
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, true, "Request", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	req.Header.Set("Content-Type", reqContentType)

	if err := utils.PopulateSecurity(ctx, req, s.sdkConfiguration.Security); err != nil {
		return nil, err
	}

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{HookContext: hookCtx}, req)
	if err != nil {
		return nil, err
	}

	httpRes, err := s.sdkConfiguration.Client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"400", "4XX", "500", "5XX"}, httpRes.StatusCode) {
		_httpRes, err := s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{HookContext: hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		} else if _httpRes != nil {
			httpRes = _httpRes
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{HookContext: hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}

	res := &operations.V3ValidateRequestResponse{
		HTTPMeta: components.HTTPMetadata{
			Request:  req,
			Response: httpRes,
		},
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out components.V3ValidateResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.V3ValidateResponse = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 500:
		switch {
		case utils.MatchContentType(httpRes.Header.Get("Content-Type"), `application/json`):
			var out sdkerrors.Error
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			return nil, &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", httpRes.Header.Get("Content-Type")), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
		return nil, sdkerrors.NewSDKError("unknown status code returned", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil

}
