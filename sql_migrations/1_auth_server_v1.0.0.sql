-- +migrate Up

-- +migrate StatementBegin

CREATE TYPE locale_language AS ENUM ('en-US', 'id-ID');
-- CREATE TYPE grant_types AS ENUM ('code', 'code_pkce', 'token', 'password', 'client_credentials');

CREATE TABLE IF NOT EXISTS users
(
    id              SERIAL PRIMARY KEY,
    username        VARCHAR(50)  NOT NULL UNIQUE,
    password        VARCHAR(256) NOT NULL,
    salt            VARCHAR(256) NOT NULL,
    email           VARCHAR(100),
    phone           VARCHAR(20),
    status			CHAR(1),
    locale			VARCHAR(3),
    client_id 		VARCHAR(256) NOT NULL,
    person_profile_id INTEGER,
    created_by      INTEGER,
    created_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by      INTEGER,
    updated_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS auth_client
(
    client_id		VARCHAR(256) PRIMAY_KEY,
    client_secret	VARCHAR(256) NOT NULL,
    secret_key		VARCHAR(256) NOT NULL,
    grant_type		VARCHAR(50) NOT NULL,
    redirect_uri	VARCHAR(256),
    created_by      INTEGER,
    created_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by      INTEGER,
    updated_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS client_resource
(
    id              SERIAL PRIMARY KEY,
    client_id		VARCHAR(256) NOT NULL,
    resource_id		VARCHAR(256) NOT NULL,
    authorities		TEXT,
    created_by      INTEGER,
    created_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by      INTEGER,
    updated_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT uq_clientressource_clientid_resourceid UNIQUE (client_id, resource_id)
    );

CREATE TABLE IF NOT EXISTS client_resource
(
    id              SERIAL PRIMARY KEY,
    client_id		VARCHAR(256) NOT NULL,
    resource_id		VARCHAR(256) NOT NULL,
    authorities		TEXT,
    created_by      INTEGER,
    created_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by      INTEGER,
    updated_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT uq_clientressource_clientid_resourceid UNIQUE (client_id, resource_id)
    );

CREATE TABLE IF NOT EXISTS resource
(
    resource_id		VARCHAR(30) PRIMARY KEY,
    descrition 		VARCHAR(256),
    created_by      INTEGER,
    created_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by      INTEGER,
    updated_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS person_profile
(
    id				SERIAL PRIMARY KEY,
    first_name 		VARCHAR(100),
    last_name		VARCHAR(100),
    address_1		VARCHAR(256),
    address_2		VARCHAR(256),
    country_id		INTEGER,
    created_by      INTEGER,
    created_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by      INTEGER,
    updated_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS user_verification
(
    id				SERIAL PRIMARY KEY,
    user_id			INTEGER,
    email			VARCHAR(100),
    email_code		VARCHAR(20),
    email_expires	INTEGER,
    email_verified_at INTEGER,
    phone			VARCHAR(20),
    phone_code		VARCHAR(20),
    phone_expires	INTEGER,
    phone_veified_at INTEGER,
    forget_code		VARCHAR(20),
    forget_expires	INTEGER,
    forget_verified_at INTEGER,
    created_by      INTEGER,
    created_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by      INTEGER,
    updated_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

INSERT INTO users
(username, "password", salt)
VALUES
    ('admin', 'abc954ec4bd81efa9fd6b9250b54a62e19758a4ee40435cf8f849847e838a3356e473b8e9380dfaa3f2c5f51eb466009a87b8192ff1a70340fb5027bc2e18e76', '32bd68cc58b542ed8f44bc42d8c68083')
    ON CONFLICT(username) DO NOTHING;
-- +migrate StatementEnd