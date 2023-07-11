# APIConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VoiceEnableDegradation** | **bool** | Unknown, probably voice optimization testing | [default to false]
**VoiceEnableReceiverLimiting** | **bool** | Unknown, probably voice optimization testing | [default to true]
**Address** | **string** | VRChat&#39;s office address | 
**Announcements** | [**[]APIConfigAnnouncement**](APIConfigAnnouncement.md) | Public Announcements | 
**AppName** | **string** | Game name | [default to "VrChat"]
**BuildVersionTag** | **string** | Build tag of the API server | 
**ClientApiKey** | **string** | apiKey to be used for all other requests | 
**ClientBPSCeiling** | **int32** | Unknown | [default to 18432]
**ClientDisconnectTimeout** | **int32** | Unknown | [default to 30000]
**ClientReservedPlayerBPS** | **int32** | Unknown | [default to 7168]
**ClientSentCountAllowance** | **int32** | Unknown | [default to 15]
**ContactEmail** | **string** | VRChat&#39;s contact email | 
**CopyrightEmail** | **string** | VRChat&#39;s copyright-issues-related email | 
**CurrentTOSVersion** | **int32** | Current version number of the Terms of Service | 
**DefaultAvatar** | **string** |  | 
**DeploymentGroup** | [**DeploymentGroup**](DeploymentGroup.md) |  | [default to BLUE]
**DevAppVersionStandalone** | **string** | Version number for game development build | 
**DevDownloadLinkWindows** | **string** | Developer Download link | 
**DevSdkUrl** | **string** | Link to download the development SDK, use downloadUrls instead | 
**DevSdkVersion** | **string** | Version of the development SDK | 
**DevServerVersionStandalone** | **string** | Version number for server development build | 
**DisCountdown** | **time.Time** | Unknown, \&quot;dis\&quot; maybe for disconnect? | 
**DisableAvatarCopying** | **bool** | Toggles if copying avatars should be disabled | [default to false]
**DisableAvatarGating** | **bool** | Toggles if avatar gating should be disabled. Avatar gating restricts uploading of avatars to people with the &#x60;system_avatar_access&#x60; Tag or &#x60;admin_avatar_access&#x60; Tag | [default to false]
**DisableCommunityLabs** | **bool** | Toggles if the Community Labs should be disabled | [default to false]
**DisableCommunityLabsPromotion** | **bool** | Toggles if promotion out of Community Labs should be disabled | [default to false]
**DisableEmail** | **bool** | Unknown | [default to false]
**DisableEventStream** | **bool** | Toggles if Analytics should be disabled. | [default to false]
**DisableFeedbackGating** | **bool** | Toggles if feedback gating should be disabled. Feedback gating restricts submission of feedback (reporting a World or User) to people with the &#x60;system_feedback_access&#x60; Tag. | [default to false]
**DisableFrontendBuilds** | **bool** | Unknown, probably toggles compilation of frontend web builds? So internal flag? | [default to false]
**DisableHello** | **bool** | Unknown | [default to false]
**DisableOculusSubs** | **bool** | Toggles if signing up for Subscriptions in Oculus is disabled or not. | [default to false]
**DisableRegistration** | **bool** | Toggles if new user account registration should be disabled. | [default to false]
**DisableSteamNetworking** | **bool** | Toggles if Steam Networking should be disabled. VRChat these days uses Photon Unity Networking (PUN) instead. | [default to true]
**DisableTwoFactorAuth** | **bool** | Toggles if 2FA should be disabled. | [default to false]
**DisableUdon** | **bool** | Toggles if Udon should be universally disabled in-game. | [default to false]
**DisableUpgradeAccount** | **bool** | Toggles if account upgrading \&quot;linking with Steam/Oculus\&quot; should be disabled. | [default to false]
**DownloadLinkWindows** | **string** | Download link for game on the Oculus Rift website. | 
**DownloadUrls** | [**APIConfigDownloadURLList**](APIConfigDownloadURLList.md) |  | 
**DynamicWorldRows** | [**[]DynamicContentRow**](DynamicContentRow.md) | Array of DynamicWorldRow objects, used by the game to display the list of world rows | 
**Events** | [**APIConfigEvents**](APIConfigEvents.md) |  | 
**GearDemoRoomId** | **string** | Unknown | 
**HomeWorldId** | **string** | WorldID be \&quot;offline\&quot; on User profiles if you are not friends with that user. | 
**HomepageRedirectTarget** | **string** | Redirect target if you try to open the base API domain in your browser | [default to "https://hello.vrchat.com"]
**HubWorldId** | **string** | WorldID be \&quot;offline\&quot; on User profiles if you are not friends with that user. | 
**JobsEmail** | **string** | VRChat&#39;s job application email | 
**MessageOfTheDay** | **string** | MOTD | 
**ModerationEmail** | **string** | VRChat&#39;s moderation related email | 
**ModerationQueryPeriod** | **int32** | Unknown | 
**NotAllowedToSelectAvatarInPrivateWorldMessage** | **string** | Used in-game to notify a user they aren&#39;t allowed to select avatars in private worlds | 
**Plugin** | **string** | Extra [plugin](https://doc.photonengine.com/en-us/server/current/plugins/manual) to run in each instance | 
**ReleaseAppVersionStandalone** | **string** | Version number for game release build | 
**ReleaseSdkUrl** | **string** | Link to download the release SDK | 
**ReleaseSdkVersion** | **string** | Version of the release SDK | 
**ReleaseServerVersionStandalone** | **string** | Version number for server release build | 
**SdkDeveloperFaqUrl** | **string** | Link to the developer FAQ | 
**SdkDiscordUrl** | **string** | Link to the official VRChat Discord | 
**SdkNotAllowedToPublishMessage** | **string** | Used in the SDK to notify a user they aren&#39;t allowed to upload avatars/worlds yet | 
**SdkUnityVersion** | **string** | Unity version supported by the SDK | 
**ServerName** | **string** | Server name of the API server currently responding | 
**SupportEmail** | **string** | VRChat&#39;s support email | 
**TimeOutWorldId** | **string** | WorldID be \&quot;offline\&quot; on User profiles if you are not friends with that user. | 
**TutorialWorldId** | **string** | WorldID be \&quot;offline\&quot; on User profiles if you are not friends with that user. | 
**UpdateRateMsMaximum** | **int32** | Unknown | 
**UpdateRateMsMinimum** | **int32** | Unknown | 
**UpdateRateMsNormal** | **int32** | Unknown | 
**UpdateRateMsUdonManual** | **int32** | Unknown | 
**UploadAnalysisPercent** | **int32** | Unknown | 
**UrlList** | **[]string** | List of allowed URLs that bypass the \&quot;Allow untrusted URL&#39;s\&quot; setting in-game | 
**UseReliableUdpForVoice** | **bool** | Unknown | [default to false]
**UserUpdatePeriod** | **int32** | Unknown | 
**UserVerificationDelay** | **int32** | Unknown | 
**UserVerificationRetry** | **int32** | Unknown | 
**UserVerificationTimeout** | **int32** | Unknown | 
**ViveWindowsUrl** | **string** | Download link for game on the Steam website. | 
**WhiteListedAssetUrls** | **[]string** | List of allowed URLs that are allowed to host avatar assets | 
**WorldUpdatePeriod** | **int32** | Unknown | 
**PlayerUrlResolverHash** | **string** | Currently used youtube-dl.exe hash in SHA-256-delimited format | 
**PlayerUrlResolverVersion** | **string** | Currently used youtube-dl.exe version | 

## Methods

### NewAPIConfig

`func NewAPIConfig(voiceEnableDegradation bool, voiceEnableReceiverLimiting bool, address string, announcements []APIConfigAnnouncement, appName string, buildVersionTag string, clientApiKey string, clientBPSCeiling int32, clientDisconnectTimeout int32, clientReservedPlayerBPS int32, clientSentCountAllowance int32, contactEmail string, copyrightEmail string, currentTOSVersion int32, defaultAvatar string, deploymentGroup DeploymentGroup, devAppVersionStandalone string, devDownloadLinkWindows string, devSdkUrl string, devSdkVersion string, devServerVersionStandalone string, disCountdown time.Time, disableAvatarCopying bool, disableAvatarGating bool, disableCommunityLabs bool, disableCommunityLabsPromotion bool, disableEmail bool, disableEventStream bool, disableFeedbackGating bool, disableFrontendBuilds bool, disableHello bool, disableOculusSubs bool, disableRegistration bool, disableSteamNetworking bool, disableTwoFactorAuth bool, disableUdon bool, disableUpgradeAccount bool, downloadLinkWindows string, downloadUrls APIConfigDownloadURLList, dynamicWorldRows []DynamicContentRow, events APIConfigEvents, gearDemoRoomId string, homeWorldId string, homepageRedirectTarget string, hubWorldId string, jobsEmail string, messageOfTheDay string, moderationEmail string, moderationQueryPeriod int32, notAllowedToSelectAvatarInPrivateWorldMessage string, plugin string, releaseAppVersionStandalone string, releaseSdkUrl string, releaseSdkVersion string, releaseServerVersionStandalone string, sdkDeveloperFaqUrl string, sdkDiscordUrl string, sdkNotAllowedToPublishMessage string, sdkUnityVersion string, serverName string, supportEmail string, timeOutWorldId string, tutorialWorldId string, updateRateMsMaximum int32, updateRateMsMinimum int32, updateRateMsNormal int32, updateRateMsUdonManual int32, uploadAnalysisPercent int32, urlList []string, useReliableUdpForVoice bool, userUpdatePeriod int32, userVerificationDelay int32, userVerificationRetry int32, userVerificationTimeout int32, viveWindowsUrl string, whiteListedAssetUrls []string, worldUpdatePeriod int32, playerUrlResolverHash string, playerUrlResolverVersion string, ) *APIConfig`

NewAPIConfig instantiates a new APIConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIConfigWithDefaults

`func NewAPIConfigWithDefaults() *APIConfig`

NewAPIConfigWithDefaults instantiates a new APIConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVoiceEnableDegradation

`func (o *APIConfig) GetVoiceEnableDegradation() bool`

GetVoiceEnableDegradation returns the VoiceEnableDegradation field if non-nil, zero value otherwise.

### GetVoiceEnableDegradationOk

`func (o *APIConfig) GetVoiceEnableDegradationOk() (*bool, bool)`

GetVoiceEnableDegradationOk returns a tuple with the VoiceEnableDegradation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceEnableDegradation

`func (o *APIConfig) SetVoiceEnableDegradation(v bool)`

SetVoiceEnableDegradation sets VoiceEnableDegradation field to given value.


### GetVoiceEnableReceiverLimiting

`func (o *APIConfig) GetVoiceEnableReceiverLimiting() bool`

GetVoiceEnableReceiverLimiting returns the VoiceEnableReceiverLimiting field if non-nil, zero value otherwise.

### GetVoiceEnableReceiverLimitingOk

`func (o *APIConfig) GetVoiceEnableReceiverLimitingOk() (*bool, bool)`

GetVoiceEnableReceiverLimitingOk returns a tuple with the VoiceEnableReceiverLimiting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceEnableReceiverLimiting

`func (o *APIConfig) SetVoiceEnableReceiverLimiting(v bool)`

SetVoiceEnableReceiverLimiting sets VoiceEnableReceiverLimiting field to given value.


### GetAddress

`func (o *APIConfig) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *APIConfig) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *APIConfig) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetAnnouncements

`func (o *APIConfig) GetAnnouncements() []APIConfigAnnouncement`

GetAnnouncements returns the Announcements field if non-nil, zero value otherwise.

### GetAnnouncementsOk

`func (o *APIConfig) GetAnnouncementsOk() (*[]APIConfigAnnouncement, bool)`

GetAnnouncementsOk returns a tuple with the Announcements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnouncements

`func (o *APIConfig) SetAnnouncements(v []APIConfigAnnouncement)`

SetAnnouncements sets Announcements field to given value.


### GetAppName

`func (o *APIConfig) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *APIConfig) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *APIConfig) SetAppName(v string)`

SetAppName sets AppName field to given value.


### GetBuildVersionTag

`func (o *APIConfig) GetBuildVersionTag() string`

GetBuildVersionTag returns the BuildVersionTag field if non-nil, zero value otherwise.

### GetBuildVersionTagOk

`func (o *APIConfig) GetBuildVersionTagOk() (*string, bool)`

GetBuildVersionTagOk returns a tuple with the BuildVersionTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildVersionTag

`func (o *APIConfig) SetBuildVersionTag(v string)`

SetBuildVersionTag sets BuildVersionTag field to given value.


### GetClientApiKey

`func (o *APIConfig) GetClientApiKey() string`

GetClientApiKey returns the ClientApiKey field if non-nil, zero value otherwise.

### GetClientApiKeyOk

`func (o *APIConfig) GetClientApiKeyOk() (*string, bool)`

GetClientApiKeyOk returns a tuple with the ClientApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientApiKey

`func (o *APIConfig) SetClientApiKey(v string)`

SetClientApiKey sets ClientApiKey field to given value.


### GetClientBPSCeiling

`func (o *APIConfig) GetClientBPSCeiling() int32`

GetClientBPSCeiling returns the ClientBPSCeiling field if non-nil, zero value otherwise.

### GetClientBPSCeilingOk

`func (o *APIConfig) GetClientBPSCeilingOk() (*int32, bool)`

GetClientBPSCeilingOk returns a tuple with the ClientBPSCeiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientBPSCeiling

`func (o *APIConfig) SetClientBPSCeiling(v int32)`

SetClientBPSCeiling sets ClientBPSCeiling field to given value.


### GetClientDisconnectTimeout

`func (o *APIConfig) GetClientDisconnectTimeout() int32`

GetClientDisconnectTimeout returns the ClientDisconnectTimeout field if non-nil, zero value otherwise.

### GetClientDisconnectTimeoutOk

`func (o *APIConfig) GetClientDisconnectTimeoutOk() (*int32, bool)`

GetClientDisconnectTimeoutOk returns a tuple with the ClientDisconnectTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientDisconnectTimeout

`func (o *APIConfig) SetClientDisconnectTimeout(v int32)`

SetClientDisconnectTimeout sets ClientDisconnectTimeout field to given value.


### GetClientReservedPlayerBPS

`func (o *APIConfig) GetClientReservedPlayerBPS() int32`

GetClientReservedPlayerBPS returns the ClientReservedPlayerBPS field if non-nil, zero value otherwise.

### GetClientReservedPlayerBPSOk

`func (o *APIConfig) GetClientReservedPlayerBPSOk() (*int32, bool)`

GetClientReservedPlayerBPSOk returns a tuple with the ClientReservedPlayerBPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientReservedPlayerBPS

`func (o *APIConfig) SetClientReservedPlayerBPS(v int32)`

SetClientReservedPlayerBPS sets ClientReservedPlayerBPS field to given value.


### GetClientSentCountAllowance

`func (o *APIConfig) GetClientSentCountAllowance() int32`

GetClientSentCountAllowance returns the ClientSentCountAllowance field if non-nil, zero value otherwise.

### GetClientSentCountAllowanceOk

`func (o *APIConfig) GetClientSentCountAllowanceOk() (*int32, bool)`

GetClientSentCountAllowanceOk returns a tuple with the ClientSentCountAllowance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSentCountAllowance

`func (o *APIConfig) SetClientSentCountAllowance(v int32)`

SetClientSentCountAllowance sets ClientSentCountAllowance field to given value.


### GetContactEmail

`func (o *APIConfig) GetContactEmail() string`

GetContactEmail returns the ContactEmail field if non-nil, zero value otherwise.

### GetContactEmailOk

`func (o *APIConfig) GetContactEmailOk() (*string, bool)`

GetContactEmailOk returns a tuple with the ContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactEmail

`func (o *APIConfig) SetContactEmail(v string)`

SetContactEmail sets ContactEmail field to given value.


### GetCopyrightEmail

`func (o *APIConfig) GetCopyrightEmail() string`

GetCopyrightEmail returns the CopyrightEmail field if non-nil, zero value otherwise.

### GetCopyrightEmailOk

`func (o *APIConfig) GetCopyrightEmailOk() (*string, bool)`

GetCopyrightEmailOk returns a tuple with the CopyrightEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyrightEmail

`func (o *APIConfig) SetCopyrightEmail(v string)`

SetCopyrightEmail sets CopyrightEmail field to given value.


### GetCurrentTOSVersion

`func (o *APIConfig) GetCurrentTOSVersion() int32`

GetCurrentTOSVersion returns the CurrentTOSVersion field if non-nil, zero value otherwise.

### GetCurrentTOSVersionOk

`func (o *APIConfig) GetCurrentTOSVersionOk() (*int32, bool)`

GetCurrentTOSVersionOk returns a tuple with the CurrentTOSVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentTOSVersion

`func (o *APIConfig) SetCurrentTOSVersion(v int32)`

SetCurrentTOSVersion sets CurrentTOSVersion field to given value.


### GetDefaultAvatar

`func (o *APIConfig) GetDefaultAvatar() string`

GetDefaultAvatar returns the DefaultAvatar field if non-nil, zero value otherwise.

### GetDefaultAvatarOk

`func (o *APIConfig) GetDefaultAvatarOk() (*string, bool)`

GetDefaultAvatarOk returns a tuple with the DefaultAvatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultAvatar

`func (o *APIConfig) SetDefaultAvatar(v string)`

SetDefaultAvatar sets DefaultAvatar field to given value.


### GetDeploymentGroup

`func (o *APIConfig) GetDeploymentGroup() DeploymentGroup`

GetDeploymentGroup returns the DeploymentGroup field if non-nil, zero value otherwise.

### GetDeploymentGroupOk

`func (o *APIConfig) GetDeploymentGroupOk() (*DeploymentGroup, bool)`

GetDeploymentGroupOk returns a tuple with the DeploymentGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentGroup

`func (o *APIConfig) SetDeploymentGroup(v DeploymentGroup)`

SetDeploymentGroup sets DeploymentGroup field to given value.


### GetDevAppVersionStandalone

`func (o *APIConfig) GetDevAppVersionStandalone() string`

GetDevAppVersionStandalone returns the DevAppVersionStandalone field if non-nil, zero value otherwise.

### GetDevAppVersionStandaloneOk

`func (o *APIConfig) GetDevAppVersionStandaloneOk() (*string, bool)`

GetDevAppVersionStandaloneOk returns a tuple with the DevAppVersionStandalone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevAppVersionStandalone

`func (o *APIConfig) SetDevAppVersionStandalone(v string)`

SetDevAppVersionStandalone sets DevAppVersionStandalone field to given value.


### GetDevDownloadLinkWindows

`func (o *APIConfig) GetDevDownloadLinkWindows() string`

GetDevDownloadLinkWindows returns the DevDownloadLinkWindows field if non-nil, zero value otherwise.

### GetDevDownloadLinkWindowsOk

`func (o *APIConfig) GetDevDownloadLinkWindowsOk() (*string, bool)`

GetDevDownloadLinkWindowsOk returns a tuple with the DevDownloadLinkWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevDownloadLinkWindows

`func (o *APIConfig) SetDevDownloadLinkWindows(v string)`

SetDevDownloadLinkWindows sets DevDownloadLinkWindows field to given value.


### GetDevSdkUrl

`func (o *APIConfig) GetDevSdkUrl() string`

GetDevSdkUrl returns the DevSdkUrl field if non-nil, zero value otherwise.

### GetDevSdkUrlOk

`func (o *APIConfig) GetDevSdkUrlOk() (*string, bool)`

GetDevSdkUrlOk returns a tuple with the DevSdkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevSdkUrl

`func (o *APIConfig) SetDevSdkUrl(v string)`

SetDevSdkUrl sets DevSdkUrl field to given value.


### GetDevSdkVersion

`func (o *APIConfig) GetDevSdkVersion() string`

GetDevSdkVersion returns the DevSdkVersion field if non-nil, zero value otherwise.

### GetDevSdkVersionOk

`func (o *APIConfig) GetDevSdkVersionOk() (*string, bool)`

GetDevSdkVersionOk returns a tuple with the DevSdkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevSdkVersion

`func (o *APIConfig) SetDevSdkVersion(v string)`

SetDevSdkVersion sets DevSdkVersion field to given value.


### GetDevServerVersionStandalone

`func (o *APIConfig) GetDevServerVersionStandalone() string`

GetDevServerVersionStandalone returns the DevServerVersionStandalone field if non-nil, zero value otherwise.

### GetDevServerVersionStandaloneOk

`func (o *APIConfig) GetDevServerVersionStandaloneOk() (*string, bool)`

GetDevServerVersionStandaloneOk returns a tuple with the DevServerVersionStandalone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevServerVersionStandalone

`func (o *APIConfig) SetDevServerVersionStandalone(v string)`

SetDevServerVersionStandalone sets DevServerVersionStandalone field to given value.


### GetDisCountdown

`func (o *APIConfig) GetDisCountdown() time.Time`

GetDisCountdown returns the DisCountdown field if non-nil, zero value otherwise.

### GetDisCountdownOk

`func (o *APIConfig) GetDisCountdownOk() (*time.Time, bool)`

GetDisCountdownOk returns a tuple with the DisCountdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisCountdown

`func (o *APIConfig) SetDisCountdown(v time.Time)`

SetDisCountdown sets DisCountdown field to given value.


### GetDisableAvatarCopying

`func (o *APIConfig) GetDisableAvatarCopying() bool`

GetDisableAvatarCopying returns the DisableAvatarCopying field if non-nil, zero value otherwise.

### GetDisableAvatarCopyingOk

`func (o *APIConfig) GetDisableAvatarCopyingOk() (*bool, bool)`

GetDisableAvatarCopyingOk returns a tuple with the DisableAvatarCopying field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAvatarCopying

`func (o *APIConfig) SetDisableAvatarCopying(v bool)`

SetDisableAvatarCopying sets DisableAvatarCopying field to given value.


### GetDisableAvatarGating

`func (o *APIConfig) GetDisableAvatarGating() bool`

GetDisableAvatarGating returns the DisableAvatarGating field if non-nil, zero value otherwise.

### GetDisableAvatarGatingOk

`func (o *APIConfig) GetDisableAvatarGatingOk() (*bool, bool)`

GetDisableAvatarGatingOk returns a tuple with the DisableAvatarGating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableAvatarGating

`func (o *APIConfig) SetDisableAvatarGating(v bool)`

SetDisableAvatarGating sets DisableAvatarGating field to given value.


### GetDisableCommunityLabs

`func (o *APIConfig) GetDisableCommunityLabs() bool`

GetDisableCommunityLabs returns the DisableCommunityLabs field if non-nil, zero value otherwise.

### GetDisableCommunityLabsOk

`func (o *APIConfig) GetDisableCommunityLabsOk() (*bool, bool)`

GetDisableCommunityLabsOk returns a tuple with the DisableCommunityLabs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableCommunityLabs

`func (o *APIConfig) SetDisableCommunityLabs(v bool)`

SetDisableCommunityLabs sets DisableCommunityLabs field to given value.


### GetDisableCommunityLabsPromotion

`func (o *APIConfig) GetDisableCommunityLabsPromotion() bool`

GetDisableCommunityLabsPromotion returns the DisableCommunityLabsPromotion field if non-nil, zero value otherwise.

### GetDisableCommunityLabsPromotionOk

`func (o *APIConfig) GetDisableCommunityLabsPromotionOk() (*bool, bool)`

GetDisableCommunityLabsPromotionOk returns a tuple with the DisableCommunityLabsPromotion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableCommunityLabsPromotion

`func (o *APIConfig) SetDisableCommunityLabsPromotion(v bool)`

SetDisableCommunityLabsPromotion sets DisableCommunityLabsPromotion field to given value.


### GetDisableEmail

`func (o *APIConfig) GetDisableEmail() bool`

GetDisableEmail returns the DisableEmail field if non-nil, zero value otherwise.

### GetDisableEmailOk

`func (o *APIConfig) GetDisableEmailOk() (*bool, bool)`

GetDisableEmailOk returns a tuple with the DisableEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEmail

`func (o *APIConfig) SetDisableEmail(v bool)`

SetDisableEmail sets DisableEmail field to given value.


### GetDisableEventStream

`func (o *APIConfig) GetDisableEventStream() bool`

GetDisableEventStream returns the DisableEventStream field if non-nil, zero value otherwise.

### GetDisableEventStreamOk

`func (o *APIConfig) GetDisableEventStreamOk() (*bool, bool)`

GetDisableEventStreamOk returns a tuple with the DisableEventStream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableEventStream

`func (o *APIConfig) SetDisableEventStream(v bool)`

SetDisableEventStream sets DisableEventStream field to given value.


### GetDisableFeedbackGating

`func (o *APIConfig) GetDisableFeedbackGating() bool`

GetDisableFeedbackGating returns the DisableFeedbackGating field if non-nil, zero value otherwise.

### GetDisableFeedbackGatingOk

`func (o *APIConfig) GetDisableFeedbackGatingOk() (*bool, bool)`

GetDisableFeedbackGatingOk returns a tuple with the DisableFeedbackGating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableFeedbackGating

`func (o *APIConfig) SetDisableFeedbackGating(v bool)`

SetDisableFeedbackGating sets DisableFeedbackGating field to given value.


### GetDisableFrontendBuilds

`func (o *APIConfig) GetDisableFrontendBuilds() bool`

GetDisableFrontendBuilds returns the DisableFrontendBuilds field if non-nil, zero value otherwise.

### GetDisableFrontendBuildsOk

`func (o *APIConfig) GetDisableFrontendBuildsOk() (*bool, bool)`

GetDisableFrontendBuildsOk returns a tuple with the DisableFrontendBuilds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableFrontendBuilds

`func (o *APIConfig) SetDisableFrontendBuilds(v bool)`

SetDisableFrontendBuilds sets DisableFrontendBuilds field to given value.


### GetDisableHello

`func (o *APIConfig) GetDisableHello() bool`

GetDisableHello returns the DisableHello field if non-nil, zero value otherwise.

### GetDisableHelloOk

`func (o *APIConfig) GetDisableHelloOk() (*bool, bool)`

GetDisableHelloOk returns a tuple with the DisableHello field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableHello

`func (o *APIConfig) SetDisableHello(v bool)`

SetDisableHello sets DisableHello field to given value.


### GetDisableOculusSubs

`func (o *APIConfig) GetDisableOculusSubs() bool`

GetDisableOculusSubs returns the DisableOculusSubs field if non-nil, zero value otherwise.

### GetDisableOculusSubsOk

`func (o *APIConfig) GetDisableOculusSubsOk() (*bool, bool)`

GetDisableOculusSubsOk returns a tuple with the DisableOculusSubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableOculusSubs

`func (o *APIConfig) SetDisableOculusSubs(v bool)`

SetDisableOculusSubs sets DisableOculusSubs field to given value.


### GetDisableRegistration

`func (o *APIConfig) GetDisableRegistration() bool`

GetDisableRegistration returns the DisableRegistration field if non-nil, zero value otherwise.

### GetDisableRegistrationOk

`func (o *APIConfig) GetDisableRegistrationOk() (*bool, bool)`

GetDisableRegistrationOk returns a tuple with the DisableRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableRegistration

`func (o *APIConfig) SetDisableRegistration(v bool)`

SetDisableRegistration sets DisableRegistration field to given value.


### GetDisableSteamNetworking

`func (o *APIConfig) GetDisableSteamNetworking() bool`

GetDisableSteamNetworking returns the DisableSteamNetworking field if non-nil, zero value otherwise.

### GetDisableSteamNetworkingOk

`func (o *APIConfig) GetDisableSteamNetworkingOk() (*bool, bool)`

GetDisableSteamNetworkingOk returns a tuple with the DisableSteamNetworking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableSteamNetworking

`func (o *APIConfig) SetDisableSteamNetworking(v bool)`

SetDisableSteamNetworking sets DisableSteamNetworking field to given value.


### GetDisableTwoFactorAuth

`func (o *APIConfig) GetDisableTwoFactorAuth() bool`

GetDisableTwoFactorAuth returns the DisableTwoFactorAuth field if non-nil, zero value otherwise.

### GetDisableTwoFactorAuthOk

`func (o *APIConfig) GetDisableTwoFactorAuthOk() (*bool, bool)`

GetDisableTwoFactorAuthOk returns a tuple with the DisableTwoFactorAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableTwoFactorAuth

`func (o *APIConfig) SetDisableTwoFactorAuth(v bool)`

SetDisableTwoFactorAuth sets DisableTwoFactorAuth field to given value.


### GetDisableUdon

`func (o *APIConfig) GetDisableUdon() bool`

GetDisableUdon returns the DisableUdon field if non-nil, zero value otherwise.

### GetDisableUdonOk

`func (o *APIConfig) GetDisableUdonOk() (*bool, bool)`

GetDisableUdonOk returns a tuple with the DisableUdon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableUdon

`func (o *APIConfig) SetDisableUdon(v bool)`

SetDisableUdon sets DisableUdon field to given value.


### GetDisableUpgradeAccount

`func (o *APIConfig) GetDisableUpgradeAccount() bool`

GetDisableUpgradeAccount returns the DisableUpgradeAccount field if non-nil, zero value otherwise.

### GetDisableUpgradeAccountOk

`func (o *APIConfig) GetDisableUpgradeAccountOk() (*bool, bool)`

GetDisableUpgradeAccountOk returns a tuple with the DisableUpgradeAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableUpgradeAccount

`func (o *APIConfig) SetDisableUpgradeAccount(v bool)`

SetDisableUpgradeAccount sets DisableUpgradeAccount field to given value.


### GetDownloadLinkWindows

`func (o *APIConfig) GetDownloadLinkWindows() string`

GetDownloadLinkWindows returns the DownloadLinkWindows field if non-nil, zero value otherwise.

### GetDownloadLinkWindowsOk

`func (o *APIConfig) GetDownloadLinkWindowsOk() (*string, bool)`

GetDownloadLinkWindowsOk returns a tuple with the DownloadLinkWindows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadLinkWindows

`func (o *APIConfig) SetDownloadLinkWindows(v string)`

SetDownloadLinkWindows sets DownloadLinkWindows field to given value.


### GetDownloadUrls

`func (o *APIConfig) GetDownloadUrls() APIConfigDownloadURLList`

GetDownloadUrls returns the DownloadUrls field if non-nil, zero value otherwise.

### GetDownloadUrlsOk

`func (o *APIConfig) GetDownloadUrlsOk() (*APIConfigDownloadURLList, bool)`

GetDownloadUrlsOk returns a tuple with the DownloadUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadUrls

`func (o *APIConfig) SetDownloadUrls(v APIConfigDownloadURLList)`

SetDownloadUrls sets DownloadUrls field to given value.


### GetDynamicWorldRows

`func (o *APIConfig) GetDynamicWorldRows() []DynamicContentRow`

GetDynamicWorldRows returns the DynamicWorldRows field if non-nil, zero value otherwise.

### GetDynamicWorldRowsOk

`func (o *APIConfig) GetDynamicWorldRowsOk() (*[]DynamicContentRow, bool)`

GetDynamicWorldRowsOk returns a tuple with the DynamicWorldRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicWorldRows

`func (o *APIConfig) SetDynamicWorldRows(v []DynamicContentRow)`

SetDynamicWorldRows sets DynamicWorldRows field to given value.


### GetEvents

`func (o *APIConfig) GetEvents() APIConfigEvents`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *APIConfig) GetEventsOk() (*APIConfigEvents, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *APIConfig) SetEvents(v APIConfigEvents)`

