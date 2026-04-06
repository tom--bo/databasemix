# File Summary

This file contains comprehensive PostgreSQL database information compiled for AI context analysis. It includes schema definitions, account configurations, system variables, and other database metadata consolidated into a single file for efficient processing.

**Database Type**: PostgreSQL  
**Database Version**: PostgreSQL 16.13 (Debian 16.13-1.pgdg13+1) on x86_64-pc-linux-gnu, compiled by gcc (Debian 14.2.0-19) 14.2.0, 64-bit

## File Structure

- Variables - PostgreSQL system variables and their current values
- Tables - Database tables with metadata and DDL definitions
- View Details - Database views with their definitions
- Stored Functions - User-defined functions with their definitions
- Stored Procedures - User-defined procedures with their definitions
- User Roles - Role definitions and assignments
- User Accounts - Database user accounts with privileges
- Extensions - Installed PostgreSQL extensions

# Variables

| Variable Name | Current Value | Source |
|---------------|---------------|--------|
| allow_in_place_tablespaces | off | default |
| allow_system_table_mods | off | default |
| application_name |  | default |
| archive_cleanup_command |  | default |
| archive_command | (disabled) | default |
| archive_library |  | default |
| archive_mode | off | default |
| archive_timeout | 0 s | default |
| array_nulls | on | default |
| authentication_timeout | 60 s | default |
| autovacuum | on | default |
| autovacuum_analyze_scale_factor | 0.1 | default |
| autovacuum_analyze_threshold | 50 | default |
| autovacuum_freeze_max_age | 200000000 | default |
| autovacuum_max_workers | 3 | default |
| autovacuum_multixact_freeze_max_age | 400000000 | default |
| autovacuum_naptime | 60 s | default |
| autovacuum_vacuum_cost_delay | 2 ms | default |
| autovacuum_vacuum_cost_limit | -1 | default |
| autovacuum_vacuum_insert_scale_factor | 0.2 | default |
| autovacuum_vacuum_insert_threshold | 1000 | default |
| autovacuum_vacuum_scale_factor | 0.2 | default |
| autovacuum_vacuum_threshold | 50 | default |
| autovacuum_work_mem | -1 kB | default |
| backend_flush_after | 0 8kB | default |
| backslash_quote | safe_encoding | default |
| backtrace_functions |  | default |
| bgwriter_delay | 200 ms | default |
| bgwriter_flush_after | 64 8kB | default |
| bgwriter_lru_maxpages | 100 | default |
| bgwriter_lru_multiplier | 2 | default |
| block_size | 8192 | default |
| bonjour | off | default |
| bonjour_name |  | default |
| bytea_output | hex | default |
| check_function_bodies | on | default |
| checkpoint_completion_target | 0.9 | default |
| checkpoint_flush_after | 32 8kB | default |
| checkpoint_timeout | 300 s | default |
| checkpoint_warning | 30 s | default |
| client_connection_check_interval | 0 ms | default |
| client_encoding | UTF8 | client |
| client_min_messages | notice | default |
| cluster_name |  | default |
| commit_delay | 0 | default |
| commit_siblings | 5 | default |
| compute_query_id | auto | default |
| config_file | /var/lib/postgresql/data/postgresql.conf | override |
| constraint_exclusion | partition | default |
| cpu_index_tuple_cost | 0.005 | default |
| cpu_operator_cost | 0.0025 | default |
| cpu_tuple_cost | 0.01 | default |
| createrole_self_grant |  | default |
| cursor_tuple_fraction | 0.1 | default |
| data_checksums | off | default |
| data_directory | /var/lib/postgresql/data | override |
| data_directory_mode | 0700 | default |
| data_sync_retry | off | default |
| DateStyle | ISO, MDY | client |
| db_user_namespace | off | default |
| deadlock_timeout | 1000 ms | default |
| debug_assertions | off | default |
| debug_discard_caches | 0 | default |
| debug_io_direct |  | default |
| debug_logical_replication_streaming | buffered | default |
| debug_parallel_query | off | default |
| debug_pretty_print | on | default |
| debug_print_parse | off | default |
| debug_print_plan | off | default |
| debug_print_rewritten | off | default |
| default_statistics_target | 100 | default |
| default_table_access_method | heap | default |
| default_tablespace |  | default |
| default_text_search_config | pg_catalog.english | configuration file |
| default_toast_compression | pglz | default |
| default_transaction_deferrable | off | default |
| default_transaction_isolation | read committed | default |
| default_transaction_read_only | off | default |
| dynamic_library_path | $libdir | default |
| dynamic_shared_memory_type | posix | configuration file |
| effective_cache_size | 524288 8kB | default |
| effective_io_concurrency | 1 | default |
| enable_async_append | on | default |
| enable_bitmapscan | on | default |
| enable_gathermerge | on | default |
| enable_hashagg | on | default |
| enable_hashjoin | on | default |
| enable_incremental_sort | on | default |
| enable_indexonlyscan | on | default |
| enable_indexscan | on | default |
| enable_material | on | default |
| enable_memoize | on | default |
| enable_mergejoin | on | default |
| enable_nestloop | on | default |
| enable_parallel_append | on | default |
| enable_parallel_hash | on | default |
| enable_partition_pruning | on | default |
| enable_partitionwise_aggregate | off | default |
| enable_partitionwise_join | off | default |
| enable_presorted_aggregate | on | default |
| enable_seqscan | on | default |
| enable_sort | on | default |
| enable_tidscan | on | default |
| escape_string_warning | on | default |
| event_source | PostgreSQL | default |
| exit_on_error | off | default |
| extension_destdir |  | default |
| external_pid_file |  | default |
| extra_float_digits | 1 | default |
| file_extend_method | posix_fallocate | default |
| from_collapse_limit | 8 | default |
| fsync | on | default |
| full_page_writes | on | default |
| geqo | on | default |
| geqo_effort | 5 | default |
| geqo_generations | 0 | default |
| geqo_pool_size | 0 | default |
| geqo_seed | 0 | default |
| geqo_selection_bias | 2 | default |
| geqo_threshold | 12 | default |
| gin_fuzzy_search_limit | 0 | default |
| gin_pending_list_limit | 4096 kB | default |
| gss_accept_delegation | off | default |
| hash_mem_multiplier | 2 | default |
| hba_file | /var/lib/postgresql/data/pg_hba.conf | override |
| hot_standby | on | default |
| hot_standby_feedback | off | default |
| huge_pages | try | default |
| huge_page_size | 0 kB | default |
| icu_validation_level | warning | default |
| ident_file | /var/lib/postgresql/data/pg_ident.conf | override |
| idle_in_transaction_session_timeout | 0 ms | default |
| idle_session_timeout | 0 ms | default |
| ignore_checksum_failure | off | default |
| ignore_invalid_pages | off | default |
| ignore_system_indexes | off | default |
| in_hot_standby | off | default |
| integer_datetimes | on | default |
| IntervalStyle | postgres | default |
| jit | on | default |
| jit_above_cost | 100000 | default |
| jit_debugging_support | off | default |
| jit_dump_bitcode | off | default |
| jit_expressions | on | default |
| jit_inline_above_cost | 500000 | default |
| jit_optimize_above_cost | 500000 | default |
| jit_profiling_support | off | default |
| jit_provider | llvmjit | default |
| jit_tuple_deforming | on | default |
| join_collapse_limit | 8 | default |
| krb_caseins_users | off | default |
| krb_server_keyfile | FILE:/etc/postgresql-common/krb5.keytab | default |
| lc_messages | en_US.utf8 | configuration file |
| lc_monetary | en_US.utf8 | configuration file |
| lc_numeric | en_US.utf8 | configuration file |
| lc_time | en_US.utf8 | configuration file |
| listen_addresses | * | configuration file |
| local_preload_libraries |  | default |
| lock_timeout | 0 ms | default |
| lo_compat_privileges | off | default |
| log_autovacuum_min_duration | 600000 ms | default |
| log_checkpoints | on | default |
| log_connections | off | default |
| log_destination | stderr | default |
| log_directory | log | default |
| log_disconnections | off | default |
| log_duration | off | default |
| log_error_verbosity | default | default |
| log_executor_stats | off | default |
| log_file_mode | 0600 | default |
| log_filename | postgresql-%Y-%m-%d_%H%M%S.log | default |
| logging_collector | off | default |
| log_hostname | off | default |
| logical_decoding_work_mem | 65536 kB | default |
| log_line_prefix | %m [%p]  | default |
| log_lock_waits | off | default |
| log_min_duration_sample | -1 ms | default |
| log_min_duration_statement | -1 ms | default |
| log_min_error_statement | error | default |
| log_min_messages | warning | default |
| log_parameter_max_length | -1 B | default |
| log_parameter_max_length_on_error | 0 B | default |
| log_parser_stats | off | default |
| log_planner_stats | off | default |
| log_recovery_conflict_waits | off | default |
| log_replication_commands | off | default |
| log_rotation_age | 1440 min | default |
| log_rotation_size | 10240 kB | default |
| log_startup_progress_interval | 10000 ms | default |
| log_statement | none | default |
| log_statement_sample_rate | 1 | default |
| log_statement_stats | off | default |
| log_temp_files | -1 kB | default |
| log_timezone | Etc/UTC | configuration file |
| log_transaction_sample_rate | 0 | default |
| log_truncate_on_rotation | off | default |
| maintenance_io_concurrency | 10 | default |
| maintenance_work_mem | 65536 kB | default |
| max_connections | 100 | configuration file |
| max_files_per_process | 1000 | default |
| max_function_args | 100 | default |
| max_identifier_length | 63 | default |
| max_index_keys | 32 | default |
| max_locks_per_transaction | 64 | default |
| max_logical_replication_workers | 4 | default |
| max_parallel_apply_workers_per_subscription | 2 | default |
| max_parallel_maintenance_workers | 2 | default |
| max_parallel_workers | 8 | default |
| max_parallel_workers_per_gather | 2 | default |
| max_pred_locks_per_page | 2 | default |
| max_pred_locks_per_relation | -2 | default |
| max_pred_locks_per_transaction | 64 | default |
| max_prepared_transactions | 0 | default |
| max_replication_slots | 10 | default |
| max_slot_wal_keep_size | -1 MB | default |
| max_stack_depth | 2048 kB | default |
| max_standby_archive_delay | 30000 ms | default |
| max_standby_streaming_delay | 30000 ms | default |
| max_sync_workers_per_subscription | 2 | default |
| max_wal_senders | 10 | default |
| max_wal_size | 1024 MB | configuration file |
| max_worker_processes | 8 | default |
| min_dynamic_shared_memory | 0 MB | default |
| min_parallel_index_scan_size | 64 8kB | default |
| min_parallel_table_scan_size | 1024 8kB | default |
| min_wal_size | 80 MB | configuration file |
| old_snapshot_threshold | -1 min | default |
| parallel_leader_participation | on | default |
| parallel_setup_cost | 1000 | default |
| parallel_tuple_cost | 0.1 | default |
| password_encryption | scram-sha-256 | default |
| plan_cache_mode | auto | default |
| port | 5432 | default |
| post_auth_delay | 0 s | default |
| pre_auth_delay | 0 s | default |
| primary_conninfo |  | default |
| primary_slot_name |  | default |
| quote_all_identifiers | off | default |
| random_page_cost | 4 | default |
| recovery_end_command |  | default |
| recovery_init_sync_method | fsync | default |
| recovery_min_apply_delay | 0 ms | default |
| recovery_prefetch | try | default |
| recovery_target |  | default |
| recovery_target_action | pause | default |
| recovery_target_inclusive | on | default |
| recovery_target_lsn |  | default |
| recovery_target_name |  | default |
| recovery_target_time |  | default |
| recovery_target_timeline | latest | default |
| recovery_target_xid |  | default |
| recursive_worktable_factor | 10 | default |
| remove_temp_files_after_crash | on | default |
| reserved_connections | 0 | default |
| restart_after_crash | on | default |
| restore_command |  | default |
| restrict_nonsystem_relation_kind |  | default |
| row_security | on | default |
| scram_iterations | 4096 | default |
| search_path | "$user", public | default |
| segment_size | 131072 8kB | default |
| send_abort_for_crash | off | default |
| send_abort_for_kill | off | default |
| seq_page_cost | 1 | default |
| server_encoding | UTF8 | default |
| server_version | 16.13 (Debian 16.13-1.pgdg13+1) | default |
| server_version_num | 160013 | default |
| session_preload_libraries |  | default |
| session_replication_role | origin | default |
| shared_buffers | 16384 8kB | configuration file |
| shared_memory_size | 143 MB | default |
| shared_memory_size_in_huge_pages | 72 | default |
| shared_memory_type | mmap | default |
| shared_preload_libraries |  | default |
| ssl | off | default |
| ssl_ca_file |  | default |
| ssl_cert_file | server.crt | default |
| ssl_ciphers | HIGH:MEDIUM:+3DES:!aNULL | default |
| ssl_crl_dir |  | default |
| ssl_crl_file |  | default |
| ssl_dh_params_file |  | default |
| ssl_ecdh_curve | prime256v1 | default |
| ssl_key_file | server.key | default |
| ssl_library | OpenSSL | default |
| ssl_max_protocol_version |  | default |
| ssl_min_protocol_version | TLSv1.2 | default |
| ssl_passphrase_command |  | default |
| ssl_passphrase_command_supports_reload | off | default |
| ssl_prefer_server_ciphers | on | default |
| standard_conforming_strings | on | default |
| statement_timeout | 0 ms | default |
| stats_fetch_consistency | cache | default |
| superuser_reserved_connections | 3 | default |
| synchronize_seqscans | on | default |
| synchronous_commit | on | default |
| synchronous_standby_names |  | default |
| syslog_facility | local0 | default |
| syslog_ident | postgres | default |
| syslog_sequence_numbers | on | default |
| syslog_split_messages | on | default |
| tcp_keepalives_count | 9 | default |
| tcp_keepalives_idle | 7200 s | default |
| tcp_keepalives_interval | 75 s | default |
| tcp_user_timeout | 0 ms | default |
| temp_buffers | 1024 8kB | default |
| temp_file_limit | -1 kB | default |
| temp_tablespaces |  | default |
| TimeZone | Etc/UTC | configuration file |
| timezone_abbreviations | Default | default |
| trace_notify | off | default |
| trace_recovery_messages | log | default |
| trace_sort | off | default |
| track_activities | on | default |
| track_activity_query_size | 1024 B | default |
| track_commit_timestamp | off | default |
| track_counts | on | default |
| track_functions | none | default |
| track_io_timing | off | default |
| track_wal_io_timing | off | default |
| transaction_deferrable | off | override |
| transaction_isolation | read committed | override |
| transaction_read_only | off | override |
| transform_null_equals | off | default |
| unix_socket_directories | /var/run/postgresql | default |
| unix_socket_group |  | default |
| unix_socket_permissions | 0777 | default |
| update_process_title | on | default |
| vacuum_buffer_usage_limit | 256 kB | default |
| vacuum_cost_delay | 0 ms | default |
| vacuum_cost_limit | 200 | default |
| vacuum_cost_page_dirty | 20 | default |
| vacuum_cost_page_hit | 1 | default |
| vacuum_cost_page_miss | 2 | default |
| vacuum_failsafe_age | 1600000000 | default |
| vacuum_freeze_min_age | 50000000 | default |
| vacuum_freeze_table_age | 150000000 | default |
| vacuum_multixact_failsafe_age | 1600000000 | default |
| vacuum_multixact_freeze_min_age | 5000000 | default |
| vacuum_multixact_freeze_table_age | 150000000 | default |
| wal_block_size | 8192 | default |
| wal_buffers | 512 8kB | default |
| wal_compression | off | default |
| wal_consistency_checking |  | default |
| wal_decode_buffer_size | 524288 B | default |
| wal_init_zero | on | default |
| wal_keep_size | 0 MB | default |
| wal_level | replica | default |
| wal_log_hints | off | default |
| wal_receiver_create_temp_slot | off | default |
| wal_receiver_status_interval | 10 s | default |
| wal_receiver_timeout | 60000 ms | default |
| wal_recycle | on | default |
| wal_retrieve_retry_interval | 5000 ms | default |
| wal_segment_size | 16777216 B | default |
| wal_sender_timeout | 60000 ms | default |
| wal_skip_threshold | 2048 kB | default |
| wal_sync_method | fdatasync | default |
| wal_writer_delay | 200 ms | default |
| wal_writer_flush_after | 128 8kB | default |
| work_mem | 4096 kB | default |
| xmlbinary | base64 | default |
| xmloption | content | default |
| zero_damaged_pages | off | default |

