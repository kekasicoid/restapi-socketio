/*
 Navicat Premium Data Transfer

 Source Server         : local-mariadb
 Source Server Type    : MySQL
 Source Server Version : 100332
 Source Host           : localhost:3306
 Source Schema         : restapi-socketio

 Target Server Type    : MySQL
 Target Server Version : 100332
 File Encoding         : 65001

 Date: 31/12/2022 15:46:33
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for assets
-- ----------------------------
DROP TABLE IF EXISTS `assets`;
CREATE TABLE `assets`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `orang_id` bigint(20) NULL DEFAULT NULL,
  `id_product` bigint(20) NULL DEFAULT NULL,
  `tittle` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `description` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `price` bigint(20) NULL DEFAULT NULL,
  `Brand` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `Category` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `thumbnail` text CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_assets_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `fk_assets_orang`(`orang_id`) USING BTREE,
  CONSTRAINT `fk_assets_orang` FOREIGN KEY (`orang_id`) REFERENCES `orang` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of assets
-- ----------------------------
INSERT INTO `assets` VALUES (1, 2, 3, 'Samsung Universe 9', 'Samsung\'s new variant which goes beyond Galaxy to the Universe', 1249, 'Samsung', 'smartphones', 'https://i.dummyjson.com/data/products/3/thumbnail.jpg', '2022-12-31 15:46:07.394', '2022-12-31 15:46:07.394', NULL);
INSERT INTO `assets` VALUES (2, 2, 7, 'Samsung Galaxy Book', 'Samsung Galaxy Book S (2020) Laptop With Intel Lakefield Chip, 8GB of RAM Launched', 1499, 'Samsung', 'laptops', 'https://i.dummyjson.com/data/products/7/thumbnail.jpg', '2022-12-31 15:46:07.394', '2022-12-31 15:46:07.394', NULL);
INSERT INTO `assets` VALUES (3, 6, 1, 'iPhone 9', 'An apple mobile which is nothing like apple', 549, 'Apple', 'smartphones', 'https://i.dummyjson.com/data/products/1/thumbnail.jpg', '2022-12-31 15:46:07.394', '2022-12-31 15:46:07.394', NULL);
INSERT INTO `assets` VALUES (4, 7, 2, 'iPhone X', 'SIM-Free, Model A19211 6.5-inch Super Retina HD display with OLED technology A12 Bionic chip with ...', 899, 'Apple', 'smartphones', 'https://i.dummyjson.com/data/products/2/thumbnail.jpg', '2022-12-31 15:46:07.394', '2022-12-31 15:46:07.394', NULL);
INSERT INTO `assets` VALUES (5, 3, 5, 'Huawei P30', 'Huawei’s re-badged P30 Pro New Edition was officially unveiled yesterday in Germany and now the device has made its way to the UK.', 499, 'Huawei', 'smartphones', 'https://i.dummyjson.com/data/products/5/thumbnail.jpg', '2022-12-31 15:46:07.394', '2022-12-31 15:46:07.394', NULL);
INSERT INTO `assets` VALUES (6, 8, 3, 'Samsung Universe 9', 'Samsung\'s new variant which goes beyond Galaxy to the Universe', 1249, 'Samsung', 'smartphones', 'https://i.dummyjson.com/data/products/3/thumbnail.jpg', '2022-12-31 15:46:07.394', '2022-12-31 15:46:07.394', NULL);
INSERT INTO `assets` VALUES (7, 9, 5, 'Huawei P30', 'Huawei’s re-badged P30 Pro New Edition was officially unveiled yesterday in Germany and now the device has made its way to the UK.', 499, 'Huawei', 'smartphones', 'https://i.dummyjson.com/data/products/5/thumbnail.jpg', '2022-12-31 15:46:07.394', '2022-12-31 15:46:07.394', NULL);
INSERT INTO `assets` VALUES (8, 9, 2, 'iPhone X', 'SIM-Free, Model A19211 6.5-inch Super Retina HD display with OLED technology A12 Bionic chip with ...', 899, 'Apple', 'smartphones', 'https://i.dummyjson.com/data/products/2/thumbnail.jpg', '2022-12-31 15:46:07.394', '2022-12-31 15:46:07.394', NULL);
INSERT INTO `assets` VALUES (9, 4, 3, 'Samsung Universe 9', 'Samsung\'s new variant which goes beyond Galaxy to the Universe', 1249, 'Samsung', 'smartphones', 'https://i.dummyjson.com/data/products/3/thumbnail.jpg', '2022-12-31 15:46:07.394', '2022-12-31 15:46:07.394', NULL);
INSERT INTO `assets` VALUES (10, 10, 7, 'Samsung Galaxy Book', 'Samsung Galaxy Book S (2020) Laptop With Intel Lakefield Chip, 8GB of RAM Launched', 1499, 'Samsung', 'laptops', 'https://i.dummyjson.com/data/products/7/thumbnail.jpg', '2022-12-31 15:46:07.394', '2022-12-31 15:46:07.394', NULL);
INSERT INTO `assets` VALUES (11, 5, 5, 'Huawei P30', 'Huawei’s re-badged P30 Pro New Edition was officially unveiled yesterday in Germany and now the device has made its way to the UK.', 499, 'Huawei', 'smartphones', 'https://i.dummyjson.com/data/products/5/thumbnail.jpg', '2022-12-31 15:46:07.394', '2022-12-31 15:46:07.394', NULL);
INSERT INTO `assets` VALUES (12, 11, 2, 'iPhone X', 'SIM-Free, Model A19211 6.5-inch Super Retina HD display with OLED technology A12 Bionic chip with ...', 899, 'Apple', 'smartphones', 'https://i.dummyjson.com/data/products/2/thumbnail.jpg', '2022-12-31 15:46:07.394', '2022-12-31 15:46:07.394', NULL);

-- ----------------------------
-- Table structure for orang
-- ----------------------------
DROP TABLE IF EXISTS `orang`;
CREATE TABLE `orang`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `nama` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `jenis_kelamin` tinyint(4) NULL DEFAULT 0 COMMENT '0 Laki-laki, 1 wanita',
  `orang_tua` bigint(20) NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_orang_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `fk_orang_anak`(`orang_tua`) USING BTREE,
  CONSTRAINT `fk_orang_anak` FOREIGN KEY (`orang_tua`) REFERENCES `orang` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of orang
-- ----------------------------
INSERT INTO `orang` VALUES (1, 'Bani', 0, NULL, '2022-12-31 15:46:07.379', '2022-12-31 15:46:07.379', NULL);
INSERT INTO `orang` VALUES (2, 'Budi', 0, 1, '2022-12-31 15:46:07.379', '2022-12-31 15:46:07.379', NULL);
INSERT INTO `orang` VALUES (3, 'Nida', 1, 1, '2022-12-31 15:46:07.379', '2022-12-31 15:46:07.379', NULL);
INSERT INTO `orang` VALUES (4, 'Andi', 0, 1, '2022-12-31 15:46:07.379', '2022-12-31 15:46:07.379', NULL);
INSERT INTO `orang` VALUES (5, 'Sigit', 0, 1, '2022-12-31 15:46:07.379', '2022-12-31 15:46:07.379', NULL);
INSERT INTO `orang` VALUES (6, 'Hari', 0, 2, '2022-12-31 15:46:07.379', '2022-12-31 15:46:07.379', NULL);
INSERT INTO `orang` VALUES (7, 'Siti', 1, 2, '2022-12-31 15:46:07.379', '2022-12-31 15:46:07.379', NULL);
INSERT INTO `orang` VALUES (8, 'Bila', 1, 3, '2022-12-31 15:46:07.379', '2022-12-31 15:46:07.379', NULL);
INSERT INTO `orang` VALUES (9, 'Lesti', 1, 3, '2022-12-31 15:46:07.379', '2022-12-31 15:46:07.379', NULL);
INSERT INTO `orang` VALUES (10, 'Diki', 0, 4, '2022-12-31 15:46:07.379', '2022-12-31 15:46:07.379', NULL);
INSERT INTO `orang` VALUES (11, 'Doni', 0, 5, '2022-12-31 15:46:07.379', '2022-12-31 15:46:07.379', NULL);
INSERT INTO `orang` VALUES (12, 'Toni', 0, 5, '2022-12-31 15:46:07.379', '2022-12-31 15:46:07.379', NULL);

SET FOREIGN_KEY_CHECKS = 1;
