-- phpMyAdmin SQL Dump
-- version 4.8.3
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: May 13, 2020 at 12:37 PM
-- Server version: 5.6.47-cll-lve
-- PHP Version: 7.2.7

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `gomysql`
--

-- --------------------------------------------------------

--
-- Table structure for table `add_admin_emails`
--

CREATE TABLE `add_admin_emails` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `uniqueid` varchar(255) DEFAULT NULL,
  `flag` varchar(255) DEFAULT NULL,
  `expirydate` datetime DEFAULT NULL,
  `role` varchar(255) DEFAULT NULL,
  `clinicmaster_id` int(10) UNSIGNED DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

--
-- Dumping data for table `add_admin_emails`
--

INSERT INTO `add_admin_emails` (`id`, `created_at`, `updated_at`, `deleted_at`, `uniqueid`, `flag`, `expirydate`, `role`, `clinicmaster_id`) VALUES
(1, '2020-05-10 11:57:22', '2020-05-10 11:59:36', NULL, '982b26433dcc9aa327e0293cdd8c576b2aaa7016', 'B', '2020-05-10 11:57:09', 'rootadmin', 1),
(2, '2020-05-10 12:01:05', '2020-05-10 12:03:00', NULL, 'cdc0bef7f5b16c80ce41a28a8e2c1b17ee9d5e00', 'B', '2020-05-10 12:00:53', 'admin', 2);

-- --------------------------------------------------------

--
-- Table structure for table `clients`
--

CREATE TABLE `clients` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

--
-- Dumping data for table `clients`
--

INSERT INTO `clients` (`id`, `created_at`, `updated_at`, `deleted_at`, `username`, `password`) VALUES
(1, '2020-04-26 16:27:32', '2020-04-26 16:27:32', NULL, 'sakib@gmail.com', '12345678');

-- --------------------------------------------------------

--
-- Table structure for table `clinicmasters`
--

CREATE TABLE `clinicmasters` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `clinic_name` varchar(255) DEFAULT NULL,
  `cityname` varchar(255) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `pincode` int(11) DEFAULT NULL,
  `district` varchar(255) DEFAULT NULL,
  `country` varchar(255) DEFAULT NULL,
  `emailid` varchar(255) DEFAULT NULL,
  `role` varchar(255) DEFAULT NULL,
  `concernperson` varchar(255) DEFAULT NULL,
  `internal_flag` varchar(255) DEFAULT NULL,
  `created_date` datetime DEFAULT NULL,
  `created_by` bigint(20) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

--
-- Dumping data for table `clinicmasters`
--

INSERT INTO `clinicmasters` (`id`, `created_at`, `updated_at`, `deleted_at`, `clinic_name`, `cityname`, `address`, `pincode`, `district`, `country`, `emailid`, `role`, `concernperson`, `internal_flag`, `created_date`, `created_by`) VALUES
(1, '2020-05-10 11:57:08', '2020-05-10 11:57:08', NULL, 'Alhamdulillha', 'NAVI MUMBAI', 'Trimurti Society, A type, A6B 1:2  Sector 8,\nKhanda colony, NEW PANVEL', 410206, 'raigad', 'India', 'sakibcoolz@gmail.com', 'rootadmin', 'sakib mulla', 'Team', '0000-00-00 00:00:00', 0),
(2, '2020-05-10 12:00:52', '2020-05-10 12:00:52', NULL, 'Sai Clinic', 'NAVI MUMBAI', 'Trimurti Society, A type, A6B 1:2  Sector 8,\nKhanda colony, NEW PANVEL', 410206, 'raigad', 'India', 'sakibcoolz@gmail.com', 'admin', 'Amit Malekar', 'clinic', '0000-00-00 00:00:00', 0);

-- --------------------------------------------------------

--
-- Table structure for table `loginmasters`
--

