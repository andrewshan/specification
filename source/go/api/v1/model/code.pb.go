// Code generated by protoc-gen-go. DO NOT EDIT.
// source: code.proto

package model // import "github.com/polarismesh/specification/source/go/api/v1/model"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Code int32

const (
	// base module status codes
	Code_Unknown                         Code = 0
	Code_ExecuteSuccess                  Code = 200000
	Code_DataNoChange                    Code = 200001
	Code_NoNeedUpdate                    Code = 200002
	Code_BadRequest                      Code = 400000
	Code_ParseException                  Code = 400001
	Code_EmptyRequest                    Code = 400002
	Code_BatchSizeOverLimit              Code = 400003
	Code_InvalidDiscoverResource         Code = 400004
	Code_InvalidRequestID                Code = 400100
	Code_InvalidUserName                 Code = 400101
	Code_InvalidUserToken                Code = 400102
	Code_InvalidParameter                Code = 400103
	Code_EmptyQueryParameter             Code = 400104
	Code_InvalidQueryInsParameter        Code = 400105
	Code_InvalidNamespaceName            Code = 400110
	Code_InvalidNamespaceOwners          Code = 400111
	Code_InvalidNamespaceToken           Code = 400112
	Code_InvalidServiceName              Code = 400120
	Code_InvalidServiceOwners            Code = 400121
	Code_InvalidServiceToken             Code = 400122
	Code_InvalidServiceMetadata          Code = 400123
	Code_InvalidServicePorts             Code = 400124
	Code_InvalidServiceBusiness          Code = 400125
	Code_InvalidServiceDepartment        Code = 400126
	Code_InvalidServiceCMDB              Code = 400127
	Code_InvalidServiceComment           Code = 400128
	Code_InvalidServiceAliasComment      Code = 400129
	Code_InvalidInstanceID               Code = 400130
	Code_InvalidInstanceHost             Code = 400131
	Code_InvalidInstancePort             Code = 400132
	Code_InvalidServiceAlias             Code = 400133
	Code_InvalidNamespaceWithAlias       Code = 400134
	Code_InvalidServiceAliasOwners       Code = 400135
	Code_InvalidInstanceProtocol         Code = 400136
	Code_InvalidInstanceVersion          Code = 400137
	Code_InvalidInstanceLogicSet         Code = 400138
	Code_InvalidInstanceIsolate          Code = 400139
	Code_HealthCheckNotOpen              Code = 400140
	Code_HeartbeatOnDisabledIns          Code = 400141
	Code_HeartbeatExceedLimit            Code = 400142
	Code_HeartbeatTypeNotFound           Code = 400143
	Code_InvalidMetadata                 Code = 400150
	Code_InvalidRateLimitID              Code = 400151
	Code_InvalidRateLimitLabels          Code = 400152
	Code_InvalidRateLimitAmounts         Code = 400153
	Code_InvalidRateLimitName            Code = 400154
	Code_InvalidCircuitBreakerID         Code = 400160
	Code_InvalidCircuitBreakerVersion    Code = 400161
	Code_InvalidCircuitBreakerName       Code = 400162
	Code_InvalidCircuitBreakerNamespace  Code = 400163
	Code_InvalidCircuitBreakerOwners     Code = 400164
	Code_InvalidCircuitBreakerToken      Code = 400165
	Code_InvalidCircuitBreakerBusiness   Code = 400166
	Code_InvalidCircuitBreakerDepartment Code = 400167
	Code_InvalidCircuitBreakerComment    Code = 400168
	Code_CircuitBreakerRuleExisted       Code = 400169
	Code_InvalidRoutingID                Code = 400700
	Code_InvalidRoutingPolicy            Code = 400701
	Code_InvalidRoutingName              Code = 400702
	Code_InvalidRoutingPriority          Code = 400703
	Code_InvalidFaultDetectID            Code = 400900
	Code_InvalidFaultDetectName          Code = 400901
	Code_InvalidFaultDetectNamespace     Code = 400902
	Code_FaultDetectRuleExisted          Code = 400903
	Code_InvalidMatchRule                Code = 400904
	// network relative codes
	Code_ServicesExistedMesh  Code = 400170
	Code_ResourcesExistedMesh Code = 400171
	Code_InvalidMeshParameter Code = 400172
	// platform relative codes
	Code_InvalidPlatformID         Code = 400180
	Code_InvalidPlatformName       Code = 400181
	Code_InvalidPlatformDomain     Code = 400182
	Code_InvalidPlatformQPS        Code = 400183
	Code_InvalidPlatformToken      Code = 400184
	Code_InvalidPlatformOwner      Code = 400185
	Code_InvalidPlatformDepartment Code = 400186
	Code_InvalidPlatformComment    Code = 400187
	Code_NotFoundPlatform          Code = 400188
	// flux relative codes
	Code_InvalidFluxRateLimitId             Code = 400190
	Code_InvalidFluxRateLimitQps            Code = 400191
	Code_InvalidFluxRateLimitSetKey         Code = 400192
	Code_ExistedResource                    Code = 400201
	Code_NotFoundResource                   Code = 400202
	Code_NamespaceExistedServices           Code = 400203
	Code_ServiceExistedInstances            Code = 400204
	Code_ServiceExistedRoutings             Code = 400205
	Code_ServiceExistedRateLimits           Code = 400206
	Code_ExistReleasedConfig                Code = 400207
	Code_SameInstanceRequest                Code = 400208
	Code_ServiceExistedCircuitBreakers      Code = 400209
	Code_ServiceExistedAlias                Code = 400210
	Code_NamespaceExistedMeshResources      Code = 400211
	Code_NamespaceExistedCircuitBreakers    Code = 400212
	Code_ServiceSubscribedByMeshes          Code = 400213
	Code_ServiceExistedFluxRateLimits       Code = 400214
	Code_NamespaceExistedConfigGroups       Code = 400219
	Code_NotFoundService                    Code = 400301
	Code_NotFoundRouting                    Code = 400302
	Code_NotFoundInstance                   Code = 400303
	Code_NotFoundServiceAlias               Code = 400304
	Code_NotFoundNamespace                  Code = 400305
	Code_NotFoundSourceService              Code = 400306
	Code_NotFoundRateLimit                  Code = 400307
	Code_NotFoundCircuitBreaker             Code = 400308
	Code_NotFoundMasterConfig               Code = 400309
	Code_NotFoundTagConfig                  Code = 400310
	Code_NotFoundTagConfigOrService         Code = 400311
	Code_ClientAPINotOpen                   Code = 400401
	Code_NotAllowBusinessService            Code = 400402
	Code_NotAllowAliasUpdate                Code = 400501
	Code_NotAllowAliasCreateInstance        Code = 400502
	Code_NotAllowAliasCreateRouting         Code = 400503
	Code_NotAllowCreateAliasForAlias        Code = 400504
	Code_NotAllowAliasCreateRateLimit       Code = 400505
	Code_NotAllowAliasBindRule              Code = 400506
	Code_NotAllowDifferentNamespaceBindRule Code = 400507
	Code_Unauthorized                       Code = 401000
	Code_NotAllowedAccess                   Code = 401001
	Code_CMDBNotFindHost                    Code = 404001
	Code_DataConflict                       Code = 409000
	Code_InstanceTooManyRequests            Code = 429001
	Code_IPRateLimit                        Code = 429002
	Code_APIRateLimit                       Code = 403003
	Code_ExecuteException                   Code = 500000
	Code_StoreLayerException                Code = 500001
	Code_CMDBPluginException                Code = 500002
	Code_ParseRoutingException              Code = 500004
	Code_ParseRateLimitException            Code = 500005
	Code_ParseCircuitBreakerException       Code = 500006
	Code_HeartbeatException                 Code = 500007
	Code_InstanceRegisTimeout               Code = 500008
	// config center status codes
	Code_InvalidConfigFileGroupName     Code = 400801
	Code_InvalidConfigFileName          Code = 400802
	Code_InvalidConfigFileContentLength Code = 400803
	Code_InvalidConfigFileFormat        Code = 400804
	Code_InvalidConfigFileTags          Code = 400805
	Code_InvalidWatchConfigFileFormat   Code = 400806
	Code_NotFoundResourceConfigFile     Code = 400807
	Code_InvalidConfigFileTemplateName  Code = 400808
	Code_EncryptConfigFileException     Code = 400809
	Code_GroupExistActiveRelease        Code = 400810
	Code_DecryptConfigFileException     Code = 400811
	// auth codes
	Code_InvalidUserOwners                      Code = 400410
	Code_InvalidUserID                          Code = 400411
	Code_InvalidUserPassword                    Code = 400412
	Code_InvalidUserMobile                      Code = 400413
	Code_InvalidUserEmail                       Code = 400414
	Code_InvalidUserGroupOwners                 Code = 400420
	Code_InvalidUserGroupID                     Code = 400421
	Code_InvalidAuthStrategyOwners              Code = 400430
	Code_InvalidAuthStrategyName                Code = 400431
	Code_InvalidAuthStrategyID                  Code = 400432
	Code_InvalidPrincipalType                   Code = 400440
	Code_UserExisted                            Code = 400215
	Code_UserGroupExisted                       Code = 400216
	Code_AuthStrategyRuleExisted                Code = 400217
	Code_SubAccountExisted                      Code = 400218
	Code_NotFoundUser                           Code = 400312
	Code_NotFoundOwnerUser                      Code = 400313
	Code_NotFoundUserGroup                      Code = 400314
	Code_NotFoundAuthStrategyRule               Code = 400315
	Code_NotAllowModifyDefaultStrategyPrincipal Code = 400508
	Code_NotAllowModifyOwnerDefaultStrategy     Code = 400509
	Code_EmptyAutToken                          Code = 401002
	Code_TokenDisabled                          Code = 401003
	Code_TokenNotExisted                        Code = 401004
	Code_AuthTokenForbidden                     Code = 403001
	Code_OperationRoleForbidden                 Code = 403002
)