SetEvents sets Events field to given value.


### GetGearDemoRoomId

`func (o *APIConfig) GetGearDemoRoomId() string`

GetGearDemoRoomId returns the GearDemoRoomId field if non-nil, zero value otherwise.

### GetGearDemoRoomIdOk

`func (o *APIConfig) GetGearDemoRoomIdOk() (*string, bool)`

GetGearDemoRoomIdOk returns a tuple with the GearDemoRoomId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGearDemoRoomId

`func (o *APIConfig) SetGearDemoRoomId(v string)`

SetGearDemoRoomId sets GearDemoRoomId field to given value.


### GetHomeWorldId

`func (o *APIConfig) GetHomeWorldId() string`

GetHomeWorldId returns the HomeWorldId field if non-nil, zero value otherwise.

### GetHomeWorldIdOk

`func (o *APIConfig) GetHomeWorldIdOk() (*string, bool)`

GetHomeWorldIdOk returns a tuple with the HomeWorldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeWorldId

`func (o *APIConfig) SetHomeWorldId(v string)`

SetHomeWorldId sets HomeWorldId field to given value.


### GetHomepageRedirectTarget

`func (o *APIConfig) GetHomepageRedirectTarget() string`

GetHomepageRedirectTarget returns the HomepageRedirectTarget field if non-nil, zero value otherwise.

