//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package generated

// AccessRight - Access rights of the access policy.
type AccessRight string

const (
	AccessRightDelete AccessRight = "Delete"
	AccessRightRead   AccessRight = "Read"
	AccessRightWrite  AccessRight = "Write"
)

// PossibleAccessRightValues returns the possible values for the AccessRight const type.
func PossibleAccessRightValues() []AccessRight {
	return []AccessRight{
		AccessRightDelete,
		AccessRightRead,
		AccessRightWrite,
	}
}

type CopyStatusType string

const (
	CopyStatusTypeAborted CopyStatusType = "aborted"
	CopyStatusTypeFailed  CopyStatusType = "failed"
	CopyStatusTypePending CopyStatusType = "pending"
	CopyStatusTypeSuccess CopyStatusType = "success"
)

// PossibleCopyStatusTypeValues returns the possible values for the CopyStatusType const type.
func PossibleCopyStatusTypeValues() []CopyStatusType {
	return []CopyStatusType{
		CopyStatusTypeAborted,
		CopyStatusTypeFailed,
		CopyStatusTypePending,
		CopyStatusTypeSuccess,
	}
}

type DeleteSnapshotsOptionType string

const (
	DeleteSnapshotsOptionTypeInclude       DeleteSnapshotsOptionType = "include"
	DeleteSnapshotsOptionTypeIncludeLeased DeleteSnapshotsOptionType = "include-leased"
)

// PossibleDeleteSnapshotsOptionTypeValues returns the possible values for the DeleteSnapshotsOptionType const type.
func PossibleDeleteSnapshotsOptionTypeValues() []DeleteSnapshotsOptionType {
	return []DeleteSnapshotsOptionType{
		DeleteSnapshotsOptionTypeInclude,
		DeleteSnapshotsOptionTypeIncludeLeased,
	}
}

type FileLastWrittenMode string

const (
	FileLastWrittenModeNow      FileLastWrittenMode = "Now"
	FileLastWrittenModePreserve FileLastWrittenMode = "Preserve"
)

// PossibleFileLastWrittenModeValues returns the possible values for the FileLastWrittenMode const type.
func PossibleFileLastWrittenModeValues() []FileLastWrittenMode {
	return []FileLastWrittenMode{
		FileLastWrittenModeNow,
		FileLastWrittenModePreserve,
	}
}

type FileRangeWriteType string

const (
	FileRangeWriteTypeClear  FileRangeWriteType = "clear"
	FileRangeWriteTypeUpdate FileRangeWriteType = "update"
)

// PossibleFileRangeWriteTypeValues returns the possible values for the FileRangeWriteType const type.
func PossibleFileRangeWriteTypeValues() []FileRangeWriteType {
	return []FileRangeWriteType{
		FileRangeWriteTypeClear,
		FileRangeWriteTypeUpdate,
	}
}

// LeaseDurationType - When a share is leased, specifies whether the lease is of infinite or fixed duration.
type LeaseDurationType string

const (
	LeaseDurationTypeFixed    LeaseDurationType = "fixed"
	LeaseDurationTypeInfinite LeaseDurationType = "infinite"
)

// PossibleLeaseDurationTypeValues returns the possible values for the LeaseDurationType const type.
func PossibleLeaseDurationTypeValues() []LeaseDurationType {
	return []LeaseDurationType{
		LeaseDurationTypeFixed,
		LeaseDurationTypeInfinite,
	}
}

// LeaseStateType - Lease state of the share.
type LeaseStateType string

const (
	LeaseStateTypeAvailable LeaseStateType = "available"
	LeaseStateTypeBreaking  LeaseStateType = "breaking"
	LeaseStateTypeBroken    LeaseStateType = "broken"
	LeaseStateTypeExpired   LeaseStateType = "expired"
	LeaseStateTypeLeased    LeaseStateType = "leased"
)

