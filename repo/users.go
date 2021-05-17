package repo

import (
	"fmt"
	"strconv"
	"strings"
)

// User -
type User struct {
	ID             int      `json:"_id"`
	URL            string   `json:"url"`
	ExternalID     string   `json:"external_id"`
	Name           string   `json:"name"`
	Alias          string   `json:"alias"`
	CreatedAt      string   `json:"created_at"`
	Active         bool     `json:"active"`
	Verified       bool     `json:"verified"`
	Shared         bool     `json:"shared"`
	Locale         string   `json:"locale"`
	Timezone       string   `json:"timezone"`
	LastLoginAt    string   `json:"last_login_at"`
	Email          string   `json:"email"`
	Phone          string   `json:"phone"`
	Signature      string   `json:"signature"`
	OrganizationID int      `json:"organization_id"`
	Tags           []string `json:"tags"`
	Suspended      bool     `json:"suspended"`
	Role           string   `json:"role"`
}

// UserIndexes -
func (d *Data) UserIndexes() {
	// map[fieldname]map[fieldvalue][]*User
	d.UserIdx = make(map[string]map[string][]*User)
	d.UserIdx["_id"] = make(map[string][]*User)
	d.UserIdx["url"] = make(map[string][]*User)
	d.UserIdx["external_id"] = make(map[string][]*User)
	d.UserIdx["name"] = make(map[string][]*User)
	d.UserIdx["alias"] = make(map[string][]*User)
	d.UserIdx["created_at"] = make(map[string][]*User)
	d.UserIdx["active"] = make(map[string][]*User)
	d.UserIdx["verified"] = make(map[string][]*User)
	d.UserIdx["shared"] = make(map[string][]*User)
	d.UserIdx["locale"] = make(map[string][]*User)
	d.UserIdx["timezone"] = make(map[string][]*User)
	d.UserIdx["last_login_at"] = make(map[string][]*User)
	d.UserIdx["email"] = make(map[string][]*User)
	d.UserIdx["phone"] = make(map[string][]*User)
	d.UserIdx["signature"] = make(map[string][]*User)
	d.UserIdx["organization_id"] = make(map[string][]*User)
	d.UserIdx["tags"] = make(map[string][]*User)
	d.UserIdx["suspended"] = make(map[string][]*User)
	d.UserIdx["role"] = make(map[string][]*User)
	d.Terms["users"] = map[string]struct{}{
		"_id":             struct{}{},
		"url":             struct{}{},
		"external_id":     struct{}{},
		"name":            struct{}{},
		"alias":           struct{}{},
		"created_at":      struct{}{},
		"active":          struct{}{},
		"verified":        struct{}{},
		"shared":          struct{}{},
		"locale":          struct{}{},
		"timezone":        struct{}{},
		"last_login_at":   struct{}{},
		"email":           struct{}{},
		"phone":           struct{}{},
		"signature":       struct{}{},
		"organization_id": struct{}{},
		"tags":            struct{}{},
		"suspended":       struct{}{},
		"role":            struct{}{},
	}

	for i := range d.Users {
		d.UserIdx["_id"][fmt.Sprintf("%d", d.Users[i].ID)] = append(d.UserIdx["_id"][fmt.Sprintf("%d", d.Users[i].ID)], d.Users[i])
		d.UserIdx["url"][strings.ToLower(d.Users[i].URL)] = append(d.UserIdx["url"][strings.ToLower(d.Users[i].URL)], d.Users[i])
		d.UserIdx["external_id"][strings.ToLower(d.Users[i].ExternalID)] = append(d.UserIdx["external_id"][strings.ToLower(d.Users[i].ExternalID)], d.Users[i])
		d.UserIdx["name"][strings.ToLower(d.Users[i].Name)] = append(d.UserIdx["name"][strings.ToLower(d.Users[i].Name)], d.Users[i])
		d.UserIdx["alias"][strings.ToLower(d.Users[i].Alias)] = append(d.UserIdx["alias"][strings.ToLower(d.Users[i].Alias)], d.Users[i])
		d.UserIdx["created_at"][strings.ToLower(d.Users[i].CreatedAt)] = append(d.UserIdx["created_at"][strings.ToLower(d.Users[i].CreatedAt)], d.Users[i])
		d.UserIdx["active"][strconv.FormatBool(d.Users[i].Active)] = append(d.UserIdx["active"][strconv.FormatBool(d.Users[i].Active)], d.Users[i])
		d.UserIdx["verified"][strconv.FormatBool(d.Users[i].Verified)] = append(d.UserIdx["verified"][strconv.FormatBool(d.Users[i].Verified)], d.Users[i])
		d.UserIdx["shared"][strconv.FormatBool(d.Users[i].Shared)] = append(d.UserIdx["shared"][strconv.FormatBool(d.Users[i].Shared)], d.Users[i])
		d.UserIdx["locale"][strings.ToLower(d.Users[i].Locale)] = append(d.UserIdx["locale"][strings.ToLower(d.Users[i].Locale)], d.Users[i])
		d.UserIdx["timezone"][strings.ToLower(d.Users[i].Timezone)] = append(d.UserIdx["timezone"][strings.ToLower(d.Users[i].Timezone)], d.Users[i])
		d.UserIdx["last_login_at"][strings.ToLower(d.Users[i].LastLoginAt)] = append(d.UserIdx["last_login_at"][strings.ToLower(d.Users[i].LastLoginAt)], d.Users[i])
		d.UserIdx["email"][strings.ToLower(d.Users[i].Email)] = append(d.UserIdx["email"][strings.ToLower(d.Users[i].Email)], d.Users[i])
		d.UserIdx["phone"][strings.ToLower(d.Users[i].Phone)] = append(d.UserIdx["phone"][strings.ToLower(d.Users[i].Phone)], d.Users[i])
		d.UserIdx["signature"][strings.ToLower(d.Users[i].Signature)] = append(d.UserIdx["signature"][strings.ToLower(d.Users[i].Signature)], d.Users[i])
		d.UserIdx["organization_id"][fmt.Sprintf("%d", d.Users[i].OrganizationID)] = append(d.UserIdx["organization_id"][fmt.Sprintf("%d", d.Users[i].OrganizationID)], d.Users[i])
		for _, tag := range d.Users[i].Tags {
			d.UserIdx["tags"][strings.ToLower(tag)] = append(d.UserIdx["Tags"][strings.ToLower(tag)], d.Users[i])
		}
		d.UserIdx["suspended"][strconv.FormatBool(d.Users[i].Suspended)] = append(d.UserIdx["suspended"][strconv.FormatBool(d.Users[i].Suspended)], d.Users[i])
		d.UserIdx["role"][strings.ToLower(d.Users[i].Role)] = append(d.UserIdx["Role"][strings.ToLower(d.Users[i].Role)], d.Users[i])
	}
}
