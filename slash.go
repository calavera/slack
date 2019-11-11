package slack

import (
	"net/http"
	"net/url"
)

// SlashCommand contains information about a request of the slash command
type SlashCommand struct {
	Token          string `json:"token"`
	TeamID         string `json:"team_id"`
	TeamDomain     string `json:"team_domain"`
	EnterpriseID   string `json:"enterprise_id,omitempty"`
	EnterpriseName string `json:"enterprise_name,omitempty"`
	ChannelID      string `json:"channel_id"`
	ChannelName    string `json:"channel_name"`
	UserID         string `json:"user_id"`
	UserName       string `json:"user_name"`
	Command        string `json:"command"`
	Text           string `json:"text"`
	ResponseURL    string `json:"response_url"`
	TriggerID      string `json:"trigger_id"`
}

// SlashCommandParse will parse the request of the slash command
func SlashCommandParse(r *http.Request) (s SlashCommand, err error) {
	if err = r.ParseForm(); err != nil {
		return s, err
	}
	
	return SlashCommandParseValues(r.PostForm), nil
}

func SlashCommandParseValues(values url.Values) SlashCommand {
	var s SlashCommand
	
	s.Token = values.Get("token")
	s.TeamID = values.Get("team_id")
	s.TeamDomain = values.Get("team_domain")
	s.EnterpriseID = values.Get("enterprise_id")
	s.EnterpriseName = values.Get("enterprise_name")
	s.ChannelID = values.Get("channel_id")
	s.ChannelName = values.Get("channel_name")
	s.UserID = values.Get("user_id")
	s.UserName = values.Get("user_name")
	s.Command = values.Get("command")
	s.Text = values.Get("text")
	s.ResponseURL = values.Get("response_url")
	s.TriggerID = values.Get("trigger_id")
	
	return s
}

// ValidateToken validates verificationTokens
func (s SlashCommand) ValidateToken(verificationTokens ...string) bool {
	for _, token := range verificationTokens {
		if s.Token == token {
			return true
		}
	}
	return false
}
