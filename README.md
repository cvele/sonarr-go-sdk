# Go API client for openapi

Sonarr API docs

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 3.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/cvele/sonarr-go-sdk"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:8989*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuthenticationApi* | [**LoginPost**](docs/AuthenticationApi.md#loginpost) | **Post** /login | 
*AuthenticationApi* | [**LogoutGet**](docs/AuthenticationApi.md#logoutget) | **Get** /logout | 
*BackupApi* | [**ApiV3SystemBackupGet**](docs/BackupApi.md#apiv3systembackupget) | **Get** /api/v3/system/backup | 
*BackupApi* | [**ApiV3SystemBackupIdDelete**](docs/BackupApi.md#apiv3systembackupiddelete) | **Delete** /api/v3/system/backup/{id} | 
*BackupApi* | [**ApiV3SystemBackupRestoreIdPost**](docs/BackupApi.md#apiv3systembackuprestoreidpost) | **Post** /api/v3/system/backup/restore/{id} | 
*BackupApi* | [**ApiV3SystemBackupRestoreUploadPost**](docs/BackupApi.md#apiv3systembackuprestoreuploadpost) | **Post** /api/v3/system/backup/restore/upload | 
*BlocklistApi* | [**ApiV3BlocklistBulkDelete**](docs/BlocklistApi.md#apiv3blocklistbulkdelete) | **Delete** /api/v3/blocklist/bulk | 
*BlocklistApi* | [**ApiV3BlocklistGet**](docs/BlocklistApi.md#apiv3blocklistget) | **Get** /api/v3/blocklist | 
*BlocklistApi* | [**ApiV3BlocklistIdDelete**](docs/BlocklistApi.md#apiv3blocklistiddelete) | **Delete** /api/v3/blocklist/{id} | 
*CalendarApi* | [**ApiV3CalendarGet**](docs/CalendarApi.md#apiv3calendarget) | **Get** /api/v3/calendar | 
*CalendarApi* | [**ApiV3CalendarIdGet**](docs/CalendarApi.md#apiv3calendaridget) | **Get** /api/v3/calendar/{id} | 
*CalendarFeedApi* | [**FeedV3CalendarSonarrIcsGet**](docs/CalendarFeedApi.md#feedv3calendarsonarricsget) | **Get** /feed/v3/calendar/sonarr.ics | 
*CommandApi* | [**ApiV3CommandGet**](docs/CommandApi.md#apiv3commandget) | **Get** /api/v3/command | 
*CommandApi* | [**ApiV3CommandIdDelete**](docs/CommandApi.md#apiv3commandiddelete) | **Delete** /api/v3/command/{id} | 
*CommandApi* | [**ApiV3CommandIdGet**](docs/CommandApi.md#apiv3commandidget) | **Get** /api/v3/command/{id} | 
*CommandApi* | [**ApiV3CommandPost**](docs/CommandApi.md#apiv3commandpost) | **Post** /api/v3/command | 
*CustomFilterApi* | [**ApiV3CustomfilterGet**](docs/CustomFilterApi.md#apiv3customfilterget) | **Get** /api/v3/customfilter | 
*CustomFilterApi* | [**ApiV3CustomfilterIdDelete**](docs/CustomFilterApi.md#apiv3customfilteriddelete) | **Delete** /api/v3/customfilter/{id} | 
*CustomFilterApi* | [**ApiV3CustomfilterIdGet**](docs/CustomFilterApi.md#apiv3customfilteridget) | **Get** /api/v3/customfilter/{id} | 
*CustomFilterApi* | [**ApiV3CustomfilterIdPut**](docs/CustomFilterApi.md#apiv3customfilteridput) | **Put** /api/v3/customfilter/{id} | 
*CustomFilterApi* | [**ApiV3CustomfilterPost**](docs/CustomFilterApi.md#apiv3customfilterpost) | **Post** /api/v3/customfilter | 
*CustomFormatApi* | [**ApiV3CustomformatGet**](docs/CustomFormatApi.md#apiv3customformatget) | **Get** /api/v3/customformat | 
*CustomFormatApi* | [**ApiV3CustomformatIdDelete**](docs/CustomFormatApi.md#apiv3customformatiddelete) | **Delete** /api/v3/customformat/{id} | 
*CustomFormatApi* | [**ApiV3CustomformatIdGet**](docs/CustomFormatApi.md#apiv3customformatidget) | **Get** /api/v3/customformat/{id} | 
*CustomFormatApi* | [**ApiV3CustomformatIdPut**](docs/CustomFormatApi.md#apiv3customformatidput) | **Put** /api/v3/customformat/{id} | 
*CustomFormatApi* | [**ApiV3CustomformatPost**](docs/CustomFormatApi.md#apiv3customformatpost) | **Post** /api/v3/customformat | 
*CustomFormatApi* | [**ApiV3CustomformatSchemaGet**](docs/CustomFormatApi.md#apiv3customformatschemaget) | **Get** /api/v3/customformat/schema | 
*CutoffApi* | [**ApiV3WantedCutoffGet**](docs/CutoffApi.md#apiv3wantedcutoffget) | **Get** /api/v3/wanted/cutoff | 
*CutoffApi* | [**ApiV3WantedCutoffIdGet**](docs/CutoffApi.md#apiv3wantedcutoffidget) | **Get** /api/v3/wanted/cutoff/{id} | 
*DelayProfileApi* | [**ApiV3DelayprofileGet**](docs/DelayProfileApi.md#apiv3delayprofileget) | **Get** /api/v3/delayprofile | 
*DelayProfileApi* | [**ApiV3DelayprofileIdDelete**](docs/DelayProfileApi.md#apiv3delayprofileiddelete) | **Delete** /api/v3/delayprofile/{id} | 
*DelayProfileApi* | [**ApiV3DelayprofileIdGet**](docs/DelayProfileApi.md#apiv3delayprofileidget) | **Get** /api/v3/delayprofile/{id} | 
*DelayProfileApi* | [**ApiV3DelayprofileIdPut**](docs/DelayProfileApi.md#apiv3delayprofileidput) | **Put** /api/v3/delayprofile/{id} | 
*DelayProfileApi* | [**ApiV3DelayprofilePost**](docs/DelayProfileApi.md#apiv3delayprofilepost) | **Post** /api/v3/delayprofile | 
*DelayProfileApi* | [**ApiV3DelayprofileReorderIdPut**](docs/DelayProfileApi.md#apiv3delayprofilereorderidput) | **Put** /api/v3/delayprofile/reorder/{id} | 
*DiskSpaceApi* | [**ApiV3DiskspaceGet**](docs/DiskSpaceApi.md#apiv3diskspaceget) | **Get** /api/v3/diskspace | 
*DownloadClientApi* | [**ApiV3DownloadclientActionNamePost**](docs/DownloadClientApi.md#apiv3downloadclientactionnamepost) | **Post** /api/v3/downloadclient/action/{name} | 
*DownloadClientApi* | [**ApiV3DownloadclientGet**](docs/DownloadClientApi.md#apiv3downloadclientget) | **Get** /api/v3/downloadclient | 
*DownloadClientApi* | [**ApiV3DownloadclientIdDelete**](docs/DownloadClientApi.md#apiv3downloadclientiddelete) | **Delete** /api/v3/downloadclient/{id} | 
*DownloadClientApi* | [**ApiV3DownloadclientIdGet**](docs/DownloadClientApi.md#apiv3downloadclientidget) | **Get** /api/v3/downloadclient/{id} | 
*DownloadClientApi* | [**ApiV3DownloadclientIdPut**](docs/DownloadClientApi.md#apiv3downloadclientidput) | **Put** /api/v3/downloadclient/{id} | 
*DownloadClientApi* | [**ApiV3DownloadclientPost**](docs/DownloadClientApi.md#apiv3downloadclientpost) | **Post** /api/v3/downloadclient | 
*DownloadClientApi* | [**ApiV3DownloadclientSchemaGet**](docs/DownloadClientApi.md#apiv3downloadclientschemaget) | **Get** /api/v3/downloadclient/schema | 
*DownloadClientApi* | [**ApiV3DownloadclientTestPost**](docs/DownloadClientApi.md#apiv3downloadclienttestpost) | **Post** /api/v3/downloadclient/test | 
*DownloadClientApi* | [**ApiV3DownloadclientTestallPost**](docs/DownloadClientApi.md#apiv3downloadclienttestallpost) | **Post** /api/v3/downloadclient/testall | 
*DownloadClientConfigApi* | [**ApiV3ConfigDownloadclientGet**](docs/DownloadClientConfigApi.md#apiv3configdownloadclientget) | **Get** /api/v3/config/downloadclient | 
*DownloadClientConfigApi* | [**ApiV3ConfigDownloadclientIdGet**](docs/DownloadClientConfigApi.md#apiv3configdownloadclientidget) | **Get** /api/v3/config/downloadclient/{id} | 
*DownloadClientConfigApi* | [**ApiV3ConfigDownloadclientIdPut**](docs/DownloadClientConfigApi.md#apiv3configdownloadclientidput) | **Put** /api/v3/config/downloadclient/{id} | 
*EpisodeApi* | [**ApiV3EpisodeGet**](docs/EpisodeApi.md#apiv3episodeget) | **Get** /api/v3/episode | 
*EpisodeApi* | [**ApiV3EpisodeIdGet**](docs/EpisodeApi.md#apiv3episodeidget) | **Get** /api/v3/episode/{id} | 
*EpisodeApi* | [**ApiV3EpisodeIdPut**](docs/EpisodeApi.md#apiv3episodeidput) | **Put** /api/v3/episode/{id} | 
*EpisodeApi* | [**ApiV3EpisodeMonitorPut**](docs/EpisodeApi.md#apiv3episodemonitorput) | **Put** /api/v3/episode/monitor | 
*EpisodeFileApi* | [**ApiV3EpisodefileBulkDelete**](docs/EpisodeFileApi.md#apiv3episodefilebulkdelete) | **Delete** /api/v3/episodefile/bulk | 
*EpisodeFileApi* | [**ApiV3EpisodefileBulkPut**](docs/EpisodeFileApi.md#apiv3episodefilebulkput) | **Put** /api/v3/episodefile/bulk | 
*EpisodeFileApi* | [**ApiV3EpisodefileEditorPut**](docs/EpisodeFileApi.md#apiv3episodefileeditorput) | **Put** /api/v3/episodefile/editor | 
*EpisodeFileApi* | [**ApiV3EpisodefileGet**](docs/EpisodeFileApi.md#apiv3episodefileget) | **Get** /api/v3/episodefile | 
*EpisodeFileApi* | [**ApiV3EpisodefileIdDelete**](docs/EpisodeFileApi.md#apiv3episodefileiddelete) | **Delete** /api/v3/episodefile/{id} | 
*EpisodeFileApi* | [**ApiV3EpisodefileIdGet**](docs/EpisodeFileApi.md#apiv3episodefileidget) | **Get** /api/v3/episodefile/{id} | 
*EpisodeFileApi* | [**ApiV3EpisodefileIdPut**](docs/EpisodeFileApi.md#apiv3episodefileidput) | **Put** /api/v3/episodefile/{id} | 
*FileSystemApi* | [**ApiV3FilesystemGet**](docs/FileSystemApi.md#apiv3filesystemget) | **Get** /api/v3/filesystem | 
*FileSystemApi* | [**ApiV3FilesystemMediafilesGet**](docs/FileSystemApi.md#apiv3filesystemmediafilesget) | **Get** /api/v3/filesystem/mediafiles | 
*FileSystemApi* | [**ApiV3FilesystemTypeGet**](docs/FileSystemApi.md#apiv3filesystemtypeget) | **Get** /api/v3/filesystem/type | 
*HealthApi* | [**ApiV3HealthGet**](docs/HealthApi.md#apiv3healthget) | **Get** /api/v3/health | 
*HealthApi* | [**ApiV3HealthIdGet**](docs/HealthApi.md#apiv3healthidget) | **Get** /api/v3/health/{id} | 
*HistoryApi* | [**ApiV3HistoryFailedIdPost**](docs/HistoryApi.md#apiv3historyfailedidpost) | **Post** /api/v3/history/failed/{id} | 
*HistoryApi* | [**ApiV3HistoryGet**](docs/HistoryApi.md#apiv3historyget) | **Get** /api/v3/history | 
*HistoryApi* | [**ApiV3HistorySeriesGet**](docs/HistoryApi.md#apiv3historyseriesget) | **Get** /api/v3/history/series | 
*HistoryApi* | [**ApiV3HistorySinceGet**](docs/HistoryApi.md#apiv3historysinceget) | **Get** /api/v3/history/since | 
*HostConfigApi* | [**ApiV3ConfigHostGet**](docs/HostConfigApi.md#apiv3confighostget) | **Get** /api/v3/config/host | 
*HostConfigApi* | [**ApiV3ConfigHostIdGet**](docs/HostConfigApi.md#apiv3confighostidget) | **Get** /api/v3/config/host/{id} | 
*HostConfigApi* | [**ApiV3ConfigHostIdPut**](docs/HostConfigApi.md#apiv3confighostidput) | **Put** /api/v3/config/host/{id} | 
*ImportListApi* | [**ApiV3ImportlistActionNamePost**](docs/ImportListApi.md#apiv3importlistactionnamepost) | **Post** /api/v3/importlist/action/{name} | 
*ImportListApi* | [**ApiV3ImportlistGet**](docs/ImportListApi.md#apiv3importlistget) | **Get** /api/v3/importlist | 
*ImportListApi* | [**ApiV3ImportlistIdDelete**](docs/ImportListApi.md#apiv3importlistiddelete) | **Delete** /api/v3/importlist/{id} | 
*ImportListApi* | [**ApiV3ImportlistIdGet**](docs/ImportListApi.md#apiv3importlistidget) | **Get** /api/v3/importlist/{id} | 
*ImportListApi* | [**ApiV3ImportlistIdPut**](docs/ImportListApi.md#apiv3importlistidput) | **Put** /api/v3/importlist/{id} | 
*ImportListApi* | [**ApiV3ImportlistPost**](docs/ImportListApi.md#apiv3importlistpost) | **Post** /api/v3/importlist | 
*ImportListApi* | [**ApiV3ImportlistSchemaGet**](docs/ImportListApi.md#apiv3importlistschemaget) | **Get** /api/v3/importlist/schema | 
*ImportListApi* | [**ApiV3ImportlistTestPost**](docs/ImportListApi.md#apiv3importlisttestpost) | **Post** /api/v3/importlist/test | 
*ImportListApi* | [**ApiV3ImportlistTestallPost**](docs/ImportListApi.md#apiv3importlisttestallpost) | **Post** /api/v3/importlist/testall | 
*ImportListExclusionApi* | [**ApiV3ImportlistexclusionGet**](docs/ImportListExclusionApi.md#apiv3importlistexclusionget) | **Get** /api/v3/importlistexclusion | 
*ImportListExclusionApi* | [**ApiV3ImportlistexclusionIdDelete**](docs/ImportListExclusionApi.md#apiv3importlistexclusioniddelete) | **Delete** /api/v3/importlistexclusion/{id} | 
*ImportListExclusionApi* | [**ApiV3ImportlistexclusionIdGet**](docs/ImportListExclusionApi.md#apiv3importlistexclusionidget) | **Get** /api/v3/importlistexclusion/{id} | 
*ImportListExclusionApi* | [**ApiV3ImportlistexclusionIdPut**](docs/ImportListExclusionApi.md#apiv3importlistexclusionidput) | **Put** /api/v3/importlistexclusion/{id} | 
*ImportListExclusionApi* | [**ApiV3ImportlistexclusionPost**](docs/ImportListExclusionApi.md#apiv3importlistexclusionpost) | **Post** /api/v3/importlistexclusion | 
*IndexerApi* | [**ApiV3IndexerActionNamePost**](docs/IndexerApi.md#apiv3indexeractionnamepost) | **Post** /api/v3/indexer/action/{name} | 
*IndexerApi* | [**ApiV3IndexerGet**](docs/IndexerApi.md#apiv3indexerget) | **Get** /api/v3/indexer | 
*IndexerApi* | [**ApiV3IndexerIdDelete**](docs/IndexerApi.md#apiv3indexeriddelete) | **Delete** /api/v3/indexer/{id} | 
*IndexerApi* | [**ApiV3IndexerIdGet**](docs/IndexerApi.md#apiv3indexeridget) | **Get** /api/v3/indexer/{id} | 
*IndexerApi* | [**ApiV3IndexerIdPut**](docs/IndexerApi.md#apiv3indexeridput) | **Put** /api/v3/indexer/{id} | 
*IndexerApi* | [**ApiV3IndexerPost**](docs/IndexerApi.md#apiv3indexerpost) | **Post** /api/v3/indexer | 
*IndexerApi* | [**ApiV3IndexerSchemaGet**](docs/IndexerApi.md#apiv3indexerschemaget) | **Get** /api/v3/indexer/schema | 
*IndexerApi* | [**ApiV3IndexerTestPost**](docs/IndexerApi.md#apiv3indexertestpost) | **Post** /api/v3/indexer/test | 
*IndexerApi* | [**ApiV3IndexerTestallPost**](docs/IndexerApi.md#apiv3indexertestallpost) | **Post** /api/v3/indexer/testall | 
*IndexerConfigApi* | [**ApiV3ConfigIndexerGet**](docs/IndexerConfigApi.md#apiv3configindexerget) | **Get** /api/v3/config/indexer | 
*IndexerConfigApi* | [**ApiV3ConfigIndexerIdGet**](docs/IndexerConfigApi.md#apiv3configindexeridget) | **Get** /api/v3/config/indexer/{id} | 
*IndexerConfigApi* | [**ApiV3ConfigIndexerIdPut**](docs/IndexerConfigApi.md#apiv3configindexeridput) | **Put** /api/v3/config/indexer/{id} | 
*InitializeJsApi* | [**InitializeJsGet**](docs/InitializeJsApi.md#initializejsget) | **Get** /initialize.js | 
*LanguageProfileApi* | [**ApiV3LanguageprofileGet**](docs/LanguageProfileApi.md#apiv3languageprofileget) | **Get** /api/v3/languageprofile | 
*LanguageProfileApi* | [**ApiV3LanguageprofileIdDelete**](docs/LanguageProfileApi.md#apiv3languageprofileiddelete) | **Delete** /api/v3/languageprofile/{id} | 
*LanguageProfileApi* | [**ApiV3LanguageprofileIdGet**](docs/LanguageProfileApi.md#apiv3languageprofileidget) | **Get** /api/v3/languageprofile/{id} | 
*LanguageProfileApi* | [**ApiV3LanguageprofileIdPut**](docs/LanguageProfileApi.md#apiv3languageprofileidput) | **Put** /api/v3/languageprofile/{id} | 
*LanguageProfileApi* | [**ApiV3LanguageprofilePost**](docs/LanguageProfileApi.md#apiv3languageprofilepost) | **Post** /api/v3/languageprofile | 
*LanguageProfileSchemaApi* | [**ApiV3LanguageprofileSchemaGet**](docs/LanguageProfileSchemaApi.md#apiv3languageprofileschemaget) | **Get** /api/v3/languageprofile/schema | 
*LogApi* | [**ApiV3LogGet**](docs/LogApi.md#apiv3logget) | **Get** /api/v3/log | 
*LogFileApi* | [**ApiV3LogFileFilenameGet**](docs/LogFileApi.md#apiv3logfilefilenameget) | **Get** /api/v3/log/file/{filename} | 
*LogFileApi* | [**ApiV3LogFileGet**](docs/LogFileApi.md#apiv3logfileget) | **Get** /api/v3/log/file | 
*ManualImportApi* | [**ApiV3ManualimportGet**](docs/ManualImportApi.md#apiv3manualimportget) | **Get** /api/v3/manualimport | 
*ManualImportApi* | [**ApiV3ManualimportPost**](docs/ManualImportApi.md#apiv3manualimportpost) | **Post** /api/v3/manualimport | 
*MediaCoverApi* | [**ApiV3MediacoverSeriesIdFilenameGet**](docs/MediaCoverApi.md#apiv3mediacoverseriesidfilenameget) | **Get** /api/v3/mediacover/{seriesId}/{filename} | 
*MediaManagementConfigApi* | [**ApiV3ConfigMediamanagementGet**](docs/MediaManagementConfigApi.md#apiv3configmediamanagementget) | **Get** /api/v3/config/mediamanagement | 
*MediaManagementConfigApi* | [**ApiV3ConfigMediamanagementIdGet**](docs/MediaManagementConfigApi.md#apiv3configmediamanagementidget) | **Get** /api/v3/config/mediamanagement/{id} | 
*MediaManagementConfigApi* | [**ApiV3ConfigMediamanagementIdPut**](docs/MediaManagementConfigApi.md#apiv3configmediamanagementidput) | **Put** /api/v3/config/mediamanagement/{id} | 
*MetadataApi* | [**ApiV3MetadataActionNamePost**](docs/MetadataApi.md#apiv3metadataactionnamepost) | **Post** /api/v3/metadata/action/{name} | 
*MetadataApi* | [**ApiV3MetadataGet**](docs/MetadataApi.md#apiv3metadataget) | **Get** /api/v3/metadata | 
*MetadataApi* | [**ApiV3MetadataIdDelete**](docs/MetadataApi.md#apiv3metadataiddelete) | **Delete** /api/v3/metadata/{id} | 
*MetadataApi* | [**ApiV3MetadataIdGet**](docs/MetadataApi.md#apiv3metadataidget) | **Get** /api/v3/metadata/{id} | 
*MetadataApi* | [**ApiV3MetadataIdPut**](docs/MetadataApi.md#apiv3metadataidput) | **Put** /api/v3/metadata/{id} | 
*MetadataApi* | [**ApiV3MetadataPost**](docs/MetadataApi.md#apiv3metadatapost) | **Post** /api/v3/metadata | 
*MetadataApi* | [**ApiV3MetadataSchemaGet**](docs/MetadataApi.md#apiv3metadataschemaget) | **Get** /api/v3/metadata/schema | 
*MetadataApi* | [**ApiV3MetadataTestPost**](docs/MetadataApi.md#apiv3metadatatestpost) | **Post** /api/v3/metadata/test | 
*MetadataApi* | [**ApiV3MetadataTestallPost**](docs/MetadataApi.md#apiv3metadatatestallpost) | **Post** /api/v3/metadata/testall | 
*MissingApi* | [**ApiV3WantedMissingGet**](docs/MissingApi.md#apiv3wantedmissingget) | **Get** /api/v3/wanted/missing | 
*MissingApi* | [**ApiV3WantedMissingIdGet**](docs/MissingApi.md#apiv3wantedmissingidget) | **Get** /api/v3/wanted/missing/{id} | 
*NamingConfigApi* | [**ApiV3ConfigNamingExamplesGet**](docs/NamingConfigApi.md#apiv3confignamingexamplesget) | **Get** /api/v3/config/naming/examples | 
*NamingConfigApi* | [**ApiV3ConfigNamingGet**](docs/NamingConfigApi.md#apiv3confignamingget) | **Get** /api/v3/config/naming | 
*NamingConfigApi* | [**ApiV3ConfigNamingIdGet**](docs/NamingConfigApi.md#apiv3confignamingidget) | **Get** /api/v3/config/naming/{id} | 
*NamingConfigApi* | [**ApiV3ConfigNamingIdPut**](docs/NamingConfigApi.md#apiv3confignamingidput) | **Put** /api/v3/config/naming/{id} | 
*NotificationApi* | [**ApiV3NotificationActionNamePost**](docs/NotificationApi.md#apiv3notificationactionnamepost) | **Post** /api/v3/notification/action/{name} | 
*NotificationApi* | [**ApiV3NotificationGet**](docs/NotificationApi.md#apiv3notificationget) | **Get** /api/v3/notification | 
*NotificationApi* | [**ApiV3NotificationIdDelete**](docs/NotificationApi.md#apiv3notificationiddelete) | **Delete** /api/v3/notification/{id} | 
*NotificationApi* | [**ApiV3NotificationIdGet**](docs/NotificationApi.md#apiv3notificationidget) | **Get** /api/v3/notification/{id} | 
*NotificationApi* | [**ApiV3NotificationIdPut**](docs/NotificationApi.md#apiv3notificationidput) | **Put** /api/v3/notification/{id} | 
*NotificationApi* | [**ApiV3NotificationPost**](docs/NotificationApi.md#apiv3notificationpost) | **Post** /api/v3/notification | 
*NotificationApi* | [**ApiV3NotificationSchemaGet**](docs/NotificationApi.md#apiv3notificationschemaget) | **Get** /api/v3/notification/schema | 
*NotificationApi* | [**ApiV3NotificationTestPost**](docs/NotificationApi.md#apiv3notificationtestpost) | **Post** /api/v3/notification/test | 
*NotificationApi* | [**ApiV3NotificationTestallPost**](docs/NotificationApi.md#apiv3notificationtestallpost) | **Post** /api/v3/notification/testall | 
*ParseApi* | [**ApiV3ParseGet**](docs/ParseApi.md#apiv3parseget) | **Get** /api/v3/parse | 
*PingApi* | [**PingGet**](docs/PingApi.md#pingget) | **Get** /ping | 
*QualityDefinitionApi* | [**ApiV3QualitydefinitionGet**](docs/QualityDefinitionApi.md#apiv3qualitydefinitionget) | **Get** /api/v3/qualitydefinition | 
*QualityDefinitionApi* | [**ApiV3QualitydefinitionIdGet**](docs/QualityDefinitionApi.md#apiv3qualitydefinitionidget) | **Get** /api/v3/qualitydefinition/{id} | 
*QualityDefinitionApi* | [**ApiV3QualitydefinitionIdPut**](docs/QualityDefinitionApi.md#apiv3qualitydefinitionidput) | **Put** /api/v3/qualitydefinition/{id} | 
*QualityDefinitionApi* | [**ApiV3QualitydefinitionUpdatePut**](docs/QualityDefinitionApi.md#apiv3qualitydefinitionupdateput) | **Put** /api/v3/qualitydefinition/update | 
*QualityProfileApi* | [**ApiV3QualityprofileGet**](docs/QualityProfileApi.md#apiv3qualityprofileget) | **Get** /api/v3/qualityprofile | 
*QualityProfileApi* | [**ApiV3QualityprofileIdDelete**](docs/QualityProfileApi.md#apiv3qualityprofileiddelete) | **Delete** /api/v3/qualityprofile/{id} | 
*QualityProfileApi* | [**ApiV3QualityprofileIdGet**](docs/QualityProfileApi.md#apiv3qualityprofileidget) | **Get** /api/v3/qualityprofile/{id} | 
*QualityProfileApi* | [**ApiV3QualityprofileIdPut**](docs/QualityProfileApi.md#apiv3qualityprofileidput) | **Put** /api/v3/qualityprofile/{id} | 
*QualityProfileApi* | [**ApiV3QualityprofilePost**](docs/QualityProfileApi.md#apiv3qualityprofilepost) | **Post** /api/v3/qualityprofile | 
*QualityProfileSchemaApi* | [**ApiV3QualityprofileSchemaGet**](docs/QualityProfileSchemaApi.md#apiv3qualityprofileschemaget) | **Get** /api/v3/qualityprofile/schema | 
*QueueApi* | [**ApiV3QueueBulkDelete**](docs/QueueApi.md#apiv3queuebulkdelete) | **Delete** /api/v3/queue/bulk | 
*QueueApi* | [**ApiV3QueueGet**](docs/QueueApi.md#apiv3queueget) | **Get** /api/v3/queue | 
*QueueApi* | [**ApiV3QueueIdDelete**](docs/QueueApi.md#apiv3queueiddelete) | **Delete** /api/v3/queue/{id} | 
*QueueApi* | [**ApiV3QueueIdGet**](docs/QueueApi.md#apiv3queueidget) | **Get** /api/v3/queue/{id} | 
*QueueActionApi* | [**ApiV3QueueGrabBulkPost**](docs/QueueActionApi.md#apiv3queuegrabbulkpost) | **Post** /api/v3/queue/grab/bulk | 
*QueueActionApi* | [**ApiV3QueueGrabIdPost**](docs/QueueActionApi.md#apiv3queuegrabidpost) | **Post** /api/v3/queue/grab/{id} | 
*QueueDetailsApi* | [**ApiV3QueueDetailsGet**](docs/QueueDetailsApi.md#apiv3queuedetailsget) | **Get** /api/v3/queue/details | 
*QueueDetailsApi* | [**ApiV3QueueDetailsIdGet**](docs/QueueDetailsApi.md#apiv3queuedetailsidget) | **Get** /api/v3/queue/details/{id} | 
*QueueStatusApi* | [**ApiV3QueueStatusGet**](docs/QueueStatusApi.md#apiv3queuestatusget) | **Get** /api/v3/queue/status | 
*QueueStatusApi* | [**ApiV3QueueStatusIdGet**](docs/QueueStatusApi.md#apiv3queuestatusidget) | **Get** /api/v3/queue/status/{id} | 
*ReleaseApi* | [**ApiV3ReleaseGet**](docs/ReleaseApi.md#apiv3releaseget) | **Get** /api/v3/release | 
*ReleaseApi* | [**ApiV3ReleaseIdGet**](docs/ReleaseApi.md#apiv3releaseidget) | **Get** /api/v3/release/{id} | 
*ReleaseApi* | [**ApiV3ReleasePost**](docs/ReleaseApi.md#apiv3releasepost) | **Post** /api/v3/release | 
*ReleaseProfileApi* | [**ApiV3ReleaseprofileGet**](docs/ReleaseProfileApi.md#apiv3releaseprofileget) | **Get** /api/v3/releaseprofile | 
*ReleaseProfileApi* | [**ApiV3ReleaseprofileIdDelete**](docs/ReleaseProfileApi.md#apiv3releaseprofileiddelete) | **Delete** /api/v3/releaseprofile/{id} | 
*ReleaseProfileApi* | [**ApiV3ReleaseprofileIdGet**](docs/ReleaseProfileApi.md#apiv3releaseprofileidget) | **Get** /api/v3/releaseprofile/{id} | 
*ReleaseProfileApi* | [**ApiV3ReleaseprofileIdPut**](docs/ReleaseProfileApi.md#apiv3releaseprofileidput) | **Put** /api/v3/releaseprofile/{id} | 
*ReleaseProfileApi* | [**ApiV3ReleaseprofilePost**](docs/ReleaseProfileApi.md#apiv3releaseprofilepost) | **Post** /api/v3/releaseprofile | 
*ReleasePushApi* | [**ApiV3ReleasePushIdGet**](docs/ReleasePushApi.md#apiv3releasepushidget) | **Get** /api/v3/release/push/{id} | 
*ReleasePushApi* | [**ApiV3ReleasePushPost**](docs/ReleasePushApi.md#apiv3releasepushpost) | **Post** /api/v3/release/push | 
*RemotePathMappingApi* | [**ApiV3RemotepathmappingGet**](docs/RemotePathMappingApi.md#apiv3remotepathmappingget) | **Get** /api/v3/remotepathmapping | 
*RemotePathMappingApi* | [**ApiV3RemotepathmappingIdDelete**](docs/RemotePathMappingApi.md#apiv3remotepathmappingiddelete) | **Delete** /api/v3/remotepathmapping/{id} | 
*RemotePathMappingApi* | [**ApiV3RemotepathmappingIdGet**](docs/RemotePathMappingApi.md#apiv3remotepathmappingidget) | **Get** /api/v3/remotepathmapping/{id} | 
*RemotePathMappingApi* | [**ApiV3RemotepathmappingIdPut**](docs/RemotePathMappingApi.md#apiv3remotepathmappingidput) | **Put** /api/v3/remotepathmapping/{id} | 
*RemotePathMappingApi* | [**ApiV3RemotepathmappingPost**](docs/RemotePathMappingApi.md#apiv3remotepathmappingpost) | **Post** /api/v3/remotepathmapping | 
*RenameEpisodeApi* | [**ApiV3RenameGet**](docs/RenameEpisodeApi.md#apiv3renameget) | **Get** /api/v3/rename | 
*RootFolderApi* | [**ApiV3RootfolderGet**](docs/RootFolderApi.md#apiv3rootfolderget) | **Get** /api/v3/rootfolder | 
*RootFolderApi* | [**ApiV3RootfolderIdDelete**](docs/RootFolderApi.md#apiv3rootfolderiddelete) | **Delete** /api/v3/rootfolder/{id} | 
*RootFolderApi* | [**ApiV3RootfolderIdGet**](docs/RootFolderApi.md#apiv3rootfolderidget) | **Get** /api/v3/rootfolder/{id} | 
*RootFolderApi* | [**ApiV3RootfolderPost**](docs/RootFolderApi.md#apiv3rootfolderpost) | **Post** /api/v3/rootfolder | 
*SeasonPassApi* | [**ApiV3SeasonpassPost**](docs/SeasonPassApi.md#apiv3seasonpasspost) | **Post** /api/v3/seasonpass | 
*SeriesApi* | [**ApiV3SeriesGet**](docs/SeriesApi.md#apiv3seriesget) | **Get** /api/v3/series | 
*SeriesApi* | [**ApiV3SeriesIdDelete**](docs/SeriesApi.md#apiv3seriesiddelete) | **Delete** /api/v3/series/{id} | 
*SeriesApi* | [**ApiV3SeriesIdGet**](docs/SeriesApi.md#apiv3seriesidget) | **Get** /api/v3/series/{id} | 
*SeriesApi* | [**ApiV3SeriesIdPut**](docs/SeriesApi.md#apiv3seriesidput) | **Put** /api/v3/series/{id} | 
*SeriesApi* | [**ApiV3SeriesPost**](docs/SeriesApi.md#apiv3seriespost) | **Post** /api/v3/series | 
*SeriesEditorApi* | [**ApiV3SeriesEditorDelete**](docs/SeriesEditorApi.md#apiv3serieseditordelete) | **Delete** /api/v3/series/editor | 
*SeriesEditorApi* | [**ApiV3SeriesEditorPut**](docs/SeriesEditorApi.md#apiv3serieseditorput) | **Put** /api/v3/series/editor | 
*SeriesImportApi* | [**ApiV3SeriesImportPost**](docs/SeriesImportApi.md#apiv3seriesimportpost) | **Post** /api/v3/series/import | 
*SeriesLookupApi* | [**ApiV3SeriesLookupGet**](docs/SeriesLookupApi.md#apiv3serieslookupget) | **Get** /api/v3/series/lookup | 
*StaticResourceApi* | [**ContentPathGet**](docs/StaticResourceApi.md#contentpathget) | **Get** /content/{path} | 
*StaticResourceApi* | [**LoginGet**](docs/StaticResourceApi.md#loginget) | **Get** /login | 
*StaticResourceApi* | [**PathGet**](docs/StaticResourceApi.md#pathget) | **Get** /{path} | 
*StaticResourceApi* | [**RootGet**](docs/StaticResourceApi.md#rootget) | **Get** / | 
*SystemApi* | [**ApiV3SystemRestartPost**](docs/SystemApi.md#apiv3systemrestartpost) | **Post** /api/v3/system/restart | 
*SystemApi* | [**ApiV3SystemRoutesDuplicateGet**](docs/SystemApi.md#apiv3systemroutesduplicateget) | **Get** /api/v3/system/routes/duplicate | 
*SystemApi* | [**ApiV3SystemRoutesGet**](docs/SystemApi.md#apiv3systemroutesget) | **Get** /api/v3/system/routes | 
*SystemApi* | [**ApiV3SystemShutdownPost**](docs/SystemApi.md#apiv3systemshutdownpost) | **Post** /api/v3/system/shutdown | 
*SystemApi* | [**ApiV3SystemStatusGet**](docs/SystemApi.md#apiv3systemstatusget) | **Get** /api/v3/system/status | 
*TagApi* | [**ApiV3TagGet**](docs/TagApi.md#apiv3tagget) | **Get** /api/v3/tag | 
*TagApi* | [**ApiV3TagIdDelete**](docs/TagApi.md#apiv3tagiddelete) | **Delete** /api/v3/tag/{id} | 
*TagApi* | [**ApiV3TagIdGet**](docs/TagApi.md#apiv3tagidget) | **Get** /api/v3/tag/{id} | 
*TagApi* | [**ApiV3TagIdPut**](docs/TagApi.md#apiv3tagidput) | **Put** /api/v3/tag/{id} | 
*TagApi* | [**ApiV3TagPost**](docs/TagApi.md#apiv3tagpost) | **Post** /api/v3/tag | 
*TagDetailsApi* | [**ApiV3TagDetailGet**](docs/TagDetailsApi.md#apiv3tagdetailget) | **Get** /api/v3/tag/detail | 
*TagDetailsApi* | [**ApiV3TagDetailIdGet**](docs/TagDetailsApi.md#apiv3tagdetailidget) | **Get** /api/v3/tag/detail/{id} | 
*TaskApi* | [**ApiV3SystemTaskGet**](docs/TaskApi.md#apiv3systemtaskget) | **Get** /api/v3/system/task | 
*TaskApi* | [**ApiV3SystemTaskIdGet**](docs/TaskApi.md#apiv3systemtaskidget) | **Get** /api/v3/system/task/{id} | 
*UiConfigApi* | [**ApiV3ConfigUiGet**](docs/UiConfigApi.md#apiv3configuiget) | **Get** /api/v3/config/ui | 
*UiConfigApi* | [**ApiV3ConfigUiIdGet**](docs/UiConfigApi.md#apiv3configuiidget) | **Get** /api/v3/config/ui/{id} | 
*UiConfigApi* | [**ApiV3ConfigUiIdPut**](docs/UiConfigApi.md#apiv3configuiidput) | **Put** /api/v3/config/ui/{id} | 
*UpdateApi* | [**ApiV3UpdateGet**](docs/UpdateApi.md#apiv3updateget) | **Get** /api/v3/update | 
*UpdateLogFileApi* | [**ApiV3LogFileUpdateFilenameGet**](docs/UpdateLogFileApi.md#apiv3logfileupdatefilenameget) | **Get** /api/v3/log/file/update/{filename} | 
*UpdateLogFileApi* | [**ApiV3LogFileUpdateGet**](docs/UpdateLogFileApi.md#apiv3logfileupdateget) | **Get** /api/v3/log/file/update | 


## Documentation For Models

 - [AddSeriesOptions](docs/AddSeriesOptions.md)
 - [AlternateTitleResource](docs/AlternateTitleResource.md)
 - [ApplyTags](docs/ApplyTags.md)
 - [AuthenticationRequiredType](docs/AuthenticationRequiredType.md)
 - [AuthenticationType](docs/AuthenticationType.md)
 - [BackupResource](docs/BackupResource.md)
 - [BackupType](docs/BackupType.md)
 - [BlocklistBulkResource](docs/BlocklistBulkResource.md)
 - [BlocklistResource](docs/BlocklistResource.md)
 - [BlocklistResourcePagingResource](docs/BlocklistResourcePagingResource.md)
 - [CertificateValidationType](docs/CertificateValidationType.md)
 - [Command](docs/Command.md)
 - [CommandPriority](docs/CommandPriority.md)
 - [CommandResource](docs/CommandResource.md)
 - [CommandStatus](docs/CommandStatus.md)
 - [CommandTrigger](docs/CommandTrigger.md)
 - [CustomFilterResource](docs/CustomFilterResource.md)
 - [CustomFormatResource](docs/CustomFormatResource.md)
 - [CustomFormatSpecificationSchema](docs/CustomFormatSpecificationSchema.md)
 - [DelayProfileResource](docs/DelayProfileResource.md)
 - [DiskSpaceResource](docs/DiskSpaceResource.md)
 - [DownloadClientConfigResource](docs/DownloadClientConfigResource.md)
 - [DownloadClientResource](docs/DownloadClientResource.md)
 - [DownloadProtocol](docs/DownloadProtocol.md)
 - [EpisodeFileListResource](docs/EpisodeFileListResource.md)
 - [EpisodeFileResource](docs/EpisodeFileResource.md)
 - [EpisodeHistoryEventType](docs/EpisodeHistoryEventType.md)
 - [EpisodeResource](docs/EpisodeResource.md)
 - [EpisodeResourcePagingResource](docs/EpisodeResourcePagingResource.md)
 - [EpisodeTitleRequiredType](docs/EpisodeTitleRequiredType.md)
 - [EpisodesMonitoredResource](docs/EpisodesMonitoredResource.md)
 - [Field](docs/Field.md)
 - [FileDateType](docs/FileDateType.md)
 - [HealthCheckResult](docs/HealthCheckResult.md)
 - [HealthResource](docs/HealthResource.md)
 - [HistoryResource](docs/HistoryResource.md)
 - [HistoryResourcePagingResource](docs/HistoryResourcePagingResource.md)
 - [HostConfigResource](docs/HostConfigResource.md)
 - [HttpUri](docs/HttpUri.md)
 - [ImportListExclusionResource](docs/ImportListExclusionResource.md)
 - [ImportListResource](docs/ImportListResource.md)
 - [ImportListType](docs/ImportListType.md)
 - [IndexerConfigResource](docs/IndexerConfigResource.md)
 - [IndexerResource](docs/IndexerResource.md)
 - [Language](docs/Language.md)
 - [LanguageProfileItemResource](docs/LanguageProfileItemResource.md)
 - [LanguageProfileResource](docs/LanguageProfileResource.md)
 - [LogFileResource](docs/LogFileResource.md)
 - [LogResource](docs/LogResource.md)
 - [LogResourcePagingResource](docs/LogResourcePagingResource.md)
 - [ManualImportReprocessResource](docs/ManualImportReprocessResource.md)
 - [ManualImportResource](docs/ManualImportResource.md)
 - [MediaCover](docs/MediaCover.md)
 - [MediaCoverTypes](docs/MediaCoverTypes.md)
 - [MediaInfoResource](docs/MediaInfoResource.md)
 - [MediaManagementConfigResource](docs/MediaManagementConfigResource.md)
 - [MetadataResource](docs/MetadataResource.md)
 - [MonitorTypes](docs/MonitorTypes.md)
 - [MonitoringOptions](docs/MonitoringOptions.md)
 - [NamingConfigResource](docs/NamingConfigResource.md)
 - [NotificationResource](docs/NotificationResource.md)
 - [PagingResourceFilter](docs/PagingResourceFilter.md)
 - [ParseResource](docs/ParseResource.md)
 - [ParsedEpisodeInfo](docs/ParsedEpisodeInfo.md)
 - [PrivacyLevel](docs/PrivacyLevel.md)
 - [ProfileFormatItemResource](docs/ProfileFormatItemResource.md)
 - [ProperDownloadTypes](docs/ProperDownloadTypes.md)
 - [ProviderMessage](docs/ProviderMessage.md)
 - [ProviderMessageType](docs/ProviderMessageType.md)
 - [ProxyType](docs/ProxyType.md)
 - [Quality](docs/Quality.md)
 - [QualityDefinitionResource](docs/QualityDefinitionResource.md)
 - [QualityModel](docs/QualityModel.md)
 - [QualityProfileQualityItemResource](docs/QualityProfileQualityItemResource.md)
 - [QualityProfileResource](docs/QualityProfileResource.md)
 - [QualitySource](docs/QualitySource.md)
 - [QueueBulkResource](docs/QueueBulkResource.md)
 - [QueueResource](docs/QueueResource.md)
 - [QueueResourcePagingResource](docs/QueueResourcePagingResource.md)
 - [QueueStatusResource](docs/QueueStatusResource.md)
 - [Ratings](docs/Ratings.md)
 - [Rejection](docs/Rejection.md)
 - [RejectionType](docs/RejectionType.md)
 - [ReleaseProfileResource](docs/ReleaseProfileResource.md)
 - [ReleaseResource](docs/ReleaseResource.md)
 - [RemotePathMappingResource](docs/RemotePathMappingResource.md)
 - [RenameEpisodeResource](docs/RenameEpisodeResource.md)
 - [RescanAfterRefreshType](docs/RescanAfterRefreshType.md)
 - [Revision](docs/Revision.md)
 - [RootFolderResource](docs/RootFolderResource.md)
 - [SeasonPassResource](docs/SeasonPassResource.md)
 - [SeasonPassSeriesResource](docs/SeasonPassSeriesResource.md)
 - [SeasonResource](docs/SeasonResource.md)
 - [SeasonStatisticsResource](docs/SeasonStatisticsResource.md)
 - [SelectOption](docs/SelectOption.md)
 - [SeriesEditorResource](docs/SeriesEditorResource.md)
 - [SeriesResource](docs/SeriesResource.md)
 - [SeriesStatisticsResource](docs/SeriesStatisticsResource.md)
 - [SeriesStatusType](docs/SeriesStatusType.md)
 - [SeriesTitleInfo](docs/SeriesTitleInfo.md)
 - [SeriesTypes](docs/SeriesTypes.md)
 - [SortDirection](docs/SortDirection.md)
 - [TagDetailsResource](docs/TagDetailsResource.md)
 - [TagResource](docs/TagResource.md)
 - [TaskResource](docs/TaskResource.md)
 - [TimeSpan](docs/TimeSpan.md)
 - [TrackedDownloadState](docs/TrackedDownloadState.md)
 - [TrackedDownloadStatus](docs/TrackedDownloadStatus.md)
 - [TrackedDownloadStatusMessage](docs/TrackedDownloadStatusMessage.md)
 - [UiConfigResource](docs/UiConfigResource.md)
 - [UnmappedFolder](docs/UnmappedFolder.md)
 - [UpdateChanges](docs/UpdateChanges.md)
 - [UpdateMechanism](docs/UpdateMechanism.md)
 - [UpdateResource](docs/UpdateResource.md)
 - [Version](docs/Version.md)


## Documentation For Authorization



### X-Api-Key

- **Type**: API key
- **API key parameter name**: X-Api-Key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-Api-Key and passed in as the auth context for each request.


### apikey

- **Type**: API key
- **API key parameter name**: apikey
- **Location**: URL query string

Note, each API key must be added to a map of `map[string]APIKey` where the key is: apikey and passed in as the auth context for each request.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



