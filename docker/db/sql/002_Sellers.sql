-- ---- drop ----
DROP TABLE IF EXISTS `Sellers`;

-- ---- create ----
CREATE TABLE `Sellers` (
    `id` BIGINT NOT NULL AUTO_INCREMENT ,
    `password` VARCHAR(16) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL ,
    `price` INT NOT NULL ,
    `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ,
    `expired_at` TIMESTAMP NULL DEFAULT NULL ,
    PRIMARY KEY (`id`, `password`),
    KEY created_at_expired_at(`created_at`, `expired_at`)
) ENGINE = InnoDB;