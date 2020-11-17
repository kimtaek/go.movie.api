-- MySQL dump 10.13  Distrib 5.7.23, for macos10.13 (x86_64)
--
-- Host: 127.0.0.1    Database: homestead
-- ------------------------------------------------------
-- Server version	5.7.18

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `actors`
--

DROP TABLE IF EXISTS `actors`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `actors` (
                          `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                          `name` varchar(255) NOT NULL COMMENT '이름',
                          `gender` enum('MALE','FEMALE') NOT NULL COMMENT '남/여',
                          `birthday` varchar(10) DEFAULT NULL COMMENT '생년월일',
                          `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '생성 일시',
                          `updated_at` timestamp NULL DEFAULT NULL COMMENT '수정 일시',
                          `deleted_at` timestamp NULL DEFAULT NULL COMMENT '삭제 일시',
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COMMENT='배우 테이블';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `actors`
--

LOCK TABLES `actors` WRITE;
/*!40000 ALTER TABLE `actors` DISABLE KEYS */;
INSERT INTO `actors` VALUES (1,'이정재','MALE','1972-12-15','2020-11-17 21:58:31',NULL,NULL),(2,'최민식','MALE','1962-05-30','2020-11-17 21:58:31',NULL,NULL),(3,'황정민','MALE','1970-09-01','2020-11-17 21:58:31',NULL,NULL),(4,'송지효','FEMALE','1981-08-15','2020-11-17 21:58:31',NULL,NULL),(5,'정우성','MALE','1973-04-22','2020-11-17 21:58:31',NULL,NULL);
/*!40000 ALTER TABLE `actors` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `movie_actors`
--

DROP TABLE IF EXISTS `movie_actors`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `movie_actors` (
                                `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                                `movie_id` int(10) unsigned NOT NULL COMMENT '영화 ID',
                                `actor_id` int(10) unsigned NOT NULL COMMENT '배우 ID',
                                `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '생성 일시',
                                `updated_at` timestamp NULL DEFAULT NULL COMMENT '수정 일시',
                                `deleted_at` timestamp NULL DEFAULT NULL COMMENT '삭제 일시',
                                PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COMMENT='영화 배우 연결 테이블';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `movie_actors`
--

LOCK TABLES `movie_actors` WRITE;
/*!40000 ALTER TABLE `movie_actors` DISABLE KEYS */;
INSERT INTO `movie_actors` VALUES (1,1,1,'2020-11-17 22:02:13',NULL,NULL),(2,1,2,'2020-11-17 22:02:13',NULL,NULL),(3,1,3,'2020-11-17 22:02:13',NULL,NULL),(4,1,4,'2020-11-17 14:58:52',NULL,NULL),(5,2,3,'2020-11-17 14:58:45',NULL,NULL),(6,2,4,'2020-11-17 14:58:45',NULL,NULL),(7,3,5,'2020-11-17 14:58:17',NULL,NULL);
/*!40000 ALTER TABLE `movie_actors` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `movies`
--

DROP TABLE IF EXISTS `movies`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `movies` (
                          `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
                          `title` varchar(255) NOT NULL COMMENT '타이틀',
                          `description` text COMMENT '소개',
                          `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '생성 일시',
                          `created_by` int(10) unsigned DEFAULT '0',
                          `updated_at` timestamp NULL DEFAULT NULL COMMENT '수정 일시',
                          `updated_by` int(10) unsigned DEFAULT '0',
                          `deleted_at` timestamp NULL DEFAULT NULL COMMENT '삭제 일시',
                          `deleted_by` int(10) unsigned DEFAULT NULL,
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='영화 테이블';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `movies`
--

LOCK TABLES `movies` WRITE;
/*!40000 ALTER TABLE `movies` DISABLE KEYS */;
INSERT INTO `movies` VALUES (1,'신세계','세 남자가 가고 싶었던 서로 다른 신세계','2020-11-17 21:36:05',0,'2020-11-17 23:51:52',0,NULL,NULL),(2,'아수라','강력계 형사 한도경(정우성)은 이권과 성공을 위해 각종 범죄를 저지르는 악덕시장 박성배(황정민)의 뒷일을 처리해주는 대가로 돈을 받는다. 악에 계속 노출되는 사이, 말기 암 환자인 아내의 병원비를 핑계로 돈 되는 건 뭐든 하는 악인의 길로 들어서게 된 한도경. 그의 약점을 쥔 독종 검사 김차인(곽도원)과 검찰수사관 도창학(정만식)은 그를 협박하고 이용해 박성배의 비리와 범죄 혐의를 캐려 한다. 각자의 이익과 목적을 위해 한도경의 목을 짓누르는 검찰과 박성배. 그 사이 태풍의 눈처럼 되어 버린 한도경은, 자신을 친형처럼 따르는 후배 형사 문선모(주지훈)를 박성배의 수하로 들여보내고, 살아남기 위해 혈안이 된 나쁜 놈들 사이에서 서로 물지 않으면 물리는 지옥도가 펼쳐진다.','2020-11-17 22:41:00',0,'2020-11-17 22:41:00',0,NULL,NULL),(3,'강철비2: 정상회담','남북미 정상회담 중, 북한 내 쿠데타로 세 정상이 납치된다!','2020-11-17 22:43:00',0,'2020-11-17 23:55:42',0,NULL,NULL);
/*!40000 ALTER TABLE `movies` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'homestead'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-11-18  0:00:59
