-- inserts documento.dominio_tipo_documento
INSERT INTO documento.dominio_tipo_documento(nombre,descripcion,activo) VALUES ('DEFAULT','por definir',TRUE);
INSERT INTO documento.dominio_tipo_documento(nombre,descripcion,codigo_abreviacion,activo,numero_orden) VALUES ('FINANCIERO','Documento que corresponde a negocios financieros','DD-FINA',TRUE,1.00);
INSERT INTO documento.dominio_tipo_documento(nombre,descripcion,codigo_abreviacion,activo,numero_orden) VALUES ('TITAN','Documentos de soporte para realizar pago de nómina mensual','DT-SOPORTE',TRUE,2.00);
INSERT INTO documento.dominio_tipo_documento(nombre,descripcion,codigo_abreviacion,activo,numero_orden) VALUES ('RESOLUCIONES','Documentos generados por la aplicación Resoluciones','RES-DVE',TRUE,3.00);

