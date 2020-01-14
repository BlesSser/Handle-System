package model

//////////////////////////////////////////////////////
//operator code
const OC_RESERVED OpCode = 0 //Reserved
const OC_RESOLUTION OpCode = 1
const OC_GET_SITEINFO OpCode = 2

const OC_CREATE_HANDLE OpCode = 100
const OC_DELETE_HANDLE OpCode = 101
const OC_ADD_VALUE OpCode = 102
const OC_REMOVE_VALUE OpCode = 103
const OC_MODIFY_VALUE OpCode = 104
const OC_LIST_HANDLE OpCode = 105
const OC_LIST_NA OpCode = 106

const OC_CHALLENGE_RESPONSE OpCode = 200
const OC_VERIFY_RESPONSE OpCode = 201

const OC_SESSION_SETUP OpCode = 2
const OC_SESSION_TERMINATE OpCode = 2
const OC_SESSION_EXCHANGEKEY OpCode = 2

/////////////////////////////////////////////////////
//response code
const RC_RESERVED RespCode = 0
const RC_SUCCESS RespCode = 1
const RC_ERROR RespCode = 2
const RC_SERVER_BUSY RespCode = 3
const RC_PROTOCOL_ERROR RespCode = 4
const RC_OPERATION_DENIED RespCode = 5
const RC_RECUR_LIMIT_EXCEEDED RespCode = 6

const RC_HANDLE_NOT_FOUND RespCode = 100
const RC_HANDLE_ALREADY_EXIST RespCode = 101
const RC_INVALID_HANDLE RespCode = 102

const RC_VALUE_NOT_FOUND RespCode = 200
const RC_VALUE_ALREADY_EXIST RespCode = 201
const RC_VALUE_INVALID RespCode = 202

const RC_EXPIRED_SITE_INFO RespCode = 300
const RC_SERVER_NOT_RESP RespCode = 301
const RC_SERVICE_REFERRAL RespCode = 302
const RC_NA_DELEGATE RespCode = 303

const RC_NOT_AUTHORIZED RespCode = 400
const RC_ACCESS_DENIED RespCode = 401
const RC_AUTHEN_NEEDED RespCode = 402
const RC_AUTHEN_FAILED RespCode = 403
const RC_INVALID_CREDENTIAL RespCode = 404
const RC_AUTHEN_TIMEOUT RespCode = 405
const RC_UNABLE_TO_AUTHEN RespCode = 406

const RC_SESSION_TIMEOUT RespCode = 500
const RC_SESSION_FAILED RespCode = 501
const RC_NO_SESSION_KEY RespCode = 502
const RC_SESSION_NO_SUPPORT RespCode = 503
const RC_SESSION_KEY_INVALID RespCode = 504

const RC_TRYING RespCode = 900
const RC_FORWARDED RespCode = 901
const RC_QUEUED RespCode = 902
