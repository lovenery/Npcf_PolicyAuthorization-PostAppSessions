/*
 * Npcf_PolicyAuthorization Service API
 *
 * This is the Policy Authorization Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// QosNotificationControlInfo - Indicates whether the QoS targets for a GRB flow are not guaranteed or guaranteed again
type QosNotificationControlInfo struct {

	NotifType *QosNotifType `json:"notifType"`

	Flows []Flows `json:"flows,omitempty"`
}