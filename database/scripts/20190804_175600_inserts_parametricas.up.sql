-- inserts documento.dominio_tipo_documento
INSERT INTO documento.dominio_tipo_documento(nombre,descripcion,codigo_abreviacion,activo,numero_orden) VALUES ('FINANCIERO','Documento que corresponde a negocios financieros','DD-FINA',TRUE,1.00);
INSERT INTO documento.dominio_tipo_documento(nombre,descripcion,codigo_abreviacion,activo,numero_orden) VALUES ('TITAN','Documentos de soporte para realizar pago de nómina mensual','DT-SOPORTE',TRUE,2.00);
INSERT INTO documento.dominio_tipo_documento(nombre,descripcion,codigo_abreviacion,activo,numero_orden) VALUES ('RESOLUCIONES','Documentos generados por la aplicación Resoluciones','RES-DVE',TRUE,3.00);
INSERT INTO documento.dominio_tipo_documento(nombre,descripcion,activo) VALUES ('DEFAULT','por definir',TRUE);
-- inserts documento.tipo_documento
INSERT INTO documento.tipo_documento(codigo_abreviacion, nombre, descripcion,activo,numero_orden,dominio_tipo_documento) VALUES ('TD-ORPA','ORDEN PAGO','Documentos Financiero para Ordenes de Pago',TRUE,1.00,1);
INSERT INTO documento.tipo_documento(codigo_abreviacion, nombre, descripcion,activo,numero_orden,dominio_tipo_documento) VALUES ('ACT','ACTA','Documentos para registro de fuentes',TRUE,1.01,1);
INSERT INTO documento.tipo_documento(codigo_abreviacion, nombre, descripcion,activo,numero_orden,dominio_tipo_documento) VALUES ('SPTE','SOPORTE','Documentos de soporte para la gestión de pago ',TRUE,1.02,2);
INSERT INTO documento.tipo_documento(codigo_abreviacion, nombre, descripcion,activo,numero_orden,dominio_tipo_documento) VALUES ('TD-ING','INGRESOS','Documentos que aplican para ingresos',TRUE,1.03,1);
INSERT INTO documento.tipo_documento(codigo_abreviacion, nombre, descripcion,activo,numero_orden,dominio_tipo_documento) VALUES ('TD-DEV','DEVOLUCIONES','Documentos que aplican para Devoluciones',TRUE,1.04,1);
INSERT INTO documento.tipo_documento(codigo_abreviacion, nombre, descripcion,activo,numero_orden,dominio_tipo_documento) VALUES ('RES-DVE','RESOLUCIONES','Resoluciones de Vinculación Especial',TRUE,1.05,3);

INSERT INTO documento.tipo_documento(nombre, descripcion, codigo_abreviacion, activo, numero_orden,tamano,extension,workspace,tipo_documento_nuxeo) VALUES ('Foto inscripción','Foto inscripción','FI', true,1,5,'png,jpg','/default-domain/workspaces/Pruebas Planestic/Inscripcion/Fotos','Picture');
INSERT INTO documento.tipo_documento(nombre, descripcion, codigo_abreviacion, activo, numero_orden,tamano,extension,workspace,tipo_documento_nuxeo) VALUES ('Documentos','Documento identificación','DI', true,2,5,'pdf','/default-domain/workspaces/Pruebas Planestic/Inscripcion/Documentos','File');
INSERT INTO documento.tipo_documento(nombre, descripcion, codigo_abreviacion, activo, numero_orden,tamano,extension,workspace,tipo_documento_nuxeo) VALUES ('Formación académica','Formación académica','FA', true,3,5,'pdf','/default-domain/workspaces/Pruebas Planestic/Inscripcion/FormacionAcademica','File');
INSERT INTO documento.tipo_documento(nombre, descripcion, codigo_abreviacion, activo, numero_orden,tamano,extension,workspace,tipo_documento_nuxeo) VALUES ('Experiencia laboral','Experiencia laboral','EL', true,4,5,'pdf','/default-domain/workspaces/Pruebas Planestic/Inscripcion/ExperienciaLaboral','File');
INSERT INTO documento.tipo_documento(nombre, descripcion, codigo_abreviacion, activo, numero_orden,tamano,extension,workspace,tipo_documento_nuxeo) VALUES ('Propuesta grado','Propuesta grado','PG', true,5,5,'pdf','/default-domain/workspaces/Pruebas Planestic/Inscripcion/PropuestaGrado','File');
INSERT INTO documento.tipo_documento(nombre, descripcion, codigo_abreviacion, activo, numero_orden,tamano,extension,workspace,tipo_documento_nuxeo) VALUES ('Documento soporte','Documento soporte','DS', true,6,5,'pdf','/default-domain/workspaces/Pruebas Planestic/Inscripcion/DocumentosSoporte','File');
INSERT INTO documento.tipo_documento(nombre, descripcion, codigo_abreviacion, activo, numero_orden,tamano,extension,workspace,tipo_documento_nuxeo) VALUES ('Descuentos matrícula','Descuentos matrícula','DM', true,7,5,'pdf','/default-domain/workspaces/Pruebas Planestic/Inscripcion/Descuentos','File');
INSERT INTO documento.tipo_documento(nombre, descripcion, codigo_abreviacion, activo, numero_orden,tamano,extension,workspace,tipo_documento_nuxeo) VALUES ('Comprobante pago','Comprobante pago','CPO', true,8,5,'pdf','/default-domain/workspaces/Pruebas Planestic/Inscripcion/ComprobantePago','File');
INSERT INTO documento.tipo_documento(nombre, descripcion, codigo_abreviacion, activo, numero_orden,tamano,extension,workspace,tipo_documento_nuxeo) VALUES ('Resolución registro calificado','Documento de resolución asociada a un registro calificado','RRC', true,9,5,'pdf','/desarrollo/workspaces/sga/Proyectos/Soportes','File');
INSERT INTO documento.tipo_documento(nombre, descripcion, codigo_abreviacion, activo, numero_orden,tamano,extension,workspace,tipo_documento_nuxeo) VALUES ('Acto adminnistrativo proyecto','Documento de acto administrativo por el que se crea un proyecto academico','AAP', true,10,5,'pdf','/desarrollo/workspaces/sga/Proyectos/Soportes','File');
INSERT INTO documento.tipo_documento(nombre, descripcion, codigo_abreviacion, activo, numero_orden,tamano,extension,workspace,tipo_documento_nuxeo) VALUES ('Resolución asignación coordinador','Documento de acto administrativo por el que se asigna un nuevo coordinador a un proyecto','RAC', true,11,5,'pdf','/desarrollo/workspaces/sga/Proyectos/Coordinacion','File');