CREATE TABLE `loginmasters` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `firstname` varchar(255) DEFAULT NULL,
  `lastname` varchar(255) DEFAULT NULL,
  `dob` datetime DEFAULT NULL,
  `rights_no` int(10) UNSIGNED DEFAULT NULL,
  `homeaddress1` varchar(255) DEFAULT NULL,
  `homeaddress2` varchar(255) DEFAULT NULL,
  `homeaddress3` varchar(255) DEFAULT NULL,
  `pincode` varchar(255) DEFAULT NULL,
  `cityname` varchar(255) DEFAULT NULL,
  `university` varchar(255) DEFAULT NULL,
  `degreename` varchar(255) DEFAULT NULL,
  `passoutyear` varchar(255) DEFAULT NULL,
  `policyaccept` tinyint(1) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `repassword` varchar(255) DEFAULT NULL,
  `oldpassword` varchar(255) DEFAULT NULL,
  `resetflag` varchar(255) DEFAULT NULL,
  `emailid` varchar(255) NOT NULL DEFAULT '',
  `contactno` varchar(255) DEFAULT NULL,
  `reffcontno` varchar(255) DEFAULT NULL,
  `uniqueid` varchar(255) DEFAULT NULL,
  `statusflag` varchar(255) DEFAULT NULL,
  `clinicmaster_id` int(10) UNSIGNED DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

--
-- Dumping data for table `loginmasters`
--

INSERT INTO `loginmasters` (`id`, `created_at`, `updated_at`, `deleted_at`, `firstname`, `lastname`, `dob`, `rights_no`, `homeaddress1`, `homeaddress2`, `homeaddress3`, `pincode`, `cityname`, `university`, `degreename`, `passoutyear`, `policyaccept`, `password`, `repassword`, `oldpassword`, `resetflag`, `emailid`, `contactno`, `reffcontno`, `uniqueid`, `statusflag`, `clinicmaster_id`) VALUES
(1, '2020-05-10 11:59:42', '2020-05-10 11:59:42', NULL, 'sakib', 'mullla', '2020-05-09 18:30:00', 1, 'Trimurti Society, A type, A6B 1:2  Sector 8,', 'Khanda colony, NEW PANVEL', '410206', '410206', 'NAVI MUMBAI', 'MAHARASHTRA', 'mca', '2017', 1, '12345678', '12345678', '', 'S', 'sakibcoolz@gmail.com', '8433854833', '09082015318', '982b26433dcc9aa327e0293cdd8c576b2aaa7016', 'A', 1),
(2, '2020-05-10 12:03:06', '2020-05-10 12:03:06', NULL, 'amit', 'malekar', '2020-05-09 18:30:00', 2, 'KHANDA COLONY', 'NEW PANVEL', '410206', '410206', 'Mumbai', 'dgfdgfhfgf', 'mca', '2017', 1, '12345678', '12345678', '', 'S', 'sakibyups@gmail.com', '9082015318', '1234567890', 'cdc0bef7f5b16c80ce41a28a8e2c1b17ee9d5e00', 'A', 2);

-- --------------------------------------------------------

--
-- Table structure for table `rights`
--

CREATE TABLE `rights` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `servicename` varchar(255) DEFAULT NULL,
  `add` tinyint(1) DEFAULT NULL,
  `edit` tinyint(1) DEFAULT NULL,
  `view` tinyint(1) DEFAULT NULL,
  `delete` tinyint(1) DEFAULT NULL,
  `comments` varchar(255) DEFAULT NULL,
  `htmltag` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

--
-- Dumping data for table `rights`
--

INSERT INTO `rights` (`id`, `created_at`, `updated_at`, `deleted_at`, `servicename`, `add`, `edit`, `view`, `delete`, `comments`, `htmltag`) VALUES
(1, '0000-00-00 00:00:00', '0000-00-00 00:00:00', NULL, 'clinic_managements', 0, 0, 0, 0, 'Clinic Management', 'Clinic Management'),
(2, '0000-00-00 00:00:00', '0000-00-00 00:00:00', NULL, 'user_managements', 0, 1, 0, 0, 'User Management', 'User Management'),
(3, '0000-00-00 00:00:00', '0000-00-00 00:00:00', NULL, 'rights', 0, 0, 0, 0, 'Rights Management', 'Rights Management');

