package singularity

import "sync"

//RTMResp is the response given by slack when we authenticate.
//TODO Create functions to retrieve each field.
type RTMResp struct {
	sync.Mutex
	OK                      bool        `json:"ok"`
	URL                     string      `json:"url"`
	Self                    Self        `json:"self"`
	Team                    Team        `json:"team"`
	LatestEventTs           string      `json:"latest_event_ts"`
	Channels                []Channel   `json:"channels"`
	Groups                  []Channel   `json:"groups"`
	IMS                     []IM        `json:"ims"`
	Users                   []User      `json:"users"`
	Bots                    []Bot       `json:"bots"`
	DND                     DND         `json:"dnd"`
	Subteams                SubTeamList `json:"subteams"`
	CacheTs                 int         `json:"cache_ts"`
	ReadOnlyChannels        []string    `json:"read_only_channels"`
	CanManageSharedChannels bool        `json:"can_manage_shared_channels"`
	CacheVersion            string      `json:"cache_version"`
	CacheTsVersion          string      `json:"cache_ts_version"`
}

//Bot represents a bot
type Bot struct {
	ID      string `json:"id"`
	Deleted bool   `json:"deleted"`
	Name    string `json:"name"`
	Icons   Icon   `json:"icons"`
}

//SubTeamList ...
type SubTeamList struct {
	Self []interface{} `json:"self"`
	All  []interface{} `json:"all"`
}

//DND ...
type DND struct {
	DndEnabled     bool `json:"dnd_enabled"`
	NextDndStartTs int  `json:"next_dnd_start_ts"`
	NextDndEndTs   int  `json:"next_dnd_end_ts"`
	SnoozeEnabled  bool `json:"snooze_enabled"`
}

//Latest ...
type Latest struct {
	Type string `json:"type"`
	User string `json:"user"`
	Text string `json:"text"`
	Ts   string `json:"ts"`
}

//IM ...
type IM struct {
	ID                 string `json:"id"`
	User               string `json:"user"`
	Created            int    `json:"created"`
	IsIm               bool   `json:"is_im"`
	IsOrgShared        bool   `json:"is_org_shared"`
	HasPins            bool   `json:"has_pins"`
	LastRead           string `json:"last_read"`
	Latest             Latest `json:"latest"`
	UnreadCount        int    `json:"unread_count"`
	UnreadCountDisplay int    `json:"unread_count_display"`
	IsOpen             bool   `json:"is_open"`
}

//Channel ...
type Channel struct {
	ID                 string   `json:"id"`
	Name               string   `json:"name"`
	IsGroup            bool     `json:"is_group"`
	Created            int      `json:"created"`
	Creator            string   `json:"creator"`
	IsArchived         bool     `json:"is_archived"`
	IsMpim             bool     `json:"is_mpim"`
	HasPins            bool     `json:"has_pins"`
	IsOpen             bool     `json:"is_open"`
	LastRead           string   `json:"last_read"`
	Latest             Latest   `json:"latest"`
	UnreadCount        int      `json:"unread_count"`
	UnreadCountDisplay int      `json:"unread_count_display"`
	Members            []string `json:"members"`
	Topic              Topic    `json:"topic"`
	Purpose            Purpose  `json:"purpose"`
	IsChannel          bool     `json:"is_channel"`
	IsGeneral          bool     `json:"is_general"`
	IsMember           bool     `json:"is_member"`
}

//Purpose ...
type Purpose struct {
	Value   string `json:"value"`
	Creator string `json:"creator"`
	LastSet int    `json:"last_set"`
}

//Topic ...
type Topic struct {
	Value   string `json:"value"`
	Creator string `json:"creator"`
	LastSet int    `json:"last_set"`
}

//Team ...
type Team struct {
	ID                    string      `json:"id"`
	Name                  string      `json:"name"`
	EmailDomain           string      `json:"email_domain"`
	Domain                string      `json:"domain"`
	MsgEditWindowMins     int         `json:"msg_edit_window_mins"`
	Preferences           Preferences `json:"prefs"`
	Icon                  Icon        `json:"icon"`
	OverStorageLimit      bool        `json:"over_storage_limit"`
	Plan                  string      `json:"plan"`
	AvatarBaseURL         string      `json:"avatar_base_url"`
	OverIntegrationsLimit bool        `json:"over_integrations_limit"`
}

//Icon ...
type Icon map[string]string

