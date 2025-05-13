-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: 13. Mai, 2025 17:33 PM
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
('96116344-b23e-49cf-868a-c67656ec17d7', 'admin');

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
('1a85843f-cfd1-44b6-972a-7d2f1c5e8529', '7fcaa398-4e28-4c6f-9bb4-b5386e7259c4', 2),
('1a85843f-cfd1-44b6-972a-7d2f1c5e8529', 'd84e12b7-4b48-4727-b1eb-dc9d9006cedc', 1),
('1a85843f-cfd1-44b6-972a-7d2f1c5e8529', 'f8028536-7055-4c3f-aea1-2f7f8fa2102d', 2),
('1b23667e-3f65-496c-92df-00de0ad36931', '56c85d36-342a-470f-8a90-482ab72774d1', 1),
('1b23667e-3f65-496c-92df-00de0ad36931', 'bd71a543-ce71-4fd1-bd45-da411f9b91fc', 2),
('96116344-b23e-49cf-868a-c67656ec17d7', '108c322d-d851-40e2-9f1c-e003fbceea3a', 2),
('96116344-b23e-49cf-868a-c67656ec17d7', 'a7d3632d-58fb-4627-a5d3-d8699c80256c', 2),
('96116344-b23e-49cf-868a-c67656ec17d7', 'd24e3e4d-5e1f-495b-aaa8-f9347bbbd6d0', 2),
('bd25f99d-e779-4eca-af91-cd7507d987ed', '28633bd5-b5ea-47c4-b8d7-b096cb2be139', 2),
('bd25f99d-e779-4eca-af91-cd7507d987ed', '9a8c7da7-c5f5-4676-8743-37aafdee8c02', 3);

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
('1a85843f-cfd1-44b6-972a-7d2f1c5e8529', 'Silver', '2025-05-13'),
('96116344-b23e-49cf-868a-c67656ec17d7', 'Gold', '2025-05-13'),
('bd25f99d-e779-4eca-af91-cd7507d987ed', 'Platinum', '2025-05-13');

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
('02734e5d-c21d-4a34-95f2-852e89c5af00', '56c85d36-342a-470f-8a90-482ab72774d1', 3),
('02734e5d-c21d-4a34-95f2-852e89c5af00', '8946f411-b179-4d73-bc9c-6b75cac1c738', 2),
('0833dc39-917c-4097-93d7-8398547aef0d', '4676c18d-433a-4002-bdb6-97c0dbb93ee3', 2),
('0833dc39-917c-4097-93d7-8398547aef0d', 'd84e12b7-4b48-4727-b1eb-dc9d9006cedc', 3),
('18b0fd98-fd82-43f8-826c-66173fbe1345', 'a7d3632d-58fb-4627-a5d3-d8699c80256c', 3),
('18b0fd98-fd82-43f8-826c-66173fbe1345', 'f2486224-7101-44a2-bfd1-23b032eb1d41', 3),
('4bde1d53-afab-4cbd-967c-ff1c2630e935', '108c322d-d851-40e2-9f1c-e003fbceea3a', 3),
('4bde1d53-afab-4cbd-967c-ff1c2630e935', 'e1d19c6d-da9b-455e-b351-e20c015706da', 3),
('682398d4-ef70-4947-8d64-f7a566d51221', '65772b1e-d9e6-4cee-a2f3-879f0a8c409c', 2),
('682398d4-ef70-4947-8d64-f7a566d51221', 'd24e3e4d-5e1f-495b-aaa8-f9347bbbd6d0', 1),
('b08732de-9f1b-4401-9183-d934c972d7cd', '49740001-b32a-4ac7-81b3-3e342b59929b', 2),
('b08732de-9f1b-4401-9183-d934c972d7cd', '7fcaa398-4e28-4c6f-9bb4-b5386e7259c4', 1),
('b8c2c8d7-497a-4cbb-b454-813b68af3d2a', 'b8c71477-5ae3-4ec5-924f-e087aed135c0', 3),
('b8c2c8d7-497a-4cbb-b454-813b68af3d2a', 'bd71a543-ce71-4fd1-bd45-da411f9b91fc', 1),
('ba5d8cfe-fc31-422c-b14e-f6ad410023db', '9a8c7da7-c5f5-4676-8743-37aafdee8c02', 3),
('ba5d8cfe-fc31-422c-b14e-f6ad410023db', 'be29513b-b9e7-4df7-8f4e-745ac2f82fcf', 3),
('c51b335a-7e4c-4859-af3e-90aface87137', '28633bd5-b5ea-47c4-b8d7-b096cb2be139', 2),
('c51b335a-7e4c-4859-af3e-90aface87137', 'b0c1aee9-b4f8-4c74-80bd-d8d61d9f0704', 3),
('e758392d-c869-4204-ad95-9afee9e10866', '1c184f45-dbec-4e34-b070-c13359798f19', 3),
('e758392d-c869-4204-ad95-9afee9e10866', 'f8028536-7055-4c3f-aea1-2f7f8fa2102d', 3);

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
  `UserID` varchar(50) NOT NULL,
  `OrderDate` date DEFAULT NULL,
  `OrderStatus` varchar(50) NOT NULL,
  `OrderTotal` decimal(10,2) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;

