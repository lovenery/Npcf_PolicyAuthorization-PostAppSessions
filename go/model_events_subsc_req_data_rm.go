/*
 * Npcf_PolicyAuthorization Service API
 *
 * This is the Policy Authorization Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// EventsSubscReqDataRm - this data type is defined in the same way as the EventsSubscReqData data type, but with the OpenAPI nullable property set to true.
type EventsSubscReqDataRm struct {

	Events []AfEventSubscription `json:"events"`

	NotifUri string `json:"notifUri,omitempty"`

	UsgThres *UsageThresholdRm `json:"usgThres,omitempty"`
}