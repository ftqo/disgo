package discord

import (
	"fmt"

	"github.com/DisgoOrg/disgo/json"
	"github.com/DisgoOrg/snowflake"
)

// ComponentType defines different Component(s)
type ComponentType int

// Supported ComponentType(s)
//goland:noinspection GoUnusedConst
const (
	ComponentTypeActionRow = iota + 1
	ComponentTypeButton
	ComponentTypeSelectMenu
	ComponentTypeInputText
)

type CustomID string

func (c CustomID) String() string {
	return string(c)
}

type Component interface {
	json.Marshaler
	Type() ComponentType
	component()
}

type ContainerComponent interface {
	Component
	Components() []InteractiveComponent
	containerComponent()
}

type InteractiveComponent interface {
	Component
	ID() CustomID
	interactiveComponent()
}

type UnmarshalComponent struct {
	Component
}

func (u *UnmarshalComponent) UnmarshalJSON(data []byte) error {
	var cType struct {
		Type ComponentType `json:"type"`
	}

	if err := json.Unmarshal(data, &cType); err != nil {
		return err
	}

	var (
		component Component
		err       error
	)

	switch cType.Type {
	case ComponentTypeActionRow:
		v := ActionRowComponent{}
		err = json.Unmarshal(data, &v)
		component = v

	case ComponentTypeButton:
		v := ButtonComponent{}
		err = json.Unmarshal(data, &v)
		component = v

	case ComponentTypeSelectMenu:
		v := SelectMenuComponent{}
		err = json.Unmarshal(data, &v)
		component = v

	default:
		err = fmt.Errorf("unkown component with type %d received", cType.Type)
	}
	if err != nil {
		return err
	}

	u.Component = component
	return nil
}

type ComponentEmoji struct {
	ID       snowflake.Snowflake `json:"id,omitempty"`
	Name     string              `json:"name,omitempty"`
	Animated bool                `json:"animated,omitempty"`
}

var (
	_ Component          = (*ActionRowComponent)(nil)
	_ ContainerComponent = (*ActionRowComponent)(nil)
)

//goland:noinspection GoUnusedExportedFunction
func NewActionRow(components ...InteractiveComponent) ActionRowComponent {
	return components
}

type ActionRowComponent []InteractiveComponent

func (c ActionRowComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type       ComponentType          `json:"type"`
		Components []InteractiveComponent `json:"components"`
	}{
		Type:       c.Type(),
		Components: c,
	})
}

func (c *ActionRowComponent) UnmarshalJSON(data []byte) error {
	var actionRow struct {
		Components []UnmarshalComponent `json:"components"`
	}

	if err := json.Unmarshal(data, &actionRow); err != nil {
		return err
	}

	if len(actionRow.Components) > 0 {
		*c = make([]InteractiveComponent, len(actionRow.Components))
		for i, component := range actionRow.Components {
			(*c)[i] = component.Component.(InteractiveComponent)
		}
	}

	return nil
}

func (c ActionRowComponent) Type() ComponentType {
	return ComponentTypeActionRow
}

func (c ActionRowComponent) component()          {}
func (c ActionRowComponent) containerComponent() {}

func (c ActionRowComponent) Components() []InteractiveComponent {
	return c
}

// Buttons returns all ButtonComponent(s) in the ActionRowComponent
func (c ActionRowComponent) Buttons() []ButtonComponent {
	var buttons []ButtonComponent
	for i := range c {
		if button, ok := c[i].(ButtonComponent); ok {
			buttons = append(buttons, button)
		}
	}
	return buttons
}

// SelectMenus returns all SelectMenuComponent(s) in the ActionRowComponent
func (c ActionRowComponent) SelectMenus() []SelectMenuComponent {
	var selectMenus []SelectMenuComponent
	for i := range c {
		if selectMenu, ok := c[i].(SelectMenuComponent); ok {
			selectMenus = append(selectMenus, selectMenu)
		}
	}
	return selectMenus
}

