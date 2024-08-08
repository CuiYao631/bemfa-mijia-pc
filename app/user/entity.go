package user

const (
	// EmailRegister 邮箱注册
	EmailRegister = "/v1/register"
	// EmailLogin 邮箱登录
	EmailLogin = "/v1/login"
	// ModifyPassword 修改密码
	ModifyPassword = "/v1/modify/password"
	// PhoneRegister 手机注册
	PhoneRegister = "/vb/api/v1/userRegister"
	// PhoneLogin 手机登录
	PhoneLogin = "/vb/api/v1/userLogin"
	// SetNewAppIDAndSecretKey 设置新的AppID和secretKey
	SetNewAppIDAndSecretKey = "/vs/web/v1/userSecretKey"
	// GetAllAppID 获取所有appID
	GetAllAppID = "/vs/web/v1/userAppID"
	// DeleteAppID 删除appID
	DeleteAppID = "/vs/web/v1/userDeleteKey"
)
