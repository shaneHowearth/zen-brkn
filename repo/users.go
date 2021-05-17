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

// ToDTO - TODO
func (u *User) ToDTO() map[string][]string {
	m := map[string][]string{}
	m["_id"] = []string{fmt.Sprintf("%d", u.ID)}
	m["url"] = []string{u.URL}
	m["external_id"] = []string{u.ExternalID}
	m["name"] = []string{u.Name}
	m["alias"] = []string{u.Alias}
	m["created_at"] = []string{u.CreatedAt}
	m["active"] = []string{strconv.FormatBool(u.Active)}
	m["verified"] = []string{strconv.FormatBool(u.Verified)}
	m["shared"] = []string{strconv.FormatBool(u.Shared)}
	m["locale"] = []string{u.Locale}
	m["timezone"] = []string{u.Timezone}
	m["last_login_at"] = []string{u.LastLoginAt}
	m["email"] = []string{u.Email}
	m["phone"] = []string{u.Phone}
	m["signature"] = []string{u.Signature}
	m["organization_id"] = []string{fmt.Sprintf("%d", u.OrganizationID)}
	m["tags"] = u.Tags
	m["suspended"] = []string{strconv.FormatBool(u.Suspended)}
	m["role"] = []string{u.Role}
	return m
}

// CreateIndex -
func (u *User) CreateIndex(in interface{}, name string) map[string]map[string][]item {
	d := in.([]*User)
	m := make(map[string]map[string][]item)
	m["_id"] = make(map[string][]item)
	m["url"] = make(map[string][]item)
	m["external_id"] = make(map[string][]item)
	m["name"] = make(map[string][]item)
	m["alias"] = make(map[string][]item)
	m["created_at"] = make(map[string][]item)
	m["active"] = make(map[string][]item)
	m["verified"] = make(map[string][]item)
	m["shared"] = make(map[string][]item)
	m["locale"] = make(map[string][]item)
	m["timezone"] = make(map[string][]item)
	m["last_login_at"] = make(map[string][]item)
	m["email"] = make(map[string][]item)
	m["phone"] = make(map[string][]item)
	m["signature"] = make(map[string][]item)
	m["organization_id"] = make(map[string][]item)
	m["tags"] = make(map[string][]item)
	m["suspended"] = make(map[string][]item)
	m["role"] = make(map[string][]item)

	for i := range d {
		m["_id"][fmt.Sprintf("%d", d[i].ID)] = append(m["_id"][fmt.Sprintf("%d", d[i].ID)], d[i])
		m["url"][strings.ToLower(d[i].URL)] = append(m["url"][strings.ToLower(d[i].URL)], d[i])
		m["external_id"][strings.ToLower(d[i].ExternalID)] = append(m["external_id"][strings.ToLower(d[i].ExternalID)], d[i])
		m["name"][strings.ToLower(d[i].Name)] = append(m["name"][strings.ToLower(d[i].Name)], d[i])
		m["alias"][strings.ToLower(d[i].Alias)] = append(m["alias"][strings.ToLower(d[i].Alias)], d[i])
		m["created_at"][strings.ToLower(d[i].CreatedAt)] = append(m["created_at"][strings.ToLower(d[i].CreatedAt)], d[i])
		m["active"][strconv.FormatBool(d[i].Active)] = append(m["active"][strconv.FormatBool(d[i].Active)], d[i])
		m["verified"][strconv.FormatBool(d[i].Verified)] = append(m["verified"][strconv.FormatBool(d[i].Verified)], d[i])
		m["shared"][strconv.FormatBool(d[i].Shared)] = append(m["shared"][strconv.FormatBool(d[i].Shared)], d[i])
		m["locale"][strings.ToLower(d[i].Locale)] = append(m["locale"][strings.ToLower(d[i].Locale)], d[i])
		m["timezone"][strings.ToLower(d[i].Timezone)] = append(m["timezone"][strings.ToLower(d[i].Timezone)], d[i])
		m["last_login_at"][strings.ToLower(d[i].LastLoginAt)] = append(m["last_login_at"][strings.ToLower(d[i].LastLoginAt)], d[i])
		m["email"][strings.ToLower(d[i].Email)] = append(m["email"][strings.ToLower(d[i].Email)], d[i])
		m["phone"][strings.ToLower(d[i].Phone)] = append(m["phone"][strings.ToLower(d[i].Phone)], d[i])
		m["signature"][strings.ToLower(d[i].Signature)] = append(m["signature"][strings.ToLower(d[i].Signature)], d[i])
		m["organization_id"][fmt.Sprintf("%d", d[i].OrganizationID)] = append(m["organization_id"][fmt.Sprintf("%d", d[i].OrganizationID)], d[i])
		for _, tag := range d[i].Tags {
			m["tags"][strings.ToLower(tag)] = append(m["Tags"][strings.ToLower(tag)], d[i])
		}
		m["suspended"][strconv.FormatBool(d[i].Suspended)] = append(m["suspended"][strconv.FormatBool(d[i].Suspended)], d[i])
		m["role"][strings.ToLower(d[i].Role)] = append(m["Role"][strings.ToLower(d[i].Role)], d[i])
	}
	return m
}
