# 彩虹聚合登录 SDK Go 版本

这是彩虹聚合登录 SDK 的 Go 语言实现版本，支持多种第三方平台的 OAuth 登录。

## 项目结构

```
聚合登录SDK/
└── sdk/
    ├── oauth.go     # OAuth 核心实现
    ├── config.go    # 配置定义
    └── go.mod       # SDK 模块定义
```

## SDK 使用

### 安装

```bash
go get oauth-sdk
```

### 基本用法

```go
import oauthsdk "oauth-sdk"

// 初始化
oauthClient := oauthsdk.NewOauth(
    "https://u.cccyun.cc/",  // API 地址
    "你的APPID",
    "你的APPKEY",
    "http://127.0.0.1:8080/connect",  // 回调地址
)

// 获取登录跳转 URL
resp, err := oauthClient.Login("qq")
if err != nil {
    // 处理错误
}
// resp.URL 即为登录跳转地址
// resp.QRCode (微信和支付宝会返回二维码地址)

// 回调处理
user, err := oauthClient.Callback(code)
if err != nil {
    // 处理错误
}
// user.SocialUID - 第三方登录 UID
// user.AccessToken - 访问令牌
// user.Nickname - 用户昵称
// user.FaceImg - 用户头像
// user.Gender - 用户性别
// user.Location - 用户位置

// 查询用户信息
userInfo, err := oauthClient.Query("qq", "social_uid")
if err != nil {
    // 处理错误
}
```

## 支持的登录方式

| 值 | 登录方式 |
|----|---------|
| qq | QQ |
| wx | 微信 |
| alipay | 支付宝 |
| sina | 新浪微博 |
| baidu | 百度 |
| douyin | 抖音 |
| huawei | 华为 |
| xiaomi | 小米 |
| google | 谷歌 |
| microsoft | 微软 |
| facebook | Facebook |
| twitter | Twitter |
| feishu | 飞书 |
| wework | 企业微信 |
| dingtalk | 钉钉 |
| gitee | Gitee |
| github | GitHub |

## API 方法

### NewOauth
创建 OAuth 客户端实例。

```go
func NewOauth(apiURL, appID, appKey, callback string) *Oauth
```

### Login
获取登录跳转地址。

```go
func (o *Oauth) Login(loginType string) (*LoginResponse, error)
```

返回值 `LoginResponse`:
- `Code`: 状态码（0 表示成功）
- `Msg`: 返回信息
- `Type`: 登录方式
- `URL`: 登录跳转地址
- `QRCode`: 登录扫码地址（微信和支付宝返回）

### Callback
通过授权码获取用户信息。

```go
func (o *Oauth) Callback(code string) (*CallbackResponse, error)
```

返回值 `CallbackResponse`:
- `Code`: 状态码（0 表示成功）
- `Msg`: 返回信息
- `Type`: 登录方式
- `SocialUID`: 第三方登录 UID
- `AccessToken`: 访问令牌
- `Nickname`: 用户昵称
- `FaceImg`: 用户头像
- `Gender`: 用户性别
- `Location`: 用户位置
- `IP`: 用户 IP

### Query
查询用户信息。

```go
func (o *Oauth) Query(loginType, socialUID string) (*QueryResponse, error)
```

返回值 `QueryResponse`:
- `Code`: 状态码（0 表示成功）
- `Msg`: 返回信息
- `Type`: 登录方式
- `SocialUID`: 第三方登录 UID
- `AccessToken`: 访问令牌
- `Nickname`: 用户昵称
- `FaceImg`: 用户头像
- `Gender`: 用户性别
- `Location`: 用户位置
- `IP`: 用户 IP

## 注意事项

1. 请确保回调地址配置正确，否则无法正常登录
2. 建议使用 HTTPS 协议
3. 请妥善保管 AppID 和 AppKey
4. 登录流程需要正确处理 state 参数以防止 CSRF 攻击

## 许可证

本项目基于 PHP 版本的彩虹聚合登录 SDK 进行改写。
