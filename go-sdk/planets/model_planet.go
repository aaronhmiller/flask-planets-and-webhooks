/*
 * Planets and Webhooks Demo API
 *
 * Simple flask-backed API showing some example API endpoints and with webhook debugging features.
 *
 * API version: 0.0.1
 * Contact: github@lornajane.net
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package planets
// Planet struct for Planet
type Planet struct {
	// Name of planet
	Name string `json:"name,omitempty"`
	// Order in place from the sun
	Position float32 `json:"position,omitempty"`
	// Number of moons, according to NASA
	Moons float32 `json:"moons,omitempty"`
}
