/*
 * NSSF NSSAI Availability
 *
 * NSSF NSSAI Availability Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type AuthorizedNssaiAvailabilityData struct {
	Tai *Tai `json:"tai" bson:"tai"`

	SupportedSnssaiList []Snssai `json:"supportedSnssaiList" bson:"supportedSnssaiList"`

	RestrictedSnssaiList []RestrictedSnssai `json:"restrictedSnssaiList,omitempty" bson:"restrictedSnssaiList"`
}
