package errno

const (
	Success                                   = 10000
	ErrParam                                  = 10001
	ErrServer                                 = 10002
	ErrNonce                                  = 10003
	ErrTimeStamp                              = 10004
	ErrRPCFailed                              = 10005
	ErrInvalidToken                           = 10006
	ErrMarshalFailed                          = 10007
	ErrUnMarshalFailed                        = 10008
	ErrMustDID                                = 10011
	ErrMustSN                                 = 10012
	ErrHttpFailed                             = 10013
	ErrRedisFailed                            = 10100
	ErrMongoFailed                            = 10101
	ErrMysqlFailed                            = 10102
	ErrRecordNotFound                         = 10103
	ErrSignError                              = 20001
	ErrRepeatRequest                          = 20002
	ErrMustLogin                              = 20003
	ErrAuthFailed                             = 20004
	ErrYetRegister                            = 20005
	ErrURLExpired                             = 20006
	ErrExistsTeam                             = 20007
	ErrMustTaskInit                           = 20008
	ErrResourceNotEnough                      = 20009
	ErrEmptyScene                             = 20010
	ErrYetPreinstall                          = 20011
	ErrReportNotFound                         = 20012
	ErrInviteCodeFailed                       = 20013
	ErrDefaultTeamFailed                      = 20014
	ErrRecordExists                           = 20015
	ErrEmptyTestCase                          = 20016
	ErrSceneCaseNameIsExist                   = 20017
	ErrApiNameAlreadyExist                    = 20018
	ErrGroupNameAlreadyExist                  = 20019
	ErrFolderNameAlreadyExist                 = 20020
	ErrSceneNameAlreadyExist                  = 20021
	ErrPlanNameAlreadyExist                   = 20022
	ErrEnvNameIsExist                         = 20023
	ErrReportInRun                            = 20024
	ErrMobileYetRegister                      = 20025
	ErrSmsCodeSendIllegal                     = 20026
	ErrSmsCodeVerifyFail                      = 20027
	ErrTeamUserNumFail                        = 20028
	ErrNotFundAvailTeam                       = 20029
	ErrAuthFailedNotRegistered                = 20030
	ErrSmsCodeSend                            = 20031
	ErrTeamOverdue                            = 20032
	ErrInvoiceAlreadyOpen                     = 20033
	ErrTeamUserOvertopLimit                   = 20034
	ErrConcurrenceLimit                       = 20035
	ErrDurationOvertopLimit                   = 20036
	ErrRoundNumOvertopLimit                   = 20037
	ErrStartConcurrencyOvertopLimit           = 20038
	ErrStepOvertopLimit                       = 20039
	ErrStepRunTimeOvertopLimit                = 20040
	ErrMaxConcurrencyOvertopLimit             = 20041
	ErrPreinstallNameIsExist                  = 20043
	ErrApiNumOvertopLimit                     = 20044
	ErrAddEmailUserNumOvertopLimit            = 20045
	ErrInviteUserNumOvertopLimit              = 20046
	ErrMachineMonitorDataPastDue              = 20047
	ErrInPlanSceneNameAlreadyExist            = 20048
	ErrPlanNameNotEmpty                       = 20049
	ErrInPlanGroupNameAlreadyExist            = 20050
	ErrVerifyFail                             = 20051
	ErrTimedTaskOverdue                       = 20052
	ErrTimedTaskTimeGtTeamTime                = 20053
	ErrWechatLoginQrCodeOverdue               = 20054
	ErrCannotDeleteRunningPlan                = 20055
	ErrCannotBatchDeleteRunningPlan           = 20056
	ErrWaitControllerOvertopLimit             = 20057
	ErrMaxConcurrencyLessThanStartConcurrency = 20058
	ErrVumNotEnough                           = 20060
	ErrEmptySceneFlow                         = 20061
	ErrEmptyTestCaseFlow                      = 20062
	ErrNameOverLength                         = 20063
	ErrCartEmpty                              = 20064
	ErrCouponDoNotCondition                   = 20065
	ErrPayFail                                = 20066
)

