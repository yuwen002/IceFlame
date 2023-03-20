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

 Date: 20/03/2023 18:01:52
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for uc_system_permission_relation
-- ----------------------------
DROP TABLE IF EXISTS `uc_system_permission_relation`;
CREATE TABLE `uc_system_permission_relation`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `permission_id` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '权限ID',
  `role_id` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '角色ID',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '后台系统菜单列表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
