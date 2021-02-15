ALTER TABLE documento.tipo_documento 
drop CONSTRAINT if EXISTS fk_tipo_documento_dominio_tipo_documento CASCADE; 

ALTER TABLE documento.tipo_documento 
ALTER COLUMN workspace DROP DEFAULT;

ALTER TABLE documento.tipo_documento 
DROP COLUMN if EXISTS dominio_tipo_documento;