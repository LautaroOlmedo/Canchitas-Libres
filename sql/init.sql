-- Crear la base de datos si no existe (esto es válido solo en psql)
SELECT 'CREATE DATABASE canchitas'
    WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'canchitas')\gexec

-- Conectar a la base de datos
    \c canchitas

-- Creación del esquema para la base de datos
CREATE SCHEMA IF NOT EXISTS canchitas;

-- Tabla Person
CREATE TABLE IF NOT EXISTS canchitas.Person
    PersonID SERIAL PRIMARY KEY,
    Firstname VARCHAR(50) NOT NULL,
    Lastname VARCHAR(50) NOT NULL,
    DNI VARCHAR(20) UNIQUE NOT NULL,
    Birthday DATE NOT NULL
);

-- Tabla User
CREATE TABLE IF NOT EXISTS canchitas.User (
    UserID SERIAL PRIMARY KEY,
    Role VARCHAR(20) NOT NULL,
    Password VARCHAR(100) NOT NULL,
    Active BOOLEAN NOT NULL,
    PersonID INT NOT NULL UNIQUE,
    CONSTRAINT fk_person FOREIGN KEY (PersonID) REFERENCES canchitas.Person(PersonID)
);

-- Tabla Phone
CREATE TABLE IF NOT EXISTS canchitas.Phone (
    PhoneID SERIAL PRIMARY KEY,
    UserID INT NOT NULL,
    Phone VARCHAR(15) NOT NULL,
    CONSTRAINT fk_user_phone FOREIGN KEY (UserID) REFERENCES canchitas.User(UserID)
);

-- Tabla Email
CREATE TABLE IF NOT EXISTS canchitas.Email (
    EmailID SERIAL PRIMARY KEY,
    UserID INT NOT NULL,
    Email VARCHAR(100) UNIQUE NOT NULL,
    CONSTRAINT fk_user_email FOREIGN KEY (UserID) REFERENCES canchitas.User(UserID)
);

-- Tabla Field
CREATE TABLE IF NOT EXISTS canchitas.Field (
    FieldID SERIAL PRIMARY KEY,
    FieldType VARCHAR(50) NOT NULL,
    Status VARCHAR(20) NOT NULL,
    FieldNumber INT NOT NULL UNIQUE,
    FieldPrice DECIMAL(10, 2) NOT NULL
);

-- Tabla Reserve
CREATE TABLE IF NOT EXISTS canchitas.Reserve (
        ReserveID SERIAL PRIMARY KEY,
        FieldID INT NOT NULL,
        UserID INT NOT NULL,
        Date TIMESTAMP NOT NULL,
        ReserveNumber VARCHAR(20) UNIQUE NOT NULL,
    CONSTRAINT fk_field FOREIGN KEY (FieldID) REFERENCES canchitas.Field(FieldID),
    CONSTRAINT fk_user_reserve FOREIGN KEY (UserID) REFERENCES canchitas.User(UserID)
);

-- Tabla Payment
CREATE TABLE IF NOT EXISTS canchitas.Payment (
    PaymentID SERIAL PRIMARY KEY,
    Amount DECIMAL(10, 2) NOT NULL,
    ReserveID INT NOT NULL,
    UserID INT NOT NULL,
    Status VARCHAR(20) NOT NULL,
    CONSTRAINT fk_reserve FOREIGN KEY (ReserveID) REFERENCES canchitas.Reserve(ReserveID),
    CONSTRAINT fk_user_payment FOREIGN KEY (UserID) REFERENCES canchitas.User(UserID)
);

-- Tabla Transaction
CREATE TABLE IF NOT EXISTS canchitas.Transaction (
        TransactionID SERIAL PRIMARY KEY,
        PaymentID INT NOT NULL,
        CONSTRAINT fk_payment FOREIGN KEY (PaymentID) REFERENCES canchitas.Payment(PaymentID)
);
