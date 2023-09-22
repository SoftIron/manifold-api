package cloud

import (
	"encoding/json"
	"encoding/xml"
)

// RawTemplate returns the XML representation of c.Template.
func (c *Cluster) RawTemplate() (string, error) {
	return marshalXML(c.Template)
}

// MarshalJSON implements the json.Marshaler interface for c. It appends an XML
// version of the template.
func (c *Cluster) MarshalJSON() ([]byte, error) {
	type Alias Cluster

	raw, err := c.RawTemplate()
	if err != nil {
		return nil, err
	}

	cluster := struct {
		*Alias
		RawTemplate string `json:"raw_template"`
	}{
		Alias:       (*Alias)(c),
		RawTemplate: raw,
	}

	return json.Marshal(cluster)
}

// MarshalYAML implements the yaml.Marshaler interface for c. It appends an XML
// version of the template.
func (c *Cluster) MarshalYAML() (interface{}, error) {
	type Alias Cluster

	raw, err := c.RawTemplate()
	if err != nil {
		return nil, err
	}

	cluster := struct {
		*Alias      `yaml:",inline"`
		RawTemplate string `yaml:"raw_template"`
	}{
		Alias:       (*Alias)(c),
		RawTemplate: raw,
	}

	return cluster, nil
}

// RawTemplate returns the XML representation of d.Template.
func (d *Datastore) RawTemplate() (string, error) {
	return marshalXML(d.Template)
}

// MarshalJSON implements the json.Marshaler interface for d. It appends an XML
// version of the template.
func (d *Datastore) MarshalJSON() ([]byte, error) {
	type Alias Datastore

	raw, err := d.RawTemplate()
	if err != nil {
		return nil, err
	}

	datastore := struct {
		*Alias
		RawTemplate string `json:"raw_template"`
	}{
		Alias:       (*Alias)(d),
		RawTemplate: raw,
	}

	return json.Marshal(datastore)
}

// MarshalYAML implements the yaml.Marshaler interface for d. It appends an XML
// version of the template.
func (d *Datastore) MarshalYAML() (interface{}, error) {
	type Alias Datastore

	raw, err := d.RawTemplate()
	if err != nil {
		return nil, err
	}

	datastore := struct {
		*Alias      `yaml:",inline"`
		RawTemplate string `yaml:"raw_template"`
	}{
		Alias:       (*Alias)(d),
		RawTemplate: raw,
	}

	return datastore, nil
}

// RawTemplate returns the XML representation of g.Template.
func (g *Group) RawTemplate() (string, error) {
	return marshalXML(g.Template)
}

// MarshalJSON implements the json.Marshaler interface for g. It appends an XML
// version of the template.
func (g *Group) MarshalJSON() ([]byte, error) {
	type Alias Group

	raw, err := g.RawTemplate()
	if err != nil {
		return nil, err
	}

	group := struct {
		*Alias
		RawTemplate string `json:"raw_template"`
	}{
		Alias:       (*Alias)(g),
		RawTemplate: raw,
	}

	return json.Marshal(group)
}

// MarshalYAML implements the yaml.Marshaler interface for g. It appends an XML
// version of the template.
func (g *Group) MarshalYAML() (interface{}, error) {
	type Alias Group

	raw, err := g.RawTemplate()
	if err != nil {
		return nil, err
	}

	group := struct {
		*Alias      `yaml:",inline"`
		RawTemplate string `yaml:"raw_template"`
	}{
		Alias:       (*Alias)(g),
		RawTemplate: raw,
	}

	return group, nil
}

// RawTemplate returns the XML representation of h.Template.
func (h *Host) RawTemplate() (string, error) {
	return marshalXML(h.Template)
}

// MarshalJSON implements the json.Marshaler interface for h. It appends an XML
// version of the template.
func (h *Host) MarshalJSON() ([]byte, error) {
	type Alias Host

	raw, err := h.RawTemplate()
	if err != nil {
		return nil, err
	}

	host := struct {
		*Alias
		RawTemplate string `json:"raw_template"`
	}{
		Alias:       (*Alias)(h),
		RawTemplate: raw,
	}

	return json.Marshal(host)
}

// MarshalYAML implements the yaml.Marshaler interface for h. It appends an XML
// version of the template.
func (h *Host) MarshalYAML() (interface{}, error) {
	type Alias Host

	raw, err := h.RawTemplate()
	if err != nil {
		return nil, err
	}

	host := struct {
		*Alias      `yaml:",inline"`
		RawTemplate string `yaml:"raw_template"`
	}{
		Alias:       (*Alias)(h),
		RawTemplate: raw,
	}

	return host, nil
}

// RawTemplate returns the XML representation of i.Template.
func (i *Image) RawTemplate() (string, error) {
	return marshalXML(i.Template)
}

// MarshalJSON implements the json.Marshaler interface for i. It appends an XML
// version of the template.
func (i *Image) MarshalJSON() ([]byte, error) {
	type Alias Image

	raw, err := i.RawTemplate()
	if err != nil {
		return nil, err
	}

	image := struct {
		*Alias
		RawTemplate string `json:"raw_template"`
	}{
		Alias:       (*Alias)(i),
		RawTemplate: raw,
	}

	return json.Marshal(image)
}

// MarshalYAML implements the yaml.Marshaler interface for i. It appends an XML
// version of the template.
func (i *Image) MarshalYAML() (interface{}, error) {
	type Alias Image

	raw, err := i.RawTemplate()
	if err != nil {
		return nil, err
	}

	image := struct {
		*Alias      `yaml:",inline"`
		RawTemplate string `yaml:"raw_template"`
	}{
		Alias:       (*Alias)(i),
		RawTemplate: raw,
	}

	return image, nil
}