//Preferences ...
type Preferences struct {
	DefaultChannels        []string `json:"default_channels"`
	InvitesOnlyAdmins      bool     `json:"invites_only_admins"`
	AllowCalls             bool     `json:"allow_calls"`
	DisplayEmailAddresses  bool     `json:"display_email_addresses"`
	HideReferers           bool     `json:"hide_referers"`
	MsgEditWindowMins      int      `json:"msg_edit_window_mins"`
	AllowMessageDeletion   bool     `json:"allow_message_deletion"`
	CallingAppName         string   `json:"calling_app_name"`
	DisplayRealNames       bool     `json:"display_real_names"`
	WhoCanAtEveryone       string   `json:"who_can_at_everyone"`
	WhoCanAtChannel        string   `json:"who_can_at_channel"`
	WhoCanCreateChannels   string   `json:"who_can_create_channels"`
	WhoCanArchiveChannels  string   `json:"who_can_archive_channels"`
	WhoCanCreateGroups     string   `json:"who_can_create_groups"`
	WhoCanPostGeneral      string   `json:"who_can_post_general"`
	WhoCanKickChannels     string   `json:"who_can_kick_channels"`
	WhoCanKickGroups       string   `json:"who_can_kick_groups"`
	RetentionType          int      `json:"retention_type"`
	RetentionDuration      int      `json:"retention_duration"`
	GroupRetentionType     int      `json:"group_retention_type"`
	GroupRetentionDuration int      `json:"group_retention_duration"`
	DmRetentionType        int      `json:"dm_retention_type"`
	DmRetentionDuration    int      `json:"dm_retention_duration"`
	FileRetentionDuration  int      `json:"file_retention_duration"`
	FileRetentionType      int      `json:"file_retention_type"`
	AllowRetentionOverride bool     `json:"allow_retention_override"`
	RequireAtForMention    bool     `json:"require_at_for_mention"`
	DefaultRxns            []string `json:"default_rxns"`
	TeamHandyRxns          struct {
		Restrict bool `json:"restrict"`
		List     []struct {
			Name  string `json:"name"`
			Title string `json:"title"`
		} `json:"list"`
	} `json:"team_handy_rxns"`
	ChannelHandyRxns                interface{} `json:"channel_handy_rxns"`
	ComplianceExportStart           int         `json:"compliance_export_start"`
	WarnBeforeAtChannel             string      `json:"warn_before_at_channel"`
	DisallowPublicFileUrls          bool        `json:"disallow_public_file_urls"`
	WhoCanCreateDeleteUserGroups    string      `json:"who_can_create_delete_user_groups"`
	WhoCanEditUserGroups            string      `json:"who_can_edit_user_groups"`
	WhoCanChangeTeamProfile         string      `json:"who_can_change_team_profile"`
	AllowSharedChannels             bool        `json:"allow_shared_channels"`
	WhoHasTeamVisibility            string      `json:"who_has_team_visibility"`
	DisableFileUploads              string      `json:"disable_file_uploads"`
	DisableFileEditing              bool        `json:"disable_file_editing"`
	DisableFileDeleting             bool        `json:"disable_file_deleting"`
	WhoCanCreateSharedChannels      string      `json:"who_can_create_shared_channels"`
	WhoCanManageSharedChannels      Types       `json:"who_can_manage_shared_channels"`
	WhoCanPostInSharedChannels      Types       `json:"who_can_post_in_shared_channels"`
	AllowSharedChannelPermsOverride bool        `json:"allow_shared_channel_perms_override"`
	DndEnabled                      bool        `json:"dnd_enabled"`
	DndStartHour                    string      `json:"dnd_start_hour"`
	DndEndHour                      string      `json:"dnd_end_hour"`
	AuthMode                        string      `json:"auth_mode"`
	WhoCanManageIntegrations        Types       `json:"who_can_manage_integrations"`
	InvitesLimit                    bool        `json:"invites_limit"`
}

//Types ...
type Types struct {
	Type []string `json:"type"`
}

//Self ...
type Self struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	Created        int      `json:"created"`
	ManualPresence string   `json:"manual_presence"`
	Preferences    SelfPref `json:"prefs"`
}

