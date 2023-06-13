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

COMMENT ON TABLE documento.firma_electronica IS 'Tabla para almacenar los datos asociados a una firma electrónica';
COMMENT ON COLUMN documento.firma_electronica.id IS 'Hash que funcionará como pk de la tabla';
COMMENT ON COLUMN documento.firma_electronica.codigo_autenticidad IS 'Código de autenticidad generado por la firma electrónica';
COMMENT ON COLUMN documento.firma_electronica.llaves IS 'Llaves para verificar la firma electrónica';
COMMENT ON COLUMN documento.firma_electronica.firmantes IS 'Personas que aparecen como firmantes en la firma electrónica';
COMMENT ON COLUMN documento.firma_electronica.firma_encriptada IS 'Firma encriptada generada por el módulo de firma electrónica';
