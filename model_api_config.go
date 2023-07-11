/*
VRChat API Documentation

![VRChat API Banner](https://vrchatapi.github.io/assets/img/api_banner_1500x400.png)  # Welcome to the VRChat API  Before we begin, we would like to state this is a **COMMUNITY DRIVEN PROJECT**. This means that everything you read on here was written by the community itself and is **not** officially supported by VRChat. The documentation is provided \"AS IS\", and any action you take towards VRChat is completely your own responsibility.  The documentation and additional libraries SHALL ONLY be used for applications interacting with VRChat's API in accordance with their [Terms of Service](https://hello.vrchat.com/legal) and [Community Guidelines](https://hello.vrchat.com/community-guidelines), and MUST NOT be used for modifying the client, \"avatar ripping\", or other illegal activities. Malicious usage or spamming the API may result in account termination. Certain parts of the API are also more sensitive than others, for example moderation, so please tread extra carefully and read the warnings when present.  ![Tupper Policy on API](https://i.imgur.com/yLlW7Ok.png)  Finally, use of the API using applications other than the approved methods (website, VRChat application, Unity SDK) is not officially supported. VRChat provides no guarantee or support for external applications using the API. Access to API endpoints may break **at any time, without notice**. Therefore, please **do not ping** VRChat Staff in the VRChat Discord if you are having API problems, as they do not provide API support. We will make a best effort in keeping this documentation and associated language libraries up to date, but things might be outdated or missing. If you find that something is no longer valid, please contact us on Discord or [create an issue](https://github.com/vrchatapi/specification/issues) and tell us so we can fix it.  # Getting Started  The VRChat API can be used to programmatically retrieve or update information regarding your profile, friends, avatars, worlds and more. The API consists of two parts, \"Photon\" which is only used in-game, and the \"Web API\" which is used by both the game and the website. This documentation focuses only on the Web API.  The API is designed around the REST ideology, providing semi-simple and usually predictable URIs to access and modify objects. Requests support standard HTTP methods like GET, PUT, POST, and DELETE and standard status codes. Response bodies are always UTF-8 encoded JSON objects, unless explicitly documented otherwise.  <div class=\"callout callout-error\">   <strong>🛑 Warning! Do not touch Photon!</strong><br>   Photon is only used by the in-game client and should <b>not</b> be touched. Doing so may result in permanent account termination. </div>  <div class=\"callout callout-info\">   <strong>ℹ️ API Key and Authentication</strong><br>   The API Key has always been the same and is currently <code>JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26</code>.   Read <a href=\"#tag--authentication\">Authentication</a> for how to log in. </div>  # Using the API  For simply exploring what the API can do it is strongly recommended to download [Insomnia](https://insomnia.rest/download), a free and open-source API client that's great for sending requests to the API in an orderly fashion. Insomnia allows you to send data in the format that's required for VRChat's API. It is also possible to try out the API in your browser, by first logging in at [vrchat.com/home](https://vrchat.com/home/) and then going to [vrchat.com/api/1/auth/user](https://vrchat.com/api/1/auth/user), but the information will be much harder to work with.  For more permanent operation such as software development it is instead recommended to use one of the existing language SDKs. This community project maintains API libraries in several languages, which allows you to interact with the API with simple function calls rather than having to implement the HTTP protocol yourself. Most of these libraries are automatically generated from the API specification, sometimes with additional helpful wrapper code to make usage easier. This allows them to be almost automatically updated and expanded upon as soon as a new feature is introduced in the specification itself. The libraries can be found on [GitHub](https://github.com/vrchatapi) or following:  * [NodeJS (JavaScript)](https://www.npmjs.com/package/vrchat) * [Dart](https://pub.dev/packages/vrchat_dart) * [Rust](https://crates.io/crates/vrchatapi) * [C#](https://github.com/vrchatapi/vrchatapi-csharp) * [Python](https://github.com/vrchatapi/vrchatapi-python)  # Pagination  Most endpoints enforce pagination, meaning they will only return 10 entries by default, and never more than 100.<br> Using both the limit and offset parameters allows you to easily paginate through a large number of objects.  | Query Parameter | Type | Description | | ----------|--|------- | | `n` | integer  | The number of objects to return. This value often defaults to 10. Highest limit is always 100.| | `offset` | integer  | A zero-based offset from the default object sorting.|  If a request returns fewer objects than the `limit` parameter, there are no more items available to return.  # Contribution  Do you want to get involved in the documentation effort? Do you want to help improve one of the language API libraries? This project is an [OPEN Open Source Project](https://openopensource.org)! This means that individuals making significant and valuable contributions are given commit-access to the project. It also means we are very open and welcoming of new people making contributions, unlike some more guarded open-source projects.  [![Discord](https://img.shields.io/static/v1?label=vrchatapi&message=discord&color=blueviolet&style=for-the-badge)](https://discord.gg/qjZE9C9fkB)

API version: 1.12.0
Contact: vrchatapi.lpv0t@aries.fyi
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vrchatapi

import (
	"encoding/json"
	"time"
)

// checks if the APIConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIConfig{}

// APIConfig 
type APIConfig struct {
	// Unknown, probably voice optimization testing
	VoiceEnableDegradation bool `json:"VoiceEnableDegradation"`
	// Unknown, probably voice optimization testing
	VoiceEnableReceiverLimiting bool `json:"VoiceEnableReceiverLimiting"`
	// VRChat's office address
	Address string `json:"address"`
	// Public Announcements
	Announcements []APIConfigAnnouncement `json:"announcements"`
	// Game name
	// Deprecated
	AppName string `json:"appName"`
	// Build tag of the API server
	BuildVersionTag string `json:"buildVersionTag"`
	// apiKey to be used for all other requests
	ClientApiKey string `json:"clientApiKey"`
	// Unknown
	ClientBPSCeiling int32 `json:"clientBPSCeiling"`
	// Unknown
	ClientDisconnectTimeout int32 `json:"clientDisconnectTimeout"`
	// Unknown
	ClientReservedPlayerBPS int32 `json:"clientReservedPlayerBPS"`
	// Unknown
	ClientSentCountAllowance int32 `json:"clientSentCountAllowance"`
	// VRChat's contact email
	ContactEmail string `json:"contactEmail"`
	// VRChat's copyright-issues-related email
	CopyrightEmail string `json:"copyrightEmail"`
	// Current version number of the Terms of Service
	CurrentTOSVersion int32 `json:"currentTOSVersion"`
	DefaultAvatar string `json:"defaultAvatar"`
	DeploymentGroup DeploymentGroup `json:"deploymentGroup"`
	// Version number for game development build
	// Deprecated
	DevAppVersionStandalone string `json:"devAppVersionStandalone"`
	// Developer Download link
	// Deprecated
	DevDownloadLinkWindows string `json:"devDownloadLinkWindows"`
	// Link to download the development SDK, use downloadUrls instead
	// Deprecated
	DevSdkUrl string `json:"devSdkUrl"`
	// Version of the development SDK
	// Deprecated
	DevSdkVersion string `json:"devSdkVersion"`
	// Version number for server development build
	// Deprecated
	DevServerVersionStandalone string `json:"devServerVersionStandalone"`
	// Unknown, \"dis\" maybe for disconnect?
	DisCountdown time.Time `json:"dis-countdown"`
	// Toggles if copying avatars should be disabled
	DisableAvatarCopying bool `json:"disableAvatarCopying"`
	// Toggles if avatar gating should be disabled. Avatar gating restricts uploading of avatars to people with the `system_avatar_access` Tag or `admin_avatar_access` Tag
	DisableAvatarGating bool `json:"disableAvatarGating"`
	// Toggles if the Community Labs should be disabled
	DisableCommunityLabs bool `json:"disableCommunityLabs"`
	// Toggles if promotion out of Community Labs should be disabled
	DisableCommunityLabsPromotion bool `json:"disableCommunityLabsPromotion"`
	// Unknown
	DisableEmail bool `json:"disableEmail"`
	// Toggles if Analytics should be disabled.
	DisableEventStream bool `json:"disableEventStream"`
	// Toggles if feedback gating should be disabled. Feedback gating restricts submission of feedback (reporting a World or User) to people with the `system_feedback_access` Tag.
	DisableFeedbackGating bool `json:"disableFeedbackGating"`
	// Unknown, probably toggles compilation of frontend web builds? So internal flag?
	DisableFrontendBuilds bool `json:"disableFrontendBuilds"`
	// Unknown
	DisableHello bool `json:"disableHello"`
	// Toggles if signing up for Subscriptions in Oculus is disabled or not.
	DisableOculusSubs bool `json:"disableOculusSubs"`
	// Toggles if new user account registration should be disabled.
	DisableRegistration bool `json:"disableRegistration"`
	// Toggles if Steam Networking should be disabled. VRChat these days uses Photon Unity Networking (PUN) instead.
	DisableSteamNetworking bool `json:"disableSteamNetworking"`
	// Toggles if 2FA should be disabled.
	// Deprecated
	DisableTwoFactorAuth bool `json:"disableTwoFactorAuth"`
	// Toggles if Udon should be universally disabled in-game.
	DisableUdon bool `json:"disableUdon"`
	// Toggles if account upgrading \"linking with Steam/Oculus\" should be disabled.
	DisableUpgradeAccount bool `json:"disableUpgradeAccount"`
	// Download link for game on the Oculus Rift website.
	DownloadLinkWindows string `json:"downloadLinkWindows"`
	DownloadUrls APIConfigDownloadURLList `json:"downloadUrls"`
	// Array of DynamicWorldRow objects, used by the game to display the list of world rows
	DynamicWorldRows []DynamicContentRow `json:"dynamicWorldRows"`
	Events APIConfigEvents `json:"events"`
	// Unknown
	// Deprecated
	GearDemoRoomId string `json:"gearDemoRoomId"`
	// WorldID be \"offline\" on User profiles if you are not friends with that user.
	HomeWorldId string `json:"homeWorldId"`
	// Redirect target if you try to open the base API domain in your browser
	HomepageRedirectTarget string `json:"homepageRedirectTarget"`
	// WorldID be \"offline\" on User profiles if you are not friends with that user.
	HubWorldId string `json:"hubWorldId"`
	// VRChat's job application email
	JobsEmail string `json:"jobsEmail"`
	// MOTD
	// Deprecated
	MessageOfTheDay string `json:"messageOfTheDay"`
	// VRChat's moderation related email
	ModerationEmail string `json:"moderationEmail"`
	// Unknown
	ModerationQueryPeriod int32 `json:"moderationQueryPeriod"`
	// Used in-game to notify a user they aren't allowed to select avatars in private worlds
	NotAllowedToSelectAvatarInPrivateWorldMessage string `json:"notAllowedToSelectAvatarInPrivateWorldMessage"`
	// Extra [plugin](https://doc.photonengine.com/en-us/server/current/plugins/manual) to run in each instance
	Plugin string `json:"plugin"`
	// Version number for game release build
	// Deprecated
	ReleaseAppVersionStandalone string `json:"releaseAppVersionStandalone"`
	// Link to download the release SDK
	// Deprecated
	ReleaseSdkUrl string `json:"releaseSdkUrl"`
	// Version of the release SDK
	// Deprecated
	ReleaseSdkVersion string `json:"releaseSdkVersion"`
	// Version number for server release build
	// Deprecated
	ReleaseServerVersionStandalone string `json:"releaseServerVersionStandalone"`
	// Link to the developer FAQ
	SdkDeveloperFaqUrl string `json:"sdkDeveloperFaqUrl"`
	// Link to the official VRChat Discord
	SdkDiscordUrl string `json:"sdkDiscordUrl"`
	// Used in the SDK to notify a user they aren't allowed to upload avatars/worlds yet
	SdkNotAllowedToPublishMessage string `json:"sdkNotAllowedToPublishMessage"`
	// Unity version supported by the SDK
	SdkUnityVersion string `json:"sdkUnityVersion"`
	// Server name of the API server currently responding
	ServerName string `json:"serverName"`
	// VRChat's support email
	SupportEmail string `json:"supportEmail"`
	// WorldID be \"offline\" on User profiles if you are not friends with that user.
	TimeOutWorldId string `json:"timeOutWorldId"`
	// WorldID be \"offline\" on User profiles if you are not friends with that user.
	TutorialWorldId string `json:"tutorialWorldId"`
	// Unknown
	UpdateRateMsMaximum int32 `json:"updateRateMsMaximum"`
	// Unknown
	UpdateRateMsMinimum int32 `json:"updateRateMsMinimum"`
	// Unknown
	UpdateRateMsNormal int32 `json:"updateRateMsNormal"`
	// Unknown
	UpdateRateMsUdonManual int32 `json:"updateRateMsUdonManual"`
	// Unknown
	UploadAnalysisPercent int32 `json:"uploadAnalysisPercent"`
	// List of allowed URLs that bypass the \"Allow untrusted URL's\" setting in-game
	UrlList []string `json:"urlList"`
	// Unknown
	UseReliableUdpForVoice bool `json:"useReliableUdpForVoice"`
	// Unknown
	UserUpdatePeriod int32 `json:"userUpdatePeriod"`
	// Unknown
	UserVerificationDelay int32 `json:"userVerificationDelay"`
	// Unknown
	UserVerificationRetry int32 `json:"userVerificationRetry"`
	// Unknown
	UserVerificationTimeout int32 `json:"userVerificationTimeout"`
	// Download link for game on the Steam website.
	ViveWindowsUrl string `json:"viveWindowsUrl"`
	// List of allowed URLs that are allowed to host avatar assets
	WhiteListedAssetUrls []string `json:"whiteListedAssetUrls"`
	// Unknown
	WorldUpdatePeriod int32 `json:"worldUpdatePeriod"`
	// Currently used youtube-dl.exe hash in SHA-256-delimited format
	PlayerUrlResolverHash string `json:"player-url-resolver-hash"`
	// Currently used youtube-dl.exe version
	PlayerUrlResolverVersion string `json:"player-url-resolver-version"`
}

// NewAPIConfig instantiates a new APIConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIConfig(voiceEnableDegradation bool, voiceEnableReceiverLimiting bool, address string, announcements []APIConfigAnnouncement, appName string, buildVersionTag string, clientApiKey string, clientBPSCeiling int32, clientDisconnectTimeout int32, clientReservedPlayerBPS int32, clientSentCountAllowance int32, contactEmail string, copyrightEmail string, currentTOSVersion int32, defaultAvatar string, deploymentGroup DeploymentGroup, devAppVersionStandalone string, devDownloadLinkWindows string, devSdkUrl string, devSdkVersion string, devServerVersionStandalone string, disCountdown time.Time, disableAvatarCopying bool, disableAvatarGating bool, disableCommunityLabs bool, disableCommunityLabsPromotion bool, disableEmail bool, disableEventStream bool, disableFeedbackGating bool, disableFrontendBuilds bool, disableHello bool, disableOculusSubs bool, disableRegistration bool, disableSteamNetworking bool, disableTwoFactorAuth bool, disableUdon bool, disableUpgradeAccount bool, downloadLinkWindows string, downloadUrls APIConfigDownloadURLList, dynamicWorldRows []DynamicContentRow, events APIConfigEvents, gearDemoRoomId string, homeWorldId string, homepageRedirectTarget string, hubWorldId string, jobsEmail string, messageOfTheDay string, moderationEmail string, moderationQueryPeriod int32, notAllowedToSelectAvatarInPrivateWorldMessage string, plugin string, releaseAppVersionStandalone string, releaseSdkUrl string, releaseSdkVersion string, releaseServerVersionStandalone string, sdkDeveloperFaqUrl string, sdkDiscordUrl string, sdkNotAllowedToPublishMessage string, sdkUnityVersion string, serverName string, supportEmail string, timeOutWorldId string, tutorialWorldId string, updateRateMsMaximum int32, updateRateMsMinimum int32, updateRateMsNormal int32, updateRateMsUdonManual int32, uploadAnalysisPercent int32, urlList []string, useReliableUdpForVoice bool, userUpdatePeriod int32, userVerificationDelay int32, userVerificationRetry int32, userVerificationTimeout int32, viveWindowsUrl string, whiteListedAssetUrls []string, worldUpdatePeriod int32, playerUrlResolverHash string, playerUrlResolverVersion string) *APIConfig {
	this := APIConfig{}
	this.VoiceEnableDegradation = voiceEnableDegradation
	this.VoiceEnableReceiverLimiting = voiceEnableReceiverLimiting
	this.Address = address
	this.Announcements = announcements
	this.AppName = appName
	this.BuildVersionTag = buildVersionTag
	this.ClientApiKey = clientApiKey
	this.ClientBPSCeiling = clientBPSCeiling
	this.ClientDisconnectTimeout = clientDisconnectTimeout
	this.ClientReservedPlayerBPS = clientReservedPlayerBPS
	this.ClientSentCountAllowance = clientSentCountAllowance
	this.ContactEmail = contactEmail
	this.CopyrightEmail = copyrightEmail
	this.CurrentTOSVersion = currentTOSVersion
	this.DefaultAvatar = defaultAvatar
	this.DeploymentGroup = deploymentGroup
	this.DevAppVersionStandalone = devAppVersionStandalone
	this.DevDownloadLinkWindows = devDownloadLinkWindows
	this.DevSdkUrl = devSdkUrl
	this.DevSdkVersion = devSdkVersion
	this.DevServerVersionStandalone = devServerVersionStandalone
	this.DisCountdown = disCountdown
	this.DisableAvatarCopying = disableAvatarCopying
	this.DisableAvatarGating = disableAvatarGating
	this.DisableCommunityLabs = disableCommunityLabs
	this.DisableCommunityLabsPromotion = disableCommunityLabsPromotion
	this.DisableEmail = disableEmail
	this.DisableEventStream = disableEventStream
	this.DisableFeedbackGating = disableFeedbackGating
	this.DisableFrontendBuilds = disableFrontendBuilds
	this.DisableHello = disableHello
	this.DisableOculusSubs = disableOculusSubs
	this.DisableRegistration = disableRegistration
	this.DisableSteamNetworking = disableSteamNetworking
	this.DisableTwoFactorAuth = disableTwoFactorAuth
	this.DisableUdon = disableUdon
	this.DisableUpgradeAccount = disableUpgradeAccount
	this.DownloadLinkWindows = downloadLinkWindows
	this.DownloadUrls = downloadUrls
	this.DynamicWorldRows = dynamicWorldRows
	this.Events = events
	this.GearDemoRoomId = gearDemoRoomId
	this.HomeWorldId = homeWorldId
	this.HomepageRedirectTarget = homepageRedirectTarget
	this.HubWorldId = hubWorldId
	this.JobsEmail = jobsEmail
	this.MessageOfTheDay = messageOfTheDay
	this.ModerationEmail = moderationEmail
	this.ModerationQueryPeriod = moderationQueryPeriod
	this.NotAllowedToSelectAvatarInPrivateWorldMessage = notAllowedToSelectAvatarInPrivateWorldMessage
	this.Plugin = plugin
	this.ReleaseAppVersionStandalone = releaseAppVersionStandalone
	this.ReleaseSdkUrl = releaseSdkUrl
	this.ReleaseSdkVersion = releaseSdkVersion
	this.ReleaseServerVersionStandalone = releaseServerVersionStandalone
	this.SdkDeveloperFaqUrl = sdkDeveloperFaqUrl
	this.SdkDiscordUrl = sdkDiscordUrl
	this.SdkNotAllowedToPublishMessage = sdkNotAllowedToPublishMessage
	this.SdkUnityVersion = sdkUnityVersion
	this.ServerName = serverName
	this.SupportEmail = supportEmail
	this.TimeOutWorldId = timeOutWorldId
	this.TutorialWorldId = tutorialWorldId
	this.UpdateRateMsMaximum = updateRateMsMaximum
	this.UpdateRateMsMinimum = updateRateMsMinimum
	this.UpdateRateMsNormal = updateRateMsNormal
	this.UpdateRateMsUdonManual = updateRateMsUdonManual
	this.UploadAnalysisPercent = uploadAnalysisPercent
	this.UrlList = urlList
	this.UseReliableUdpForVoice = useReliableUdpForVoice
	this.UserUpdatePeriod = userUpdatePeriod
	this.UserVerificationDelay = userVerificationDelay
	this.UserVerificationRetry = userVerificationRetry
	this.UserVerificationTimeout = userVerificationTimeout
	this.ViveWindowsUrl = viveWindowsUrl
	this.WhiteListedAssetUrls = whiteListedAssetUrls
	this.WorldUpdatePeriod = worldUpdatePeriod
	this.PlayerUrlResolverHash = playerUrlResolverHash
	this.PlayerUrlResolverVersion = playerUrlResolverVersion
	return &this
}

// NewAPIConfigWithDefaults instantiates a new APIConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIConfigWithDefaults() *APIConfig {
	this := APIConfig{}
	var voiceEnableDegradation bool = false
	this.VoiceEnableDegradation = voiceEnableDegradation
	var voiceEnableReceiverLimiting bool = true
	this.VoiceEnableReceiverLimiting = voiceEnableReceiverLimiting
	var appName string = "VrChat"
	this.AppName = appName
	var clientBPSCeiling int32 = 18432
	this.ClientBPSCeiling = clientBPSCeiling
	var clientDisconnectTimeout int32 = 30000
	this.ClientDisconnectTimeout = clientDisconnectTimeout
	var clientReservedPlayerBPS int32 = 7168
	this.ClientReservedPlayerBPS = clientReservedPlayerBPS
	var clientSentCountAllowance int32 = 15
	this.ClientSentCountAllowance = clientSentCountAllowance
	var deploymentGroup DeploymentGroup = BLUE
	this.DeploymentGroup = deploymentGroup
	var disableAvatarCopying bool = false
	this.DisableAvatarCopying = disableAvatarCopying
	var disableAvatarGating bool = false
	this.DisableAvatarGating = disableAvatarGating
	var disableCommunityLabs bool = false
	this.DisableCommunityLabs = disableCommunityLabs
	var disableCommunityLabsPromotion bool = false
	this.DisableCommunityLabsPromotion = disableCommunityLabsPromotion
	var disableEmail bool = false
	this.DisableEmail = disableEmail
	var disableEventStream bool = false
	this.DisableEventStream = disableEventStream
	var disableFeedbackGating bool = false
	this.DisableFeedbackGating = disableFeedbackGating
	var disableFrontendBuilds bool = false
	this.DisableFrontendBuilds = disableFrontendBuilds
	var disableHello bool = false
	this.DisableHello = disableHello
	var disableOculusSubs bool = false
	this.DisableOculusSubs = disableOculusSubs
	var disableRegistration bool = false
	this.DisableRegistration = disableRegistration
	var disableSteamNetworking bool = true
	this.DisableSteamNetworking = disableSteamNetworking
	var disableTwoFactorAuth bool = false
	this.DisableTwoFactorAuth = disableTwoFactorAuth
	var disableUdon bool = false
	this.DisableUdon = disableUdon
	var disableUpgradeAccount bool = false
	this.DisableUpgradeAccount = disableUpgradeAccount
	var homepageRedirectTarget string = "https://hello.vrchat.com"
	this.HomepageRedirectTarget = homepageRedirectTarget
	var useReliableUdpForVoice bool = false
	this.UseReliableUdpForVoice = useReliableUdpForVoice
	return &this
}

// GetVoiceEnableDegradation returns the VoiceEnableDegradation field value
func (o *APIConfig) GetVoiceEnableDegradation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.VoiceEnableDegradation
}

// GetVoiceEnableDegradationOk returns a tuple with the VoiceEnableDegradation field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetVoiceEnableDegradationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VoiceEnableDegradation, true
}

// SetVoiceEnableDegradation sets field value
func (o *APIConfig) SetVoiceEnableDegradation(v bool) {
	o.VoiceEnableDegradation = v
}

// GetVoiceEnableReceiverLimiting returns the VoiceEnableReceiverLimiting field value
func (o *APIConfig) GetVoiceEnableReceiverLimiting() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.VoiceEnableReceiverLimiting
}

// GetVoiceEnableReceiverLimitingOk returns a tuple with the VoiceEnableReceiverLimiting field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetVoiceEnableReceiverLimitingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VoiceEnableReceiverLimiting, true
}

// SetVoiceEnableReceiverLimiting sets field value
func (o *APIConfig) SetVoiceEnableReceiverLimiting(v bool) {
	o.VoiceEnableReceiverLimiting = v
}

// GetAddress returns the Address field value
func (o *APIConfig) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *APIConfig) SetAddress(v string) {
	o.Address = v
}

// GetAnnouncements returns the Announcements field value
func (o *APIConfig) GetAnnouncements() []APIConfigAnnouncement {
	if o == nil {
		var ret []APIConfigAnnouncement
		return ret
	}

	return o.Announcements
}

// GetAnnouncementsOk returns a tuple with the Announcements field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetAnnouncementsOk() ([]APIConfigAnnouncement, bool) {
	if o == nil {
		return nil, false
	}
	return o.Announcements, true
}

// SetAnnouncements sets field value
func (o *APIConfig) SetAnnouncements(v []APIConfigAnnouncement) {
	o.Announcements = v
}

// GetAppName returns the AppName field value
// Deprecated
func (o *APIConfig) GetAppName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *APIConfig) GetAppNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppName, true
}

// SetAppName sets field value
// Deprecated
func (o *APIConfig) SetAppName(v string) {
	o.AppName = v
}

// GetBuildVersionTag returns the BuildVersionTag field value
func (o *APIConfig) GetBuildVersionTag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BuildVersionTag
}

// GetBuildVersionTagOk returns a tuple with the BuildVersionTag field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetBuildVersionTagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuildVersionTag, true
}

// SetBuildVersionTag sets field value
func (o *APIConfig) SetBuildVersionTag(v string) {
	o.BuildVersionTag = v
}

// GetClientApiKey returns the ClientApiKey field value
func (o *APIConfig) GetClientApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientApiKey
}

// GetClientApiKeyOk returns a tuple with the ClientApiKey field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetClientApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientApiKey, true
}

// SetClientApiKey sets field value
func (o *APIConfig) SetClientApiKey(v string) {
	o.ClientApiKey = v
}

// GetClientBPSCeiling returns the ClientBPSCeiling field value
func (o *APIConfig) GetClientBPSCeiling() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClientBPSCeiling
}

// GetClientBPSCeilingOk returns a tuple with the ClientBPSCeiling field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetClientBPSCeilingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientBPSCeiling, true
}

// SetClientBPSCeiling sets field value
func (o *APIConfig) SetClientBPSCeiling(v int32) {
	o.ClientBPSCeiling = v
}

// GetClientDisconnectTimeout returns the ClientDisconnectTimeout field value
func (o *APIConfig) GetClientDisconnectTimeout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClientDisconnectTimeout
}

// GetClientDisconnectTimeoutOk returns a tuple with the ClientDisconnectTimeout field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetClientDisconnectTimeoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientDisconnectTimeout, true
}

// SetClientDisconnectTimeout sets field value
func (o *APIConfig) SetClientDisconnectTimeout(v int32) {
	o.ClientDisconnectTimeout = v
}

// GetClientReservedPlayerBPS returns the ClientReservedPlayerBPS field value
func (o *APIConfig) GetClientReservedPlayerBPS() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClientReservedPlayerBPS
}

// GetClientReservedPlayerBPSOk returns a tuple with the ClientReservedPlayerBPS field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetClientReservedPlayerBPSOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientReservedPlayerBPS, true
}

// SetClientReservedPlayerBPS sets field value
func (o *APIConfig) SetClientReservedPlayerBPS(v int32) {
	o.ClientReservedPlayerBPS = v
}

// GetClientSentCountAllowance returns the ClientSentCountAllowance field value
func (o *APIConfig) GetClientSentCountAllowance() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ClientSentCountAllowance
}

// GetClientSentCountAllowanceOk returns a tuple with the ClientSentCountAllowance field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetClientSentCountAllowanceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSentCountAllowance, true
}

// SetClientSentCountAllowance sets field value
func (o *APIConfig) SetClientSentCountAllowance(v int32) {
	o.ClientSentCountAllowance = v
}

// GetContactEmail returns the ContactEmail field value
func (o *APIConfig) GetContactEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContactEmail
}

// GetContactEmailOk returns a tuple with the ContactEmail field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetContactEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContactEmail, true
}

// SetContactEmail sets field value
func (o *APIConfig) SetContactEmail(v string) {
	o.ContactEmail = v
}

// GetCopyrightEmail returns the CopyrightEmail field value
func (o *APIConfig) GetCopyrightEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CopyrightEmail
}

// GetCopyrightEmailOk returns a tuple with the CopyrightEmail field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetCopyrightEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CopyrightEmail, true
}

// SetCopyrightEmail sets field value
func (o *APIConfig) SetCopyrightEmail(v string) {
	o.CopyrightEmail = v
}

// GetCurrentTOSVersion returns the CurrentTOSVersion field value
func (o *APIConfig) GetCurrentTOSVersion() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CurrentTOSVersion
}

// GetCurrentTOSVersionOk returns a tuple with the CurrentTOSVersion field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetCurrentTOSVersionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentTOSVersion, true
}

// SetCurrentTOSVersion sets field value
func (o *APIConfig) SetCurrentTOSVersion(v int32) {
	o.CurrentTOSVersion = v
}

// GetDefaultAvatar returns the DefaultAvatar field value
func (o *APIConfig) GetDefaultAvatar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultAvatar
}

// GetDefaultAvatarOk returns a tuple with the DefaultAvatar field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDefaultAvatarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultAvatar, true
}

// SetDefaultAvatar sets field value
func (o *APIConfig) SetDefaultAvatar(v string) {
	o.DefaultAvatar = v
}

// GetDeploymentGroup returns the DeploymentGroup field value
func (o *APIConfig) GetDeploymentGroup() DeploymentGroup {
	if o == nil {
		var ret DeploymentGroup
		return ret
	}

	return o.DeploymentGroup
}

// GetDeploymentGroupOk returns a tuple with the DeploymentGroup field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDeploymentGroupOk() (*DeploymentGroup, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeploymentGroup, true
}

// SetDeploymentGroup sets field value
func (o *APIConfig) SetDeploymentGroup(v DeploymentGroup) {
	o.DeploymentGroup = v
}

// GetDevAppVersionStandalone returns the DevAppVersionStandalone field value
// Deprecated
func (o *APIConfig) GetDevAppVersionStandalone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DevAppVersionStandalone
}

// GetDevAppVersionStandaloneOk returns a tuple with the DevAppVersionStandalone field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *APIConfig) GetDevAppVersionStandaloneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DevAppVersionStandalone, true
}

// SetDevAppVersionStandalone sets field value
// Deprecated
func (o *APIConfig) SetDevAppVersionStandalone(v string) {
	o.DevAppVersionStandalone = v
}

// GetDevDownloadLinkWindows returns the DevDownloadLinkWindows field value
// Deprecated
func (o *APIConfig) GetDevDownloadLinkWindows() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DevDownloadLinkWindows
}

// GetDevDownloadLinkWindowsOk returns a tuple with the DevDownloadLinkWindows field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *APIConfig) GetDevDownloadLinkWindowsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DevDownloadLinkWindows, true
}

// SetDevDownloadLinkWindows sets field value
// Deprecated
func (o *APIConfig) SetDevDownloadLinkWindows(v string) {
	o.DevDownloadLinkWindows = v
}

// GetDevSdkUrl returns the DevSdkUrl field value
// Deprecated
func (o *APIConfig) GetDevSdkUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DevSdkUrl
}

// GetDevSdkUrlOk returns a tuple with the DevSdkUrl field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *APIConfig) GetDevSdkUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DevSdkUrl, true
}

// SetDevSdkUrl sets field value
// Deprecated
func (o *APIConfig) SetDevSdkUrl(v string) {
	o.DevSdkUrl = v
}

// GetDevSdkVersion returns the DevSdkVersion field value
// Deprecated
func (o *APIConfig) GetDevSdkVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DevSdkVersion
}

// GetDevSdkVersionOk returns a tuple with the DevSdkVersion field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *APIConfig) GetDevSdkVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DevSdkVersion, true
}

// SetDevSdkVersion sets field value
// Deprecated
func (o *APIConfig) SetDevSdkVersion(v string) {
	o.DevSdkVersion = v
}

// GetDevServerVersionStandalone returns the DevServerVersionStandalone field value
// Deprecated
func (o *APIConfig) GetDevServerVersionStandalone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DevServerVersionStandalone
}

// GetDevServerVersionStandaloneOk returns a tuple with the DevServerVersionStandalone field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *APIConfig) GetDevServerVersionStandaloneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DevServerVersionStandalone, true
}

// SetDevServerVersionStandalone sets field value
// Deprecated
func (o *APIConfig) SetDevServerVersionStandalone(v string) {
	o.DevServerVersionStandalone = v
}

// GetDisCountdown returns the DisCountdown field value
func (o *APIConfig) GetDisCountdown() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DisCountdown
}

// GetDisCountdownOk returns a tuple with the DisCountdown field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisCountdownOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisCountdown, true
}

// SetDisCountdown sets field value
func (o *APIConfig) SetDisCountdown(v time.Time) {
	o.DisCountdown = v
}

// GetDisableAvatarCopying returns the DisableAvatarCopying field value
func (o *APIConfig) GetDisableAvatarCopying() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableAvatarCopying
}

// GetDisableAvatarCopyingOk returns a tuple with the DisableAvatarCopying field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableAvatarCopyingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableAvatarCopying, true
}

// SetDisableAvatarCopying sets field value
func (o *APIConfig) SetDisableAvatarCopying(v bool) {
	o.DisableAvatarCopying = v
}

// GetDisableAvatarGating returns the DisableAvatarGating field value
func (o *APIConfig) GetDisableAvatarGating() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableAvatarGating
}

// GetDisableAvatarGatingOk returns a tuple with the DisableAvatarGating field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableAvatarGatingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableAvatarGating, true
}

// SetDisableAvatarGating sets field value
func (o *APIConfig) SetDisableAvatarGating(v bool) {
	o.DisableAvatarGating = v
}

// GetDisableCommunityLabs returns the DisableCommunityLabs field value
func (o *APIConfig) GetDisableCommunityLabs() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableCommunityLabs
}

// GetDisableCommunityLabsOk returns a tuple with the DisableCommunityLabs field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableCommunityLabsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableCommunityLabs, true
}

// SetDisableCommunityLabs sets field value
func (o *APIConfig) SetDisableCommunityLabs(v bool) {
	o.DisableCommunityLabs = v
}

// GetDisableCommunityLabsPromotion returns the DisableCommunityLabsPromotion field value
func (o *APIConfig) GetDisableCommunityLabsPromotion() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableCommunityLabsPromotion
}

// GetDisableCommunityLabsPromotionOk returns a tuple with the DisableCommunityLabsPromotion field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableCommunityLabsPromotionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableCommunityLabsPromotion, true
}

// SetDisableCommunityLabsPromotion sets field value
func (o *APIConfig) SetDisableCommunityLabsPromotion(v bool) {
	o.DisableCommunityLabsPromotion = v
}

// GetDisableEmail returns the DisableEmail field value
func (o *APIConfig) GetDisableEmail() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableEmail
}

// GetDisableEmailOk returns a tuple with the DisableEmail field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableEmailOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableEmail, true
}

// SetDisableEmail sets field value
func (o *APIConfig) SetDisableEmail(v bool) {
	o.DisableEmail = v
}

// GetDisableEventStream returns the DisableEventStream field value
func (o *APIConfig) GetDisableEventStream() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableEventStream
}

// GetDisableEventStreamOk returns a tuple with the DisableEventStream field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableEventStreamOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableEventStream, true
}

// SetDisableEventStream sets field value
func (o *APIConfig) SetDisableEventStream(v bool) {
	o.DisableEventStream = v
}

// GetDisableFeedbackGating returns the DisableFeedbackGating field value
func (o *APIConfig) GetDisableFeedbackGating() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableFeedbackGating
}

// GetDisableFeedbackGatingOk returns a tuple with the DisableFeedbackGating field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableFeedbackGatingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableFeedbackGating, true
}

// SetDisableFeedbackGating sets field value
func (o *APIConfig) SetDisableFeedbackGating(v bool) {
	o.DisableFeedbackGating = v
}

// GetDisableFrontendBuilds returns the DisableFrontendBuilds field value
func (o *APIConfig) GetDisableFrontendBuilds() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableFrontendBuilds
}

// GetDisableFrontendBuildsOk returns a tuple with the DisableFrontendBuilds field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableFrontendBuildsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableFrontendBuilds, true
}

// SetDisableFrontendBuilds sets field value
func (o *APIConfig) SetDisableFrontendBuilds(v bool) {
	o.DisableFrontendBuilds = v
}

// GetDisableHello returns the DisableHello field value
func (o *APIConfig) GetDisableHello() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableHello
}

// GetDisableHelloOk returns a tuple with the DisableHello field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableHelloOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableHello, true
}

// SetDisableHello sets field value
func (o *APIConfig) SetDisableHello(v bool) {
	o.DisableHello = v
}

// GetDisableOculusSubs returns the DisableOculusSubs field value
func (o *APIConfig) GetDisableOculusSubs() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableOculusSubs
}

// GetDisableOculusSubsOk returns a tuple with the DisableOculusSubs field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableOculusSubsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableOculusSubs, true
}

// SetDisableOculusSubs sets field value
func (o *APIConfig) SetDisableOculusSubs(v bool) {
	o.DisableOculusSubs = v
}

// GetDisableRegistration returns the DisableRegistration field value
func (o *APIConfig) GetDisableRegistration() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableRegistration
}

// GetDisableRegistrationOk returns a tuple with the DisableRegistration field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableRegistrationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableRegistration, true
}

// SetDisableRegistration sets field value
func (o *APIConfig) SetDisableRegistration(v bool) {
	o.DisableRegistration = v
}

// GetDisableSteamNetworking returns the DisableSteamNetworking field value
func (o *APIConfig) GetDisableSteamNetworking() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableSteamNetworking
}

// GetDisableSteamNetworkingOk returns a tuple with the DisableSteamNetworking field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableSteamNetworkingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableSteamNetworking, true
}

// SetDisableSteamNetworking sets field value
func (o *APIConfig) SetDisableSteamNetworking(v bool) {
	o.DisableSteamNetworking = v
}

// GetDisableTwoFactorAuth returns the DisableTwoFactorAuth field value
// Deprecated
func (o *APIConfig) GetDisableTwoFactorAuth() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableTwoFactorAuth
}

// GetDisableTwoFactorAuthOk returns a tuple with the DisableTwoFactorAuth field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *APIConfig) GetDisableTwoFactorAuthOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableTwoFactorAuth, true
}

// SetDisableTwoFactorAuth sets field value
// Deprecated
func (o *APIConfig) SetDisableTwoFactorAuth(v bool) {
	o.DisableTwoFactorAuth = v
}

// GetDisableUdon returns the DisableUdon field value
func (o *APIConfig) GetDisableUdon() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableUdon
}

// GetDisableUdonOk returns a tuple with the DisableUdon field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableUdonOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableUdon, true
}

// SetDisableUdon sets field value
func (o *APIConfig) SetDisableUdon(v bool) {
	o.DisableUdon = v
}

// GetDisableUpgradeAccount returns the DisableUpgradeAccount field value
func (o *APIConfig) GetDisableUpgradeAccount() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisableUpgradeAccount
}

// GetDisableUpgradeAccountOk returns a tuple with the DisableUpgradeAccount field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDisableUpgradeAccountOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisableUpgradeAccount, true
}

// SetDisableUpgradeAccount sets field value
func (o *APIConfig) SetDisableUpgradeAccount(v bool) {
	o.DisableUpgradeAccount = v
}

// GetDownloadLinkWindows returns the DownloadLinkWindows field value
func (o *APIConfig) GetDownloadLinkWindows() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DownloadLinkWindows
}

// GetDownloadLinkWindowsOk returns a tuple with the DownloadLinkWindows field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDownloadLinkWindowsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DownloadLinkWindows, true
}

// SetDownloadLinkWindows sets field value
func (o *APIConfig) SetDownloadLinkWindows(v string) {
	o.DownloadLinkWindows = v
}

// GetDownloadUrls returns the DownloadUrls field value
func (o *APIConfig) GetDownloadUrls() APIConfigDownloadURLList {
	if o == nil {
		var ret APIConfigDownloadURLList
		return ret
	}

	return o.DownloadUrls
}

// GetDownloadUrlsOk returns a tuple with the DownloadUrls field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDownloadUrlsOk() (*APIConfigDownloadURLList, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DownloadUrls, true
}

// SetDownloadUrls sets field value
func (o *APIConfig) SetDownloadUrls(v APIConfigDownloadURLList) {
	o.DownloadUrls = v
}

// GetDynamicWorldRows returns the DynamicWorldRows field value
func (o *APIConfig) GetDynamicWorldRows() []DynamicContentRow {
	if o == nil {
		var ret []DynamicContentRow
		return ret
	}

	return o.DynamicWorldRows
}

// GetDynamicWorldRowsOk returns a tuple with the DynamicWorldRows field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetDynamicWorldRowsOk() ([]DynamicContentRow, bool) {
	if o == nil {
		return nil, false
	}
	return o.DynamicWorldRows, true
}

// SetDynamicWorldRows sets field value
func (o *APIConfig) SetDynamicWorldRows(v []DynamicContentRow) {
	o.DynamicWorldRows = v
}

// GetEvents returns the Events field value
func (o *APIConfig) GetEvents() APIConfigEvents {
	if o == nil {
		var ret APIConfigEvents
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetEventsOk() (*APIConfigEvents, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Events, true
}

// SetEvents sets field value
func (o *APIConfig) SetEvents(v APIConfigEvents) {
	o.Events = v
}

// GetGearDemoRoomId returns the GearDemoRoomId field value
// Deprecated
func (o *APIConfig) GetGearDemoRoomId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GearDemoRoomId
}

// GetGearDemoRoomIdOk returns a tuple with the GearDemoRoomId field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *APIConfig) GetGearDemoRoomIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GearDemoRoomId, true
}

// SetGearDemoRoomId sets field value
// Deprecated
func (o *APIConfig) SetGearDemoRoomId(v string) {
	o.GearDemoRoomId = v
}

// GetHomeWorldId returns the HomeWorldId field value
func (o *APIConfig) GetHomeWorldId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HomeWorldId
}

// GetHomeWorldIdOk returns a tuple with the HomeWorldId field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetHomeWorldIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HomeWorldId, true
}

// SetHomeWorldId sets field value
func (o *APIConfig) SetHomeWorldId(v string) {
	o.HomeWorldId = v
}

// GetHomepageRedirectTarget returns the HomepageRedirectTarget field value
func (o *APIConfig) GetHomepageRedirectTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HomepageRedirectTarget
}

// GetHomepageRedirectTargetOk returns a tuple with the HomepageRedirectTarget field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetHomepageRedirectTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HomepageRedirectTarget, true
}

// SetHomepageRedirectTarget sets field value
func (o *APIConfig) SetHomepageRedirectTarget(v string) {
	o.HomepageRedirectTarget = v
}

// GetHubWorldId returns the HubWorldId field value
func (o *APIConfig) GetHubWorldId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HubWorldId
}

// GetHubWorldIdOk returns a tuple with the HubWorldId field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetHubWorldIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HubWorldId, true
}

// SetHubWorldId sets field value
func (o *APIConfig) SetHubWorldId(v string) {
	o.HubWorldId = v
}

// GetJobsEmail returns the JobsEmail field value
func (o *APIConfig) GetJobsEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobsEmail
}

// GetJobsEmailOk returns a tuple with the JobsEmail field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetJobsEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobsEmail, true
}

// SetJobsEmail sets field value
func (o *APIConfig) SetJobsEmail(v string) {
	o.JobsEmail = v
}

// GetMessageOfTheDay returns the MessageOfTheDay field value
// Deprecated
func (o *APIConfig) GetMessageOfTheDay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageOfTheDay
}

// GetMessageOfTheDayOk returns a tuple with the MessageOfTheDay field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *APIConfig) GetMessageOfTheDayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageOfTheDay, true
}

// SetMessageOfTheDay sets field value
// Deprecated
func (o *APIConfig) SetMessageOfTheDay(v string) {
	o.MessageOfTheDay = v
}

// GetModerationEmail returns the ModerationEmail field value
func (o *APIConfig) GetModerationEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModerationEmail
}

// GetModerationEmailOk returns a tuple with the ModerationEmail field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetModerationEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModerationEmail, true
}

// SetModerationEmail sets field value
func (o *APIConfig) SetModerationEmail(v string) {
	o.ModerationEmail = v
}

// GetModerationQueryPeriod returns the ModerationQueryPeriod field value
func (o *APIConfig) GetModerationQueryPeriod() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ModerationQueryPeriod
}

// GetModerationQueryPeriodOk returns a tuple with the ModerationQueryPeriod field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetModerationQueryPeriodOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModerationQueryPeriod, true
}

// SetModerationQueryPeriod sets field value
func (o *APIConfig) SetModerationQueryPeriod(v int32) {
	o.ModerationQueryPeriod = v
}

// GetNotAllowedToSelectAvatarInPrivateWorldMessage returns the NotAllowedToSelectAvatarInPrivateWorldMessage field value
func (o *APIConfig) GetNotAllowedToSelectAvatarInPrivateWorldMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotAllowedToSelectAvatarInPrivateWorldMessage
}

// GetNotAllowedToSelectAvatarInPrivateWorldMessageOk returns a tuple with the NotAllowedToSelectAvatarInPrivateWorldMessage field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetNotAllowedToSelectAvatarInPrivateWorldMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotAllowedToSelectAvatarInPrivateWorldMessage, true
}

// SetNotAllowedToSelectAvatarInPrivateWorldMessage sets field value
func (o *APIConfig) SetNotAllowedToSelectAvatarInPrivateWorldMessage(v string) {
	o.NotAllowedToSelectAvatarInPrivateWorldMessage = v
}

// GetPlugin returns the Plugin field value
func (o *APIConfig) GetPlugin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Plugin
}

// GetPluginOk returns a tuple with the Plugin field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetPluginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plugin, true
}

// SetPlugin sets field value
func (o *APIConfig) SetPlugin(v string) {
	o.Plugin = v
}

// GetReleaseAppVersionStandalone returns the ReleaseAppVersionStandalone field value
// Deprecated
func (o *APIConfig) GetReleaseAppVersionStandalone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseAppVersionStandalone
}

// GetReleaseAppVersionStandaloneOk returns a tuple with the ReleaseAppVersionStandalone field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *APIConfig) GetReleaseAppVersionStandaloneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseAppVersionStandalone, true
}

// SetReleaseAppVersionStandalone sets field value
// Deprecated
func (o *APIConfig) SetReleaseAppVersionStandalone(v string) {
	o.ReleaseAppVersionStandalone = v
}

// GetReleaseSdkUrl returns the ReleaseSdkUrl field value
// Deprecated
func (o *APIConfig) GetReleaseSdkUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseSdkUrl
}

// GetReleaseSdkUrlOk returns a tuple with the ReleaseSdkUrl field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *APIConfig) GetReleaseSdkUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseSdkUrl, true
}

// SetReleaseSdkUrl sets field value
// Deprecated
func (o *APIConfig) SetReleaseSdkUrl(v string) {
	o.ReleaseSdkUrl = v
}

// GetReleaseSdkVersion returns the ReleaseSdkVersion field value
// Deprecated
func (o *APIConfig) GetReleaseSdkVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseSdkVersion
}

// GetReleaseSdkVersionOk returns a tuple with the ReleaseSdkVersion field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *APIConfig) GetReleaseSdkVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseSdkVersion, true
}

// SetReleaseSdkVersion sets field value
// Deprecated
func (o *APIConfig) SetReleaseSdkVersion(v string) {
	o.ReleaseSdkVersion = v
}

// GetReleaseServerVersionStandalone returns the ReleaseServerVersionStandalone field value
// Deprecated
func (o *APIConfig) GetReleaseServerVersionStandalone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseServerVersionStandalone
}

// GetReleaseServerVersionStandaloneOk returns a tuple with the ReleaseServerVersionStandalone field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *APIConfig) GetReleaseServerVersionStandaloneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseServerVersionStandalone, true
}

// SetReleaseServerVersionStandalone sets field value
// Deprecated
func (o *APIConfig) SetReleaseServerVersionStandalone(v string) {
	o.ReleaseServerVersionStandalone = v
}

// GetSdkDeveloperFaqUrl returns the SdkDeveloperFaqUrl field value
func (o *APIConfig) GetSdkDeveloperFaqUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SdkDeveloperFaqUrl
}

// GetSdkDeveloperFaqUrlOk returns a tuple with the SdkDeveloperFaqUrl field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetSdkDeveloperFaqUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SdkDeveloperFaqUrl, true
}

// SetSdkDeveloperFaqUrl sets field value
func (o *APIConfig) SetSdkDeveloperFaqUrl(v string) {
	o.SdkDeveloperFaqUrl = v
}

// GetSdkDiscordUrl returns the SdkDiscordUrl field value
func (o *APIConfig) GetSdkDiscordUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SdkDiscordUrl
}

// GetSdkDiscordUrlOk returns a tuple with the SdkDiscordUrl field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetSdkDiscordUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SdkDiscordUrl, true
}

// SetSdkDiscordUrl sets field value
func (o *APIConfig) SetSdkDiscordUrl(v string) {
	o.SdkDiscordUrl = v
}

// GetSdkNotAllowedToPublishMessage returns the SdkNotAllowedToPublishMessage field value
func (o *APIConfig) GetSdkNotAllowedToPublishMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SdkNotAllowedToPublishMessage
}

// GetSdkNotAllowedToPublishMessageOk returns a tuple with the SdkNotAllowedToPublishMessage field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetSdkNotAllowedToPublishMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SdkNotAllowedToPublishMessage, true
}

// SetSdkNotAllowedToPublishMessage sets field value
func (o *APIConfig) SetSdkNotAllowedToPublishMessage(v string) {
	o.SdkNotAllowedToPublishMessage = v
}

// GetSdkUnityVersion returns the SdkUnityVersion field value
func (o *APIConfig) GetSdkUnityVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SdkUnityVersion
}

// GetSdkUnityVersionOk returns a tuple with the SdkUnityVersion field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetSdkUnityVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SdkUnityVersion, true
}

// SetSdkUnityVersion sets field value
func (o *APIConfig) SetSdkUnityVersion(v string) {
	o.SdkUnityVersion = v
}

// GetServerName returns the ServerName field value
func (o *APIConfig) GetServerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetServerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerName, true
}

// SetServerName sets field value
func (o *APIConfig) SetServerName(v string) {
	o.ServerName = v
}

// GetSupportEmail returns the SupportEmail field value
func (o *APIConfig) GetSupportEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupportEmail
}

// GetSupportEmailOk returns a tuple with the SupportEmail field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetSupportEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupportEmail, true
}

// SetSupportEmail sets field value
func (o *APIConfig) SetSupportEmail(v string) {
	o.SupportEmail = v
}

// GetTimeOutWorldId returns the TimeOutWorldId field value
func (o *APIConfig) GetTimeOutWorldId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeOutWorldId
}

// GetTimeOutWorldIdOk returns a tuple with the TimeOutWorldId field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetTimeOutWorldIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeOutWorldId, true
}

// SetTimeOutWorldId sets field value
func (o *APIConfig) SetTimeOutWorldId(v string) {
	o.TimeOutWorldId = v
}

// GetTutorialWorldId returns the TutorialWorldId field value
func (o *APIConfig) GetTutorialWorldId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TutorialWorldId
}

// GetTutorialWorldIdOk returns a tuple with the TutorialWorldId field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetTutorialWorldIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TutorialWorldId, true
}

// SetTutorialWorldId sets field value
func (o *APIConfig) SetTutorialWorldId(v string) {
	o.TutorialWorldId = v
}

// GetUpdateRateMsMaximum returns the UpdateRateMsMaximum field value
func (o *APIConfig) GetUpdateRateMsMaximum() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdateRateMsMaximum
}

// GetUpdateRateMsMaximumOk returns a tuple with the UpdateRateMsMaximum field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetUpdateRateMsMaximumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateRateMsMaximum, true
}

// SetUpdateRateMsMaximum sets field value
func (o *APIConfig) SetUpdateRateMsMaximum(v int32) {
	o.UpdateRateMsMaximum = v
}

// GetUpdateRateMsMinimum returns the UpdateRateMsMinimum field value
func (o *APIConfig) GetUpdateRateMsMinimum() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdateRateMsMinimum
}

// GetUpdateRateMsMinimumOk returns a tuple with the UpdateRateMsMinimum field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetUpdateRateMsMinimumOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateRateMsMinimum, true
}

// SetUpdateRateMsMinimum sets field value
func (o *APIConfig) SetUpdateRateMsMinimum(v int32) {
	o.UpdateRateMsMinimum = v
}

// GetUpdateRateMsNormal returns the UpdateRateMsNormal field value
func (o *APIConfig) GetUpdateRateMsNormal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdateRateMsNormal
}

// GetUpdateRateMsNormalOk returns a tuple with the UpdateRateMsNormal field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetUpdateRateMsNormalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateRateMsNormal, true
}

// SetUpdateRateMsNormal sets field value
func (o *APIConfig) SetUpdateRateMsNormal(v int32) {
	o.UpdateRateMsNormal = v
}

// GetUpdateRateMsUdonManual returns the UpdateRateMsUdonManual field value
func (o *APIConfig) GetUpdateRateMsUdonManual() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UpdateRateMsUdonManual
}

// GetUpdateRateMsUdonManualOk returns a tuple with the UpdateRateMsUdonManual field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetUpdateRateMsUdonManualOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdateRateMsUdonManual, true
}

// SetUpdateRateMsUdonManual sets field value
func (o *APIConfig) SetUpdateRateMsUdonManual(v int32) {
	o.UpdateRateMsUdonManual = v
}

// GetUploadAnalysisPercent returns the UploadAnalysisPercent field value
func (o *APIConfig) GetUploadAnalysisPercent() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UploadAnalysisPercent
}

// GetUploadAnalysisPercentOk returns a tuple with the UploadAnalysisPercent field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetUploadAnalysisPercentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UploadAnalysisPercent, true
}

// SetUploadAnalysisPercent sets field value
func (o *APIConfig) SetUploadAnalysisPercent(v int32) {
	o.UploadAnalysisPercent = v
}

// GetUrlList returns the UrlList field value
func (o *APIConfig) GetUrlList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.UrlList
}

// GetUrlListOk returns a tuple with the UrlList field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetUrlListOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UrlList, true
}

// SetUrlList sets field value
func (o *APIConfig) SetUrlList(v []string) {
	o.UrlList = v
}

// GetUseReliableUdpForVoice returns the UseReliableUdpForVoice field value
func (o *APIConfig) GetUseReliableUdpForVoice() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseReliableUdpForVoice
}

// GetUseReliableUdpForVoiceOk returns a tuple with the UseReliableUdpForVoice field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetUseReliableUdpForVoiceOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseReliableUdpForVoice, true
}

// SetUseReliableUdpForVoice sets field value
func (o *APIConfig) SetUseReliableUdpForVoice(v bool) {
	o.UseReliableUdpForVoice = v
}

// GetUserUpdatePeriod returns the UserUpdatePeriod field value
func (o *APIConfig) GetUserUpdatePeriod() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserUpdatePeriod
}

// GetUserUpdatePeriodOk returns a tuple with the UserUpdatePeriod field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetUserUpdatePeriodOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserUpdatePeriod, true
}

// SetUserUpdatePeriod sets field value
func (o *APIConfig) SetUserUpdatePeriod(v int32) {
	o.UserUpdatePeriod = v
}

// GetUserVerificationDelay returns the UserVerificationDelay field value
func (o *APIConfig) GetUserVerificationDelay() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserVerificationDelay
}

// GetUserVerificationDelayOk returns a tuple with the UserVerificationDelay field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetUserVerificationDelayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserVerificationDelay, true
}

// SetUserVerificationDelay sets field value
func (o *APIConfig) SetUserVerificationDelay(v int32) {
	o.UserVerificationDelay = v
}

// GetUserVerificationRetry returns the UserVerificationRetry field value
func (o *APIConfig) GetUserVerificationRetry() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserVerificationRetry
}

// GetUserVerificationRetryOk returns a tuple with the UserVerificationRetry field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetUserVerificationRetryOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserVerificationRetry, true
}

// SetUserVerificationRetry sets field value
func (o *APIConfig) SetUserVerificationRetry(v int32) {
	o.UserVerificationRetry = v
}

// GetUserVerificationTimeout returns the UserVerificationTimeout field value
func (o *APIConfig) GetUserVerificationTimeout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UserVerificationTimeout
}

// GetUserVerificationTimeoutOk returns a tuple with the UserVerificationTimeout field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetUserVerificationTimeoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserVerificationTimeout, true
}

// SetUserVerificationTimeout sets field value
func (o *APIConfig) SetUserVerificationTimeout(v int32) {
	o.UserVerificationTimeout = v
}

// GetViveWindowsUrl returns the ViveWindowsUrl field value
func (o *APIConfig) GetViveWindowsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ViveWindowsUrl
}

// GetViveWindowsUrlOk returns a tuple with the ViveWindowsUrl field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetViveWindowsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ViveWindowsUrl, true
}

// SetViveWindowsUrl sets field value
func (o *APIConfig) SetViveWindowsUrl(v string) {
	o.ViveWindowsUrl = v
}

// GetWhiteListedAssetUrls returns the WhiteListedAssetUrls field value
func (o *APIConfig) GetWhiteListedAssetUrls() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.WhiteListedAssetUrls
}

// GetWhiteListedAssetUrlsOk returns a tuple with the WhiteListedAssetUrls field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetWhiteListedAssetUrlsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WhiteListedAssetUrls, true
}

// SetWhiteListedAssetUrls sets field value
func (o *APIConfig) SetWhiteListedAssetUrls(v []string) {
	o.WhiteListedAssetUrls = v
}

// GetWorldUpdatePeriod returns the WorldUpdatePeriod field value
func (o *APIConfig) GetWorldUpdatePeriod() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WorldUpdatePeriod
}

// GetWorldUpdatePeriodOk returns a tuple with the WorldUpdatePeriod field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetWorldUpdatePeriodOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorldUpdatePeriod, true
}

// SetWorldUpdatePeriod sets field value
func (o *APIConfig) SetWorldUpdatePeriod(v int32) {
	o.WorldUpdatePeriod = v
}

// GetPlayerUrlResolverHash returns the PlayerUrlResolverHash field value
func (o *APIConfig) GetPlayerUrlResolverHash() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlayerUrlResolverHash
}

// GetPlayerUrlResolverHashOk returns a tuple with the PlayerUrlResolverHash field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetPlayerUrlResolverHashOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlayerUrlResolverHash, true
}

// SetPlayerUrlResolverHash sets field value
func (o *APIConfig) SetPlayerUrlResolverHash(v string) {
	o.PlayerUrlResolverHash = v
}

// GetPlayerUrlResolverVersion returns the PlayerUrlResolverVersion field value
func (o *APIConfig) GetPlayerUrlResolverVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlayerUrlResolverVersion
}

// GetPlayerUrlResolverVersionOk returns a tuple with the PlayerUrlResolverVersion field value
// and a boolean to check if the value has been set.
func (o *APIConfig) GetPlayerUrlResolverVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlayerUrlResolverVersion, true
}

// SetPlayerUrlResolverVersion sets field value
func (o *APIConfig) SetPlayerUrlResolverVersion(v string) {
	o.PlayerUrlResolverVersion = v
}

func (o APIConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["VoiceEnableDegradation"] = o.VoiceEnableDegradation
	toSerialize["VoiceEnableReceiverLimiting"] = o.VoiceEnableReceiverLimiting
	toSerialize["address"] = o.Address
	toSerialize["announcements"] = o.Announcements
	toSerialize["appName"] = o.AppName
	toSerialize["buildVersionTag"] = o.BuildVersionTag
	toSerialize["clientApiKey"] = o.ClientApiKey
	toSerialize["clientBPSCeiling"] = o.ClientBPSCeiling
	toSerialize["clientDisconnectTimeout"] = o.ClientDisconnectTimeout
	toSerialize["clientReservedPlayerBPS"] = o.ClientReservedPlayerBPS
	toSerialize["clientSentCountAllowance"] = o.ClientSentCountAllowance
	toSerialize["contactEmail"] = o.ContactEmail
	toSerialize["copyrightEmail"] = o.CopyrightEmail
	toSerialize["currentTOSVersion"] = o.CurrentTOSVersion
	toSerialize["defaultAvatar"] = o.DefaultAvatar
	toSerialize["deploymentGroup"] = o.DeploymentGroup
	toSerialize["devAppVersionStandalone"] = o.DevAppVersionStandalone
	toSerialize["devDownloadLinkWindows"] = o.DevDownloadLinkWindows
	toSerialize["devSdkUrl"] = o.DevSdkUrl
	toSerialize["devSdkVersion"] = o.DevSdkVersion
	toSerialize["devServerVersionStandalone"] = o.DevServerVersionStandalone
	toSerialize["dis-countdown"] = o.DisCountdown
	toSerialize["disableAvatarCopying"] = o.DisableAvatarCopying
	toSerialize["disableAvatarGating"] = o.DisableAvatarGating
	toSerialize["disableCommunityLabs"] = o.DisableCommunityLabs
	toSerialize["disableCommunityLabsPromotion"] = o.DisableCommunityLabsPromotion
	toSerialize["disableEmail"] = o.DisableEmail
	toSerialize["disableEventStream"] = o.DisableEventStream
	toSerialize["disableFeedbackGating"] = o.DisableFeedbackGating
	toSerialize["disableFrontendBuilds"] = o.DisableFrontendBuilds
	toSerialize["disableHello"] = o.DisableHello
	toSerialize["disableOculusSubs"] = o.DisableOculusSubs
	toSerialize["disableRegistration"] = o.DisableRegistration
	toSerialize["disableSteamNetworking"] = o.DisableSteamNetworking
	toSerialize["disableTwoFactorAuth"] = o.DisableTwoFactorAuth
	toSerialize["disableUdon"] = o.DisableUdon
	toSerialize["disableUpgradeAccount"] = o.DisableUpgradeAccount
	toSerialize["downloadLinkWindows"] = o.DownloadLinkWindows
	toSerialize["downloadUrls"] = o.DownloadUrls
	toSerialize["dynamicWorldRows"] = o.DynamicWorldRows
	toSerialize["events"] = o.Events
	toSerialize["gearDemoRoomId"] = o.GearDemoRoomId
	toSerialize["homeWorldId"] = o.HomeWorldId
	toSerialize["homepageRedirectTarget"] = o.HomepageRedirectTarget
	toSerialize["hubWorldId"] = o.HubWorldId
	toSerialize["jobsEmail"] = o.JobsEmail
	toSerialize["messageOfTheDay"] = o.MessageOfTheDay
	toSerialize["moderationEmail"] = o.ModerationEmail
	toSerialize["moderationQueryPeriod"] = o.ModerationQueryPeriod
	toSerialize["notAllowedToSelectAvatarInPrivateWorldMessage"] = o.NotAllowedToSelectAvatarInPrivateWorldMessage
	toSerialize["plugin"] = o.Plugin
	toSerialize["releaseAppVersionStandalone"] = o.ReleaseAppVersionStandalone
	toSerialize["releaseSdkUrl"] = o.ReleaseSdkUrl
	toSerialize["releaseSdkVersion"] = o.ReleaseSdkVersion
	toSerialize["releaseServerVersionStandalone"] = o.ReleaseServerVersionStandalone
	toSerialize["sdkDeveloperFaqUrl"] = o.SdkDeveloperFaqUrl
	toSerialize["sdkDiscordUrl"] = o.SdkDiscordUrl
	toSerialize["sdkNotAllowedToPublishMessage"] = o.SdkNotAllowedToPublishMessage
	toSerialize["sdkUnityVersion"] = o.SdkUnityVersion
	toSerialize["serverName"] = o.ServerName
	toSerialize["supportEmail"] = o.SupportEmail
	toSerialize["timeOutWorldId"] = o.TimeOutWorldId
	toSerialize["tutorialWorldId"] = o.TutorialWorldId
	toSerialize["updateRateMsMaximum"] = o.UpdateRateMsMaximum
	toSerialize["updateRateMsMinimum"] = o.UpdateRateMsMinimum
	toSerialize["updateRateMsNormal"] = o.UpdateRateMsNormal
	toSerialize["updateRateMsUdonManual"] = o.UpdateRateMsUdonManual
	toSerialize["uploadAnalysisPercent"] = o.UploadAnalysisPercent
	toSerialize["urlList"] = o.UrlList
	toSerialize["useReliableUdpForVoice"] = o.UseReliableUdpForVoice
	toSerialize["userUpdatePeriod"] = o.UserUpdatePeriod
	toSerialize["userVerificationDelay"] = o.UserVerificationDelay
	toSerialize["userVerificationRetry"] = o.UserVerificationRetry
	toSerialize["userVerificationTimeout"] = o.UserVerificationTimeout
	toSerialize["viveWindowsUrl"] = o.ViveWindowsUrl
	toSerialize["whiteListedAssetUrls"] = o.WhiteListedAssetUrls
	toSerialize["worldUpdatePeriod"] = o.WorldUpdatePeriod
	toSerialize["player-url-resolver-hash"] = o.PlayerUrlResolverHash
	toSerialize["player-url-resolver-version"] = o.PlayerUrlResolverVersion
	return toSerialize, nil
}

type NullableAPIConfig struct {
	value *APIConfig
	isSet bool
}

func (v NullableAPIConfig) Get() *APIConfig {
	return v.value
}

func (v *NullableAPIConfig) Set(val *APIConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIConfig(val *APIConfig) *NullableAPIConfig {
	return &NullableAPIConfig{value: val, isSet: true}
}

func (v NullableAPIConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


