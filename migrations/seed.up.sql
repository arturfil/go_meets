INSERT INTO subjects (id, name, description)
VALUES 
    ('0bd925f4-bbe7-4f03-8796-9e8badc4101e', 'Calculus 101', 'Differential and Integral calculus'),
    ('cdf38e88-c6b7-4bcf-b699-b41ba60a1e8c', 'Physics 101', 'Thermodynamics and laws of momevent'),
    ('a78aaa5a-5e0a-4c7c-9944-741a76099750', 'Finance 204', 'Coroprate Finance')
;

INSERT INTO users (id, first_name, last_name, email, is_admin, is_teacher, password)
VALUES 
    ('9ff612d2-f2d0-4118-a9fe-b4dee796d2c3', 'Test', 'User', 'test@user.com', true, true, '$2a$12$jqPQIqtZR5gZBupqvRzUfO7IAOCJJnjaAYhSwPxUS7P6DtyAeHUwu'),
    ('32c1755b-15e6-49b9-ad38-dc6a2243b54b', 'Santiago', 'Test', 'santiago@test.com', true, true, '$2a$12$t/CCDiZ8vcyV2Gf9N0uyTeGYzYjsxGDRt4dCBEJWYtr5hJkh6TJJu')
;

-- insert a meeting later
