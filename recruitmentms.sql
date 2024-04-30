-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Apr 30, 2024 at 09:39 AM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `recruitmentms`
--

DELIMITER $$
--
-- Procedures
--
CREATE DEFINER=`root`@`localhost` PROCEDURE `add_applicants` (IN `position_id` INT(3), IN `first_name` VARCHAR(255), IN `middle_name` VARCHAR(255), IN `last_name` VARCHAR(255), IN `extension_name` VARCHAR(255), IN `birthdate` DATE, IN `age` INT(3), IN `present_address` VARCHAR(255), IN `highest_education` VARCHAR(255), IN `email_address` VARCHAR(255), IN `facebook_link` VARCHAR(255), IN `bpo_exp` VARCHAR(255), IN `shift_sched` VARCHAR(255), IN `work_report` VARCHAR(255), IN `work_site_location` VARCHAR(255), IN `platform_id` INT(3), IN `ref_full_name` VARCHAR(255), IN `ref_company` VARCHAR(255), IN `ref_position` VARCHAR(255), IN `ref_contact_num` VARCHAR(255), IN `ref_email` VARCHAR(255), IN `applicant_cv` VARCHAR(255), IN `applicant_portfolio_link` VARCHAR(255), IN `createdAt` VARCHAR(255))   BEGIN
INSERT INTO applicants_data (
position_id, 
first_name, 
middle_name, 
last_name, 
extension_name, 
birthdate, 
age, 
present_address, 
highest_education, 
email_address, 
facebook_link, 
bpo_exp, 
shift_sched, 
work_report, 
work_site_location, 
platform_id, 
ref_full_name, 
ref_company, 
ref_position, 
ref_contact_num, 
ref_email, 
applicant_cv, 
applicant_portfolio_link,
createdAt) 
VALUES 
(
position_id, 
first_name, 
middle_name, 
last_name, 
extension_name, 
birthdate, 
age, 
present_address, 
highest_education, 
email_address, 
facebook_link, 
bpo_exp, 
shift_sched, 
work_report, 
work_site_location, 
platform_id, 
ref_full_name, 
ref_company, 
ref_position, 
ref_contact_num, 
ref_email, 
applicant_cv, 
applicant_portfolio_link,createdAt
);
END$$

CREATE DEFINER=`root`@`localhost` PROCEDURE `add_job_slot` (IN `position_name` VARCHAR(255), IN `department_id` INT(3), IN `position_status` VARCHAR(255), IN `available_slot` INT(3))   BEGIN
INSERT INTO job_position(position_name, department_id, position_status, available_slot)
VALUES(position_name, department_id, position_status, available_slot);
END$$

CREATE DEFINER=`root`@`localhost` PROCEDURE `add_user_accounts` (IN `username` VARCHAR(255), IN `password` VARCHAR(255), IN `user_role` VARCHAR(255), IN `user_name` VARCHAR(255), IN `user_position` INT(3), IN `department_id` INT(3), IN `createdAt` VARCHAR(255))   BEGIN
INSERT INTO user_accounts (username, password, user_role, user_name, user_position, department_id, createdAt)
VALUES (username, password, user_role, user_name, user_position, department_id, createdAt);
END$$

CREATE DEFINER=`root`@`localhost` PROCEDURE `fetch_applicants_data` ()   BEGIN
	SELECT * FROM applicants_data;
END$$

CREATE DEFINER=`root`@`localhost` PROCEDURE `fetch_application_status` ()   BEGIN 
SELECT application_status.applicant_id,applicants_data.first_name, applicants_data.middle_name, applicants_data.last_name,
	job_position.position_name,
	department.department_name,
	application_status_list.application_status_name,
	user_accounts.user_name
	FROM application_status
	INNER JOIN applicants_data ON application_status.applicant_id = applicants_data.applicant_id
	INNER JOIN application_status_list ON application_status.application_status_id = application_status_list.application_status_id
	INNER JOIN job_position ON applicants_data.position_id = job_position.position_id
	INNER JOIN department ON job_position.department_id = department.department_id
	INNER JOIN user_accounts ON job_position.department_id = user_accounts.department_id;
END$$

CREATE DEFINER=`root`@`localhost` PROCEDURE `fetch_platform_data` ()   BEGIN
SELECT job_posting_platform.platform_id, job_posting_platform.platform_name, 
	COUNT(*) FROM applicants_data
	INNER JOIN Job_posting_platform ON job_posting_platform.platform_id = applicants_data.platform_id
	GROUP BY platform_name;
