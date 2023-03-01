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

 Date: 01/03/2023 16:21:08
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for uc_system_master_visitor_logs
-- ----------------------------
DROP TABLE IF EXISTS `uc_system_master_visitor_logs`;
CREATE TABLE `uc_system_master_visitor_logs`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `account_id` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '关联用户ID',
  `os_category` tinyint(4) UNSIGNED NULL DEFAULT NULL COMMENT '系统访问类别（1=web端，2=android端，3=IOS端）',
  `visit_category` smallint(5) UNSIGNED NULL DEFAULT NULL COMMENT '访问类型',
  `union_id` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '关联ID',
  `description` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '访问信息描述',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of uc_system_master_visitor_logs
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
