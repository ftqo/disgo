package discord

// GatewayEventType wraps all GatewayEventType types
type GatewayEventType string

// Constants for the gateway events
const (
	GatewayEventTypeReady                         GatewayEventType = "READY"
	GatewayEventTypeResumed                       GatewayEventType = "RESUMED"
	GatewayEventTypeChannelCreate                 GatewayEventType = "CHANNEL_CREATE"
	GatewayEventTypeChannelUpdate                 GatewayEventType = "CHANNEL_UPDATE"
	GatewayEventTypeChannelDelete                 GatewayEventType = "CHANNEL_DELETE"
	GatewayEventTypeChannelPinsUpdate             GatewayEventType = "CHANNEL_PINS_UPDATE"
	GatewayEventTypeThreadCreate                  GatewayEventType = "THREAD_CREATE"
	GatewayEventTypeThreadUpdate                  GatewayEventType = "THREAD_UPDATE"
	GatewayEventTypeThreadDelete                  GatewayEventType = "THREAD_DELETE"
	GatewayEventTypeThreadListSync                GatewayEventType = "THREAD_LIST_SYNC"
	GatewayEventTypeThreadMemberUpdate            GatewayEventType = "THREAD_MEMBER_UPDATE"
	GatewayEventTypeThreadMembersUpdate           GatewayEventType = "THREAD_MEMBERS_UPDATE"
	GatewayEventTypeGuildCreate                   GatewayEventType = "GUILD_CREATE"
	GatewayEventTypeGuildUpdate                   GatewayEventType = "GUILD_UPDATE"
	GatewayEventTypeGuildDelete                   GatewayEventType = "GUILD_DELETE"
	GatewayEventTypeGuildBanAdd                   GatewayEventType = "GUILD_BAN_ADD"
	GatewayEventTypeGuildBanRemove                GatewayEventType = "GUILD_BAN_REMOVE"
	GatewayEventTypeGuildEmojisUpdate             GatewayEventType = "GUILD_EMOJIS_UPDATE"
	GatewayEventTypeGuildStickersUpdate           GatewayEventType = "GUILD_STICKERS_UPDATE"
	GatewayEventTypeGuildIntegrationsUpdate       GatewayEventType = "GUILD_INTEGRATIONS_UPDATE"
	GatewayEventTypeGuildMemberAdd                GatewayEventType = "GUILD_MEMBER_ADD"
	GatewayEventTypeGuildMemberRemove             GatewayEventType = "GUILD_MEMBER_REMOVE"
	GatewayEventTypeGuildMemberUpdate             GatewayEventType = "GUILD_MEMBER_UPDATE"
	GatewayEventTypeGuildMembersChunk             GatewayEventType = "GUILD_MEMBERS_CHUNK"
	GatewayEventTypeGuildRoleCreate               GatewayEventType = "GUILD_ROLE_CREATE"
	GatewayEventTypeGuildRoleUpdate               GatewayEventType = "GUILD_ROLE_UPDATE"
	GatewayEventTypeGuildRoleDelete               GatewayEventType = "GUILD_ROLE_DELETE"
	GatewayEventTypeGuildScheduledEventCreate     GatewayEventType = "GUILD_SCHEDULED_EVENT_CREATE"
	GatewayEventTypeGuildScheduledEventUpdate     GatewayEventType = "GUILD_SCHEDULED_EVENT_UPDATE"
	GatewayEventTypeGuildScheduledEventDelete     GatewayEventType = "GUILD_SCHEDULED_EVENT_DELETE"
	GatewayEventTypeGuildScheduledEventUserAdd    GatewayEventType = "GUILD_SCHEDULED_EVENT_USER_ADD"
	GatewayEventTypeGuildScheduledEventUserRemove GatewayEventType = "GUILD_SCHEDULED_EVENT_USER_REMOVE"
	GatewayEventTypeIntegrationCreate             GatewayEventType = "INTEGRATION_CREATE"
	GatewayEventTypeIntegrationUpdate             GatewayEventType = "INTEGRATION_UPDATE"
	GatewayEventTypeIntegrationDelete             GatewayEventType = "INTEGRATION_DELETE"
	GatewayEventTypeInteractionCreate             GatewayEventType = "INTERACTION_CREATE"
	GatewayEventTypeInviteCreate                  GatewayEventType = "INVITE_CREATE"
	GatewayEventTypeInviteDelete                  GatewayEventType = "INVITE_DELETE"
	GatewayEventTypeMessageCreate                 GatewayEventType = "MESSAGE_CREATE"
	GatewayEventTypeMessageUpdate                 GatewayEventType = "MESSAGE_UPDATE"
	GatewayEventTypeMessageDelete                 GatewayEventType = "MESSAGE_DELETE"
	GatewayEventTypeMessageDeleteBulk             GatewayEventType = "MESSAGE_DELETE_BULK"
	GatewayEventTypeMessageReactionAdd            GatewayEventType = "MESSAGE_REACTION_ADD"
	GatewayEventTypeMessageReactionRemove         GatewayEventType = "MESSAGE_REACTION_REMOVE"
	GatewayEventTypeMessageReactionRemoveAll      GatewayEventType = "MESSAGE_REACTION_REMOVE_ALL"
	GatewayEventTypeMessageReactionRemoveEmoji    GatewayEventType = "MESSAGE_REACTION_REMOVE_EMOJI"
	GatewayEventTypePresenceUpdate                GatewayEventType = "PRESENCE_UPDATE"
	GatewayEventTypeStageInstanceCreate           GatewayEventType = "STAGE_INSTANCE_CREATE"
	GatewayEventTypeStageInstanceDelete           GatewayEventType = "STAGE_INSTANCE_DELETE"
	GatewayEventTypeStageInstanceUpdate           GatewayEventType = "STAGE_INSTANCE_UPDATE"
	GatewayEventTypeTypingStart                   GatewayEventType = "TYPING_START"
	GatewayEventTypeUserUpdate                    GatewayEventType = "USER_UPDATE"
	GatewayEventTypeVoiceStateUpdate              GatewayEventType = "VOICE_STATE_UPDATE"
	GatewayEventTypeVoiceServerUpdate             GatewayEventType = "VOICE_SERVER_UPDATE"
	GatewayEventTypeWebhooksUpdate                GatewayEventType = "WEBHOOKS_UPDATE"
)