### GetHomepageRedirectTargetOk

`func (o *APIConfig) GetHomepageRedirectTargetOk() (*string, bool)`

GetHomepageRedirectTargetOk returns a tuple with the HomepageRedirectTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepageRedirectTarget

`func (o *APIConfig) SetHomepageRedirectTarget(v string)`

SetHomepageRedirectTarget sets HomepageRedirectTarget field to given value.


### GetHubWorldId

`func (o *APIConfig) GetHubWorldId() string`

GetHubWorldId returns the HubWorldId field if non-nil, zero value otherwise.

### GetHubWorldIdOk

`func (o *APIConfig) GetHubWorldIdOk() (*string, bool)`

GetHubWorldIdOk returns a tuple with the HubWorldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHubWorldId

`func (o *APIConfig) SetHubWorldId(v string)`

SetHubWorldId sets HubWorldId field to given value.


### GetJobsEmail

`func (o *APIConfig) GetJobsEmail() string`

GetJobsEmail returns the JobsEmail field if non-nil, zero value otherwise.

### GetJobsEmailOk

`func (o *APIConfig) GetJobsEmailOk() (*string, bool)`

GetJobsEmailOk returns a tuple with the JobsEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobsEmail

`func (o *APIConfig) SetJobsEmail(v string)`

SetJobsEmail sets JobsEmail field to given value.