// CodeAlertMap 错图码映射错误提示，展示给用户
var CodeAlertMap = map[int]string{
	Success:                                   "成功",
	ErrServer:                                 "服务器错误",
	ErrParam:                                  "参数校验错误",
	ErrSignError:                              "签名错误",
	ErrRepeatRequest:                          "重放请求",
	ErrNonce:                                  "_nonce参数错误",
	ErrTimeStamp:                              "_timestamp参数错误",
	ErrRecordNotFound:                         "数据库记录不存在",
	ErrRPCFailed:                              "请求下游服务失败",
	ErrInvalidToken:                           "无效的token",
	ErrMarshalFailed:                          "序列化失败",
	ErrUnMarshalFailed:                        "反序列化失败",
	ErrRedisFailed:                            "redis操作失败",
	ErrMongoFailed:                            "mongo操作失败",
	ErrMysqlFailed:                            "mysql操作失败",
	ErrMustLogin:                              "没有获取到登录态",
	ErrMustDID:                                "缺少设备DID信息",
	ErrMustSN:                                 "缺少设备SN信息",
	ErrHttpFailed:                             "请求下游Http服务失败",
	ErrAuthFailed:                             "用户名或密码错误",
	ErrYetRegister:                            "用户邮箱已注册",
	ErrURLExpired:                             "邀请链接已过期",
	ErrExistsTeam:                             "用户已在此团队",
	ErrMustTaskInit:                           "请填写任务配置并保存",
	ErrResourceNotEnough:                      "资源不足",
	ErrEmptyScene:                             "场景不能为空",
	ErrYetPreinstall:                          "预设配置名称已存在",
	ErrReportNotFound:                         "报告不存在",
	ErrInviteCodeFailed:                       "邀请码验证失败",
	ErrDefaultTeamFailed:                      "当前默认团队错误",
	ErrRecordExists:                           "数据库记录已存在",
	ErrEmptyTestCase:                          "场景用例不能为空",
	ErrSceneCaseNameIsExist:                   "同一场景下用例名称不能重复",
	ErrApiNameAlreadyExist:                    "接口名称已存在",
	ErrGroupNameAlreadyExist:                  "分组名称已存在",
	ErrFolderNameAlreadyExist:                 "文件夹名称已存在",
	ErrSceneNameAlreadyExist:                  "场景名称已存在",
	ErrPlanNameAlreadyExist:                   "计划名称已存在",
	ErrEnvNameIsExist:                         "环境名称已存在",
	ErrReportInRun:                            "报告数据正在生成中，请稍后再查看",
	ErrMobileYetRegister:                      "手机号已注册",
	ErrSmsCodeSendIllegal:                     "验证码发送不合法",
	ErrSmsCodeVerifyFail:                      "验证码不正确",
	ErrTeamUserNumFail:                        "抱歉，此团队购买人数不在范围之内",
	ErrNotFundAvailTeam:                       "没有找到可用的团队",
	ErrAuthFailedNotRegistered:                "账号未注册",
	ErrSmsCodeSend:                            "",
	ErrTeamOverdue:                            "该团队已过期，请及时续费",
	ErrInvoiceAlreadyOpen:                     "发票不能重复开具，如需重新开具，请联系客服",
	ErrTeamUserOvertopLimit:                   "团队人数已经超过上限",
	ErrConcurrenceLimit:                       "并发数超过上限",
	ErrDurationOvertopLimit:                   "持续时长超过上限",
	ErrRoundNumOvertopLimit:                   "运行轮次超过上限",
	ErrStartConcurrencyOvertopLimit:           "不可超过最大并发数",
	ErrStepOvertopLimit:                       "不可超过最大并发数",
	ErrStepRunTimeOvertopLimit:                "已超出该团队套餐的最大时长",
	ErrMaxConcurrencyOvertopLimit:             "已超出该团队套餐的最大限额",
	ErrPreinstallNameIsExist:                  "预设配置名称已存在",
	ErrApiNumOvertopLimit:                     "单场景接口数已达该团队套餐的最大限额",
	ErrAddEmailUserNumOvertopLimit:            "单次只可添加1-50个收件人进行发送",
	ErrInviteUserNumOvertopLimit:              "受邀人数已超出团队成员限额，请升级团队以完成邀请",
	ErrMachineMonitorDataPastDue:              "只能查询15天以内的压力机监控数据",
	ErrInPlanSceneNameAlreadyExist:            "计划内场景不可重名",
	ErrPlanNameNotEmpty:                       "计划名称不能为空",
	ErrInPlanGroupNameAlreadyExist:            "计划内分组不可重名",
	ErrVerifyFail:                             "验证失败",
	ErrTimedTaskOverdue:                       "开始或结束时间不能早于当前时间",
	ErrTimedTaskTimeGtTeamTime:                "定时任务时间不可晚于团队过期时间",
	ErrWechatLoginQrCodeOverdue:               "当前微信二维码过期",
	ErrCannotDeleteRunningPlan:                "该计划正在运行，无法删除",
	ErrCannotBatchDeleteRunningPlan:           "存在运行中的计划，无法删除",
	ErrWaitControllerOvertopLimit:             "等待控制器时间不能超过20000毫秒",
	ErrMaxConcurrencyLessThanStartConcurrency: "最大并发数不能小于起始并发数",
	ErrVumNotEnough:                           "VUM余额不足",
	ErrEmptySceneFlow:                         "场景flow不能为空",
	ErrEmptyTestCaseFlow:                      "场景用例flow不能为空",
	ErrNameOverLength:                         "名称过长！不可超出30字符",
	ErrCartEmpty:                              "购物车已空",
	ErrCouponDoNotCondition:                   "优惠券不满足使用条件",
	ErrPayFail:                                "支付失败",
}