END$$

CREATE DEFINER=`root`@`localhost` PROCEDURE `fetch_posting_data` ()   BEGIN
SELECT job_posting_platform.platform_id, job_posting_platform.platform_name, 
	COUNT(*) FROM applicants_data
	INNER JOIN Job_posting_platform ON job_posting_platform.platform_id = applicants_data.platform_id
	GROUP BY platform_name;
END$$

CREATE DEFINER=`root`@`localhost` PROCEDURE `fetch_user_accounts` ()   BEGIN
SELECT * FROM user_accounts;
END$$

CREATE DEFINER=`root`@`localhost` PROCEDURE `manager_applicants_data` (IN `account_id` INT(3))   BEGIN
SELECT applicants_data.* FROM applicants_data
	INNER JOIN application_status ON applicants_data.applicant_id = application_status.applicant_id
	INNER JOIN user_accounts ON application_status.user_interviewee_id = user_accounts.account_id
	WHERE user_accounts.account_id = account_id;
END$$

CREATE DEFINER=`root`@`localhost` PROCEDURE `update_account_status` (IN `up_account_id` INT(3), IN `up_account_status` VARCHAR(255))   BEGIN 
UPDATE user_accounts
	SET account_status = up_account_status
	WHERE account_id = up_account_id;
END$$

CREATE DEFINER=`root`@`localhost` PROCEDURE `update_application_status` (IN `app_id` INT(3), IN `app_status` INT(3), IN `updated_at` VARCHAR(255))   BEGIN
UPDATE application_status
 SET application_status_id = app_status, updatedAt = updated_at
WHERE status_id = app_id;
END$$

CREATE DEFINER=`root`@`localhost` PROCEDURE `update_job_position` (IN `up_position_id` INT(3), IN `up_position_name` VARCHAR(255), IN `up_department_id` INT(3), IN `up_position_status` VARCHAR(255), IN `up_available_slot` INT(3))   BEGIN
UPDATE job_position
SET position_name = up_position_name, department_id = up_department_id, position_status = up_position_status, available_slot = up_available_slot
WHERE position_id = up_position_id;
END$$

CREATE DEFINER=`root`@`localhost` PROCEDURE `user_login` (IN `username` VARCHAR(255))   BEGIN
SELECT account_id, username, password FROM user_accounts WHERE username = username;
END$$

DELIMITER ;

-- --------------------------------------------------------

--
-- Table structure for table `applicants_data`
--

