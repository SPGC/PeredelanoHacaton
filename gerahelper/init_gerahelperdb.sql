CREATE TABLE organisations
(
    Id       SERIAL PRIMARY KEY,
    Country  CHARACTER VARYING(50),
    Name     CHARACTER VARYING(50) NOT NULL,
    Contacts CHARACTER VARYING(100),
    Type     CHARACTER VARYING(50)
);

CREATE TABLE users
(
    Id      SERIAL PRIMARY KEY,
    Name    CHARACTER VARYING(30),
    Contact CHARACTER VARYING(30) UNIQUE NOT NULL
);

CREATE TABLE issues
(
    Id              SERIAL PRIMARY KEY,
    Status          CHARACTER VARYING(30),
    Description     TEXT,
    Organisation_id INTEGER,
    User_id         INTEGER,
    Validation      BOOLEAN,
    CONSTRAINT fk_organisation FOREIGN KEY (Organisation_id) REFERENCES organisations (Id),
    CONSTRAINT fk_user FOREIGN KEY (User_id) REFERENCES users (Id)
);

CREATE TABLE messages
(
    Id       SERIAL PRIMARY KEY,
    Data     TEXT,
    Date     DATE,
    Issue_id INTEGER,
    CONSTRAINT fk_author FOREIGN KEY (Issue_id) REFERENCES issues (Id)
);

CREATE PROCEDURE insert_issue(UserName CHARACTER VARYING(50), UserContactInfo CHARACTER VARYING(100),
                              Description TEXT, OrganisationName CHARACTER VARYING(50),
                              OrganisationCountry CHARACTER VARYING(50),
                              OrganisationContactInfo CHARACTER VARYING(100),
                              OrganisationType CHARACTER VARYING(50))
    LANGUAGE sql AS
$$

INSERT INTO users (Name, Contact)
VALUES (UserName, UserContactInfo)
ON CONFLICT DO NOTHING;

INSERT INTO organisations (Country, Name, Contacts, Type)
SELECT OrganisationCountry,
       OrganisationName,
       OrganisationContactInfo,
       OrganisationType
WHERE NOT EXISTS (SELECT Id FROM organisations WHERE Country = OrganisationCountry AND Name = OrganisationName);

INSERT INTO issues (Status, Description, Organisation_id, Validation, User_id)
SELECT 'New', Description, Id, FALSE, (SELECT id from users where Contact=UserContactInfo)
FROM organisations
WHERE Country = OrganisationCountry
  AND Name = OrganisationName;
$$;

create procedure update_organisation(    P_Id       INT,
                                         P_Country  CHARACTER VARYING(50),
                                         P_Name     CHARACTER VARYING(50),
                                         P_Contacts CHARACTER VARYING(100),
                                         P_Type     CHARACTER VARYING(50))
    language plpgsql as $$
declare
    new_name_id int;
begin
    if exists(select * from organisations o where o.country=P_Country and o.name=P_Name) then
        if P_Id = (select id from organisations o where o.country=P_Country and o.name=P_Name) then
            update organisations set
                                     country = P_Country, name=P_Name, contacts=P_Contacts, type=P_Type
            where id = P_Id;
        else
            new_name_id := (select id from organisations o where o.country=P_Country and o.name=P_Name);
            update issues set organisation_id=new_name_id where organisation_id=P_Id;
            delete from organisations where id = P_Id;
        end if;
    else
        update organisations set
                                 country = P_Country, name=P_Name, contacts=P_Contacts, type=P_Type
        where id = P_Id;
    end if;
end;
$$;