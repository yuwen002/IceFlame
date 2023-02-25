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

 Date: 25/02/2023 16:39:35
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for uc_account
-- ----------------------------
DROP TABLE IF EXISTS `uc_account`;
CREATE TABLE `uc_account`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `identity_card_id` int(11) UNSIGNED NULL DEFAULT 0 COMMENT '关联用户身份证信息ID',
  `username` char(40) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户名',
  `password_hash` char(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户密码',
  `tel` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户手机号码',
  `status` tinyint(1) NULL DEFAULT 0 COMMENT '用户状态(0.启用1.停用)',
  `real_name_type` tinyint(1) NULL DEFAULT 1 COMMENT '实名状态1=未实名2=已上传未审核3=审核驳回4=实名成功',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `identity_card_id`(`identity_card_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户中心' ROW_FORMAT = DYNAMIC;

SET FOREIGN_KEY_CHECKS = 1;
