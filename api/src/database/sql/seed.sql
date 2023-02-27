INSERT INTO users (fullname, nick, email, password) VALUES
("staging1", "st1", "staging1@gmail.com", "$2a$10$lBCaPQ8M.Fr3/5HmI068e.3CuZbzQnzIdofL0REvCTJ1/2gazwwv2"),
("staging2", "st2", "staging2@gmail.com", "$2a$10$lBCaPQ8M.Fr3/5HmI068e.3CuZbzQnzIdofL0REvCTJ1/2gazwwv2"),
("staging3", "st3", "staging3@gmail.com", "$2a$10$lBCaPQ8M.Fr3/5HmI068e.3CuZbzQnzIdofL0REvCTJ1/2gazwwv2");

INSERT INTO followers (user_id, follower_id) VALUES
(1, 2),
(1, 2),
(3, 1),
(2, 3);
