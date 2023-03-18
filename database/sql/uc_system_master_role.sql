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

 Date: 18/03/2023 16:28:50
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for uc_system_master_role
-- ----------------------------
DROP TABLE IF EXISTS `uc_system_master_role`;
CREATE TABLE `uc_system_master_role`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '角色名称',
  `remark` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注',
  `supper_master` tinyint(4) NULL DEFAULT 0 COMMENT '1为超级管理员模块',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '管理员角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of uc_system_master_role
-- ----------------------------
INSERT INTO `uc_system_master_role` VALUES (1, '超级管理员', '222', NULL, '2023-03-07 16:53:22', '2023-03-16 15:41:16');
INSERT INTO `uc_system_master_role` VALUES (2, '普通管理员', '测试修改1', NULL, '2023-03-07 16:53:49', '2023-03-07 17:19:20');

SET FOREIGN_KEY_CHECKS = 1;
