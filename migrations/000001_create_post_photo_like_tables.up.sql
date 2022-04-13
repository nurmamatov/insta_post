CREATE TABLE IF NOT EXISTS post_photo (
  image_id UUID PRIMARY KEY,
  type varchar,
  base_code text,
  created_at timestamp,
  updated_at timestamp,
  deleted_at timestamp
);
CREATE TABLE IF NOT EXISTS post (
  post_id UUID PRIMARY KEY,
  user_id UUID NOT NULL,
  title varchar(25),
  description varchar(150),
  image UUID REFERENCES post_photo(image_id),
  created_at timestamp,
  updated_at timestamp,
  deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS likes (
  like_id UUID,
  post_id UUID REFERENCES post(post_id),
  user_id UUID NOT NULL,
  likee   boolean
);