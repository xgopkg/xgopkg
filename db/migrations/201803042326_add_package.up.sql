CREATE TABLE xgp_package (
  id BIGINT(20) NOT NULL AUTO_INCREMENT,
  pkg_name CHAR(50) NOT NULL,
  pkg_source VARCHAR(100) NOT NULL,
  pkg_desc VARCHAR(255) DEFAULT NULL,
  updated_at DATETIME NOT NULL,
  created_at DATETIME NOT NULL,
  PRIMARY KEY (id)
);