var Code_name = map[int32]string{
	0:      "Unknown",
	200000: "ExecuteSuccess",
	200001: "DataNoChange",
	200002: "NoNeedUpdate",
	400000: "BadRequest",
	400001: "ParseException",
	400002: "EmptyRequest",
	400003: "BatchSizeOverLimit",
	400004: "InvalidDiscoverResource",
	400100: "InvalidRequestID",
	400101: "InvalidUserName",
	400102: "InvalidUserToken",
	400103: "InvalidParameter",
	400104: "EmptyQueryParameter",
	400105: "InvalidQueryInsParameter",
	400110: "InvalidNamespaceName",
	400111: "InvalidNamespaceOwners",
	400112: "InvalidNamespaceToken",
	400120: "InvalidServiceName",
	400121: "InvalidServiceOwners",
	400122: "InvalidServiceToken",
	400123: "InvalidServiceMetadata",
	400124: "InvalidServicePorts",
	400125: "InvalidServiceBusiness",
	400126: "InvalidServiceDepartment",
	400127: "InvalidServiceCMDB",
	400128: "InvalidServiceComment",
	400129: "InvalidServiceAliasComment",
	400130: "InvalidInstanceID",
	400131: "InvalidInstanceHost",
	400132: "InvalidInstancePort",
	400133: "InvalidServiceAlias",
	400134: "InvalidNamespaceWithAlias",
	400135: "InvalidServiceAliasOwners",
	400136: "InvalidInstanceProtocol",
	400137: "InvalidInstanceVersion",
	400138: "InvalidInstanceLogicSet",
	400139: "InvalidInstanceIsolate",
	400140: "HealthCheckNotOpen",
	400141: "HeartbeatOnDisabledIns",
	400142: "HeartbeatExceedLimit",
	400143: "HeartbeatTypeNotFound",
	400150: "InvalidMetadata",
	400151: "InvalidRateLimitID",
	400152: "InvalidRateLimitLabels",
	400153: "InvalidRateLimitAmounts",
	400154: "InvalidRateLimitName",
	400160: "InvalidCircuitBreakerID",
	400161: "InvalidCircuitBreakerVersion",
	400162: "InvalidCircuitBreakerName",
	400163: "InvalidCircuitBreakerNamespace",
	400164: "InvalidCircuitBreakerOwners",
	400165: "InvalidCircuitBreakerToken",
	400166: "InvalidCircuitBreakerBusiness",
	400167: "InvalidCircuitBreakerDepartment",
	400168: "InvalidCircuitBreakerComment",
	400169: "CircuitBreakerRuleExisted",
	400700: "InvalidRoutingID",
	400701: "InvalidRoutingPolicy",
	400702: "InvalidRoutingName",
	400703: "InvalidRoutingPriority",
	400900: "InvalidFaultDetectID",
	400901: "InvalidFaultDetectName",
	400902: "InvalidFaultDetectNamespace",
	400903: "FaultDetectRuleExisted",
	400904: "InvalidMatchRule",
	400170: "ServicesExistedMesh",
	400171: "ResourcesExistedMesh",
	400172: "InvalidMeshParameter",
	400180: "InvalidPlatformID",
	400181: "InvalidPlatformName",
	400182: "InvalidPlatformDomain",
	400183: "InvalidPlatformQPS",
	400184: "InvalidPlatformToken",
	400185: "InvalidPlatformOwner",
	400186: "InvalidPlatformDepartment",
	400187: "InvalidPlatformComment",
	400188: "NotFoundPlatform",
	400190: "InvalidFluxRateLimitId",
	400191: "InvalidFluxRateLimitQps",
	400192: "InvalidFluxRateLimitSetKey",
	400201: "ExistedResource",
	400202: "NotFoundResource",
	400203: "NamespaceExistedServices",
	400204: "ServiceExistedInstances",
	400205: "ServiceExistedRoutings",
	400206: "ServiceExistedRateLimits",
	400207: "ExistReleasedConfig",
	400208: "SameInstanceRequest",
	400209: "ServiceExistedCircuitBreakers",
	400210: "ServiceExistedAlias",
	400211: "NamespaceExistedMeshResources",
	400212: "NamespaceExistedCircuitBreakers",
	400213: "ServiceSubscribedByMeshes",
	400214: "ServiceExistedFluxRateLimits",
	400219: "NamespaceExistedConfigGroups",
	400301: "NotFoundService",
	400302: "NotFoundRouting",
	400303: "NotFoundInstance",
	400304: "NotFoundServiceAlias",
	400305: "NotFoundNamespace",
	400306: "NotFoundSourceService",
	400307: "NotFoundRateLimit",
	400308: "NotFoundCircuitBreaker",
	400309: "NotFoundMasterConfig",
	400310: "NotFoundTagConfig",
	400311: "NotFoundTagConfigOrService",
	400401: "ClientAPINotOpen",
	400402: "NotAllowBusinessService",
	400501: "NotAllowAliasUpdate",
	400502: "NotAllowAliasCreateInstance",
	400503: "NotAllowAliasCreateRouting",
	400504: "NotAllowCreateAliasForAlias",
	400505: "NotAllowAliasCreateRateLimit",
	400506: "NotAllowAliasBindRule",
	400507: "NotAllowDifferentNamespaceBindRule",
	401000: "Unauthorized",
	401001: "NotAllowedAccess",
	404001: "CMDBNotFindHost",
	409000: "DataConflict",
	429001: "InstanceTooManyRequests",
	429002: "IPRateLimit",
	403003: "APIRateLimit",
	500000: "ExecuteException",
	500001: "StoreLayerException",
	500002: "CMDBPluginException",
	500004: "ParseRoutingException",
	500005: "ParseRateLimitException",
	500006: "ParseCircuitBreakerException",
	500007: "HeartbeatException",
	500008: "InstanceRegisTimeout",
	400801: "InvalidConfigFileGroupName",
	400802: "InvalidConfigFileName",
	400803: "InvalidConfigFileContentLength",
	400804: "InvalidConfigFileFormat",
	400805: "InvalidConfigFileTags",
	400806: "InvalidWatchConfigFileFormat",
	400807: "NotFoundResourceConfigFile",
	400808: "InvalidConfigFileTemplateName",
	400809: "EncryptConfigFileException",
	400810: "GroupExistActiveRelease",
	400811: "DecryptConfigFileException",
	400410: "InvalidUserOwners",
	400411: "InvalidUserID",
	400412: "InvalidUserPassword",
	400413: "InvalidUserMobile",
	400414: "InvalidUserEmail",
	400420: "InvalidUserGroupOwners",
	400421: "InvalidUserGroupID",
	400430: "InvalidAuthStrategyOwners",
	400431: "InvalidAuthStrategyName",
	400432: "InvalidAuthStrategyID",
	400440: "InvalidPrincipalType",
	400215: "UserExisted",
	400216: "UserGroupExisted",
	400217: "AuthStrategyRuleExisted",
	400218: "SubAccountExisted",
	400312: "NotFoundUser",
	400313: "NotFoundOwnerUser",
	400314: "NotFoundUserGroup",
	400315: "NotFoundAuthStrategyRule",
	400508: "NotAllowModifyDefaultStrategyPrincipal",
	400509: "NotAllowModifyOwnerDefaultStrategy",
	401002: "EmptyAutToken",
	401003: "TokenDisabled",
	401004: "TokenNotExisted",
	403001: "AuthTokenForbidden",
	403002: "OperationRoleForbidden",
}
var Code_value = map[string]int32{
	"Unknown":                                0,
	"ExecuteSuccess":                         200000,
	"DataNoChange":                           200001,
	"NoNeedUpdate":                           200002,
	"BadRequest":                             400000,
	"ParseException":                         400001,
	"EmptyRequest":                           400002,
	"BatchSizeOverLimit":                     400003,
	"InvalidDiscoverResource":                400004,
	"InvalidRequestID":                       400100,
	"InvalidUserName":                        400101,
	"InvalidUserToken":                       400102,
	"InvalidParameter":                       400103,
	"EmptyQueryParameter":                    400104,
	"InvalidQueryInsParameter":               400105,
	"InvalidNamespaceName":                   400110,
	"InvalidNamespaceOwners":                 400111,
	"InvalidNamespaceToken":                  400112,
	"InvalidServiceName":                     400120,
	"InvalidServiceOwners":                   400121,
	"InvalidServiceToken":                    400122,
	"InvalidServiceMetadata":                 400123,
	"InvalidServicePorts":                    400124,
	"InvalidServiceBusiness":                 400125,
	"InvalidServiceDepartment":               400126,
	"InvalidServiceCMDB":                     400127,
	"InvalidServiceComment":                  400128,
	"InvalidServiceAliasComment":             400129,
	"InvalidInstanceID":                      400130,
	"InvalidInstanceHost":                    400131,
	"InvalidInstancePort":                    400132,
	"InvalidServiceAlias":                    400133,
	"InvalidNamespaceWithAlias":              400134,
	"InvalidServiceAliasOwners":              400135,
	"InvalidInstanceProtocol":                400136,
	"InvalidInstanceVersion":                 400137,
	"InvalidInstanceLogicSet":                400138,
	"InvalidInstanceIsolate":                 400139,
	"HealthCheckNotOpen":                     400140,
	"HeartbeatOnDisabledIns":                 400141,
	"HeartbeatExceedLimit":                   400142,
	"HeartbeatTypeNotFound":                  400143,
	"InvalidMetadata":                        400150,
	"InvalidRateLimitID":                     400151,
	"InvalidRateLimitLabels":                 400152,
	"InvalidRateLimitAmounts":                400153,
	"InvalidRateLimitName":                   400154,
	"InvalidCircuitBreakerID":                400160,
	"InvalidCircuitBreakerVersion":           400161,
	"InvalidCircuitBreakerName":              400162,
	"InvalidCircuitBreakerNamespace":         400163,
	"InvalidCircuitBreakerOwners":            400164,
	"InvalidCircuitBreakerToken":             400165,
	"InvalidCircuitBreakerBusiness":          400166,
	"InvalidCircuitBreakerDepartment":        400167,
	"InvalidCircuitBreakerComment":           400168,
	"CircuitBreakerRuleExisted":              400169,
	"InvalidRoutingID":                       400700,
	"InvalidRoutingPolicy":                   400701,
	"InvalidRoutingName":                     400702,
	"InvalidRoutingPriority":                 400703,
	"InvalidFaultDetectID":                   400900,
	"InvalidFaultDetectName":                 400901,
	"InvalidFaultDetectNamespace":            400902,
	"FaultDetectRuleExisted":                 400903,
	"InvalidMatchRule":                       400904,
	"ServicesExistedMesh":                    400170,
	"ResourcesExistedMesh":                   400171,
	"InvalidMeshParameter":                   400172,
	"InvalidPlatformID":                      400180,
	"InvalidPlatformName":                    400181,
	"InvalidPlatformDomain":                  400182,
	"InvalidPlatformQPS":                     400183,
	"InvalidPlatformToken":                   400184,
	"InvalidPlatformOwner":                   400185,
	"InvalidPlatformDepartment":              400186,
	"InvalidPlatformComment":                 400187,
	"NotFoundPlatform":                       400188,
	"InvalidFluxRateLimitId":                 400190,
	"InvalidFluxRateLimitQps":                400191,
	"InvalidFluxRateLimitSetKey":             400192,
	"ExistedResource":                        400201,
	"NotFoundResource":                       400202,
	"NamespaceExistedServices":               400203,
	"ServiceExistedInstances":                400204,
	"ServiceExistedRoutings":                 400205,
	"ServiceExistedRateLimits":               400206,
	"ExistReleasedConfig":                    400207,
	"SameInstanceRequest":                    400208,
	"ServiceExistedCircuitBreakers":          400209,
	"ServiceExistedAlias":                    400210,
	"NamespaceExistedMeshResources":          400211,
	"NamespaceExistedCircuitBreakers":        400212,
	"ServiceSubscribedByMeshes":              400213,
	"ServiceExistedFluxRateLimits":           400214,
	"NamespaceExistedConfigGroups":           400219,
	"NotFoundService":                        400301,
	"NotFoundRouting":                        400302,
	"NotFoundInstance":                       400303,
	"NotFoundServiceAlias":                   400304,
	"NotFoundNamespace":                      400305,
	"NotFoundSourceService":                  400306,
	"NotFoundRateLimit":                      400307,
	"NotFoundCircuitBreaker":                 400308,
	"NotFoundMasterConfig":                   400309,
	"NotFoundTagConfig":                      400310,
	"NotFoundTagConfigOrService":             400311,
	"ClientAPINotOpen":                       400401,
	"NotAllowBusinessService":                400402,
	"NotAllowAliasUpdate":                    400501,
	"NotAllowAliasCreateInstance":            400502,
	"NotAllowAliasCreateRouting":             400503,
	"NotAllowCreateAliasForAlias":            400504,
	"NotAllowAliasCreateRateLimit":           400505,
	"NotAllowAliasBindRule":                  400506,
	"NotAllowDifferentNamespaceBindRule":     400507,
	"Unauthorized":                           401000,
	"NotAllowedAccess":                       401001,
	"CMDBNotFindHost":                        404001,
	"DataConflict":                           409000,
	"InstanceTooManyRequests":                429001,
	"IPRateLimit":                            429002,
	"APIRateLimit":                           403003,
	"ExecuteException":                       500000,
	"StoreLayerException":                    500001,
	"CMDBPluginException":                    500002,
	"ParseRoutingException":                  500004,
	"ParseRateLimitException":                500005,
	"ParseCircuitBreakerException":           500006,
	"HeartbeatException":                     500007,
	"InstanceRegisTimeout":                   500008,
	"InvalidConfigFileGroupName":             400801,
	"InvalidConfigFileName":                  400802,
	"InvalidConfigFileContentLength":         400803,
	"InvalidConfigFileFormat":                400804,
	"InvalidConfigFileTags":                  400805,
	"InvalidWatchConfigFileFormat":           400806,
	"NotFoundResourceConfigFile":             400807,
	"InvalidConfigFileTemplateName":          400808,
	"EncryptConfigFileException":             400809,
	"GroupExistActiveRelease":                400810,
	"DecryptConfigFileException":             400811,
	"InvalidUserOwners":                      400410,
	"InvalidUserID":                          400411,
	"InvalidUserPassword":                    400412,
	"InvalidUserMobile":                      400413,
	"InvalidUserEmail":                       400414,
	"InvalidUserGroupOwners":                 400420,
	"InvalidUserGroupID":                     400421,
	"InvalidAuthStrategyOwners":              400430,
	"InvalidAuthStrategyName":                400431,
	"InvalidAuthStrategyID":                  400432,
	"InvalidPrincipalType":                   400440,
	"UserExisted":                            400215,
	"UserGroupExisted":                       400216,
	"AuthStrategyRuleExisted":                400217,
	"SubAccountExisted":                      400218,
	"NotFoundUser":                           400312,
	"NotFoundOwnerUser":                      400313,
	"NotFoundUserGroup":                      400314,
	"NotFoundAuthStrategyRule":               400315,
	"NotAllowModifyDefaultStrategyPrincipal": 400508,
	"NotAllowModifyOwnerDefaultStrategy":     400509,
	"EmptyAutToken":                          401002,
	"TokenDisabled":                          401003,
	"TokenNotExisted":                        401004,
	"AuthTokenForbidden":                     403001,
	"OperationRoleForbidden":                 403002,
}

