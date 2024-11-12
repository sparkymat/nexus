--
-- PostgreSQL database dump
--

-- Dumped from database version 16.3
-- Dumped by pg_dump version 16.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: moddatetime; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS moddatetime WITH SCHEMA public;


--
-- Name: EXTENSION moddatetime; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION moddatetime IS 'functions for tracking last modification time';


--
-- Name: property_type; Type: TYPE; Schema: public; Owner: -
--

CREATE TYPE public.property_type AS ENUM (
    'string',
    'integer',
    'float',
    'boolean',
    'object',
    'image',
    'date'
);


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: objects; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.objects (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text NOT NULL,
    alternate_names_csv text DEFAULT ''::text NOT NULL,
    is_template boolean DEFAULT false NOT NULL,
    template_id uuid,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


--
-- Name: properties; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.properties (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text NOT NULL,
    object_id uuid NOT NULL,
    template_id uuid,
    property_type public.property_type NOT NULL,
    string_value text,
    integer_value integer,
    float_value double precision,
    object_value_id uuid,
    date_value date,
    boolean_value boolean,
    image_path text,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text NOT NULL,
    email text NOT NULL,
    encrypted_password text NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


--
-- Name: objects objects_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.objects
    ADD CONSTRAINT objects_pkey PRIMARY KEY (id);


--
-- Name: properties properties_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.properties
    ADD CONSTRAINT properties_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: objects_alternate_names_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX objects_alternate_names_idx ON public.objects USING btree (alternate_names_csv);


--
-- Name: objects_name_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX objects_name_idx ON public.objects USING btree (name);


--
-- Name: properties_boolean_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX properties_boolean_idx ON public.properties USING btree (boolean_value);


--
-- Name: properties_date_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX properties_date_idx ON public.properties USING btree (date_value);


--
-- Name: properties_float_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX properties_float_idx ON public.properties USING btree (float_value);


--
-- Name: properties_image_path_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX properties_image_path_idx ON public.properties USING btree (image_path);


--
-- Name: properties_integer_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX properties_integer_idx ON public.properties USING btree (integer_value);


--
-- Name: properties_name_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX properties_name_idx ON public.properties USING btree (name);


--
-- Name: properties_object_id_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX properties_object_id_idx ON public.properties USING btree (object_id);


--
-- Name: properties_string_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX properties_string_idx ON public.properties USING btree (string_value);


--
-- Name: properties_template_id_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX properties_template_id_idx ON public.properties USING btree (template_id);


--
-- Name: users users_updated_at; Type: TRIGGER; Schema: public; Owner: -
--

CREATE TRIGGER users_updated_at BEFORE UPDATE ON public.users FOR EACH ROW EXECUTE FUNCTION public.moddatetime('updated_at');


--
-- Name: objects objects_template_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.objects
    ADD CONSTRAINT objects_template_id_fkey FOREIGN KEY (template_id) REFERENCES public.objects(id) ON DELETE SET NULL;


--
-- Name: properties properties_object_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.properties
    ADD CONSTRAINT properties_object_id_fkey FOREIGN KEY (object_id) REFERENCES public.objects(id) ON DELETE CASCADE;


--
-- Name: properties properties_object_value_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.properties
    ADD CONSTRAINT properties_object_value_id_fkey FOREIGN KEY (object_value_id) REFERENCES public.objects(id) ON DELETE SET NULL;


--
-- Name: properties properties_template_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.properties
    ADD CONSTRAINT properties_template_id_fkey FOREIGN KEY (template_id) REFERENCES public.properties(id) ON DELETE SET NULL;


--
-- PostgreSQL database dump complete
--

