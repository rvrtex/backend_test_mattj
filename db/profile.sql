
CREATE TABLE IF NOT EXISTS "default_language"(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),  
  "name" VARCHAR(128) NOT NULL UNIQUE,
  code VARCHAR(128) NOT NULL UNIQUE    
);


CREATE TABLE IF NOT EXISTS "phone_number" (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),  
  ext VARCHAR(128) NOT NULL UNIQUE,
  country_code VARCHAR(128) NOT NULL UNIQUE, 
  phone_number VARCHAR(128) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS "profile"(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),  
first_name VARCHAR(128) NOT NULL DEFAULT '',
  last_name VARCHAR(128) NOT NULL DEFAULT '',
  email varchar(128) NOT NULL DEFAULT '',
  default_language UUID NOT NULL REFERENCES default_language(id),
 street_address_primary VARCHAR(128) NOT NULL DEFAULT '',
 street_address_secondary VARCHAR(128) DEFAULT '',
  phone_number_primary UUID NOT NULL REFERENCES phone_number(id),
    phone_number_secondary UUID REFERENCES phone_number(id),
    city VARCHAR(128) NOT NULL DEFAULT '',
    state_providence VARCHAR(128) NOT NULL DEFAULT '',
    zip_code VARCHAR(128) NOT NULL DEFAULT '',
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


INSERT INTO public.default_language
(id, "name", code)
VALUES(gen_random_uuid(), 'english', 'EN-US');

INSERT INTO public.phone_number
(id, ext, country_code, phone_number)
VALUES(gen_random_uuid(), '1', '1', '867-5309');

INSERT INTO public.profile
(id, first_name, last_name, email, default_language, street_address_primary,  phone_number_primary, city, state_providence, zip_code, created_at, updated_at)
VALUES(gen_random_uuid(), 'John'::character varying, 'Doe'::character varying, 'johnDoe@gmail.com'::character varying, 'e248bd56-5cfd-4b34-98fc-62754926ae15'::uuid, '2020 Fake street'::character varying, 'fd8a3939-d5bd-4ae4-bd23-727854ab0bb1'::uuid,  'Fake City'::character varying, 'PA'::character varying, '00000'::character varying,  now(), now());