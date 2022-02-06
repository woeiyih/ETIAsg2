# ETIAsg2

Author: Toh Woei Yih

# 1. Introduction
This project was created as part of an assignment for the Emerging Trends in IT (ETI) module in my diploma course here in Ngee Ann Polytechnic. It involves working with fellow peers to design, implement and containerise a number of microservices and REST APIs to bring alive a simulated version of an online institution learning management system - we called it EduFi. This online platform was designed with many different features, and the set of features that I am assigned to is the management of student account system, as described below.

3.2 Management of Student Accounts
  3.2.1.	Create, view, update, delete student accounts. Info includes
    3.2.1.1.	Student ID
    3.2.1.2.	Name
    3.2.1.3.	Date of Birth
    3.2.1.4.	Address
   3.2.1.5.	Phone number
  3.2.2.	List students
  3.2.3.	Search for students
  3.2.4.	View student’s modules, results, timetable
 
 # 2. Design Consideration of Mircoservices
 
 From all the requirements stated above, this microservice requires a front-end application to allow users to view student's information. 
 
 To implement the functions stated above, there needs to be a REST API developed that contains the information of the students. From the requirements, there needs to be at least four types of API request types: GET, PUT, POST and DELETE. Although there is no requirement to rely on another microservice to perform the functions, you still require to request API to obtain the information needed for student modules, results and timetable.
 
 Last but not least, a database has to be set up in order to provide a persistent platform where the student account information can be stored. It is recommended to use MySQL Workbench to host the database based on the way I have set it up. The structure of the DB is simple and can be found below.
 
 StudentID VARCHAR (30) NOT NULL PRIMARY KEY, StudentName VARCHAR (30), DOB VARCHAR(30), Address VARCHAR (30), PhoneNumber VARCHAR (30)
 
 # 3. Architecture Diagram
 This microservice's structure is made up of a front-end application, the student microservice and a student database. It is very simple and a diagram will be appended below very soon.
 
 # 4. Set-Up Instructions
 Step 1. On the server, pull the three docker images for the project from DockerHub. 
 
   docker pull woeiyih/asg2
 
   docker pull woeiyih/woeiyih-students-database
 
   docker pull woeiyih/woeiyih-students-frontend
 
 Step 2. Run the containers.
 
  docker run -d woeiyih/woeiyih-students-database
 
   docker run -d woeiyih/asg2
 
   docker run -d woeiyih/woeiyih-students-frontend
 
 Step 3. The Front-End Page should be accessbile on http://10.31.11.11:8150/
 
 
