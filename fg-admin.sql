/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80015
 Source Host           : localhost
 Source Database       : fg-admin

 Target Server Type    : MySQL
 Target Server Version : 80015
 File Encoding         : utf-8

 Date: 03/27/2020 16:31:45 PM
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `operation_logs`
-- ----------------------------
DROP TABLE IF EXISTS `operation_logs`;
CREATE TABLE `operation_logs` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `uid` int(11) DEFAULT NULL,
  `method` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `path` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `params` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci,
  `ip` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `status_code` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
--  Table structure for `permissions`
-- ----------------------------
DROP TABLE IF EXISTS `permissions`;
CREATE TABLE `permissions` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `description` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `auth_path` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `pid` int(11) DEFAULT NULL,
  `icon` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `is_show` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `idx_permissions_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=61 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
--  Records of `permissions`
-- ----------------------------
BEGIN;
INSERT INTO `permissions` VALUES ('1', '2019-11-06 10:55:24', '2019-11-06 10:55:24', null, '权限管理', '权限管理', '/auth/manage', '0', 'fa-id-card', '1'), ('4', '2019-11-06 10:55:24', '2019-11-06 10:55:24', null, '用户列表', '用户列表', '/user/list', '1', 'fa-user-o', '1'), ('5', '2019-11-06 10:55:24', '2019-11-06 10:55:24', null, '角色列表', '角色列表', '/role/list', '1', 'fa-user-circle-o', '1'), ('6', '2019-11-06 10:55:24', '2019-11-06 10:55:24', null, '权限列表', '权限列表', '/auth/list', '1', 'fa-key', '1'), ('7', '2019-11-06 10:55:24', '2019-11-06 10:55:24', null, '服务器管理', '服务器管理', '/server/index', '0', 'fa-server', '1'), ('8', '2019-11-06 10:55:24', '2019-11-11 14:29:10', null, '服务器列表', '登录服列表', '/server/list', '7', 'fa-server', '1'), ('9', '2019-11-06 10:55:24', '2019-11-09 17:13:53', '2019-12-18 10:44:38', '服务器表单', '登录表单', '/server/table', '7', 'fa-server', '0'), ('10', '2019-11-06 10:55:24', '2019-11-13 18:34:22', '2019-12-18 10:44:16', '用户列表表单', '用户列表表单', '/user/table', '1', 'fa-server', '0'), ('11', '2019-11-06 10:55:24', '2019-11-09 17:31:57', null, '服务器新增/修改', '服务器修改', '/server/upsert', '7', 'fa-server', '0'), ('12', '2019-11-06 10:55:24', '2019-11-06 10:55:24', null, '服务器删除', '服务器删除', '/server/delete', '7', 'fa-server', '0'), ('13', '2019-11-06 10:55:24', '2019-11-12 17:40:30', null, '侧边菜单', '侧边菜单', '/home/menu', '39', 'fa-server', '0'), ('14', '2019-11-06 10:55:24', '2019-11-12 17:32:42', null, 'GM添加资源', 'GM添加资源', '/server/cmd', '39', 'fa-laptop', '1'), ('15', '2019-11-06 10:55:24', '2019-11-12 17:32:49', null, 'GM发邮件', 'GM发邮件', '/server/mail', '39', 'fa-paper-plane', '1'), ('16', '2019-11-06 10:55:24', '2019-11-09 17:13:21', '2019-12-18 10:44:21', '角色表格', '角色表格', '/role/table', '1', 'fa-mail', '0'), ('17', '2019-11-06 10:55:24', '2019-11-11 14:41:20', null, '用户新增/修改', '用户编辑', '/user/upsert', '1', 'fa-mail', '0'), ('18', '2019-11-06 10:55:24', '2019-11-09 17:14:07', null, '用户删除', '用户删除', '/user/delete', '1', 'fa-mail', '0'), ('19', '2019-11-06 10:55:24', '2019-11-09 17:13:01', null, '权限节点', '权限节点', '/auth/nodes', '1', 'fa-mail', '0'), ('20', '2019-11-06 10:55:24', '2019-11-09 17:13:10', null, '权限单个节点', '权限单个节点', '/auth/node', '1', 'fa-mail', '0'), ('21', '2019-11-06 10:55:24', '2019-11-11 13:57:57', null, '权限新增/修改', '权限添加更新', '/auth/upsert', '1', 'fa-mail', '0'), ('23', '2019-11-09 17:17:19', '2019-11-09 17:17:19', null, '权限删除', '权限删除', '/auth/delete', '1', 'fa-mail', '0'), ('32', '2019-11-09 17:47:55', '2019-11-09 17:47:55', null, '角色新增/修改', '角色新增/修改', '/role/upsert', '1', 'fa-mail', '0'), ('33', '2019-11-11 11:16:21', '2019-11-11 11:16:21', '2019-11-11 15:21:36', '权限节点data', '', '/auth/data', '1', '', '0'), ('37', '2019-11-06 10:55:24', '2019-11-11 15:23:30', null, '权限树', '权限树', '/auth/tree', '1', 'fa-mail', '0'), ('38', '2019-11-11 17:08:17', '2019-11-11 17:08:17', null, '角色删除', '', '/role/delete', '1', '', '0'), ('39', '2019-11-12 17:32:18', '2019-11-12 17:32:18', null, 'GM命令', '', '/gm/index', '0', '', '1'), ('40', '2019-11-13 18:07:37', '2019-11-13 18:08:01', null, '服务器配置修改', '', '/server/conf/edit', '7', 'fa-server', '1'), ('41', '2019-11-14 10:22:25', '2019-11-14 10:22:25', null, '上传Items配表', '', '/server/upload', '7', 'fa-server', '1'), ('42', '2019-11-14 17:17:53', '2019-11-14 17:43:09', null, '操作日志', '', '/server/operate/list', '7', 'fa-server', '1'), ('43', '2019-11-14 17:18:17', '2019-11-14 17:18:17', '2019-12-18 10:48:54', '操作日志表格', '', '/server/operate/table', '7', 'fa-server', '0'), ('44', '2019-12-11 15:17:47', '2020-03-05 11:41:34', null, '活动管理', '', '/season/list', '39', 'fa-user-o', '1'), ('45', '2019-12-11 15:18:26', '2019-12-18 10:32:46', null, '赛季新增修改', '', '/season/upsert', '39', 'fa-user-o', '0'), ('46', '2019-12-11 15:18:41', '2019-12-11 15:18:41', null, '赛季删除', '', '/season/delete', '39', 'fa-user-o', '0'), ('47', '2019-12-11 16:23:17', '2019-12-16 14:11:31', '2019-12-18 10:32:31', '赛季数据', '', '/season/table', '39', 'fa-user-o', '1'), ('48', '2020-03-03 14:50:59', '2020-03-04 15:04:02', null, '兑换码生成', '', '/server/code/gen', '39', 'fa-user-o', '0'), ('49', '2020-03-03 16:19:37', '2020-03-03 16:19:37', '2020-03-03 16:46:45', '兑换码导出', '', '/server/exportcode', '39', 'fa-user-o', '0'), ('50', '2020-03-04 15:03:55', '2020-03-04 15:03:55', null, '礼包兑换码', '', '/server/code/list', '39', 'fa-user-o', '1'), ('51', '2020-03-06 14:37:14', '2020-03-06 14:37:14', null, '兑换码删除', '', '/server/code/delete', '39', 'fa-user-o', '0'), ('52', '2020-03-09 14:06:06', '2020-03-09 14:06:06', null, '活动配置上传', '', '/season/upload', '39', 'fa-user-o', '0'), ('53', '2020-03-09 17:45:39', '2020-03-09 17:45:39', null, '活动信息修改', '', '/season/edit', '39', 'fa-user-o', '0'), ('54', '2020-03-10 14:25:15', '2020-03-10 14:25:15', null, '活动配置加载', '', '/season/load', '39', 'fa-user-o', '0'), ('55', '2020-03-11 11:24:14', '2020-03-11 11:24:14', null, '数据监测', '', '/', '0', '', '1'), ('56', '2020-03-11 11:26:38', '2020-03-11 11:28:28', null, '玩家数据', '', '/monitor/player', '55', 'fa-user-o', '1'), ('57', '2020-03-17 17:39:09', '2020-03-17 17:39:09', null, '在线玩家', '', '/monitor/online', '55', 'fa-user-o', '1'), ('58', '2020-03-20 18:11:13', '2020-03-20 18:12:37', null, '在线数据导出', '', '/monitor/online/export', '55', 'fa-user-o', '0'), ('59', '2020-03-24 14:36:20', '2020-03-24 17:32:56', null, '玩家资源流水', '', '/monitor/flow', '55', 'fa-user-o', '1'), ('60', '2020-03-27 14:45:53', '2020-03-27 14:45:53', null, '服务器状态', '', '/server/status', '7', 'fa-server', '1');
COMMIT;

-- ----------------------------
--  Table structure for `role_perms`
-- ----------------------------
DROP TABLE IF EXISTS `role_perms`;
CREATE TABLE `role_perms` (
  `role_id` int(10) unsigned NOT NULL,
  `permission_id` int(10) unsigned NOT NULL,
  PRIMARY KEY (`role_id`,`permission_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