# Tables

# Tables

## public.data_types_test

- Row Format: 16 kB

```sql
CREATE TABLE public.data_types_test (
    id integer NOT NULL DEFAULT nextval('data_types_test_id_seq'::regclass),
    col_smallint smallint,
    col_integer integer,
    col_bigint bigint,
    col_numeric numeric(15,5),
    col_real real,
    col_double double precision,
    col_varchar varchar(255),
    col_char char(10),
    col_text text,
    col_boolean boolean,
    col_date date,
    col_time time without time zone,
    col_timestamp timestamp without time zone,
    col_timestamptz timestamp with time zone,
    col_interval interval,
    col_uuid uuid,
    col_json json,
    col_jsonb jsonb,
    col_bytea bytea,
    col_inet inet,
    col_cidr cidr,
    col_macaddr macaddr,
    col_int_array _int4,
    col_text_array _text,
    CONSTRAINT data_types_test_pkey PRIMARY KEY (id)
);
```

## public.logs

- Row Format: 0 bytes

```sql
CREATE TABLE public.logs (
    id integer NOT NULL DEFAULT nextval('logs_id_seq'::regclass),
    log_date date NOT NULL,
    message text,
    level varchar(20)
);
```

## public.logs_2024

- Row Format: 16 kB

```sql
CREATE TABLE public.logs_2024 (
    id integer NOT NULL DEFAULT nextval('logs_id_seq'::regclass),
    log_date date NOT NULL,
    message text,
    level varchar(20)
);
```