// PossibleLeaseStateTypeValues returns the possible values for the LeaseStateType const type.
func PossibleLeaseStateTypeValues() []LeaseStateType {
	return []LeaseStateType{
		LeaseStateTypeAvailable,
		LeaseStateTypeBreaking,
		LeaseStateTypeBroken,
		LeaseStateTypeExpired,
		LeaseStateTypeLeased,
	}
}

// LeaseStatusType - The current lease status of the share.
type LeaseStatusType string

const (
	LeaseStatusTypeLocked   LeaseStatusType = "locked"
	LeaseStatusTypeUnlocked LeaseStatusType = "unlocked"
)

// PossibleLeaseStatusTypeValues returns the possible values for the LeaseStatusType const type.
func PossibleLeaseStatusTypeValues() []LeaseStatusType {
	return []LeaseStatusType{
		LeaseStatusTypeLocked,
		LeaseStatusTypeUnlocked,
	}
}

type ListFilesIncludeType string

const (
	ListFilesIncludeTypeAttributes    ListFilesIncludeType = "Attributes"
	ListFilesIncludeTypeEtag          ListFilesIncludeType = "Etag"
	ListFilesIncludeTypePermissionKey ListFilesIncludeType = "PermissionKey"
	ListFilesIncludeTypeTimestamps    ListFilesIncludeType = "Timestamps"
)

// PossibleListFilesIncludeTypeValues returns the possible values for the ListFilesIncludeType const type.
func PossibleListFilesIncludeTypeValues() []ListFilesIncludeType {
	return []ListFilesIncludeType{
		ListFilesIncludeTypeAttributes,
		ListFilesIncludeTypeEtag,
		ListFilesIncludeTypePermissionKey,
		ListFilesIncludeTypeTimestamps,
	}
}

type ListSharesIncludeType string

const (
	ListSharesIncludeTypeDeleted   ListSharesIncludeType = "deleted"
	ListSharesIncludeTypeMetadata  ListSharesIncludeType = "metadata"
	ListSharesIncludeTypeSnapshots ListSharesIncludeType = "snapshots"
)

// PossibleListSharesIncludeTypeValues returns the possible values for the ListSharesIncludeType const type.
func PossibleListSharesIncludeTypeValues() []ListSharesIncludeType {
	return []ListSharesIncludeType{
		ListSharesIncludeTypeDeleted,
		ListSharesIncludeTypeMetadata,
		ListSharesIncludeTypeSnapshots,
	}
}

type PermissionCopyModeType string

const (
	PermissionCopyModeTypeOverride PermissionCopyModeType = "override"
	PermissionCopyModeTypeSource   PermissionCopyModeType = "source"
)

// PossiblePermissionCopyModeTypeValues returns the possible values for the PermissionCopyModeType const type.
func PossiblePermissionCopyModeTypeValues() []PermissionCopyModeType {
	return []PermissionCopyModeType{
		PermissionCopyModeTypeOverride,
		PermissionCopyModeTypeSource,
	}
}

type ShareAccessTier string

const (
	ShareAccessTierCool                 ShareAccessTier = "Cool"
	ShareAccessTierHot                  ShareAccessTier = "Hot"
	ShareAccessTierTransactionOptimized ShareAccessTier = "TransactionOptimized"
)

// PossibleShareAccessTierValues returns the possible values for the ShareAccessTier const type.
func PossibleShareAccessTierValues() []ShareAccessTier {
	return []ShareAccessTier{
		ShareAccessTierCool,
		ShareAccessTierHot,
		ShareAccessTierTransactionOptimized,
	}
}

type ShareRootSquash string

const (
	ShareRootSquashAllSquash    ShareRootSquash = "AllSquash"
	ShareRootSquashNoRootSquash ShareRootSquash = "NoRootSquash"
	ShareRootSquashRootSquash   ShareRootSquash = "RootSquash"
)

// PossibleShareRootSquashValues returns the possible values for the ShareRootSquash const type.
func PossibleShareRootSquashValues() []ShareRootSquash {
	return []ShareRootSquash{
		ShareRootSquashAllSquash,
		ShareRootSquashNoRootSquash,
		ShareRootSquashRootSquash,
	}
}

