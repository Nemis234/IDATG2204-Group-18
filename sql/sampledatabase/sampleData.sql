-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: 22. Mai, 2025 14:27 PM
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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dataark for tabell `administrators`
--

INSERT INTO `administrators` (`UserID`, `RoleName`) VALUES
('ef0bafcb-a085-497e-9416-5b57e0bafdba', 'admin');

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `brand`
--

CREATE TABLE `brand` (
  `BrandName` varchar(50) NOT NULL,
  `BrandDesc` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dataark for tabell `cartitem`
--

INSERT INTO `cartitem` (`UserID`, `ProductID`, `Quantity`) VALUES
('03f1d279-32dc-4e63-91c2-3469c00a4f74', '8e303112-32ca-4ccb-9c79-805c19ca258f', 2),
('03f1d279-32dc-4e63-91c2-3469c00a4f74', '9eb2fe46-e138-44d4-82c3-2fa49fa6a22b', 3),
('a25bf5c9-9cc5-489e-9915-d41fbdf1af18', '4bd208fc-fa53-4ad2-a189-ed441764b480', 2),
('a25bf5c9-9cc5-489e-9915-d41fbdf1af18', 'eaa93045-066a-466b-b963-c388d98f513d', 2),
('a25bf5c9-9cc5-489e-9915-d41fbdf1af18', 'f99440a9-5f6b-4c64-8fd2-9aeb3d4d31e8', 1),
('c8c5c503-8a24-4f42-850f-b6b2ebc1d2f3', '20aedb8c-505c-4572-aa13-49b43b13de3f', 1),
('c8c5c503-8a24-4f42-850f-b6b2ebc1d2f3', 'a5bedf7c-dd9b-4105-9e09-00002e4dcb7d', 3),
('ef0bafcb-a085-497e-9416-5b57e0bafdba', '54236647-467c-4d51-b083-d3c4a8a64c30', 3),
('ef0bafcb-a085-497e-9416-5b57e0bafdba', '68edfc0b-a26b-4eab-bc01-009180f9b848', 2),
('ef0bafcb-a085-497e-9416-5b57e0bafdba', 'ddfcfda4-7466-4f1e-99f5-8d53351127dd', 2);

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `category`
--

CREATE TABLE `category` (
  `CategoryName` varchar(50) NOT NULL,
  `CategoryDesc` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dataark for tabell `members`
--

INSERT INTO `members` (`UserID`, `MembershipLevel`, `MembershipStart`) VALUES
('03f1d279-32dc-4e63-91c2-3469c00a4f74', 'Platinum', '2025-05-22'),
('a25bf5c9-9cc5-489e-9915-d41fbdf1af18', 'Silver', '2025-05-22'),
('ef0bafcb-a085-497e-9416-5b57e0bafdba', 'Gold', '2025-05-22');

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `orderitem`
--

CREATE TABLE `orderitem` (
  `OrderID` varchar(50) NOT NULL,
  `ProductID` varchar(50) NOT NULL,
  `Quantity` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dataark for tabell `orderitem`
--

INSERT INTO `orderitem` (`OrderID`, `ProductID`, `Quantity`) VALUES
('4c65cb71-667c-49d5-9367-b29ca3fe9662', '9cd90103-073f-4b3b-9ff6-12d95e0c32ce', 3),
('4c65cb71-667c-49d5-9367-b29ca3fe9662', 'eaa93045-066a-466b-b963-c388d98f513d', 1),
('7e7632a2-0619-42df-85f6-0279b3c5f5db', '4bd208fc-fa53-4ad2-a189-ed441764b480', 3),
('7e7632a2-0619-42df-85f6-0279b3c5f5db', '83452b25-2fa5-4122-b199-dca28e4cb9fa', 1),
('817e560f-b8f6-4b00-8ed7-510ccc7ba614', '4413dcca-7177-4965-8c74-8383fc544889', 3),
('817e560f-b8f6-4b00-8ed7-510ccc7ba614', 'ddfcfda4-7466-4f1e-99f5-8d53351127dd', 1),
('a46b2f1a-c565-418e-a783-98b8175c33b8', 'a5bedf7c-dd9b-4105-9e09-00002e4dcb7d', 1),
('a46b2f1a-c565-418e-a783-98b8175c33b8', 'b9be6ce4-1396-425f-96fd-1fce01008bfe', 2),
('ae162dcf-69eb-46f2-bdcb-679a18c9811e', '45f31c1b-a451-46e5-a669-0c6d63d6ee4c', 2),
('ae162dcf-69eb-46f2-bdcb-679a18c9811e', 'f99440a9-5f6b-4c64-8fd2-9aeb3d4d31e8', 3),
('b365e7b4-1490-474d-aa55-70b45bebe68d', '68edfc0b-a26b-4eab-bc01-009180f9b848', 2),
('b365e7b4-1490-474d-aa55-70b45bebe68d', 'c40392d1-c1fa-454c-8439-8d429fa6f0a0', 2),
('cd5e1a9d-4590-4019-b3c7-f6d66a242919', '11a26214-c0e6-4b5a-b833-aaeff626ee64', 3),
('cd5e1a9d-4590-4019-b3c7-f6d66a242919', '8e303112-32ca-4ccb-9c79-805c19ca258f', 3),
('d746a681-3f5d-4d50-a0a2-c12d1696b808', '54236647-467c-4d51-b083-d3c4a8a64c30', 3),
('d746a681-3f5d-4d50-a0a2-c12d1696b808', 'd0aebd57-2f98-4e6d-b975-4f198ab7d0d2', 3),
('e908f29b-f2d5-4cdd-95f4-e479542f64c0', '00c405fb-10d1-4075-9417-c171409d9daf', 2),
('e908f29b-f2d5-4cdd-95f4-e479542f64c0', '20aedb8c-505c-4572-aa13-49b43b13de3f', 3),
('f94aa0ee-5aa0-4200-9263-f6b04cd277b7', '9eb2fe46-e138-44d4-82c3-2fa49fa6a22b', 3),
('f94aa0ee-5aa0-4200-9263-f6b04cd277b7', 'f9019fe4-a316-4bef-b466-948416b44e89', 1);

-- --------------------------------------------------------

--
-- Tabellstruktur for tabell `orderstatus`
--

CREATE TABLE `orderstatus` (
  `StatusName` varchar(50) NOT NULL,
  `StatusDesc` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

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
  `OrderStatus` varchar(50) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dataark for tabell `ordertable`
--

INSERT INTO `ordertable` (`OrderID`, `UserID`, `OrderDate`, `OrderStatus`) VALUES
('4c65cb71-667c-49d5-9367-b29ca3fe9662', 'a25bf5c9-9cc5-489e-9915-d41fbdf1af18', '2025-04-08', 'Returned'),
('7e7632a2-0619-42df-85f6-0279b3c5f5db', 'a25bf5c9-9cc5-489e-9915-d41fbdf1af18', '2025-03-31', 'Processing'),
('817e560f-b8f6-4b00-8ed7-510ccc7ba614', 'ef0bafcb-a085-497e-9416-5b57e0bafdba', '2025-03-24', 'Pending'),
('a46b2f1a-c565-418e-a783-98b8175c33b8', 'c8c5c503-8a24-4f42-850f-b6b2ebc1d2f3', '2025-03-13', 'Failed'),
('ae162dcf-69eb-46f2-bdcb-679a18c9811e', 'a25bf5c9-9cc5-489e-9915-d41fbdf1af18', '2025-04-24', 'Processing'),
('b365e7b4-1490-474d-aa55-70b45bebe68d', 'ef0bafcb-a085-497e-9416-5b57e0bafdba', '2025-05-17', 'Cancelled'),
('cd5e1a9d-4590-4019-b3c7-f6d66a242919', '03f1d279-32dc-4e63-91c2-3469c00a4f74', '2025-02-12', 'Shipped'),
('d746a681-3f5d-4d50-a0a2-c12d1696b808', 'ef0bafcb-a085-497e-9416-5b57e0bafdba', '2025-02-20', 'Pending'),
('e908f29b-f2d5-4cdd-95f4-e479542f64c0', 'c8c5c503-8a24-4f42-850f-b6b2ebc1d2f3', '2025-04-16', 'Delivered'),
('f94aa0ee-5aa0-4200-9263-f6b04cd277b7', '03f1d279-32dc-4e63-91c2-3469c00a4f74', '2025-03-27', 'Refunded');

-- --------------------------------------------------------

--
-- Erstatningsstruktur for visning `ordertotal`
-- (See below for the actual view)
--
CREATE TABLE `ordertotal` (
`OrderID` varchar(50)
,`UserID` varchar(50)
,`OrderDate` date
,`OrderStatus` varchar(50)
,`OrderTotal` decimal(42,2)
);

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dataark for tabell `payment`
--

INSERT INTO `payment` (`PaymentID`, `OrderID`, `PaymentMethod`, `Amount`, `PaymentDate`, `PaymentStatus`) VALUES
('2b417ed9-a43c-41cd-b675-25d989d0ce5f', '817e560f-b8f6-4b00-8ed7-510ccc7ba614', 'bank_transfer', 2099.96, '2025-04-06', 'pending'),
('57a2385e-97d3-447e-ae54-6174ab3cb2cd', 'ae162dcf-69eb-46f2-bdcb-679a18c9811e', 'card', 53999.95, '2025-05-22', 'pending'),
('5e26ffea-13d6-436b-8698-ace15dd09c08', 'e908f29b-f2d5-4cdd-95f4-e479542f64c0', 'card', 18999.95, '2025-03-30', 'successful'),
('71779b78-00b2-49c2-ab12-6919f7ebe918', 'd746a681-3f5d-4d50-a0a2-c12d1696b808', 'card', 29999.94, '2025-05-06', 'pending'),
('77d1ad09-a2a8-4217-ba13-ed98720f0fba', 'a46b2f1a-c565-418e-a783-98b8175c33b8', 'vipps', 10499.97, '2025-04-09', 'failed'),
('82b8f9d9-45f3-4df0-9df8-03d1f2d8e5c4', 'b365e7b4-1490-474d-aa55-70b45bebe68d', 'vipps', 5999.96, '2025-03-22', 'failed'),
('958692a1-cc31-48f5-930e-ad02605ef367', 'cd5e1a9d-4590-4019-b3c7-f6d66a242919', 'bank_transfer', 32999.94, '2025-05-15', 'successful'),
('bc1c7206-88e6-4f2e-807f-cd5477548322', '7e7632a2-0619-42df-85f6-0279b3c5f5db', 'vipps', 39799.96, '2025-04-18', 'pending'),
('bdb7e88c-a62f-4c30-94cc-52439a13c2e0', 'f94aa0ee-5aa0-4200-9263-f6b04cd277b7', 'card', 12299.96, '2025-04-25', 'failed'),
('f270f2fe-807b-403e-9c21-5381dceb394a', '4c65cb71-667c-49d5-9367-b29ca3fe9662', 'bank_transfer', 2999.96, '2025-05-02', 'failed');

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dataark for tabell `product`
--

INSERT INTO `product` (`ProductID`, `ProductName`, `ProductDesc`, `ProductImgUrl`, `Price`, `StockQuantity`, `Brand`, `Category`) VALUES
('00c405fb-10d1-4075-9417-c171409d9daf', 'E-Gadgets 4K Projector', 'Cinema-quality projector for home.', 'https://example.com/img14.jpg', 4999.99, 45, 'E-Gadgets', 'TV & Home Theater'),
('11a26214-c0e6-4b5a-b833-aaeff626ee64', 'TechBrave Drone Pro', 'Professional drone with 4K camera.', 'https://example.com/img13.jpg', 8999.99, 25, 'TechBrave', 'Cameras & Photography'),
('20aedb8c-505c-4572-aa13-49b43b13de3f', 'TechBrave Smartwatch 2', 'Fitness tracking smartwatch.', 'https://example.com/img4.jpg', 2999.99, 150, 'TechBrave', 'Wearable Technology'),
('4413dcca-7177-4965-8c74-8383fc544889', 'TechBrave Power Bank 20000mAh', 'Portable charger for mobile devices.', 'https://example.com/img19.jpg', 499.99, 250, 'TechBrave', 'Mobile Phones & Tablets'),
('45f31c1b-a451-46e5-a669-0c6d63d6ee4c', 'E-Gadgets Robot Vacuum', 'Smart robot vacuum cleaner.', 'https://example.com/img20.jpg', 2999.99, 30, 'E-Gadgets', 'Smart Home'),
('4bd208fc-fa53-4ad2-a189-ed441764b480', 'E-Gadgets Ultra HD TV 55\"', 'Ultra HD smart TV with vibrant colors.', 'https://example.com/img2.jpg', 12999.99, 30, 'E-Gadgets', 'TV & Home Theater'),
('54236647-467c-4d51-b083-d3c4a8a64c30', 'TechBrave X100 Smartphone', 'High-end smartphone with AI features.', 'https://example.com/img1.jpg', 8999.99, 100, NULL, 'Mobile Phones & Tablets'),
('68edfc0b-a26b-4eab-bc01-009180f9b848', 'E-Gadgets Home Assistant', 'Voice-controlled smart home hub.', 'https://example.com/img5.jpg', 1499.99, 50, 'E-Gadgets', 'Smart Home'),
('83452b25-2fa5-4122-b199-dca28e4cb9fa', 'DigiTech Gaming Mouse', 'Precision gaming mouse.', 'https://example.com/img12.jpg', 799.99, 120, 'DigiTech', 'Gaming'),
('8e303112-32ca-4ccb-9c79-805c19ca258f', 'DigiTech Gaming Headset Pro', 'Immersive gaming headset with surround sound.', 'https://example.com/img3.jpg', 1999.99, 80, 'DigiTech', 'Audio & Headphones'),
('9cd90103-073f-4b3b-9ff6-12d95e0c32ce', 'TechBrave Fitness Band', 'Basic fitness tracker with heart monitor.', 'https://example.com/img16.jpg', 699.99, 110, 'TechBrave', 'Wearable Technology'),
('9eb2fe46-e138-44d4-82c3-2fa49fa6a22b', 'TechBrave Tablet 10\"', 'Lightweight tablet for work and play.', 'https://example.com/img7.jpg', 3999.99, 90, 'TechBrave', 'Mobile Phones & Tablets'),
('a5bedf7c-dd9b-4105-9e09-00002e4dcb7d', 'E-Gadgets Noise Cancelling Headphones', 'Noise-free music experience.', 'https://example.com/img8.jpg', 2499.99, 70, 'E-Gadgets', 'Audio & Headphones'),
('b9be6ce4-1396-425f-96fd-1fce01008bfe', 'DigiTech VR Headset', 'Immersive VR gaming experience.', 'https://example.com/img18.jpg', 3999.99, 60, 'DigiTech', 'Gaming'),
('c40392d1-c1fa-454c-8439-8d429fa6f0a0', 'DigiTech Portable SSD 1TB', 'High-speed portable storage.', 'https://example.com/img15.jpg', 1499.99, 70, 'DigiTech', 'Computers & Accessories'),
('d0aebd57-2f98-4e6d-b975-4f198ab7d0d2', 'E-Gadgets Security Camera', 'Smart WiFi home security camera.', 'https://example.com/img11.jpg', 999.99, 100, 'E-Gadgets', 'Smart Home'),
('ddfcfda4-7466-4f1e-99f5-8d53351127dd', 'DigiTech Wireless Charger', 'Fast wireless charging pad.', 'https://example.com/img9.jpg', 599.99, 300, 'DigiTech', 'Mobile Phones & Tablets'),
('eaa93045-066a-466b-b963-c388d98f513d', 'DigiTech Bluetooth Speaker', 'Portable speaker with rich bass.', 'https://example.com/img6.jpg', 899.99, 200, 'DigiTech', 'Audio & Headphones'),
('f9019fe4-a316-4bef-b466-948416b44e89', 'E-Gadgets Smart Light Bulb', 'WiFi-enabled color-changing bulb.', 'https://example.com/img17.jpg', 299.99, 500, 'E-Gadgets', 'Smart Home'),
('f99440a9-5f6b-4c64-8fd2-9aeb3d4d31e8', 'TechBrave Laptop Pro', 'Powerful laptop for professionals.', 'https://example.com/img10.jpg', 15999.99, 40, 'TechBrave', 'Computers & Accessories');

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dataark for tabell `review`
--

INSERT INTO `review` (`UserID`, `ProductID`, `Comment`, `Rating`, `PostDate`) VALUES
('03f1d279-32dc-4e63-91c2-3469c00a4f74', '8e303112-32ca-4ccb-9c79-805c19ca258f', 'It\'s okay, does the job.', 5, '2025-04-25'),
('03f1d279-32dc-4e63-91c2-3469c00a4f74', '9eb2fe46-e138-44d4-82c3-2fa49fa6a22b', 'Good, but could be better.', 6, '2025-04-26'),
('a25bf5c9-9cc5-489e-9915-d41fbdf1af18', '4bd208fc-fa53-4ad2-a189-ed441764b480', 'Broke after a few days.', 3, '2025-03-19'),
('a25bf5c9-9cc5-489e-9915-d41fbdf1af18', 'eaa93045-066a-466b-b963-c388d98f513d', 'Amazing quality!', 10, '2025-05-07'),
('a25bf5c9-9cc5-489e-9915-d41fbdf1af18', 'f99440a9-5f6b-4c64-8fd2-9aeb3d4d31e8', 'Amazing quality!', 10, '2025-04-17'),
('c8c5c503-8a24-4f42-850f-b6b2ebc1d2f3', '20aedb8c-505c-4572-aa13-49b43b13de3f', 'Good, but could be better.', 6, '2025-02-24'),
('c8c5c503-8a24-4f42-850f-b6b2ebc1d2f3', 'a5bedf7c-dd9b-4105-9e09-00002e4dcb7d', 'Broke after a few days.', 1, '2025-03-13'),
('ef0bafcb-a085-497e-9416-5b57e0bafdba', '54236647-467c-4d51-b083-d3c4a8a64c30', 'Good, but could be better.', 6, '2025-03-13'),
('ef0bafcb-a085-497e-9416-5b57e0bafdba', '68edfc0b-a26b-4eab-bc01-009180f9b848', 'It\'s okay, does the job.', 5, '2025-04-19'),
('ef0bafcb-a085-497e-9416-5b57e0bafdba', 'ddfcfda4-7466-4f1e-99f5-8d53351127dd', 'Terrible quality!', 3, '2025-02-25');

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
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dataark for tabell `users`
--

INSERT INTO `users` (`UserID`, `Username`, `Password`, `Email`, `FirstName`, `LastName`, `Address`) VALUES
('03f1d279-32dc-4e63-91c2-3469c00a4f74', 'alice123', '$2a$10$M9Dm49icKRaE6b5eTYap5uwvRNz8HNjl9Dk9IH53THfLFS2n2bbq2', 'alice@gmail.com', 'Alice', 'Smith', 'River Road 42'),
('a25bf5c9-9cc5-489e-9915-d41fbdf1af18', 'janedoe', '$2a$10$UhZy9gj6uXO1XvPNDWtRZu.fvKfbOeSjSfFfGpAJWZHI1ORF7CKyq', 'jane_doe@gmail.com', 'Jane', 'Doe', 'Main Street 5'),
('c8c5c503-8a24-4f42-850f-b6b2ebc1d2f3', 'bob88', '$2a$10$WRVGYMPGJNjJPdv.klAYxuC4JNBFstPvr7bYOy3jo/LYj6LGC11eS', 'bob88@gmail.com', 'Bob', 'Johnson', 'Mountain View 10'),
('ef0bafcb-a085-497e-9416-5b57e0bafdba', 'helloWorld', '$2a$10$Kq8AlPb.3n.OGkAbKkMPA.MyTaX3QbQDjULXM467rob44tzzWCrfy', 'john_doe@gmail.com', 'John', 'Doe', 'Yolostreet 15');

-- --------------------------------------------------------

--
-- Visningsstruktur `ordertotal`
--
DROP TABLE IF EXISTS `ordertotal`;

CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `ordertotal`  AS SELECT `o`.`OrderID` AS `OrderID`, `o`.`UserID` AS `UserID`, `o`.`OrderDate` AS `OrderDate`, `o`.`OrderStatus` AS `OrderStatus`, sum(`oi`.`Quantity` * `p`.`Price`) AS `OrderTotal` FROM ((`ordertable` `o` join `orderitem` `oi` on(`o`.`OrderID` = `oi`.`OrderID`)) join `product` `p` on(`oi`.`ProductID` = `p`.`ProductID`)) GROUP BY `o`.`OrderID`, `o`.`UserID`, `o`.`OrderDate`, `o`.`OrderStatus` ;

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
  ADD KEY `OrderStatus` (`OrderStatus`),
  ADD KEY `ind_ordertable_userid` (`UserID`);

--
-- Indexes for table `payment`
--
ALTER TABLE `payment`
  ADD PRIMARY KEY (`PaymentID`),
  ADD KEY `idx_payment_orderid` (`OrderID`);

--
-- Indexes for table `product`
--
ALTER TABLE `product`
  ADD PRIMARY KEY (`ProductID`),
  ADD KEY `idx_product_brand` (`Brand`),
  ADD KEY `idx_product_category` (`Category`);

--
-- Indexes for table `review`
--
ALTER TABLE `review`
  ADD PRIMARY KEY (`UserID`,`ProductID`),
  ADD KEY `idx_review_productid` (`ProductID`);

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