// UpdateComponent returns a new ActionRowComponent with the Component which has the customID replaced
func (c ActionRowComponent) UpdateComponent(customID CustomID, component InteractiveComponent) ActionRowComponent {
	for i, cc := range c {
		if cc.ID() == customID {
			c[i] = component
			return c
		}
	}
	return c
}

// AddComponents returns a new ActionRowComponent with the provided Component(s) added
func (c ActionRowComponent) AddComponents(components ...InteractiveComponent) ActionRowComponent {
	return append(c, components...)
}

// RemoveComponent returns a new ActionRowComponent with the provided Component at the index removed
func (c ActionRowComponent) RemoveComponent(index int) ActionRowComponent {
	if len(c) > index {
		return append(c[:index], c[index+1:]...)
	}
	return c
}

// ButtonStyle defines how the ButtonComponent looks like (https://discord.com/assets/7bb017ce52cfd6575e21c058feb3883b.png)
type ButtonStyle int

// Supported ButtonStyle(s)
const (
	ButtonStylePrimary = iota + 1
	ButtonStyleSecondary
	ButtonStyleSuccess
	ButtonStyleDanger
	ButtonStyleLink
)

// NewButton creates a new ButtonComponent with the provided parameters. Link ButtonComponent(s) need a URL and other ButtonComponent(s) need a customID
//goland:noinspection GoUnusedExportedFunction
func NewButton(style ButtonStyle, label string, customID CustomID, url string) ButtonComponent {
	return ButtonComponent{
		Style:    style,
		CustomID: customID,
		URL:      url,
		Label:    label,
	}
}

// NewPrimaryButton creates a new ButtonComponent with ButtonStylePrimary & the provided parameters
//goland:noinspection GoUnusedExportedFunction
func NewPrimaryButton(label string, customID CustomID) ButtonComponent {
	return NewButton(ButtonStylePrimary, label, customID, "")
}

// NewSecondaryButton creates a new ButtonComponent with ButtonStyleSecondary & the provided parameters
//goland:noinspection GoUnusedExportedFunction
func NewSecondaryButton(label string, customID CustomID) ButtonComponent {
	return NewButton(ButtonStyleSecondary, label, customID, "")
}

// NewSuccessButton creates a new ButtonComponent with ButtonStyleSuccess & the provided parameters
//goland:noinspection GoUnusedExportedFunction
func NewSuccessButton(label string, customID CustomID) ButtonComponent {
	return NewButton(ButtonStyleSuccess, label, customID, "")
}

// NewDangerButton creates a new ButtonComponent with ButtonStyleDanger & the provided parameters
//goland:noinspection GoUnusedExportedFunction
func NewDangerButton(label string, customID CustomID) ButtonComponent {
	return NewButton(ButtonStyleDanger, label, customID, "")
}

// NewLinkButton creates a new link ButtonComponent with ButtonStyleLink & the provided parameters
//goland:noinspection GoUnusedExportedFunction
func NewLinkButton(label string, url string) ButtonComponent {
	return NewButton(ButtonStyleLink, label, "", url)
}

var (
	_ Component            = (*ButtonComponent)(nil)
	_ InteractiveComponent = (*ButtonComponent)(nil)
)

type ButtonComponent struct {
	Style    ButtonStyle     `json:"style"`
	Label    string          `json:"label,omitempty"`
	Emoji    *ComponentEmoji `json:"emoji,omitempty"`
	CustomID CustomID        `json:"custom_id,omitempty"`
	URL      string          `json:"url,omitempty"`
	Disabled bool            `json:"disabled,omitempty"`
}

func (c ButtonComponent) MarshalJSON() ([]byte, error) {
	type buttonComponent ButtonComponent
	return json.Marshal(struct {
		Type ComponentType `json:"type"`
		buttonComponent
	}{
		Type:            c.Type(),
		buttonComponent: buttonComponent(c),
	})
}

func (c ButtonComponent) Type() ComponentType {
	return ComponentTypeButton
}

func (c ButtonComponent) ID() CustomID {
	return c.CustomID
}

func (c ButtonComponent) component()            {}
func (c ButtonComponent) interactiveComponent() {}

