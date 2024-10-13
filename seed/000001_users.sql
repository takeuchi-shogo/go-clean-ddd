-- ユーザーテーブルにサンプルデータを挿入
INSERT INTO `users` (`created_at`, `updated_at`)
VALUES (NOW(), NOW());

-- 挿入されたユーザーのIDを取得
SET @user_id = LAST_INSERT_ID();

-- ユーザー詳細テーブルにサンプルデータを挿入
INSERT INTO `user_details` (`user_id`, `first_name`, `last_name`, `created_at`, `updated_at`)
VALUES (@user_id, '太郎', '山田', NOW(), NOW());

-- ユーザーアカウント情報テーブルにサンプルデータを挿入
INSERT INTO `user_accounts` (`user_id`, `email`, `password`, `created_at`, `updated_at`)
VALUES (@user_id, 'sample@example.com', 'hashed_password', NOW(), NOW());
