-- +goose Up
-- +goose StatementBegin
--
-- TOC entry 3384 (class 0 OID 16666)
-- Dependencies: 213
-- Data for Name: subjects; Type: TABLE DATA; Schema: public; Owner: postgres
--
INSERT INTO public.subjects VALUES ('0bd925f4-bbe7-4f03-8796-9e8badc4101e', 'Calculus 101', 'Differential and Integral calculus', '2024-05-20 00:23:36.857134+00', '2024-05-20 00:23:36.857134+00');
INSERT INTO public.subjects VALUES ('cdf38e88-c6b7-4bcf-b699-b41ba60a1e8c', 'Physics 101', 'Thermodynamics and laws of momevent', '2024-05-20 00:23:36.857134+00', '2024-05-20 00:23:36.857134+00');
INSERT INTO public.subjects VALUES ('a78aaa5a-5e0a-4c7c-9944-741a76099750', 'Finance 204', 'Coroprate Finance', '2024-05-20 00:23:36.857134+00', '2024-05-20 00:23:36.857134+00');

--
-- TOC entry 3381 (class 0 OID 16632)
-- Dependencies: 210
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--
INSERT INTO public.users VALUES ('aef0a244-f285-4948-a8b7-68b610a879d0', 'arturo', 'filio', 'arturo@test.com', '$2a$12$ktYT0l..0Qii2JYijKA5JOJvo4fHmUlBQtZMSxc6Z3nUrLhfm8Ax2', '2024-04-22 03:11:44.021814+00', '2024-04-22 03:11:44.021814+00');
INSERT INTO public.users VALUES ('efebbf58-1946-4d9b-a5cf-c2d3c8f4aa0c', 'Adri', 'Test', 'adri@test.com', '$2a$12$2H.cLKBZTanDcE45n4dtcur2NXe.GwwSwOAf4X9ebZvsdJ4L95ewq', '2024-05-21 16:34:07.074783+00', '2024-05-21 16:34:07.074783+00');

--
-- TOC entry 3382 (class 0 OID 16646)
-- Dependencies: 211
-- Data for Name: requests; Type: TABLE DATA; Schema: public; Owner: postgres
--
INSERT INTO public.requests VALUES ('aef0a244-f285-4948-a8b7-68b610a879d0', 'pending', 'teach request');
INSERT INTO public.requests VALUES ('efebbf58-1946-4d9b-a5cf-c2d3c8f4aa0c', 'pending', 'teach request');

--
-- TOC entry 3383 (class 0 OID 16658)
-- Dependencies: 212
-- Data for Name: meetings; Type: TABLE DATA; Schema: public; Owner: postgres
--
INSERT INTO roles (id, description) VALUES ('52431aac-82d8-46df-8676-97b748e7bea4', 'admin');
INSERT INTO roles (id, description) VALUES ('71dc50c1-1934-4da1-91a5-2fb73fadb39e', 'teacher');
INSERT INTO roles (id, description) VALUES ('22b3f2ca-3e98-447f-a807-9609fa496ae9', 'user');

INSERT INTO role_relations VALUES ('efebbf58-1946-4d9b-a5cf-c2d3c8f4aa0c', '71dc50c1-1934-4da1-91a5-2fb73fadb39e');
INSERT INTO role_relations VALUES ('efebbf58-1946-4d9b-a5cf-c2d3c8f4aa0c', '22b3f2ca-3e98-447f-a807-9609fa496ae9');
INSERT INTO role_relations VALUES ('aef0a244-f285-4948-a8b7-68b610a879d0', '52431aac-82d8-46df-8676-97b748e7bea4');
INSERT INTO role_relations VALUES ('aef0a244-f285-4948-a8b7-68b610a879d0', '22b3f2ca-3e98-447f-a807-9609fa496ae9');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE subjects CASCADE;
TRUNCATE TABLE users CASCADE;
TRUNCATE TABLE requests CASCADE;
TRUNCATE TABLE meetings CASCADE;
TRUNCATE TABLE roles CASCADE;
TRUNCATE TABLE role_relations CASCADE;
-- +goose StatementEnd
