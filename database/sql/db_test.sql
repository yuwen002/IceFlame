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

 Date: 25/02/2023 16:42:02
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for db_test
-- ----------------------------
DROP TABLE IF EXISTS `db_test`;
CREATE TABLE `db_test`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `test_data` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '数据',
  `title` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '标题',
  `classify` int(5) NULL DEFAULT 1,
  `num` int(11) NULL DEFAULT NULL,
  `union_id` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
