# ******************************************************************************
# Settings
# ******************************************************************************
SET foreign_key_checks = 1;
SET time_zone = '+00:00';

# ******************************************************************************
# Create new tables
# ******************************************************************************
# CREATE DATABASE IF NOT EXISTS blueprint DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;
# USE blueprint;

# ******************************************************************************
# Create tables
# ******************************************************************************
CREATE TABLE user_status (
    id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    
    status VARCHAR(25) NOT NULL,
    
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    
    PRIMARY KEY (id)
);

CREATE TABLE user (
    id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password CHAR(60) NOT NULL,
    
    status_id INT(10) UNSIGNED NOT NULL DEFAULT 1,
    
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    
    UNIQUE KEY (email),
    CONSTRAINT `f_user_status` FOREIGN KEY (`status_id`) REFERENCES `user_status` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    
    PRIMARY KEY (id)
);

INSERT INTO `user_status` (`id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'active',   CURRENT_TIMESTAMP,  NULL,  NULL),
(2, 'inactive', CURRENT_TIMESTAMP,  NULL,  NULL);

CREATE TABLE note (
    id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    
    firstname TEXT NOT NULL,
    middlename TEXT NOT NULL,
    lastname TEXT NOT NULL,
    manager TEXT NOT NULL,
    personalemail TEXT,
    businessemail TEXT,
    homephone TEXT,
    mobilephone TEXT,
    workphone TEXT,
    addressline1 TEXT NOT NULL,
    housenamenumber TEXT,
    addressline2 TEXT,
    apartmentsuite TEXT,
    town TEXT NOT NULL,
    county TEXT,
    postcode TEXT NOT NULL,
    country TEXT NOT NULL,
    noktitle TEXT NOT NULL,
    nokfirstname TEXT NOT NULL,
    noklastname TEXT NOT NULL,
    nokrelationship TEXT,
    nokaddressline1 TEXT NOT NULL,
    nokhousenamenumber TEXT,
    nokaddressline2 TEXT,
    nokapartmentsuite TEXT,
    noktown TEXT NOT NULL,
    nokcounty TEXT,
    nokpostcode TEXT NOT NULL,
    nokcountry TEXT NOT NULL,
    nokhomephone TEXT,
    nokworkphone TEXT,
    ec1title TEXT NOT NULL,
    ec1firstname TEXT NOT NULL,
    ec1lastname TEXT NOT NULL,
    ec1relationship TEXT,
    ec1addressline1 TEXT NOT NULL,
    ec1housenamenumber TEXT,
    ec1addressline2 TEXT,
    ec1apartmentsuite TEXT,
    ec1town TEXT NOT NULL,
    ec1county TEXT,
    ec1postcode TEXT NOT NULL,
    ec1country TEXT NOT NULL,
    ec1homephone TEXT,
    ec1workphone TEXT,
    ec2title TEXT NOT NULL,
    ec2firstname TEXT NOT NULL,
    ec2lastname TEXT NOT NULL,
    ec2relationship TEXT,
    ec2addressline1 TEXT NOT NULL,
    ec2housenamenumber TEXT,
    ec2addressline2 TEXT,
    ec2apartmentsuite TEXT,
    ec2town TEXT NOT NULL,
    ec2county TEXT,
    ec2postcode TEXT NOT NULL,
    ec2country TEXT NOT NULL,
    ec2homephone TEXT,
    ec2workphone TEXT,
    ec3title TEXT NOT NULL,
    ec3firstname TEXT NOT NULL,
    ec3lastname TEXT NOT NULL,
    ec3relationship TEXT,
    ec3addressline1 TEXT NOT NULL,
    ec3housenamenumber TEXT,
    ec3addressline2 TEXT,
    ec3apartmentsuite TEXT,
    ec3town TEXT NOT NULL,
    ec3county TEXT,
    ec3postcode TEXT NOT NULL,
    ec3country TEXT NOT NULL,
    ec3homephone TEXT,
    ec3workphone TEXT,
    futureuse1 TEXT,
    futureuse2 TEXT,
    futureuse3 TEXT,
    futureuse4 TEXT,






    user_id INT(10) UNSIGNED NOT NULL,
    
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    
    CONSTRAINT `f_note_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    
    PRIMARY KEY (id)
);