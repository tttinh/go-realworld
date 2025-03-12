begin;

create table if not exists users (
    id bigserial,
    username text not null unique,
    email text not null unique,
    password text not null,
    bio text,
    image text,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp,

    primary key (id)
);

create table if not exists articles (
    id bigserial,
    author_id bigint not null,
    slug text not null unique,
    title text not null,
    description text not null,
    body text not null,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp,

    primary key (id),
    foreign key (author_id) references users(id)
);

create table if not exists comments (
    id bigserial,
    body text not null,
    author_id bigint not null,
    article_id bigint not null,
    created_at timestamptz not null default current_timestamp,
    updated_at timestamptz not null default current_timestamp,

    primary key (id),
    foreign key (author_id) references users(id),
    foreign key (article_id) references articles(id) on delete cascade
);

create table if not exists follows (
    follower_id bigint,
    following_id bigint,

    primary key (follower_id, following_id),
    foreign key (follower_id) references users(id),
    foreign key (following_id) references users(id)
);

create table if not exists favorites (
    user_id bigint,
    article_id bigint,

    primary key (user_id, article_id),
    foreign key (user_id) references users(id),
    foreign key (article_id) references articles(id) on delete cascade
);

create table if not exists tags (
    id bigserial,
    name text not null unique,

    primary key (id)
);

create table if not exists article_tags (
    article_id bigint,
    tag_id bigint,

    primary key (article_id, tag_id),
    foreign key (tag_id) references tags(id),
    foreign key (article_id) references articles(id) on delete cascade
);

commit;