//User ...
type User struct {
	ID                string      `json:"id"`
	Name              string      `json:"name"`
	Deleted           bool        `json:"deleted"`
	Color             string      `json:"color"`
	Profile           Profile     `json:"profile"`
	IsAdmin           bool        `json:"is_admin"`
	IsOwner           bool        `json:"is_owner"`
	IsPrimaryOwner    bool        `json:"is_primary_owner"`
	IsRestricted      bool        `json:"is_restricted"`
	IsUltraRestricted bool        `json:"is_ultra_restricted"`
	Has2Fa            bool        `json:"has_2fa"`
	TwoFactorType     string      `json:"two_factor_type"`
	HasFiles          bool        `json:"has_files"`
	TeamID            string      `json:"team_id"`
	Status            interface{} `json:"status,omitempty"`
	RealName          string      `json:"real_name,omitempty"`
	Tz                string      `json:"tz,omitempty"`
	TzLabel           string      `json:"tz_label,omitempty"`
	TzOffset          int         `json:"tz_offset,omitempty"`
	IsBot             bool        `json:"is_bot"`
	Presence          string      `json:"presence"`
}

//Profile ...
type Profile struct {
	FirstName          string      `json:"first_name"`
	LastName           string      `json:"last_name"`
	AvatarHash         string      `json:"avatar_hash"`
	RealName           string      `json:"real_name"`
	RealNameNormalized string      `json:"real_name_normalized"`
	Email              string      `json:"email"`
	Image24            string      `json:"image_24"`
	Image32            string      `json:"image_32"`
	Image48            string      `json:"image_48"`
	Image72            string      `json:"image_72"`
	Image192           string      `json:"image_192"`
	Image512           string      `json:"image_512"`
	Fields             interface{} `json:"fields"`
	Skype              string      `json:"skype"`
	Phone              string      `json:"phone"`
}