type ShareTokenIntent string

const (
	ShareTokenIntentBackup ShareTokenIntent = "backup"
)

// PossibleShareTokenIntentValues returns the possible values for the ShareTokenIntent const type.
func PossibleShareTokenIntentValues() []ShareTokenIntent {
	return []ShareTokenIntent{
		ShareTokenIntentBackup,
	}
}

// StorageErrorCode - Error codes returned by the service
type StorageErrorCode string

const (
	StorageErrorCodeAccountAlreadyExists                 StorageErrorCode = "AccountAlreadyExists"
	StorageErrorCodeAccountBeingCreated                  StorageErrorCode = "AccountBeingCreated"
	StorageErrorCodeAccountIsDisabled                    StorageErrorCode = "AccountIsDisabled"
	StorageErrorCodeAuthenticationFailed                 StorageErrorCode = "AuthenticationFailed"
	StorageErrorCodeAuthorizationFailure                 StorageErrorCode = "AuthorizationFailure"
	StorageErrorCodeAuthorizationPermissionMismatch      StorageErrorCode = "AuthorizationPermissionMismatch"
	StorageErrorCodeAuthorizationProtocolMismatch        StorageErrorCode = "AuthorizationProtocolMismatch"
	StorageErrorCodeAuthorizationResourceTypeMismatch    StorageErrorCode = "AuthorizationResourceTypeMismatch"
	StorageErrorCodeAuthorizationServiceMismatch         StorageErrorCode = "AuthorizationServiceMismatch"
	StorageErrorCodeAuthorizationSourceIPMismatch        StorageErrorCode = "AuthorizationSourceIPMismatch"
	StorageErrorCodeCannotDeleteFileOrDirectory          StorageErrorCode = "CannotDeleteFileOrDirectory"
	StorageErrorCodeClientCacheFlushDelay                StorageErrorCode = "ClientCacheFlushDelay"
	StorageErrorCodeConditionHeadersNotSupported         StorageErrorCode = "ConditionHeadersNotSupported"
	StorageErrorCodeConditionNotMet                      StorageErrorCode = "ConditionNotMet"
	StorageErrorCodeContainerQuotaDowngradeNotAllowed    StorageErrorCode = "ContainerQuotaDowngradeNotAllowed"
	StorageErrorCodeDeletePending                        StorageErrorCode = "DeletePending"
	StorageErrorCodeDirectoryNotEmpty                    StorageErrorCode = "DirectoryNotEmpty"
	StorageErrorCodeEmptyMetadataKey                     StorageErrorCode = "EmptyMetadataKey"
	StorageErrorCodeFeatureVersionMismatch               StorageErrorCode = "FeatureVersionMismatch"
	StorageErrorCodeFileLockConflict                     StorageErrorCode = "FileLockConflict"
	StorageErrorCodeInsufficientAccountPermissions       StorageErrorCode = "InsufficientAccountPermissions"
	StorageErrorCodeInternalError                        StorageErrorCode = "InternalError"
	StorageErrorCodeInvalidAuthenticationInfo            StorageErrorCode = "InvalidAuthenticationInfo"
	StorageErrorCodeInvalidFileOrDirectoryPathName       StorageErrorCode = "InvalidFileOrDirectoryPathName"
	StorageErrorCodeInvalidHTTPVerb                      StorageErrorCode = "InvalidHttpVerb"
	StorageErrorCodeInvalidHeaderValue                   StorageErrorCode = "InvalidHeaderValue"
	StorageErrorCodeInvalidInput                         StorageErrorCode = "InvalidInput"
	StorageErrorCodeInvalidMD5                           StorageErrorCode = "InvalidMd5"
	StorageErrorCodeInvalidMetadata                      StorageErrorCode = "InvalidMetadata"
	StorageErrorCodeInvalidQueryParameterValue           StorageErrorCode = "InvalidQueryParameterValue"
	StorageErrorCodeInvalidRange                         StorageErrorCode = "InvalidRange"
	StorageErrorCodeInvalidResourceName                  StorageErrorCode = "InvalidResourceName"
	StorageErrorCodeInvalidURI                           StorageErrorCode = "InvalidUri"
	StorageErrorCodeInvalidXMLDocument                   StorageErrorCode = "InvalidXmlDocument"
	StorageErrorCodeInvalidXMLNodeValue                  StorageErrorCode = "InvalidXmlNodeValue"
	StorageErrorCodeMD5Mismatch                          StorageErrorCode = "Md5Mismatch"
	StorageErrorCodeMetadataTooLarge                     StorageErrorCode = "MetadataTooLarge"
	StorageErrorCodeMissingContentLengthHeader           StorageErrorCode = "MissingContentLengthHeader"
	StorageErrorCodeMissingRequiredHeader                StorageErrorCode = "MissingRequiredHeader"
	StorageErrorCodeMissingRequiredQueryParameter        StorageErrorCode = "MissingRequiredQueryParameter"
	StorageErrorCodeMissingRequiredXMLNode               StorageErrorCode = "MissingRequiredXmlNode"
	StorageErrorCodeMultipleConditionHeadersNotSupported StorageErrorCode = "MultipleConditionHeadersNotSupported"
	StorageErrorCodeOperationTimedOut                    StorageErrorCode = "OperationTimedOut"
	StorageErrorCodeOutOfRangeInput                      StorageErrorCode = "OutOfRangeInput"
	StorageErrorCodeOutOfRangeQueryParameterValue        StorageErrorCode = "OutOfRangeQueryParameterValue"
	StorageErrorCodeParentNotFound                       StorageErrorCode = "ParentNotFound"
	StorageErrorCodePreviousSnapshotNotFound             StorageErrorCode = "PreviousSnapshotNotFound"
	StorageErrorCodeReadOnlyAttribute                    StorageErrorCode = "ReadOnlyAttribute"
	StorageErrorCodeRequestBodyTooLarge                  StorageErrorCode = "RequestBodyTooLarge"
	StorageErrorCodeRequestURLFailedToParse              StorageErrorCode = "RequestUrlFailedToParse"
	StorageErrorCodeResourceAlreadyExists                StorageErrorCode = "ResourceAlreadyExists"
	StorageErrorCodeResourceNotFound                     StorageErrorCode = "ResourceNotFound"
	StorageErrorCodeResourceTypeMismatch                 StorageErrorCode = "ResourceTypeMismatch"
	StorageErrorCodeServerBusy                           StorageErrorCode = "ServerBusy"
	StorageErrorCodeShareAlreadyExists                   StorageErrorCode = "ShareAlreadyExists"
	StorageErrorCodeShareBeingDeleted                    StorageErrorCode = "ShareBeingDeleted"
	StorageErrorCodeShareDisabled                        StorageErrorCode = "ShareDisabled"
	StorageErrorCodeShareHasSnapshots                    StorageErrorCode = "ShareHasSnapshots"
	StorageErrorCodeShareNotFound                        StorageErrorCode = "ShareNotFound"
	StorageErrorCodeShareSnapshotCountExceeded           StorageErrorCode = "ShareSnapshotCountExceeded"
	StorageErrorCodeShareSnapshotInProgress              StorageErrorCode = "ShareSnapshotInProgress"
	StorageErrorCodeShareSnapshotOperationNotSupported   StorageErrorCode = "ShareSnapshotOperationNotSupported"
	StorageErrorCodeSharingViolation                     StorageErrorCode = "SharingViolation"
	StorageErrorCodeUnsupportedHTTPVerb                  StorageErrorCode = "UnsupportedHttpVerb"
	StorageErrorCodeUnsupportedHeader                    StorageErrorCode = "UnsupportedHeader"
	StorageErrorCodeUnsupportedQueryParameter            StorageErrorCode = "UnsupportedQueryParameter"
	StorageErrorCodeUnsupportedXMLNode                   StorageErrorCode = "UnsupportedXmlNode"
)

