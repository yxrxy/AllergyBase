namespace go model

// 基本响应结构
struct BaseResp {
    1: i64 code,    // 错误码，0表示成功
    2: string msg,  // 错误信息
}

// 用户模型
struct User {
    1: required i64 id,              // 用户ID
    2: required string username,     // 用户名
    4: optional string avatar,       // 头像URL
    6: optional i64 followCount,     // 关注数
    7: optional i64 followerCount,   // 粉丝数
    8: optional bool isFollow,       // 是否已关注
    9: optional i64 likeCount,       // 获赞数量
    10: optional i64 videoCount,     // 视频数量
}
