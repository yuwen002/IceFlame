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

 Date: 24/03/2023 17:44:35
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article_channel
-- ----------------------------
DROP TABLE IF EXISTS `article_channel`;
CREATE TABLE `article_channel`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '栏目名称',
  `remark` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  `sort` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '排序顺序',
  `status` tinyint(3) UNSIGNED NULL DEFAULT 0 COMMENT '显示状态（0=显示，1=隐藏）',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = 'Summary 频道分类' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of article_channel
-- ----------------------------
INSERT INTO `article_channel` VALUES (1, '测试频道', '测试频道修改', 1, 0, '2023-03-24 13:40:42', '2023-03-24 14:45:13');
INSERT INTO `article_channel` VALUES (2, '频道2', NULL, NULL, 0, '2023-03-24 17:25:30', '2023-03-24 17:25:30');

SET FOREIGN_KEY_CHECKS = 1;
