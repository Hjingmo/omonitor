/*
Navicat MySQL Data Transfer

Source Server         : 192.168.1.214
Source Server Version : 50630
Source Host           : 192.168.1.214:3306
Source Database       : jz_omserver

Target Server Type    : MYSQL
Target Server Version : 50630
File Encoding         : 65001

Date: 2016-11-19 10:10:19
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `alarm_groups`
-- ----------------------------
DROP TABLE IF EXISTS `alarm_groups`;
CREATE TABLE `alarm_groups` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `group` varchar(64) DEFAULT NULL,
  `sms` tinyint(11) DEFAULT '0' COMMENT '短信发送开关0：不发送，1：发送',
  `email` tinyint(11) DEFAULT '0' COMMENT '邮件发送开关0：不发送，1：发送',
  `comment` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `gname` (`group`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for `alarm_group_users`
-- ----------------------------
DROP TABLE IF EXISTS `alarm_group_users`;
CREATE TABLE `alarm_group_users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `groupid` int(11) DEFAULT NULL,
  `userid` bigint(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `group` (`groupid`) USING BTREE,
  KEY `user` (`userid`) USING BTREE,
  CONSTRAINT `agid` FOREIGN KEY (`groupid`) REFERENCES `alarm_groups` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `auid` FOREIGN KEY (`userid`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for `kafka_consumer_groups`
-- ----------------------------
DROP TABLE IF EXISTS `kafka_consumer_groups`;
CREATE TABLE `kafka_consumer_groups` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `groupname` varchar(64) DEFAULT NULL COMMENT '消费组名称',
  `monitoring` tinyint(11) DEFAULT '0' COMMENT '是否监控(0: 不监控, 1: 监控)',
  `alarmval` int(64) DEFAULT '0' COMMENT '告警阀值(0:阀值无效)',
  `alarmgroup` int(11) DEFAULT '1' COMMENT '告警发送组',
  `alerts` tinyint(11) DEFAULT '1' COMMENT '报警次数',
  `comment` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `group` (`groupname`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=33 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for `kafka_consumer_group_topics`
-- ----------------------------
DROP TABLE IF EXISTS `kafka_consumer_group_topics`;
CREATE TABLE `kafka_consumer_group_topics` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `topicname` varchar(64) DEFAULT NULL,
  `groupid` int(11) DEFAULT '0' COMMENT '组ID',
  `monitoring` tinyint(11) DEFAULT '0' COMMENT '是否监控',
  `already` tinyint(11) DEFAULT '0' COMMENT '已告警次数',
  `alarmval` int(64) DEFAULT '0' COMMENT '告警阀值',
  `comment` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `topic` (`topicname`) USING BTREE,
  KEY `group` (`groupid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=127 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for `permissions`
-- ----------------------------
DROP TABLE IF EXISTS `permissions`;
CREATE TABLE `permissions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `codename` varchar(50) DEFAULT NULL,
  `comment` varchar(128) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `name` (`codename`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for `users`
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(32) DEFAULT NULL,
  `firstname` varchar(32) DEFAULT NULL,
  `lastname` varchar(32) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `avatar` varchar(100) DEFAULT NULL,
  `status` tinyint(10) DEFAULT '1',
  `superuser` tinyint(10) DEFAULT '0',
  `lastlogin` varchar(50) DEFAULT NULL,
  `mobile` bigint(20) DEFAULT '0',
  `email` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `INDEX` (`username`,`status`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES ('1', 'admin', '超级', '管理员', '202cb962ac59075b964b07152d234b70', 'root.png', '1', '1', '2016-11-18 15:27:30.810819308 +0800 CST', '0', 'admin@9zhitx.com');

-- ----------------------------
-- Table structure for `user_permissions`
-- ----------------------------
DROP TABLE IF EXISTS `user_permissions`;
CREATE TABLE `user_permissions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `userid` bigint(20) DEFAULT NULL,
  `permissionid` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `uid` (`userid`) USING BTREE,
  KEY `pid` (`permissionid`) USING BTREE,
  CONSTRAINT `per_id` FOREIGN KEY (`permissionid`) REFERENCES `permissions` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `user_id` FOREIGN KEY (`userid`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=278 DEFAULT CHARSET=utf8;
