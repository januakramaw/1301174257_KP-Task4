/*
SQLyog Ultimate v12.4.3 (64 bit)
MySQL - 5.6.21 : Database - colorThumbnail
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

/*Table structure for table `colorThumbnail` */

DROP TABLE IF EXISTS `colorThumbnail`;

CREATE TABLE `colorThumbnail` (
  `ColorID` VARCHAR(25) DEFAULT NULL, 
  `Color` VARCHAR(25) DEFAULT NULL,
  `Category` VARCHAR(25) DEFAULT NULL,
  `Type` VARCHAR(25) DEFAULT NULL,
  `Rgba` VARCHAR(25) DEFAULT NULL,
  `Hex` VARCHAR(30) DEFAULT NULL,
  `ThumbnailURL` VARCHAR(25) DEFAULT NULL,
  `ThumbnailHeight` INT(25) DEFAULT NULL,
  `ThumbnailWidth` INT(35) DEFAULT NULL,
  PRIMARY KEY (`ColorID`)
) ENGINE=INNODB AUTO_INCREMENT=125 DEFAULT CHARSET=utf8;

/*Data for the table `food` */

