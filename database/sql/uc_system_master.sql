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

 Date: 18/03/2023 16:28:23
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for uc_system_master
-- ----------------------------
DROP TABLE IF EXISTS `uc_system_master`;
CREATE TABLE `uc_system_master`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `account_id` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '关联account id',
  `name` varchar(32) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户姓名',
  `tel` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户电话',
  `supper_master` tinyint(1) NULL DEFAULT 0 COMMENT '超级管理员(=1为超级管理员)',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '管理员用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of uc_system_master
-- ----------------------------
INSERT INTO `uc_system_master` VALUES (1, 1, '流星雨', '15566036902', 1, '2023-02-22 23:36:37', '2023-02-22 23:36:37');
INSERT INTO `uc_system_master` VALUES (2, 2, '测试修改33', '15566006661', 0, '2023-02-24 11:31:00', '2023-02-28 13:42:11');
INSERT INTO `uc_system_master` VALUES (3, 3, '流星雨11', '15566036904', 0, '2023-02-25 23:48:15', '2023-02-25 23:48:15');

SET FOREIGN_KEY_CHECKS = 1;
