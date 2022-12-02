SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Table `bbdate`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `bbdate`.`users` (
  `id` BIGINT(10) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` VARCHAR(512) NOT NULL COMMENT 'ユーザー名',
  `email_address` VARCHAR(512) NULL DEFAULT NULL COMMENT 'Eメールアドレス',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '生成日時',
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新日時',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `email_address_UNIQUE` (`email_address` ASC))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_unicode_ci;

-- -----------------------------------------------------
-- Table `bbdate`.`auth_users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `bbdate`.`auth_users` (
  `id` BIGINT(10) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` VARCHAR(512) NOT NULL COMMENT 'ユーザー名',
  `email_address` VARCHAR(512) NOT NULL COMMENT 'メールアドレス',
  `password` VARCHAR(512) NOT NULL COMMENT 'パスワード',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '生成日時',
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新日時',
  PRIMARY KEY (`id`),
  UNIQUE INDEX `email_address_UNIQUE` (`email_address` ASC))
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4
COLLATE = utf8mb4_unicode_ci;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;