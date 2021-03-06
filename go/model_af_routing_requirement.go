/*
 * Npcf_PolicyAuthorization Service API
 *
 * This is the Policy Authorization Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// AfRoutingRequirement - describes the event information delivered in the subscription
type AfRoutingRequirement struct {

	AppReloc bool `json:"appReloc,omitempty"`

	RouteToLocs []RouteToLocation `json:"routeToLocs,omitempty"`

	SpVal *SpatialValidity `json:"spVal,omitempty"`

	TempVals []TemporalValidity `json:"tempVals,omitempty"`

	UpPathChgSub *UpPathChgEvent `json:"upPathChgSub,omitempty"`
}
