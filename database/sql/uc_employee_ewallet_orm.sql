/*
 Navicat MySQL Data Transfer

 Source Server         : test
 Source Server Type    : MySQL
 Source Server Version : 50738
 Source Host           : 82.157.248.230:3306
 Source Schema         : go_test

 Target Server Type    : MySQL
 Target Server Version : 50738
 File Encoding         : 65001

 Date: 29/03/2023 23:49:06
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for uc_employee_ewallet_orm
-- ----------------------------
DROP TABLE IF EXISTS `uc_employee_ewallet_orm`;
CREATE TABLE `uc_employee_ewallet_orm`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `account_id` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '关联登入ID',
  `union_id` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '关联钱包充值、消费ID',
  `union_type` tinyint(2) UNSIGNED NULL DEFAULT NULL COMMENT '关联类型（1.充值、0.消费）',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '员工钱包入账、消费ORM表' ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
