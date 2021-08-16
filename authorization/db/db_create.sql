DROP TABLE IF EXISTS accounts;
DROP TABLE IF EXISTS actions;

CREATE TABLE accounts (
    id integer NOT NULL PRIMARY KEY,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    email text,
    password text,
    token text
);

ALTER TABLE accounts OWNER TO postgres;

CREATE SEQUENCE accounts_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE accounts_id_seq OWNER TO postgres;

ALTER SEQUENCE accounts_id_seq OWNED BY accounts.id;

ALTER TABLE ONLY accounts ALTER COLUMN id SET DEFAULT nextval('accounts_id_seq'::regclass);
	
CREATE TABLE actions (
    id integer NOT NULL PRIMARY KEY,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    action text,
    date_time text,
	user_id integer NOT NULL,
	FOREIGN KEY (user_id) REFERENCES accounts (id) ON DELETE CASCADE
);

ALTER TABLE actions OWNER TO postgres;

CREATE SEQUENCE actions_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER TABLE actions_id_seq OWNER TO postgres;

ALTER SEQUENCE actions_id_seq OWNED BY actions.id;

ALTER TABLE ONLY actions ALTER COLUMN id SET DEFAULT nextval('actions_id_seq'::regclass);

INSERT INTO accounts (created_at, updated_at, deleted_at, email, password) VALUES
  ('2021-08-16 14:58:58.677865+03', '2021-08-16 14:58:58.677865+03', NULL, 'test1@gmail.com', '$2a$10$XY18bjI7/e09OI1Mdxn3weVXOfn4AdHkzArG/pIugV5GBnQ..51Im'),
  ('2021-08-16 14:59:14.430128+03', '2021-08-16 14:59:14.430128+03', NULL, 'test2@gmail.com', '$2a$10$meD.YWM5POw./9Q89EiYeukwLqlQeT5GVW.cmkAjD5TbjYB8njgbq');

INSERT INTO actions (created_at, updated_at, deleted_at, action, date_time, user_id) VALUES
  ('2021-08-16 15:07:10.565546+03', '2021-08-16 15:07:10.565546+03', NULL, 'Log in', '2021-08-16 15:07:10', '1'),
  ('2021-08-16 15:07:12.35871+03', '2021-08-16 15:07:12.35871+03', NULL, 'Log out', '2021-08-16 15:07:12', '1'),
  ('2021-08-16 15:07:25.960248+03', '2021-08-16 15:07:25.960248+03', NULL, 'Log in', '2021-08-16 15:07:25', '1'),
  ('2021-08-16 15:07:40.772148+03', '2021-08-16 15:07:40.772148+03', NULL, 'Log in', '2021-08-16 15:07:40', '1'),
  ('2021-08-16 15:07:52.579216+03', '2021-08-16 15:07:52.579216+03', NULL, 'Log in', '2021-08-16 15:07:52', '2'),
  ('2021-08-16 15:07:54.082654+03', '2021-08-16 15:07:54.082654+03', NULL, 'Log out', '2021-08-16 15:07:54', '2'),
  ('2021-08-16 15:08:01.801402+03', '2021-08-16 15:08:01.801402+03', NULL, 'Log in', '2021-08-16 15:08:01', '2'),
  ('2021-08-16 15:08:03.318391+03', '2021-08-16 15:08:03.318391+03', NULL, 'Log out', '2021-08-16 15:08:03', '2');
	