### GetMessageOfTheDay

`func (o *APIConfig) GetMessageOfTheDay() string`

GetMessageOfTheDay returns the MessageOfTheDay field if non-nil, zero value otherwise.

### GetMessageOfTheDayOk

`func (o *APIConfig) GetMessageOfTheDayOk() (*string, bool)`

GetMessageOfTheDayOk returns a tuple with the MessageOfTheDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageOfTheDay

`func (o *APIConfig) SetMessageOfTheDay(v string)`

SetMessageOfTheDay sets MessageOfTheDay field to given value.


### GetModerationEmail

`func (o *APIConfig) GetModerationEmail() string`

GetModerationEmail returns the ModerationEmail field if non-nil, zero value otherwise.

### GetModerationEmailOk

`func (o *APIConfig) GetModerationEmailOk() (*string, bool)`

GetModerationEmailOk returns a tuple with the ModerationEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationEmail

`func (o *APIConfig) SetModerationEmail(v string)`

SetModerationEmail sets ModerationEmail field to given value.


### GetModerationQueryPeriod

`func (o *APIConfig) GetModerationQueryPeriod() int32`

GetModerationQueryPeriod returns the ModerationQueryPeriod field if non-nil, zero value otherwise.

### GetModerationQueryPeriodOk

`func (o *APIConfig) GetModerationQueryPeriodOk() (*int32, bool)`

