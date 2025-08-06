create table if not exists users  (
	id uuid primary key,
	email text unique not null,
	full_name text,
	password_hash text,
	user_role text check (user_role in ('applicant', 'employer')),
	created_at timestamp,
	updated_at timestamp
);

create table if not exists company (
	id uuid primary key,
	full_name text,
	email text unique not null,
	website text,
	owner_id uuid references users(id),
	user_rating float,
	created_at timestamp,
	updated_at timestamp
);

create table if not exists resume(
	id uuid primary key,
	user_id uuid references users(id),
	title text,
	experience text,
	skills text,
	education text,
	about text,
	updated_at timestamp
);

create table if not exists vacancy(
	id UUID PRIMARY KEY,
	company_id UUID REFERENCES Company(id),
	title TEXT,
	description TEXT,
	location TEXT,
	experience_required text,
	skills_required TEXT,
	salary_from INT,
	salary_to INT,
	created_at TIMESTAMP,
	updated_at TIMESTAMP

);

create table if not exists responses (
	id uuid primary key,
	resume_id uuid references resume(id),
	vacnacy_id uuid references vacancy(id),
	status text check (status in ('new', 'viewed','rejected','invited')),
	created_at timestamp,
	updated_at timestamp
);