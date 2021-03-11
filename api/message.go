package api

import (
	"time"

	"github.com/chebyrash/promise"
)

type Message struct {
	Disgo           Disgo
	ID              Snowflake             `json:"id"`
	GuildId         Snowflake             `json:"guild_id"`
	Reactions       []Reactions           `json:"reactions"`
	Attachments     []interface{}         `json:"attachments"`
	Tts             bool                  `json:"tts"`
	Embeds          []interface{}         `json:"embeds"`
	Timestamp       time.Time             `json:"timestamp"`
	MentionEveryone bool                  `json:"mention_everyone"`
	Pinned          bool                  `json:"pinned"`
	EditedTimestamp interface{}           `json:"edited_timestamp"`
	Author          User                  `json:"author"`
	MentionRoles    []interface{}         `json:"mention_roles"`
	Content         string                `json:"content"`
	ChannelID       Snowflake             `json:"channel_id"`
	Mentions        []interface{}         `json:"mentions"`
	ChannelType     ChannelType `json:"type"`
}

func (m Message) AddReactionByEmote(emote Emote) *promise.Promise {
	return m.AddReaction(emote.Reaction())
}

func (m Message) AddReaction(emoji string) *promise.Promise {
	return m.Disgo.RestClient().AddReaction(m.ChannelID, m.ID, emoji)
}

type Reactions struct {
	Count int   `json:"count"`
	Me    bool  `json:"me"`
	Emoji Emote `json:"emoji"`
}