GetModerationQueryPeriodOk returns a tuple with the ModerationQueryPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModerationQueryPeriod

`func (o *APIConfig) SetModerationQueryPeriod(v int32)`

SetModerationQueryPeriod sets ModerationQueryPeriod field to given value.


### GetNotAllowedToSelectAvatarInPrivateWorldMessage

`func (o *APIConfig) GetNotAllowedToSelectAvatarInPrivateWorldMessage() string`

GetNotAllowedToSelectAvatarInPrivateWorldMessage returns the NotAllowedToSelectAvatarInPrivateWorldMessage field if non-nil, zero value otherwise.

### GetNotAllowedToSelectAvatarInPrivateWorldMessageOk

`func (o *APIConfig) GetNotAllowedToSelectAvatarInPrivateWorldMessageOk() (*string, bool)`

GetNotAllowedToSelectAvatarInPrivateWorldMessageOk returns a tuple with the NotAllowedToSelectAvatarInPrivateWorldMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAllowedToSelectAvatarInPrivateWorldMessage

`func (o *APIConfig) SetNotAllowedToSelectAvatarInPrivateWorldMessage(v string)`

SetNotAllowedToSelectAvatarInPrivateWorldMessage sets NotAllowedToSelectAvatarInPrivateWorldMessage field to given value.


