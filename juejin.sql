-- MySQL dump 10.13  Distrib 8.0.31, for Win64 (x86_64)
--
-- Host: localhost    Database: juejin
-- ------------------------------------------------------
-- Server version	8.0.31

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `article_counter`
--

DROP TABLE IF EXISTS `article_counter`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `article_counter` (
  `digg_count` int DEFAULT '0',
  `view_count` int DEFAULT '0',
  `collect_count` int DEFAULT '0',
  `comment_count` int DEFAULT '0',
  `article_id` int NOT NULL,
  PRIMARY KEY (`article_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `article_counter`
--

LOCK TABLES `article_counter` WRITE;
/*!40000 ALTER TABLE `article_counter` DISABLE KEYS */;
INSERT INTO `article_counter` VALUES (5,14,1,0,13),(0,0,0,2,14);
/*!40000 ALTER TABLE `article_counter` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `article_major`
--

DROP TABLE IF EXISTS `article_major`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `article_major` (
  `content` text,
  `brief_content` varchar(255) DEFAULT NULL,
  `cover` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `title` varchar(100) DEFAULT NULL,
  `category_id` varchar(19) DEFAULT NULL,
  `user_id` bigint DEFAULT NULL,
  `article_id` int NOT NULL AUTO_INCREMENT,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`article_id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `article_major`
--

LOCK TABLES `article_major` WRITE;
/*!40000 ALTER TABLE `article_major` DISABLE KEYS */;
INSERT INTO `article_major` VALUES ('后端小菜鸡做梦会梦到进红岩吗','114514','image/article/cover/d11d0a2f-a921-4cd4-8a85-9eedd41fa74c.png','逸一时，误意识','6809637769959178254',1621161939394629632,13,'2023-02-06 23:40:34'),('','114514','image/article/cover/d11d0a2f-a921-4cd4-8a85-9eedd41fa74c.png','go是世界上最好的编程语言','6809637769959178254',1621161939394629632,14,'2023-02-07 01:26:28');
/*!40000 ALTER TABLE `article_major` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `category`
--

DROP TABLE IF EXISTS `category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `category` (
  `category_id` varchar(19) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `category_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`category_id`),
  UNIQUE KEY `category_id` (`category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category`
--

LOCK TABLES `category` WRITE;
/*!40000 ALTER TABLE `category` DISABLE KEYS */;
INSERT INTO `category` VALUES ('6809635626661445640','IOS'),('6809635626879549454','Android'),('6809637767543259144','前端'),('6809637769959178254','后端'),('6809637771511070734','开发工具'),('6809637772874219534','阅读'),('6809637773935378440','人工智能'),('6809637776263217160','代码人生');
/*!40000 ALTER TABLE `category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `collection_select_articles`
--

DROP TABLE IF EXISTS `collection_select_articles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `collection_select_articles` (
  `id` int NOT NULL AUTO_INCREMENT,
  `collection_id` bigint NOT NULL,
  `article_id` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `collection_id` (`collection_id`) USING BTREE,
  CONSTRAINT `collection_select_articles_ibfk_1` FOREIGN KEY (`collection_id`) REFERENCES `collection_set` (`collection_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `collection_select_articles`
--

LOCK TABLES `collection_select_articles` WRITE;
/*!40000 ALTER TABLE `collection_select_articles` DISABLE KEYS */;
INSERT INTO `collection_select_articles` VALUES (15,1620069447278530560,13),(16,1624280846963838976,13);
/*!40000 ALTER TABLE `collection_select_articles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `collection_set`
--

DROP TABLE IF EXISTS `collection_set`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `collection_set` (
  `collection_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `collection_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '',
  `permisssion` tinyint(1) NOT NULL,
  `create_time` datetime NOT NULL,
  `update_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  `post_article_count` int unsigned NOT NULL DEFAULT '0',
  KEY `collection_id` (`collection_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `collection_set`
--

LOCK TABLES `collection_set` WRITE;
/*!40000 ALTER TABLE `collection_set` DISABLE KEYS */;
INSERT INTO `collection_set` VALUES (1620069447278530560,1619766784787746816,'圆今飞题什','们长律眼集场想层克局技千。业给论观己个教好业果强率高。林山称证里质级风率样因价时。制广部据几酸花效质色来你如委确。',0,'2023-01-30 22:40:30','2023-02-03 01:11:24',1),(1620070077179105280,1619766784787746816,'四值工','收该现队儿价白品看关运部可。阶世际准始据家土调制转边实到意东最。速习照北性也真厂明管门式能新节江。平约大东会么二决阶需高周列。位需化所须克的省容素上名省任子利状回。具原度西联这光写队干米联认所六。',1,'2023-01-30 22:43:01','2023-01-30 22:43:01',0),(1624280846963838976,1621161939394629632,'正持现式离老','作约上年西山音高音积之近。手受务平我院今放究改直位青发。大准都深得表现亲志装周非儿。清历以党边式民方程受交争算。京往地张明华品事到往水集程党须五。着例王示生机周点高图住日干。分受常统群律年便意到张劳员。',1,'2023-02-11 13:35:06','2023-02-11 13:35:47',1),(1624292061287026688,1621161939394629632,'样不或院上','马支参公程成三县价包强形象。万议处国阶地史金报交上委放然必角。明就派则党深划集前也断白起眼争农。',1,'2023-02-11 14:19:40','2023-02-11 14:19:40',0);
/*!40000 ALTER TABLE `collection_set` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `comment`
--

DROP TABLE IF EXISTS `comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `comment` (
  `comment_id` int NOT NULL AUTO_INCREMENT,
  `comment_content` varchar(255) DEFAULT NULL,
  `digg_count` int DEFAULT '0',
  `reply_count` int DEFAULT '0',
  `create_time` datetime DEFAULT NULL,
  `item_id` int DEFAULT NULL,
  `item_type` int DEFAULT NULL,
  `user_id` bigint NOT NULL,
  PRIMARY KEY (`comment_id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment`
--

LOCK TABLES `comment` WRITE;
/*!40000 ALTER TABLE `comment` DISABLE KEYS */;
INSERT INTO `comment` VALUES (5,'enim sed',6,1,'2023-02-07 20:34:26',14,2,1621161939394629632),(7,'6',0,0,'2023-02-11 14:20:29',14,2,1621161939394629632);
/*!40000 ALTER TABLE `comment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `digg`
--

DROP TABLE IF EXISTS `digg`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `digg` (
  `user_id` bigint NOT NULL,
  `item_id` int NOT NULL,
  `item_type` varchar(255) NOT NULL,
  `id` int NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `digg`
--

LOCK TABLES `digg` WRITE;
/*!40000 ALTER TABLE `digg` DISABLE KEYS */;
INSERT INTO `digg` VALUES (1621161939394629632,13,'2',3),(1621161939394629632,5,'5',4),(1621161939394629632,6,'6',5),(1621161939394629632,7,'6',6),(1621161939394629632,9,'6',7);
/*!40000 ALTER TABLE `digg` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `draft`
--

DROP TABLE IF EXISTS `draft`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `draft` (
  `content` text,
  `brief_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '',
  `cover` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '',
  `category_id` varchar(19) NOT NULL,
  `draft_id` int NOT NULL AUTO_INCREMENT,
  `user_id` bigint DEFAULT NULL,
  `create_time` datetime NOT NULL,
  `update_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`draft_id`)
) ENGINE=InnoDB AUTO_INCREMENT=35 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `draft`
--

LOCK TABLES `draft` WRITE;
/*!40000 ALTER TABLE `draft` DISABLE KEYS */;
INSERT INTO `draft` VALUES ('后端小菜鸡做梦会梦到进红岩吗','114514','image/article/cover/d11d0a2f-a921-4cd4-8a85-9eedd41fa74c.png','逸一时，误意识','6809637769959178254',32,1621161939394629632,'2023-02-07 18:22:18','2023-02-07 18:22:18'),('后端小菜鸡做梦会梦到进红岩吗','114514','image/article/cover/d11d0a2f-a921-4cd4-8a85-9eedd41fa74c.png','逸一时，误意识','6809637769959178254',33,1621161939394629632,'2023-02-11 13:31:03','2023-02-11 13:31:03'),('后端小菜鸡做梦会梦到进红岩吗','114514','image/article/cover/d11d0a2f-a921-4cd4-8a85-9eedd41fa74c.png','逸一时，误意识','6809637769959178254',34,1621161939394629632,'2023-02-11 14:18:24','2023-02-11 14:18:24');
/*!40000 ALTER TABLE `draft` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `follow`
--

DROP TABLE IF EXISTS `follow`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `follow` (
  `follower` bigint DEFAULT NULL,
  `followee` bigint DEFAULT NULL,
  `id` int NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `follow`
--

LOCK TABLES `follow` WRITE;
/*!40000 ALTER TABLE `follow` DISABLE KEYS */;
INSERT INTO `follow` VALUES (1621161939394629632,1619766784787746816,10);
/*!40000 ALTER TABLE `follow` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `item_tag`
--

DROP TABLE IF EXISTS `item_tag`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `item_tag` (
  `tag_id` varchar(19) NOT NULL,
  `item_id` varchar(19) NOT NULL,
  `item_type` int NOT NULL,
  `id` int NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=116 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `item_tag`
--

LOCK TABLES `item_tag` WRITE;
/*!40000 ALTER TABLE `item_tag` DISABLE KEYS */;
INSERT INTO `item_tag` VALUES ('6809640408797167623','6809637769959178254',8,95),('6809640445233070094','6809637769959178254',8,96),('7147583745896218628','6809637769959178254',8,97),('6809640364677267469','6809637769959178254',8,98),('6809640398105870343','6809637767543259144',8,99),('6809640407484334093','6809637767543259144',8,100),('6809640364677267469','11',2,106),('6809640364677267469','12',2,107),('6809640364677267469','13',2,108),('6809640364677267469','14',2,109),('6809640364677267469','32',4,110),('6809640364677267469','32',4,111),('6809640364677267469','33',4,112),('6809640364677267469','33',4,113),('6809640364677267469','34',4,114),('6809640364677267469','34',4,115);
/*!40000 ALTER TABLE `item_tag` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `reply`
--

DROP TABLE IF EXISTS `reply`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `reply` (
  `reply_id` int NOT NULL AUTO_INCREMENT,
  `reply_comment_id` int DEFAULT NULL,
  `reply_content` varchar(255) DEFAULT NULL,
  `user_id` bigint DEFAULT NULL,
  `item_id` int DEFAULT NULL,
  `item_type` int DEFAULT NULL,
  `digg_count` int DEFAULT '0',
  `create_time` datetime DEFAULT NULL,
  `parent_reply_id` int NOT NULL DEFAULT '0',
  `reply_user_id` bigint DEFAULT '0',
  PRIMARY KEY (`reply_id`),
  KEY `parent_reply_id` (`parent_reply_id`),
  KEY `reply_comment_id` (`reply_comment_id`),
  CONSTRAINT `reply_ibfk_2` FOREIGN KEY (`reply_comment_id`) REFERENCES `comment` (`comment_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `reply`
--

LOCK TABLES `reply` WRITE;
/*!40000 ALTER TABLE `reply` DISABLE KEYS */;
INSERT INTO `reply` VALUES (9,5,'kokodayo',1621161939394629632,13,2,3,'2023-02-11 14:20:51',0,1621161939394629632);
/*!40000 ALTER TABLE `reply` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tag_info`
--

DROP TABLE IF EXISTS `tag_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `tag_info` (
  `tag_id` varchar(19) NOT NULL,
  `tag_name` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `create_time` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`tag_id`),
  UNIQUE KEY `tag_id` (`tag_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tag_info`
--

LOCK TABLES `tag_info` WRITE;
/*!40000 ALTER TABLE `tag_info` DISABLE KEYS */;
INSERT INTO `tag_info` VALUES ('6809640357354012685','React.js','2023-02-06 22:25:10'),('6809640364677267469','Go','2023-02-06 22:22:25'),('6809640369764958215','Vue.js','2023-02-06 22:24:44'),('6809640398105870343','JavaScript','2023-02-06 22:23:58'),('6809640407484334093','前端','2023-02-06 22:23:41'),('6809640408797167623','后端','2023-02-06 22:21:00'),('6809640445233070094','Java','2023-02-06 22:21:18'),('6809640501776482317','架构','2023-02-06 22:22:51'),('6809640600502009863','数据库','2023-02-06 22:23:11'),('7147583745896218628','掘金·日新计划','2023-02-06 22:22:00');
/*!40000 ALTER TABLE `tag_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_auth`
--

DROP TABLE IF EXISTS `user_auth`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_auth` (
  `id` bigint NOT NULL,
  `username` varchar(20) NOT NULL,
  `password` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `email` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `github_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_auth`
--

LOCK TABLES `user_auth` WRITE;
/*!40000 ALTER TABLE `user_auth` DISABLE KEYS */;
INSERT INTO `user_auth` VALUES (1619766784787746816,'lwj','$2a$10$XbcAsaJc13FAhVnFvOw2OeQ7noHlT23ApL1qENTecES3u/UgdR1vG','liangweijian666@outlook.com','18167577178','2023-01-30 02:37:50','2023-01-30 02:37:50',NULL),(1621161939394629632,'马勇','$2a$10$e0JSlb4lOXHCcstLMWY16O7/2k6ACWFq4fU6yXLTZMrp6P//95mj6','1960441553@qq.com','18193596329','2023-02-02 23:01:41','2023-02-02 23:01:41',NULL);
/*!40000 ALTER TABLE `user_auth` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_basic`
--

DROP TABLE IF EXISTS `user_basic`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_basic` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '',
  `company` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '',
  `job_title` varchar(60) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '',
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `user_id` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_basic`
--

LOCK TABLES `user_basic` WRITE;
/*!40000 ALTER TABLE `user_basic` DISABLE KEYS */;
INSERT INTO `user_basic` VALUES (36,1619766784787746816,'http://dummyimage.com/100x100','','必看铁而教总资受我只管信局。们西支立九价整料须严身然查心的明好。验水然已示及西或克电如代。','般究火细身','冬之花'),(37,1621161939394629632,'http://dummyimage.com/100x100','','多斗价京难方使厂门明政影可人。少将国社图战美党验技议究。义工标热利速圆这阶率海公联群道平用。题调时社件始一花易山却商展路管片连。计列如拉西织明受六人转近什接千半统。三过前原期上中马求知华矿产。','经通音器过确','先辈');
/*!40000 ALTER TABLE `user_basic` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_counter`
--

DROP TABLE IF EXISTS `user_counter`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user_counter` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `digg_article_count` int DEFAULT '0',
  `digg_shortmsg_count` int DEFAULT '0',
  `followee_count` int DEFAULT '0',
  `follower_count` int DEFAULT '0',
  `got_digg_count` int DEFAULT '0',
  `got_view_count` int DEFAULT '0',
  `post_article_count` int DEFAULT '0',
  `post_shortmsg_count` int DEFAULT '0',
  `select_online_course_count` int DEFAULT '0',
  `collection_set_count` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_counter`
--

LOCK TABLES `user_counter` WRITE;
/*!40000 ALTER TABLE `user_counter` DISABLE KEYS */;
INSERT INTO `user_counter` VALUES (40,1619766784787746816,0,0,0,1,11,11,0,0,0,2),(41,1621161939394629632,14,0,1,0,14,13,0,0,0,3);
/*!40000 ALTER TABLE `user_counter` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-03-10 23:11:30
