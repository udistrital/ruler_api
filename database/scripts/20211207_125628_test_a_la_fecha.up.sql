--
-- PostgreSQL database dump
--

-- Dumped from database version 12.7
-- Dumped by pg_dump version 12.9 (Ubuntu 12.9-0ubuntu0.20.04.1)

-- Started on 2021-12-10 18:16:06 -05

-- SET statement_timeout = 0;
-- SET lock_timeout = 0;
-- SET idle_in_transaction_session_timeout = 0;
-- SET client_encoding = 'UTF8';
-- SET standard_conforming_strings = on;
-- SELECT pg_catalog.set_config('search_path', '', false);
-- SET check_function_bodies = false;
-- SET xmloption = content;
-- SET client_min_messages = warning;
-- SET row_security = off;

--
-- TOC entry 10 (class 2615 OID 39763)
-- Name: ruler; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA ruler;


--
-- TOC entry 1739 (class 1259 OID 757263)
-- Name: dominio_id_seq; Type: SEQUENCE; Schema: ruler; Owner: -
--

CREATE SEQUENCE ruler.dominio_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


-- SET default_table_access_method = heap;

--
-- TOC entry 1740 (class 1259 OID 757269)
-- Name: dominio; Type: TABLE; Schema: ruler; Owner: -
--

CREATE TABLE ruler.dominio (
    id integer DEFAULT nextval('ruler.dominio_id_seq'::regclass) NOT NULL,
    nombre character varying(35) NOT NULL,
    descripcion character varying(150)
);


--
-- TOC entry 9280 (class 0 OID 0)
-- Dependencies: 1740
-- Name: TABLE dominio; Type: COMMENT; Schema: ruler; Owner: -
--

COMMENT ON TABLE ruler.dominio IS 'Tabla de dominios del ruler';


--
-- TOC entry 9281 (class 0 OID 0)
-- Dependencies: 1740
-- Name: COLUMN dominio.id; Type: COMMENT; Schema: ruler; Owner: -
--

COMMENT ON COLUMN ruler.dominio.id IS 'Identificador del dominio';


--
-- TOC entry 9282 (class 0 OID 0)
-- Dependencies: 1740
-- Name: COLUMN dominio.nombre; Type: COMMENT; Schema: ruler; Owner: -
--

COMMENT ON COLUMN ruler.dominio.nombre IS 'Nombre del dominio';


--
-- TOC entry 9283 (class 0 OID 0)
-- Dependencies: 1740
-- Name: COLUMN dominio.descripcion; Type: COMMENT; Schema: ruler; Owner: -
--

COMMENT ON COLUMN ruler.dominio.descripcion IS 'Descripcion del dominio';


--
-- TOC entry 1741 (class 1259 OID 757273)
-- Name: predicado_id_seq; Type: SEQUENCE; Schema: ruler; Owner: -
--

CREATE SEQUENCE ruler.predicado_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 1742 (class 1259 OID 757275)
-- Name: predicado; Type: TABLE; Schema: ruler; Owner: -
--

CREATE TABLE ruler.predicado (
    id integer DEFAULT nextval('ruler.predicado_id_seq'::regclass) NOT NULL,
    dominio integer NOT NULL,
    nombre character varying NOT NULL,
    descripcion character varying(150) NOT NULL,
    tipo_predicado smallint NOT NULL
);


--
-- TOC entry 9284 (class 0 OID 0)
-- Dependencies: 1742
-- Name: TABLE predicado; Type: COMMENT; Schema: ruler; Owner: -
--

COMMENT ON TABLE ruler.predicado IS 'Tabla de predicados del ruler';


--
-- TOC entry 9285 (class 0 OID 0)
-- Dependencies: 1742
-- Name: COLUMN predicado.id; Type: COMMENT; Schema: ruler; Owner: -
--

COMMENT ON COLUMN ruler.predicado.id IS 'Identificador del predicado';


--
-- TOC entry 9286 (class 0 OID 0)
-- Dependencies: 1742
-- Name: COLUMN predicado.dominio; Type: COMMENT; Schema: ruler; Owner: -
--

COMMENT ON COLUMN ruler.predicado.dominio IS 'Identificador del dominio';


--
-- TOC entry 9287 (class 0 OID 0)
-- Dependencies: 1742
-- Name: COLUMN predicado.nombre; Type: COMMENT; Schema: ruler; Owner: -
--

COMMENT ON COLUMN ruler.predicado.nombre IS 'Texto del predicado';


--
-- TOC entry 9288 (class 0 OID 0)
-- Dependencies: 1742
-- Name: COLUMN predicado.descripcion; Type: COMMENT; Schema: ruler; Owner: -
--

COMMENT ON COLUMN ruler.predicado.descripcion IS 'Descripcion del predicado';


--
-- TOC entry 9289 (class 0 OID 0)
-- Dependencies: 1742
-- Name: COLUMN predicado.tipo_predicado; Type: COMMENT; Schema: ruler; Owner: -
--

COMMENT ON COLUMN ruler.predicado.tipo_predicado IS 'Define si es un hecho o una regla';


--
-- TOC entry 1743 (class 1259 OID 757282)
-- Name: tipo_predicado_id_seq; Type: SEQUENCE; Schema: ruler; Owner: -
--

CREATE SEQUENCE ruler.tipo_predicado_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 1744 (class 1259 OID 757284)
-- Name: tipo_predicado; Type: TABLE; Schema: ruler; Owner: -
--

CREATE TABLE ruler.tipo_predicado (
    id integer DEFAULT nextval('ruler.tipo_predicado_id_seq'::regclass) NOT NULL,
    nombre character varying(20)
);


--
-- TOC entry 9290 (class 0 OID 0)
-- Dependencies: 1744
-- Name: TABLE tipo_predicado; Type: COMMENT; Schema: ruler; Owner: -
--

COMMENT ON TABLE ruler.tipo_predicado IS 'Tabla de tipos de predicados del ruler';


--
-- TOC entry 9291 (class 0 OID 0)
-- Dependencies: 1744
-- Name: COLUMN tipo_predicado.id; Type: COMMENT; Schema: ruler; Owner: -
--

COMMENT ON COLUMN ruler.tipo_predicado.id IS 'Identificador del tipo de predicado';


--
-- TOC entry 9292 (class 0 OID 0)
-- Dependencies: 1744
-- Name: COLUMN tipo_predicado.nombre; Type: COMMENT; Schema: ruler; Owner: -
--

COMMENT ON COLUMN ruler.tipo_predicado.nombre IS 'Nombre del tipo de predicado';


-- Completed on 2021-12-10 18:16:16 -05

--
-- PostgreSQL database dump complete
--
