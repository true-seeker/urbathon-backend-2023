BEGIN;
ALTER TABLE appeal_themes
    DROP COLUMN appeal_category_id;
DROP TABLE appeal_categories;
ALTER TABLE appeal_types
    RENAME COLUMN appeal_theme_id TO appeal_category_id;
ALTER TABLE appeal_themes
    RENAME TO appeal_categories;
COMMIT;