### GetPlugin

`func (o *APIConfig) GetPlugin() string`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *APIConfig) GetPluginOk() (*string, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *APIConfig) SetPlugin(v string)`

SetPlugin sets Plugin field to given value.


### GetReleaseAppVersionStandalone

`func (o *APIConfig) GetReleaseAppVersionStandalone() string`

GetReleaseAppVersionStandalone returns the ReleaseAppVersionStandalone field if non-nil, zero value otherwise.

### GetReleaseAppVersionStandaloneOk

`func (o *APIConfig) GetReleaseAppVersionStandaloneOk() (*string, bool)`

GetReleaseAppVersionStandaloneOk returns a tuple with the ReleaseAppVersionStandalone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseAppVersionStandalone

`func (o *APIConfig) SetReleaseAppVersionStandalone(v string)`

SetReleaseAppVersionStandalone sets ReleaseAppVersionStandalone field to given value.


### GetReleaseSdkUrl

`func (o *APIConfig) GetReleaseSdkUrl() string`

GetReleaseSdkUrl returns the ReleaseSdkUrl field if non-nil, zero value otherwise.

### GetReleaseSdkUrlOk

`func (o *APIConfig) GetReleaseSdkUrlOk() (*string, bool)`

GetReleaseSdkUrlOk returns a tuple with the ReleaseSdkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseSdkUrl

`func (o *APIConfig) SetReleaseSdkUrl(v string)`

SetReleaseSdkUrl sets ReleaseSdkUrl field to given value.


### GetReleaseSdkVersion

`func (o *APIConfig) GetReleaseSdkVersion() string`

GetReleaseSdkVersion returns the ReleaseSdkVersion field if non-nil, zero value otherwise.

### GetReleaseSdkVersionOk

`func (o *APIConfig) GetReleaseSdkVersionOk() (*string, bool)`

GetReleaseSdkVersionOk returns a tuple with the ReleaseSdkVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseSdkVersion

`func (o *APIConfig) SetReleaseSdkVersion(v string)`

SetReleaseSdkVersion sets ReleaseSdkVersion field to given value.


### GetReleaseServerVersionStandalone

`func (o *APIConfig) GetReleaseServerVersionStandalone() string`

GetReleaseServerVersionStandalone returns the ReleaseServerVersionStandalone field if non-nil, zero value otherwise.

### GetReleaseServerVersionStandaloneOk

`func (o *APIConfig) GetReleaseServerVersionStandaloneOk() (*string, bool)`

GetReleaseServerVersionStandaloneOk returns a tuple with the ReleaseServerVersionStandalone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseServerVersionStandalone

`func (o *APIConfig) SetReleaseServerVersionStandalone(v string)`

SetReleaseServerVersionStandalone sets ReleaseServerVersionStandalone field to given value.


### GetSdkDeveloperFaqUrl

`func (o *APIConfig) GetSdkDeveloperFaqUrl() string`

GetSdkDeveloperFaqUrl returns the SdkDeveloperFaqUrl field if non-nil, zero value otherwise.

### GetSdkDeveloperFaqUrlOk

`func (o *APIConfig) GetSdkDeveloperFaqUrlOk() (*string, bool)`

GetSdkDeveloperFaqUrlOk returns a tuple with the SdkDeveloperFaqUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkDeveloperFaqUrl

`func (o *APIConfig) SetSdkDeveloperFaqUrl(v string)`

SetSdkDeveloperFaqUrl sets SdkDeveloperFaqUrl field to given value.


### GetSdkDiscordUrl

`func (o *APIConfig) GetSdkDiscordUrl() string`

GetSdkDiscordUrl returns the SdkDiscordUrl field if non-nil, zero value otherwise.

### GetSdkDiscordUrlOk

`func (o *APIConfig) GetSdkDiscordUrlOk() (*string, bool)`

GetSdkDiscordUrlOk returns a tuple with the SdkDiscordUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkDiscordUrl

`func (o *APIConfig) SetSdkDiscordUrl(v string)`

SetSdkDiscordUrl sets SdkDiscordUrl field to given value.


### GetSdkNotAllowedToPublishMessage

`func (o *APIConfig) GetSdkNotAllowedToPublishMessage() string`

GetSdkNotAllowedToPublishMessage returns the SdkNotAllowedToPublishMessage field if non-nil, zero value otherwise.

### GetSdkNotAllowedToPublishMessageOk

`func (o *APIConfig) GetSdkNotAllowedToPublishMessageOk() (*string, bool)`

GetSdkNotAllowedToPublishMessageOk returns a tuple with the SdkNotAllowedToPublishMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkNotAllowedToPublishMessage

`func (o *APIConfig) SetSdkNotAllowedToPublishMessage(v string)`

SetSdkNotAllowedToPublishMessage sets SdkNotAllowedToPublishMessage field to given value.


### GetSdkUnityVersion

`func (o *APIConfig) GetSdkUnityVersion() string`

GetSdkUnityVersion returns the SdkUnityVersion field if non-nil, zero value otherwise.

### GetSdkUnityVersionOk

`func (o *APIConfig) GetSdkUnityVersionOk() (*string, bool)`

GetSdkUnityVersionOk returns a tuple with the SdkUnityVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdkUnityVersion

`func (o *APIConfig) SetSdkUnityVersion(v string)`

SetSdkUnityVersion sets SdkUnityVersion field to given value.


### GetServerName

`func (o *APIConfig) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *APIConfig) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *APIConfig) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### GetSupportEmail

