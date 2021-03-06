-- MySQL dump 10.13  Distrib 8.0.28, for Win64 (x86_64)
--
-- Host: localhost    Database: dna_matching
-- ------------------------------------------------------
-- Server version	8.0.28

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
-- Current Database: `dna_matching`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `dna_matching` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `dna_matching`;

--
-- Table structure for table `disease`
--

DROP TABLE IF EXISTS `disease`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `disease` (
  `id_disease` int NOT NULL AUTO_INCREMENT,
  `disease_name` varchar(30) NOT NULL,
  `dna_sequence` varchar(100) NOT NULL,
  PRIMARY KEY (`id_disease`),
  UNIQUE KEY `id_disease_UNIQUE` (`id_disease`),
  UNIQUE KEY `disease_name_UNIQUE` (`disease_name`)
) ENGINE=InnoDB AUTO_INCREMENT=76 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `disease`
--

LOCK TABLES `disease` WRITE;
/*!40000 ALTER TABLE `disease` DISABLE KEYS */;
INSERT INTO `disease` (`id_disease`, `disease_name`, `dna_sequence`) VALUES (1,'warfarin sodium','TGGGAGGTAATAGAGGAAGTCTCGACCATGGTCTTAAATATGCCAAACAT'),(2,'TITANIUM DIOXIDE','AGTAATCTGGCGACTTCGGAGTAGCTCGTCTGGAGGAAGGCTGTCGACAACG'),(3,'Pramipexole','CACTTGAGGCAATAATCCCGTCTAAGACTGATCTTAATT'),(4,'Loperamides','CCAGCGTTCTATAATCTGTTGGATTTCTTCCAGATTAGTCTTAAAATGCATGGATAATCTACGCTGCTCGGCTCATTCGCTGTAAGGGTCCCCTAC'),(5,'ethyl alcohol','GTGGAAAAATAGACCGGAAATGAATTTTGGACGACGCAACAATGCGGTACCCC'),(6,'Bupropion Hydrochloride','AGACTGGGGAAAGCATCGCCTGAAGGATATGGTAGG'),(7,'TRIFLUOPERAZIN','CCTCTTTTGGCGTCGCATCAGAAAGTATCTAACTGGAGGCGAGGGTACTACCTTAGGGTTCGTGGTGGGAAGAGCATAGTGCGCATTACTGGTCCAC'),(8,'PREDNISONE','TCATTAGATACGCC'),(9,'Walnut English Pollen','TGCATTGTGCCAATTTT'),(10,'Chlorhexidine Gluconate','GTTTAGCAGCACGTTACGCTTTTCTGATCTATACATTTCCTGTAATAGTAGGT'),(11,'Hydroxyzine','ACACGATGTCCAAGGTATCCCCCACCCCGAACTCCCAACGCTTAGGGGGAAAATCTTGGGTACCTCTTGTACGTCC'),(12,'MUPIROCIN','CATTACGTACATATTGCTAGCACGTGTCTTTCATCCTACTATGCCTGAACG'),(13,'Furosemide','ATGTACCTTGCAAGTAAACCTTTCAGAGTCATTGCGCCTTCACAGTAATGATGGCCTCGATT'),(14,'lamivudin','CCAGGTATCTACCCACCAGTAAGCGGGCAGCGTA'),(15,'Spironolactone','GCTAGCGCTAGGGGCAGGTAATCGTGGTCCAAGGTTGGACGTTTACAGTGACTGACTTCTTGGGACACACGCGACCGGTTAGCCTCCGTTCAATA'),(16,'Simvastatin','CGGTCATCGGTCGCTCTACGGAAGAGTGAGCAACTCGAACATTGACGGATACTGGCTACTCTGGGCGCCGTAGCACG'),(17,'CEFTAZIDIME','TTGGAGGAAAGACCAGACAGGCGGTTATGGCACCGCTAGAGTCCG'),(18,'triamcinolone acetonide','ACGGTACACGTCCCTGCCGCTTCCGAAATTAACTACCAAATCGTCTCTTGTCCTACATTTAGGGAGTTACGT'),(19,'Triclosan','AAGGGTATTGGACATTTGCCCACTGACTCATCGCGTGCTCTTAGGAAATAGATGAGCTAGCAACCGCCGCCGAGGGTTAATG'),(20,'Torsemide','TTAGGGGATTGCCAAGAGCACGCCATGGGATGGTTCCTTCTATGGCGACAGCGGAGTTTTGCGCG'),(21,'Titanium Dioxide Ferguso','GACTCGGGACATCTAGTACGTAAATCGCAGGGTAGTACCGTAATCAAAGGACTAAATCATTCAGGCGGCCTGCTGCTTGTCAGGTA'),(22,'Minoxidil','GGCCCGTATGCATGGTCTCTAAAGGGTACCA'),(23,'Aspirin','CCACGGATAATTGGATCCTTGCGGGCATTTTTGGTTAGCTCCAGCTCATGCCTGGGGGAAGGC'),(24,'Tobramycin','GCTGGTAAGGTATAAATGCGC'),(25,'Sulfamethoxazole','AGGTGTCCGCTATTCCCACCCTCTGCGTGTACTATAAAGCTGTAGGAACTGCGCTCGTTTCTAGCTGTAACGCGGAGTTGTCACGAC'),(26,'Nateglinide','GCTTCTCCACTAATGGTTATGGAGTCTAGGAACACCACCTGCGCAACCTTGCATCTCCCGGCCGC'),(27,'Potassium Chloride','ACTTTAGAGGTTGGATGGCAGCCTGGTCGGAGCGGTTAATCTGCAGGACCGACACTGTCACGGTACTCAGACATCGCGAACGGGA'),(28,'Honey Bee hymenoptera','GGGTGACAGATTCTGCCGTGTAACTGCTAAAGGAGCCTCCCAGCAC'),(29,'Sodium Fluoride','AAGGATTGGG'),(30,'Italian Rye Grass','CACCTGACCGGTAAAGACTTAGATGAAGTACGGCCTTGCACCAGCACTGCTCCAAGGAGATTAGCGTCATCATTAAGATTTTCCG'),(31,'Avobenzone','GTGGGGAATCTTGTTAAAATTGCATG'),(32,'ACONITUM NAPELLUS','TCCGATGGCAAATATAGCTGTCATGGACAAGCTGTGGTACGCATCTCTTCTAAGGCCCAATTCAAAACCCGGAGCATG'),(33,'Pravastatin Sodium','TATTTCGTTAGTCCGGGGGACAGAGGGAATGCCCACGTAGGACTAACGGTTAAATATAA'),(34,'ziprasidone hydrochloride','TTTCCGCCAAGGTGAAGAATGGACGTTGTTAATCTGCAACTTAGCACTGAGCTGCCTG'),(35,'amlodipine besylate','CCTGCCGACACGCGTATTCACTGGCCACTCGC'),(36,'ARNICA WHOLE PLANT','CGTATGCGTTTCGGCTAAACCAATCAGGAGATTGACTGGTCGGTTCCTTACGGAGCATCG'),(37,'Titanium Dioxide Roxanne','GAACTCTCACCCTCGAATTCTGAACCCGTCGGCTTATTTCCTAAGTGTCCTTTGAATGAACT'),(39,'Arnica montana','TCCAGATATGGGCCCACGGAAAGGAGCGCTTATGCGCATCCCGGGTTTAAACATCCCGCAGCTCAGCTGCACCGAGAATGCGCACCCCATGAT'),(40,'HOMOSALATE','AAGCGCGTGCTCGTTAGTGAGTGCAGGTCACCATGGGCACTACTACGGAGCTTTAGCGTGATGGAAATGGTTTTGCACCT'),(41,'choriogonadotropin alfa','ATTAGGCACGGAAAACACTTAACCCCTTGGCTCGATTGACTCTCCTTAGGCGGAATTCTAGCCGTGCTGTCAATATCGGTTCAAAATTCTTAACT'),(42,'Urea','AAGAAAAACGCTACACACAGCGGCGTAACTTCTAT'),(43,'Ibuprofen','AAGACATTTCGAACAGTTACCTCACCTCATCTCTGCTCGTTTGGAATTGAATAGCCCTTACACCCG'),(44,'Hydroquinone','TTGTGCAAGCAGTTCATAAGGATGCTAACTTACCCGGCCCCGGCCC'),(45,'Acetaminophen','CGGCCAATAGAGTT'),(46,'CARBAMIDE PEROXIDE','GCAGTTCGTATCGCTGCAGTCGTCTAGTGGTGGCGTAGGAAAGGCGTGTAGGCGCCACTCTTTCACGTAATATGACGGACA'),(47,'Octinoxate','CCCCCTGCCTCAAATCAAATTCTCAGTCATGGCATCGTGAAATTCCGAGCGCCACCGATGAGTTCGAAACCGAAGAAGTGTTCC'),(48,'Salicylic Acid','AAAAGGCACATTGTTTCAGATACGCATA'),(49,'Octinoxate ARNICA','TATGGTGCAGGGCGAGCGACTAAAGGCGAGGATCTTCGGTAAATTTCTTCAGTTGGTCCTTTGAACTGCGTTCGACGTGTTTAGGTATCAGTGAT'),(50,'losartan potassium','TGCGACTAAGCTAAATTCGATCGCATAGACTTTTAGCAGGACGGGTACTCGTGAACACTATAATGCACCGTAATCC'),(51,'Levetiracetam','ATAGACCTGGGATGATTTCCGCGGACGACCGACTGTATACAT'),(52,'Carisoprodol','GGAAGTTCTACCTACGTCACG'),(53,'Alendronate Sodium','GGTTATACTCTCCGTCACTAAAACAACTTTGTCACTGCCCG'),(54,'Venlafaxine Hydrochloride','GAGGGCCGAAAACGTTCGC'),(55,'Salsalate','CGTCACAAATGAGCCCTTAGGCTCGACCAAGATACTAGACATTCAATACAATCACATACGCCTAGGCCCGAGAACGGGTAAA'),(56,'Avobenzo','CACTGACCGTTGGACATCTAGTTCTTGGGGCAGGGGAGTACGTTTCTACTCGAACGCTGCATATTGG'),(57,'Methylprednisolone Acetate','GGAACATAGGCATTTGGACAAACGTACCGTTCTAGACTTGCCTTACTTGGTACTCGAGAATG'),(58,'Bacitracin Zinc','TGGCCGTGACGTGGACTATGGGCTGAGATAGAGCGTCGTTCTACCTGTATTGAATATGCGTTCGGAAAGC'),(59,'Meperidine Hydrochloride','AAGCGCTAGAACCTTAGGCAGGCCCAGCGCACCAATCTGTAACGCAACGTCGCACTCGGAAA'),(60,'NALOXONE HYDROCHLORIDE','GTTTTTTCCAG'),(61,'PANTOPRAZOLE SODIUM','ACGGTTGAATGTGCTCAGGACCAAAATACACCTACTTCGGTAGGGCCGCACCAACTGCAAGGGTA'),(62,'borage oil','TGGCCCTTGCCGCCTTCTCTAGATGACGAGTTATGCGTAGGAGTGACATTGTATGTAGGTAGCCAGCGTCT'),(63,'Fibrinogen Huma','CATTATCTCCGTTGA'),(64,'Tadalafil','CCAAACGTGCGATCTCTCGTGGTCGTTGTCTTAAGTACACTTCGTCATACCAGATATGGATCTTACTTGCGGTTGTGCTATTCGGC'),(65,'clobetasol propionate','AAGTCCTAACAATCATTGCATATTATCCACTAGAGTCGGAGAAAGAGAATCGCATTTATGAGATCCACC'),(66,'Potassium Chlori','AATTTGGTTCGGCCGGTGGGGTCCGGGCT'),(67,'Lidocaine Hydrochloride','GAGAAGCCAAACTACGTAATCGATAATAGAACGAATCAAAATTGA'),(68,'Hydroquinone Arnica','TCTACAACATGTGGGCAGTCCGCGGATGCTGCCC'),(69,'ERYTHROMYCIN','GTTTCTGTAACCGACAGATATA'),(70,'piroxicam','ATTACTACCGAGCTCGGTAAATTTGTTGGGCCGCC'),(71,'Epicoccum nigrum','TAGATCTAGCAACCAACTGCCAGGCCACGGCGAACTGCGAGAGATGTGCCGCTAGCCTGTGTCCATAACCT'),(72,'silver sulfadiazine','TGCCAAGAATGGCACGATGCGCTGCCTACTTGTAAGCTATGCCGATATTCACGCCCCTGCA'),(73,'Nicotine Polacrilex','GGAACCCGCATTTCCTTACTCGCCGTATAA'),(74,'HYDROQUINONE besylate','TTAAATCGTGAAGCGGACGTACAAAGTCCTGAAGGAAACCTC'),(75,'clobetasol propionate ACONITUM','GGCCTTCTATTTCGCCATTGCTCCGGGATAGGGATCGACCGAGCCAGGCGAAATTAGGCACCGGCCGTACTT');
/*!40000 ALTER TABLE `disease` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `prediction_result`
--

DROP TABLE IF EXISTS `prediction_result`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `prediction_result` (
  `id_prediction` int NOT NULL AUTO_INCREMENT,
  `prediction_date` date NOT NULL,
  `patient_name` varchar(45) NOT NULL,
  `id_disease` int NOT NULL,
  `prediction_status` tinyint NOT NULL,
  PRIMARY KEY (`id_prediction`),
  UNIQUE KEY `id_prediction_UNIQUE` (`id_prediction`),
  KEY `id_disease_foreign_idx` (`id_disease`),
  CONSTRAINT `id_disease_foreign` FOREIGN KEY (`id_disease`) REFERENCES `disease` (`id_disease`)
) ENGINE=InnoDB AUTO_INCREMENT=76 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `prediction_result`
--

LOCK TABLES `prediction_result` WRITE;
/*!40000 ALTER TABLE `prediction_result` DISABLE KEYS */;
INSERT INTO `prediction_result` (`id_prediction`, `prediction_date`, `patient_name`, `id_disease`, `prediction_status`) VALUES (1,'2022-04-24','Nicholas King',25,1),(2,'2022-02-12','Robin Brown PhD',67,0),(3,'2022-04-09','Nicholas Simpson',74,0),(4,'2022-01-28','Jennifer Sanchez',69,1),(5,'2022-02-24','Anthony Allen',14,0),(6,'2022-01-28','Jennifer Leonard',67,1),(7,'2022-01-19','Jessica Mcintosh',75,1),(8,'2022-02-18','Victor Roberts',25,1),(9,'2022-04-14','Dr. Amber Guzman PhD',50,1),(10,'2022-01-29','Laura Fitzpatrick',33,0),(11,'2022-03-02','Samuel Pierce DVM',28,0),(12,'2022-01-20','Russell Schwartz',75,0),(13,'2022-01-15','Alison Guerra',39,0),(14,'2022-03-06','Kelly Rhodes',31,0),(15,'2022-04-02','Rebecca Howard',72,1),(16,'2022-04-19','Abigail Butler',63,1),(17,'2022-01-20','Jason Hansen',74,0),(18,'2022-03-10','Laurie Wright',9,1),(19,'2022-03-26','Ashley Garcia',14,1),(20,'2022-03-03','Daniel Lynch',34,0),(21,'2022-02-09','Ross Valentine',53,1),(31,'2022-03-19','Tiffany Webb',55,1),(32,'2022-01-01','John Solis',21,1),(33,'2022-04-06','Dean Harris',70,1),(34,'2022-03-06','Dr. Ronald Smith',39,1),(35,'2022-02-17','David Smith',46,1),(36,'2022-03-05','Sharon Delgado',46,0),(37,'2022-01-13','Daniel Reed',64,1),(38,'2022-02-07','Regina Hall',60,0),(39,'2022-04-21','Ann Nguyen',9,0),(40,'2022-02-13','Harry Flowers',18,0),(41,'2022-02-16','Miss Lisa Valdez',56,0),(42,'2022-01-28','Jeremy Dominguez',1,0),(43,'2022-04-16','Daniel Craig',29,0),(44,'2022-03-28','Keith Goodman',14,0),(45,'2022-01-18','Derek Miller',47,1),(46,'2022-01-16','Jacob Garrett',68,0),(47,'2022-03-04','Adrian Benton',57,0),(48,'2022-04-09','Patricia Miller',71,1),(49,'2022-04-06','Natalie Cochran',6,1),(50,'2022-01-22','Kelly Wood',52,1),(51,'2022-02-28','Thomas Phillips',45,1),(52,'2022-04-18','Jennifer Mcdonald',49,1),(53,'2022-02-26','Beth Martinez',35,1),(54,'2022-01-25','Patrick Ballard',17,1),(55,'2022-04-17','Larry White',36,0),(56,'2022-03-06','Rachel Raymond',70,0),(57,'2022-04-15','Jennifer Sweeney',49,1),(58,'2022-02-09','Meredith Deleon',13,1),(59,'2022-04-17','Colton Martinez',57,0),(60,'2022-02-25','Bryce Stokes',69,0),(61,'2022-03-28','Kevin Montgomery',21,0),(62,'2022-01-18','Steven Edwards',44,0),(63,'2022-03-16','Jennifer Vargas',49,1),(64,'2022-03-03','Dawn Harris',74,1),(66,'2022-02-17','Justin Velez',36,1),(67,'2022-02-26','Sharon Williams',21,0),(68,'2022-02-05','Taylor Neal',7,1),(69,'2022-02-06','James Lara',8,0),(70,'2022-02-06','Megan Mitchell',58,1),(71,'2022-01-07','Jeffrey Garner',21,1),(72,'2022-02-20','Amy Fowler',61,0),(73,'2022-04-03','Anthony Schmidt',34,1),(74,'2022-01-14','Linda Raymond',7,1),(75,'2022-01-18','Derek Stewart',23,1);
/*!40000 ALTER TABLE `prediction_result` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-04-25  9:48:27