## public.logs_2025

- Row Format: 16 kB

```sql
CREATE TABLE public.logs_2025 (
    id integer NOT NULL DEFAULT nextval('logs_id_seq'::regclass),
    log_date date NOT NULL,
    message text,
    level varchar(20)
);
```

## public.order_items

- Row Format: 56 kB

```sql
CREATE TABLE public.order_items (
    id integer NOT NULL DEFAULT nextval('order_items_id_seq'::regclass),
    order_id integer NOT NULL,
    product_id integer NOT NULL,
    quantity integer NOT NULL DEFAULT 1,
    unit_price numeric(10,2) NOT NULL,
    CONSTRAINT order_items_order_id_fkey FOREIGN KEY (order_id) REFERENCES orders(id),
    CONSTRAINT order_items_product_id_fkey FOREIGN KEY (product_id) REFERENCES products(id),
    CONSTRAINT order_items_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_order_items_order_id ON public.order_items USING btree (order_id);
CREATE INDEX idx_order_items_product_id ON public.order_items USING btree (product_id);
```

## public.orders

- Row Format: 64 kB

```sql
CREATE TABLE public.orders (
    id integer NOT NULL DEFAULT nextval('orders_id_seq'::regclass),
    user_id integer NOT NULL,
    order_date timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    total_amount numeric(10,2) NOT NULL DEFAULT 0.00,
    status order_status DEFAULT 'pending'::order_status,
    notes text,
    CONSTRAINT orders_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT orders_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_orders_user_id ON public.orders USING btree (user_id);
CREATE INDEX idx_orders_status ON public.orders USING btree (status);
```

