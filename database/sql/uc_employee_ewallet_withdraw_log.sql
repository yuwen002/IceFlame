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

 Date: 29/03/2023 23:49:12
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for uc_employee_ewallet_withdraw_log
-- ----------------------------
DROP TABLE IF EXISTS `uc_employee_ewallet_withdraw_log`;
CREATE TABLE `uc_employee_ewallet_withdraw_log`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `account_id` int(11) UNSIGNED NOT NULL COMMENT '关联员工ID',
  `money` int(11) UNSIGNED NULL DEFAULT NULL COMMENT '提现金额（单位：里）',
  `balance` int(11) NULL DEFAULT NULL COMMENT '提现前余额（单位：里）',
  `real_money` int(11) NULL DEFAULT 0 COMMENT '实际入账金额（单位：里）',
  `charge_fee` int(11) NULL DEFAULT 0 COMMENT '入账扣点金额（单位：里）',
  `withdrawal_status` tinyint(1) NULL DEFAULT 1 COMMENT '1=申请中2=已通过待到账3=已到账99=已驳回',
  `reject_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '驳回内容',
  `hash_withdraw` varchar(80) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '入账hash',
  `partner_trade_no` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '' COMMENT '提现到银行卡订单号',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '员工系统提现记录' ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
