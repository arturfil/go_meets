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
('a78aaa5a-5e0a-4c7c-9944-741a76099750', 'Finance 204', 'Corporate Finance', 'c9f771e7-1293-4c72-9d5d-01b4c3e7e0f9', '2024-05-20 00:23:36.857134+00', '2024-05-20 00:23:36.857134+00'),
('d8c2d9e4-1a3b-4c5c-9d8e-7f6b5c4d3a2b', 'Linear Algebra', 'Vectors, matrices, and linear transformations', 'f47ac10b-58cc-4372-a567-0e02b2c3d479', '2024-05-20 00:23:36.857134+00', '2024-05-20 00:23:36.857134+00'),
('e9d3e0f5-2b4c-4a6d-ae9f-8f7c6d5e4b3c', 'Chemistry 101', 'Introduction to chemical principles and reactions', 'b6ea4744-b192-4a39-8b4f-f9e7f6eb40bc', '2024-05-20 00:23:36.857134+00', '2024-05-20 00:23:36.857134+00'),
('f0e4f1a5-3c5d-4b7e-bf0a-9a8d7e6f5c4d', 'Economics 101', 'Principles of microeconomics and macroeconomics', 'c9f771e7-1293-4c72-9d5d-01b4c3e7e0f9', '2024-05-20 00:23:36.857134+00', '2024-05-20 00:23:36.857134+00');

-- Insert users (unchanged)
INSERT INTO public.users VALUES 
('aef0a244-f285-4948-a8b7-68b610a879d0', 'arturo', 'filio', 'arturo@test.com', '$2a$12$ktYT0l..0Qii2JYijKA5JOJvo4fHmUlBQtZMSxc6Z3nUrLhfm8Ax2', '2024-04-22 03:11:44.021814+00', '2024-04-22 03:11:44.021814+00'),
('efebbf58-1946-4d9b-a5cf-c2d3c8f4aa0c', 'Adri', 'Test', 'adri@test.com', '$2a$12$2H.cLKBZTanDcE45n4dtcur2NXe.GwwSwOAf4X9ebZvsdJ4L95ewq', '2024-05-21 16:34:07.074783+00', '2024-05-21 16:34:07.074783+00'),
('c21c7824-f3c8-4b2d-9b4e-d7a325b6a115', 'Sarah', 'Johnson', 'sarah.j@test.com', '$2a$12$ktYT0l..0Qii2JYijKA5JOJvo4fHmUlBQtZMSxc6Z3nUrLhfm8Ax2', '2024-04-22 03:11:44.021814+00', '2024-04-22 03:11:44.021814+00'),
('57bf2843-ea1d-461e-8591-5376235f0c43', 'Michael', 'Chen', 'michael.c@test.com', '$2a$12$ktYT0l..0Qii2JYijKA5JOJvo4fHmUlBQtZMSxc6Z3nUrLhfm8Ax2', '2024-04-22 03:11:44.021814+00', '2024-04-22 03:11:44.021814+00'),
('6c271773-b516-40ab-b91c-1eafba25bc53', 'Emily', 'Brown', 'emily.b@test.com', '$2a$12$ktYT0l..0Qii2JYijKA5JOJvo4fHmUlBQtZMSxc6Z3nUrLhfm8Ax2', '2024-04-22 03:11:44.021814+00', '2024-04-22 03:11:44.021814+00'),
('23478866-27a8-4e5e-beda-fd1c39c30c72', 'David', 'Wilson', 'david.w@test.com', '$2a$12$ktYT0l..0Qii2JYijKA5JOJvo4fHmUlBQtZMSxc6Z3nUrLhfm8Ax2', '2024-04-22 03:11:44.021814+00', '2024-04-22 03:11:44.021814+00'),
('869d5d58-cf5d-4edb-b974-cff9bc230524', 'Maria', 'Garcia', 'maria.g@test.com', '$2a$12$ktYT0l..0Qii2JYijKA5JOJvo4fHmUlBQtZMSxc6Z3nUrLhfm8Ax2', '2024-04-22 03:11:44.021814+00', '2024-04-22 03:11:44.021814+00');

INSERT INTO public.teachings (id, teacher_id, subject_id, created_at, updated_at) VALUES
-- Sarah Johnson teaches Calculus 101 and Physics 101
('d4a29b36-8394-4d6a-b2a1-51c8f7876a1d', 'c21c7824-f3c8-4b2d-9b4e-d7a325b6a115', '0bd925f4-bbe7-4f03-8796-9e8badc4101e', '2024-05-20 00:23:36.857134+00', '2024-05-20 00:23:36.857134+00'),
('f8b24631-f755-4ee1-a536-42e6c88f68c7', 'c21c7824-f3c8-4b2d-9b4e-d7a325b6a115', 'cdf38e88-c6b7-4bcf-b699-b41ba60a1e8c', '2024-05-20 00:23:36.857134+00', '2024-05-20 00:23:36.857134+00'),

