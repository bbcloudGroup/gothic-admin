/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50647
 Source Host           : 127.0.0.1:3306
 Source Schema         : admin

 Target Server Type    : MySQL
 Target Server Version : 50647
 File Encoding         : 65001

 Date: 21/05/2020 09:54:31
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for log
-- ----------------------------
DROP TABLE IF EXISTS `log`;
CREATE TABLE `log` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `user_id` int(10) unsigned DEFAULT NULL,
  `menu_id` int(10) unsigned DEFAULT NULL,
  `data` varchar(1024) COLLATE utf8_bin DEFAULT NULL,
  `method_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_log_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Table structure for menu
-- ----------------------------
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `tag` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `name` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `type` int(11) DEFAULT NULL,
  `parent_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_menu_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of menu
-- ----------------------------
BEGIN;
INSERT INTO `menu` VALUES (1, '2020-05-14 11:09:06', '2020-05-18 17:27:14', NULL, 'admin', '管理页', 1, 0);
INSERT INTO `menu` VALUES (2, '2020-05-14 11:09:26', '2020-05-19 09:45:55', NULL, 'user', '用户管理', 2, 1);
INSERT INTO `menu` VALUES (3, '2020-05-14 11:09:26', '2020-05-19 09:45:55', NULL, 'add', '创建', 3, 2);
INSERT INTO `menu` VALUES (4, '2020-05-14 11:09:26', '2020-05-19 09:45:55', NULL, 'update', '更新', 3, 2);
INSERT INTO `menu` VALUES (5, '2020-05-14 11:09:26', '2020-05-19 09:45:55', NULL, 'delete', '删除', 3, 2);
INSERT INTO `menu` VALUES (6, '2020-05-14 11:09:26', '2020-05-19 09:45:55', NULL, 'get', '查询', 3, 2);
INSERT INTO `menu` VALUES (12, '2020-05-15 16:24:03', '2020-05-19 09:45:55', NULL, 'resetPassword', '重置密码', 3, 2);
INSERT INTO `menu` VALUES (13, '2020-05-15 16:24:47', '2020-05-19 09:45:55', NULL, 'statusChange', '用户状态', 3, 2);
INSERT INTO `menu` VALUES (14, '2020-05-18 17:24:10', '2020-05-18 17:27:14', NULL, 'role', '角色管理', 2, 1);
INSERT INTO `menu` VALUES (15, '2020-05-18 17:24:10', '2020-05-18 17:27:14', NULL, 'delete', '删除', 3, 14);
INSERT INTO `menu` VALUES (16, '2020-05-18 17:24:10', '2020-05-18 17:27:14', NULL, 'get', '查询', 3, 14);
INSERT INTO `menu` VALUES (17, '2020-05-18 17:24:11', '2020-05-18 17:27:14', NULL, 'add', '创建', 3, 14);
INSERT INTO `menu` VALUES (18, '2020-05-18 17:24:11', '2020-05-18 17:27:14', NULL, 'update', '更新', 3, 14);
INSERT INTO `menu` VALUES (19, '2020-05-18 17:24:44', '2020-05-18 17:27:14', NULL, 'menu', '菜单管理', 2, 1);
INSERT INTO `menu` VALUES (20, '2020-05-18 17:24:44', '2020-05-18 17:27:14', NULL, 'get', '查询', 3, 19);
INSERT INTO `menu` VALUES (21, '2020-05-18 17:24:44', '2020-05-18 17:27:14', NULL, 'add', '创建', 3, 19);
INSERT INTO `menu` VALUES (22, '2020-05-18 17:24:44', '2020-05-18 17:27:14', NULL, 'update', '更新', 3, 19);
INSERT INTO `menu` VALUES (23, '2020-05-18 17:24:44', '2020-05-18 17:27:14', NULL, 'delete', '删除', 3, 19);
INSERT INTO `menu` VALUES (24, '2020-05-19 10:01:04', '2020-05-21 09:25:35', NULL, 'log', '操作日志', 2, 1);
INSERT INTO `menu` VALUES (25, '2020-05-19 10:01:04', '2020-05-21 09:25:35', NULL, 'get', '查询', 3, 24);
INSERT INTO `menu` VALUES (26, '2020-05-19 10:01:04', '2020-05-21 09:25:35', NULL, 'add', '创建', 3, 24);
INSERT INTO `menu` VALUES (27, '2020-05-19 10:01:04', '2020-05-21 09:25:35', NULL, 'update', '更新', 3, 24);
INSERT INTO `menu` VALUES (28, '2020-05-19 10:01:04', '2020-05-21 09:25:35', NULL, 'delete', '删除', 3, 24);
COMMIT;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(100) COLLATE utf8_bin NOT NULL,
  `tag` varchar(100) COLLATE utf8_bin NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_role_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of role
-- ----------------------------
BEGIN;
INSERT INTO `role` VALUES (1, '2020-05-11 19:15:48', '2020-05-21 09:25:52', NULL, '管理员', 'admin');
INSERT INTO `role` VALUES (2, '2020-05-11 19:17:23', '2020-05-21 09:25:35', NULL, '普通用户', 'user');
INSERT INTO `role` VALUES (3, '2020-05-11 19:25:18', '2020-05-21 09:31:25', NULL, '客人', 'guest');
COMMIT;

-- ----------------------------
-- Table structure for role_menu
-- ----------------------------
DROP TABLE IF EXISTS `role_menu`;
CREATE TABLE `role_menu` (
  `role_id` int(10) unsigned NOT NULL DEFAULT '0',
  `menu_id` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`role_id`,`menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of role_menu
-- ----------------------------
BEGIN;
INSERT INTO `role_menu` VALUES (2, 24);
INSERT INTO `role_menu` VALUES (2, 25);
INSERT INTO `role_menu` VALUES (2, 26);
INSERT INTO `role_menu` VALUES (2, 27);
INSERT INTO `role_menu` VALUES (2, 28);
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `password` varchar(100) COLLATE utf8_bin NOT NULL,
  `name` varchar(100) COLLATE utf8_bin NOT NULL,
  `avatar` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  `mail` varchar(50) COLLATE utf8_bin DEFAULT NULL,
  `mobile` varchar(11) COLLATE utf8_bin DEFAULT NULL,
  `status` tinyint(1) DEFAULT NULL,
  `remember_token` varchar(100) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uix_user_mail` (`mail`),
  UNIQUE KEY `uix_user_mobile` (`mobile`),
  KEY `idx_user_deleted_at` (`deleted_at`),
  KEY `idx_user_mail` (`mail`),
  KEY `idx_user_mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (1, '2020-05-07 16:32:53', '2020-05-21 09:25:52', NULL, 'e10adc3949ba59abbe56e057f20f883e', 'admin', 'https://secure.gravatar.com/avatar/6b57c6120439ec8fe7e5e83533d38af5?d=identicon', 'admin', '10000', 1, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJuYW1lIjoiYWRtaW4iLCJtYWlsIjoiYWRtaW4iLCJhdmF');
COMMIT;

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role` (
  `user_id` int(10) unsigned NOT NULL DEFAULT '0',
  `role_id` int(10) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (`user_id`,`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of user_role
-- ----------------------------
BEGIN;
INSERT INTO `user_role` VALUES (1, 1);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
