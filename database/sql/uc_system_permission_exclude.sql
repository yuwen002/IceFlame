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

 Date: 24/03/2023 17:43:58
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for uc_system_permission_exclude
-- ----------------------------
DROP TABLE IF EXISTS `uc_system_permission_exclude`;
CREATE TABLE `uc_system_permission_exclude`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '菜单名称',
  `module` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '对应的程序模块',
  `uri` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '对应的URI地址',
  `remark` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '权限排除，公共模块，任何管理员都可以访问' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of uc_system_permission_exclude
-- ----------------------------
INSERT INTO `uc_system_permission_exclude` VALUES (1, '管理员修改密码', 'manage.UcSystemMaster.EditPassword', 'manage/master/edit_password', '修改11', '2023-03-19 00:13:53', '2023-03-19 00:19:43');
INSERT INTO `uc_system_permission_exclude` VALUES (2, '所有频道', 'manage.Article.GetArticleChannelAll', 'manage/article/channel/get_all', NULL, '2023-03-24 17:28:16', '2023-03-24 17:28:16');
INSERT INTO `uc_system_permission_exclude` VALUES (3, '所有分类', 'manage.Article.GetArticleCategoryAll', 'manage/article/category/get_all', NULL, '2023-03-24 17:36:53', '2023-03-24 17:36:53');

SET FOREIGN_KEY_CHECKS = 1;
