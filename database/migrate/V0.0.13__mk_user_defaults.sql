CREATE TABLE public.users_defaults (
    -- Main defaults
    user_id INTEGER NOT NULL PRIMARY KEY,
    project_id INTEGER,
    individual_id INTEGER,
    lifestage TEXT,
    sex TEXT,
    body_length TEXT,
    created_at TEXT,
    language_id INTEGER,
    latitude NUMERIC,
    longitude NUMERIC,
    place_name TEXT,
    note TEXT,

    -- Classification defaults
    classification_id INTEGER,
    classification_species TEXT,
    classification_genus TEXT,
    classification_family TEXT,
    classification_order TEXT,
    classification_class TEXT,
    classification_phylum TEXT,
    classification_kingdom TEXT,
    classification_others TEXT,

    -- Observation defaults
    observation_id INTEGER,
    observation_user_id INTEGER,
    observation_method_id INTEGER,
    observation_method_name TEXT,
    observation_page_id INTEGER,
    observation_behavior TEXT,
    observation_observed_at TEXT,

    -- Specimen defaults
    specimen_method_id INTEGER,
    specimen_method_name TEXT,
    specimen_page_id INTEGER,

    -- Identification defaults
    identification_id INTEGER,
    identification_user_id INTEGER,
    identification_identified_at TEXT,
    identification_source_info TEXT,

    -- usersテーブルへの外部キー制約。ユーザーが削除されたら、このデフォルト設定も一緒に消えるのだ。
    CONSTRAINT fk_users
        FOREIGN KEY(user_id)
        REFERENCES public.users(user_id)
        ON DELETE CASCADE
);
