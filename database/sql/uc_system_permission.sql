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

 Date: 15/03/2023 17:59:56
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for uc_system_permission
-- ----------------------------
DROP TABLE IF EXISTS `uc_system_permission`;
CREATE TABLE `uc_system_permission`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `fid` int(10) UNSIGNED NULL DEFAULT 0 COMMENT '父级ID',
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '菜单名称',
  `module` varchar(80) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '对应的程序模块',
  `type` tinyint(4) NULL DEFAULT NULL COMMENT '类型（1=菜单，2=按钮）',
  `status` tinyint(3) UNSIGNED NULL DEFAULT 0 COMMENT '程序模块状态（0=启用，1=停用）',
  `remark` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '后台系统权限模块列表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of uc_system_permission
-- ----------------------------
INSERT INTO `uc_system_permission` VALUES (1, 0, '权限管理', NULL, 1, 0, '一级导航', '2023-03-13 11:19:24', '2023-03-13 11:19:24');
INSERT INTO `uc_system_permission` VALUES (2, 1, '管理员角色列表', 'manage.UcSystemMasterAuth.ListRole', 1, 0, '二级导航', '2023-03-13 11:45:46', '2023-03-13 11:45:46');
INSERT INTO `uc_system_permission` VALUES (3, 1, '添加管理员角色', 'manage.UcSystemMasterAuth.AddRole', 2, 0, '二级按钮', '2023-03-13 12:00:47', '2023-03-13 12:00:47');
INSERT INTO `uc_system_permission` VALUES (4, 1, '编辑管理员角色', 'manage.UcSystemMasterAuth.EditRole', 2, 0, '二级按钮', '2023-03-13 12:01:06', '2023-03-13 12:01:06');
INSERT INTO `uc_system_permission` VALUES (5, 1, '管理员角色绑定列表', 'manage.UcSystemMasterAuth.ListRoleRelation', 1, 0, '二级导航', '2023-03-13 13:46:48', '2023-03-13 13:46:48');
INSERT INTO `uc_system_permission` VALUES (6, 1, '添加管理员角色绑定', 'manage.UcSystemMasterAuth.AddRoleRelation', 2, 0, '二级按钮', '2023-03-13 13:47:30', '2023-03-13 13:47:30');
INSERT INTO `uc_system_permission` VALUES (7, 1, '编辑管理员角色绑定', 'manage.UcSystemMasterAuth.EditRoleRelation', 2, 0, '二级按钮', '2023-03-13 13:49:48', '2023-03-13 13:49:48');
INSERT INTO `uc_system_permission` VALUES (8, 1, '删除管理员角色绑定', 'manage.UcSystemMasterAuth.DeleteRoleRelation', 2, 0, '二级按钮', '2023-03-13 13:50:14', '2023-03-13 13:50:14');
INSERT INTO `uc_system_permission` VALUES (9, 1, '权限列表', 'manage.UcSystemMasterAuth.ListPermission', 1, 0, '二级菜单', '2023-03-13 14:30:56', '2023-03-13 14:30:56');
INSERT INTO `uc_system_permission` VALUES (10, 1, '添加权限', 'manage.UcSystemMasterAuth.AddPermission', 1, 0, '二级按钮', '2023-03-13 14:31:57', '2023-03-13 14:31:57');
INSERT INTO `uc_system_permission` VALUES (11, 1, '编辑权限', 'manage.UcSystemMasterAuth.EditPermission', 1, 0, '二级按钮', '2023-03-13 14:32:15', '2023-03-13 14:32:15');
INSERT INTO `uc_system_permission` VALUES (12, 1, '查看权限分配', 'manage.UcSystemMasterAuth.ListPermissionRelation', 2, 0, '二级按钮', '2023-03-15 17:57:35', '2023-03-15 17:57:35');
INSERT INTO `uc_system_permission` VALUES (13, 1, '编辑权限分配', 'manage.UcSystemMasterAuth.ListPermissionRelation', 2, 0, '二级按钮', '2023-03-15 17:58:59', '2023-03-15 17:58:59');

SET FOREIGN_KEY_CHECKS = 1;