## public.products

- Row Format: 64 kB

```sql
CREATE TABLE public.products (
    id integer NOT NULL DEFAULT nextval('products_id_seq'::regclass),
    name varchar(200) NOT NULL,
    description text,
    price numeric(10,2) NOT NULL DEFAULT 0.00,
    stock_quantity integer NOT NULL DEFAULT 0,
    category varchar(50),
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT products_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_products_category ON public.products USING btree (category);
CREATE INDEX idx_products_price ON public.products USING btree (price);
```

## public.users

- Row Format: 80 kB

```sql
CREATE TABLE public.users (
    id integer NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    username varchar(50) NOT NULL,
    email varchar(100) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    is_active boolean DEFAULT true,
    metadata jsonb DEFAULT '{}'::jsonb,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_username_key UNIQUE (username)
);
CREATE INDEX idx_users_email ON public.users USING btree (email);
CREATE INDEX idx_users_active ON public.users USING btree (is_active);
```

# View info details

## public.active_users

```sql
CREATE VIEW public.active_users AS
 SELECT id,
    username,
    email,
    created_at
   FROM users
  WHERE is_active = true;
```

## public.order_summary

```sql
CREATE VIEW public.order_summary AS
 SELECT u.username,
    count(o.id) AS total_orders,
    COALESCE(sum(o.total_amount), 0::numeric) AS total_spent,
    max(o.order_date) AS last_order_date
   FROM users u
     LEFT JOIN orders o ON u.id = o.user_id
  GROUP BY u.username;
```

