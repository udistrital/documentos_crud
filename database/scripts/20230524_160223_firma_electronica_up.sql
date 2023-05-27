CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS documento.firma_electronica (
	id uuid DEFAULT uuid_generate_v4 (),
	codigo_autenticidad TEXT NOT NULL,
	llaves JSON,
	firmantes JSON,
	firma_encriptada TEXT,
    activo BOOLEAN NOT NULL,
    fecha_creacion TIMESTAMP NOT NULL,
    fecha_modificacion TIMESTAMP NOT NULL,
    CONSTRAINT pk_firma_electronica PRIMARY KEY (id)
);
