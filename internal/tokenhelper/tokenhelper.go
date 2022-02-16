package tokenhelper

import (
	"encoding/base64"
	"strings"

	"github.com/DisgoOrg/disgo/discord"
	"github.com/DisgoOrg/snowflake"
)

// IDFromToken returns the applicationID from the BotToken
//goland:noinspection GoUnusedExportedFunction
func IDFromToken(token string) (*snowflake.Snowflake, error) {
	strs := strings.Split(token, ".")
	if len(strs) == 0 {
		return nil, discord.ErrInvalidBotToken
	}
	byteID, err := base64.StdEncoding.DecodeString(strs[0])
	if err != nil {
		return nil, err
	}
	strID := snowflake.Snowflake(byteID)
	return &strID, nil
}