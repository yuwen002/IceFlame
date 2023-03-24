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

 Date: 24/03/2023 17:43:38
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
  `uri` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '对应的URI地址',
  `type` tinyint(4) NULL DEFAULT NULL COMMENT '类型（1=菜单，2=按钮, 3=提交程序）',
  `status` tinyint(3) UNSIGNED NULL DEFAULT 0 COMMENT '程序模块状态（0=启用，1=停用）',
  `sort` int(11) NULL DEFAULT NULL COMMENT '排序',
  `supper_master` tinyint(1) NULL DEFAULT 0 COMMENT '1为超级管理员模块',
  `remark` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '备注信息',
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 45 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '后台系统权限模块列表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of uc_system_permission
-- ----------------------------
INSERT INTO `uc_system_permission` VALUES (1, 0, '权限管理', NULL, NULL, 1, 0, NULL, 1, '一级导航', '2023-03-13 11:19:24', '2023-03-13 11:19:24');
INSERT INTO `uc_system_permission` VALUES (2, 1, '管理员角色列表', 'manage.UcSystemMasterAuth.ListRole', 'manage/master/auth/show_role', 1, 0, NULL, 1, '二级导航', '2023-03-13 11:45:46', '2023-03-13 11:45:46');
INSERT INTO `uc_system_permission` VALUES (3, 1, '添加管理员角色', 'manage.UcSystemMasterAuth.AddRole', 'manage/master/auth/add_role', 2, 0, NULL, 1, '二级按钮', '2023-03-13 12:00:47', '2023-03-13 12:00:47');
INSERT INTO `uc_system_permission` VALUES (4, 1, '编辑管理员角色', 'manage.UcSystemMasterAuth.EditRole', 'manage/master/auth/edit_role', 2, 0, NULL, 1, '二级程序', '2023-03-13 12:01:06', '2023-03-13 12:01:06');
INSERT INTO `uc_system_permission` VALUES (5, 1, '管理员角色绑定列表', 'manage.UcSystemMasterAuth.ListRoleRelation', 'manage/master/auth/show_role_relation', 1, 0, NULL, 1, '二级导航', '2023-03-13 13:46:48', '2023-03-13 13:46:48');
INSERT INTO `uc_system_permission` VALUES (6, 1, '添加管理员角色绑定', 'manage.UcSystemMasterAuth.AddRoleRelation', 'manage/master/auth/add_role_relation', 2, 0, NULL, 1, '二级按钮', '2023-03-13 13:47:30', '2023-03-13 13:47:30');
INSERT INTO `uc_system_permission` VALUES (7, 1, '编辑管理员角色绑定', 'manage.UcSystemMasterAuth.EditRoleRelation', 'manage/master/auth/edit_role_relation', 3, 0, NULL, 1, '处理程序', '2023-03-13 13:49:48', '2023-03-13 13:49:48');
INSERT INTO `uc_system_permission` VALUES (8, 1, '删除管理员角色绑定', 'manage.UcSystemMasterAuth.DeleteRoleRelation', 'manage/master/auth/delete_role_relation', 2, 0, NULL, 1, '二级按钮', '2023-03-13 13:50:14', '2023-03-13 13:50:14');
INSERT INTO `uc_system_permission` VALUES (9, 1, '权限列表', 'manage.UcSystemMasterAuth.ListPermission', 'manage/master/auth/show_permission', 1, 0, NULL, 1, '二级菜单', '2023-03-13 14:30:56', '2023-03-13 14:30:56');
INSERT INTO `uc_system_permission` VALUES (10, 1, '添加权限', 'manage.UcSystemMasterAuth.AddPermission', 'manage/master/auth/add_permission', 2, 0, NULL, 1, '二级按钮', '2023-03-13 14:31:57', '2023-03-13 14:31:57');
INSERT INTO `uc_system_permission` VALUES (11, 1, '编辑权限', 'manage.UcSystemMasterAuth.EditPermission', 'manage/master/auth/edit_permission', 3, 0, NULL, 1, '处理程序', '2023-03-13 14:32:15', '2023-03-13 14:32:15');
INSERT INTO `uc_system_permission` VALUES (12, 1, '查看权限分配', 'manage.UcSystemMasterAuth.ListPermissionRelation', 'manage/master/auth/show_permission_relation', 2, 0, NULL, 1, '二级按钮', '2023-03-15 17:57:35', '2023-03-15 17:57:35');
INSERT INTO `uc_system_permission` VALUES (13, 1, '编辑权限分配', 'manage.UcSystemMasterAuth.EditPermissionRelation', 'manage/master/auth/show_permission_relation', 3, 0, NULL, 1, '处理程序', '2023-03-15 17:58:59', '2023-03-15 17:58:59');
INSERT INTO `uc_system_permission` VALUES (14, 1, '管理员用户列表', 'manage.UcSystemMaster.ListSystemMaster', 'manage/master/show_system_master', 1, 0, NULL, 1, '二级菜单', '2023-03-15 18:06:12', '2023-03-15 18:06:12');
INSERT INTO `uc_system_permission` VALUES (15, 1, '新建管理员用户', 'manage.UcSystemMaster.CreateSystemMaster', 'manage/master/create_system_master', 2, 0, NULL, 1, '二级按钮', '2023-03-15 18:07:06', '2023-03-15 18:07:06');
INSERT INTO `uc_system_permission` VALUES (16, 1, '编辑管理员用户', 'manage.UcSystemMaster.EditSystemMaster', 'manage/master/edit_system_master', 3, 0, NULL, 1, '处理程序', '2023-03-15 18:07:48', '2023-03-15 18:07:48');
INSERT INTO `uc_system_permission` VALUES (17, 1, '重置管理员用户密码', 'manage.UcSystemMaster.ResetPassword', 'manage/master/reset_password', 2, 0, NULL, 1, '二级按钮', '2023-03-15 18:08:23', '2023-03-15 18:08:23');
INSERT INTO `uc_system_permission` VALUES (18, 1, '编辑管理员用户状态', 'manage.UcSystemMaster.EditStatus', 'manage/master/edit_status', 3, 0, NULL, 1, '处理程序', '2023-03-15 18:09:34', '2023-03-15 18:09:34');
INSERT INTO `uc_system_permission` VALUES (19, 1, '解锁管理员用户限制登入', 'manage.UcSystemMaster.UnlockSystemMaster', 'manage/master/unlock_system_master', 2, 0, NULL, 1, '二级按钮', '2023-03-15 18:10:09', '2023-03-15 18:10:09');
INSERT INTO `uc_system_permission` VALUES (20, 1, '查看管理员角色', 'manage.UcSystemMasterAuth.GetEditRole', 'manage/master/auth/edit_role', 2, 0, NULL, 1, '二级按钮', '2023-03-16 17:21:26', '2023-03-16 17:21:26');
INSERT INTO `uc_system_permission` VALUES (21, 1, '查看管理员角色绑定', 'manage.UcSystemMasterAuth.GetEditRoleRelation', 'manage/master/auth/edit_role_relation', 2, 0, 1, 1, '二级按钮', '2023-03-16 17:23:42', '2023-03-16 17:23:42');
INSERT INTO `uc_system_permission` VALUES (22, 1, '查看编辑权限信息', 'manage.UcSystemMasterAuth.GetEditRoleRelation', 'manage/master/auth/edit_permission', 2, 0, NULL, 1, '处理程序', '2023-03-18 14:26:20', '2023-03-18 14:26:20');
INSERT INTO `uc_system_permission` VALUES (23, 1, '编辑管理员用户', 'manage.UcSystemMaster.GetEditSystemMaster', 'manage/master/auth/edit_permission', 2, 0, NULL, 1, '二级按钮', '2023-03-18 15:28:23', '2023-03-18 15:28:23');
INSERT INTO `uc_system_permission` VALUES (24, 1, '访问类型列表', 'manage.UcSystemMasterVisitor.ListVisitCategory', 'manage/visitor_logs/show', 1, 0, NULL, 1, '二级按钮', '2023-03-18 16:42:25', '2023-03-18 16:42:25');
INSERT INTO `uc_system_permission` VALUES (25, 1, '添加访问类型', 'manage.UcSystemMasterVisitor.AddVisitCategory', 'manage/visit_category/add', 2, 0, NULL, 1, '二级按钮', '2023-03-18 16:44:08', '2023-03-18 16:44:08');
INSERT INTO `uc_system_permission` VALUES (26, 1, '编辑访问类型', 'manage.UcSystemMasterVisitor.GetEditVisitCategory', 'manage/visit_category/edit', 2, 0, NULL, 1, '二级按钮', '2023-03-18 16:45:49', '2023-03-18 16:45:49');
INSERT INTO `uc_system_permission` VALUES (27, 1, '编辑访问类型', 'manage.UcSystemMasterVisitor.EditVisitCategory', 'manage/visit_category/edit', 3, 0, NULL, 1, '处理程序', '2023-03-18 16:48:29', '2023-03-18 16:48:29');
INSERT INTO `uc_system_permission` VALUES (28, 1, '删除访问类型缓存', 'manage.UcSystemMasterVisitor.DeleteCacheVisitCategory', 'manage/visit_category/delete', 2, 0, NULL, 1, '二级按钮', '2023-03-18 17:01:07', '2023-03-18 17:01:07');
INSERT INTO `uc_system_permission` VALUES (29, 1, '访问日志列表', 'manage.UcSystemMasterVisitor.ListVisitorLogs', 'manage/visitor_logs/show', 1, 0, NULL, 1, '二级菜单', '2023-03-18 17:03:39', '2023-03-18 17:03:39');
INSERT INTO `uc_system_permission` VALUES (30, 0, '文章管理', '', NULL, 1, 0, NULL, 0, '一级菜单', '2023-03-22 16:39:34', '2023-03-22 16:39:34');
INSERT INTO `uc_system_permission` VALUES (31, 30, '添加单页', 'manage.Article.AddSinglePage', 'manage/article/single_page/add', 2, 0, NULL, 0, '二级按钮', '2023-03-22 17:04:57', '2023-03-22 17:04:57');
INSERT INTO `uc_system_permission` VALUES (32, 30, '编辑单页', 'manage.Article.GetSinglePage', 'manage/article/single_page/edit', 2, 0, NULL, 0, '二级按钮', '2023-03-22 17:07:56', '2023-03-22 17:07:56');
INSERT INTO `uc_system_permission` VALUES (33, 30, '编辑单页', 'manage.Article.EditSinglePage', 'manage/article/single_page/edit', 3, 0, NULL, 0, '处理程序', '2023-03-22 18:03:57', '2023-03-22 18:03:57');
INSERT INTO `uc_system_permission` VALUES (34, 30, '单页列表', 'manage.Article.ListSinglePage', 'manage/article/single_page/show', 1, 0, NULL, 0, '二级按钮', '2023-03-22 18:11:52', '2023-03-22 18:11:52');
INSERT INTO `uc_system_permission` VALUES (35, 30, '编辑单页', 'manage.Article.DeleteSinglePage', 'manage/article/single_page/delete', 2, 0, NULL, 0, '二级按钮', '2023-03-22 23:29:56', '2023-03-22 23:29:56');
INSERT INTO `uc_system_permission` VALUES (36, 30, '添加频道', 'manage.Article.AddArticleChannel', 'manage/article/channel/add', 2, 0, NULL, 0, '二级按钮', '2023-03-23 14:38:40', '2023-03-23 14:38:40');
INSERT INTO `uc_system_permission` VALUES (37, 30, '编辑频道', 'manage.Article.GetArticleChannel', 'manage/article/channel/edit', 2, 0, NULL, 0, '二级按钮', '2023-03-23 15:57:17', '2023-03-23 15:57:17');
INSERT INTO `uc_system_permission` VALUES (38, 30, '编辑频道', 'manage.Article.EditArticleChannel', 'manage/article/channel/edit', 3, 0, NULL, 0, '二级按钮', '2023-03-23 15:59:29', '2023-03-23 15:59:29');
INSERT INTO `uc_system_permission` VALUES (39, 30, '频道列表', 'manage.Article.ListArticleChannel', 'manage/article/channel/show', 1, 0, NULL, 0, '二级导航', '2023-03-23 16:25:00', '2023-03-23 16:25:00');
INSERT INTO `uc_system_permission` VALUES (40, 30, '删除频道', 'manage.Article.DelArticleChannel', 'manage/article/channel/delete', 2, 0, NULL, 0, '二级导航', '2023-03-23 16:35:02', '2023-03-23 16:35:02');
INSERT INTO `uc_system_permission` VALUES (41, 30, '添加分类', 'manage.Article.AddArticleCategory', 'manage/article/category/add', 2, 0, NULL, 0, '二级按钮', '2023-03-24 16:34:06', '2023-03-24 16:34:06');
INSERT INTO `uc_system_permission` VALUES (42, 30, '编辑分类', 'manage.Article.GetArticleCategory', 'manage/article/category/edit', 2, 0, NULL, 0, '二级按钮', '2023-03-24 16:36:16', '2023-03-24 16:36:16');
INSERT INTO `uc_system_permission` VALUES (43, 30, '编辑分类', 'manage.Article.EditArticleCategory', 'manage/article/category/edit', 3, 0, NULL, 0, '处理程序', '2023-03-24 16:36:46', '2023-03-24 16:36:46');
INSERT INTO `uc_system_permission` VALUES (44, 30, '分类列表', 'manage.Article.ListArticleCategory', 'manage/article/category/show', 1, 0, NULL, 0, '二级菜单', '2023-03-24 16:39:01', '2023-03-24 16:39:01');

SET FOREIGN_KEY_CHECKS = 1;
