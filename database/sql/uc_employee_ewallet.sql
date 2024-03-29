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

 Date: 02/04/2023 22:32:11
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for uc_employee_ewallet
-- ----------------------------
DROP TABLE IF EXISTS `uc_employee_ewallet`;
CREATE TABLE `uc_employee_ewallet`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `account_id` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '关联员工ID',
  `money` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '账户余额（单位：里）（冻结金额）',
  `balance` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '上次入账余额（单位：里）',
  `total_money` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '账户累计金额（单位：里）',
  `unfreeze_amount` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '解冻金额 ',
  `hash_ewallet` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'hash钱包',
  `wx_openid` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '微信提现openid',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `account_id`(`account_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '员工钱包' ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
