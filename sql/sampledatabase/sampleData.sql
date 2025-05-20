-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: 15. Mai, 2025 17:50 PM
-- Tjener-versjon: 11.7.2-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `idatg2204`
--

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `administrators`
--

CREATE TABLE `administrators` (
  `UserID` varchar(50) NOT NULL,
  `RoleName` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

--
-- Dataark for tabell `administrators`
--

INSERT INTO `administrators` (`UserID`, `RoleName`) VALUES
('c01ab0a2-5036-4273-a1c3-7f19d4aa257c', 'admin');

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `brand`
--

CREATE TABLE `brand` (
  `BrandName` varchar(50) NOT NULL,
  `BrandDesc` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

--
-- Dataark for tabell `brand`
--

INSERT INTO `brand` (`BrandName`, `BrandDesc`) VALUES
('DigiTech', 'DigiTech powers your lifestyle with a full range of smart electronics — from smartphones to home appliances — all designed for seamless digital living.'),
('E-Gadgets', 'E-Gadget brings everyday innovation to your fingertips with a wide selection of electronics and accessories that make life simpler, smarter, and more exciting.'),
('TechBrave', 'TechBrave delivers bold, cutting-edge technology across smartphones, laptops, cameras, and more, empowering pioneers in every part of life.');

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `cartitem`
--

CREATE TABLE `cartitem` (
  `UserID` varchar(50) NOT NULL,
  `ProductID` varchar(50) NOT NULL,
  `Quantity` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

--
-- Dataark for tabell `cartitem`
--

INSERT INTO `cartitem` (`UserID`, `ProductID`, `Quantity`) VALUES
('34010147-490a-444a-ba96-9550b0126006', '628a074c-817f-4d69-ba30-6aa04b023b43', 2),
('34010147-490a-444a-ba96-9550b0126006', 'fc84edd2-93ba-46c6-a453-61e974d0d8aa', 2),
('4b8b39c2-e536-4e1b-985b-77e9989d3de4', '8ced2470-c515-403f-b4a1-69cb2c8434c1', 3),
('4b8b39c2-e536-4e1b-985b-77e9989d3de4', 'c522e1d3-1da6-4915-8431-e10d9bb35a2c', 2),
('c01ab0a2-5036-4273-a1c3-7f19d4aa257c', '0f27563b-9c7c-4fc9-8fe0-17c668f7d865', 2),
('c01ab0a2-5036-4273-a1c3-7f19d4aa257c', '4da637b6-f102-4d5e-b5cb-3bf209acd705', 1),
('c01ab0a2-5036-4273-a1c3-7f19d4aa257c', 'aa185448-3e8d-49a2-9184-f5f3a6559f25', 1),
('d2054caf-1155-47a6-8791-3222c9d6beb8', '306ec8e0-6cd0-4af2-84f2-a2c446e59ea7', 3),
('d2054caf-1155-47a6-8791-3222c9d6beb8', '41499fa1-f67a-4293-a6b3-bc469e045d6b', 3),
('d2054caf-1155-47a6-8791-3222c9d6beb8', '63e238fc-9a93-47e7-8ee0-df418a1f2aac', 2);

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `category`
--

CREATE TABLE `category` (
  `CategoryName` varchar(50) NOT NULL,
  `CategoryDesc` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

--
-- Dataark for tabell `category`
--

INSERT INTO `category` (`CategoryName`, `CategoryDesc`) VALUES
('Audio & Headphones', 'Speakers, headphones, earphones, and audio systems.'),
('Cameras & Photography', 'Digital cameras, lenses, tripods, and photography gear.'),
('Computers & Accessories', 'Desktops, laptops, components, and peripherals like keyboards and mice.'),
('Electronics', 'Gadgets, smartphones, computers, and accessories.'),
('Gaming', 'Consoles, video games, gaming accessories, and VR devices.'),
('Home Appliances', 'Refrigerators, washing machines, microwaves, and small kitchen devices.'),
('Mobile Phones & Tablets', 'Smartphones, tablets, and mobile accessories.'),
('Smart Home', 'Smart lights, thermostats, security cameras, and home automation.'),
('TV & Home Theater', 'Televisions, projectors, sound systems, and accessories.'),
('Wearable Technology', 'Smartwatches, fitness trackers, and health tech devices.');

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `members`
--

CREATE TABLE `members` (
  `UserID` varchar(50) NOT NULL,
  `MembershipLevel` enum('Silver','Gold','Platinum') NOT NULL,
  `MembershipStart` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

--
-- Dataark for tabell `members`
--

INSERT INTO `members` (`UserID`, `MembershipLevel`, `MembershipStart`) VALUES
('34010147-490a-444a-ba96-9550b0126006', 'Platinum', '2025-05-15'),
('c01ab0a2-5036-4273-a1c3-7f19d4aa257c', 'Gold', '2025-05-15'),
('d2054caf-1155-47a6-8791-3222c9d6beb8', 'Silver', '2025-05-15');

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `orderitem`
--

CREATE TABLE `orderitem` (
  `OrderID` varchar(50) NOT NULL,
  `ProductID` varchar(50) NOT NULL,
  `Quantity` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

--
-- Dataark for tabell `orderitem`
--

INSERT INTO `orderitem` (`OrderID`, `ProductID`, `Quantity`) VALUES
('1673dc26-8ab8-428d-bc33-e00f567f5514', 'aa185448-3e8d-49a2-9184-f5f3a6559f25', 2),
('1673dc26-8ab8-428d-bc33-e00f567f5514', 'f33c2b0a-17e7-4f05-b4ab-bc95dd5f8c0c', 2),
('16c9ec5c-68bf-415f-b7d5-6ffc5edce1ad', '30b72f14-f0e6-4eb5-94c1-80371627ab22', 1),
('16c9ec5c-68bf-415f-b7d5-6ffc5edce1ad', 'c522e1d3-1da6-4915-8431-e10d9bb35a2c', 3),
('2636954b-3803-4b04-a46b-2ae9a3a764bc', '0f27563b-9c7c-4fc9-8fe0-17c668f7d865', 1),
('2636954b-3803-4b04-a46b-2ae9a3a764bc', 'bc9a1acf-375c-445d-a657-51206a952938', 3),
('470b0d9f-8eb4-4f03-aeb4-4fdf4672c6a0', '8ced2470-c515-403f-b4a1-69cb2c8434c1', 2),
('470b0d9f-8eb4-4f03-aeb4-4fdf4672c6a0', 'c02ce86b-a251-4154-900e-9759619c02f5', 3),
('a6be4785-17ad-45ea-9cdf-c7d146050dec', '41499fa1-f67a-4293-a6b3-bc469e045d6b', 2),
('a6be4785-17ad-45ea-9cdf-c7d146050dec', '903b683a-a94b-4849-9c07-8fb33fa43fbd', 3),
('a70e0fe5-5371-4598-bda2-4e2fdff3858e', '53e355eb-97d6-4f04-aebc-d87ddd3d0ae1', 1),
('a70e0fe5-5371-4598-bda2-4e2fdff3858e', 'fc84edd2-93ba-46c6-a453-61e974d0d8aa', 1),
('a87df88e-737d-4ac2-8330-beab8a65a396', '4da637b6-f102-4d5e-b5cb-3bf209acd705', 3),
('a87df88e-737d-4ac2-8330-beab8a65a396', '944461b2-856f-45bd-a5a5-57dcd55ec145', 2),
('c11590f1-e88b-4b66-ab5c-44eaa24b489d', '306ec8e0-6cd0-4af2-84f2-a2c446e59ea7', 3),
('c11590f1-e88b-4b66-ab5c-44eaa24b489d', 'ec8fad3a-edbc-4623-9728-55247bf43997', 1),
('e66c4a2c-74b4-4d41-b503-9697ff11faed', '63e238fc-9a93-47e7-8ee0-df418a1f2aac', 1),
('e66c4a2c-74b4-4d41-b503-9697ff11faed', '94a4e493-03dd-4bba-a353-01ebf4a20809', 3),
('f19f1338-b16a-48f5-b66a-8618933bb58b', '628a074c-817f-4d69-ba30-6aa04b023b43', 3),
('f19f1338-b16a-48f5-b66a-8618933bb58b', 'bc5a29e4-7a41-4bd0-8498-10863982cad4', 1);

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `orderstatus`
--

CREATE TABLE `orderstatus` (
  `StatusName` varchar(50) NOT NULL,
  `StatusDesc` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

--
-- Dataark for tabell `orderstatus`
--

INSERT INTO `orderstatus` (`StatusName`, `StatusDesc`) VALUES
('Cancelled', 'Order was cancelled by the customer or store.'),
('Delivered', 'Order successfully delivered to the customer.'),
('Failed', 'Order payment failed or could not be processed.'),
('Pending', 'Order received but not yet processed.'),
('Processing', 'Order is currently being prepared.'),
('Refunded', 'Customer has been refunded for the order.'),
('Returned', 'Customer has returned the order.'),
('Shipped', 'Order has been shipped to the customer.');

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `ordertable`
--

CREATE TABLE `ordertable` (
  `OrderID` varchar(50) NOT NULL,
  `UserID` varchar(50) DEFAULT NULL,
  `OrderDate` date DEFAULT NULL,
  `OrderStatus` varchar(50) NOT NULL,
  `OrderTotal` decimal(10,2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

--
-- Dataark for tabell `ordertable`
--

INSERT INTO `ordertable` (`OrderID`, `UserID`, `OrderDate`, `OrderStatus`, `OrderTotal`) VALUES
('1673dc26-8ab8-428d-bc33-e00f567f5514', 'c01ab0a2-5036-4273-a1c3-7f19d4aa257c', '2025-03-04', 'Pending', 19999.96),
('16c9ec5c-68bf-415f-b7d5-6ffc5edce1ad', '4b8b39c2-e536-4e1b-985b-77e9989d3de4', '2025-02-10', 'Delivered', 13999.96),
('2636954b-3803-4b04-a46b-2ae9a3a764bc', 'c01ab0a2-5036-4273-a1c3-7f19d4aa257c', '2025-02-15', 'Cancelled', 5999.96),
('470b0d9f-8eb4-4f03-aeb4-4fdf4672c6a0', '4b8b39c2-e536-4e1b-985b-77e9989d3de4', '2025-05-01', 'Failed', 16999.95),
('a6be4785-17ad-45ea-9cdf-c7d146050dec', 'd2054caf-1155-47a6-8791-3222c9d6beb8', '2025-03-27', 'Returned', 3899.95),
('a70e0fe5-5371-4598-bda2-4e2fdff3858e', '34010147-490a-444a-ba96-9550b0126006', '2025-04-23', 'Shipped', 10999.98),
('a87df88e-737d-4ac2-8330-beab8a65a396', 'c01ab0a2-5036-4273-a1c3-7f19d4aa257c', '2025-03-04', 'Pending', 2799.95),
('c11590f1-e88b-4b66-ab5c-44eaa24b489d', 'd2054caf-1155-47a6-8791-3222c9d6beb8', '2025-03-18', 'Processing', 39799.96),
('e66c4a2c-74b4-4d41-b503-9697ff11faed', 'd2054caf-1155-47a6-8791-3222c9d6beb8', '2025-02-21', 'Processing', 24999.96),
('f19f1338-b16a-48f5-b66a-8618933bb58b', '34010147-490a-444a-ba96-9550b0126006', '2025-03-17', 'Refunded', 12299.96);

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `payment`
--

CREATE TABLE `payment` (
  `PaymentID` varchar(50) NOT NULL,
  `OrderID` varchar(50) NOT NULL,
  `PaymentMethod` enum('card','vipps','bank_transfer') NOT NULL,
  `Amount` decimal(10,2) NOT NULL,
  `PaymentDate` date NOT NULL,
  `PaymentStatus` enum('pending','successful','failed') DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

--
-- Dataark for tabell `payment`
--

INSERT INTO `payment` (`PaymentID`, `OrderID`, `PaymentMethod`, `Amount`, `PaymentDate`, `PaymentStatus`) VALUES
('217d7cde-1a7a-449d-8222-84e9f859f5a0', 'a6be4785-17ad-45ea-9cdf-c7d146050dec', 'bank_transfer', 3899.95, '2025-03-31', 'failed'),
('36fed162-f630-4b4c-a5be-f296a3735ded', '2636954b-3803-4b04-a46b-2ae9a3a764bc', 'vipps', 5999.96, '2025-03-25', 'failed'),
('835f8d71-6749-4f0b-8956-27a2f6e4c546', '470b0d9f-8eb4-4f03-aeb4-4fdf4672c6a0', 'vipps', 16999.95, '2025-04-09', 'failed'),
('b650c4b7-0d3c-4f6f-8bd7-3acee4a2cb24', 'e66c4a2c-74b4-4d41-b503-9697ff11faed', 'card', 24999.96, '2025-02-06', 'pending'),
('c3fdbc5a-b82e-491f-bcb8-1ef01bfdfd58', 'f19f1338-b16a-48f5-b66a-8618933bb58b', 'card', 12299.96, '2025-04-12', 'failed'),
('c7744ce5-6ddb-4e45-a939-be87dba57c64', 'a87df88e-737d-4ac2-8330-beab8a65a396', 'bank_transfer', 2799.95, '2025-02-23', 'pending'),
('c7c1540e-9847-4f60-9cfa-274470adae64', 'c11590f1-e88b-4b66-ab5c-44eaa24b489d', 'vipps', 39799.96, '2025-03-02', 'pending'),
('d0c93737-1f6f-448f-b954-b8088df72fd9', 'a70e0fe5-5371-4598-bda2-4e2fdff3858e', 'bank_transfer', 10999.98, '2025-04-06', 'successful'),
('fb9f4fef-6a2a-4666-be58-a62a1889f7d2', '1673dc26-8ab8-428d-bc33-e00f567f5514', 'card', 19999.96, '2025-04-18', 'pending'),
('ff369a08-4f14-496d-a79a-871abeb6759d', '16c9ec5c-68bf-415f-b7d5-6ffc5edce1ad', 'card', 13999.96, '2025-04-07', 'successful');

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `product`
--

CREATE TABLE `product` (
  `ProductID` varchar(50) NOT NULL,
  `ProductName` varchar(100) NOT NULL,
  `ProductDesc` varchar(255) DEFAULT NULL,
  `ProductImgUrl` varchar(255) DEFAULT NULL,
  `Price` decimal(10,2) NOT NULL,
  `StockQuantity` int(11) NOT NULL,
  `Brand` varchar(50) DEFAULT NULL,
  `Category` varchar(50) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

--
-- Dataark for tabell `product`
--

INSERT INTO `product` (`ProductID`, `ProductName`, `ProductDesc`, `ProductImgUrl`, `Price`, `StockQuantity`, `Brand`, `Category`) VALUES
('0f27563b-9c7c-4fc9-8fe0-17c668f7d865', 'E-Gadgets Home Assistant', 'Voice-controlled smart home hub.', 'https://example.com/img5.jpg', 1499.99, 50, 'E-Gadgets', 'Smart Home'),
('306ec8e0-6cd0-4af2-84f2-a2c446e59ea7', 'E-Gadgets Ultra HD TV 55\"', 'Ultra HD smart TV with vibrant colors.', 'https://example.com/img2.jpg', 12999.99, 30, 'E-Gadgets', 'TV & Home Theater'),
('30b72f14-f0e6-4eb5-94c1-80371627ab22', 'E-Gadgets 4K Projector', 'Cinema-quality projector for home.', 'https://example.com/img14.jpg', 4999.99, 45, 'E-Gadgets', 'TV & Home Theater'),
('41499fa1-f67a-4293-a6b3-bc469e045d6b', 'DigiTech Bluetooth Speaker', 'Portable speaker with rich bass.', 'https://example.com/img6.jpg', 899.99, 200, 'DigiTech', 'Audio & Headphones'),
('4da637b6-f102-4d5e-b5cb-3bf209acd705', 'DigiTech Wireless Charger', 'Fast wireless charging pad.', 'https://example.com/img9.jpg', 599.99, 300, 'DigiTech', 'Mobile Phones & Tablets'),
('53e355eb-97d6-4f04-aebc-d87ddd3d0ae1', 'TechBrave Drone Pro', 'Professional drone with 4K camera.', 'https://example.com/img13.jpg', 8999.99, 25, 'TechBrave', 'Cameras & Photography'),
('628a074c-817f-4d69-ba30-6aa04b023b43', 'TechBrave Tablet 10\"', 'Lightweight tablet for work and play.', 'https://example.com/img7.jpg', 3999.99, 90, 'TechBrave', 'Mobile Phones & Tablets'),
('63e238fc-9a93-47e7-8ee0-df418a1f2aac', 'TechBrave Laptop Pro', 'Powerful laptop for professionals.', 'https://example.com/img10.jpg', 15999.99, 40, 'TechBrave', 'Computers & Accessories'),
('8ced2470-c515-403f-b4a1-69cb2c8434c1', 'E-Gadgets Noise Cancelling Headphones', 'Noise-free music experience.', 'https://example.com/img8.jpg', 2499.99, 70, 'E-Gadgets', 'Audio & Headphones'),
('903b683a-a94b-4849-9c07-8fb33fa43fbd', 'TechBrave Fitness Band', 'Basic fitness tracker with heart monitor.', 'https://example.com/img16.jpg', 699.99, 110, 'TechBrave', 'Wearable Technology'),
('944461b2-856f-45bd-a5a5-57dcd55ec145', 'TechBrave Power Bank 20000mAh', 'Portable charger for mobile devices.', 'https://example.com/img19.jpg', 499.99, 250, 'TechBrave', 'Mobile Phones & Tablets'),
('94a4e493-03dd-4bba-a353-01ebf4a20809', 'E-Gadgets Robot Vacuum', 'Smart robot vacuum cleaner.', 'https://example.com/img20.jpg', 2999.99, 30, 'E-Gadgets', 'Smart Home'),
('aa185448-3e8d-49a2-9184-f5f3a6559f25', 'TechBrave X100 Smartphone', 'High-end smartphone with AI features.', 'https://example.com/img1.jpg', 8999.99, 100, NULL, 'Mobile Phones & Tablets'),
('bc5a29e4-7a41-4bd0-8498-10863982cad4', 'E-Gadgets Smart Light Bulb', 'WiFi-enabled color-changing bulb.', 'https://example.com/img17.jpg', 299.99, 500, 'E-Gadgets', 'Smart Home'),
('bc9a1acf-375c-445d-a657-51206a952938', 'DigiTech Portable SSD 1TB', 'High-speed portable storage.', 'https://example.com/img15.jpg', 1499.99, 70, 'DigiTech', 'Computers & Accessories'),
('c02ce86b-a251-4154-900e-9759619c02f5', 'DigiTech VR Headset', 'Immersive VR gaming experience.', 'https://example.com/img18.jpg', 3999.99, 60, 'DigiTech', 'Gaming'),
('c522e1d3-1da6-4915-8431-e10d9bb35a2c', 'TechBrave Smartwatch 2', 'Fitness tracking smartwatch.', 'https://example.com/img4.jpg', 2999.99, 150, 'TechBrave', 'Wearable Technology'),
('ec8fad3a-edbc-4623-9728-55247bf43997', 'DigiTech Gaming Mouse', 'Precision gaming mouse.', 'https://example.com/img12.jpg', 799.99, 120, 'DigiTech', 'Gaming'),
('f33c2b0a-17e7-4f05-b4ab-bc95dd5f8c0c', 'E-Gadgets Security Camera', 'Smart WiFi home security camera.', 'https://example.com/img11.jpg', 999.99, 100, 'E-Gadgets', 'Smart Home'),
('fc84edd2-93ba-46c6-a453-61e974d0d8aa', 'DigiTech Gaming Headset Pro', 'Immersive gaming headset with surround sound.', 'https://example.com/img3.jpg', 1999.99, 80, 'DigiTech', 'Audio & Headphones');

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `review`
--

CREATE TABLE `review` (
  `UserID` varchar(50) NOT NULL,
  `ProductID` varchar(50) NOT NULL,
  `Comment` varchar(500) DEFAULT NULL,
  `Rating` int(11) DEFAULT NULL CHECK (`Rating` between 1 and 10),
  `PostDate` date NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

--
-- Dataark for tabell `review`
--

INSERT INTO `review` (`UserID`, `ProductID`, `Comment`, `Rating`, `PostDate`) VALUES
('34010147-490a-444a-ba96-9550b0126006', '628a074c-817f-4d69-ba30-6aa04b023b43', 'Not worth the money.', 1, '2025-03-26'),
('34010147-490a-444a-ba96-9550b0126006', 'fc84edd2-93ba-46c6-a453-61e974d0d8aa', 'It\'s okay, does the job.', 6, '2025-05-14'),
('4b8b39c2-e536-4e1b-985b-77e9989d3de4', '8ced2470-c515-403f-b4a1-69cb2c8434c1', 'Good, but could be better.', 6, '2025-02-19'),
('4b8b39c2-e536-4e1b-985b-77e9989d3de4', 'c522e1d3-1da6-4915-8431-e10d9bb35a2c', 'Good, but could be better.', 5, '2025-03-12'),
('c01ab0a2-5036-4273-a1c3-7f19d4aa257c', '0f27563b-9c7c-4fc9-8fe0-17c668f7d865', 'Highly recommend it!', 8, '2025-02-26'),
('c01ab0a2-5036-4273-a1c3-7f19d4aa257c', '4da637b6-f102-4d5e-b5cb-3bf209acd705', 'It\'s okay, does the job.', 7, '2025-02-21'),
('c01ab0a2-5036-4273-a1c3-7f19d4aa257c', 'aa185448-3e8d-49a2-9184-f5f3a6559f25', 'Good, but could be better.', 7, '2025-02-21'),
('d2054caf-1155-47a6-8791-3222c9d6beb8', '306ec8e0-6cd0-4af2-84f2-a2c446e59ea7', 'Amazing quality!', 10, '2025-02-10'),
('d2054caf-1155-47a6-8791-3222c9d6beb8', '41499fa1-f67a-4293-a6b3-bc469e045d6b', 'Average, nothing special.', 4, '2025-04-29'),
('d2054caf-1155-47a6-8791-3222c9d6beb8', '63e238fc-9a93-47e7-8ee0-df418a1f2aac', 'Exceeded my expectations.', 8, '2025-03-07');

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `users`
--

CREATE TABLE `users` (
  `UserID` varchar(50) NOT NULL,
  `Username` varchar(100) NOT NULL,
  `Password` varchar(100) NOT NULL,
  `Email` varchar(100) NOT NULL,
  `FirstName` varchar(50) NOT NULL,
  `LastName` varchar(50) NOT NULL,
  `Address` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

--
-- Dataark for tabell `users`
--

INSERT INTO `users` (`UserID`, `Username`, `Password`, `Email`, `FirstName`, `LastName`, `Address`) VALUES
('34010147-490a-444a-ba96-9550b0126006', 'alice123', '$2a$10$KJpUQ3eMtnYZOIjC9zK9dOQQGuJ6NGs60hlLoV8/kHWsDsLcfcdEG', 'alice@gmail.com', 'Alice', 'Smith', 'River Road 42'),
('4b8b39c2-e536-4e1b-985b-77e9989d3de4', 'bob88', '$2a$10$CaOXZVlTX8IUch9HzgMJGur2TjaUp2kDV/zJZOhsv/yCctaILGMYG', 'bob88@gmail.com', 'Bob', 'Johnson', 'Mountain View 10'),
('c01ab0a2-5036-4273-a1c3-7f19d4aa257c', 'helloWorld', '$2a$10$l0HAqaB9k3fXEK9zmPeaX.UMUlmrF6nusYMWx28a8ObqWmQ/tlcp2', 'john_doe@gmail.com', 'John', 'Doe', 'Yolostreet 15'),
('d2054caf-1155-47a6-8791-3222c9d6beb8', 'janedoe', '$2a$10$jsWXVECuVrT8tyqjXrN1JO9Hj.gtOQcI5tjtEWugqEcujONpImEM2', 'jane_doe@gmail.com', 'Jane', 'Doe', 'Main Street 5');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `administrators`
--
ALTER TABLE `administrators`
  ADD PRIMARY KEY (`UserID`);

--
-- Indexes for table `brand`
--
ALTER TABLE `brand`
  ADD PRIMARY KEY (`BrandName`);

--
-- Indexes for table `cartitem`
--
ALTER TABLE `cartitem`
  ADD PRIMARY KEY (`UserID`,`ProductID`),
  ADD KEY `ProductID` (`ProductID`);

--
-- Indexes for table `category`
--
ALTER TABLE `category`
  ADD PRIMARY KEY (`CategoryName`);

--
-- Indexes for table `members`
--
ALTER TABLE `members`
  ADD PRIMARY KEY (`UserID`);

--
-- Indexes for table `orderitem`
--
ALTER TABLE `orderitem`
  ADD PRIMARY KEY (`OrderID`,`ProductID`),
  ADD KEY `ProductID` (`ProductID`);

--
-- Indexes for table `orderstatus`
--
ALTER TABLE `orderstatus`
  ADD PRIMARY KEY (`StatusName`);

--
-- Indexes for table `ordertable`
--
ALTER TABLE `ordertable`
  ADD PRIMARY KEY (`OrderID`),
  ADD KEY `UserID` (`UserID`),
  ADD KEY `OrderStatus` (`OrderStatus`);

--
-- Indexes for table `payment`
--
ALTER TABLE `payment`
  ADD PRIMARY KEY (`PaymentID`),
  ADD KEY `OrderID` (`OrderID`);

--
-- Indexes for table `product`
--
ALTER TABLE `product`
  ADD PRIMARY KEY (`ProductID`),
  ADD KEY `Brand` (`Brand`),
  ADD KEY `Category` (`Category`);

--
-- Indexes for table `review`
--
ALTER TABLE `review`
  ADD PRIMARY KEY (`UserID`,`ProductID`),
  ADD KEY `ProductID` (`ProductID`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`UserID`);

--
-- Begrensninger for dumpede tabeller
--

--
-- Begrensninger for tabell `administrators`
--
ALTER TABLE `administrators`
  ADD CONSTRAINT `administrators_ibfk_1` FOREIGN KEY (`UserID`) REFERENCES `users` (`UserID`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Begrensninger for tabell `cartitem`
--
ALTER TABLE `cartitem`
  ADD CONSTRAINT `cartitem_ibfk_1` FOREIGN KEY (`UserID`) REFERENCES `users` (`UserID`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `cartitem_ibfk_2` FOREIGN KEY (`ProductID`) REFERENCES `product` (`ProductID`) ON UPDATE CASCADE;

--
-- Begrensninger for tabell `members`
--
ALTER TABLE `members`
  ADD CONSTRAINT `members_ibfk_1` FOREIGN KEY (`UserID`) REFERENCES `users` (`UserID`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Begrensninger for tabell `orderitem`
--
ALTER TABLE `orderitem`
  ADD CONSTRAINT `orderitem_ibfk_1` FOREIGN KEY (`OrderID`) REFERENCES `ordertable` (`OrderID`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `orderitem_ibfk_2` FOREIGN KEY (`ProductID`) REFERENCES `product` (`ProductID`) ON UPDATE CASCADE;

--
-- Begrensninger for tabell `ordertable`
--
ALTER TABLE `ordertable`
  ADD CONSTRAINT `ordertable_ibfk_1` FOREIGN KEY (`UserID`) REFERENCES `users` (`UserID`) ON UPDATE CASCADE,
  ADD CONSTRAINT `ordertable_ibfk_2` FOREIGN KEY (`OrderStatus`) REFERENCES `orderstatus` (`StatusName`) ON UPDATE CASCADE;

--
-- Begrensninger for tabell `payment`
--
ALTER TABLE `payment`
  ADD CONSTRAINT `payment_ibfk_1` FOREIGN KEY (`OrderID`) REFERENCES `ordertable` (`OrderID`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Begrensninger for tabell `product`
--
ALTER TABLE `product`
  ADD CONSTRAINT `product_ibfk_1` FOREIGN KEY (`Brand`) REFERENCES `brand` (`BrandName`) ON UPDATE CASCADE,
  ADD CONSTRAINT `product_ibfk_2` FOREIGN KEY (`Category`) REFERENCES `category` (`CategoryName`) ON UPDATE CASCADE;

--
-- Begrensninger for tabell `review`
--
ALTER TABLE `review`
  ADD CONSTRAINT `review_ibfk_1` FOREIGN KEY (`UserID`) REFERENCES `users` (`UserID`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `review_ibfk_2` FOREIGN KEY (`ProductID`) REFERENCES `product` (`ProductID`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
