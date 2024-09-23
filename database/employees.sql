/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

DROP TABLE IF EXISTS `employees`;
CREATE TABLE `employees` (
  `employee_id` varchar(50) NOT NULL,
  `fullname` varchar(50) NOT NULL,
  `address` varchar(50) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`employee_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `employees` (`employee_id`, `fullname`, `address`, `created_at`, `updated_at`) VALUES
('1001', 'Budi', 'Jakarta', '2024-09-23 11:47:51', '2024-09-23 11:47:51');
INSERT INTO `employees` (`employee_id`, `fullname`, `address`, `created_at`, `updated_at`) VALUES
('1002', 'Adi', 'Jakarta', '2024-09-23 11:48:24', '2024-09-23 11:48:24');
INSERT INTO `employees` (`employee_id`, `fullname`, `address`, `created_at`, `updated_at`) VALUES
('1003', 'Muhammad', 'Tangerang Selatan', '2024-09-23 11:48:46', '2024-09-23 12:40:32');
INSERT INTO `employees` (`employee_id`, `fullname`, `address`, `created_at`, `updated_at`) VALUES
('1004', 'Ilham', 'Bandung', '2024-09-23 12:41:06', '2024-09-23 12:41:06');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;