## public.product_inventory

```sql
CREATE VIEW public.product_inventory AS
 SELECT name,
    stock_quantity,
    price,
        CASE
            WHEN stock_quantity = 0 THEN 'Out of Stock'::text
            WHEN stock_quantity < 10 THEN 'Low Stock'::text
            ELSE 'In Stock'::text
        END AS stock_status
   FROM products;
```

# Stored Functions

## public.calculate_user_total

- Specific Name: calculate_user_total
- Routine Catalog: def
- Character Set: utf8mb4
- Collation: utf8mb4_unicode_ci
- Routine Body: SQL
- External Language: SQL
- Parameter Style: SQL
- Is Deterministic: YES
- SQL Data Access: STABLE
- Security Type: INVOKER
- SQL Mode: ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION

```sql
CREATE OR REPLACE FUNCTION public.calculate_user_total(p_user_id integer)
 RETURNS numeric
 LANGUAGE plpgsql
 STABLE
AS $function$
DECLARE
    total NUMERIC(10,2);
BEGIN
    SELECT COALESCE(SUM(total_amount), 0) INTO total
    FROM orders
    WHERE user_id = p_user_id;
    RETURN total;
END;
$function$

```

## public.format_currency

- Specific Name: format_currency
- Routine Catalog: def
- Character Set: utf8mb4
- Collation: utf8mb4_unicode_ci
- Routine Body: SQL
- External Language: SQL
- Parameter Style: SQL
- Is Deterministic: YES
- SQL Data Access: IMMUTABLE
- Security Type: INVOKER
- SQL Mode: ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION

