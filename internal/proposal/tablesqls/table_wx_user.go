package tablesqls

//CREATE TABLE `admin` (
//`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
//`username` varchar(32) NOT NULL DEFAULT '' COMMENT '用户名',
//`password` varchar(100) NOT NULL DEFAULT '' COMMENT '密码',
//`nickname` varchar(60) NOT NULL DEFAULT '' COMMENT '昵称',
//`mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
//`is_used` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否启用 1:是  -1:否',
//`is_deleted` tinyint(1) NOT NULL DEFAULT '-1' COMMENT '是否删除 1:是  -1:否',
//`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
//`created_user` varchar(60) NOT NULL DEFAULT '' COMMENT '创建人',
//`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
//`updated_user` varchar(60) NOT NULL DEFAULT '' COMMENT '更新人',
//PRIMARY KEY (`id`),
//UNIQUE KEY `unique_username` (`username`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='管理员表';

func CreateWxUserTableSql() (sql string) {
	sql = "CREATE TABLE `wx_user` ("
	sql += "`id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',"
	sql += "`openid` varchar(60) NOT NULL DEFAULT '' COMMENT '微信openid',"
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
