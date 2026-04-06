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

| Variable Name | Current Value |
|---------------|---------------|
| allow_in_place_tablespaces | off |
| allow_system_table_mods | off |
| application_name |  |
| archive_cleanup_command |  |
| archive_command | (disabled) |
| archive_library |  |
| archive_mode | off |
| archive_timeout | 0 s |
| array_nulls | on |
| authentication_timeout | 60 s |
| autovacuum | on |
| autovacuum_analyze_scale_factor | 0.1 |
| autovacuum_analyze_threshold | 50 |
| autovacuum_freeze_max_age | 200000000 |
| autovacuum_max_workers | 3 |
| autovacuum_multixact_freeze_max_age | 400000000 |
| autovacuum_naptime | 60 s |
| autovacuum_vacuum_cost_delay | 2 ms |
| autovacuum_vacuum_cost_limit | -1 |
| autovacuum_vacuum_insert_scale_factor | 0.2 |
| autovacuum_vacuum_insert_threshold | 1000 |
| autovacuum_vacuum_scale_factor | 0.2 |
| autovacuum_vacuum_threshold | 50 |
| autovacuum_work_mem | -1 kB |
| backend_flush_after | 0 8kB |
| backslash_quote | safe_encoding |
| backtrace_functions |  |
| bgwriter_delay | 200 ms |
| bgwriter_flush_after | 64 8kB |
| bgwriter_lru_maxpages | 100 |
| bgwriter_lru_multiplier | 2 |
| block_size | 8192 |
| bonjour | off |
| bonjour_name |  |
| bytea_output | hex |
| check_function_bodies | on |
| checkpoint_completion_target | 0.9 |
| checkpoint_flush_after | 32 8kB |
| checkpoint_timeout | 300 s |
| checkpoint_warning | 30 s |
| client_connection_check_interval | 0 ms |
| client_encoding | UTF8 |
| client_min_messages | notice |
| cluster_name |  |
| commit_delay | 0 |
| commit_siblings | 5 |
| compute_query_id | auto |
| config_file | /var/lib/postgresql/data/postgresql.conf |
| constraint_exclusion | partition |
| cpu_index_tuple_cost | 0.005 |
| cpu_operator_cost | 0.0025 |
| cpu_tuple_cost | 0.01 |
| createrole_self_grant |  |
| cursor_tuple_fraction | 0.1 |
| data_checksums | off |
| data_directory | /var/lib/postgresql/data |
| data_directory_mode | 0700 |
| data_sync_retry | off |
| DateStyle | ISO, MDY |
| db_user_namespace | off |
| deadlock_timeout | 1000 ms |
| debug_assertions | off |
| debug_discard_caches | 0 |
| debug_io_direct |  |
| debug_logical_replication_streaming | buffered |
| debug_parallel_query | off |
| debug_pretty_print | on |
| debug_print_parse | off |
| debug_print_plan | off |
| debug_print_rewritten | off |
| default_statistics_target | 100 |
| default_table_access_method | heap |
| default_tablespace |  |
| default_text_search_config | pg_catalog.english |
| default_toast_compression | pglz |
| default_transaction_deferrable | off |
| default_transaction_isolation | read committed |
| default_transaction_read_only | off |
| dynamic_library_path | $libdir |
| dynamic_shared_memory_type | posix |
| effective_cache_size | 524288 8kB |
| effective_io_concurrency | 1 |
| enable_async_append | on |
| enable_bitmapscan | on |
| enable_gathermerge | on |
| enable_hashagg | on |
| enable_hashjoin | on |
| enable_incremental_sort | on |
| enable_indexonlyscan | on |
| enable_indexscan | on |
| enable_material | on |
| enable_memoize | on |
| enable_mergejoin | on |
| enable_nestloop | on |
| enable_parallel_append | on |
| enable_parallel_hash | on |
| enable_partition_pruning | on |
| enable_partitionwise_aggregate | off |
| enable_partitionwise_join | off |
| enable_presorted_aggregate | on |
| enable_seqscan | on |
| enable_sort | on |
| enable_tidscan | on |
| escape_string_warning | on |
| event_source | PostgreSQL |
| exit_on_error | off |
| extension_destdir |  |
| external_pid_file |  |
| extra_float_digits | 1 |
| file_extend_method | posix_fallocate |
| from_collapse_limit | 8 |
| fsync | on |
| full_page_writes | on |
| geqo | on |
| geqo_effort | 5 |
| geqo_generations | 0 |
| geqo_pool_size | 0 |
| geqo_seed | 0 |
| geqo_selection_bias | 2 |
| geqo_threshold | 12 |
| gin_fuzzy_search_limit | 0 |
| gin_pending_list_limit | 4096 kB |
| gss_accept_delegation | off |
| hash_mem_multiplier | 2 |
| hba_file | /var/lib/postgresql/data/pg_hba.conf |
| hot_standby | on |
| hot_standby_feedback | off |
| huge_pages | try |
| huge_page_size | 0 kB |
| icu_validation_level | warning |
| ident_file | /var/lib/postgresql/data/pg_ident.conf |
| idle_in_transaction_session_timeout | 0 ms |
| idle_session_timeout | 0 ms |
| ignore_checksum_failure | off |
| ignore_invalid_pages | off |
| ignore_system_indexes | off |
| in_hot_standby | off |
| integer_datetimes | on |
| IntervalStyle | postgres |
| jit | on |
| jit_above_cost | 100000 |
| jit_debugging_support | off |
| jit_dump_bitcode | off |
| jit_expressions | on |
| jit_inline_above_cost | 500000 |
| jit_optimize_above_cost | 500000 |
| jit_profiling_support | off |
| jit_provider | llvmjit |
| jit_tuple_deforming | on |
| join_collapse_limit | 8 |
| krb_caseins_users | off |
| krb_server_keyfile | FILE:/etc/postgresql-common/krb5.keytab |
| lc_messages | en_US.utf8 |
| lc_monetary | en_US.utf8 |
| lc_numeric | en_US.utf8 |
| lc_time | en_US.utf8 |
| listen_addresses | * |
| local_preload_libraries |  |
| lock_timeout | 0 ms |
| lo_compat_privileges | off |
| log_autovacuum_min_duration | 600000 ms |
| log_checkpoints | on |
| log_connections | off |
| log_destination | stderr |
| log_directory | log |
| log_disconnections | off |
| log_duration | off |
| log_error_verbosity | default |
| log_executor_stats | off |
| log_file_mode | 0600 |
| log_filename | postgresql-%Y-%m-%d_%H%M%S.log |
| logging_collector | off |
| log_hostname | off |
| logical_decoding_work_mem | 65536 kB |
| log_line_prefix | %m [%p]  |
| log_lock_waits | off |
| log_min_duration_sample | -1 ms |
| log_min_duration_statement | -1 ms |
| log_min_error_statement | error |
| log_min_messages | warning |
| log_parameter_max_length | -1 B |
| log_parameter_max_length_on_error | 0 B |
| log_parser_stats | off |
| log_planner_stats | off |
| log_recovery_conflict_waits | off |
| log_replication_commands | off |
| log_rotation_age | 1440 min |
| log_rotation_size | 10240 kB |
| log_startup_progress_interval | 10000 ms |
| log_statement | none |
| log_statement_sample_rate | 1 |
| log_statement_stats | off |
| log_temp_files | -1 kB |
| log_timezone | Etc/UTC |
| log_transaction_sample_rate | 0 |
| log_truncate_on_rotation | off |
| maintenance_io_concurrency | 10 |
| maintenance_work_mem | 65536 kB |
| max_connections | 100 |
| max_files_per_process | 1000 |
| max_function_args | 100 |
| max_identifier_length | 63 |
| max_index_keys | 32 |
| max_locks_per_transaction | 64 |
| max_logical_replication_workers | 4 |
| max_parallel_apply_workers_per_subscription | 2 |
| max_parallel_maintenance_workers | 2 |
| max_parallel_workers | 8 |
| max_parallel_workers_per_gather | 2 |
| max_pred_locks_per_page | 2 |
| max_pred_locks_per_relation | -2 |
| max_pred_locks_per_transaction | 64 |
| max_prepared_transactions | 0 |
| max_replication_slots | 10 |
| max_slot_wal_keep_size | -1 MB |
| max_stack_depth | 2048 kB |
| max_standby_archive_delay | 30000 ms |
| max_standby_streaming_delay | 30000 ms |
| max_sync_workers_per_subscription | 2 |
| max_wal_senders | 10 |
| max_wal_size | 1024 MB |
| max_worker_processes | 8 |
| min_dynamic_shared_memory | 0 MB |
| min_parallel_index_scan_size | 64 8kB |
| min_parallel_table_scan_size | 1024 8kB |
| min_wal_size | 80 MB |
| old_snapshot_threshold | -1 min |
| parallel_leader_participation | on |
| parallel_setup_cost | 1000 |
| parallel_tuple_cost | 0.1 |
| password_encryption | scram-sha-256 |
| plan_cache_mode | auto |
| port | 5432 |
| post_auth_delay | 0 s |
| pre_auth_delay | 0 s |
| primary_conninfo |  |
| primary_slot_name |  |
| quote_all_identifiers | off |
| random_page_cost | 4 |
| recovery_end_command |  |
| recovery_init_sync_method | fsync |
| recovery_min_apply_delay | 0 ms |
| recovery_prefetch | try |
| recovery_target |  |
| recovery_target_action | pause |
| recovery_target_inclusive | on |
| recovery_target_lsn |  |
| recovery_target_name |  |
| recovery_target_time |  |
| recovery_target_timeline | latest |
| recovery_target_xid |  |
| recursive_worktable_factor | 10 |
| remove_temp_files_after_crash | on |
| reserved_connections | 0 |
| restart_after_crash | on |
| restore_command |  |
| restrict_nonsystem_relation_kind |  |
| row_security | on |
| scram_iterations | 4096 |
| search_path | "$user", public |
| segment_size | 131072 8kB |
| send_abort_for_crash | off |
| send_abort_for_kill | off |
| seq_page_cost | 1 |
| server_encoding | UTF8 |
| server_version | 16.13 (Debian 16.13-1.pgdg13+1) |
| server_version_num | 160013 |
| session_preload_libraries |  |
| session_replication_role | origin |
| shared_buffers | 16384 8kB |
| shared_memory_size | 143 MB |
| shared_memory_size_in_huge_pages | 72 |
| shared_memory_type | mmap |
| shared_preload_libraries |  |
| ssl | off |
| ssl_ca_file |  |
| ssl_cert_file | server.crt |
| ssl_ciphers | HIGH:MEDIUM:+3DES:!aNULL |
| ssl_crl_dir |  |
| ssl_crl_file |  |
| ssl_dh_params_file |  |
| ssl_ecdh_curve | prime256v1 |
| ssl_key_file | server.key |
| ssl_library | OpenSSL |
| ssl_max_protocol_version |  |
| ssl_min_protocol_version | TLSv1.2 |
| ssl_passphrase_command |  |
| ssl_passphrase_command_supports_reload | off |
| ssl_prefer_server_ciphers | on |
| standard_conforming_strings | on |
| statement_timeout | 0 ms |
| stats_fetch_consistency | cache |
| superuser_reserved_connections | 3 |
| synchronize_seqscans | on |
| synchronous_commit | on |
| synchronous_standby_names |  |
| syslog_facility | local0 |
| syslog_ident | postgres |
| syslog_sequence_numbers | on |
| syslog_split_messages | on |
| tcp_keepalives_count | 9 |
| tcp_keepalives_idle | 7200 s |
| tcp_keepalives_interval | 75 s |
| tcp_user_timeout | 0 ms |
| temp_buffers | 1024 8kB |
| temp_file_limit | -1 kB |
| temp_tablespaces |  |
| TimeZone | Etc/UTC |
| timezone_abbreviations | Default |
| trace_notify | off |
| trace_recovery_messages | log |
| trace_sort | off |
| track_activities | on |
| track_activity_query_size | 1024 B |
| track_commit_timestamp | off |
| track_counts | on |
| track_functions | none |
| track_io_timing | off |
| track_wal_io_timing | off |
| transaction_deferrable | off |
| transaction_isolation | read committed |
| transaction_read_only | off |
| transform_null_equals | off |
| unix_socket_directories | /var/run/postgresql |
| unix_socket_group |  |
| unix_socket_permissions | 0777 |
| update_process_title | on |
| vacuum_buffer_usage_limit | 256 kB |
| vacuum_cost_delay | 0 ms |
| vacuum_cost_limit | 200 |
| vacuum_cost_page_dirty | 20 |
| vacuum_cost_page_hit | 1 |
| vacuum_cost_page_miss | 2 |
| vacuum_failsafe_age | 1600000000 |
| vacuum_freeze_min_age | 50000000 |
| vacuum_freeze_table_age | 150000000 |
| vacuum_multixact_failsafe_age | 1600000000 |
| vacuum_multixact_freeze_min_age | 5000000 |
| vacuum_multixact_freeze_table_age | 150000000 |
| wal_block_size | 8192 |
| wal_buffers | 512 8kB |
| wal_compression | off |
| wal_consistency_checking |  |
| wal_decode_buffer_size | 524288 B |
| wal_init_zero | on |
| wal_keep_size | 0 MB |
| wal_level | replica |
| wal_log_hints | off |
| wal_receiver_create_temp_slot | off |
| wal_receiver_status_interval | 10 s |
| wal_receiver_timeout | 60000 ms |
| wal_recycle | on |
| wal_retrieve_retry_interval | 5000 ms |
| wal_segment_size | 16777216 B |
| wal_sender_timeout | 60000 ms |
| wal_skip_threshold | 2048 kB |
| wal_sync_method | fdatasync |
| wal_writer_delay | 200 ms |
| wal_writer_flush_after | 128 8kB |
| work_mem | 4096 kB |
| xmlbinary | base64 |
| xmloption | content |
| zero_damaged_pages | off |

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

