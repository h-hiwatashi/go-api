# 記事データ 2 つ
insert into articles (title, contents, username, nice, created_at) values
('firstPost', 'This is my first blog', 'user name', 2, now());
insert into articles (title, contents, username, nice) values
('2nd', 'Second blog post', 'user name', 4);
# コメントデータ 2 つ
insert into comments (article_id, message, created_at) values
(1, '1st comment yeah', now());
insert into comments (article_id, message) values
(1, 'welcome');