create table if not exists users (
    id bigserial primary key,
    username text not null unique,
    email text not null unique,
    pwd text not null,
    bio text,
    img text,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp
);

create table if not exists articles (
    id bigserial primary key,
    author_id bigint not null,
    slug text not null unique,
    title text not null,
    description text not null,
    body text not null,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp,

    foreign key (author_id) references users(id)
);

create table if not exists comments (
    id bigserial primary key,
    body text not null,
    author_id bigint not null,
    article_id bigint not null,

    foreign key (author_id) references users(id),
    foreign key (article_id) references articles(id)
);