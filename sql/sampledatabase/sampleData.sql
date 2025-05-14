-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: 14. Mai, 2025 13:21 PM
-- Tjener-versjon: 10.4.28-MariaDB
-- PHP Version: 8.2.4

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
-- Tabellstruktur for tabell `Administrators`
--

CREATE TABLE `Administrators` (
  `UserID` varchar(50) NOT NULL,
  `RoleName` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dataark for tabell `Administrators`
--

INSERT INTO `Administrators` (`UserID`, `RoleName`) VALUES
('e156ef59-bc21-4c72-9699-42a80be9f97a', 'admin');

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `Brand`
--

CREATE TABLE `Brand` (
  `BrandName` varchar(50) NOT NULL,
  `BrandDesc` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dataark for tabell `Brand`
--

INSERT INTO `Brand` (`BrandName`, `BrandDesc`) VALUES
('DigiTech', 'DigiTech powers your lifestyle with a full range of smart electronics — from smartphones to home appliances — all designed for seamless digital living.'),
('E-Gadgets', 'E-Gadget brings everyday innovation to your fingertips with a wide selection of electronics and accessories that make life simpler, smarter, and more exciting.'),
('TechBrave', 'TechBrave delivers bold, cutting-edge technology across smartphones, laptops, cameras, and more, empowering pioneers in every part of life.');

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `CartItem`
--

CREATE TABLE `CartItem` (
  `UserID` varchar(50) NOT NULL,
  `ProductID` varchar(50) NOT NULL,
  `Quantity` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dataark for tabell `CartItem`
--

INSERT INTO `CartItem` (`UserID`, `ProductID`, `Quantity`) VALUES
('2e20bf33-a493-4741-8047-99bcdbf7188d', '7fc631a0-0c4d-4e9b-9176-40064d92b1dc', 3),
('2e20bf33-a493-4741-8047-99bcdbf7188d', 'b920bf4a-b8f7-40b4-898e-1184700d5fb0', 2),
('9744e8eb-4c01-44c1-8a56-3a3767f1cfe1', '00ca2bb5-f537-418d-8c7e-db595a9e17b6', 1),
('9744e8eb-4c01-44c1-8a56-3a3767f1cfe1', '35107d01-c99b-43d3-8203-38ab776ad589', 2),
('9744e8eb-4c01-44c1-8a56-3a3767f1cfe1', 'dc6127b0-e5f8-4f0f-bdc8-afe6394217b8', 2),
('a3e8eabb-ee71-4892-ba06-7f011a42954c', '99c699bd-05ff-4e84-80ed-973fa3d8b7b6', 2),
('a3e8eabb-ee71-4892-ba06-7f011a42954c', 'c9ada754-017c-4762-95c4-893e8bc108f4', 2),
('e156ef59-bc21-4c72-9699-42a80be9f97a', '41af6698-6dff-44dc-b5a7-13c145adba57', 3),
('e156ef59-bc21-4c72-9699-42a80be9f97a', '98471267-f1c6-4306-b37f-dfdcec1c96e7', 3),
('e156ef59-bc21-4c72-9699-42a80be9f97a', 'b147738f-6f21-4565-b81e-000df750b4df', 1);

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `Category`
--

CREATE TABLE `Category` (
  `CategoryName` varchar(50) NOT NULL,
  `CategoryDesc` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dataark for tabell `Category`
--

INSERT INTO `Category` (`CategoryName`, `CategoryDesc`) VALUES
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
-- Tabellstruktur for tabell `Members`
--

CREATE TABLE `Members` (
  `UserID` varchar(50) NOT NULL,
  `MembershipLevel` enum('Silver','Gold','Platinum') NOT NULL,
  `MembershipStart` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dataark for tabell `Members`
--

INSERT INTO `Members` (`UserID`, `MembershipLevel`, `MembershipStart`) VALUES
('2e20bf33-a493-4741-8047-99bcdbf7188d', 'Platinum', '2025-05-14'),
('9744e8eb-4c01-44c1-8a56-3a3767f1cfe1', 'Silver', '2025-05-14'),
('e156ef59-bc21-4c72-9699-42a80be9f97a', 'Gold', '2025-05-14');

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `OrderItem`
--

CREATE TABLE `OrderItem` (
  `OrderID` varchar(50) NOT NULL,
  `ProductID` varchar(50) NOT NULL,
  `Quantity` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dataark for tabell `OrderItem`
--

INSERT INTO `OrderItem` (`OrderID`, `ProductID`, `Quantity`) VALUES
('04f652ff-d4b9-4755-ad74-946fadfdbefb', '502ef50e-11d4-4198-9edd-800b87833a28', 3),
('04f652ff-d4b9-4755-ad74-946fadfdbefb', '7fc631a0-0c4d-4e9b-9176-40064d92b1dc', 3),
('0688ed91-2b84-4d30-894b-665dbcf2b8df', '35107d01-c99b-43d3-8203-38ab776ad589', 1),
('0688ed91-2b84-4d30-894b-665dbcf2b8df', 'c4b1ad2a-98c8-4bf0-bb77-70c52c8a756d', 1),
('2d820e70-d46d-4da0-bb6e-039a11935fe7', '41af6698-6dff-44dc-b5a7-13c145adba57', 3),
('2d820e70-d46d-4da0-bb6e-039a11935fe7', '9d5acb94-b1a6-4d9c-b7c3-3ea080b74cac', 3),
('49138e06-cf6f-4bb2-bae1-d254f0455b3c', '1fdb86ac-1dd8-4022-a19a-a10aefef393b', 3),
('49138e06-cf6f-4bb2-bae1-d254f0455b3c', 'dc6127b0-e5f8-4f0f-bdc8-afe6394217b8', 3),
('8e012d47-ac88-435d-8983-dc750d550408', '00ca2bb5-f537-418d-8c7e-db595a9e17b6', 2),
('8e012d47-ac88-435d-8983-dc750d550408', 'e4d7426d-770f-4e4b-8177-024462a9913a', 1),
('bcdeb2bb-4f26-44cb-9659-cc193bc254ac', 'b920bf4a-b8f7-40b4-898e-1184700d5fb0', 1),
('bcdeb2bb-4f26-44cb-9659-cc193bc254ac', 'e1d61285-db0c-4d34-af41-828d7debc6ed', 2),
('d659d330-6aa7-4f35-beff-10d0e544cca5', '307b17d5-8dfb-4ccb-aafe-c5bc30f4ddaa', 2),
('d659d330-6aa7-4f35-beff-10d0e544cca5', '98471267-f1c6-4306-b37f-dfdcec1c96e7', 2),
('e811443c-1822-47c6-8fb2-2f40da3278b2', '864f398f-301b-434a-87a9-ff56a0d8be56', 1),
('e811443c-1822-47c6-8fb2-2f40da3278b2', 'c9ada754-017c-4762-95c4-893e8bc108f4', 1),
('ea5c094c-a603-4c8e-a724-9582f7f47099', '596b777a-74ba-4919-ab2e-53d61fc006c0', 2),
('ea5c094c-a603-4c8e-a724-9582f7f47099', '99c699bd-05ff-4e84-80ed-973fa3d8b7b6', 1),
('ea969b83-9f44-44fa-887f-cb00398c94e3', 'a628c7ce-f23b-40fe-a8b9-ac6ec650b623', 1),
('ea969b83-9f44-44fa-887f-cb00398c94e3', 'b147738f-6f21-4565-b81e-000df750b4df', 2);

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `OrderStatus`
--

CREATE TABLE `OrderStatus` (
  `StatusName` varchar(50) NOT NULL,
  `StatusDesc` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dataark for tabell `OrderStatus`
--

INSERT INTO `OrderStatus` (`StatusName`, `StatusDesc`) VALUES
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
-- Tabellstruktur for tabell `OrderTable`
--

CREATE TABLE `OrderTable` (
  `OrderID` varchar(50) NOT NULL,
  `UserID` varchar(50) NOT NULL,
  `OrderDate` date DEFAULT NULL,
  `OrderStatus` varchar(50) NOT NULL,
  `OrderTotal` decimal(10,2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dataark for tabell `OrderTable`
--

INSERT INTO `OrderTable` (`OrderID`, `UserID`, `OrderDate`, `OrderStatus`, `OrderTotal`) VALUES
('04f652ff-d4b9-4755-ad74-946fadfdbefb', '2e20bf33-a493-4741-8047-99bcdbf7188d', '2025-03-02', 'Refunded', 12899.94),
('0688ed91-2b84-4d30-894b-665dbcf2b8df', '9744e8eb-4c01-44c1-8a56-3a3767f1cfe1', '2025-03-08', 'Processing', 18999.98),
('2d820e70-d46d-4da0-bb6e-039a11935fe7', 'e156ef59-bc21-4c72-9699-42a80be9f97a', '2025-02-17', 'Pending', 29999.94),
('49138e06-cf6f-4bb2-bae1-d254f0455b3c', '9744e8eb-4c01-44c1-8a56-3a3767f1cfe1', '2025-03-19', 'Processing', 41399.94),
('8e012d47-ac88-435d-8983-dc750d550408', '9744e8eb-4c01-44c1-8a56-3a3767f1cfe1', '2025-02-17', 'Returned', 2499.97),
('bcdeb2bb-4f26-44cb-9659-cc193bc254ac', '2e20bf33-a493-4741-8047-99bcdbf7188d', '2025-03-31', 'Shipped', 19999.97),
('d659d330-6aa7-4f35-beff-10d0e544cca5', 'e156ef59-bc21-4c72-9699-42a80be9f97a', '2025-02-21', 'Cancelled', 5999.96),
('e811443c-1822-47c6-8fb2-2f40da3278b2', 'a3e8eabb-ee71-4892-ba06-7f011a42954c', '2025-02-13', 'Failed', 6499.98),
('ea5c094c-a603-4c8e-a724-9582f7f47099', 'a3e8eabb-ee71-4892-ba06-7f011a42954c', '2025-03-27', 'Delivered', 12999.97),
('ea969b83-9f44-44fa-887f-cb00398c94e3', 'e156ef59-bc21-4c72-9699-42a80be9f97a', '2025-02-23', 'Pending', 1699.97);

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `Payment`
--

CREATE TABLE `Payment` (
  `PaymentID` varchar(50) NOT NULL,
  `OrderID` varchar(50) NOT NULL,
  `PaymentMethod` enum('card','vipps','bank_transfer') NOT NULL,
  `Amount` decimal(10,2) NOT NULL,
  `PaymentDate` date NOT NULL,
  `PaymentStatus` enum('pending','successful','failed') DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dataark for tabell `Payment`
--

INSERT INTO `Payment` (`PaymentID`, `OrderID`, `PaymentMethod`, `Amount`, `PaymentDate`, `PaymentStatus`) VALUES
('0173108d-675f-43b0-b0bc-648c48f88de8', '8e012d47-ac88-435d-8983-dc750d550408', 'bank_transfer', 2499.97, '2025-04-28', 'failed'),
('29762b35-2d99-41cd-9cc0-e3cbb0cedb04', 'ea969b83-9f44-44fa-887f-cb00398c94e3', 'bank_transfer', 1699.97, '2025-04-06', 'pending'),
('32a7e313-7447-4132-aeeb-76a251972c79', 'bcdeb2bb-4f26-44cb-9659-cc193bc254ac', 'bank_transfer', 19999.97, '2025-02-20', 'successful'),
('36c36e71-a91a-4b09-b726-809cdc55bfce', '04f652ff-d4b9-4755-ad74-946fadfdbefb', 'card', 12899.94, '2025-02-18', 'failed'),
('461d12a6-26f3-460b-b65f-5aa8fa32eb60', '49138e06-cf6f-4bb2-bae1-d254f0455b3c', 'vipps', 41399.94, '2025-04-30', 'pending'),
('4b74408c-4165-4c59-9c39-7b2ac0a02d75', '2d820e70-d46d-4da0-bb6e-039a11935fe7', 'card', 29999.94, '2025-04-19', 'pending'),
('b19e8338-c47d-4b75-8c5b-46a7698f4306', '0688ed91-2b84-4d30-894b-665dbcf2b8df', 'card', 18999.98, '2025-03-21', 'pending'),
('c05faedd-d456-462c-a4d3-15955eb63253', 'd659d330-6aa7-4f35-beff-10d0e544cca5', 'vipps', 5999.96, '2025-04-30', 'failed'),
('c6f9c513-1528-47dd-becb-948e6b258fb8', 'e811443c-1822-47c6-8fb2-2f40da3278b2', 'vipps', 6499.98, '2025-04-14', 'failed'),
('de0f36a0-9fbc-4388-8abc-296012b981bf', 'ea5c094c-a603-4c8e-a724-9582f7f47099', 'card', 12999.97, '2025-04-09', 'successful');

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `Product`
--

CREATE TABLE `Product` (
  `ProductID` varchar(50) NOT NULL,
  `ProductName` varchar(100) NOT NULL,
  `ProductDesc` varchar(255) DEFAULT NULL,
  `ProductImgUrl` varchar(255) DEFAULT NULL,
  `Price` decimal(10,2) NOT NULL,
  `StockQuantity` int(11) NOT NULL,
  `Brand` varchar(50) DEFAULT NULL,
  `Category` varchar(50) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dataark for tabell `Product`
--

INSERT INTO `Product` (`ProductID`, `ProductName`, `ProductDesc`, `ProductImgUrl`, `Price`, `StockQuantity`, `Brand`, `Category`) VALUES
('00ca2bb5-f537-418d-8c7e-db595a9e17b6', 'DigiTech Bluetooth Speaker', 'Portable speaker with rich bass.', 'https://example.com/img6.jpg', 899.99, 200, 'DigiTech', 'Audio & Headphones'),
('1fdb86ac-1dd8-4022-a19a-a10aefef393b', 'DigiTech Gaming Mouse', 'Precision gaming mouse.', 'https://example.com/img12.jpg', 799.99, 120, 'DigiTech', 'Gaming'),
('307b17d5-8dfb-4ccb-aafe-c5bc30f4ddaa', 'DigiTech Portable SSD 1TB', 'High-speed portable storage.', 'https://example.com/img15.jpg', 1499.99, 70, 'DigiTech', 'Computers & Accessories'),
('35107d01-c99b-43d3-8203-38ab776ad589', 'TechBrave Laptop Pro', 'Powerful laptop for professionals.', 'https://example.com/img10.jpg', 15999.99, 40, 'TechBrave', 'Computers & Accessories'),
('41af6698-6dff-44dc-b5a7-13c145adba57', 'TechBrave X100 Smartphone', 'High-end smartphone with AI features.', 'https://example.com/img1.jpg', 8999.99, 100, NULL, 'Mobile Phones & Tablets'),
('502ef50e-11d4-4198-9edd-800b87833a28', 'E-Gadgets Smart Light Bulb', 'WiFi-enabled color-changing bulb.', 'https://example.com/img17.jpg', 299.99, 500, 'E-Gadgets', 'Smart Home'),
('596b777a-74ba-4919-ab2e-53d61fc006c0', 'E-Gadgets 4K Projector', 'Cinema-quality projector for home.', 'https://example.com/img14.jpg', 4999.99, 45, 'E-Gadgets', 'TV & Home Theater'),
('7fc631a0-0c4d-4e9b-9176-40064d92b1dc', 'TechBrave Tablet 10\"', 'Lightweight tablet for work and play.', 'https://example.com/img7.jpg', 3999.99, 90, 'TechBrave', 'Mobile Phones & Tablets'),
('864f398f-301b-434a-87a9-ff56a0d8be56', 'DigiTech VR Headset', 'Immersive VR gaming experience.', 'https://example.com/img18.jpg', 3999.99, 60, 'DigiTech', 'Gaming'),
('98471267-f1c6-4306-b37f-dfdcec1c96e7', 'E-Gadgets Home Assistant', 'Voice-controlled smart home hub.', 'https://example.com/img5.jpg', 1499.99, 50, 'E-Gadgets', 'Smart Home'),
('99c699bd-05ff-4e84-80ed-973fa3d8b7b6', 'TechBrave Smartwatch 2', 'Fitness tracking smartwatch.', 'https://example.com/img4.jpg', 2999.99, 150, 'TechBrave', 'Wearable Technology'),
('9d5acb94-b1a6-4d9c-b7c3-3ea080b74cac', 'E-Gadgets Security Camera', 'Smart WiFi home security camera.', 'https://example.com/img11.jpg', 999.99, 100, 'E-Gadgets', 'Smart Home'),
('a628c7ce-f23b-40fe-a8b9-ac6ec650b623', 'TechBrave Power Bank 20000mAh', 'Portable charger for mobile devices.', 'https://example.com/img19.jpg', 499.99, 250, 'TechBrave', 'Mobile Phones & Tablets'),
('b147738f-6f21-4565-b81e-000df750b4df', 'DigiTech Wireless Charger', 'Fast wireless charging pad.', 'https://example.com/img9.jpg', 599.99, 300, 'DigiTech', 'Mobile Phones & Tablets'),
('b920bf4a-b8f7-40b4-898e-1184700d5fb0', 'DigiTech Gaming Headset Pro', 'Immersive gaming headset with surround sound.', 'https://example.com/img3.jpg', 1999.99, 80, 'DigiTech', 'Audio & Headphones'),
('c4b1ad2a-98c8-4bf0-bb77-70c52c8a756d', 'E-Gadgets Robot Vacuum', 'Smart robot vacuum cleaner.', 'https://example.com/img20.jpg', 2999.99, 30, 'E-Gadgets', 'Smart Home'),
('c9ada754-017c-4762-95c4-893e8bc108f4', 'E-Gadgets Noise Cancelling Headphones', 'Noise-free music experience.', 'https://example.com/img8.jpg', 2499.99, 70, 'E-Gadgets', 'Audio & Headphones'),
('dc6127b0-e5f8-4f0f-bdc8-afe6394217b8', 'E-Gadgets Ultra HD TV 55\"', 'Ultra HD smart TV with vibrant colors.', 'https://example.com/img2.jpg', 12999.99, 30, 'E-Gadgets', 'TV & Home Theater'),
('e1d61285-db0c-4d34-af41-828d7debc6ed', 'TechBrave Drone Pro', 'Professional drone with 4K camera.', 'https://example.com/img13.jpg', 8999.99, 25, 'TechBrave', 'Cameras & Photography'),
('e4d7426d-770f-4e4b-8177-024462a9913a', 'TechBrave Fitness Band', 'Basic fitness tracker with heart monitor.', 'https://example.com/img16.jpg', 699.99, 110, 'TechBrave', 'Wearable Technology');

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `Review`
--

CREATE TABLE `Review` (
  `UserID` varchar(50) NOT NULL,
  `ProductID` varchar(50) NOT NULL,
  `Comment` varchar(500) DEFAULT NULL,
  `Rating` int(11) DEFAULT NULL CHECK (`Rating` between 1 and 10),
  `PostDate` date NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dataark for tabell `Review`
--

INSERT INTO `Review` (`UserID`, `ProductID`, `Comment`, `Rating`, `PostDate`) VALUES
('2e20bf33-a493-4741-8047-99bcdbf7188d', '7fc631a0-0c4d-4e9b-9176-40064d92b1dc', 'Good, but could be better.', 4, '2025-04-30'),
('2e20bf33-a493-4741-8047-99bcdbf7188d', 'b920bf4a-b8f7-40b4-898e-1184700d5fb0', 'Terrible quality!', 2, '2025-03-27'),
('9744e8eb-4c01-44c1-8a56-3a3767f1cfe1', '00ca2bb5-f537-418d-8c7e-db595a9e17b6', 'Very disappointing.', 2, '2025-04-19'),
('9744e8eb-4c01-44c1-8a56-3a3767f1cfe1', '35107d01-c99b-43d3-8203-38ab776ad589', 'Fantastic product!', 10, '2025-04-12'),
('9744e8eb-4c01-44c1-8a56-3a3767f1cfe1', 'dc6127b0-e5f8-4f0f-bdc8-afe6394217b8', 'Average, nothing special.', 4, '2025-03-09'),
('a3e8eabb-ee71-4892-ba06-7f011a42954c', '99c699bd-05ff-4e84-80ed-973fa3d8b7b6', 'Average, nothing special.', 6, '2025-03-13'),
('a3e8eabb-ee71-4892-ba06-7f011a42954c', 'c9ada754-017c-4762-95c4-893e8bc108f4', 'Broke after a few days.', 3, '2025-04-25'),
('e156ef59-bc21-4c72-9699-42a80be9f97a', '41af6698-6dff-44dc-b5a7-13c145adba57', 'Exceeded my expectations.', 8, '2025-04-09'),
('e156ef59-bc21-4c72-9699-42a80be9f97a', '98471267-f1c6-4306-b37f-dfdcec1c96e7', 'Good, but could be better.', 4, '2025-05-12'),
('e156ef59-bc21-4c72-9699-42a80be9f97a', 'b147738f-6f21-4565-b81e-000df750b4df', 'Highly recommend it!', 8, '2025-03-10');

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `Users`
--

CREATE TABLE `Users` (
  `UserID` varchar(50) NOT NULL,
  `Username` varchar(100) NOT NULL,
  `Password` varchar(100) NOT NULL,
  `Email` varchar(100) NOT NULL,
  `FirstName` varchar(50) NOT NULL,
  `LastName` varchar(50) NOT NULL,
  `Address` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dataark for tabell `Users`
--

INSERT INTO `Users` (`UserID`, `Username`, `Password`, `Email`, `FirstName`, `LastName`, `Address`) VALUES
('2e20bf33-a493-4741-8047-99bcdbf7188d', 'alice123', '$2a$10$9ePvNDa2g8HqjC6oj6zQkejE0NMvcaj0EkyhJf1rmTM/m8pKjA0GS', 'alice@gmail.com', 'Alice', 'Smith', 'River Road 42'),
('9744e8eb-4c01-44c1-8a56-3a3767f1cfe1', 'janedoe', '$2a$10$.9cKdixENsUjSAg32OC8x.xRQZeHJLm4ORXTez.HQfgEpuqw5FLKO', 'jane_doe@gmail.com', 'Jane', 'Doe', 'Main Street 5'),
('a3e8eabb-ee71-4892-ba06-7f011a42954c', 'bob88', '$2a$10$i..HvUxZiN.oIugZcmdSaOJTR9Qm98I2AJUzVl3tOYbd4lWlPsGeS', 'bob88@gmail.com', 'Bob', 'Johnson', 'Mountain View 10'),
('e156ef59-bc21-4c72-9699-42a80be9f97a', 'helloWorld', '$2a$10$iPefuaAX/QfRCW55rVBIkOlPntB9CSvRrZ6jWICE4SU8J0P/R06jG', 'john_doe@gmail.com', 'John', 'Doe', 'Yolostreet 15');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `Administrators`
--
ALTER TABLE `Administrators`
  ADD PRIMARY KEY (`UserID`);

--
-- Indexes for table `Brand`
--
ALTER TABLE `Brand`
  ADD PRIMARY KEY (`BrandName`);

--
-- Indexes for table `CartItem`
--
ALTER TABLE `CartItem`
  ADD PRIMARY KEY (`UserID`,`ProductID`),
  ADD KEY `ProductID` (`ProductID`);

--
-- Indexes for table `Category`
--
ALTER TABLE `Category`
  ADD PRIMARY KEY (`CategoryName`);

--
-- Indexes for table `Members`
--
ALTER TABLE `Members`
  ADD PRIMARY KEY (`UserID`);

--
-- Indexes for table `OrderItem`
--
ALTER TABLE `OrderItem`
  ADD PRIMARY KEY (`OrderID`,`ProductID`),
  ADD KEY `ProductID` (`ProductID`);

--
-- Indexes for table `OrderStatus`
--
ALTER TABLE `OrderStatus`
  ADD PRIMARY KEY (`StatusName`);

--
-- Indexes for table `OrderTable`
--
ALTER TABLE `OrderTable`
  ADD PRIMARY KEY (`OrderID`),
  ADD KEY `UserID` (`UserID`),
  ADD KEY `OrderStatus` (`OrderStatus`);

--
-- Indexes for table `Payment`
--
ALTER TABLE `Payment`
  ADD PRIMARY KEY (`PaymentID`),
  ADD KEY `OrderID` (`OrderID`);

--
-- Indexes for table `Product`
--
ALTER TABLE `Product`
  ADD PRIMARY KEY (`ProductID`),
  ADD KEY `Brand` (`Brand`),
  ADD KEY `Category` (`Category`);

--
-- Indexes for table `Review`
--
ALTER TABLE `Review`
  ADD PRIMARY KEY (`UserID`,`ProductID`),
  ADD KEY `ProductID` (`ProductID`);

--
-- Indexes for table `Users`
--
ALTER TABLE `Users`
  ADD PRIMARY KEY (`UserID`);

--
-- Begrensninger for dumpede tabeller
--

--
-- Begrensninger for tabell `Administrators`
--
ALTER TABLE `Administrators`
  ADD CONSTRAINT `Administrators_ibfk_1` FOREIGN KEY (`UserID`) REFERENCES `Users` (`UserID`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Begrensninger for tabell `CartItem`
--
ALTER TABLE `CartItem`
  ADD CONSTRAINT `CartItem_ibfk_1` FOREIGN KEY (`UserID`) REFERENCES `Users` (`UserID`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `CartItem_ibfk_2` FOREIGN KEY (`ProductID`) REFERENCES `Product` (`ProductID`) ON UPDATE CASCADE;

--
-- Begrensninger for tabell `Members`
--
ALTER TABLE `Members`
  ADD CONSTRAINT `Members_ibfk_1` FOREIGN KEY (`UserID`) REFERENCES `Users` (`UserID`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Begrensninger for tabell `OrderItem`
--
ALTER TABLE `OrderItem`
  ADD CONSTRAINT `OrderItem_ibfk_1` FOREIGN KEY (`OrderID`) REFERENCES `OrderTable` (`OrderID`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `OrderItem_ibfk_2` FOREIGN KEY (`ProductID`) REFERENCES `Product` (`ProductID`) ON UPDATE CASCADE;

--
-- Begrensninger for tabell `OrderTable`
--
ALTER TABLE `OrderTable`
  ADD CONSTRAINT `OrderTable_ibfk_1` FOREIGN KEY (`UserID`) REFERENCES `Users` (`UserID`) ON UPDATE CASCADE,
  ADD CONSTRAINT `OrderTable_ibfk_2` FOREIGN KEY (`OrderStatus`) REFERENCES `OrderStatus` (`StatusName`) ON UPDATE CASCADE;

--
-- Begrensninger for tabell `Payment`
--
ALTER TABLE `Payment`
  ADD CONSTRAINT `Payment_ibfk_1` FOREIGN KEY (`OrderID`) REFERENCES `OrderTable` (`OrderID`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Begrensninger for tabell `Product`
--
ALTER TABLE `Product`
  ADD CONSTRAINT `Product_ibfk_1` FOREIGN KEY (`Brand`) REFERENCES `Brand` (`BrandName`) ON UPDATE CASCADE,
  ADD CONSTRAINT `Product_ibfk_2` FOREIGN KEY (`Category`) REFERENCES `Category` (`CategoryName`) ON UPDATE CASCADE;

--
-- Begrensninger for tabell `Review`
--
ALTER TABLE `Review`
  ADD CONSTRAINT `Review_ibfk_1` FOREIGN KEY (`UserID`) REFERENCES `Users` (`UserID`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `Review_ibfk_2` FOREIGN KEY (`ProductID`) REFERENCES `Product` (`ProductID`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
