begin;

drop index if exists comments_article_idx;
drop index if exists comments_author_idx;

drop index if exists articles_slug_idx;
drop index if exists articles_author_idx;

drop index if exists users_email_idx;

commit;