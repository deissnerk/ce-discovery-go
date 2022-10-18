package cediscovery

import "net/url"

type Endpoint struct {
	Header
	Group

	Channel   string `json:"channel,omitempty"`
	Authscope string `json:"authscope,omitempty"`
	Usage     string `json:"usage"`
}

type Group struct {
	Header
	Deprecated  `json:"deprecated,omitempty"`
	Definitions []Definition `json:"definitions,omitempty"`
	Groups      []url.URL    `json:"groups,omitempty"`
}

type Definition struct {
	Header
	Format     string  `json:"format"`
	OwnerGroup url.URL `json:"ownergroup"`
	Metadata   struct {
		Attributes map[string]Attribute `json:"attributes"`
	} `json:"metadata"`
	SchemaUrl url.URL
}

// Common Resource Properties
// https://github.com/cloudevents/spec/blob/d775bc17bca915a18ab3386463a549d45e0630a0/discovery/spec.md#common-resource-properties
type Header struct {
	Id          string            `json:"id"`
	Name        string            `json:"name"`
	Self        url.URL           `json:"self"`
	Origin      url.URL           `json:"origin,omitempty"`
	Epoch       uint64            `json:"epoch"`
	Description string            `json:"description,omitempty"`
	Tags        map[string]string `json:"tags,omitempty"`
	Docs        url.URL           `json:"docs,omitempty"`
}

type Attribute struct {
	Required bool   `json:"required"`
	Value    string `json:"value,omitempty"`
}

type Deprecated struct {
	Effective   string  `json:"effective,omitempty"`
	Removal     string  `json:"removal,omitempty"`
	Alternative url.URL `json:"alternative,omitempty"`
	Docs        url.URL `json:"docs,omitempty"`
}
