CREATE TABLE `users` (
  `id` integer PRIMARY KEY,
  `user_id` varchar(255),
  `password` varchar(255),
  `nickname` varchar(255),
  `credits` integer
);

CREATE TABLE `volunteerPosts` (
  `id` integer PRIMARY KEY,
  `category` tinyint,
  `content` varchar(255),
  `createdAt` datetime,
  `userPK` integer,
  `DISCD` tinyint
);

CREATE TABLE `volunteerPostComments` (
  `id` integer PRIMARY KEY,
  `content` varchar(255),
  `userPK` integer,
  `volunteerPostPK` integer,
  `commentPK` integer
);

CREATE TABLE `chatRooms` (
  `id` integer PRIMARY KEY,
  `chatRoomName` varchar(255),
  `chatType` integer,
  `ownerPK` integer,
  `DISCD` tinyint
);

CREATE TABLE `chats` (
  `id` integer PRIMARY KEY,
  `message` varchar(255),
  `isRead` tinyint,
  `chatRoomPK` integer,
  `senderPK` integer,
  `createdAt` datetime,
  `DISCD` tinyint
);

CREATE TABLE `products` (
  `id` integer PRIMARY KEY,
  `category` integer,
  `productName` varchar(255),
  `unitPrice` integer,
  `isSoldOut` tinyint,
  `createdAt` datetime,
  `DISCD` tinyint,
  `maxDiscountRate` integer,
  `couponApplyable` tinyint,
  `creditsApplyable` tinyint
);

CREATE TABLE `productReview` (
  `id` integer PRIMARY KEY,
  `content` varchar(255),
  `productPK` integer,
  `userPK` integer,
  `rate` integer
);

CREATE TABLE `productReviewComment` (
  `id` integer PRIMARY KEY,
  `content` varchar(255),
  `userPK` integer,
  `reviewPK` integer
);

CREATE TABLE `product_order` (
  `id` integer PRIMARY KEY,
  `orderCount` integer,
  `productPK` integer,
  `orderPK` integer,
  `DISCD` tinyint
);

CREATE TABLE `order` (
  `id` integer PRIMARY KEY,
  `userPK` integer,
  `discountCouponPK` integer,
  `usedCredits` integer,
  `orderStatus` tinyint,
  `createdAt` datetime
);

CREATE TABLE `payment` (
  `id` integer PRIMARY KEY,
  `orderPK` integer,
  `paymentMethod` integer,
  `payAmount` integer
);

CREATE TABLE `discountCoupon` (
  `id` integer PRIMARY KEY,
  `discountCouponName` varchar(255),
  `discountRate` integer,
  `applyableProductCategory` int,
  `createdAt` datetime,
  `expiresAt` datetime
);

ALTER TABLE `volunteerPosts` ADD FOREIGN KEY (`userPK`) REFERENCES `users` (`id`);

ALTER TABLE `chatRooms` ADD FOREIGN KEY (`ownerPK`) REFERENCES `users` (`id`);

ALTER TABLE `chats` ADD FOREIGN KEY (`senderPK`) REFERENCES `users` (`id`);

ALTER TABLE `chats` ADD FOREIGN KEY (`chatRoomPK`) REFERENCES `chatRooms` (`id`);

ALTER TABLE `volunteerPostComments` ADD FOREIGN KEY (`userPK`) REFERENCES `users` (`id`);

ALTER TABLE `volunteerPostComments` ADD FOREIGN KEY (`volunteerPostPK`) REFERENCES `volunteerPosts` (`id`);

ALTER TABLE `volunteerPostComments` ADD FOREIGN KEY (`commentPK`) REFERENCES `volunteerPostComments` (`id`);

ALTER TABLE `product_order` ADD FOREIGN KEY (`productPK`) REFERENCES `products` (`id`);

ALTER TABLE `product_order` ADD FOREIGN KEY (`orderPK`) REFERENCES `order` (`id`);

ALTER TABLE `order` ADD FOREIGN KEY (`userPK`) REFERENCES `users` (`id`);

ALTER TABLE `productReview` ADD FOREIGN KEY (`productPK`) REFERENCES `products` (`id`);

ALTER TABLE `productReview` ADD FOREIGN KEY (`userPK`) REFERENCES `users` (`id`);

ALTER TABLE `productReviewComment` ADD FOREIGN KEY (`reviewPK`) REFERENCES `productReview` (`id`);

ALTER TABLE `productReviewComment` ADD FOREIGN KEY (`userPK`) REFERENCES `users` (`id`);

ALTER TABLE `order` ADD FOREIGN KEY (`discountCouponPK`) REFERENCES `discountCoupon` (`id`);

ALTER TABLE `payment` ADD FOREIGN KEY (`orderPK`) REFERENCES `order` (`id`);