--
-- Dataark for tabell `ordertable`
--

INSERT INTO `ordertable` (`OrderID`, `UserID`, `OrderDate`, `OrderStatus`, `OrderTotal`) VALUES
('02734e5d-c21d-4a34-95f2-852e89c5af00', '1b23667e-3f65-496c-92df-00de0ad36931', '2025-04-30', 'Failed', 15499.95),
('0833dc39-917c-4097-93d7-8398547aef0d', '1a85843f-cfd1-44b6-972a-7d2f1c5e8529', '2025-04-14', 'Processing', 40599.95),
('18b0fd98-fd82-43f8-826c-66173fbe1345', '96116344-b23e-49cf-868a-c67656ec17d7', '2025-02-18', 'Cancelled', 8999.94),
('4bde1d53-afab-4cbd-967c-ff1c2630e935', '96116344-b23e-49cf-868a-c67656ec17d7', '2025-03-28', 'Pending', 29999.94),
('682398d4-ef70-4947-8d64-f7a566d51221', '96116344-b23e-49cf-868a-c67656ec17d7', '2025-05-08', 'Pending', 1599.97),
('b08732de-9f1b-4401-9183-d934c972d7cd', '1a85843f-cfd1-44b6-972a-7d2f1c5e8529', '2025-02-04', 'Returned', 2299.97),
('b8c2c8d7-497a-4cbb-b454-813b68af3d2a', '1b23667e-3f65-496c-92df-00de0ad36931', '2025-04-29', 'Delivered', 17999.96),
('ba5d8cfe-fc31-422c-b14e-f6ad410023db', 'bd25f99d-e779-4eca-af91-cd7507d987ed', '2025-03-06', 'Shipped', 32999.94),
('c51b335a-7e4c-4859-af3e-90aface87137', 'bd25f99d-e779-4eca-af91-cd7507d987ed', '2025-02-13', 'Refunded', 8899.95),
('e758392d-c869-4204-ad95-9afee9e10866', '1a85843f-cfd1-44b6-972a-7d2f1c5e8529', '2025-03-24', 'Processing', 56999.94);

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
('07930cbc-061d-4b48-a3fe-446b3b44fab4', '682398d4-ef70-4947-8d64-f7a566d51221', 'bank_transfer', 1599.97, '2025-02-26', 'pending'),
('268a59ae-2765-462b-81e9-23b4538cfdf5', 'b08732de-9f1b-4401-9183-d934c972d7cd', 'bank_transfer', 2299.97, '2025-02-18', 'failed'),
('3563f618-fe17-46f7-af5a-fac1f13e3947', '0833dc39-917c-4097-93d7-8398547aef0d', 'vipps', 40599.95, '2025-05-07', 'pending'),
('56009f0a-086b-4995-a1f9-dc7d1b83044c', '18b0fd98-fd82-43f8-826c-66173fbe1345', 'vipps', 8999.94, '2025-04-18', 'failed'),
('5c0f49b1-1790-49d5-a05d-6376961d96c1', '4bde1d53-afab-4cbd-967c-ff1c2630e935', 'card', 29999.94, '2025-03-08', 'pending'),
('6483b204-49e8-4416-9e60-8210d2828c6b', 'ba5d8cfe-fc31-422c-b14e-f6ad410023db', 'bank_transfer', 32999.94, '2025-02-26', 'successful'),
('a6bed5e3-f642-4f6b-98bc-3fe438f27984', 'e758392d-c869-4204-ad95-9afee9e10866', 'card', 56999.94, '2025-04-20', 'pending'),
('a70e6e03-cef4-41cc-92d4-17521868f7d2', 'b8c2c8d7-497a-4cbb-b454-813b68af3d2a', 'card', 17999.96, '2025-03-16', 'successful'),
('c624f74c-835a-47a4-97b1-e75f252f6696', 'c51b335a-7e4c-4859-af3e-90aface87137', 'card', 8899.95, '2025-04-10', 'failed'),
('f700d8bd-c0db-4c1f-b9bc-c9b483edff07', '02734e5d-c21d-4a34-95f2-852e89c5af00', 'vipps', 15499.95, '2025-04-04', 'failed');

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
('108c322d-d851-40e2-9f1c-e003fbceea3a', 'TechBrave X100 Smartphone', 'High-end smartphone with AI features.', 'https://example.com/img1.jpg', 8999.99, 100, NULL, 'Mobile Phones & Tablets'),
('1c184f45-dbec-4e34-b070-c13359798f19', 'E-Gadgets Robot Vacuum', 'Smart robot vacuum cleaner.', 'https://example.com/img20.jpg', 2999.99, 30, 'E-Gadgets', 'Smart Home'),
('28633bd5-b5ea-47c4-b8d7-b096cb2be139', 'TechBrave Tablet 10\"', 'Lightweight tablet for work and play.', 'https://example.com/img7.jpg', 3999.99, 90, 'TechBrave', 'Mobile Phones & Tablets'),
('4676c18d-433a-4002-bdb6-97c0dbb93ee3', 'DigiTech Gaming Mouse', 'Precision gaming mouse.', 'https://example.com/img12.jpg', 799.99, 120, 'DigiTech', 'Gaming'),
('49740001-b32a-4ac7-81b3-3e342b59929b', 'TechBrave Fitness Band', 'Basic fitness tracker with heart monitor.', 'https://example.com/img16.jpg', 699.99, 110, 'TechBrave', 'Wearable Technology'),
('56c85d36-342a-470f-8a90-482ab72774d1', 'E-Gadgets Noise Cancelling Headphones', 'Noise-free music experience.', 'https://example.com/img8.jpg', 2499.99, 70, 'E-Gadgets', 'Audio & Headphones'),
('65772b1e-d9e6-4cee-a2f3-879f0a8c409c', 'TechBrave Power Bank 20000mAh', 'Portable charger for mobile devices.', 'https://example.com/img19.jpg', 499.99, 250, 'TechBrave', 'Mobile Phones & Tablets'),
('7fcaa398-4e28-4c6f-9bb4-b5386e7259c4', 'DigiTech Bluetooth Speaker', 'Portable speaker with rich bass.', 'https://example.com/img6.jpg', 899.99, 200, 'DigiTech', 'Audio & Headphones'),
('8946f411-b179-4d73-bc9c-6b75cac1c738', 'DigiTech VR Headset', 'Immersive VR gaming experience.', 'https://example.com/img18.jpg', 3999.99, 60, 'DigiTech', 'Gaming'),
('9a8c7da7-c5f5-4676-8743-37aafdee8c02', 'DigiTech Gaming Headset Pro', 'Immersive gaming headset with surround sound.', 'https://example.com/img3.jpg', 1999.99, 80, 'DigiTech', 'Audio & Headphones'),
('a7d3632d-58fb-4627-a5d3-d8699c80256c', 'E-Gadgets Home Assistant', 'Voice-controlled smart home hub.', 'https://example.com/img5.jpg', 1499.99, 50, 'E-Gadgets', 'Smart Home'),
('b0c1aee9-b4f8-4c74-80bd-d8d61d9f0704', 'E-Gadgets Smart Light Bulb', 'WiFi-enabled color-changing bulb.', 'https://example.com/img17.jpg', 299.99, 500, 'E-Gadgets', 'Smart Home'),
('b8c71477-5ae3-4ec5-924f-e087aed135c0', 'E-Gadgets 4K Projector', 'Cinema-quality projector for home.', 'https://example.com/img14.jpg', 4999.99, 45, 'E-Gadgets', 'TV & Home Theater'),
('bd71a543-ce71-4fd1-bd45-da411f9b91fc', 'TechBrave Smartwatch 2', 'Fitness tracking smartwatch.', 'https://example.com/img4.jpg', 2999.99, 150, 'TechBrave', 'Wearable Technology'),
('be29513b-b9e7-4df7-8f4e-745ac2f82fcf', 'TechBrave Drone Pro', 'Professional drone with 4K camera.', 'https://example.com/img13.jpg', 8999.99, 25, 'TechBrave', 'Cameras & Photography'),
('d24e3e4d-5e1f-495b-aaa8-f9347bbbd6d0', 'DigiTech Wireless Charger', 'Fast wireless charging pad.', 'https://example.com/img9.jpg', 599.99, 300, 'DigiTech', 'Mobile Phones & Tablets'),
('d84e12b7-4b48-4727-b1eb-dc9d9006cedc', 'E-Gadgets Ultra HD TV 55\"', 'Ultra HD smart TV with vibrant colors.', 'https://example.com/img2.jpg', 12999.99, 30, 'E-Gadgets', 'TV & Home Theater'),
('e1d19c6d-da9b-455e-b351-e20c015706da', 'E-Gadgets Security Camera', 'Smart WiFi home security camera.', 'https://example.com/img11.jpg', 999.99, 100, 'E-Gadgets', 'Smart Home'),
('f2486224-7101-44a2-bfd1-23b032eb1d41', 'DigiTech Portable SSD 1TB', 'High-speed portable storage.', 'https://example.com/img15.jpg', 1499.99, 70, 'DigiTech', 'Computers & Accessories'),
('f8028536-7055-4c3f-aea1-2f7f8fa2102d', 'TechBrave Laptop Pro', 'Powerful laptop for professionals.', 'https://example.com/img10.jpg', 15999.99, 40, 'TechBrave', 'Computers & Accessories');

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
('1a85843f-cfd1-44b6-972a-7d2f1c5e8529', '7fcaa398-4e28-4c6f-9bb4-b5386e7259c4', 'Good, but could be better.', 4, '2025-05-02'),
('1a85843f-cfd1-44b6-972a-7d2f1c5e8529', 'd84e12b7-4b48-4727-b1eb-dc9d9006cedc', 'Very disappointing.', 2, '2025-05-06'),
('1a85843f-cfd1-44b6-972a-7d2f1c5e8529', 'f8028536-7055-4c3f-aea1-2f7f8fa2102d', 'Amazing quality!', 9, '2025-02-04'),
('1b23667e-3f65-496c-92df-00de0ad36931', '56c85d36-342a-470f-8a90-482ab72774d1', 'Highly recommend it!', 10, '2025-05-05'),
('1b23667e-3f65-496c-92df-00de0ad36931', 'bd71a543-ce71-4fd1-bd45-da411f9b91fc', 'Average, nothing special.', 7, '2025-04-19'),
('96116344-b23e-49cf-868a-c67656ec17d7', '108c322d-d851-40e2-9f1c-e003fbceea3a', 'Average, nothing special.', 4, '2025-03-10'),
('96116344-b23e-49cf-868a-c67656ec17d7', 'a7d3632d-58fb-4627-a5d3-d8699c80256c', 'Exceeded my expectations.', 10, '2025-02-20'),
('96116344-b23e-49cf-868a-c67656ec17d7', 'd24e3e4d-5e1f-495b-aaa8-f9347bbbd6d0', 'Good, but could be better.', 4, '2025-02-18'),
('bd25f99d-e779-4eca-af91-cd7507d987ed', '28633bd5-b5ea-47c4-b8d7-b096cb2be139', 'Very disappointing.', 1, '2025-04-22'),
('bd25f99d-e779-4eca-af91-cd7507d987ed', '9a8c7da7-c5f5-4676-8743-37aafdee8c02', 'Fantastic product!', 8, '2025-03-16');

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
('1a85843f-cfd1-44b6-972a-7d2f1c5e8529', 'janedoe', '$2a$10$yaEiDbQUZfC/8jadySvIwu.j7Dp5u9fHOdp7tfMqW1rIl24cB9dy2', 'jane_doe@gmail.com', 'Jane', 'Doe', 'Main Street 5'),
('1b23667e-3f65-496c-92df-00de0ad36931', 'bob88', '$2a$10$5FVzzxKcOwylCqXOf.E2Zun/8WN8OXnqBasIiMCcpsZd0rNIkqN8y', 'bob88@gmail.com', 'Bob', 'Johnson', 'Mountain View 10'),
('96116344-b23e-49cf-868a-c67656ec17d7', 'helloWorld', '$2a$10$YkNok2pbmVg4ivuEv5G42Og455.kqQ84iMaZhdTu1tOaeZ4LwgA/6', 'john_doe@gmail.com', 'John', 'Doe', 'Yolostreet 15'),
('bd25f99d-e779-4eca-af91-cd7507d987ed', 'alice123', '$2a$10$78vQiqW1fw5S/eDnFxKvz.Rc9McOs2/neAuhsB5xC2KSmEHTElJq2', 'alice@gmail.com', 'Alice', 'Smith', 'River Road 42');

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
