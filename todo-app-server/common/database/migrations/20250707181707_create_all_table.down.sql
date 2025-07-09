DO $$ DECLARE
r RECORD;
BEGIN
FOR r IN (SELECT tablename FROM pg_tables WHERE schemaname = 'public') LOOP
        EXECUTE 'DROP TABLE IF EXISTS public.' || quote_ident(r.tablename) || ' CASCADE';
END LOOP;
END $$;

DROP TRIGGER IF EXISTS trigger_update_user_updated_at ON users;
DROP TRIGGER IF EXISTS trigger_update_todo_updated_at ON todos;

DROP FUNCTION IF EXISTS update_updated_at_column;

DROP TYPE IF EXISTS tag_enum;

DROP INDEX IF EXISTS idx_user_todos_id;
DROP INDEX IF EXISTS idx_todos_tag_id;
DROP INDEX IF EXISTS idx_todos_title_id;