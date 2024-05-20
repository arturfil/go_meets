-- insert subjects
INSERT INTO subjects (id, name, description)
VALUES 
    ('0bd925f4-bbe7-4f03-8796-9e8badc4101e', 'Calculus 101', 'Differential and Integral calculus'),
    ('cdf38e88-c6b7-4bcf-b699-b41ba60a1e8c', 'Physics 101', 'Thermodynamics and laws of momevent'),
    ('a78aaa5a-5e0a-4c7c-9944-741a76099750', 'Finance 204', 'Coroprate Finance')
;

-- insert users
INSERT INTO users (id, first_name, last_name, email, is_admin, is_teacher, password)
VALUES 
    ('9ff612d2-f2d0-4118-a9fe-b4dee796d2c3', 'Test', 'User', 'test@user.com', true, true, '$2a$12$jqPQIqtZR5gZBupqvRzUfO7IAOCJJnjaAYhSwPxUS7P6DtyAeHUwu'),
    ('32c1755b-15e6-49b9-ad38-dc6a2243b54b', 'Santiago', 'Test', 'santiago@test.com', true, true, '$2a$12$t/CCDiZ8vcyV2Gf9N0uyTeGYzYjsxGDRt4dCBEJWYtr5hJkh6TJJu')
;

INSERT INTO public.meetings VALUES ('37007114-b1c8-42bb-8fdf-cd1eef2f6c73', 'cdf38e88-c6b7-4bcf-b699-b41ba60a1e8c', '9ff612d2-f2d0-4118-a9fe-b4dee796d2c3', '32c1755b-15e6-49b9-ad38-dc6a2243b54b', false, '2024-03-10 10:00:00+00', '2024-03-10 10:30:00+00', '2024-03-10 02:47:30.501952+00', '2024-03-10 02:47:30.501952+00');
INSERT INTO public.meetings VALUES ('89dc32bf-aa12-453e-868b-f2791467c2a2', 'cdf38e88-c6b7-4bcf-b699-b41ba60a1e8c', '9ff612d2-f2d0-4118-a9fe-b4dee796d2c3', '32c1755b-15e6-49b9-ad38-dc6a2243b54b', false, '2024-03-10 10:00:00+00', '2024-03-10 10:30:00+00', '2024-03-10 02:48:34.948496+00', '2024-03-10 02:48:34.948496+00');
INSERT INTO public.meetings VALUES ('9a2d6b3f-7783-4fdb-ac86-d53e27c12fe5', 'cdf38e88-c6b7-4bcf-b699-b41ba60a1e8c', '9ff612d2-f2d0-4118-a9fe-b4dee796d2c3', '32c1755b-15e6-49b9-ad38-dc6a2243b54b', false, '2024-03-10 10:00:00+00', '2024-03-10 10:30:00+00', '2024-03-10 02:48:44.421846+00', '2024-03-10 02:48:44.421846+00');
INSERT INTO public.meetings VALUES ('6b01ffca-b3a3-48f9-b766-8223a9c7e9da', 'cdf38e88-c6b7-4bcf-b699-b41ba60a1e8c', '9ff612d2-f2d0-4118-a9fe-b4dee796d2c3', '32c1755b-15e6-49b9-ad38-dc6a2243b54b', true, '2024-03-10 10:00:00+00', '2024-03-10 10:30:00+00', '2024-03-10 02:17:47.601559+00', '2024-03-10 02:17:47.601559+00');
INSERT INTO public.meetings VALUES ('d8f988e5-5409-449b-94d5-fe6590f103b2', 'cdf38e88-c6b7-4bcf-b699-b41ba60a1e8c', '9ff612d2-f2d0-4118-a9fe-b4dee796d2c3', '32c1755b-15e6-49b9-ad38-dc6a2243b54b', true, '2024-03-10 10:00:00+00', '2024-03-10 10:30:00+00', '2024-03-10 02:45:42.758144+00', '2024-03-10 02:45:42.758144+00');
INSERT INTO public.meetings VALUES ('fecb6ec1-fd05-41da-bd11-fdafe5a69693', '0bd925f4-bbe7-4f03-8796-9e8badc4101e', '9ff612d2-f2d0-4118-a9fe-b4dee796d2c3', '32c1755b-15e6-49b9-ad38-dc6a2243b54b', false, '2024-04-24 16:00:00+00', '2024-04-24 18:00:00+00', '2024-04-07 23:47:21.1+00', '2024-04-07 23:47:21.1+00');
INSERT INTO public.meetings VALUES ('9605dc9a-f130-435a-897e-9332ebd52e62', '0bd925f4-bbe7-4f03-8796-9e8badc4101e', 'aef0a244-f285-4948-a8b7-68b610a879d0', '9ff612d2-f2d0-4118-a9fe-b4dee796d2c3', false, '2024-05-17 15:30:00+00', '2024-05-17 16:30:00+00', '0001-01-01 00:00:00+00', '0001-01-01 00:00:00+00');
INSERT INTO public.meetings VALUES ('030524de-cb9d-45a8-901c-84df7c7220d6', '0bd925f4-bbe7-4f03-8796-9e8badc4101e', 'aef0a244-f285-4948-a8b7-68b610a879d0', '9ff612d2-f2d0-4118-a9fe-b4dee796d2c3', false, '2024-05-20 18:30:00+00', '2024-05-20 07:30:00+00', '0001-01-01 00:00:00+00', '0001-01-01 00:00:00+00');

-- INSERT INTO public.subjects VALUES ('0bd925f4-bbe7-4f03-8796-9e8badc4101e', 'Calculus 101', 'Differential and Integral calculus', '2024-03-09 19:54:06.359352+00', '2024-03-09 19:54:06.359352+00');
-- INSERT INTO public.subjects VALUES ('cdf38e88-c6b7-4bcf-b699-b41ba60a1e8c', 'Physics 101', 'Thermodynamics and laws of momevent', '2024-03-09 19:54:06.359352+00', '2024-03-09 19:54:06.359352+00');
-- INSERT INTO public.subjects VALUES ('a78aaa5a-5e0a-4c7c-9944-741a76099750', 'Finance 204', 'Coroprate Finance', '2024-03-09 19:54:06.359352+00', '2024-03-09 19:54:06.359352+00');


-- INSERT INTO public.users VALUES ('9ff612d2-f2d0-4118-a9fe-b4dee796d2c3', 'Test', 'User', true, true, 'test@user.com', '$2a$12$jqPQIqtZR5gZBupqvRzUfO7IAOCJJnjaAYhSwPxUS7P6DtyAeHUwu', '2024-03-09 19:54:53.041467+00', '2024-03-09 19:54:53.041467+00');
-- INSERT INTO public.users VALUES ('32c1755b-15e6-49b9-ad38-dc6a2243b54b', 'Santiago', 'Test', false, false, 'santiago@test.com', '$2a$12$t/CCDiZ8vcyV2Gf9N0uyTeGYzYjsxGDRt4dCBEJWYtr5hJkh6TJJu', '2024-03-09 20:34:51.389439+00', '2024-03-09 20:34:51.389439+00');
INSERT INTO public.users VALUES ('aef0a244-f285-4948-a8b7-68b610a879d0', 'arturo', 'filio', false, false, 'arturo@test.com', '$2a$12$ktYT0l..0Qii2JYijKA5JOJvo4fHmUlBQtZMSxc6Z3nUrLhfm8Ax2', '2024-04-22 03:11:44.021814+00', '2024-04-22 03:11:44.021814+00');

-- insert a meeting later