`func (o *APIConfig) GetSupportEmail() string`

GetSupportEmail returns the SupportEmail field if non-nil, zero value otherwise.

### GetSupportEmailOk

`func (o *APIConfig) GetSupportEmailOk() (*string, bool)`

GetSupportEmailOk returns a tuple with the SupportEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportEmail

`func (o *APIConfig) SetSupportEmail(v string)`

SetSupportEmail sets SupportEmail field to given value.


### GetTimeOutWorldId

`func (o *APIConfig) GetTimeOutWorldId() string`

GetTimeOutWorldId returns the TimeOutWorldId field if non-nil, zero value otherwise.

### GetTimeOutWorldIdOk

`func (o *APIConfig) GetTimeOutWorldIdOk() (*string, bool)`

GetTimeOutWorldIdOk returns a tuple with the TimeOutWorldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeOutWorldId

`func (o *APIConfig) SetTimeOutWorldId(v string)`

SetTimeOutWorldId sets TimeOutWorldId field to given value.


### GetTutorialWorldId

`func (o *APIConfig) GetTutorialWorldId() string`

GetTutorialWorldId returns the TutorialWorldId field if non-nil, zero value otherwise.

### GetTutorialWorldIdOk

`func (o *APIConfig) GetTutorialWorldIdOk() (*string, bool)`

GetTutorialWorldIdOk returns a tuple with the TutorialWorldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTutorialWorldId

`func (o *APIConfig) SetTutorialWorldId(v string)`

SetTutorialWorldId sets TutorialWorldId field to given value.


### GetUpdateRateMsMaximum

`func (o *APIConfig) GetUpdateRateMsMaximum() int32`

GetUpdateRateMsMaximum returns the UpdateRateMsMaximum field if non-nil, zero value otherwise.

### GetUpdateRateMsMaximumOk

`func (o *APIConfig) GetUpdateRateMsMaximumOk() (*int32, bool)`

GetUpdateRateMsMaximumOk returns a tuple with the UpdateRateMsMaximum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateRateMsMaximum

`func (o *APIConfig) SetUpdateRateMsMaximum(v int32)`

SetUpdateRateMsMaximum sets UpdateRateMsMaximum field to given value.


### GetUpdateRateMsMinimum

`func (o *APIConfig) GetUpdateRateMsMinimum() int32`

GetUpdateRateMsMinimum returns the UpdateRateMsMinimum field if non-nil, zero value otherwise.

### GetUpdateRateMsMinimumOk

`func (o *APIConfig) GetUpdateRateMsMinimumOk() (*int32, bool)`

GetUpdateRateMsMinimumOk returns a tuple with the UpdateRateMsMinimum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateRateMsMinimum

`func (o *APIConfig) SetUpdateRateMsMinimum(v int32)`

SetUpdateRateMsMinimum sets UpdateRateMsMinimum field to given value.


### GetUpdateRateMsNormal

`func (o *APIConfig) GetUpdateRateMsNormal() int32`

GetUpdateRateMsNormal returns the UpdateRateMsNormal field if non-nil, zero value otherwise.

### GetUpdateRateMsNormalOk

`func (o *APIConfig) GetUpdateRateMsNormalOk() (*int32, bool)`

GetUpdateRateMsNormalOk returns a tuple with the UpdateRateMsNormal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateRateMsNormal

`func (o *APIConfig) SetUpdateRateMsNormal(v int32)`

SetUpdateRateMsNormal sets UpdateRateMsNormal field to given value.


### GetUpdateRateMsUdonManual

`func (o *APIConfig) GetUpdateRateMsUdonManual() int32`

GetUpdateRateMsUdonManual returns the UpdateRateMsUdonManual field if non-nil, zero value otherwise.

### GetUpdateRateMsUdonManualOk

`func (o *APIConfig) GetUpdateRateMsUdonManualOk() (*int32, bool)`

GetUpdateRateMsUdonManualOk returns a tuple with the UpdateRateMsUdonManual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateRateMsUdonManual

`func (o *APIConfig) SetUpdateRateMsUdonManual(v int32)`

SetUpdateRateMsUdonManual sets UpdateRateMsUdonManual field to given value.


### GetUploadAnalysisPercent

`func (o *APIConfig) GetUploadAnalysisPercent() int32`

GetUploadAnalysisPercent returns the UploadAnalysisPercent field if non-nil, zero value otherwise.

### GetUploadAnalysisPercentOk

`func (o *APIConfig) GetUploadAnalysisPercentOk() (*int32, bool)`

