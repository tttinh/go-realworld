begin;

create table if not exists users (
    id bigserial primary key,
    username text not null unique,
    email text not null unique,
    password text not null,
    bio text,
    image text,
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
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp,

    foreign key (author_id) references users(id),
    foreign key (article_id) references articles(id)
);

create table if not exists follows (
    follower_id bigint,
    following_id bigint,

    foreign key (follower_id) references users(id),
    foreign key (following_id) references users(id),

    primary key (follower_id, following_id)
);

create table if not exists favorites (
    user_id bigint,
    article_id bigint,

    foreign key (user_id) references users(id),
    foreign key (article_id) references articles(id),

    primary key (user_id, article_id)
);

create table if not exists tags (
    id bigserial primary key,
    tag text not null,
    article_id bigint,

    foreign key (article_id) references articles(id),

    unique (tag, article_id)
);

commit;