--  Records of `role_perms`
-- ----------------------------
BEGIN;
INSERT INTO `role_perms` VALUES ('1', '1'), ('1', '4'), ('1', '5'), ('1', '6'), ('1', '7'), ('1', '8'), ('1', '11'), ('1', '13'), ('1', '14'), ('1', '15'), ('1', '17'), ('1', '18'), ('1', '19'), ('1', '20'), ('1', '21'), ('1', '23'), ('1', '32'), ('1', '37'), ('1', '38'), ('1', '39'), ('1', '40'), ('1', '41'), ('1', '42'), ('1', '44'), ('1', '45'), ('1', '46'), ('1', '48'), ('1', '50'), ('1', '51'), ('1', '52'), ('1', '53'), ('1', '54'), ('1', '55'), ('1', '56'), ('1', '57'), ('1', '58'), ('1', '59'), ('1', '60'), ('2', '1'), ('2', '4'), ('2', '5'), ('2', '13'), ('2', '19'), ('2', '39'), ('3', '1'), ('3', '13'), ('4', '1'), ('4', '16'), ('4', '17'), ('4', '18'), ('5', '1'), ('5', '4'), ('5', '5'), ('5', '6'), ('5', '10'), ('5', '13'), ('5', '16'), ('5', '17'), ('5', '18'), ('5', '19'), ('5', '20'), ('5', '21'), ('5', '23'), ('5', '32'), ('5', '37'), ('5', '38'), ('6', '7'), ('6', '8'), ('6', '13'), ('6', '14'), ('6', '15'), ('6', '39'), ('6', '44'), ('6', '45'), ('6', '46'), ('6', '48'), ('6', '50'), ('6', '51'), ('6', '52'), ('6', '53'), ('6', '54');
COMMIT;

