-- +migrate Up

-- +migrate StatementBegin

DO $$ BEGIN
CREATE TYPE locale_language AS ENUM ('en-US', 'id-ID');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;
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
    locale			VARCHAR(5),
    client_id 		VARCHAR(256) NOT NULL,
    person_profile_id INTEGER,
    created_by      INTEGER,
    created_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_by      INTEGER,
    updated_at 		TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

CREATE TABLE IF NOT EXISTS auth_client
(
    client_id		VARCHAR(256) PRIMARY KEY,
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

CREATE TABLE IF NOT EXISTS auth_token (
  id 				SERIAL PRIMARY KEY,
  tokens 			TEXT,
  expired_time 	    INTEGER,
  value_token 			TEXT,
  created_at		TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS access_permissions
(
    user_id         BIGINT PRIMARY KEY,
    company_id      BIGINT NOT NULL,
    branch_id       BIGINT[] NOT NULL,
    is_admin        boolean NOT NULL DEFAULT FALSE
);

INSERT INTO users
(username, "password", salt, client_id)
VALUES
    ('admin', '58250cd9882e9ed8a9ca5c1b47aca886c904821eb17e6db1ddc28fd68e6a3a630edca05b471e85743c69e8c770ed304390046cdaac8724e146872da80ac062ab', '496c42e1c3484cd687962fb63332c868', 'b789a2c2f2f14e3085f556fc9de7da2a')
    ON CONFLICT(username) DO NOTHING;

INSERT INTO auth_client
(client_id, client_secret, secret_key, grant_type, created_by, updated_by)
VALUES
    ('b789a2c2f2f14e3085f556fc9de7da2a', 'secret-test', 'key-test', 'password', 1, 1);

INSERT INTO resource
(resource_id, descrition, created_by, updated_by)
VALUES
    ('auth', 'Server auth', 1, 1);

INSERT INTO client_resource
(client_id, resource_id, authorities, created_by, updated_by)
VALUES
    ('b789a2c2f2f14e3085f556fc9de7da2a', 'auth', '{"ALL":"ALL"}', 1, 1);
-- +migrate StatementEnd