// WithStyle returns a new ButtonComponent with the provided style
func (c ButtonComponent) WithStyle(style ButtonStyle) ButtonComponent {
	c.Style = style
	return c
}

// WithLabel returns a new ButtonComponent with the provided label
func (c ButtonComponent) WithLabel(label string) ButtonComponent {
	c.Label = label
	return c
}

// WithEmoji returns a new ButtonComponent with the provided Emoji
func (c ButtonComponent) WithEmoji(emoji ComponentEmoji) ButtonComponent {
	c.Emoji = &emoji
	return c
}

// WithCustomID returns a new ButtonComponent with the provided custom id
func (c ButtonComponent) WithCustomID(customID CustomID) ButtonComponent {
	c.CustomID = customID
	return c
}

// WithURL returns a new ButtonComponent with the provided URL
func (c ButtonComponent) WithURL(url string) ButtonComponent {
	c.URL = url
	return c
}

// AsEnabled returns a new ButtonComponent but enabled
func (c ButtonComponent) AsEnabled() ButtonComponent {
	c.Disabled = false
	return c
}

// AsDisabled returns a new ButtonComponent but disabled
func (c ButtonComponent) AsDisabled() ButtonComponent {
	c.Disabled = true
	return c
}

// WithDisabled returns a new ButtonComponent but disabled/enabled
func (c ButtonComponent) WithDisabled(disabled bool) ButtonComponent {
	c.Disabled = disabled
	return c
}

// NewSelectMenu builds a new SelectMenuComponent from the provided values
//goland:noinspection GoUnusedExportedFunction
func NewSelectMenu(customID CustomID, placeholder string, options ...SelectMenuOption) SelectMenuComponent {
	return SelectMenuComponent{
		CustomID:    customID,
		Placeholder: placeholder,
		Options:     options,
	}
}

var (
	_ Component            = (*SelectMenuComponent)(nil)
	_ InteractiveComponent = (*SelectMenuComponent)(nil)
)

type SelectMenuComponent struct {
	CustomID    CustomID           `json:"custom_id"`
	Placeholder string             `json:"placeholder,omitempty"`
	MinValues   json.NullInt       `json:"min_values,omitempty"`
	MaxValues   json.NullInt       `json:"max_values,omitempty"`
	Disabled    bool               `json:"disabled,omitempty"`
	Options     []SelectMenuOption `json:"options,omitempty"`
}

func (c SelectMenuComponent) MarshalJSON() ([]byte, error) {
	type selectMenuComponent SelectMenuComponent
	return json.Marshal(struct {
		Type ComponentType `json:"type"`
		selectMenuComponent
	}{
		Type:                c.Type(),
		selectMenuComponent: selectMenuComponent(c),
	})
}

func (c SelectMenuComponent) Type() ComponentType {
	return ComponentTypeSelectMenu
}

func (c SelectMenuComponent) ID() CustomID {
	return c.CustomID
}

func (c SelectMenuComponent) component()            {}
func (c SelectMenuComponent) interactiveComponent() {}

// WithCustomID returns a new SelectMenuComponent with the provided customID
func (c SelectMenuComponent) WithCustomID(customID CustomID) SelectMenuComponent {
	c.CustomID = customID
	return c
}

// WithPlaceholder returns a new SelectMenuComponent with the provided placeholder
func (c SelectMenuComponent) WithPlaceholder(placeholder string) SelectMenuComponent {
	c.Placeholder = placeholder
	return c
}

// WithMinValues returns a new SelectMenuComponent with the provided minValue
func (c SelectMenuComponent) WithMinValues(minValue int) SelectMenuComponent {
	c.MinValues = *json.NewInt(minValue)
	return c
}

// WithMaxValues returns a new SelectMenuComponent with the provided maxValue
func (c SelectMenuComponent) WithMaxValues(maxValue int) SelectMenuComponent {
	c.MaxValues = *json.NewInt(maxValue)
	return c
}

// AsEnabled returns a new SelectMenuComponent but enabled
func (c SelectMenuComponent) AsEnabled() SelectMenuComponent {
	c.Disabled = false
	return c
}

