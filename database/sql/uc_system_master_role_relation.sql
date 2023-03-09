/*
 Navicat MySQL Data Transfer

 Source Server         : kd_test
 Source Server Type    : MySQL
 Source Server Version : 50738
 Source Host           : 82.157.248.230:3306
 Source Schema         : go_test

 Target Server Type    : MySQL
 Target Server Version : 50738
 File Encoding         : 65001

 Date: 09/03/2023 15:31:42
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for uc_system_master_role_relation
-- ----------------------------
DROP TABLE IF EXISTS `uc_system_master_role_relation`;
CREATE TABLE `uc_system_master_role_relation`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `account_id` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '管理员ID',
  `role_id` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '身份ID',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '角色信息关联管理ID' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of uc_system_master_role_relation
-- ----------------------------
INSERT INTO `uc_system_master_role_relation` VALUES (1, 1, 1, '2023-03-09 13:27:27', '2023-03-09 14:26:49');


SET FOREIGN_KEY_CHECKS = 1;
