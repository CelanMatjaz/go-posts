-- gen_random_uuid ()

ALTER TABLE users
ALTER id SET DEFAULT gen_random_uuid();

ALTER TABLE posts
ALTER id SET DEFAULT gen_random_uuid();
