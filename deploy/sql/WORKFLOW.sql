/*
 Navicat Premium Data Transfer

 Source Server         : LOCAL
 Source Server Type    : MySQL
 Source Server Version : 80034
 Source Host           : localhost:3306
 Source Schema         : WORKFLOW

 Target Server Type    : MySQL
 Target Server Version : 80034
 File Encoding         : 65001

 Date: 28/08/2023 19:16:50
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for oxygen_records
-- ----------------------------
DROP TABLE IF EXISTS `oxygen_records`;
CREATE TABLE `oxygen_records` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `uid` int unsigned NOT NULL,
  `wid` int unsigned NOT NULL,
  `raw_data` decimal(10,3) unsigned NOT NULL DEFAULT '0.000' COMMENT '原始数据',
  `int_map` decimal(4,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '整数位映射数值',
  `float_map` decimal(4,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '浮点数映射数值',
  `count_data` decimal(4,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '计算数据',
  `fixed_data` decimal(4,2) unsigned NOT NULL DEFAULT '0.00' COMMENT '纠正数据',
  `is_fixed` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否被纠正过数据\n0：未被纠正过数据（自动计算得出）\n1：已被纠正过数据（人为主动修改）',
  `is_auto_push` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '是否为自动推到的数据\n0：不是自动推到的数据\n1：是自动推到的数据',
  `state` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '数据状态\n0：未上报\n1：已上报',
  `effect_time` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `wid` (`wid`),
  CONSTRAINT `wid` FOREIGN KEY (`wid`) REFERENCES `weekdays` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=157 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `unionid` varchar(255) DEFAULT NULL,
  `openid` varchar(255) NOT NULL,
  `session_key` varchar(48) NOT NULL,
  `session_refresh` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for weekdays
-- ----------------------------
DROP TABLE IF EXISTS `weekdays`;
CREATE TABLE `weekdays` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `uid` int unsigned NOT NULL,
  `state` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '工作日状态\n0：进行中\n1：已完成',
  `mark_time` datetime NOT NULL,
  `start_time` datetime NOT NULL,
  `end_time` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=62 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS = 1;
