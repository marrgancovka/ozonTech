CREATE TABLE IF NOT EXISTS "user" (
                                      id SERIAL PRIMARY KEY,
                                      name TEXT CONSTRAINT name_length CHECK (char_length(name) <= 30) NOT NULL UNIQUE,
                                      password_hash TEXT CONSTRAINT passwordHash_length CHECK (char_length(password_hash) <= 64) NOT NULL
);


CREATE TABLE IF NOT EXISTS post (
                                    id SERIAL PRIMARY KEY,
                                    user_id INTEGER NOT NULL,
                                    title TEXT CONSTRAINT title_length CHECK (char_length(title) <= 500) NOT NULL,
                                    content TEXT CONSTRAINT content_length CHECK (char_length(content) <= 5000) NOT NULL,
                                    comments_allowed BOOLEAN DEFAULT true NOT NULL,
                                    FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS comment (
                                       id SERIAL PRIMARY KEY,
                                       post_id INTEGER NOT NULL,
                                       user_id INTEGER NOT NULL,
                                       text TEXT NOT NULL,
                                       parent_comment_id INTEGER DEFAULT 0,
                                       child_comments INTEGER[] DEFAULT ARRAY[]::INTEGER[],
                                       FOREIGN KEY (post_id) REFERENCES post(id) ON DELETE CASCADE,
                                       FOREIGN KEY (user_id) REFERENCES "user"(id) ON DELETE CASCADE
);