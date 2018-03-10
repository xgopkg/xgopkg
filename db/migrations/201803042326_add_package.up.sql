CREATE TABLE `xgp_package` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `pkg_name` char(50) NOT NULL,
  `pkg_source` varchar(100) NOT NULL,
  `pkg_desc` varchar(255) DEFAULT NULL,
  `updated_at` datetime NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
);