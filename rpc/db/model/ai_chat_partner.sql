DROP TABLE IF EXISTS `ac_user`;
CREATE TABLE `ac_user`  (
  `uin` int(255) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '唯一id',
  `role` int(10) NOT NULL COMMENT '用户类型',
  `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `nickname` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '昵称',
  `sex` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '性别',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '上次更新数据时间',
  `version` int(30) UNSIGNED NOT NULL DEFAULT 1 COMMENT '版本号',
  PRIMARY KEY (`uin`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

CREATE TRIGGER before_insert_ac_user
BEFORE INSERT ON ac_user
FOR EACH ROW
BEGIN
    SET NEW.version = (SELECT IFNULL(MAX(version), 0) + 1 FROM ac_user);
END