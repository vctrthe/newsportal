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

INSERT INTO `users` (`id`, `name`, `email`, `password`, `created_at`, `updated_at`) VALUES
(1, 'John Doe', 'john@example.com', 'hashedpassword1', NOW(), NOW()),
(2, 'Jane Smith', 'jane@example.com', 'hashedpassword2', NOW(), NOW()),
(3, 'Alice Johnson', 'alice@example.com', 'hashedpassword3', NOW(), NOW()),
(4, 'Bob Brown', 'bob@example.com', 'hashedpassword4', NOW(), NOW()),
(5, 'Charlie White', 'charlie@example.com', 'hashedpassword5', NOW(), NOW()),
(6, 'David Black', 'david@example.com', 'hashedpassword6', NOW(), NOW()),
(7, 'Eve Green', 'eve@example.com', 'hashedpassword7', NOW(), NOW()),
(8, 'Frank Blue', 'frank@example.com', 'hashedpassword8', NOW(), NOW()),
(9, 'Grace Red', 'grace@example.com', 'hashedpassword9', NOW(), NOW()),
(10, 'Hank Yellow', 'hank@example.com', 'hashedpassword10', NOW(), NOW());

CREATE TABLE `categories` (
  `id` int PRIMARY KEY,
  `created_by_id` int,
  `title` varchar(200),
  `slug` varchar(200),
  `created_at` timestamp,
  `updated_at` timestamp,
  FOREIGN KEY (`created_by_id`) REFERENCES `users` (`id`)
);

INSERT INTO `categories` (`id`, `created_by_id`, `title`, `slug`, `created_at`, `updated_at`) VALUES
(1, 1, 'Technology', 'technology', NOW(), NOW()),
(2, 2, 'Health', 'health', NOW(), NOW()),
(3, 3, 'Entertainment', 'entertainment', NOW(), NOW()),
(4, 4, 'Sports', 'sports', NOW(), NOW()),
(5, 5, 'Science', 'science', NOW(), NOW()),
(6, 6, 'Business', 'business', NOW(), NOW()),
(7, 7, 'Lifestyle', 'lifestyle', NOW(), NOW()),
(8, 8, 'Education', 'education', NOW(), NOW()),
(9, 9, 'Politics', 'politics', NOW(), NOW()),
(10, 10, 'Travel', 'travel', NOW(), NOW());

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

INSERT INTO `contents` (`id`, `category_id`, `created_by_id`, `title`, `excerpt`, `description`, `image`, `status`, `tags`, `created_at`, `updated_at`) VALUES
(1, 1, 1, 'The Future of AI', 'AI is evolving rapidly...', 'Full content about AI advancements...', 'ai_image.jpg', 'published', 'AI, Tech, Future', NOW(), NOW()),
(2, 2, 2, '10 Tips for a Healthy Life', 'Living a healthy life...', 'Full content with health tips...', 'health_tips.jpg', 'published', 'Health, Wellness', NOW(), NOW()),
(3, 3, 3, 'Upcoming Movie Releases', 'Exciting movies are coming...', 'Full content about movies...', 'movies.jpg', 'draft', 'Movies, Entertainment', NOW(), NOW()),
(4, 4, 4, 'The Rise of Esports', 'Esports is gaining popularity...', 'Full content about esports...', 'esports.jpg', 'published', 'Esports, Gaming', NOW(), NOW()),
(5, 5, 5, 'New Space Discoveries', 'NASA found new planets...', 'Full content about space...', 'space.jpg', 'published', 'Space, Science', NOW(), NOW()),
(6, 6, 6, 'Stock Market Trends', 'Analyzing recent market trends...', 'Full content on stock trends...', 'stocks.jpg', 'draft', 'Finance, Business', NOW(), NOW()),
(7, 7, 7, 'Minimalist Living Tips', 'How to live with less...', 'Full content about minimalism...', 'minimalism.jpg', 'published', 'Lifestyle, Minimalism', NOW(), NOW()),
(8, 8, 8, 'Best Online Learning Platforms', 'Education is going digital...', 'Full content on e-learning...', 'elearning.jpg', 'published', 'Education, Online', NOW(), NOW()),
(9, 9, 9, 'Election 2025: What to Expect', 'Politics is heating up...', 'Full political analysis...', 'election.jpg', 'draft', 'Politics, Government', NOW(), NOW()),
(10, 10, 10, 'Top 10 Travel Destinations', 'Best places to visit...', 'Full travel guide...', 'travel.jpg', 'published', 'Travel, Adventure', NOW(), NOW());