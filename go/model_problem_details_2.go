/*
 * Npcf_PolicyAuthorization Service API
 *
 * This is the Policy Authorization Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ProblemDetails2 struct {

	Type string `json:"type,omitempty"`

	Title string `json:"title,omitempty"`

	Status int32 `json:"status,omitempty"`

	Detail string `json:"detail,omitempty"`

	Instance string `json:"instance,omitempty"`

	Cause string `json:"cause,omitempty"`

	InvalidParams []InvalidParam2 `json:"invalidParams,omitempty"`
}
