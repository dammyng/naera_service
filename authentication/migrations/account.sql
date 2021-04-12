-- -------------------------------------------------------------
-- TablePlus 3.12.4(360)
--
-- https://tableplus.com/
--
-- Database: naeraauth
-- Generation Time: 2021-04-12 03:58:41.1180
-- -------------------------------------------------------------
/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
DROP TABLE IF EXISTS `accounts`;
CREATE TABLE `accounts` (
  `id` varchar(255) NOT NULL,
  `user_name` longtext,
  `first_name` longtext,
  `surname` longtext,
  `email` longtext,
  `phone_number` longtext,
  `dob` bigint DEFAULT NULL,
  `ref_code` longtext,
  `photo` longtext,
  `bvn` longtext,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;
INSERT INTO
  `accounts` (
    `id`,
    `user_name`,
    `first_name`,
    `surname`,
    `email`,
    `phone_number`,
    `dob`,
    `ref_code`,
    `photo`,
    `bvn`
  )
VALUES
  (
    'd1a4e8f3-925e-444b-9195-4062742815b1',
    '',
    'first',
    'last',
    'email',
    '',
    '0',
    '',
    '',
    ''
  );
  /*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
  /*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
  /*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
  /*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
  /*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
  /*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
  /*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;