```sql
CREATE OR REPLACE FUNCTION public.format_currency(amount numeric)
 RETURNS text
 LANGUAGE plpgsql
 IMMUTABLE
AS $function$
BEGIN
    RETURN '$' || TO_CHAR(amount, 'FM999,999,990.00');
END;
$function$

```

## public.get_product_availability

- Specific Name: get_product_availability
- Routine Catalog: def
- Character Set: utf8mb4
- Collation: utf8mb4_unicode_ci
- Routine Body: SQL
- External Language: SQL
- Parameter Style: SQL
- Is Deterministic: YES
- SQL Data Access: STABLE
- Security Type: INVOKER
- SQL Mode: ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION

```sql
CREATE OR REPLACE FUNCTION public.get_product_availability(p_product_id integer)
 RETURNS text
 LANGUAGE plpgsql
 STABLE
AS $function$
DECLARE
    qty INTEGER;
BEGIN
    SELECT stock_quantity INTO qty
    FROM products
    WHERE id = p_product_id;

    IF qty IS NULL THEN
        RETURN 'Not Found';
    ELSIF qty = 0 THEN
        RETURN 'Out of Stock';
    ELSIF qty < 10 THEN
        RETURN 'Low Stock';
    ELSE
        RETURN 'In Stock';
    END IF;
END;
$function$

```

