/*
 Navicat MySQL Data Transfer

 Source Server         : VM-UbuntuServer-docker-mysql-latest
 Source Server Type    : MySQL
 Source Server Version : 90100
 Source Host           : 127.0.0.1:3306
 Source Schema         : ai_chat_partner

 Target Server Type    : MySQL
 Target Server Version : 90100
 File Encoding         : 65001

 Date: 29/12/2024 02:01:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for ac_user
-- ----------------------------
DROP TABLE IF EXISTS `ac_user`;
CREATE TABLE `ac_user`  (
  `uin` int UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '唯一id',
  `role` int NOT NULL COMMENT '用户类型',
  `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `nickname` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '昵称',
  `sex` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '性别',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '上次更新数据时间',
  `version` int UNSIGNED NULL DEFAULT 1 COMMENT '版本号',
  PRIMARY KEY (`uin`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Triggers structure for table ac_user
-- ----------------------------
DROP TRIGGER IF EXISTS `before_insert_ac_user`;
delimiter ;;
CREATE TRIGGER `before_insert_ac_user` BEFORE INSERT ON `ac_user` FOR EACH ROW BEGIN
    SET NEW.version = (SELECT IFNULL(MAX(version), 0) + 1 FROM ac_user);
END
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table ac_user
-- ----------------------------
DROP TRIGGER IF EXISTS `before_update_ac_user_version`;
delimiter ;;
CREATE TRIGGER `before_update_ac_user_version` BEFORE UPDATE ON `ac_user` FOR EACH ROW BEGIN
    -- 在更新操作之前，将 NEW 行的 version 字段值增加 1
    SET NEW.version = OLD.version + 1;
END
;;
delimiter ;

SET FOREIGN_KEY_CHECKS = 1;
