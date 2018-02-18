--
-- PostgreSQL database dump
--

-- Dumped from database version 10.2
-- Dumped by pg_dump version 10.2

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: history; Type: TABLE; Schema: public; Owner: Junhwan
--

CREATE TABLE history (
    sequence integer NOT NULL,
    id integer NOT NULL,
    location character varying(20) NOT NULL,
    description character varying(60),
    adminid integer NOT NULL,
    scantime timestamp without time zone NOT NULL
);


ALTER TABLE history OWNER TO "Junhwan";

--
-- Name: student; Type: TABLE; Schema: public; Owner: Junhwan
--

CREATE TABLE student (
    id integer NOT NULL,
    number integer NOT NULL,
    name character varying(6) NOT NULL,
    points integer DEFAULT 0 NOT NULL,
    is_admin boolean NOT NULL,
    email character varying(500),
    password character(60),
    coupons integer DEFAULT 0 NOT NULL
);


ALTER TABLE student OWNER TO "Junhwan";

--
-- Data for Name: history; Type: TABLE DATA; Schema: public; Owner: Junhwan
--

COPY history (sequence, id, location, description, adminid, scantime) FROM stdin;
1	1	Main Entrance	late for school	1	2018-02-17 18:26:18
2	1	Gym	Entering	1	2018-02-17 21:15:20
\.


--
-- Data for Name: student; Type: TABLE DATA; Schema: public; Owner: Junhwan
--

COPY student (id, number, name, points, is_admin, email, password, coupons) FROM stdin;
2	10717	구영서	9	f	\N	\N	0
3	10101	잘몰라	0	f	\N	\N	0
4	20323	김창호	12	f	\N	\N	0
1	10423	범준환	609	t	solkblte@icloud.com	$2a$14$4ynRTwSKtIEZoKNE.j8yp.6oHFJHEBKgY1QwDrj2.ZYKfKNiebN6W	0
\.


--
-- Name: history history_pkey; Type: CONSTRAINT; Schema: public; Owner: Junhwan
--

ALTER TABLE ONLY history
    ADD CONSTRAINT history_pkey PRIMARY KEY (sequence);


--
-- Name: student student_pkey; Type: CONSTRAINT; Schema: public; Owner: Junhwan
--

ALTER TABLE ONLY student
    ADD CONSTRAINT student_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

