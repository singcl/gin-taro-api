package tablesqls

// CREATE TABLE `weixin` (
// `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
// `openid` varchar(60) NOT NULL DEFAULT '' COMMENT '微信openid',
// `unionid` varchar(60) NOT NULL DEFAULT '' COMMENT '微信unionid',
// `session_key` varchar(60) NOT NULL DEFAULT '' COMMENT '微信session_key',
// `username` varchar(32) NOT NULL DEFAULT '' COMMENT '微信用户名',
// `nickname` varchar(60) NOT NULL DEFAULT '' COMMENT '昵称',
// `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
// `is_used` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否启用 1:是  -1:否',
// `is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT '是否删除 1:是  -1:否',
// `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
// `created_user` varchar(60) NOT NULL DEFAULT '' COMMENT '创建人',
// `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
// `updated_user` varchar(60) NOT NULL DEFAULT '' COMMENT '更新人',
// PRIMARY KEY (`id`, `openid`),
// UNIQUE KEY `unique_openid` (`openid`)
// ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='微信小程序登录用户表';

func CreateWxUserTableSql() (sql string) {
	sql = "CREATE TABLE `weixin` ("
	sql += "`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',"
	sql += "`openid` varchar(60) NOT NULL DEFAULT '' COMMENT '微信openid',"
	sql += "`unionid` varchar(60) NOT NULL DEFAULT '' COMMENT '微信unionid',"
	sql += "`session_key` varchar(60) NOT NULL DEFAULT '' COMMENT '微信session_key',"
	sql += "`username` varchar(32) NOT NULL DEFAULT '' COMMENT '微信用户名',"
	sql += "`nickname` varchar(60) NOT NULL DEFAULT '' COMMENT '昵称',"
	sql += "`mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',"
	sql += "`is_used` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否启用 1:是  -1:否',"
	sql += "`is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT '是否删除 1:是  -1:否',"
	sql += "`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',"
	sql += "`created_user` varchar(60) NOT NULL DEFAULT '' COMMENT '创建人',"
	sql += "`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',"
	sql += "`updated_user` varchar(60) NOT NULL DEFAULT '' COMMENT '更新人',"
	sql += "PRIMARY KEY (`id`, `openid`),"
	sql += "UNIQUE KEY `unique_openid` (`openid`)"
	sql += ") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='微信小程序登录用户表';"

	return
}
