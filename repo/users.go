package repo

import (
	"fmt"
	"strconv"
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
	d.UserIdx["ID"] = make(map[string][]*User)
	d.UserIdx["URL"] = make(map[string][]*User)
	d.UserIdx["ExternalID"] = make(map[string][]*User)
	d.UserIdx["Name"] = make(map[string][]*User)
	d.UserIdx["Alias"] = make(map[string][]*User)
	d.UserIdx["CreatedAt"] = make(map[string][]*User)
	d.UserIdx["Active"] = make(map[string][]*User)
	d.UserIdx["Verified"] = make(map[string][]*User)
	d.UserIdx["Shared"] = make(map[string][]*User)
	d.UserIdx["Locale"] = make(map[string][]*User)
	d.UserIdx["Timezone"] = make(map[string][]*User)
	d.UserIdx["LastLoginAt"] = make(map[string][]*User)
	d.UserIdx["Email"] = make(map[string][]*User)
	d.UserIdx["Phone"] = make(map[string][]*User)
	d.UserIdx["Signature"] = make(map[string][]*User)
	d.UserIdx["OrganizationID"] = make(map[string][]*User)
	d.UserIdx["Tags"] = make(map[string][]*User)
	d.UserIdx["Suspended"] = make(map[string][]*User)
	d.UserIdx["Role"] = make(map[string][]*User)

	for i := range d.Users {
		d.UserIdx["ID"][fmt.Sprintf("%d", d.Users[i].ID)] = append(d.UserIdx["ID"][fmt.Sprintf("%d", d.Users[i].ID)], d.Users[i])
		d.UserIdx["URL"][d.Users[i].URL] = append(d.UserIdx["URL"][d.Users[i].URL], d.Users[i])
		d.UserIdx["ExternalID"][d.Users[i].ExternalID] = append(d.UserIdx["ExternalID"][d.Users[i].ExternalID], d.Users[i])
		d.UserIdx["Name"][d.Users[i].Name] = append(d.UserIdx["Name"][d.Users[i].Name], d.Users[i])
		d.UserIdx["Alias"][d.Users[i].Alias] = append(d.UserIdx["Alias"][d.Users[i].Alias], d.Users[i])
		d.UserIdx["CreatedAt"][d.Users[i].CreatedAt] = append(d.UserIdx["CreatedAt"][d.Users[i].CreatedAt], d.Users[i])
		d.UserIdx["Active"][strconv.FormatBool(d.Users[i].Active)] = append(d.UserIdx["Active"][strconv.FormatBool(d.Users[i].Active)], d.Users[i])
		d.UserIdx["Verified"][strconv.FormatBool(d.Users[i].Verified)] = append(d.UserIdx["Verified"][strconv.FormatBool(d.Users[i].Verified)], d.Users[i])
		d.UserIdx["Shared"][strconv.FormatBool(d.Users[i].Shared)] = append(d.UserIdx["Shared"][strconv.FormatBool(d.Users[i].Shared)], d.Users[i])
		d.UserIdx["Locale"][d.Users[i].Locale] = append(d.UserIdx["Locale"][d.Users[i].Locale], d.Users[i])
		d.UserIdx["Timezone"][d.Users[i].Timezone] = append(d.UserIdx["Timezone"][d.Users[i].Timezone], d.Users[i])
		d.UserIdx["LastLoginAt"][d.Users[i].LastLoginAt] = append(d.UserIdx["LastLoginAt"][d.Users[i].LastLoginAt], d.Users[i])
		d.UserIdx["Email"][d.Users[i].Email] = append(d.UserIdx["Email"][d.Users[i].Email], d.Users[i])
		d.UserIdx["Phone"][d.Users[i].Phone] = append(d.UserIdx["Phone"][d.Users[i].Phone], d.Users[i])
		d.UserIdx["Signature"][d.Users[i].Signature] = append(d.UserIdx["Signature"][d.Users[i].Signature], d.Users[i])
		d.UserIdx["OrganizationID"][fmt.Sprintf("%d", d.Users[i].OrganizationID)] = append(d.UserIdx["OrganizationID"][fmt.Sprintf("%d", d.Users[i].OrganizationID)], d.Users[i])
		for _, tag := range d.Users[i].Tags {
			d.UserIdx["Tags"][tag] = append(d.UserIdx["Tags"][tag], d.Users[i])
		}
		d.UserIdx["Suspended"][strconv.FormatBool(d.Users[i].Suspended)] = append(d.UserIdx["Suspended"][strconv.FormatBool(d.Users[i].Suspended)], d.Users[i])
		d.UserIdx["Role"][d.Users[i].Role] = append(d.UserIdx["Role"][d.Users[i].Role], d.Users[i])
	}
}