-- ----------------------------
--  Table structure for `roles`
-- ----------------------------
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `description` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  KEY `idx_roles_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
--  Records of `roles`
-- ----------------------------
BEGIN;
INSERT INTO `roles` VALUES ('1', '2019-11-06 10:55:24', '2020-03-27 14:46:03', null, 'admin', '超级管理员'), ('2', '2019-11-06 10:55:24', '2019-11-13 14:50:22', null, 'QA', '测试权限'), ('3', '2019-11-09 18:19:17', '2019-11-09 18:19:17', '2019-11-11 17:18:46', 'yey', 'ferer'), ('4', '2019-11-11 17:04:16', '2019-11-11 17:04:16', '2019-11-11 17:18:43', 'ces', 'ces'), ('5', '2019-11-11 17:16:48', '2019-11-11 17:16:48', '2019-11-11 17:18:32', 'jfjf', 'fdfdf'), ('6', '2019-11-12 17:30:33', '2020-03-10 14:32:36', null, 'GM', '供策划使用，拥有所有GM权限');
COMMIT;

-- ----------------------------
--  Table structure for `users`
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `name` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `username` varchar(32) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `password` varchar(200) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `role_id` int(10) unsigned DEFAULT NULL,
  `email` varchar(32) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `phone` varchar(32) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
--  Records of `users`
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES ('1', '2019-11-06 10:53:24', '2019-11-12 17:37:00', null, 'ding', 'ding', '$2a$10$wZENwl97NKrr0FGQOHIazOZFZI8CRY3TchVUHR3uxiL5h3r6F3ABm', '6', 'b835399559@126.com', '18352332612'), ('2', '2019-11-06 10:53:24', '2019-11-06 10:53:24', null, 'admin', 'admin', '$2a$10$wZENwl97NKrr0FGQOHIazOZFZI8CRY3TchVUHR3uxiL5h3r6F3ABm', '1', 'a835399559@gmail.com', null), ('3', '2019-11-06 10:53:24', '2019-11-13 13:52:23', null, 'test1223', 'test1', '$2a$10$yGoJ6zWNuBlYlPnSh4uXKuSRx7maV.rRXwuSLqv2fKzP/fUAoZEdi', '6', 'b835399559@126.com', '18352332612'), ('4', '2019-11-06 10:53:24', '2019-11-13 14:25:14', null, 'test266', 'test1', '$2a$10$QBMHRzo.FoOpFZCFGXOcjO1O5L1jWdEf6b.vsUkoMk8Ngyo5Nqn16', '1', 'b835399559@126.com', '18356363445'), ('5', '2019-11-06 10:53:24', '2019-11-06 10:53:24', null, 'test3', 'test1', '$2a$10$wZENwl97NKrr0FGQOHIazOZFZI8CRY3TchVUHR3uxiL5h3r6F3ABm', '1', 'b835399559@126.com', null), ('6', '2019-11-06 10:53:24', '2019-11-06 10:53:24', null, 'test4', 'test1', '$2a$10$wZENwl97NKrr0FGQOHIazOZFZI8CRY3TchVUHR3uxiL5h3r6F3ABm', '1', 'b835399559@126.com', null), ('7', '2019-11-06 10:53:24', '2019-11-14 11:02:56', null, 'test5', 'test1', '$2a$10$j2xUdqqCF2NWVSfSFGve1.XeDijJ2nnj4wgWzOZvesr2CXUsoSn1S', '2', 'b835399559@126.com', '18938378827'), ('8', '2019-11-06 10:53:24', '2019-11-14 13:44:19', null, 'test67', 'test1', '$2a$10$m0VrVJYNhiGHVmOsK82auO0TQAXh9iFa2xJSaRss7PlfUuRrcpkR.', '2', 'b835399559@126.com', '17787633343'), ('9', '2019-11-06 10:53:24', '2019-11-06 10:53:24', null, 'test7', 'test1', '$2a$10$wZENwl97NKrr0FGQOHIazOZFZI8CRY3TchVUHR3uxiL5h3r6F3ABm', '1', 'b835399559@126.com', null), ('10', '2019-11-06 10:53:24', '2019-11-06 10:53:24', null, 'test8', 'test1', '$2a$10$wZENwl97NKrr0FGQOHIazOZFZI8CRY3TchVUHR3uxiL5h3r6F3ABm', '1', 'b835399559@126.com', null), ('11', '2019-11-06 10:53:24', '2019-11-06 10:53:24', null, 'test9', 'test1', '$2a$10$wZENwl97NKrr0FGQOHIazOZFZI8CRY3TchVUHR3uxiL5h3r6F3ABm', '1', 'b835399559@126.com', null), ('12', '2019-11-06 10:53:24', '2019-11-06 10:53:24', '2020-03-18 16:40:06', 'test10', 'test1', '$2a$10$wZENwl97NKrr0FGQOHIazOZFZI8CRY3TchVUHR3uxiL5h3r6F3ABm', '1', 'b835399559@126.com', null), ('13', '2019-11-09 10:44:59', '2019-11-09 10:44:59', '2019-11-09 10:53:39', 'a7748', '', '$2a$10$amvLH1yDxATHZcYYT5LwbO1.6imO2M0HCqRaZhGEIidYghV9XLUse', '0', '', ''), ('14', '2019-11-09 10:48:50', '2019-11-09 10:48:50', '2019-11-09 10:53:44', 'a7788', '', '$2a$10$NtsmOmIqH/EQz4gzCDIWU.Y5HdGZPKuh/d5FkexeMvPtu/Du8yf..', '2', 'b887373@qq.com', '18352332612'), ('15', '2019-11-12 18:16:30', '2019-11-12 18:16:30', null, 'cehua', 'cehua', '$2a$10$wZENwl97NKrr0FGQOHIazOZFZI8CRY3TchVUHR3uxiL5h3r6F3ABm', '6', 'b84392728@126.com', '18033849932');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