## public.update_timestamp

- Specific Name: update_timestamp
- Routine Catalog: def
- Character Set: utf8mb4
- Collation: utf8mb4_unicode_ci
- Routine Body: SQL
- External Language: SQL
- Parameter Style: SQL
- Is Deterministic: YES
- SQL Data Access: VOLATILE
- Security Type: INVOKER
- SQL Mode: ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION

```sql
CREATE OR REPLACE FUNCTION public.update_timestamp()
 RETURNS trigger
 LANGUAGE plpgsql
AS $function$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$function$

```

# Stored Procedures

## public.generate_sales_report

- Specific Name: generate_sales_report
- Routine Catalog: def
- Routine Body: SQL
- External Language: SQL
- Parameter Style: SQL
- Is Deterministic: YES
- SQL Data Access: VOLATILE
- Security Type: INVOKER
- SQL Mode: ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION

```sql
CREATE OR REPLACE PROCEDURE public.generate_sales_report(IN p_start_date date, IN p_end_date date)
 LANGUAGE plpgsql
AS $procedure$
DECLARE
    rec RECORD;
BEGIN
    FOR rec IN
        SELECT order_date::date as sale_date,
               COUNT(*) as order_count,
               SUM(total_amount) as daily_total
        FROM orders
        WHERE order_date >= p_start_date AND order_date < p_end_date
        GROUP BY order_date::date
        ORDER BY order_date::date
    LOOP
        RAISE NOTICE 'Date: %, Orders: %, Total: %', rec.sale_date, rec.order_count, rec.daily_total;
    END LOOP;
END;
$procedure$

```

## public.update_product_stock

- Specific Name: update_product_stock
- Routine Catalog: def
- Routine Body: SQL
- External Language: SQL
- Parameter Style: SQL
- Is Deterministic: YES
- SQL Data Access: VOLATILE
- Security Type: INVOKER
- SQL Mode: ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION

```sql
CREATE OR REPLACE PROCEDURE public.update_product_stock(IN p_product_id integer, IN p_quantity integer)
 LANGUAGE plpgsql
AS $procedure$
BEGIN
    UPDATE products
    SET stock_quantity = stock_quantity + p_quantity
    WHERE id = p_product_id;

    IF NOT FOUND THEN
        RAISE EXCEPTION 'Product % not found', p_product_id;
    END IF;
END;
$procedure$

```

# User Roles

## analyst@

## app_admin@

- **Default Role for some users**

## app_read@

- **Default Role for some users**

## app_write@

- **Default Role for some users**

## developer@

- **Default Role for some users**

# User List

## admin

- Attributes: CREATEDB, CREATEROLE
- Grants:
  - MEMBER OF developer
  - DATABASE testdb: {=Tc/postgres,postgres=CTc/postgres,testuser=CTc/postgres,readonly=c/postgres,admin=CTc/postgres}

## postgres

- Attributes: SUPERUSER, CREATEDB, CREATEROLE, REPLICATION
- Grants:
  - DATABASE template1: {=c/postgres,postgres=CTc/postgres}
  - DATABASE template0: {=c/postgres,postgres=CTc/postgres}
  - DATABASE testdb: {=Tc/postgres,postgres=CTc/postgres,testuser=CTc/postgres,readonly=c/postgres,admin=CTc/postgres}

## readonly

- Grants:
  - MEMBER OF app_read
  - DATABASE testdb: {=Tc/postgres,postgres=CTc/postgres,testuser=CTc/postgres,readonly=c/postgres,admin=CTc/postgres}

## testuser

- Grants:
  - MEMBER OF app_write
  - DATABASE testdb: {=Tc/postgres,postgres=CTc/postgres,testuser=CTc/postgres,readonly=c/postgres,admin=CTc/postgres}

# Extensions

| Name | Version | Description |
|------|---------|-------------|
| plpgsql | 1.0 | PL/pgSQL procedural language |

