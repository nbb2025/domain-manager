// 错误码枚举
export const ErrorCode = {
  Success: 0,
  InvalidParams: 1001,
  Unauthorized: 1002,
  Forbidden: 1003,
  NotFound: 1004,
  InternalError: 1005,
  InvalidPassword: 1006,
  UserExists: 1007,
  DomainExists: 1008,
  ProviderKeyError: 1009
}

// 错误信息映射
export const ErrorMessage = {
  [ErrorCode.Success]: '操作成功',
  [ErrorCode.InvalidParams]: '请求参数无效',
  [ErrorCode.Unauthorized]: '未授权访问',
  [ErrorCode.Forbidden]: '禁止访问',
  [ErrorCode.NotFound]: '资源不存在',
  [ErrorCode.InternalError]: '服务器内部错误',
  [ErrorCode.InvalidPassword]: '密码错误',
  [ErrorCode.UserExists]: '用户已存在',
  [ErrorCode.DomainExists]: '域名已存在',
  [ErrorCode.ProviderKeyError]: '服务商密钥错误'
}

// 获取错误信息
export const getErrorMessage = (code) => {
  return ErrorMessage[code] || '未知错误'
}