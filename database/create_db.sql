-- Active: 1739867389985@@192.168.100.41@3306
DROP DATABASE IF EXISTS newsportal;

CREATE DATABASE newsportal;

USE newsportal;

CREATE TABLE `users` (
  `id` int PRIMARY KEY,
  `name` varchar(100),
  `email` varchar(100),
  `password` varchar(100),
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `categories` (
  `id` int PRIMARY KEY,
  `created_by_id` int,
  `title` varchar(200),
  `slug` varchar(200),
  `created_at` timestamp,
  `updated_at` timestamp,
  FOREIGN KEY (`created_by_id`) REFERENCES `users` (`id`)
);

CREATE TABLE `contents` (
  `id` int PRIMARY KEY,
  `category_id` int,
  `created_by_id` int,
  `title` varchar(200),
  `excerpt` varchar(250),
  `description` text,
  `image` text,
  `status` varchar(20),
  `tags` text,
  `created_at` timestamp,
  `updated_at` timestamp,
  FOREIGN KEY (`created_by_id`) REFERENCES `users` (`id`),
  FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`)
);