-- Michael Chen teaches Physics 101 and Finance 204
('7c5d1e93-2b6f-4c9a-b4d8-9e56f8d3a941', '57bf2843-ea1d-461e-8591-5376235f0c43', 'cdf38e88-c6b7-4bcf-b699-b41ba60a1e8c', '2024-05-20 00:23:36.857134+00', '2024-05-20 00:23:36.857134+00'),
('9e3a7d24-5c8b-4f1a-ae2d-6b9c12f45e83', '57bf2843-ea1d-461e-8591-5376235f0c43', 'a78aaa5a-5e0a-4c7c-9944-741a76099750', '2024-05-20 00:23:36.857134+00', '2024-05-20 00:23:36.857134+00'),

-- Maria Garcia teaches Calculus 101 and Finance 204
('b2c1f459-7d3e-4e8a-95f6-3c4d8a2f1b5e', '869d5d58-cf5d-4edb-b974-cff9bc230524', '0bd925f4-bbe7-4f03-8796-9e8badc4101e', '2024-05-20 00:23:36.857134+00', '2024-05-20 00:23:36.857134+00'),
('e5d4c3b2-a1f0-4d9e-8c7b-6b5a4f3e2d1c', '869d5d58-cf5d-4edb-b974-cff9bc230524', 'a78aaa5a-5e0a-4c7c-9944-741a76099750', '2024-05-20 00:23:36.857134+00', '2024-05-20 00:23:36.857134+00'),

-- Arturo teaches Linear Algebra, Chemistry 101, and Economics 101
('a1b2c3d4-e5f6-4a7b-8c9d-0e1f2a3b4c5d', 'aef0a244-f285-4948-a8b7-68b610a879d0', 'd8c2d9e4-1a3b-4c5c-9d8e-7f6b5c4d3a2b', '2024-05-20 00:23:36.857134+00', '2024-05-20 00:23:36.857134+00'),
('b2c3d4e5-f6a7-4b8c-9d0e-1f2a3b4c5d6e', 'aef0a244-f285-4948-a8b7-68b610a879d0', 'e9d3e0f5-2b4c-4a6d-ae9f-8f7c6d5e4b3c', '2024-05-20 00:23:36.857134+00', '2024-05-20 00:23:36.857134+00'),
('c3d4e5f6-a7b8-4c9d-0e1f-2a3b4c5d6e7f', 'aef0a244-f285-4948-a8b7-68b610a879d0', 'f0e4f1a5-3c5d-4b7e-bf0a-9a8d7e6f5c4d', '2024-05-20 00:23:36.857134+00', '2024-05-20 00:23:36.857134+00');

-- Insert requests (unchanged)
INSERT INTO requests(user_id, type) VALUES 
('aef0a244-f285-4948-a8b7-68b610a879d0', 'teach request'),
('efebbf58-1946-4d9b-a5cf-c2d3c8f4aa0c', 'teach request');

-- Insert roles (unchanged)
INSERT INTO roles (id, description) VALUES 
('52431aac-82d8-46df-8676-97b748e7bea4', 'admin'),
('71dc50c1-1934-4da1-91a5-2fb73fadb39e', 'teacher'),
('22b3f2ca-3e98-447f-a807-9609fa496ae9', 'user');

-- Insert role relations (unchanged)
INSERT INTO role_relations VALUES 
('efebbf58-1946-4d9b-a5cf-c2d3c8f4aa0c', '22b3f2ca-3e98-447f-a807-9609fa496ae9'),
('aef0a244-f285-4948-a8b7-68b610a879d0', '52431aac-82d8-46df-8676-97b748e7bea4'),
('aef0a244-f285-4948-a8b7-68b610a879d0', '22b3f2ca-3e98-447f-a807-9609fa496ae9'),
-- Sarah Johnson (Teacher + User)
('c21c7824-f3c8-4b2d-9b4e-d7a325b6a115', '71dc50c1-1934-4da1-91a5-2fb73fadb39e'),
('c21c7824-f3c8-4b2d-9b4e-d7a325b6a115', '22b3f2ca-3e98-447f-a807-9609fa496ae9'),
-- Michael Chen (Teacher + User)
('57bf2843-ea1d-461e-8591-5376235f0c43', '71dc50c1-1934-4da1-91a5-2fb73fadb39e'),
('57bf2843-ea1d-461e-8591-5376235f0c43', '22b3f2ca-3e98-447f-a807-9609fa496ae9'),
-- Emily Brown (User only)
('6c271773-b516-40ab-b91c-1eafba25bc53', '22b3f2ca-3e98-447f-a807-9609fa496ae9'),
-- David Wilson (User only)
('23478866-27a8-4e5e-beda-fd1c39c30c72', '22b3f2ca-3e98-447f-a807-9609fa496ae9'),
-- Maria Garcia (Teacher + User)
('869d5d58-cf5d-4edb-b974-cff9bc230524', '71dc50c1-1934-4da1-91a5-2fb73fadb39e'),
('869d5d58-cf5d-4edb-b974-cff9bc230524', '22b3f2ca-3e98-447f-a807-9609fa496ae9');

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
TRUNCATE TABLE teachings CASCADE;
-- +goose StatementEnd
