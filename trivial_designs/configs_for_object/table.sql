create table if not exists object_configs(
      id bigserial not null primary key,
      object_id varchar(128) not null references "object"("id"),
      name varchar(64) not null,
      category varchar(32) not null default '',
      switch varchar(16) not null default 'on',
      content_type varchar(32) not null default 'json',
      content text not null,
      description text not null default '',
      created_date            timestamp with time zone                      DEFAULT now(),
      changed_date            timestamp with time zone                      DEFAULT now(),
      deleted_date            timestamp with time zone                      NULL
);


create unique index if not exists object_configs_object_id_name_unique_idx on object_configs(object_id, name);