-- --------------------------------------------------------

--
-- Table structure for table `rightsmasters`
--

CREATE TABLE `rightsmasters` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `clinicmaster_id` int(10) UNSIGNED DEFAULT NULL,
  `rightsname` varchar(255) DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

--
-- Dumping data for table `rightsmasters`
--

INSERT INTO `rightsmasters` (`id`, `created_at`, `updated_at`, `deleted_at`, `clinicmaster_id`, `rightsname`) VALUES
(1, '2020-05-10 11:59:39', '2020-05-10 11:59:39', NULL, 1, 'rootadmin'),
(2, '2020-05-10 12:03:03', '2020-05-10 12:03:03', NULL, 2, 'admin'),
(6, '2020-05-13 19:35:02', '2020-05-13 19:35:02', NULL, 1, 'sakib');

-- --------------------------------------------------------

--
-- Table structure for table `rightsservicemappers`
--

CREATE TABLE `rightsservicemappers` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `servicename` varchar(255) DEFAULT NULL,
  `add` tinyint(1) DEFAULT NULL,
  `edit` tinyint(1) DEFAULT NULL,
  `view` tinyint(1) DEFAULT NULL,
  `delete` tinyint(1) DEFAULT NULL,
  `rightsmaster_id` int(10) UNSIGNED DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

--
-- Dumping data for table `rightsservicemappers`
--

INSERT INTO `rightsservicemappers` (`id`, `created_at`, `updated_at`, `deleted_at`, `servicename`, `add`, `edit`, `view`, `delete`, `rightsmaster_id`) VALUES
(1, '2020-05-10 11:59:39', '2020-05-10 11:59:39', NULL, 'clinic_managements', 1, 1, 1, 1, 1),
(2, '2020-05-10 11:59:40', '2020-05-10 11:59:40', NULL, 'user_management', 1, 1, 1, 1, 1),
(3, '2020-05-10 11:59:40', '2020-05-10 11:59:40', NULL, 'todayappointment', 1, 1, 1, 1, 1),
(4, '2020-05-10 11:59:40', '2020-05-10 11:59:40', NULL, 'newappointment', 1, 1, 1, 1, 1),
(5, '2020-05-10 11:59:41', '2020-05-10 11:59:41', NULL, 'opdappointmenthistory', 1, 1, 1, 1, 1),
(6, '2020-05-10 12:03:03', '2020-05-10 12:03:03', NULL, 'user_management', 1, 1, 1, 1, 2),
(7, '2020-05-10 12:03:04', '2020-05-10 12:03:04', NULL, 'todayappointment', 1, 1, 1, 1, 2),
(8, '2020-05-10 12:03:04', '2020-05-10 12:03:04', NULL, 'newappointment', 1, 1, 1, 1, 2),
(9, '2020-05-10 12:03:05', '2020-05-10 12:03:05', NULL, 'opdappointmenthistory', 1, 1, 1, 1, 2),
(10, '0000-00-00 00:00:00', '0000-00-00 00:00:00', NULL, 'rights', 1, 1, 1, 1, 1),
(13, '0000-00-00 00:00:00', '0000-00-00 00:00:00', NULL, 'rights', 1, 0, 1, 0, 2),
(25, '2020-05-13 19:35:02', '2020-05-13 19:35:02', NULL, 'rights', 1, 1, 1, 0, 6),
(24, '2020-05-13 19:35:02', '2020-05-13 19:35:02', NULL, 'user_managements', 1, 1, 1, 0, 6);

-- --------------------------------------------------------

--
-- Table structure for table `servicemasters`
--

CREATE TABLE `servicemasters` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `htmlname` varchar(255) DEFAULT NULL,
  `paths` varchar(255) DEFAULT NULL,
  `comments` varchar(200) NOT NULL
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

--
-- Dumping data for table `servicemasters`
--

