
-- se agrega valor default para el campo workspace en la tabla tipo_documento  
ALTER TABLE documento.tipo_documento 
ALTER COLUMN workspace set DEFAULT '/desarrollo/pruebas/';

-- se agrega columna dominio_tipo_documento a la tabla tipo documento 
ALTER TABLE documento.tipo_documento
add COLUMN dominio_tipo_documento INTEGER default 1; 

-- object: fk_tipo_documento_dominio_tipo_documento | type: CONSTRAINT --
-- ALTER TABLE documento.tipo_documento DROP CONSTRAINT IF EXISTS fk_tipo_documento_dominio_tipo_documento CASCADE;
ALTER TABLE documento.tipo_documento ADD CONSTRAINT fk_tipo_documento_dominio_tipo_documento FOREIGN KEY (dominio_tipo_documento)
REFERENCES documento.dominio_tipo_documento (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --
