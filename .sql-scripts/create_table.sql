CREATE TABLE account 
  ( 
     id       INT(11) NOT NULL AUTO_INCREMENT, 
     document VARCHAR(14) NOT NULL, 
     PRIMARY KEY (id), 
     CONSTRAINT document_unique UNIQUE (document)
  ); 

CREATE TABLE operation_type 
  ( 
     id          INT(11) NOT NULL AUTO_INCREMENT,
     description VARCHAR(40) NOT NULL, 
     PRIMARY KEY (id) 
  ); 

CREATE TABLE transaction 
  ( 
     id                INT(11) NOT NULL AUTO_INCREMENT,
     account_id        INT(11) NOT NULL, 
     operation_type_id INT(11) NOT NULL, 
     amount            DECIMAL(15, 2) NOT NULL, 
     created_at        TIMESTAMP NOT NULL, 
     PRIMARY KEY (id), 
     FOREIGN KEY (account_id) REFERENCES account(id),
     FOREIGN KEY (operation_type_id) REFERENCES operation_type(id)
  ); 