INSERT INTO `servicemasters` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`, `htmlname`, `paths`, `comments`) VALUES
(1, '0000-00-00 00:00:00', '0000-00-00 00:00:00', NULL, 'clinic_managements', 'Add Clinic', '/ClinicManagement', '0'),
(2, '0000-00-00 00:00:00', '0000-00-00 00:00:00', NULL, 'user_management', 'Add user', '/UserManagement', '0'),
(3, '0000-00-00 00:00:00', '0000-00-00 00:00:00', NULL, 'todayappointment', 'Today Appointment', '/opdTodayappointment', '0'),
(4, '0000-00-00 00:00:00', '0000-00-00 00:00:00', NULL, 'newappointment', 'New Opd Appointment', '/opdnewappointment', '0'),
(5, '0000-00-00 00:00:00', '0000-00-00 00:00:00', NULL, 'opdappointmenthistory', 'Opd Appointment History', '/opdhisappointment', '0'),
(6, '0000-00-00 00:00:00', '0000-00-00 00:00:00', NULL, 'rights', 'Rights', '/rightsmanagement', '0');

-- --------------------------------------------------------

--
-- Table structure for table `telnumbers`
--

CREATE TABLE `telnumbers` (
  `id` int(10) UNSIGNED NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `telenumber` bigint(20) DEFAULT NULL,
  `clinicmaster_id` int(10) UNSIGNED DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

--
-- Dumping data for table `telnumbers`
--

INSERT INTO `telnumbers` (`id`, `created_at`, `updated_at`, `deleted_at`, `telenumber`, `clinicmaster_id`) VALUES
(1, '2020-05-10 11:57:08', '2020-05-10 11:57:08', NULL, 9082015318, 1),
(2, '2020-05-10 12:00:52', '2020-05-10 12:00:52', NULL, 4444488888, 2);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `add_admin_emails`
--
ALTER TABLE `add_admin_emails`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_add_admin_emails_deleted_at` (`deleted_at`);

--
-- Indexes for table `clients`
--
ALTER TABLE `clients`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_clients_deleted_at` (`deleted_at`);

--
-- Indexes for table `clinicmasters`
--
ALTER TABLE `clinicmasters`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_clinicmasters_deleted_at` (`deleted_at`);

--
-- Indexes for table `loginmasters`
--
ALTER TABLE `loginmasters`
  ADD PRIMARY KEY (`id`,`emailid`),
  ADD UNIQUE KEY `contactno` (`contactno`),
  ADD KEY `idx_loginmasters_deleted_at` (`deleted_at`);

--
-- Indexes for table `rights`
--
ALTER TABLE `rights`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_rights_deleted_at` (`deleted_at`);

--
-- Indexes for table `rightsmasters`
--
ALTER TABLE `rightsmasters`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_rightsmasters_deleted_at` (`deleted_at`);

--
-- Indexes for table `rightsservicemappers`
--
ALTER TABLE `rightsservicemappers`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_rightsservicemappers_deleted_at` (`deleted_at`),
  ADD KEY `idx_rightsservicemappers_rightsmaster_id` (`rightsmaster_id`);

--
-- Indexes for table `servicemasters`
--
ALTER TABLE `servicemasters`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_servicemasters_deleted_at` (`deleted_at`);

--
-- Indexes for table `telnumbers`
--
ALTER TABLE `telnumbers`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_telnumbers_deleted_at` (`deleted_at`),
  ADD KEY `idx_telnumbers_clinicmaster_id` (`clinicmaster_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `add_admin_emails`
--
ALTER TABLE `add_admin_emails`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `clients`
--
ALTER TABLE `clients`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `clinicmasters`
--
ALTER TABLE `clinicmasters`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `loginmasters`
--
ALTER TABLE `loginmasters`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `rights`
--
ALTER TABLE `rights`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `rightsmasters`
--
ALTER TABLE `rightsmasters`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `rightsservicemappers`
--
ALTER TABLE `rightsservicemappers`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=26;

--
-- AUTO_INCREMENT for table `servicemasters`
--
ALTER TABLE `servicemasters`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT for table `telnumbers`
--
ALTER TABLE `telnumbers`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