GetUploadAnalysisPercentOk returns a tuple with the UploadAnalysisPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadAnalysisPercent

`func (o *APIConfig) SetUploadAnalysisPercent(v int32)`

SetUploadAnalysisPercent sets UploadAnalysisPercent field to given value.


### GetUrlList

`func (o *APIConfig) GetUrlList() []string`

GetUrlList returns the UrlList field if non-nil, zero value otherwise.

### GetUrlListOk

`func (o *APIConfig) GetUrlListOk() (*[]string, bool)`

GetUrlListOk returns a tuple with the UrlList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlList

`func (o *APIConfig) SetUrlList(v []string)`

SetUrlList sets UrlList field to given value.


### GetUseReliableUdpForVoice

`func (o *APIConfig) GetUseReliableUdpForVoice() bool`

GetUseReliableUdpForVoice returns the UseReliableUdpForVoice field if non-nil, zero value otherwise.

### GetUseReliableUdpForVoiceOk

`func (o *APIConfig) GetUseReliableUdpForVoiceOk() (*bool, bool)`

GetUseReliableUdpForVoiceOk returns a tuple with the UseReliableUdpForVoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseReliableUdpForVoice

`func (o *APIConfig) SetUseReliableUdpForVoice(v bool)`

SetUseReliableUdpForVoice sets UseReliableUdpForVoice field to given value.


### GetUserUpdatePeriod

`func (o *APIConfig) GetUserUpdatePeriod() int32`

GetUserUpdatePeriod returns the UserUpdatePeriod field if non-nil, zero value otherwise.

### GetUserUpdatePeriodOk

`func (o *APIConfig) GetUserUpdatePeriodOk() (*int32, bool)`

GetUserUpdatePeriodOk returns a tuple with the UserUpdatePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUpdatePeriod

`func (o *APIConfig) SetUserUpdatePeriod(v int32)`

SetUserUpdatePeriod sets UserUpdatePeriod field to given value.


### GetUserVerificationDelay

`func (o *APIConfig) GetUserVerificationDelay() int32`

GetUserVerificationDelay returns the UserVerificationDelay field if non-nil, zero value otherwise.

### GetUserVerificationDelayOk

`func (o *APIConfig) GetUserVerificationDelayOk() (*int32, bool)`

GetUserVerificationDelayOk returns a tuple with the UserVerificationDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVerificationDelay

`func (o *APIConfig) SetUserVerificationDelay(v int32)`

SetUserVerificationDelay sets UserVerificationDelay field to given value.


### GetUserVerificationRetry

`func (o *APIConfig) GetUserVerificationRetry() int32`

GetUserVerificationRetry returns the UserVerificationRetry field if non-nil, zero value otherwise.

### GetUserVerificationRetryOk

`func (o *APIConfig) GetUserVerificationRetryOk() (*int32, bool)`

GetUserVerificationRetryOk returns a tuple with the UserVerificationRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVerificationRetry

`func (o *APIConfig) SetUserVerificationRetry(v int32)`

SetUserVerificationRetry sets UserVerificationRetry field to given value.


### GetUserVerificationTimeout

`func (o *APIConfig) GetUserVerificationTimeout() int32`

GetUserVerificationTimeout returns the UserVerificationTimeout field if non-nil, zero value otherwise.

### GetUserVerificationTimeoutOk

`func (o *APIConfig) GetUserVerificationTimeoutOk() (*int32, bool)`

GetUserVerificationTimeoutOk returns a tuple with the UserVerificationTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVerificationTimeout

`func (o *APIConfig) SetUserVerificationTimeout(v int32)`

SetUserVerificationTimeout sets UserVerificationTimeout field to given value.


### GetViveWindowsUrl

`func (o *APIConfig) GetViveWindowsUrl() string`

GetViveWindowsUrl returns the ViveWindowsUrl field if non-nil, zero value otherwise.

### GetViveWindowsUrlOk

`func (o *APIConfig) GetViveWindowsUrlOk() (*string, bool)`

GetViveWindowsUrlOk returns a tuple with the ViveWindowsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViveWindowsUrl

`func (o *APIConfig) SetViveWindowsUrl(v string)`

SetViveWindowsUrl sets ViveWindowsUrl field to given value.


### GetWhiteListedAssetUrls

`func (o *APIConfig) GetWhiteListedAssetUrls() []string`

GetWhiteListedAssetUrls returns the WhiteListedAssetUrls field if non-nil, zero value otherwise.

### GetWhiteListedAssetUrlsOk

`func (o *APIConfig) GetWhiteListedAssetUrlsOk() (*[]string, bool)`

GetWhiteListedAssetUrlsOk returns a tuple with the WhiteListedAssetUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhiteListedAssetUrls

`func (o *APIConfig) SetWhiteListedAssetUrls(v []string)`

SetWhiteListedAssetUrls sets WhiteListedAssetUrls field to given value.


### GetWorldUpdatePeriod

`func (o *APIConfig) GetWorldUpdatePeriod() int32`

GetWorldUpdatePeriod returns the WorldUpdatePeriod field if non-nil, zero value otherwise.

### GetWorldUpdatePeriodOk

`func (o *APIConfig) GetWorldUpdatePeriodOk() (*int32, bool)`

GetWorldUpdatePeriodOk returns a tuple with the WorldUpdatePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorldUpdatePeriod

`func (o *APIConfig) SetWorldUpdatePeriod(v int32)`

SetWorldUpdatePeriod sets WorldUpdatePeriod field to given value.


### GetPlayerUrlResolverHash

`func (o *APIConfig) GetPlayerUrlResolverHash() string`

GetPlayerUrlResolverHash returns the PlayerUrlResolverHash field if non-nil, zero value otherwise.

### GetPlayerUrlResolverHashOk

`func (o *APIConfig) GetPlayerUrlResolverHashOk() (*string, bool)`

GetPlayerUrlResolverHashOk returns a tuple with the PlayerUrlResolverHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerUrlResolverHash

`func (o *APIConfig) SetPlayerUrlResolverHash(v string)`

SetPlayerUrlResolverHash sets PlayerUrlResolverHash field to given value.


### GetPlayerUrlResolverVersion

`func (o *APIConfig) GetPlayerUrlResolverVersion() string`

GetPlayerUrlResolverVersion returns the PlayerUrlResolverVersion field if non-nil, zero value otherwise.

### GetPlayerUrlResolverVersionOk

`func (o *APIConfig) GetPlayerUrlResolverVersionOk() (*string, bool)`

GetPlayerUrlResolverVersionOk returns a tuple with the PlayerUrlResolverVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerUrlResolverVersion

`func (o *APIConfig) SetPlayerUrlResolverVersion(v string)`

SetPlayerUrlResolverVersion sets PlayerUrlResolverVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


