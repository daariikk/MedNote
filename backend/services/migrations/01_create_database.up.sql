DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'mednote') THEN
        CREATE DATABASE "MedNote";
    END IF;
END $$;