//SelfPref ...
type SelfPref struct {
	HighlightWords                     string      `json:"highlight_words"`
	UserColors                         string      `json:"user_colors"`
	ColorNamesInList                   bool        `json:"color_names_in_list"`
	GrowlsEnabled                      bool        `json:"growls_enabled"`
	Tz                                 string      `json:"tz"`
	PushDmAlert                        bool        `json:"push_dm_alert"`
	PushMentionAlert                   bool        `json:"push_mention_alert"`
	MsgReplies                         string      `json:"msg_replies"`
	PushEverything                     bool        `json:"push_everything"`
	PushIdleWait                       int         `json:"push_idle_wait"`
	PushSound                          string      `json:"push_sound"`
	PushLoudChannels                   string      `json:"push_loud_channels"`
	PushMentionChannels                string      `json:"push_mention_channels"`
	PushLoudChannelsSet                string      `json:"push_loud_channels_set"`
	EmailAlerts                        string      `json:"email_alerts"`
	EmailAlertsSleepUntil              int         `json:"email_alerts_sleep_until"`
	EmailMisc                          bool        `json:"email_misc"`
	EmailWeekly                        bool        `json:"email_weekly"`
	WelcomeMessageHidden               bool        `json:"welcome_message_hidden"`
	AllChannelsLoud                    bool        `json:"all_channels_loud"`
	LoudChannels                       string      `json:"loud_channels"`
	NeverChannels                      string      `json:"never_channels"`
	LoudChannelsSet                    string      `json:"loud_channels_set"`
	SearchSort                         string      `json:"search_sort"`
	ExpandInlineImgs                   bool        `json:"expand_inline_imgs"`
	ExpandInternalInlineImgs           bool        `json:"expand_internal_inline_imgs"`
	ExpandSnippets                     bool        `json:"expand_snippets"`
	PostsFormattingGuide               bool        `json:"posts_formatting_guide"`
	SeenWelcome2                       bool        `json:"seen_welcome_2"`
	SeenSsbPrompt                      bool        `json:"seen_ssb_prompt"`
	SpacesNewXpBannerDismissed         bool        `json:"spaces_new_xp_banner_dismissed"`
	SearchOnlyMyChannels               bool        `json:"search_only_my_channels"`
	SearchOnlyCurrentTeam              bool        `json:"search_only_current_team"`
	EmojiMode                          string      `json:"emoji_mode"`
	EmojiUse                           string      `json:"emoji_use"`
	HasInvited                         bool        `json:"has_invited"`
	HasUploaded                        bool        `json:"has_uploaded"`
	HasCreatedChannel                  bool        `json:"has_created_channel"`
	HasSearched                        bool        `json:"has_searched"`
	SearchExcludeChannels              string      `json:"search_exclude_channels"`
	MessagesTheme                      string      `json:"messages_theme"`
	WebappSpellcheck                   bool        `json:"webapp_spellcheck"`
	NoJoinedOverlays                   bool        `json:"no_joined_overlays"`
	NoCreatedOverlays                  bool        `json:"no_created_overlays"`
	DropboxEnabled                     bool        `json:"dropbox_enabled"`
	SeenDomainInviteReminder           bool        `json:"seen_domain_invite_reminder"`
	SeenMemberInviteReminder           bool        `json:"seen_member_invite_reminder"`
	MuteSounds                         bool        `json:"mute_sounds"`
	ArrowHistory                       bool        `json:"arrow_history"`
	TabUIReturnSelects                 bool        `json:"tab_ui_return_selects"`
	ObeyInlineImgLimit                 bool        `json:"obey_inline_img_limit"`
	NewMsgSnd                          string      `json:"new_msg_snd"`
	RequireAt                          bool        `json:"require_at"`
	SsbSpaceWindow                     string      `json:"ssb_space_window"`
	MacSsbBounce                       string      `json:"mac_ssb_bounce"`
	MacSsbBullet                       bool        `json:"mac_ssb_bullet"`
	ExpandNonMediaAttachments          bool        `json:"expand_non_media_attachments"`
	ShowTyping                         bool        `json:"show_typing"`
	PagekeysHandled                    bool        `json:"pagekeys_handled"`
	LastSnippetType                    string      `json:"last_snippet_type"`
	DisplayRealNamesOverride           int         `json:"display_real_names_override"`
	DisplayPreferredNames              bool        `json:"display_preferred_names"`
	Time24                             bool        `json:"time24"`
	EnterIsSpecialInTbt                bool        `json:"enter_is_special_in_tbt"`
	GraphicEmoticons                   bool        `json:"graphic_emoticons"`
	ConvertEmoticons                   bool        `json:"convert_emoticons"`
	SsEmojis                           bool        `json:"ss_emojis"`
	SidebarBehavior                    string      `json:"sidebar_behavior"`
	SeenOnboardingStart                bool        `json:"seen_onboarding_start"`
	OnboardingCancelled                bool        `json:"onboarding_cancelled"`
	SeenOnboardingSlackbotConversation bool        `json:"seen_onboarding_slackbot_conversation"`
	SeenOnboardingChannels             bool        `json:"seen_onboarding_channels"`
	SeenOnboardingDirectMessages       bool        `json:"seen_onboarding_direct_messages"`
	SeenOnboardingInvites              bool        `json:"seen_onboarding_invites"`
	SeenOnboardingSearch               bool        `json:"seen_onboarding_search"`
	SeenOnboardingRecentMentions       bool        `json:"seen_onboarding_recent_mentions"`
	SeenOnboardingStarredItems         bool        `json:"seen_onboarding_starred_items"`
	SeenOnboardingPrivateGroups        bool        `json:"seen_onboarding_private_groups"`
	OnboardingSlackbotConversationStep int         `json:"onboarding_slackbot_conversation_step"`
	DndEnabled                         bool        `json:"dnd_enabled"`
	DndStartHour                       string      `json:"dnd_start_hour"`
	DndEndHour                         string      `json:"dnd_end_hour"`
	MarkMsgsReadImmediately            bool        `json:"mark_msgs_read_immediately"`
	StartScrollAtOldest                bool        `json:"start_scroll_at_oldest"`
	SnippetEditorWrapLongLines         bool        `json:"snippet_editor_wrap_long_lines"`
	LsDisabled                         bool        `json:"ls_disabled"`
	SidebarTheme                       string      `json:"sidebar_theme"`
	SidebarThemeCustomValues           string      `json:"sidebar_theme_custom_values"`
	FKeySearch                         bool        `json:"f_key_search"`
	KKeyOmnibox                        bool        `json:"k_key_omnibox"`
	SpeakGrowls                        bool        `json:"speak_growls"`
	MacSpeakVoice                      string      `json:"mac_speak_voice"`
	MacSpeakSpeed                      int         `json:"mac_speak_speed"`
	CommaKeyPrefs                      bool        `json:"comma_key_prefs"`
	AtChannelSuppressedChannels        string      `json:"at_channel_suppressed_channels"`
	PushAtChannelSuppressedChannels    string      `json:"push_at_channel_suppressed_channels"`
	PromptedForEmailDisabling          bool        `json:"prompted_for_email_disabling"`
	FullTextExtracts                   bool        `json:"full_text_extracts"`
	NoTextInNotifications              bool        `json:"no_text_in_notifications"`
	MutedChannels                      string      `json:"muted_channels"`
	NoMacelectronBanner                bool        `json:"no_macelectron_banner"`
	NoMacssb1Banner                    bool        `json:"no_macssb1_banner"`
	NoMacssb2Banner                    bool        `json:"no_macssb2_banner"`
	NoWinssb1Banner                    bool        `json:"no_winssb1_banner"`
	NoInvitesWidgetInSidebar           bool        `json:"no_invites_widget_in_sidebar"`
	NoOmniboxInChannels                bool        `json:"no_omnibox_in_channels"`
	KKeyOmniboxAutoHideCount           int         `json:"k_key_omnibox_auto_hide_count"`
	HideUserGroupInfoPane              bool        `json:"hide_user_group_info_pane"`
	MentionsExcludeAtUserGroups        bool        `json:"mentions_exclude_at_user_groups"`
	PrivacyPolicySeen                  bool        `json:"privacy_policy_seen"`
	EnterpriseMigrationSeen            bool        `json:"enterprise_migration_seen"`
	SearchExcludeBots                  bool        `json:"search_exclude_bots"`
	LoadLato2                          bool        `json:"load_lato_2"`
	FullerTimestamps                   bool        `json:"fuller_timestamps"`
	LastSeenAtChannelWarning           int         `json:"last_seen_at_channel_warning"`
	FlexResizeWindow                   bool        `json:"flex_resize_window"`
	MsgPreview                         bool        `json:"msg_preview"`
	MsgPreviewPersistent               bool        `json:"msg_preview_persistent"`
	EmojiAutocompleteBig               bool        `json:"emoji_autocomplete_big"`
	WinssbRunFromTray                  bool        `json:"winssb_run_from_tray"`
	WinssbWindowFlashBehavior          string      `json:"winssb_window_flash_behavior"`
	TwoFactorAuthEnabled               bool        `json:"two_factor_auth_enabled"`
	TwoFactorType                      interface{} `json:"two_factor_type"`
	TwoFactorBackupType                interface{} `json:"two_factor_backup_type"`
	ClientLogsPri                      string      `json:"client_logs_pri"`
	EnhancedDebugging                  bool        `json:"enhanced_debugging"`
	FlannelLazyMembers                 bool        `json:"flannel_lazy_members"`
	FlannelServerPool                  string      `json:"flannel_server_pool"`
	MentionsExcludeAtChannels          bool        `json:"mentions_exclude_at_channels"`
	ConfirmClearAllUnreads             bool        `json:"confirm_clear_all_unreads"`
	ConfirmUserMarkedAway              bool        `json:"confirm_user_marked_away"`
	BoxEnabled                         bool        `json:"box_enabled"`
	SeenSingleEmojiMsg                 bool        `json:"seen_single_emoji_msg"`
	ConfirmShCallStart                 bool        `json:"confirm_sh_call_start"`
	PreferredSkinTone                  string      `json:"preferred_skin_tone"`
	ShowAllSkinTones                   bool        `json:"show_all_skin_tones"`
	SeparatePrivateChannels            bool        `json:"separate_private_channels"`
	WhatsNewRead                       int         `json:"whats_new_read"`
	Hotness                            bool        `json:"hotness"`
	FrecencyJumper                     string      `json:"frecency_jumper"`
	FrecencyEntJumper                  string      `json:"frecency_ent_jumper"`
	Jumbomoji                          bool        `json:"jumbomoji"`
	NoFlexInHistory                    bool        `json:"no_flex_in_history"`
	NewxpSeenLastMessage               int         `json:"newxp_seen_last_message"`
	AttachmentsWithBorders             bool        `json:"attachments_with_borders"`
	ShowMemoryInstrument               bool        `json:"show_memory_instrument"`
	EnableUnreadView                   bool        `json:"enable_unread_view"`
	SeenUnreadViewCoachmark            bool        `json:"seen_unread_view_coachmark"`
	SeenCallsVideoBetaCoachmark        bool        `json:"seen_calls_video_beta_coachmark"`
	MeasureCSSUsage                    bool        `json:"measure_css_usage"`
	SeenRepliesCoachmark               bool        `json:"seen_replies_coachmark"`
	AllUnreadsSortOrder                string      `json:"all_unreads_sort_order"`
	GdriveAuthed                       bool        `json:"gdrive_authed"`
	GdriveEnabled                      bool        `json:"gdrive_enabled"`
	SeenGdriveCoachmark                bool        `json:"seen_gdrive_coachmark"`
	ChannelSort                        string      `json:"channel_sort"`
	OverloadedMessageEnabled           bool        `json:"overloaded_message_enabled"`
	A11YFontSize                       string      `json:"a11y_font_size"`
	A11YAnimations                     bool        `json:"a11y_animations"`
}

//Message ...
type Message struct {
	Type      string `json:"type"`
	Channel   string `json:"channel"`
	User      string `json:"user"`
	Text      string `json:"text"`
	TimeStamp string `json:"ts"`
	SubType   string `json:"subtype"`
}
