CREATE TABLE user
(
    id        bigint AUTO_INCREMENT,
    username  varchar(36) default '',
    password  varchar(64) default '',
    UNIQUE name_index (username),
    PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci;
