ALTER TABLE documento.firma_electronica
ADD COLUMN IF NOT EXISTS documento_id INTEGER NOT NULL;

ALTER TABLE documento.firma_electronica
    ADD CONSTRAINT fk_documento_firma_electronica FOREIGN KEY (documento_id)
    REFERENCES documento.documento (id) MATCH FULL
    ON DELETE RESTRICT ON UPDATE CASCADE;

COMMENT ON COLUMN documento.firma_electronica.documento_id IS 'Id del documento al que corresponde la firma.';