// CodeMsgMap 错误码映射错误信息，不展示给用户
var CodeMsgMap = map[int]string{
	Success:                                   "success",
	ErrServer:                                 "internal server error",
	ErrParam:                                  "param error",
	ErrSignError:                              "signature error",
	ErrRepeatRequest:                          "repeat request",
	ErrNonce:                                  "nonce error",
	ErrTimeStamp:                              "timestamp error",
	ErrRecordNotFound:                         "record not found",
	ErrRPCFailed:                              "rpc failed",
	ErrInvalidToken:                           "invalid token",
	ErrMarshalFailed:                          "marshal failed",
	ErrUnMarshalFailed:                        "unmarshal failed",
	ErrRedisFailed:                            "redis operate failed",
	ErrMongoFailed:                            "mongo operate failed",
	ErrMysqlFailed:                            "mysql operate failed",
	ErrMustLogin:                              "must login",
	ErrMustDID:                                "must DID",
	ErrMustSN:                                 "must SN",
	ErrHttpFailed:                             "http failed",
	ErrAuthFailed:                             "username/password failed",
	ErrYetRegister:                            "email yet register",
	ErrURLExpired:                             "invite url expired",
	ErrExistsTeam:                             "invite user exists team",
	ErrMustTaskInit:                           "fill in the task allocation and save it",
	ErrResourceNotEnough:                      "resource not enough",
	ErrEmptyScene:                             "the scene cannot be empty",
	ErrYetPreinstall:                          "preinstall yet exists",
	ErrReportNotFound:                         "report not found",
	ErrInviteCodeFailed:                       "invite code failed",
	ErrDefaultTeamFailed:                      "default team failed",
	ErrRecordExists:                           "record exists",
	ErrEmptyTestCase:                          "scenario cases cannot be empty",
	ErrSceneCaseNameIsExist:                   "scene case name is exist",
	ErrApiNameAlreadyExist:                    "api name already exist",
	ErrGroupNameAlreadyExist:                  "group name already exist",
	ErrFolderNameAlreadyExist:                 "folder name already exist",
	ErrSceneNameAlreadyExist:                  "scene name already exist",
	ErrPlanNameAlreadyExist:                   "plan name already exist",
	ErrEnvNameIsExist:                         "environment name is exist",
	ErrReportInRun:                            "report in run",
	ErrMobileYetRegister:                      "mobile yet register",
	ErrSmsCodeSendIllegal:                     "ErrSmsCodeSendIllegal",
	ErrSmsCodeVerifyFail:                      "ErrSmsCodeVerifyFail",
	ErrTeamUserNumFail:                        "team user num fail",
	ErrNotFundAvailTeam:                       "not fund available team",
	ErrAuthFailedNotRegistered:                "account not registered",
	ErrSmsCodeSend:                            "ErrSmsCodeSend",
	ErrTeamOverdue:                            "ErrTeamOverdue",
	ErrInvoiceAlreadyOpen:                     "invoice already open",
	ErrTeamUserOvertopLimit:                   "team user overtop limit",
	ErrConcurrenceLimit:                       "ErrConcurrenceLimit",
	ErrDurationOvertopLimit:                   "ErrDurationOvertopLimit",
	ErrRoundNumOvertopLimit:                   "ErrRoundNumOvertopLimit",
	ErrStartConcurrencyOvertopLimit:           "ErrStartConcurrencyOvertopLimit",
	ErrStepOvertopLimit:                       "ErrStepOvertopLimit",
	ErrStepRunTimeOvertopLimit:                "ErrStepRunTimeOvertopLimit",
	ErrMaxConcurrencyOvertopLimit:             "ErrMaxConcurrencyOvertopLimit",
	ErrPreinstallNameIsExist:                  "preinstall name is exist",
	ErrApiNumOvertopLimit:                     "api num overtop limit",
	ErrAddEmailUserNumOvertopLimit:            "ErrAddEmailUserNumOvertopLimit",
	ErrInviteUserNumOvertopLimit:              "ErrInviteUserNumOvertopLimit",
	ErrMachineMonitorDataPastDue:              "ErrMachineMonitorDataPastDue",
	ErrInPlanSceneNameAlreadyExist:            "ErrInPlanSceneNameAlreadyExist",
	ErrPlanNameNotEmpty:                       "ErrPlanNameNotEmpty",
	ErrInPlanGroupNameAlreadyExist:            "ErrInPlanGroupNameAlreadyExist",
	ErrVerifyFail:                             "ErrVerifyFail",
	ErrTimedTaskOverdue:                       "ErrTimedTaskOverdue",
	ErrTimedTaskTimeGtTeamTime:                "ErrTimedTaskTimeGtTeamTime",
	ErrWechatLoginQrCodeOverdue:               "ErrWechatLoginQrCodeOverdue",
	ErrCannotDeleteRunningPlan:                "ErrCannotDeleteRunningPlan",
	ErrCannotBatchDeleteRunningPlan:           "ErrCannotBatchDeleteRunningPlan",
	ErrWaitControllerOvertopLimit:             "ErrWaitControllerOvertopLimit",
	ErrMaxConcurrencyLessThanStartConcurrency: "ErrMaxConcurrencyLessThanStartConcurrency",
	ErrVumNotEnough:                           "ErrVumNotEnough",
	ErrEmptySceneFlow:                         "ErrEmptySceneFlow",
	ErrEmptyTestCaseFlow:                      "ErrEmptyTestCaseFlow",
	ErrNameOverLength:                         "ErrNameOverLength",
	ErrCartEmpty:                              "ErrCartEmpty",
	ErrCouponDoNotCondition:                   "ErrCouponDoNotCondition",
	ErrPayFail:                                "ErrPayFail",
}
