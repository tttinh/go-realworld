begin;

create index if not exists users_email_idx on users (email);

create index if not exists articles_author_idx on articles (author_id);
create index if not exists articles_slug_idx on articles (slug);

create index if not exists comments_author_idx on comments (author_id);
create index if not exists comments_article_idx on comments (article_id);

commit;