/*
 * Npcf_PolicyAuthorization Service API
 *
 * This is the Policy Authorization Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// AppSessionContextRespData - Describes the authorization data of an Individual Application Session Context created by the PCF.
type AppSessionContextRespData struct {

	ServAuthInfo *ServAuthInfo `json:"servAuthInfo,omitempty"`

	SuppFeat string `json:"suppFeat,omitempty"`
}
