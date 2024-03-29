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

 Date: 24/03/2023 17:44:26
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article_category
-- ----------------------------
DROP TABLE IF EXISTS `article_category`;
CREATE TABLE `article_category`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `fid` int(10) UNSIGNED NULL DEFAULT NULL COMMENT '父级ID',
  `name` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '分类名称',
  `remark` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  `sort` int(255) UNSIGNED NULL DEFAULT NULL COMMENT '排序顺序',
  `status` tinyint(3) UNSIGNED NULL DEFAULT 0 COMMENT '显示状态（0=显示，1=隐藏）',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of article_category
-- ----------------------------
INSERT INTO `article_category` VALUES (1, 0, '测试分类', '测试分类修改11', 1, 0, '2023-03-24 16:44:09', '2023-03-24 16:54:41');

SET FOREIGN_KEY_CHECKS = 1;
