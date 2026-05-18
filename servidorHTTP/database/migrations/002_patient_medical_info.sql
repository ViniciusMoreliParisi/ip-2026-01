-- Run if patients table already exists without medical fields (PostgreSQL)

ALTER TABLE patients ADD COLUMN IF NOT EXISTS cpf VARCHAR(11);
ALTER TABLE patients ADD COLUMN IF NOT EXISTS blood_type VARCHAR(5);
ALTER TABLE patients ADD COLUMN IF NOT EXISTS allergies TEXT;
ALTER TABLE patients ADD COLUMN IF NOT EXISTS diseases TEXT;
ALTER TABLE patients ADD COLUMN IF NOT EXISTS medicine_usage TEXT;

-- Backfill only when upgrading old rows (adjust manually if needed)
UPDATE patients SET cpf = '' WHERE cpf IS NULL;
UPDATE patients SET blood_type = 'O+' WHERE blood_type IS NULL;
UPDATE patients SET allergies = '' WHERE allergies IS NULL;
UPDATE patients SET diseases = '' WHERE diseases IS NULL;
UPDATE patients SET medicine_usage = '' WHERE medicine_usage IS NULL;

ALTER TABLE patients ALTER COLUMN cpf SET NOT NULL;
ALTER TABLE patients ALTER COLUMN blood_type SET NOT NULL;
ALTER TABLE patients ALTER COLUMN allergies SET DEFAULT '';
ALTER TABLE patients ALTER COLUMN diseases SET DEFAULT '';
ALTER TABLE patients ALTER COLUMN medicine_usage SET DEFAULT '';

CREATE UNIQUE INDEX IF NOT EXISTS idx_patients_cpf ON patients (cpf);

ALTER TABLE patients DROP CONSTRAINT IF EXISTS chk_blood_type;
ALTER TABLE patients ADD CONSTRAINT chk_blood_type
    CHECK (blood_type IN ('A+', 'A-', 'B+', 'B-', 'AB+', 'AB-', 'O+', 'O-'));
