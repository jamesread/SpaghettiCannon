-- +migrate Up
CREATE TABLE `cvars` (
  `cvar_key` varchar(255) NOT NULL,
  `cvar_value_int` tinyint(4) DEFAULT NULL,
  `cvar_value_string` varchar(255) DEFAULT NULL,
  `cvar_main_type` varchar(255) NOT NULL,
  `cvar_title` varchar(255) NOT NULL DEFAULT '',
  `cvar_description` varchar(512) NOT NULL DEFAULT '',
  `cvar_category` varchar(255) NOT NULL DEFAULT '',
  `cvar_ordinal` int(11) NOT NULL DEFAULT 0,
  PRIMARY KEY (`cvar_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- +migrate Down
DROP TABLE IF EXISTS `cvars`;
