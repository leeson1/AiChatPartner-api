/*
 Navicat MySQL Data Transfer

 Source Server         : VM-UbuntuServer_docker_mysql5.7
 Source Server Type    : MySQL
 Source Server Version : 50743
 Source Host           : 175.178.219.9:3306
 Source Schema         : ai_chat_partner

 Target Server Type    : MySQL
 Target Server Version : 50743
 File Encoding         : 65001

 Date: 24/11/2024 14:54:00
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for ac_user
-- ----------------------------
DROP TABLE IF EXISTS `ac_user`;
CREATE TABLE `ac_user`  (
  `uin` int(255) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '唯一id',
  `role` int(10) NOT NULL COMMENT '用户类型',
  `username` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '邮箱',
  `nickname` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '昵称',
  `sex` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '性别',
  `create_time` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '注册时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '上次更新数据时间',
  `version` int(30) UNSIGNED NULL DEFAULT NULL COMMENT '版本号',
  PRIMARY KEY (`uin`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
