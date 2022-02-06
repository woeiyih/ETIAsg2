CREATE database a2_db;

USE a2_db;

CREATE TABLE Students (StudentID VARCHAR (30) NOT NULL PRIMARY KEY, StudentName VARCHAR (30), DOB VARCHAR(30), Address VARCHAR (30), PhoneNumber VARCHAR (30)); 

INSERT INTO Students (StudentID, StudentName, DOB, Address, PhoneNumber) VALUES ("1", "Jake", "02/2/2002", "12 Chestnut Drive", "81112222");

INSERT INTO Students (StudentID, StudentName, DOB, Address, PhoneNumber) VALUES (2, "Woei Yih", "12/02/2001", "88 ALkaff Drive", "82223333");

SELECT * FROM Students

