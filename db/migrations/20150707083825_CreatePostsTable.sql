-- The whole purpose of this table is a demonstration

-- +goose Up
CREATE TABLE post (
    id INT NOT NULL,
    title text,
    body text,
    PRIMARY KEY(id)
);
 
-- +goose Down
DROP TABLE post;
- See more at: https://labs.chie.do/using-goose-for-golang-with-mysql/#sthash.IZn7r0Og.dpuf
