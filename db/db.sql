--
-- PostgreSQL database dump
--

-- Dumped from database version 17.5
-- Dumped by pg_dump version 17.5

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


--
-- Name: product_type; Type: TYPE; Schema: public; Owner: postgres
--

CREATE TYPE public.product_type AS ENUM (
    'UNIT',
    'AREA',
    'LINEAR'
);


ALTER TYPE public.product_type OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: accesses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.accesses (
    role_id uuid NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public.accesses OWNER TO postgres;

--
-- Name: product_categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.product_categories (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    group_id uuid NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public.product_categories OWNER TO postgres;

--
-- Name: product_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.product_groups (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public.product_groups OWNER TO postgres;

--
-- Name: product_subcategories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.product_subcategories (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    category_id uuid NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public.product_subcategories OWNER TO postgres;

--
-- Name: products; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.products (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    subcategory_id uuid NOT NULL,
    name character varying NOT NULL,
    type character varying NOT NULL
);


ALTER TABLE public.products OWNER TO postgres;

--
-- Name: COLUMN products.type; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.products.type IS 'UNIT/AREA/LINEAR';


--
-- Name: role_groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.role_groups (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid NOT NULL,
    name character varying NOT NULL
);


ALTER TABLE public.role_groups OWNER TO postgres;

--
-- Name: roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.roles (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    role_group_id uuid NOT NULL,
    name character varying NOT NULL,
    label character varying NOT NULL
);


ALTER TABLE public.roles OWNER TO postgres;

--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    name character varying NOT NULL,
    email character varying NOT NULL,
    phone character varying NOT NULL,
    username character varying NOT NULL,
    password character varying NOT NULL,
    token text
);


ALTER TABLE public.users OWNER TO postgres;

--
-- Name: workspace_members; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.workspace_members (
    workspace_id uuid NOT NULL,
    user_id uuid NOT NULL,
    role_id uuid NOT NULL
);


ALTER TABLE public.workspace_members OWNER TO postgres;

--
-- Name: workspaces; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.workspaces (
    id uuid DEFAULT public.uuid_generate_v4() NOT NULL,
    user_id uuid NOT NULL,
    role_group_id uuid,
    name character varying NOT NULL,
    product_group_id uuid
);


ALTER TABLE public.workspaces OWNER TO postgres;

--
-- Name: roles group_role_details_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT group_role_details_pkey PRIMARY KEY (id);


--
-- Name: role_groups group_roles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role_groups
    ADD CONSTRAINT group_roles_pkey PRIMARY KEY (id);


--
-- Name: product_categories product_categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_categories
    ADD CONSTRAINT product_categories_pkey PRIMARY KEY (id);


--
-- Name: product_groups product_groups_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_groups
    ADD CONSTRAINT product_groups_pkey PRIMARY KEY (id);


--
-- Name: product_subcategories product_subcategories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_subcategories
    ADD CONSTRAINT product_subcategories_pkey PRIMARY KEY (id);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_phone_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_phone_key UNIQUE (phone);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- Name: workspace_members workspace_member_workspace_id_user_id_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.workspace_members
    ADD CONSTRAINT workspace_member_workspace_id_user_id_key UNIQUE (workspace_id, user_id);


--
-- Name: workspaces workspace_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.workspaces
    ADD CONSTRAINT workspace_pkey PRIMARY KEY (id);


--
-- Name: category_id_for_product_subcategories; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX category_id_for_product_subcategories ON public.product_subcategories USING btree (category_id);


--
-- Name: group_role_detail_id_for_workspace_member; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX group_role_detail_id_for_workspace_member ON public.workspace_members USING btree (role_id);


--
-- Name: group_role_id_for_workspace; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX group_role_id_for_workspace ON public.workspaces USING btree (role_group_id);


--
-- Name: group_roles_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX group_roles_id ON public.roles USING btree (role_group_id);


--
-- Name: idx_product_categories_group_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_product_categories_group_id ON public.product_categories USING btree (group_id);


--
-- Name: product_group_id_for_workspace; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX product_group_id_for_workspace ON public.workspaces USING btree (product_group_id);


--
-- Name: role_id_for_accesses; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX role_id_for_accesses ON public.accesses USING btree (role_id);


--
-- Name: subcategory_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX subcategory_id ON public.products USING btree (subcategory_id);


--
-- Name: user_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX user_id ON public.role_groups USING btree (user_id);


--
-- Name: user_id_for_workspace; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX user_id_for_workspace ON public.workspaces USING btree (user_id);


--
-- Name: user_id_for_workspace_member; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX user_id_for_workspace_member ON public.workspace_members USING btree (user_id);


--
-- Name: workspace_id_for_workspace_member; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX workspace_id_for_workspace_member ON public.workspace_members USING btree (workspace_id);


--
-- Name: workspace_members FK2_user_id_for_workspace_member; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.workspace_members
    ADD CONSTRAINT "FK2_user_id_for_workspace_member" FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: workspace_members FK3_group_role_detail_id_for_workspace_member; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.workspace_members
    ADD CONSTRAINT "FK3_group_role_detail_id_for_workspace_member" FOREIGN KEY (role_id) REFERENCES public.roles(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: accesses FK__roles_for_accesses; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.accesses
    ADD CONSTRAINT "FK__roles_for_accesses" FOREIGN KEY (role_id) REFERENCES public.roles(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: roles FK_group_role_details_group_roles; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT "FK_group_role_details_group_roles" FOREIGN KEY (role_group_id) REFERENCES public.role_groups(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: role_groups FK_group_roles_users; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role_groups
    ADD CONSTRAINT "FK_group_roles_users" FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: product_categories FK_product_categories_product_groups; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_categories
    ADD CONSTRAINT "FK_product_categories_product_groups" FOREIGN KEY (group_id) REFERENCES public.product_groups(id) ON UPDATE CASCADE;


--
-- Name: product_subcategories FK_product_subcategories_product_categories; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.product_subcategories
    ADD CONSTRAINT "FK_product_subcategories_product_categories" FOREIGN KEY (category_id) REFERENCES public.product_categories(id);


--
-- Name: products FK_products_product_subcategories; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT "FK_products_product_subcategories" FOREIGN KEY (subcategory_id) REFERENCES public.product_subcategories(id);


--
-- Name: workspaces FK_workspace_group_roles; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.workspaces
    ADD CONSTRAINT "FK_workspace_group_roles" FOREIGN KEY (role_group_id) REFERENCES public.role_groups(id) ON UPDATE SET NULL ON DELETE SET NULL;


--
-- Name: workspace_members FK_workspace_id_for_workspace_member; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.workspace_members
    ADD CONSTRAINT "FK_workspace_id_for_workspace_member" FOREIGN KEY (workspace_id) REFERENCES public.workspaces(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: workspaces FK_workspace_users_owner; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.workspaces
    ADD CONSTRAINT "FK_workspace_users_owner" FOREIGN KEY (user_id) REFERENCES public.users(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: workspaces fk_workspace_product_group; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.workspaces
    ADD CONSTRAINT fk_workspace_product_group FOREIGN KEY (product_group_id) REFERENCES public.product_groups(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- PostgreSQL database dump complete
--