CREATE TABLE `applicants_data` (
  `applicant_id` int(3) NOT NULL,
  `position_id` int(3) NOT NULL,
  `first_name` varchar(255) NOT NULL,
  `middle_name` varchar(255) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `extension_name` varchar(255) NOT NULL,
  `birthdate` date NOT NULL,
  `age` int(3) NOT NULL,
  `present_address` varchar(255) NOT NULL,
  `highest_education` varchar(255) NOT NULL,
  `email_address` varchar(255) NOT NULL,
  `facebook_link` varchar(255) NOT NULL,
  `bpo_exp` varchar(255) NOT NULL,
  `shift_sched` varchar(255) NOT NULL,
  `work_report` varchar(255) NOT NULL,
  `work_site_location` varchar(255) NOT NULL,
  `platform_id` int(3) NOT NULL,
  `ref_full_name` varchar(255) NOT NULL,
  `ref_company` varchar(255) NOT NULL,
  `ref_position` varchar(255) NOT NULL,
  `ref_contact_num` varchar(255) NOT NULL,
  `ref_email` varchar(255) NOT NULL,
  `applicant_cv` varchar(255) NOT NULL,
  `applicant_portfolio_link` varchar(255) NOT NULL,
  `createdAt` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `applicants_data`
--

INSERT INTO `applicants_data` (`applicant_id`, `position_id`, `first_name`, `middle_name`, `last_name`, `extension_name`, `birthdate`, `age`, `present_address`, `highest_education`, `email_address`, `facebook_link`, `bpo_exp`, `shift_sched`, `work_report`, `work_site_location`, `platform_id`, `ref_full_name`, `ref_company`, `ref_position`, `ref_contact_num`, `ref_email`, `applicant_cv`, `applicant_portfolio_link`, `createdAt`) VALUES
(2, 2, 'Samantha', 'Purificacion', 'Eduardo', '', '2002-10-23', 22, 'Talavera', 'College', 'samantha@gmail.com', 'SamanthaEduardo', 'No', 'Yes', 'Yes', 'Yes', 3, '', '', '', '', '', '', '', '2024-04-29 11:15:06'),
(3, 0, 'Samantha', 'Purificacion', 'Eduardo', '', '2002-10-23', 22, 'Talavera', 'College', 'samantha@gmail.com', 'SamanthaEduardo', 'No', 'Yes', 'Yes', 'Yes', 1, '', '', '', '', '', '', '', '2024-04-29 11:15:09'),
(4, 0, 'Samantha', 'Purificacion', 'Eduardo', '', '2002-10-23', 22, 'Talavera', 'College', 'samantha@gmail.com', 'SamanthaEduardo', 'No', 'Yes', 'Yes', 'Yes', 1, '', '', '', '', '', '', '', '2024-04-29 11:15:10'),
(5, 0, 'Samantha', 'Purificacion', 'Eduardo', '', '2002-10-23', 22, 'Talavera', 'College', 'samantha@gmail.com', 'SamanthaEduardo', 'No', 'Yes', 'Yes', 'Yes', 2, '', '', '', '', '', '', '', '2024-04-29 11:15:13'),
(6, 0, 'Samantha', 'Purificacion', 'Eduardo', '', '2002-10-23', 22, 'Talavera', 'College', 'samantha@gmail.com', 'SamanthaEduardo', 'No', 'Yes', 'Yes', 'Yes', 2, '', '', '', '', '', '', '', '2024-04-29 14:07:54'),
(7, 0, 'Samantha', 'Purificacion', 'Eduardo', '', '2002-10-23', 22, 'Talavera', 'College', 'samantha@gmail.com', 'SamanthaEduardo', 'No', 'Yes', 'Yes', 'Yes', 2, '', '', '', '', '', '', '', '2024-04-29 14:08:04');

-- --------------------------------------------------------

--
-- Table structure for table `application_status`
--

CREATE TABLE `application_status` (
  `status_id` int(3) NOT NULL,
  `applicant_id` int(3) NOT NULL,
  `position_id` int(3) NOT NULL,
  `user_interviewee_id` int(3) NOT NULL,
  `interview_date` date NOT NULL,
  `interview_time` varchar(255) NOT NULL,
  `application_status_id` int(3) NOT NULL,
  `createdAt` varchar(255) DEFAULT NULL,
  `updatedAt` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `application_status`
--

INSERT INTO `application_status` (`status_id`, `applicant_id`, `position_id`, `user_interviewee_id`, `interview_date`, `interview_time`, `application_status_id`, `createdAt`, `updatedAt`) VALUES
(1, 2, 2, 2, '2024-04-09', '', 3, NULL, '2024-04-29 15:12:08');

-- --------------------------------------------------------

--
-- Table structure for table `application_status_list`
--

CREATE TABLE `application_status_list` (
  `application_status_id` int(3) NOT NULL,
  `application_status_name` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `application_status_list`
--

INSERT INTO `application_status_list` (`application_status_id`, `application_status_name`) VALUES
(1, 'Received'),
(2, 'Incomplete Application'),
(3, 'In Progress'),
(4, 'Interviewing'),
(5, 'Position Closed'),
(6, 'Shortlisted'),
(7, 'Progress Completed');

-- --------------------------------------------------------

--
-- Table structure for table `department`
--

CREATE TABLE `department` (
  `department_id` int(3) NOT NULL,
  `department_name` varchar(255) NOT NULL,
  `department_status` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `department`
--

INSERT INTO `department` (`department_id`, `department_name`, `department_status`) VALUES
(1, 'Human Resources and Administration', 'Active'),
(2, 'Project Management', 'Active'),
(3, 'Software Development', 'Active'),
(4, 'Creative Services', 'Active'),
(5, 'Global Service Desk', 'Active'),
(6, 'Network Operations Center', 'Active');

-- --------------------------------------------------------

--
-- Table structure for table `job_position`
--

CREATE TABLE `job_position` (
  `position_id` int(3) NOT NULL,
  `position_name` varchar(255) NOT NULL,
  `department_id` int(3) NOT NULL,
  `position_status` varchar(255) NOT NULL,
  `available_slot` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `job_position`
--

INSERT INTO `job_position` (`position_id`, `position_name`, `department_id`, `position_status`, `available_slot`) VALUES
(1, 'Customer Service Representative', 5, '', NULL),
(2, 'Software Developer', 2, 'Inactive', '4'),
(3, 'Project Management Associate', 2, '', NULL),
(4, 'QC Analyst', 3, '', NULL),
(5, 'Support Specialist T1', 5, '', NULL),
(6, 'Support Specialist T2', 5, '', NULL),
(7, 'Support Specialist T3', 5, '', NULL),
(8, 'Web Developer', 3, '', NULL),
(9, 'Linux Systems Administrator', 3, '', NULL),
(10, '', 0, '', '0'),
(11, '', 0, '', '0'),
(12, '', 0, '', '0'),
(13, '', 0, '', '0'),
(14, '', 0, '', '0'),
(15, '', 0, '', '0'),
(16, '', 0, '', '0'),
(17, 'Software Developer', 2, 'Available', '3');

-- --------------------------------------------------------

--
-- Table structure for table `job_posting_platform`
--

CREATE TABLE `job_posting_platform` (
  `platform_id` int(3) NOT NULL,
  `platform_name` varchar(255) NOT NULL,
  `platform_status` varchar(255) NOT NULL,
  `platform_others` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `job_posting_platform`
--

INSERT INTO `job_posting_platform` (`platform_id`, `platform_name`, `platform_status`, `platform_others`) VALUES
(1, 'MotivIT Website', 'Active', ''),
(2, 'Facebook', 'Active', ''),
(3, 'LinkedIn', 'Active', ''),
(4, 'Referral', 'Active', ''),
(5, 'Others', 'Active', '');

-- --------------------------------------------------------

--
-- Table structure for table `user_accounts`
--

CREATE TABLE `user_accounts` (
  `account_id` int(3) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `user_role` varchar(255) NOT NULL,
  `user_name` varchar(255) NOT NULL,
  `user_position` varchar(255) NOT NULL,
  `department_id` int(3) NOT NULL,
  `account_status` varchar(255) NOT NULL,
  `createdAt` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `user_accounts`
--

INSERT INTO `user_accounts` (`account_id`, `username`, `password`, `user_role`, `user_name`, `user_position`, `department_id`, `account_status`, `createdAt`) VALUES
(1, 'MichaelJR022', '$2a$10$fkuZDWadQzw8WK/1eWJ22O0RUA3tK2FPhhuo0iKOBIWE3OXi3K3.y', 'Manager', 'Michael Eduardo Jr', '0', 1, 'Inactive', '2024-04-29 10:39:49'),
(2, 'Samantha022', '$2a$10$A.3si9wSRpVn6Sb3rBzjy.HqpeAug57lo9Y6Ky1IABU3iCwjGLCBi', 'Manager', 'Michael Eduardo Jr', '0', 1, 'ACTIVE', '2024-04-29 10:40:12');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `applicants_data`
--
ALTER TABLE `applicants_data`
  ADD PRIMARY KEY (`applicant_id`);

--
-- Indexes for table `application_status`
--
ALTER TABLE `application_status`
  ADD PRIMARY KEY (`status_id`);

--
-- Indexes for table `application_status_list`
--
ALTER TABLE `application_status_list`
  ADD PRIMARY KEY (`application_status_id`);

--
-- Indexes for table `department`
--
ALTER TABLE `department`
  ADD PRIMARY KEY (`department_id`);

--
-- Indexes for table `job_position`
--
ALTER TABLE `job_position`
  ADD PRIMARY KEY (`position_id`);

--
-- Indexes for table `job_posting_platform`
--
ALTER TABLE `job_posting_platform`
  ADD PRIMARY KEY (`platform_id`);

--
-- Indexes for table `user_accounts`
--
ALTER TABLE `user_accounts`
  ADD PRIMARY KEY (`account_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `applicants_data`
--
ALTER TABLE `applicants_data`
  MODIFY `applicant_id` int(3) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `application_status`
--
ALTER TABLE `application_status`
  MODIFY `status_id` int(3) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `application_status_list`
--
ALTER TABLE `application_status_list`
  MODIFY `application_status_id` int(3) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `department`
--
ALTER TABLE `department`
  MODIFY `department_id` int(3) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `job_position`
--
ALTER TABLE `job_position`
  MODIFY `position_id` int(3) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;

--
-- AUTO_INCREMENT for table `job_posting_platform`
--
ALTER TABLE `job_posting_platform`
  MODIFY `platform_id` int(3) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `user_accounts`
--
ALTER TABLE `user_accounts`
  MODIFY `account_id` int(3) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