// AsDisabled returns a new SelectMenuComponent but disabled
func (c SelectMenuComponent) AsDisabled() SelectMenuComponent {
	c.Disabled = true
	return c
}

// WithDisabled returns a new SelectMenuComponent with the provided disabled
func (c SelectMenuComponent) WithDisabled(disabled bool) SelectMenuComponent {
	c.Disabled = disabled
	return c
}

// SetOptions returns a new SelectMenuComponent with the provided SelectMenuOption(s)
func (c SelectMenuComponent) SetOptions(options ...SelectMenuOption) SelectMenuComponent {
	c.Options = options
	return c
}

// SetOption returns a new SelectMenuComponent with the SelectMenuOption which has the value replaced
func (c SelectMenuComponent) SetOption(value string, option SelectMenuOption) SelectMenuComponent {
	for i, o := range c.Options {
		if o.Value == value {
			c.Options[i] = option
			break
		}
	}
	return c
}

// AddOptions returns a new SelectMenuComponent with the provided SelectMenuOption(s) added
func (c SelectMenuComponent) AddOptions(options ...SelectMenuOption) SelectMenuComponent {
	c.Options = append(c.Options, options...)
	return c
}

// RemoveOption returns a new SelectMenuComponent with the provided SelectMenuOption at the index removed
func (c SelectMenuComponent) RemoveOption(index int) SelectMenuComponent {
	if len(c.Options) > index {
		c.Options = append(c.Options[:index], c.Options[index+1:]...)
	}
	return c
}

// NewSelectMenuOption builds a new SelectMenuOption
//goland:noinspection GoUnusedExportedFunction
func NewSelectMenuOption(label string, value string) SelectMenuOption {
	return SelectMenuOption{
		Label: label,
		Value: value,
	}
}

// SelectMenuOption represents an option in a SelectMenuComponent
type SelectMenuOption struct {
	Label       string          `json:"label"`
	Value       string          `json:"value"`
	Description string          `json:"description,omitempty"`
	Emoji       *ComponentEmoji `json:"emoji,omitempty"`
	Default     bool            `json:"default,omitempty"`
}

// WithLabel returns a new SelectMenuOption with the provided label
func (o SelectMenuOption) WithLabel(label string) SelectMenuOption {
	o.Label = label
	return o
}

// WithValue returns a new SelectMenuOption with the provided value
func (o SelectMenuOption) WithValue(value string) SelectMenuOption {
	o.Value = value
	return o
}

// WithDescription returns a new SelectMenuOption with the provided description
func (o SelectMenuOption) WithDescription(description string) SelectMenuOption {
	o.Description = description
	return o
}

// WithEmoji returns a new SelectMenuOption with the provided Emoji
func (o SelectMenuOption) WithEmoji(emoji ComponentEmoji) SelectMenuOption {
	o.Emoji = &emoji
	return o
}

// WithDefault returns a new SelectMenuOption as default/non-default
func (o SelectMenuOption) WithDefault(defaultOption bool) SelectMenuOption {
	o.Default = defaultOption
	return o
}

var (
	_ Component            = (*InputText)(nil)
	_ InteractiveComponent = (*InputText)(nil)
)

type InputText struct {
	TextStyle   TextStyle    `json:"text_style"`
	CustomID    CustomID     `json:"custom_id"`
	Label       string       `json:"label"`
	Placeholder string       `json:"placeholder,omitempty"`
	MinLength   json.NullInt `json:"min_length,omitempty"`
	MaxLength   json.NullInt `json:"max_length,omitempty"`
}

func (t InputText) MarshalJSON() ([]byte, error) {
	type inputText InputText
	return json.Marshal(struct {
		Type ComponentType `json:"type"`
		inputText
	}{
		Type:      t.Type(),
		inputText: inputText(t),
	})
}

func (t InputText) Type() ComponentType {
	return ComponentTypeInputText
}

func (t InputText) ID() CustomID {
	return t.CustomID
}

func (t InputText) component()            {}
func (t InputText) interactiveComponent() {}

type TextStyle int

//goland:noinspection GoUnusedConst
const (
	TextStyleShort = iota + 1
	TextStyleParagraph
)
