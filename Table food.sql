/*
SQLyog Ultimate v12.4.3 (64 bit)
MySQL - 5.6.21 : Database - fooddetail
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`task4` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE `task4`;

/*Table structure for table `employees` */

DROP TABLE IF EXISTS `fooddetail`;

CREATE TABLE `fooddetail` (
  `FoodID` VARCHAR(25) DEFAULT NULL,
  `ImageHeight` INT(25) DEFAULT NULL,
  `ImageURL` VARCHAR(25) DEFAULT NULL,
  `ImageWidth` INT(25) DEFAULT NULL,
  `Name` VARCHAR(30) DEFAULT NULL,
  `ThumbnailHeight` INT(25) DEFAULT NULL,
  `ThumbnailURL` VARCHAR(25) DEFAULT NULL,
  `ThumbnailWidth` INT(35) DEFAULT NULL,
  `Type` VARCHAR(25) DEFAULT NULL,
  PRIMARY KEY (`FoodID`)
) ENGINE=INNODB AUTO_INCREMENT=125 DEFAULT CHARSET=utf8;

/*Data for the table `fooddetail` */

