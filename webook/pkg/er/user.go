package er

const (
	UserExist                        ErrCode = 1001 //用户已存在
	UserInvalidInput                 ErrCode = 1003 //输入错误
	UserAuthFailed                   ErrCode = 1004 //账号密码错误
	UserOperationTooFrequent         ErrCode = 1005 //操作太频繁
	Code_NotFind                     ErrCode = 2001 //验证码不存在,请重新发送
	Code_VerifyFail                  ErrCode = 2002 //验证失败
	Code_TooManyVerificationAttempts ErrCode = 2003 //验证次数过多,请重新发送
)
