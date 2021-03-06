

SET search_path TO pg_catalog,public,documento;


-- object: documento.dominio_tipo_documento | type: TABLE --
-- DROP TABLE IF EXISTS documento.dominio_tipo_documento CASCADE;
CREATE TABLE documento.dominio_tipo_documento (
	id serial NOT NULL,
	nombre character varying(150) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),	
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	CONSTRAINT pk_dominio_tipo_documento PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE documento.dominio_tipo_documento IS 'Tabla que almacena los dominios que abarcan diferentes tipos de documento';
-- ddl-end --
COMMENT ON COLUMN documento.dominio_tipo_documento.id IS 'Identificador del dominio';
-- ddl-end --
COMMENT ON COLUMN documento.dominio_tipo_documento.nombre IS 'Título del dominio';
-- ddl-end --
COMMENT ON COLUMN documento.dominio_tipo_documento.descripcion IS 'Descripción del dominio';
-- ddl-end --
COMMENT ON COLUMN documento.dominio_tipo_documento.codigo_abreviacion IS 'Código de abreviación, sigla, acrónimo u otra representación corta del registro si se requiere';
-- ddl-end --
COMMENT ON COLUMN documento.dominio_tipo_documento.activo IS 'Campo para indicar el estado del registro';
-- ddl-end --
COMMENT ON COLUMN documento.dominio_tipo_documento.numero_orden IS 'En dado caso que se necesite establecer un orden a los registros que no se encuentre definido por aplicación ni por BD. Allí se almacena permitiendo realizar subitems mediante la precisión.';
-- ddl-end --
COMMENT ON COLUMN documento.dominio_tipo_documento.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN documento.dominio_tipo_documento.fecha_modificacion IS 'Fecha de la última modificación del registro';
-- ddl-end --
COMMENT ON CONSTRAINT pk_dominio_tipo_documento ON documento.dominio_tipo_documento  IS 'Llave primaria de la tabla dominio_tipo_documento';
-- ddl-end --



-- object: documento.subtipo_documento | type: TABLE --
-- DROP TABLE IF EXISTS documento.subtipo_documento CASCADE;
CREATE TABLE documento.subtipo_documento (
	id serial NOT NULL,
	tipo_documento_padre integer NOT NULL,
	tipo_documento_hijo integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL DEFAULT now(),
	fecha_modificacion timestamp NOT NULL DEFAULT now(),
	CONSTRAINT pk_subtipo_documento PRIMARY KEY (id)
);
COMMENT ON TABLE documento.subtipo_documento IS 'Tabla paramétrica para identificar las relaciones de jerarquia entre los tipos de documento';
-- ddl-end --
COMMENT ON COLUMN documento.subtipo_documento.id IS 'Identificador de la tabla subtipo_documento';
-- ddl-end --
COMMENT ON COLUMN documento.subtipo_documento.id IS 'Identificador de la tabla subtipo_documento';
-- ddl-end --
COMMENT ON COLUMN documento.subtipo_documento.tipo_documento_padre  IS 'Referencia foránea a la tabla tipo_documento que vincula el tipo de documento padre';
-- ddl-end --
COMMENT ON COLUMN documento.subtipo_documento.tipo_documento_hijo  IS 'Referencia foránea a la tabla tipo_documento que vincula el tipo de documento hijo';
-- ddl-end --
COMMENT ON COLUMN documento.subtipo_documento.activo IS 'Campo que indica si el parámetro está activo';
-- ddl-end --
COMMENT ON COLUMN documento.subtipo_documento.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN documento.subtipo_documento.fecha_modificacion IS 'Fecha de la última modificación del registro';
-- ddl-end --
COMMENT ON CONSTRAINT pk_subtipo_documento ON documento.subtipo_documento  IS 'Llave primaria de la tabla subtipo_documento';
-- ddl-end --



-- object: fk_subtipo_documento_tipo_documento_padre | type: CONSTRAINT --
-- ALTER TABLE documento.subtipo_documento DROP CONSTRAINT IF EXISTS fk_subtipo_documento_tipo_documento_padre CASCADE;
ALTER TABLE documento.subtipo_documento ADD CONSTRAINT fk_subtipo_documento_tipo_documento_padre FOREIGN KEY (tipo_documento_padre)
REFERENCES documento.tipo_documento (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_subtipo_documento_tipo_documento_hijo | type: CONSTRAINT --
-- ALTER TABLE documento.subtipo_documento DROP CONSTRAINT IF EXISTS fk_subtipo_documento_tipo_documento_hijo CASCADE;
ALTER TABLE documento.subtipo_documento ADD CONSTRAINT fk_subtipo_documento_tipo_documento_hijo FOREIGN KEY (tipo_documento_hijo)
REFERENCES documento.tipo_documento (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;

-- Permisos de usuario
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA documento TO desarrollooas;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA documento TO desarrollooas;
