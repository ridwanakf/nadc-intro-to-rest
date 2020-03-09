-- noinspection SqlNoDataSourceInspectionForFile

-- Create Table

CREATE TABLE public.nadc_mst_book (
      book_id serial NOT NULL,
      book_name varchar(250) NULL DEFAULT ''::character varying,
      book_rate float4 NULL DEFAULT 0.0,
      book_author_name varchar(250) NULL DEFAULT ''::character varying,
      book_category varchar(250) NULL DEFAULT ''::character varying
);

-- Permissions

ALTER TABLE public.nadc_mst_book OWNER TO postgres;
GRANT ALL ON TABLE public.nadc_mst_book TO postgres;

-- Populate sample data

INSERT INTO nadc_mst_book (book_name, book_rate, book_author_name, book_category) VALUES ('What If?: Serious Scientific Answers to Absurd Hypothetical Questions', 4.16, 'Munroe, Randall', 'Nonfiction') RETURNING book_id;
INSERT INTO nadc_mst_book (book_name, book_rate, book_author_name, book_category) VALUES ('The Alchemist', 3.86, 'Coelho, Paulo', 'Fiction') RETURNING book_id;
INSERT INTO nadc_mst_book (book_name, book_rate, book_author_name, book_category) VALUES ('We Were the Lucky Ones', 4.40, 'Hunter, Georgia', 'Historical') RETURNING book_id;
INSERT INTO nadc_mst_book (book_name, book_rate, book_author_name, book_category) VALUES ('Everything Is F*cked: A Book About Hope', 3.77, 'Manson, Mark', 'Nonfiction') RETURNING book_id;
INSERT INTO nadc_mst_book (book_name, book_rate, book_author_name, book_category) VALUES ('Looking for Alaska', 4.04, 'Green, John 	', 'Fiction') RETURNING book_id;
INSERT INTO nadc_mst_book (book_name, book_rate, book_author_name, book_category) VALUES ('Pillow Thoughts', 3.83, 'Peppernell, Courtney', 'Poetry') RETURNING book_id;
INSERT INTO nadc_mst_book (book_name, book_rate, book_author_name, book_category) VALUES ('The Subtle Art of Not Giving a F*ck', 3.96, 'Manson, Mark', 'Nonfiction') RETURNING book_id;
INSERT INTO nadc_mst_book (book_name, book_rate, book_author_name, book_category) VALUES ('Bad Blood: Secrets and Lies in a Silicon Valley Startup ', 4.45, 'Carreyrou, John', 'Nonfiction') RETURNING book_id;
