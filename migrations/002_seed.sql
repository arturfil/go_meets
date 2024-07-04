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
INSERT INTO public.users VALUES ('9ff612d2-f2d0-4118-a9fe-b4dee796d2c3', 'Test', 'User', true, true, 'test@user.com', '$2a$12$jqPQIqtZR5gZBupqvRzUfO7IAOCJJnjaAYhSwPxUS7P6DtyAeHUwu', '2024-05-20 00:23:36.861504+00', '2024-05-20 00:23:36.861504+00');
INSERT INTO public.users VALUES ('32c1755b-15e6-49b9-ad38-dc6a2243b54b', 'Santiago', 'Test', true, true, 'santiago@test.com', '$2a$12$t/CCDiZ8vcyV2Gf9N0uyTeGYzYjsxGDRt4dCBEJWYtr5hJkh6TJJu', '2024-05-20 00:23:36.861504+00', '2024-05-20 00:23:36.861504+00');
INSERT INTO public.users VALUES ('aef0a244-f285-4948-a8b7-68b610a879d0', 'arturo', 'filio', false, false, 'arturo@test.com', '$2a$12$ktYT0l..0Qii2JYijKA5JOJvo4fHmUlBQtZMSxc6Z3nUrLhfm8Ax2', '2024-04-22 03:11:44.021814+00', '2024-04-22 03:11:44.021814+00');
INSERT INTO public.users VALUES ('ef8070fe-af77-4fc3-9a45-8a52f033a287', 'arturo', 'filio', false, false, 'test2@test.com', '$2a$12$yzX/A6J5SurUYjTUJn3lQu94IX6iyCU1Kr5gB4pIahkK9wyXbnhJa', '2024-05-21 06:39:19.123593+00', '2024-05-21 06:39:19.123593+00');
INSERT INTO public.users VALUES ('0c699137-16de-4623-9e6a-4817d292338f', '', '', false, false, '', '$2a$12$C7NvJv9ZzRSuSddF5rOoEuby.6TwV0RrsidE/jHDxFRKQIWHWSfLu', '2024-05-21 06:40:02.693764+00', '2024-05-21 06:40:02.693764+00');
INSERT INTO public.users VALUES ('728414d3-2b3f-4fbc-8682-10c02c82fee0', 'arturo', 'filio', false, false, 'arturo2@test.com', '$2a$12$EMETXYpCSMTtZUkw26xTaOGHigAJePzAB19/KAhVMuTOlT3En6tfm', '2024-05-21 16:31:15.425828+00', '2024-05-21 16:31:15.425828+00');
INSERT INTO public.users VALUES ('efebbf58-1946-4d9b-a5cf-c2d3c8f4aa0c', 'Adri', 'Test', false, false, 'adri@test.com', '$2a$12$2H.cLKBZTanDcE45n4dtcur2NXe.GwwSwOAf4X9ebZvsdJ4L95ewq', '2024-05-21 16:34:07.074783+00', '2024-05-21 16:34:07.074783+00');

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
INSERT INTO public.meetings VALUES ('37007114-b1c8-42bb-8fdf-cd1eef2f6c73', 'cdf38e88-c6b7-4bcf-b699-b41ba60a1e8c', '9ff612d2-f2d0-4118-a9fe-b4dee796d2c3', '32c1755b-15e6-49b9-ad38-dc6a2243b54b', false, '2024-03-10 10:00:00+00', '2024-03-10 10:30:00+00', '2024-03-10 02:47:30.501952+00', '2024-03-10 02:47:30.501952+00');
INSERT INTO public.meetings VALUES ('89dc32bf-aa12-453e-868b-f2791467c2a2', 'cdf38e88-c6b7-4bcf-b699-b41ba60a1e8c', '9ff612d2-f2d0-4118-a9fe-b4dee796d2c3', '32c1755b-15e6-49b9-ad38-dc6a2243b54b', false, '2024-03-10 10:00:00+00', '2024-03-10 10:30:00+00', '2024-03-10 02:48:34.948496+00', '2024-03-10 02:48:34.948496+00');
INSERT INTO public.meetings VALUES ('9a2d6b3f-7783-4fdb-ac86-d53e27c12fe5', 'cdf38e88-c6b7-4bcf-b699-b41ba60a1e8c', '9ff612d2-f2d0-4118-a9fe-b4dee796d2c3', '32c1755b-15e6-49b9-ad38-dc6a2243b54b', false, '2024-03-10 10:00:00+00', '2024-03-10 10:30:00+00', '2024-03-10 02:48:44.421846+00', '2024-03-10 02:48:44.421846+00');
INSERT INTO public.meetings VALUES ('6b01ffca-b3a3-48f9-b766-8223a9c7e9da', 'cdf38e88-c6b7-4bcf-b699-b41ba60a1e8c', '9ff612d2-f2d0-4118-a9fe-b4dee796d2c3', '32c1755b-15e6-49b9-ad38-dc6a2243b54b', true, '2024-03-10 10:00:00+00', '2024-03-10 10:30:00+00', '2024-03-10 02:17:47.601559+00', '2024-03-10 02:17:47.601559+00');
INSERT INTO public.meetings VALUES ('d8f988e5-5409-449b-94d5-fe6590f103b2', 'cdf38e88-c6b7-4bcf-b699-b41ba60a1e8c', '9ff612d2-f2d0-4118-a9fe-b4dee796d2c3', '32c1755b-15e6-49b9-ad38-dc6a2243b54b', true, '2024-03-10 10:00:00+00', '2024-03-10 10:30:00+00', '2024-03-10 02:45:42.758144+00', '2024-03-10 02:45:42.758144+00');
INSERT INTO public.meetings VALUES ('fecb6ec1-fd05-41da-bd11-fdafe5a69693', '0bd925f4-bbe7-4f03-8796-9e8badc4101e', '9ff612d2-f2d0-4118-a9fe-b4dee796d2c3', '32c1755b-15e6-49b9-ad38-dc6a2243b54b', false, '2024-04-24 16:00:00+00', '2024-04-24 18:00:00+00', '2024-04-07 23:47:21.1+00', '2024-04-07 23:47:21.1+00');
INSERT INTO public.meetings VALUES ('9605dc9a-f130-435a-897e-9332ebd52e62', '0bd925f4-bbe7-4f03-8796-9e8badc4101e', 'aef0a244-f285-4948-a8b7-68b610a879d0', '9ff612d2-f2d0-4118-a9fe-b4dee796d2c3', false, '2024-05-17 15:30:00+00', '2024-05-17 16:30:00+00', '0001-01-01 00:00:00+00', '0001-01-01 00:00:00+00');
INSERT INTO public.meetings VALUES ('030524de-cb9d-45a8-901c-84df7c7220d6', '0bd925f4-bbe7-4f03-8796-9e8badc4101e', 'aef0a244-f285-4948-a8b7-68b610a879d0', '9ff612d2-f2d0-4118-a9fe-b4dee796d2c3', false, '2024-05-20 18:30:00+00', '2024-05-20 07:30:00+00', '0001-01-01 00:00:00+00', '0001-01-01 00:00:00+00');

INSERT INTO roles (id, description) VALUES ('52431aac-82d8-46df-8676-97b748e7bea4', 'admin');
INSERT INTO roles (id, description) VALUES ('71dc50c1-1934-4da1-91a5-2fb73fadb39e', 'teacher');
INSERT INTO roles (id, description) VALUES ('22b3f2ca-3e98-447f-a807-9609fa496ae9', 'user');

INSERT INTO role_relations VALUES ('9ff612d2-f2d0-4118-a9fe-b4dee796d2c3', '52431aac-82d8-46df-8676-97b748e7bea4');
INSERT INTO role_relations VALUES ('aef0a244-f285-4948-a8b7-68b610a879d0', '52431aac-82d8-46df-8676-97b748e7bea4');
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
