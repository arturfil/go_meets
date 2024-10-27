-- +goose Up
-- +goose StatementBegin

-- Insert subject categories
INSERT INTO public.subject_categories (id, name, description) VALUES
('f47ac10b-58cc-4372-a567-0e02b2c3d479', 'Mathematics', 'Courses related to mathematical concepts and applications'),
('b6ea4744-b192-4a39-8b4f-f9e7f6eb40bc', 'Science', 'Courses covering various scientific disciplines'),
('c9f771e7-1293-4c72-9d5d-01b4c3e7e0f9', 'Business', 'Courses focusing on business and finance topics');

-- Insert subjects with categories
INSERT INTO public.subjects VALUES 
('0bd925f4-bbe7-4f03-8796-9e8badc4101e', 'Calculus 101', 'Differential and Integral calculus', 'f47ac10b-58cc-4372-a567-0e02b2c3d479', '2024-05-20 00:23:36.857134+00', '2024-05-20 00:23:36.857134+00'),
('cdf38e88-c6b7-4bcf-b699-b41ba60a1e8c', 'Physics 101', 'Thermodynamics and laws of movement', 'b6ea4744-b192-4a39-8b4f-f9e7f6eb40bc', '2024-05-20 00:23:36.857134+00', '2024-05-20 00:23:36.857134+00'),
('a78aaa5a-5e0a-4c7c-9944-741a76099750', 'Finance 204', 'Corporate Finance', 'c9f771e7-1293-4c72-9d5d-01b4c3e7e0f9', '2024-05-20 00:23:36.857134+00', '2024-05-20 00:23:36.857134+00');

-- Insert users (unchanged)
INSERT INTO public.users VALUES 
('aef0a244-f285-4948-a8b7-68b610a879d0', 'arturo', 'filio', 'arturo@test.com', '$2a$12$ktYT0l..0Qii2JYijKA5JOJvo4fHmUlBQtZMSxc6Z3nUrLhfm8Ax2', '2024-04-22 03:11:44.021814+00', '2024-04-22 03:11:44.021814+00'),
('efebbf58-1946-4d9b-a5cf-c2d3c8f4aa0c', 'Adri', 'Test', 'adri@test.com', '$2a$12$2H.cLKBZTanDcE45n4dtcur2NXe.GwwSwOAf4X9ebZvsdJ4L95ewq', '2024-05-21 16:34:07.074783+00', '2024-05-21 16:34:07.074783+00');

-- Insert requests (unchanged)
INSERT INTO public.requests VALUES 
('aef0a244-f285-4948-a8b7-68b610a879d0', 'pending', 'teach request'),
('efebbf58-1946-4d9b-a5cf-c2d3c8f4aa0c', 'pending', 'teach request');

-- Insert roles (unchanged)
INSERT INTO roles (id, description) VALUES 
('52431aac-82d8-46df-8676-97b748e7bea4', 'admin'),
('71dc50c1-1934-4da1-91a5-2fb73fadb39e', 'teacher'),
('22b3f2ca-3e98-447f-a807-9609fa496ae9', 'user');

-- Insert role relations (unchanged)
INSERT INTO role_relations VALUES 
('efebbf58-1946-4d9b-a5cf-c2d3c8f4aa0c', '71dc50c1-1934-4da1-91a5-2fb73fadb39e'),
('efebbf58-1946-4d9b-a5cf-c2d3c8f4aa0c', '22b3f2ca-3e98-447f-a807-9609fa496ae9'),
('aef0a244-f285-4948-a8b7-68b610a879d0', '52431aac-82d8-46df-8676-97b748e7bea4'),
('aef0a244-f285-4948-a8b7-68b610a879d0', '22b3f2ca-3e98-447f-a807-9609fa496ae9');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
TRUNCATE TABLE subject_categories CASCADE;
TRUNCATE TABLE subjects CASCADE;
TRUNCATE TABLE users CASCADE;
TRUNCATE TABLE requests CASCADE;
TRUNCATE TABLE meetings CASCADE;
TRUNCATE TABLE roles CASCADE;
TRUNCATE TABLE role_relations CASCADE;
-- +goose StatementEnd
