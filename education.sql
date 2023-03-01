-- phpMyAdmin SQL Dump
-- version 4.2.11
-- http://www.phpmyadmin.net
--
-- Host: 127.0.0.1
-- Generation Time: Mar 01, 2023 at 12:00 PM
-- Server version: 5.6.21
-- PHP Version: 5.6.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;

--
-- Database: `education`
--

-- --------------------------------------------------------

--
-- Table structure for table `admin`
--

CREATE TABLE IF NOT EXISTS `admin` (
`id` int(11) NOT NULL,
  `name` varchar(300) NOT NULL,
  `email` varchar(100) NOT NULL,
  `username` varchar(300) NOT NULL,
  `password` varchar(100) NOT NULL
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

--
-- Dumping data for table `admin`
--

INSERT INTO `admin` (`id`, `name`, `email`, `username`, `password`) VALUES
(1, 'admin', 'admin@yopmail.com', 'gtAdmin', '$2a$10$R9v2nAWDpLq4rA2agV2.GOKsl00odpDmpI67tB2/ieKdwPQ9Bvbxm'),
(2, 'sami', 'sami@yopmail.com', 'sami', '$2a$10$fbbXJbLewvepFfAknC5mB.u1vidjtS2Y3qWFQlN1pp06.qKRv/C.e');

-- --------------------------------------------------------

--
-- Table structure for table `assignments`
--

CREATE TABLE IF NOT EXISTS `assignments` (
`id` int(11) NOT NULL,
  `title` varchar(100) NOT NULL,
  `assignment` text NOT NULL,
  `duedate` date NOT NULL,
  `totalmarks` varchar(10) NOT NULL,
  `submit` varchar(100) NOT NULL,
  `result` varchar(11) NOT NULL
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=latin1;

--
-- Dumping data for table `assignments`
--

INSERT INTO `assignments` (`id`, `title`, `assignment`, `duedate`, `totalmarks`, `submit`, `result`) VALUES
(17, 'Assignment no 1', 'Create Responsive Website', '2023-02-28', '20', 'Create Responsive Website', 'Fail');

-- --------------------------------------------------------

--
-- Table structure for table `course`
--

CREATE TABLE IF NOT EXISTS `course` (
`id` int(11) NOT NULL,
  `coursename` varchar(100) NOT NULL,
  `lesson` varchar(100) NOT NULL,
  `week` varchar(100) NOT NULL,
  `price` varchar(100) NOT NULL,
  `image` varchar(300) NOT NULL,
  `description` varchar(300) NOT NULL
) ENGINE=InnoDB AUTO_INCREMENT=55 DEFAULT CHARSET=latin1;

--
-- Dumping data for table `course`
--

INSERT INTO `course` (`id`, `coursename`, `lesson`, `week`, `price`, `image`, `description`) VALUES
(51, 'PHP Course', '4 Lessons', '10 Week', '12', '1677481016296465900.jpg', 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Porro eius suscipit delectus enim iusto tempora, adipisci at provident.'),
(52, 'C# Course', '12 Lesson', '10 Week', '12', '1677481035408182000.jpg', 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Porro eius suscipit delectus enim iusto tempora, adipisci at provident.'),
(53, 'CSS', '12 Lesson', '10 Week', '12', '1677481073472457900.jpg', 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Porro eius suscipit delectus enim iusto tempora, adipisci at provident.'),
(54, 'Javascript course', '12 Lesson', '10 Week', '12', '1677666263828743100.jpg', 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Porro eius suscipit delectus enim iusto tempora, adipisci at provident. ');

-- --------------------------------------------------------

--
-- Table structure for table `students`
--

CREATE TABLE IF NOT EXISTS `students` (
`id` int(11) NOT NULL,
  `name` varchar(300) NOT NULL,
  `email` varchar(100) NOT NULL,
  `address` text NOT NULL,
  `phone` varchar(50) NOT NULL,
  `password` varchar(300) NOT NULL
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=latin1;

--
-- Dumping data for table `students`
--

INSERT INTO `students` (`id`, `name`, `email`, `address`, `phone`, `password`) VALUES
(13, 'husnainraza', 'admin@yopmail.com', 'Nazimabad #street no 1', '35345345', '$2a$10$fa649aFyzqoSrKKtnL758erAr9/SVQnVQEBJAIw5RC82HxpX18vHm'),
(14, 'ishraq', 'webadmin@genetechsolutions.com', 'Nazimabad #street no 1', '5646464', '$2a$10$JYXlE02DekoX4QycBo6xr.rIarG3DaaY9POzZ.8K6qm928yRVOrwa'),
(15, 'Mazdoor Online', 'admin@yopmail.com', 'Nazimabad #street no 1', '515641564', '$2a$10$pX5q6yRVbRrjrD0TZxz5LuA3ncZSJN8JJPsXjx1iRFdM5YZFSZzMO'),
(16, 'Abdulsami', 'admin@yopmail.com', 'Nazimabad #street no 1', '65464', '$2a$10$gbJonWu.uEQkiybYtM3ndejAXB9MJ.JXwNXVjY6SAan3eQntSfnV2'),
(17, 'husnainraza', 'admin@yopmail.com', 'Nazimabad #street no 1', '5456456456546', '$2a$10$GqFViisH6xwo4h4Ny5/mQO6L0HGmiBTzl9uny0Gw0PKfIuieZTf0y'),
(18, 'Abdulsami', 'admin@yopmail.com', 'Nazimabad #street no 1', '564654564', '$2a$10$6c0.aHpS6EPqkY/BMqIt2.SfBwtprPZiOUQC40ca2WOBT7cl7zDvi'),
(19, 'Abdulsami', 'webadmin@genetechsolutions.com', 'Nazimabad #street no 1', '56464', '$2a$10$Hjp3P4cBg8hdAgD6IzuZ1.oGhTJ9kGQr9edMiEJw//3PTLuGmtBEO'),
(20, 'Abdulsami', 'admin@yopmail.com', 'Nazimabad #street no 1', '5646465', '$2a$10$SvaflKDC4YolxbbicjrBROhOLqjPz121N9rNdCowvlzPV50iNyqPK'),
(21, 'husnainraza', 'admin@yopmail.com', 'Nazimabad #street no 1', '564564', '$2a$10$maJzbCI1BxePiUf52FV14OjBZEAzh1xOth0p2TPnOaemBsikARe6e'),
(22, 'Abdulsami', 'admin@yopmail.com', 'Nazimabad #street no 1', '132556165', '$2a$10$z7.OhI2bXvtQOc1j4G4bKOUkkVQHdm2PMTVM299fv2G/3ZqSkPIxe');

-- --------------------------------------------------------

--
-- Table structure for table `teachers`
--

CREATE TABLE IF NOT EXISTS `teachers` (
`id` int(11) NOT NULL,
  `teachername` varchar(100) NOT NULL,
  `position` varchar(100) NOT NULL,
  `description` varchar(300) NOT NULL,
  `image` varchar(300) NOT NULL
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=latin1;

--
-- Dumping data for table `teachers`
--

INSERT INTO `teachers` (`id`, `teachername`, `position`, `description`, `image`) VALUES
(9, 'Abdul Sami', 'PHP Teacher', 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Porro eius suscipit delectus enim iusto tempora, adipisci at provident.', '1677480893385477600.jpg'),
(10, 'Subhan Nasir', 'Physics Teacher', 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Porro eius suscipit delectus enim iusto tempora, adipisci at provident.', '1677480918625382900.jpg'),
(11, 'Haider Abbas', 'PHP Teacher', 'Lorem ipsum dolor sit amet, consectetur adipisicing elit. Porro eius suscipit delectus enim iusto tempora, adipisci at provident.', '1677666325421919600.jpg');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `admin`
--
ALTER TABLE `admin`
 ADD PRIMARY KEY (`id`);

--
-- Indexes for table `assignments`
--
ALTER TABLE `assignments`
 ADD PRIMARY KEY (`id`);

--
-- Indexes for table `course`
--
ALTER TABLE `course`
 ADD PRIMARY KEY (`id`);

--
-- Indexes for table `students`
--
ALTER TABLE `students`
 ADD PRIMARY KEY (`id`);

--
-- Indexes for table `teachers`
--
ALTER TABLE `teachers`
 ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `admin`
--
ALTER TABLE `admin`
MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=3;
--
-- AUTO_INCREMENT for table `assignments`
--
ALTER TABLE `assignments`
MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=19;
--
-- AUTO_INCREMENT for table `course`
--
ALTER TABLE `course`
MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=55;
--
-- AUTO_INCREMENT for table `students`
--
ALTER TABLE `students`
MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=23;
--
-- AUTO_INCREMENT for table `teachers`
--
ALTER TABLE `teachers`
MODIFY `id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=13;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
