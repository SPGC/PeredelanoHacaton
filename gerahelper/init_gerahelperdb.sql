CREATE TABLE organisations
(
    Id       SERIAL PRIMARY KEY,
    Country  CHARACTER VARYING(50),
    Name     CHARACTER VARYING(50) NO NULL,
    Contacts CHARACTER VARYING(100),
    Type     CHARACTER VARYING(50)
);

CREATE TABLE users
(
    Id      SERIAL PRIMARY KEY,
    Name    CHARACTER VARYING(30),
    Contact CHARACTER VARYING(30) UNIQUE NO NULL
);

CREATE TABLE issues
(
    Id              SERIAL PRIMARY KEY,
    Status          CHARACTER VARYING(30),
    Description     CHARACTER VARYING(300),
    Organisation_id INTEGER,
    Validation      BOOLEAN,
    CONSTRAINT fk_organisation FOREIGN KEY (Organisation_id) REFERENCES organisations (Id)
);

CREATE TABLE messages
(
    Id       SERIAL PRIMARY KEY,
    Data     TIMESTAMP,
    Date     CHARACTER VARYING(4000),
    Issue_id INTEGER,
    CONSTRAINT fk_author FOREIGN KEY (Issue_id) REFERENCES issues (Id)
);

CREATE PROCEDURE insert_issue(UserName CHARACTER VARYING (50), UserContactInfo CHARACTER VARYING (100),
                              Description CHARACTER VARYING (300), OrganisationName CHARACTER VARYING (50),
                              OrganisationCountry CHARACTER VARYING (50),
                              OrganisationContactInfo CHARACTER VARYING (100),
                              OrganisationType CHARACTER VARYING (50))
    LANGUAGE sql AS
$$
INSERT INTO users (Name, Contact)
VALUES (UserName, UserContactInfo)
ON CONFLICT DO NOTHING;

INSERT INTO organisations (Country, Name, Contacts, Type)
SELECT OrganisationCountry,
       OrganisationName,
       OrganisationContactInfo,
       OrganisationType WHERE NOT EXISTS (SELECT Id FROM organisations WHERE Country = OrganisationCountry AND Name = OrganisationName);

INSERT INTO issues (Status, Description, Organisation_id, Validation)
SELECT 'New', Description, Id, FALSE
FROM organisations
WHERE Country = OrganisationCountry
  AND Name = OrganisationName;
$$;