func (x Code) String() string {
	return proto.EnumName(Code_name, int32(x))
}
func (Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_code_cc736450b8dbd25a, []int{0}
}

func init() {
	proto.RegisterEnum("v1.Code", Code_name, Code_value)
}

func init() { proto.RegisterFile("code.proto", fileDescriptor_code_cc736450b8dbd25a) }

var fileDescriptor_code_cc736450b8dbd25a = []byte{
	// 2018 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x58, 0x5b, 0x90, 0x24, 0x45,
	0x15, 0x35, 0x0c, 0x43, 0xc3, 0x02, 0xf5, 0x52, 0xbc, 0x8a, 0x05, 0x06, 0x45, 0x31, 0x0c, 0xc3,
	0x98, 0x89, 0x0d, 0x3f, 0xfd, 0x9a, 0xe9, 0x9e, 0x71, 0x3b, 0xdc, 0x99, 0xe9, 0x9d, 0x9e, 0x95,
	0x08, 0xff, 0xb2, 0xab, 0xee, 0x74, 0x67, 0x6c, 0x75, 0x65, 0x5b, 0x99, 0x35, 0x3b, 0xcd, 0x17,
	0x8f, 0xe5, 0xe1, 0x3b, 0xf4, 0xc3, 0xd5, 0xf5, 0x11, 0x2c, 0xec, 0x0b, 0x50, 0x40, 0xa5, 0xab,
	0x6b, 0xa7, 0x5b, 0xc0, 0xf9, 0x03, 0xd7, 0x37, 0x3e, 0xf9, 0xf5, 0x01, 0x6a, 0xa8, 0x3f, 0x2a,
	0xb0, 0xa0, 0x46, 0x66, 0x56, 0x56, 0x55, 0xf6, 0xf6, 0xfa, 0xd9, 0x79, 0xee, 0xbd, 0x79, 0x1f,
	0xe7, 0xde, 0xba, 0xd9, 0x8e, 0xe3, 0xb3, 0x00, 0xe7, 0xfb, 0x31, 0x13, 0xcc, 0x7d, 0xf3, 0xf6,
	0xfe, 0x0f, 0xbe, 0x78, 0x9b, 0xf3, 0x96, 0x1a, 0x0b, 0xd0, 0xbd, 0xc2, 0x79, 0xdb, 0xe1, 0xe8,
	0x48, 0xc4, 0x8e, 0x46, 0xf0, 0x26, 0xf7, 0x1a, 0xe7, 0x9d, 0xcb, 0x3b, 0xe8, 0x27, 0x02, 0x5b,
	0x89, 0xef, 0x23, 0xe7, 0xf0, 0xec, 0x89, 0x2b, 0x5d, 0xd7, 0xb9, 0xb2, 0x4e, 0x04, 0x59, 0x63,
	0xb5, 0x2e, 0x89, 0x3a, 0x08, 0x3f, 0xd0, 0x67, 0x6b, 0x6c, 0x0d, 0x31, 0x38, 0xdc, 0x0f, 0x88,
	0x40, 0xd8, 0x3b, 0x71, 0xa5, 0x0b, 0x8e, 0xb3, 0x44, 0x82, 0x0d, 0xfc, 0x64, 0x82, 0x5c, 0xc0,
	0x9d, 0x43, 0x4f, 0xda, 0x6b, 0x92, 0x98, 0xe3, 0xf2, 0x8e, 0x8f, 0x7d, 0x41, 0x59, 0x04, 0x77,
	0x0d, 0x3d, 0xa9, 0xbb, 0xdc, 0xeb, 0x8b, 0x81, 0x91, 0xbc, 0x7b, 0xe8, 0xb9, 0x9e, 0xe3, 0x2e,
	0x11, 0xe1, 0x77, 0x5b, 0xf4, 0x0e, 0x5c, 0xdf, 0xc6, 0xf8, 0x20, 0xed, 0x51, 0x01, 0xf7, 0x0c,
	0x3d, 0xf7, 0x66, 0xe7, 0xfa, 0x46, 0xb4, 0x4d, 0x42, 0x1a, 0xd4, 0x29, 0xf7, 0xd9, 0x36, 0xc6,
	0x1b, 0xc8, 0x59, 0x12, 0xfb, 0x08, 0xc7, 0x86, 0x9e, 0x7b, 0x9d, 0x03, 0x39, 0x9c, 0x9b, 0x6b,
	0xd4, 0xe1, 0xf7, 0x43, 0xcf, 0xbd, 0xd6, 0x79, 0x57, 0x7e, 0x7e, 0x98, 0x63, 0xbc, 0x46, 0x7a,
	0x08, 0x7f, 0xb0, 0xc4, 0xe5, 0xf1, 0x26, 0x3b, 0x82, 0x11, 0xfc, 0xd1, 0x3a, 0x6f, 0x92, 0x98,
	0xf4, 0x50, 0x60, 0x0c, 0x7f, 0x1a, 0x7a, 0xee, 0x0d, 0xce, 0xd5, 0xca, 0xd7, 0x43, 0x09, 0xc6,
	0x83, 0x12, 0x7a, 0x69, 0xe8, 0xb9, 0x73, 0x8e, 0x97, 0xab, 0x28, 0xb0, 0x11, 0xf1, 0x12, 0x7f,
	0x79, 0xe8, 0xb9, 0xfb, 0x9c, 0x6b, 0x72, 0x5c, 0xde, 0xce, 0xfb, 0xc4, 0x47, 0xe5, 0xc6, 0xdf,
	0x86, 0x9e, 0x7b, 0x93, 0x73, 0xdd, 0x34, 0xb6, 0x7e, 0x34, 0xc2, 0x98, 0xc3, 0xdf, 0x87, 0x9e,
	0x7b, 0xa3, 0x73, 0xed, 0x34, 0xaa, 0x3d, 0xfd, 0x87, 0xce, 0x54, 0x0e, 0xb6, 0x30, 0xde, 0xa6,
	0xb9, 0xd1, 0x57, 0xac, 0x0b, 0x73, 0x24, 0x37, 0xf9, 0xaa, 0x8e, 0xc3, 0xc6, 0xb4, 0xc1, 0xd7,
	0x2c, 0x5f, 0x72, 0x68, 0x15, 0x05, 0x09, 0x88, 0x20, 0x70, 0x71, 0x96, 0x62, 0x93, 0xc5, 0x82,
	0xc3, 0xeb, 0xb3, 0x14, 0x97, 0x12, 0x4e, 0x23, 0xc9, 0x9a, 0x37, 0xac, 0xf4, 0xe4, 0x68, 0x1d,
	0xfb, 0x24, 0x16, 0x3d, 0x8c, 0x04, 0xfc, 0x67, 0x56, 0x1c, 0xb5, 0xd5, 0xfa, 0x12, 0xfc, 0xd7,
	0x0a, 0xdf, 0x20, 0xac, 0xa7, 0xd4, 0xee, 0x4c, 0x3d, 0xf7, 0xdd, 0xce, 0x3e, 0x1b, 0x5c, 0x0c,
	0x29, 0xe1, 0x46, 0xe2, 0xae, 0xd4, 0x73, 0xaf, 0x77, 0xae, 0xca, 0x25, 0x1a, 0x11, 0x17, 0x24,
	0xf2, 0xb1, 0x51, 0x87, 0xbb, 0xd3, 0x6a, 0x28, 0x06, 0x38, 0xc0, 0xb8, 0x80, 0x7b, 0x66, 0x42,
	0x32, 0x4c, 0x38, 0x96, 0xce, 0x48, 0x80, 0xba, 0x10, 0xee, 0x4d, 0x3d, 0xf7, 0x16, 0xe7, 0x86,
	0xe9, 0x3a, 0xdd, 0x4e, 0x45, 0x57, 0x0b, 0xdc, 0x67, 0x09, 0x54, 0x75, 0xf3, 0xb2, 0xdc, 0x9f,
	0x56, 0xc9, 0x5d, 0xdc, 0x2b, 0x5b, 0xd4, 0x67, 0x21, 0x3c, 0x90, 0x56, 0x33, 0x6c, 0xe0, 0x8f,
	0x63, 0xcc, 0x65, 0x1f, 0x7d, 0x6a, 0xa6, 0xf2, 0x41, 0xd6, 0xa1, 0x7e, 0x0b, 0x05, 0x7c, 0x7a,
	0xa6, 0x72, 0x83, 0xb3, 0x50, 0x36, 0xeb, 0x67, 0x52, 0x95, 0xfe, 0x03, 0x48, 0x42, 0xd1, 0xad,
	0x75, 0xd1, 0x3f, 0xb2, 0xc6, 0xc4, 0x7a, 0x1f, 0x23, 0xf8, 0xac, 0xd6, 0x3b, 0x80, 0x24, 0x16,
	0x6d, 0x24, 0x62, 0x3d, 0xaa, 0x53, 0x4e, 0xda, 0x21, 0x4a, 0x1b, 0xf0, 0xb9, 0x54, 0x91, 0xac,
	0x40, 0x65, 0x5b, 0x63, 0xa0, 0x5b, 0xf5, 0xf3, 0xa9, 0x2a, 0x5c, 0x81, 0x6d, 0x0e, 0xfa, 0xb8,
	0xc6, 0xc4, 0x0a, 0x4b, 0xa2, 0x00, 0xbe, 0x90, 0x56, 0x1b, 0xb2, 0xe0, 0xd7, 0x97, 0xd3, 0x2a,
	0x0d, 0x36, 0x88, 0x40, 0x65, 0xab, 0x51, 0x87, 0xe3, 0x96, 0xff, 0x05, 0x72, 0x90, 0xb4, 0x31,
	0xe4, 0xf0, 0x15, 0x2b, 0xf8, 0x02, 0x5d, 0xec, 0xb1, 0x24, 0x12, 0x1c, 0xbe, 0x9a, 0x56, 0x7b,
	0xa1, 0x80, 0x55, 0x9f, 0x9c, 0xb0, 0x54, 0x6b, 0x34, 0xf6, 0x13, 0x2a, 0x96, 0x62, 0x24, 0x47,
	0x30, 0x6e, 0xd4, 0xe1, 0xc1, 0xd4, 0x73, 0x6f, 0x75, 0x6e, 0x9a, 0x09, 0x9b, 0xd4, 0x9f, 0xb4,
	0x0a, 0x6b, 0xcb, 0xa8, 0x3b, 0x1e, 0x4a, 0x3d, 0xf7, 0x7d, 0xce, 0xdc, 0x65, 0x05, 0x14, 0x51,
	0xe0, 0xe1, 0xd4, 0x73, 0xdf, 0xe3, 0xdc, 0x38, 0x53, 0x2a, 0x67, 0xc8, 0x29, 0x8b, 0xef, 0xb6,
	0x88, 0xee, 0xdf, 0xd3, 0xa9, 0xe7, 0xbe, 0xd7, 0xb9, 0x79, 0xa6, 0x44, 0xd1, 0x8d, 0x67, 0x52,
	0xcf, 0xbd, 0xcd, 0xb9, 0x65, 0xa6, 0x50, 0xa5, 0x29, 0xcf, 0xfe, 0x9f, 0xd8, 0x4d, 0x7f, 0x9d,
	0xd3, 0xb1, 0xdb, 0xe0, 0x46, 0x12, 0xe2, 0xf2, 0x0e, 0xe5, 0x02, 0x03, 0x78, 0x24, 0xb5, 0x46,
	0x32, 0x4b, 0x04, 0x8d, 0x3a, 0x8d, 0x3a, 0x4c, 0x76, 0xad, 0x9a, 0xe8, 0xf3, 0x26, 0x0b, 0xa9,
	0x3f, 0x80, 0xef, 0xef, 0x5a, 0x34, 0xd0, 0x98, 0xca, 0xe4, 0xd3, 0xbb, 0x16, 0x0d, 0x72, 0xad,
	0x98, 0xb2, 0x98, 0x8a, 0x01, 0x3c, 0x63, 0xd9, 0x5c, 0x21, 0x49, 0x28, 0xea, 0x28, 0xd0, 0x97,
	0x04, 0x3a, 0x36, 0xa9, 0x6a, 0x56, 0x30, 0x65, 0xf7, 0xde, 0x49, 0x35, 0xf7, 0x53, 0xa8, 0x2e,
	0xcf, 0x7d, 0xda, 0x40, 0x05, 0xab, 0x86, 0x79, 0xff, 0xa4, 0x1a, 0xe6, 0xaa, 0xfc, 0x72, 0x49,
	0x18, 0x1e, 0x98, 0xa8, 0x81, 0x91, 0x77, 0x3b, 0xcf, 0xc5, 0x57, 0x91, 0x77, 0xe1, 0x51, 0xcd,
	0x4a, 0xf3, 0xf1, 0xb2, 0xb0, 0xc7, 0x2c, 0xc6, 0xca, 0xa3, 0xf2, 0x53, 0xf2, 0x2d, 0x6b, 0xa4,
	0x35, 0x43, 0x22, 0xb6, 0x58, 0xdc, 0x6b, 0xd4, 0xe1, 0x29, 0x6b, 0x38, 0x19, 0x40, 0xc5, 0x37,
	0x4c, 0xab, 0x53, 0xd4, 0x40, 0x75, 0xd6, 0x23, 0x34, 0x82, 0xd4, 0xea, 0x3a, 0x03, 0x1e, 0x6a,
	0xb6, 0x60, 0x64, 0xb9, 0x61, 0x10, 0xcd, 0xb4, 0x6c, 0x26, 0xa6, 0x88, 0x0a, 0xe7, 0xad, 0x8e,
	0x28, 0xae, 0x2b, 0xa9, 0xb5, 0x6b, 0xb5, 0xb3, 0x11, 0x30, 0xa4, 0x1a, 0x6b, 0xce, 0x98, 0x69,
	0x61, 0x60, 0x98, 0x58, 0x5a, 0x2b, 0x61, 0xb2, 0x53, 0x8e, 0x88, 0x00, 0x9e, 0xb6, 0x3a, 0xd9,
	0x42, 0x0f, 0xf5, 0x39, 0x3c, 0x63, 0xf5, 0x8e, 0x05, 0xb7, 0x50, 0x7c, 0x0c, 0x07, 0xf0, 0xac,
	0x1e, 0x4a, 0x79, 0x1d, 0x8a, 0xa5, 0xe2, 0x39, 0xdb, 0x9b, 0xe2, 0xfc, 0xf9, 0x54, 0x7d, 0xd3,
	0x0a, 0x86, 0xe4, 0x7a, 0xa6, 0xd4, 0xf0, 0x43, 0xed, 0x4f, 0xfe, 0x3b, 0x47, 0xcd, 0xe4, 0xe5,
	0x70, 0x41, 0x07, 0x63, 0xc3, 0x39, 0xa3, 0x39, 0xfc, 0x48, 0x1b, 0x9f, 0x42, 0x8d, 0xc3, 0x1c,
	0x7e, 0xac, 0x6b, 0xad, 0x80, 0x0d, 0x0c, 0x91, 0x70, 0x0c, 0x6a, 0x2c, 0xda, 0xa2, 0x1d, 0xf8,
	0x89, 0x86, 0x5a, 0xa4, 0x87, 0xe6, 0x36, 0xb3, 0x58, 0xfd, 0x54, 0x4f, 0x07, 0xdb, 0xaa, 0xdd,
	0xbb, 0x1c, 0x7e, 0x96, 0x56, 0x29, 0x9b, 0x0b, 0xe9, 0x4f, 0xd8, 0xcf, 0xb5, 0xfe, 0x74, 0xc8,
	0x92, 0x9f, 0x05, 0x8d, 0xe1, 0x17, 0x7a, 0xba, 0x4c, 0x0b, 0x4d, 0x5f, 0xf3, 0x82, 0xe6, 0x48,
	0x7e, 0x4d, 0x2b, 0x69, 0x73, 0x3f, 0xa6, 0x6d, 0x0c, 0x96, 0x06, 0xd2, 0x1c, 0x72, 0xf8, 0xa5,
	0x1e, 0x3f, 0xb6, 0x1f, 0x56, 0xdd, 0x38, 0xfc, 0x4a, 0xcb, 0x5c, 0x72, 0x97, 0xca, 0xc4, 0x47,
	0x63, 0x96, 0xf4, 0x39, 0xbc, 0xa8, 0xcb, 0x6a, 0xea, 0x97, 0xdb, 0x83, 0x6f, 0x8f, 0xac, 0xe3,
	0x3c, 0xf3, 0xf0, 0xf8, 0xc8, 0xaa, 0xb6, 0xc9, 0x20, 0x3c, 0x31, 0x52, 0x74, 0x9f, 0xb2, 0xa2,
	0xd3, 0xf2, 0xe4, 0x48, 0x75, 0xa4, 0xc1, 0xca, 0x99, 0xf1, 0x9d, 0x91, 0x6a, 0xbb, 0x42, 0x49,
	0x65, 0xc8, 0x38, 0xf0, 0x5d, 0x5b, 0xab, 0x88, 0x0a, 0xbe, 0x37, 0x52, 0xcc, 0x30, 0x80, 0x9d,
	0x38, 0x78, 0xca, 0x76, 0x64, 0x95, 0x70, 0x21, 0xc7, 0xb1, 0x2a, 0xfd, 0xd0, 0x36, 0xb9, 0x49,
	0x3a, 0x39, 0x90, 0x8e, 0x14, 0xf9, 0x2f, 0x01, 0xd6, 0x63, 0xe3, 0xcd, 0x48, 0xc7, 0x5d, 0x0b,
	0x29, 0x46, 0x62, 0xb1, 0xd9, 0x30, 0x0b, 0xc0, 0x17, 0x33, 0xc5, 0xe2, 0x35, 0x26, 0x16, 0xc3,
	0x90, 0x1d, 0x35, 0xdf, 0x10, 0xa3, 0xf6, 0xa5, 0x4c, 0x91, 0xc5, 0xc0, 0x2a, 0x1f, 0xf9, 0x0b,
	0xe0, 0x9f, 0x99, 0x9a, 0xa9, 0x16, 0x54, 0x8b, 0x91, 0x88, 0x82, 0x96, 0xf0, 0xaf, 0xcc, 0xb8,
	0x35, 0x2d, 0x62, 0xca, 0xf1, 0x6f, 0xdb, 0x88, 0x06, 0x95, 0xdc, 0x0a, 0x8b, 0x75, 0xf6, 0x5f,
	0xc9, 0x34, 0x07, 0x66, 0x18, 0x29, 0x52, 0xfa, 0x6a, 0x66, 0x0a, 0x51, 0xca, 0x2c, 0xd1, 0x28,
	0x50, 0x33, 0xfa, 0xb5, 0xcc, 0x73, 0x3f, 0xe0, 0xdc, 0x6a, 0xc0, 0x3a, 0xdd, 0xda, 0xc2, 0x18,
	0xa3, 0x72, 0xf6, 0x17, 0x92, 0x17, 0x33, 0xf5, 0x58, 0x39, 0x1c, 0x91, 0x44, 0x74, 0x59, 0x4c,
	0xef, 0xc0, 0x00, 0x5e, 0x9a, 0x18, 0xc2, 0x28, 0x6d, 0x0c, 0x16, 0xf5, 0x43, 0xe9, 0xe5, 0x89,
	0xe2, 0x97, 0x5c, 0x62, 0x65, 0xda, 0x69, 0x14, 0xa8, 0xe5, 0xf2, 0xe4, 0x0b, 0x9e, 0x79, 0x3f,
	0xc9, 0x12, 0x84, 0xd4, 0x17, 0x70, 0xee, 0x62, 0x3e, 0xb9, 0x74, 0x5a, 0x36, 0x19, 0x5b, 0x25,
	0x91, 0x79, 0x0d, 0x71, 0x78, 0xee, 0xf8, 0x3e, 0xf7, 0x2a, 0xe7, 0x8a, 0x46, 0xb3, 0x8c, 0xe7,
	0xf9, 0xe3, 0xfb, 0xa4, 0x95, 0xc5, 0x66, 0xa3, 0x3c, 0x1b, 0x5f, 0x50, 0x8e, 0xe4, 0xef, 0xb5,
	0xf2, 0x85, 0xf5, 0xe0, 0xde, 0x9c, 0xea, 0x67, 0xc1, 0x62, 0x3c, 0x48, 0x06, 0x18, 0x97, 0xd0,
	0x49, 0x0d, 0x49, 0x1f, 0x9b, 0x61, 0xd2, 0xa1, 0x51, 0x09, 0x3d, 0xb4, 0x37, 0x27, 0x33, 0xa6,
	0x5e, 0x6b, 0x79, 0x31, 0x4a, 0xf0, 0xd4, 0xde, 0x9c, 0x74, 0x58, 0x83, 0xc6, 0x81, 0x12, 0x3e,
	0xbd, 0x37, 0x27, 0x2b, 0xa2, 0x60, 0x9b, 0xbd, 0xa5, 0xcc, 0x99, 0xbd, 0xb9, 0x7c, 0xe5, 0x2c,
	0x57, 0x47, 0x8d, 0x9c, 0xdd, 0x9b, 0xd3, 0x1f, 0x16, 0x33, 0xbb, 0x3a, 0x94, 0x6f, 0xd2, 0x1e,
	0xb2, 0x44, 0xc0, 0xb9, 0xbd, 0xb9, 0xea, 0x02, 0xa4, 0x38, 0xbc, 0x42, 0x43, 0x54, 0xad, 0xae,
	0xbe, 0x74, 0x27, 0xc7, 0xd5, 0x2f, 0x5d, 0x29, 0xa1, 0x17, 0xb1, 0xb1, 0xb5, 0x88, 0x15, 0x60,
	0x8d, 0x45, 0x02, 0x23, 0x71, 0x10, 0xa3, 0x8e, 0xe8, 0xc2, 0xc3, 0x63, 0x6b, 0x25, 0x2c, 0xa4,
	0x56, 0x58, 0xdc, 0x23, 0x02, 0x4e, 0x5d, 0xe6, 0x86, 0x4d, 0xd2, 0xe1, 0x70, 0x7a, 0x5c, 0xdd,
	0x99, 0x6e, 0x97, 0x7b, 0xc0, 0x25, 0x06, 0xce, 0x8c, 0xad, 0x66, 0x34, 0x93, 0xb3, 0x94, 0x83,
	0xb3, 0x63, 0x6b, 0x8b, 0x2b, 0xaf, 0xc0, 0x5e, 0x5f, 0x2e, 0xec, 0x2a, 0x98, 0x73, 0xda, 0xcc,
	0x72, 0xe4, 0xc7, 0x83, 0xbe, 0x28, 0x85, 0xca, 0x4c, 0x3e, 0xa2, 0x03, 0x51, 0xc9, 0x51, 0x93,
	0x71, 0xd1, 0x17, 0x74, 0x1b, 0xf3, 0xef, 0x05, 0x3c, 0xaa, 0x0d, 0xd4, 0xf1, 0xb2, 0x06, 0x1e,
	0x1b, 0x57, 0x57, 0x0d, 0xf9, 0x40, 0xce, 0x17, 0xd1, 0x13, 0x99, 0xe7, 0x5e, 0xed, 0xbc, 0xa3,
	0x02, 0x34, 0xea, 0xf0, 0xb5, 0xac, 0xba, 0x7f, 0xc8, 0xc3, 0x26, 0xe1, 0xfc, 0x28, 0x8b, 0x03,
	0xf8, 0x7a, 0x36, 0x6d, 0x68, 0x95, 0xb5, 0x65, 0xa4, 0xdf, 0xc8, 0xa6, 0x9f, 0xe0, 0xcb, 0x3d,
	0x42, 0x43, 0xf8, 0x66, 0x56, 0xfd, 0xd4, 0xcb, 0x73, 0x15, 0x85, 0xd9, 0x83, 0xb3, 0xea, 0xc6,
	0x52, 0xa0, 0x8d, 0x3a, 0x9c, 0xce, 0xaa, 0x9b, 0xc7, 0x62, 0x22, 0xba, 0x2d, 0x11, 0x13, 0x81,
	0x9d, 0x41, 0xae, 0xfa, 0x78, 0x56, 0x2d, 0x6e, 0x55, 0x40, 0x25, 0xf5, 0x89, 0xac, 0x5a, 0xdc,
	0x2a, 0xdc, 0xa8, 0xc3, 0x93, 0x99, 0xb5, 0xf2, 0xc4, 0x34, 0xf2, 0x69, 0x9f, 0x84, 0xf2, 0x65,
	0x03, 0x59, 0xe6, 0xc9, 0x26, 0x55, 0x11, 0xe4, 0x3b, 0xe1, 0xaf, 0xf5, 0xe2, 0x50, 0xb8, 0x67,
	0xce, 0x7f, 0xa3, 0x17, 0x83, 0xaa, 0xf1, 0xea, 0x2a, 0xf9, 0x5b, 0xbd, 0xdf, 0xb5, 0x92, 0xf6,
	0xa2, 0xef, 0xcb, 0x07, 0x8c, 0x01, 0x7e, 0x97, 0x7a, 0xfa, 0x6f, 0x16, 0xcd, 0x1b, 0x69, 0x17,
	0x32, 0x7b, 0xe2, 0xab, 0x20, 0x15, 0x70, 0xde, 0x06, 0x0a, 0x27, 0x60, 0x77, 0xa4, 0xd7, 0x96,
	0x1c, 0x98, 0xf6, 0x02, 0xc6, 0x23, 0xcf, 0xfd, 0x90, 0xf3, 0x7e, 0x33, 0xcf, 0x56, 0x59, 0x40,
	0xb7, 0x06, 0x75, 0xdc, 0x92, 0x7b, 0xaf, 0x11, 0x2c, 0x42, 0x87, 0xd7, 0xed, 0xd9, 0xa9, 0xa5,
	0x95, 0x17, 0x53, 0x2a, 0xf0, 0x86, 0xa6, 0x8c, 0xfa, 0xf3, 0x64, 0x31, 0x11, 0x7a, 0x89, 0xfc,
	0xf3, 0x44, 0x1d, 0xaa, 0x1f, 0xe6, 0x61, 0x09, 0x7f, 0xd1, 0x93, 0x53, 0x1d, 0xae, 0xb1, 0x22,
	0xfc, 0xbf, 0x4e, 0x54, 0xd1, 0xa5, 0xc3, 0x0a, 0x5a, 0x61, 0x71, 0x9b, 0x06, 0x01, 0x46, 0x70,
	0xfe, 0x82, 0x22, 0xcb, 0x7a, 0x1f, 0x63, 0x22, 0x79, 0xbb, 0xc1, 0x54, 0xaf, 0xe5, 0xe8, 0xee,
	0x05, 0x6f, 0x69, 0xc7, 0x99, 0xf7, 0x59, 0x6f, 0x5e, 0x60, 0xe4, 0x63, 0x24, 0xe6, 0xfb, 0x2c,
	0x24, 0x31, 0xe5, 0xf3, 0xbc, 0x8f, 0x3e, 0xdd, 0xa2, 0xbe, 0xd2, 0x9a, 0x27, 0x7d, 0x3a, 0xbf,
	0xbd, 0x7f, 0xbe, 0xc7, 0x02, 0x0c, 0x97, 0xde, 0x5e, 0x63, 0x81, 0x7e, 0x7b, 0x7f, 0xe2, 0x23,
	0x1d, 0x2a, 0xba, 0x49, 0x5b, 0x5a, 0x58, 0xc8, 0x35, 0x7b, 0xc8, 0xbb, 0x0b, 0x96, 0xf6, 0x82,
	0x6e, 0xe2, 0x85, 0x0e, 0x5b, 0x20, 0x7d, 0xba, 0xb0, 0xbd, 0x7f, 0x41, 0xd9, 0x69, 0xbf, 0x55,
	0xfd, 0xc5, 0xf6, 0xe1, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0x9e, 0xf0, 0x22, 0xba, 0x70, 0x13,
	0x00, 0x00,
}
