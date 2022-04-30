/*
 Navicat Premium Data Transfer

 Source Server         : 42.193.136.192
 Source Server Type    : MySQL
 Source Server Version : 50735
 Source Host           : 42.193.136.192:23306
 Source Schema         : gf

 Target Server Type    : MySQL
 Target Server Version : 50735
 File Encoding         : 65001

 Date: 30/04/2022 20:30:47
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'User ID',
  `passport` varchar(45) NOT NULL COMMENT 'User Passport',
  `password` varchar(45) NOT NULL COMMENT 'User Password',
  `nickname` varchar(45) NOT NULL COMMENT 'User Nickname',
  `create_at` datetime DEFAULT NULL COMMENT 'Created Time',
  `update_at` datetime DEFAULT NULL COMMENT 'Updated Time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` (`id`, `passport`, `password`, `nickname`, `create_at`, `update_at`) VALUES (1, 'dd', 'ddd', '黄仪鹏', '2022-02-21 15:57:33', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `nickname`, `create_at`, `update_at`) VALUES (2, 'hhh', 'hhh', '111', '2022-02-21 16:32:11', NULL);
INSERT INTO `user` (`id`, `passport`, `password`, `nickname`, `create_at`, `update_at`) VALUES (3, 'sss', '222', 'hhh', '2022-02-21 16:33:38', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
