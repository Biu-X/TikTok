/*
 Navicat Premium Data Transfer

 Source Server         : Mac
 Source Server Type    : MySQL
 Source Server Version : 80033
 Source Host           : localhost:3306
 Source Schema         : tiktok

 Target Server Type    : MySQL
 Target Server Version : 80033
 File Encoding         : 65001

 Date: 06/08/2023 21:14:35
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL COMMENT '用户id',
  `video_id` bigint NOT NULL COMMENT '视频id',
  `content` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '评论内容',
  `created_at` datetime(3) DEFAULT NULL COMMENT '评论发布日期',
  `deleted_at` datetime(3) DEFAULT NULL COMMENT '评论删除日期',
  PRIMARY KEY (`id`),
  KEY `idx_comment_video_id` (`video_id`),
  KEY `idx_comment_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for favorite
-- ----------------------------
DROP TABLE IF EXISTS `favorite`;
CREATE TABLE `favorite` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL COMMENT '用户id',
  `video_id` bigint NOT NULL COMMENT '视频id',
  `cancel` tinyint NOT NULL COMMENT '取消赞是1，默认0',
  PRIMARY KEY (`id`),
  KEY `idx_favorite_user_id` (`user_id`),
  KEY `idx_favorite_video_id` (`video_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for follow
-- ----------------------------
DROP TABLE IF EXISTS `follow`;
CREATE TABLE `follow` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL COMMENT '用户id',
  `follower_id` bigint NOT NULL COMMENT '被关注用户id',
  `cancel` tinyint NOT NULL COMMENT '取消关注是1，默认0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_follow_user_id_follow_id` (`user_id`,`follower_id`),
  KEY `idx_follow_user_id` (`user_id`),
  KEY `idx_follow_follow_id` (`follower_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `to_user_id` bigint NOT NULL COMMENT '消息接收者id',
  `from_user_id` bigint NOT NULL COMMENT '消息发送者id',
  `content` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '消息内容',
  `created_at` datetime(3) DEFAULT NULL COMMENT '消息发送时间',
  PRIMARY KEY (`id`),
  KEY `idx_message_from_user_id_to_user_id` (`to_user_id`,`from_user_id`),
  KEY `idx_message_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `signature` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '个人简介',
  `avatar` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '头像',
  `background_image` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户个人页顶部大图',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uq_user_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Table structure for video
-- ----------------------------
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `author_id` bigint NOT NULL COMMENT '作者信息',
  `play_url` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '播放地址',
  `cover_url` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '封面地址',
  `created_at` datetime(3) DEFAULT NULL COMMENT '发布时间',
  `title` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '标题',
  PRIMARY KEY (`id`),
  KEY `idx_video_author_id` (`author_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

SET FOREIGN_KEY_CHECKS = 1;