// PossibleStorageErrorCodeValues returns the possible values for the StorageErrorCode const type.
func PossibleStorageErrorCodeValues() []StorageErrorCode {
	return []StorageErrorCode{
		StorageErrorCodeAccountAlreadyExists,
		StorageErrorCodeAccountBeingCreated,
		StorageErrorCodeAccountIsDisabled,
		StorageErrorCodeAuthenticationFailed,
		StorageErrorCodeAuthorizationFailure,
		StorageErrorCodeAuthorizationPermissionMismatch,
		StorageErrorCodeAuthorizationProtocolMismatch,
		StorageErrorCodeAuthorizationResourceTypeMismatch,
		StorageErrorCodeAuthorizationServiceMismatch,
		StorageErrorCodeAuthorizationSourceIPMismatch,
		StorageErrorCodeCannotDeleteFileOrDirectory,
		StorageErrorCodeClientCacheFlushDelay,
		StorageErrorCodeConditionHeadersNotSupported,
		StorageErrorCodeConditionNotMet,
		StorageErrorCodeContainerQuotaDowngradeNotAllowed,
		StorageErrorCodeDeletePending,
		StorageErrorCodeDirectoryNotEmpty,
		StorageErrorCodeEmptyMetadataKey,
		StorageErrorCodeFeatureVersionMismatch,
		StorageErrorCodeFileLockConflict,
		StorageErrorCodeInsufficientAccountPermissions,
		StorageErrorCodeInternalError,
		StorageErrorCodeInvalidAuthenticationInfo,
		StorageErrorCodeInvalidFileOrDirectoryPathName,
		StorageErrorCodeInvalidHTTPVerb,
		StorageErrorCodeInvalidHeaderValue,
		StorageErrorCodeInvalidInput,
		StorageErrorCodeInvalidMD5,
		StorageErrorCodeInvalidMetadata,
		StorageErrorCodeInvalidQueryParameterValue,
		StorageErrorCodeInvalidRange,
		StorageErrorCodeInvalidResourceName,
		StorageErrorCodeInvalidURI,
		StorageErrorCodeInvalidXMLDocument,
		StorageErrorCodeInvalidXMLNodeValue,
		StorageErrorCodeMD5Mismatch,
		StorageErrorCodeMetadataTooLarge,
		StorageErrorCodeMissingContentLengthHeader,
		StorageErrorCodeMissingRequiredHeader,
		StorageErrorCodeMissingRequiredQueryParameter,
		StorageErrorCodeMissingRequiredXMLNode,
		StorageErrorCodeMultipleConditionHeadersNotSupported,
		StorageErrorCodeOperationTimedOut,
		StorageErrorCodeOutOfRangeInput,
		StorageErrorCodeOutOfRangeQueryParameterValue,
		StorageErrorCodeParentNotFound,
		StorageErrorCodePreviousSnapshotNotFound,
		StorageErrorCodeReadOnlyAttribute,
		StorageErrorCodeRequestBodyTooLarge,
		StorageErrorCodeRequestURLFailedToParse,
		StorageErrorCodeResourceAlreadyExists,
		StorageErrorCodeResourceNotFound,
		StorageErrorCodeResourceTypeMismatch,
		StorageErrorCodeServerBusy,
		StorageErrorCodeShareAlreadyExists,
		StorageErrorCodeShareBeingDeleted,
		StorageErrorCodeShareDisabled,
		StorageErrorCodeShareHasSnapshots,
		StorageErrorCodeShareNotFound,
		StorageErrorCodeShareSnapshotCountExceeded,
		StorageErrorCodeShareSnapshotInProgress,
		StorageErrorCodeShareSnapshotOperationNotSupported,
		StorageErrorCodeSharingViolation,
		StorageErrorCodeUnsupportedHTTPVerb,
		StorageErrorCodeUnsupportedHeader,
		StorageErrorCodeUnsupportedQueryParameter,
		StorageErrorCodeUnsupportedXMLNode,
	}
}
