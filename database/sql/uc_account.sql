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

 Date: 18/03/2023 16:26:52
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
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci COMMENT = '用户中心' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of uc_account
-- ----------------------------
INSERT INTO `uc_account` VALUES (1, 0, 'SA_15566036902', '$2a$10$CCHP.HVTh.MK.drUV4d3hOOIzfuRYms9QG/lEAruVhxJVSn9WF/Le', 'SA_15566036902', 0, 1, '2023-02-22 23:36:37', '2023-02-28 16:27:47');
INSERT INTO `uc_account` VALUES (2, 0, 'SA_yuwen002', '$2a$10$gWCqZgx6sEm4hHJEWQkaEuj3ukX3OFlpv1ZspqjWH9cQhDYTFT6YK', 'SA_15566006666', 0, 1, '2023-02-24 11:31:00', '2023-02-28 15:42:56');
INSERT INTO `uc_account` VALUES (3, 0, 'SA_yuwen003', '$2a$10$UwsnYTpFIxPHPWm7ufGQr..kw0H7Q7BN3cPFb4wml./AF48fbJJ1K', 'SA_15566036904', 0, 1, '2023-02-25 23:48:15', '2023-02-25 23:48:15');

SET FOREIGN_KEY_CHECKS = 1;
