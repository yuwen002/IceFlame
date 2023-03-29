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

 Date: 29/03/2023 23:49:00
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for uc_employee_ewallet_entry_log
-- ----------------------------
DROP TABLE IF EXISTS `uc_employee_ewallet_entry_log`;
CREATE TABLE `uc_employee_ewallet_entry_log`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `account_id` int(11) UNSIGNED NOT NULL COMMENT '关联员工ID',
  `store_id` int(11) UNSIGNED NULL DEFAULT NULL COMMENT '关联店铺ID',
  `separate_classify` tinyint(3) UNSIGNED NULL DEFAULT NULL COMMENT '入账类型（1线下扫码支付 2线上 3 店铺推广会员消费分佣 4 合伙人推广会员消费分佣 5配送费 6 店铺额外合伙人分销佣金 7 WiFi分佣）',
  `money` int(11) UNSIGNED NULL DEFAULT NULL COMMENT '入账金额（单位：里）',
  `balance` int(11) NULL DEFAULT NULL COMMENT '入账前余额（单位：里）',
  `entry_status` tinyint(3) UNSIGNED NULL DEFAULT 0 COMMENT '入账状态（99为提现驳回入账记录）',
  `order_id` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '关联订单ID',
  `hash_entry` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '入账hash',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '员工系统收入记账' ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