// RawTemplate returns the XML representation of l.Template.
func (l *LockedInstance) RawTemplate() (string, error) {
	return marshalXML(l.Template)
}

// MarshalJSON implements the json.Marshaler interface for l. It appends an XML
// version of the template.
func (l *LockedInstance) MarshalJSON() ([]byte, error) {
	type Alias LockedInstance

	raw, err := l.RawTemplate()
	if err != nil {
		return nil, err
	}

	inst := struct {
		*Alias
		RawTemplate string `json:"raw_template"`
	}{
		Alias:       (*Alias)(l),
		RawTemplate: raw,
	}

	return json.Marshal(inst)
}

// MarshalYAML implements the yaml.Marshaler interface for l. It appends an XML
// version of the template.
func (l *LockedInstance) MarshalYAML() (interface{}, error) {
	type Alias LockedInstance

	raw, err := l.RawTemplate()
	if err != nil {
		return nil, err
	}

	inst := struct {
		*Alias      `yaml:",inline"`
		RawTemplate string `yaml:"raw_template"`
	}{
		Alias:       (*Alias)(l),
		RawTemplate: raw,
	}

	return inst, nil
}

// RawTemplate returns the XML representation of u.Template.
func (u *User) RawTemplate() (string, error) {
	return marshalXML(u.Template)
}

// MarshalJSON implements the json.Marshaler interface for u. It appends an XML
// version of the template.
func (u *User) MarshalJSON() ([]byte, error) {
	type Alias User

	raw, err := u.RawTemplate()
	if err != nil {
		return nil, err
	}

	user := struct {
		*Alias
		RawTemplate string `json:"raw_template"`
	}{
		Alias:       (*Alias)(u),
		RawTemplate: raw,
	}

	return json.Marshal(user)
}

// MarshalYAML implements the yaml.Marshaler interface for u. It appends an XML
// version of the template.
func (u *User) MarshalYAML() (interface{}, error) {
	type Alias User

	raw, err := u.RawTemplate()
	if err != nil {
		return nil, err
	}

	user := struct {
		*Alias      `yaml:",inline"`
		RawTemplate string `yaml:"raw_template"`
	}{
		Alias:       (*Alias)(u),
		RawTemplate: raw,
	}

	return user, nil
}

// RawTemplate returns the XML representation of v.Template.
func (v *DataCenter) RawTemplate() (string, error) {
	return marshalXML(v.Template)
}

// MarshalJSON implements the json.Marshaler interface for v. It appends an XML
// version of the template.
func (v *DataCenter) MarshalJSON() ([]byte, error) {
	type Alias DataCenter

	raw, err := v.RawTemplate()
	if err != nil {
		return nil, err
	}

	vdc := struct {
		*Alias
		RawTemplate string `json:"raw_template"`
	}{
		Alias:       (*Alias)(v),
		RawTemplate: raw,
	}

	return json.Marshal(vdc)
}

// MarshalYAML implements the yaml.Marshaler interface for v. It appends an XML
// version of the template.
func (v *DataCenter) MarshalYAML() (interface{}, error) {
	type Alias DataCenter

	raw, err := v.RawTemplate()
	if err != nil {
		return nil, err
	}

	vdc := struct {
		*Alias      `yaml:",inline"`
		RawTemplate string `yaml:"raw_template"`
	}{
		Alias:       (*Alias)(v),
		RawTemplate: raw,
	}

	return vdc, nil
}

// RawTemplate returns the XML representation of v.Template.
func (v *Network) RawTemplate() (string, error) {
	return marshalXML(v.Template)
}

// MarshalJSON implements the json.Marshaler interface for v. It appends an XML
// version of the template.
func (v *Network) MarshalJSON() ([]byte, error) {
	type Alias Network

	raw, err := v.RawTemplate()
	if err != nil {
		return nil, err
	}

	vnet := struct {
		*Alias
		RawTemplate string `json:"raw_template"`
	}{
		Alias:       (*Alias)(v),
		RawTemplate: raw,
	}

	return json.Marshal(vnet)
}

// MarshalYAML implements the yaml.Marshaler interface for v. It appends an XML
// version of the template.
func (v *Network) MarshalYAML() (interface{}, error) {
	type Alias Network

	raw, err := v.RawTemplate()
	if err != nil {
		return nil, err
	}

	vnet := struct {
		*Alias      `yaml:",inline"`
		RawTemplate string `yaml:"raw_template"`
	}{
		Alias:       (*Alias)(v),
		RawTemplate: raw,
	}

	return vnet, nil
}

// RawTemplate returns the XML representation of z.Template.
func (z *Zone) RawTemplate() (string, error) {
	return marshalXML(z.Template)
}

// MarshalJSON implements the json.Marshaler interface for z. It appends an XML
// version of the template.
func (z *Zone) MarshalJSON() ([]byte, error) {
	type Alias Zone

	raw, err := z.RawTemplate()
	if err != nil {
		return nil, err
	}

	zone := struct {
		*Alias
		RawTemplate string `json:"raw_template"`
	}{
		Alias:       (*Alias)(z),
		RawTemplate: raw,
	}

	return json.Marshal(zone)
}

// MarshalYAML implements the yaml.Marshaler interface for z. It appends an XML
// version of the template.
func (z *Zone) MarshalYAML() (interface{}, error) {
	type Alias Zone

	raw, err := z.RawTemplate()
	if err != nil {
		return nil, err
	}

	zone := struct {
		*Alias      `yaml:",inline"`
		RawTemplate string `yaml:"raw_template"`
	}{
		Alias:       (*Alias)(z),
		RawTemplate: raw,
	}

	return zone, nil
}

func marshalXML(v any) (string, error) {
	raw, err := xml.MarshalIndent(v, "", "    ")
	if err != nil {
		return "", err
	}

	return string(raw), nil
}
