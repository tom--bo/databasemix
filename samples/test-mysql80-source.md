# File Summary

This file contains comprehensive MySQL database information compiled for AI context analysis. It includes schema definitions, account configurations, system variables, and other database metadata consolidated into a single file for efficient processing.

**Database Type**: MySQL  
**Database Version**: 8.0.45

## File Structure

- Variables - MySQL system variables and their current values
- Tables - Database tables with metadata and DDL definitions
- View Details - Database views with their definitions
- Stored Functions - User-defined functions with their definitions
- Stored Procedures - User-defined procedures with their definitions
- User Roles - Role definitions and assignments
- User Accounts - Database user accounts with privileges

# Variables

| Variable Name | Current Value | Source |
|---------------|---------------|--------|
| activate_all_roles_on_login | OFF | COMPILED |
| admin_address |  | COMPILED |
| admin_port | 33062 | COMPILED |
| admin_ssl_ca |  | COMPILED |
| admin_ssl_capath |  | COMPILED |
| admin_ssl_cert |  | COMPILED |
| admin_ssl_cipher |  | COMPILED |
| admin_ssl_crl |  | COMPILED |
| admin_ssl_crlpath |  | COMPILED |
| admin_ssl_key |  | COMPILED |
| admin_tls_ciphersuites |  | COMPILED |
| admin_tls_version | TLSv1.2,TLSv1.3 | COMPILED |
| authentication_policy | *,, | COMPILED |
| auto_generate_certs | ON | COMPILED |
| auto_increment_increment | 1 | COMPILED |
| auto_increment_offset | 1 | COMPILED |
| autocommit | ON | COMPILED |
| automatic_sp_privileges | ON | COMPILED |
| avoid_temporal_upgrade | OFF | COMPILED |
| back_log | 151 | COMPILED |
| basedir | /usr/ | COMPILED |
| big_tables | OFF | COMPILED |
| bind_address | * | COMPILED |
| binlog_cache_size | 32768 | COMPILED |
| binlog_checksum | CRC32 | COMPILED |
| binlog_direct_non_transactional_updates | OFF | COMPILED |
| binlog_encryption | OFF | COMPILED |
| binlog_error_action | ABORT_SERVER | COMPILED |
| binlog_expire_logs_auto_purge | ON | COMPILED |
| binlog_expire_logs_seconds | 2592000 | COMPILED |
| binlog_format | ROW | COMPILED |
| binlog_group_commit_sync_delay | 0 | COMPILED |
| binlog_group_commit_sync_no_delay_count | 0 | COMPILED |
| binlog_gtid_simple_recovery | ON | COMPILED |
| binlog_max_flush_queue_time | 0 | COMPILED |
| binlog_order_commits | ON | COMPILED |
| binlog_rotate_encryption_master_key_at_startup | OFF | COMPILED |
| binlog_row_event_max_size | 8192 | COMPILED |
| binlog_row_image | FULL | COMPILED |
| binlog_row_metadata | MINIMAL | COMPILED |
| binlog_row_value_options |  | COMPILED |
| binlog_rows_query_log_events | OFF | COMPILED |
| binlog_stmt_cache_size | 32768 | COMPILED |
| binlog_transaction_compression | OFF | COMPILED |
| binlog_transaction_compression_level_zstd | 3 | COMPILED |
| binlog_transaction_dependency_history_size | 25000 | COMPILED |
| binlog_transaction_dependency_tracking | COMMIT_ORDER | COMPILED |
| block_encryption_mode | aes-128-ecb | COMPILED |
| build_id | 3a62bc18e392ecf8a4ac5e882191b79d9382f04b | COMPILED |
| bulk_insert_buffer_size | 8388608 | COMPILED |
| caching_sha2_password_auto_generate_rsa_keys | ON | COMPILED |
| caching_sha2_password_digest_rounds | 5000 | COMPILED |
| caching_sha2_password_private_key_path | private_key.pem | COMPILED |
| caching_sha2_password_public_key_path | public_key.pem | COMPILED |
| character_set_client | utf8mb4 | COMPILED |
| character_set_connection | utf8mb4 | COMPILED |
| character_set_database | utf8mb4 | COMPILED |
| character_set_filesystem | binary | COMPILED |
| character_set_results | utf8mb4 | COMPILED |
| character_set_server | utf8mb4 | COMPILED |
| character_set_system | utf8mb3 | COMPILED |
| character_sets_dir | /usr/share/mysql-8.0/charsets/ | COMPILED |
| check_proxy_users | OFF | COMPILED |
| collation_connection | utf8mb4_unicode_ci | COMPILED |
| collation_database | utf8mb4_unicode_ci | COMPILED |
| collation_server | utf8mb4_unicode_ci | COMPILED |
| completion_type | NO_CHAIN | COMPILED |
| concurrent_insert | AUTO | COMPILED |
| connect_timeout | 10 | COMPILED |
| connection_memory_chunk_size | 8192 | COMPILED |
| connection_memory_limit | 18446744073709551615 | COMPILED |
| core_file | OFF | COMPILED |
| create_admin_listener_thread | OFF | COMPILED |
| cte_max_recursion_depth | 1000 | COMPILED |
| datadir | /var/lib/mysql/ | GLOBAL (/etc/my.cnf) |
| default_authentication_plugin | mysql_native_password | COMMAND_LINE |
| default_collation_for_utf8mb4 | utf8mb4_0900_ai_ci | COMPILED |
| default_password_lifetime | 0 | COMPILED |
| default_storage_engine | InnoDB | COMPILED |
| default_table_encryption | OFF | COMPILED |
| default_tmp_storage_engine | InnoDB | COMPILED |
| default_week_format | 0 | COMPILED |
| delay_key_write | ON | COMPILED |
| delayed_insert_limit | 100 | COMPILED |
| delayed_insert_timeout | 300 | COMPILED |
| delayed_queue_size | 1000 | COMPILED |
| disabled_storage_engines |  | COMPILED |
| disconnect_on_expired_password | ON | COMPILED |
| div_precision_increment | 4 | COMPILED |
| end_markers_in_json | OFF | COMPILED |
| enforce_gtid_consistency | OFF | COMPILED |
| eq_range_index_dive_limit | 200 | COMPILED |
| event_scheduler | ON | COMPILED |
| expire_logs_days | 0 | COMPILED |
| explain_format | TRADITIONAL | COMPILED |
| explicit_defaults_for_timestamp | ON | COMPILED |
| flush | OFF | COMPILED |
| flush_time | 0 | COMPILED |
| foreign_key_checks | ON | DYNAMIC |
| ft_boolean_syntax | + -><()~*:""&| | COMPILED |
| ft_max_word_len | 84 | COMPILED |
| ft_min_word_len | 4 | COMPILED |
| ft_query_expansion_limit | 20 | COMPILED |
| ft_stopword_file | (built-in) | COMPILED |
| general_log | OFF | COMPILED |
| general_log_file | /var/lib/mysql/cd786d93dc11.log | COMPILED |
| generated_random_password_length | 20 | COMPILED |
| global_connection_memory_limit | 18446744073709551615 | COMPILED |
| global_connection_memory_tracking | OFF | COMPILED |
| group_concat_max_len | 1024 | COMPILED |
| group_replication_consistency | EVENTUAL | COMPILED |
| gtid_executed |  | COMPILED |
| gtid_executed_compression_period | 0 | COMPILED |
| gtid_mode | OFF | COMPILED |
| gtid_owned |  | COMPILED |
| gtid_purged |  | COMPILED |
| have_compress | YES | COMPILED |
| have_dynamic_loading | YES | COMPILED |
| have_geometry | YES | COMPILED |
| have_openssl | YES | COMPILED |
| have_profiling | YES | COMPILED |
| have_query_cache | NO | COMPILED |
| have_rtree_keys | YES | COMPILED |
| have_ssl | YES | COMPILED |
| have_statement_timeout | YES | COMPILED |
| have_symlink | DISABLED | COMPILED |
| histogram_generation_max_mem_size | 20000000 | COMPILED |
| host_cache_size | 279 | COMPILED |
| hostname | cd786d93dc11 | COMPILED |
| information_schema_stats_expiry | 86400 | COMPILED |
| init_connect |  | COMPILED |
| init_file |  | COMPILED |
| init_replica |  | COMPILED |
| init_slave |  | COMPILED |
| innodb_adaptive_flushing | ON | COMPILED |
| innodb_adaptive_flushing_lwm | 10 | COMPILED |
| innodb_adaptive_hash_index | ON | COMPILED |
| innodb_adaptive_hash_index_parts | 8 | COMPILED |
| innodb_adaptive_max_sleep_delay | 150000 | COMPILED |
| innodb_api_bk_commit_interval | 5 | COMPILED |
| innodb_api_disable_rowlock | OFF | COMPILED |
| innodb_api_enable_binlog | OFF | COMPILED |
| innodb_api_enable_mdl | OFF | COMPILED |
| innodb_api_trx_level | 0 | COMPILED |
| innodb_autoextend_increment | 64 | COMPILED |
| innodb_autoinc_lock_mode | 2 | COMPILED |
| innodb_buffer_pool_chunk_size | 134217728 | COMPILED |
| innodb_buffer_pool_dump_at_shutdown | ON | COMPILED |
| innodb_buffer_pool_dump_now | OFF | COMPILED |
| innodb_buffer_pool_dump_pct | 25 | COMPILED |
| innodb_buffer_pool_filename | ib_buffer_pool | COMPILED |
| innodb_buffer_pool_in_core_file | ON | COMPILED |
| innodb_buffer_pool_instances | 1 | COMPILED |
| innodb_buffer_pool_load_abort | OFF | COMPILED |
| innodb_buffer_pool_load_at_startup | ON | COMPILED |
| innodb_buffer_pool_load_now | OFF | COMPILED |
| innodb_buffer_pool_size | 134217728 | COMPILED |
| innodb_change_buffer_max_size | 25 | COMPILED |
| innodb_change_buffering | all | COMPILED |
| innodb_checksum_algorithm | crc32 | COMPILED |
| innodb_cmp_per_index_enabled | OFF | COMPILED |
| innodb_commit_concurrency | 0 | COMPILED |
| innodb_compression_failure_threshold_pct | 5 | COMPILED |
| innodb_compression_level | 6 | COMPILED |
| innodb_compression_pad_pct_max | 50 | COMPILED |
| innodb_concurrency_tickets | 5000 | COMPILED |
| innodb_data_file_path | ibdata1:12M:autoextend | COMPILED |
| innodb_data_home_dir |  | COMPILED |
| innodb_ddl_buffer_size | 1048576 | COMPILED |
| innodb_ddl_threads | 4 | COMPILED |
| innodb_deadlock_detect | ON | COMPILED |
| innodb_dedicated_server | OFF | COMPILED |
| innodb_default_row_format | dynamic | COMPILED |
| innodb_directories |  | COMPILED |
| innodb_disable_sort_file_cache | OFF | COMPILED |
| innodb_doublewrite | ON | COMPILED |
| innodb_doublewrite_batch_size | 0 | COMPILED |
| innodb_doublewrite_dir |  | COMPILED |
| innodb_doublewrite_files | 2 | COMPILED |
| innodb_doublewrite_pages | 4 | COMPILED |
| innodb_extend_and_initialize | ON | COMPILED |
| innodb_fast_shutdown | 1 | COMPILED |
| innodb_file_per_table | ON | COMPILED |
| innodb_fill_factor | 100 | COMPILED |
| innodb_flush_log_at_timeout | 1 | COMPILED |
| innodb_flush_log_at_trx_commit | 1 | COMPILED |
| innodb_flush_method | fsync | COMPILED |
| innodb_flush_neighbors | 0 | COMPILED |
| innodb_flush_sync | ON | COMPILED |
| innodb_flushing_avg_loops | 30 | COMPILED |
| innodb_force_load_corrupted | OFF | COMPILED |
| innodb_force_recovery | 0 | COMPILED |
| innodb_fsync_threshold | 0 | COMPILED |
| innodb_ft_aux_table |  | COMPILED |
| innodb_ft_cache_size | 8000000 | COMPILED |
| innodb_ft_enable_diag_print | OFF | COMPILED |
| innodb_ft_enable_stopword | ON | COMPILED |
| innodb_ft_max_token_size | 84 | COMPILED |
| innodb_ft_min_token_size | 3 | COMPILED |
| innodb_ft_num_word_optimize | 2000 | COMPILED |
| innodb_ft_result_cache_limit | 2000000000 | COMPILED |
| innodb_ft_server_stopword_table |  | COMPILED |
| innodb_ft_sort_pll_degree | 2 | COMPILED |
| innodb_ft_total_cache_size | 640000000 | COMPILED |
| innodb_ft_user_stopword_table |  | COMPILED |
| innodb_idle_flush_pct | 100 | COMPILED |
| innodb_io_capacity | 200 | COMPILED |
| innodb_io_capacity_max | 2000 | COMPILED |
| innodb_lock_wait_timeout | 50 | COMPILED |
| innodb_log_buffer_size | 16777216 | COMPILED |
| innodb_log_checksums | ON | COMPILED |
| innodb_log_compressed_pages | ON | COMPILED |
| innodb_log_file_size | 50331648 | COMPILED |
| innodb_log_files_in_group | 2 | COMPILED |
| innodb_log_group_home_dir | ./ | COMPILED |
| innodb_log_spin_cpu_abs_lwm | 80 | COMPILED |
| innodb_log_spin_cpu_pct_hwm | 50 | COMPILED |
| innodb_log_wait_for_flush_spin_hwm | 400 | COMPILED |
| innodb_log_write_ahead_size | 8192 | COMPILED |
| innodb_log_writer_threads | ON | COMPILED |
| innodb_lru_scan_depth | 1024 | COMPILED |
| innodb_max_dirty_pages_pct | 90.000000 | COMPILED |
| innodb_max_dirty_pages_pct_lwm | 10.000000 | COMPILED |
| innodb_max_purge_lag | 0 | COMPILED |
| innodb_max_purge_lag_delay | 0 | COMPILED |
| innodb_max_undo_log_size | 1073741824 | COMPILED |
| innodb_monitor_disable |  | COMPILED |
| innodb_monitor_enable |  | COMPILED |
| innodb_monitor_reset |  | COMPILED |
| innodb_monitor_reset_all |  | COMPILED |
| innodb_old_blocks_pct | 37 | COMPILED |
| innodb_old_blocks_time | 1000 | COMPILED |
| innodb_online_alter_log_max_size | 134217728 | COMPILED |
| innodb_open_files | 4000 | COMPILED |
| innodb_optimize_fulltext_only | OFF | COMPILED |
| innodb_page_cleaners | 1 | COMPILED |
| innodb_page_size | 16384 | COMPILED |
| innodb_parallel_read_threads | 4 | COMPILED |
| innodb_print_all_deadlocks | OFF | COMPILED |
| innodb_print_ddl_logs | OFF | COMPILED |
| innodb_purge_batch_size | 300 | COMPILED |
| innodb_purge_rseg_truncate_frequency | 128 | COMPILED |
| innodb_purge_threads | 4 | COMPILED |
| innodb_random_read_ahead | OFF | COMPILED |
| innodb_read_ahead_threshold | 56 | COMPILED |
| innodb_read_io_threads | 4 | COMPILED |
| innodb_read_only | OFF | COMPILED |
| innodb_redo_log_archive_dirs |  | COMPILED |
| innodb_redo_log_capacity | 104857600 | COMPILED |
| innodb_redo_log_encrypt | OFF | COMPILED |
| innodb_replication_delay | 0 | COMPILED |
| innodb_rollback_on_timeout | OFF | COMPILED |
| innodb_rollback_segments | 128 | COMPILED |
| innodb_segment_reserve_factor | 12.500000 | COMPILED |
| innodb_sort_buffer_size | 1048576 | COMPILED |
| innodb_spin_wait_delay | 6 | COMPILED |
| innodb_spin_wait_pause_multiplier | 50 | COMPILED |
| innodb_stats_auto_recalc | ON | COMPILED |
| innodb_stats_include_delete_marked | OFF | COMPILED |
| innodb_stats_method | nulls_equal | COMPILED |
| innodb_stats_on_metadata | OFF | COMPILED |
| innodb_stats_persistent | ON | COMPILED |
| innodb_stats_persistent_sample_pages | 20 | COMPILED |
| innodb_stats_transient_sample_pages | 8 | COMPILED |
| innodb_status_output | OFF | COMPILED |
| innodb_status_output_locks | OFF | COMPILED |
| innodb_strict_mode | ON | COMPILED |
| innodb_sync_array_size | 1 | COMPILED |
| innodb_sync_spin_loops | 30 | COMPILED |
| innodb_table_locks | ON | COMPILED |
| innodb_temp_data_file_path | ibtmp1:12M:autoextend | COMPILED |
| innodb_temp_tablespaces_dir | ./#innodb_temp/ | COMPILED |
| innodb_thread_concurrency | 0 | COMPILED |
| innodb_thread_sleep_delay | 10000 | COMPILED |
| innodb_tmpdir |  | COMPILED |
| innodb_undo_directory | ./ | COMPILED |
| innodb_undo_log_encrypt | OFF | COMPILED |
| innodb_undo_log_truncate | ON | COMPILED |
| innodb_undo_tablespaces | 2 | COMPILED |
| innodb_use_fdatasync | OFF | COMPILED |
| innodb_use_native_aio | ON | COMPILED |
| innodb_validate_tablespace_paths | ON | COMPILED |
| innodb_version | 8.0.45 | COMPILED |
| innodb_write_io_threads | 4 | COMPILED |
| interactive_timeout | 28800 | COMPILED |
| internal_tmp_mem_storage_engine | TempTable | COMPILED |
| join_buffer_size | 262144 | COMPILED |
| keep_files_on_create | OFF | COMPILED |
| key_buffer_size | 8388608 | COMPILED |
| key_cache_age_threshold | 300 | COMPILED |
| key_cache_block_size | 1024 | COMPILED |
| key_cache_division_limit | 100 | COMPILED |
| keyring_operations | ON | COMPILED |
| large_files_support | ON | COMPILED |
| large_page_size | 0 | COMPILED |
| large_pages | OFF | COMPILED |
| lc_messages | en_US | COMPILED |
| lc_messages_dir | /usr/share/mysql-8.0/ | COMPILED |
| lc_time_names | en_US | COMPILED |
| license | GPL | COMPILED |
| local_infile | OFF | COMPILED |
| lock_wait_timeout | 31536000 | COMPILED |
| locked_in_memory | OFF | COMPILED |
| log_bin | ON | COMPILED |
| log_bin_basename | /var/lib/mysql/mysql-bin | COMPILED |
| log_bin_index | /var/lib/mysql/mysql-bin.index | COMPILED |
| log_bin_trust_function_creators | OFF | COMPILED |
| log_bin_use_v1_row_events | OFF | COMPILED |
| log_error | stderr | COMPILED |
| log_error_services | log_filter_internal; log_sink_internal | COMPILED |
| log_error_suppression_list |  | COMPILED |
| log_error_verbosity | 2 | COMPILED |
| log_output | FILE | COMPILED |
| log_queries_not_using_indexes | OFF | COMPILED |
| log_raw | OFF | COMPILED |
| log_replica_updates | ON | COMPILED |
| log_slave_updates | ON | COMPILED |
| log_slow_admin_statements | OFF | COMPILED |
| log_slow_extra | OFF | COMPILED |
| log_slow_replica_statements | OFF | COMPILED |
| log_slow_slave_statements | OFF | COMPILED |
| log_statements_unsafe_for_binlog | ON | COMPILED |
| log_throttle_queries_not_using_indexes | 0 | COMPILED |
| log_timestamps | UTC | COMPILED |
| long_query_time | 10.000000 | COMPILED |
| low_priority_updates | OFF | COMPILED |
| lower_case_file_system | OFF | COMPILED |
| lower_case_table_names | 0 | COMPILED |
| mandatory_roles |  | COMPILED |
| master_info_repository | TABLE | COMPILED |
| master_verify_checksum | OFF | COMPILED |
| max_allowed_packet | 67108864 | COMPILED |
| max_binlog_cache_size | 18446744073709547520 | COMPILED |
| max_binlog_size | 1073741824 | COMPILED |
| max_binlog_stmt_cache_size | 18446744073709547520 | COMPILED |
| max_connect_errors | 100 | COMPILED |
| max_connections | 151 | COMPILED |
| max_delayed_threads | 20 | COMPILED |
| max_digest_length | 1024 | COMPILED |
| max_error_count | 1024 | COMPILED |
| max_execution_time | 0 | COMPILED |
| max_heap_table_size | 16777216 | COMPILED |
| max_insert_delayed_threads | 20 | COMPILED |
| max_join_size | 18446744073709551615 | COMPILED |
| max_length_for_sort_data | 4096 | COMPILED |
| max_points_in_geometry | 65536 | COMPILED |
| max_prepared_stmt_count | 16382 | COMPILED |
| max_relay_log_size | 0 | COMPILED |
| max_seeks_for_key | 18446744073709551615 | COMPILED |
| max_sort_length | 1024 | COMPILED |
| max_sp_recursion_depth | 0 | COMPILED |
| max_user_connections | 0 | COMPILED |
| max_write_lock_count | 18446744073709551615 | COMPILED |
| min_examined_row_limit | 0 | COMPILED |
| myisam_data_pointer_size | 6 | COMPILED |
| myisam_max_sort_file_size | 9223372036853727232 | COMPILED |
| myisam_mmap_size | 18446744073709551615 | COMPILED |
| myisam_recover_options | OFF | COMPILED |
| myisam_sort_buffer_size | 8388608 | COMPILED |
| myisam_stats_method | nulls_unequal | COMPILED |
| myisam_use_mmap | OFF | COMPILED |
| mysql_native_password_proxy_users | OFF | COMPILED |
| mysqlx_bind_address | * | COMPILED |
| mysqlx_compression_algorithms | DEFLATE_STREAM,LZ4_MESSAGE,ZSTD_STREAM | COMPILED |
| mysqlx_connect_timeout | 30 | COMPILED |
| mysqlx_deflate_default_compression_level | 3 | COMPILED |
| mysqlx_deflate_max_client_compression_level | 5 | COMPILED |
| mysqlx_document_id_unique_prefix | 0 | COMPILED |
| mysqlx_enable_hello_notice | ON | COMPILED |
| mysqlx_idle_worker_thread_timeout | 60 | COMPILED |
| mysqlx_interactive_timeout | 28800 | COMPILED |
| mysqlx_lz4_default_compression_level | 2 | COMPILED |
| mysqlx_lz4_max_client_compression_level | 8 | COMPILED |
| mysqlx_max_allowed_packet | 67108864 | COMPILED |
| mysqlx_max_connections | 100 | COMPILED |
| mysqlx_min_worker_threads | 2 | COMPILED |
| mysqlx_port | 33060 | COMPILED |
| mysqlx_port_open_timeout | 0 | COMPILED |
| mysqlx_read_timeout | 30 | COMPILED |
| mysqlx_socket | /var/run/mysqld/mysqlx.sock | COMPILED |
| mysqlx_ssl_ca |  | COMPILED |
| mysqlx_ssl_capath |  | COMPILED |
| mysqlx_ssl_cert |  | COMPILED |
| mysqlx_ssl_cipher |  | COMPILED |
| mysqlx_ssl_crl |  | COMPILED |
| mysqlx_ssl_crlpath |  | COMPILED |
| mysqlx_ssl_key |  | COMPILED |
| mysqlx_wait_timeout | 28800 | COMPILED |
| mysqlx_write_timeout | 60 | COMPILED |
| mysqlx_zstd_default_compression_level | 3 | COMPILED |
| mysqlx_zstd_max_client_compression_level | 11 | COMPILED |
| net_buffer_length | 16384 | COMPILED |
| net_read_timeout | 30 | COMPILED |
| net_retry_count | 10 | COMPILED |
| net_write_timeout | 60 | COMPILED |
| new | OFF | COMPILED |
| ngram_token_size | 2 | COMPILED |
| offline_mode | OFF | COMPILED |
| old | OFF | COMPILED |
| old_alter_table | OFF | COMPILED |
| open_files_limit | 1048576 | COMPILED |
| optimizer_max_subgraph_pairs | 100000 | COMPILED |
| optimizer_prune_level | 1 | COMPILED |
| optimizer_search_depth | 62 | COMPILED |
| optimizer_switch | index_merge=on,index_merge_union=on,index_merge_sort_union=on,index_merge_intersection=on,engine_condition_pushdown=on,index_condition_pushdown=on,mrr=on,mrr_cost_based=on,block_nested_loop=on,batched_key_access=off,materialization=on,semijoin=on,loosescan=on,firstmatch=on,duplicateweedout=on,subquery_materialization_cost_based=on,use_index_extensions=on,condition_fanout_filter=on,derived_merge=on,use_invisible_indexes=off,skip_scan=on,hash_join=on,subquery_to_derived=off,prefer_ordering_index=on,hypergraph_optimizer=off,derived_condition_pushdown=on | COMPILED |
| optimizer_trace | enabled=off,one_line=off | COMPILED |
| optimizer_trace_features | greedy_search=on,range_optimizer=on,dynamic_range=on,repeated_subselect=on | COMPILED |
| optimizer_trace_limit | 1 | COMPILED |
| optimizer_trace_max_mem_size | 1048576 | COMPILED |
| optimizer_trace_offset | -1 | COMPILED |
| parser_max_mem_size | 18446744073709551615 | COMPILED |
| partial_revokes | OFF | COMPILED |
| password_history | 0 | COMPILED |
| password_require_current | OFF | COMPILED |
| password_reuse_interval | 0 | COMPILED |
| performance_schema | ON | COMPILED |
| performance_schema_accounts_size | -1 | COMPILED |
| performance_schema_digests_size | 10000 | COMPILED |
| performance_schema_error_size | 5325 | COMPILED |
| performance_schema_events_stages_history_long_size | 10000 | COMPILED |
| performance_schema_events_stages_history_size | 10 | COMPILED |
| performance_schema_events_statements_history_long_size | 10000 | COMPILED |
| performance_schema_events_statements_history_size | 10 | COMPILED |
| performance_schema_events_transactions_history_long_size | 10000 | COMPILED |
| performance_schema_events_transactions_history_size | 10 | COMPILED |
| performance_schema_events_waits_history_long_size | 10000 | COMPILED |
| performance_schema_events_waits_history_size | 10 | COMPILED |
| performance_schema_hosts_size | -1 | COMPILED |
| performance_schema_max_cond_classes | 150 | COMPILED |
| performance_schema_max_cond_instances | -1 | COMPILED |
| performance_schema_max_digest_length | 1024 | COMPILED |
| performance_schema_max_digest_sample_age | 60 | COMPILED |
| performance_schema_max_file_classes | 80 | COMPILED |
| performance_schema_max_file_handles | 32768 | COMPILED |
| performance_schema_max_file_instances | -1 | COMPILED |
| performance_schema_max_index_stat | -1 | COMPILED |
| performance_schema_max_memory_classes | 450 | COMPILED |
| performance_schema_max_metadata_locks | -1 | COMPILED |
| performance_schema_max_mutex_classes | 350 | COMPILED |
| performance_schema_max_mutex_instances | -1 | COMPILED |
| performance_schema_max_prepared_statements_instances | -1 | COMPILED |
| performance_schema_max_program_instances | -1 | COMPILED |
| performance_schema_max_rwlock_classes | 60 | COMPILED |
| performance_schema_max_rwlock_instances | -1 | COMPILED |
| performance_schema_max_socket_classes | 10 | COMPILED |
| performance_schema_max_socket_instances | -1 | COMPILED |
| performance_schema_max_sql_text_length | 1024 | COMPILED |
| performance_schema_max_stage_classes | 175 | COMPILED |
| performance_schema_max_statement_classes | 219 | COMPILED |
| performance_schema_max_statement_stack | 10 | COMPILED |
| performance_schema_max_table_handles | -1 | COMPILED |
| performance_schema_max_table_instances | -1 | COMPILED |
| performance_schema_max_table_lock_stat | -1 | COMPILED |
| performance_schema_max_thread_classes | 100 | COMPILED |
| performance_schema_max_thread_instances | -1 | COMPILED |
| performance_schema_session_connect_attrs_size | 512 | COMPILED |
| performance_schema_setup_actors_size | -1 | COMPILED |
| performance_schema_setup_objects_size | -1 | COMPILED |
| performance_schema_show_processlist | OFF | COMPILED |
| performance_schema_users_size | -1 | COMPILED |
| persist_only_admin_x509_subject |  | COMPILED |
| persist_sensitive_variables_in_plaintext | ON | COMPILED |
| persisted_globals_load | ON | COMPILED |
| pid_file | /var/run/mysqld/mysqld.pid | GLOBAL (/etc/my.cnf) |
| plugin_dir | /usr/lib64/mysql/plugin/ | COMPILED |
| port | 3306 | COMPILED |
| preload_buffer_size | 32768 | COMPILED |
| print_identified_with_as_hex | OFF | COMPILED |
| profiling | OFF | COMPILED |
| profiling_history_size | 15 | COMPILED |
| protocol_compression_algorithms | zlib,zstd,uncompressed | COMPILED |
| protocol_version | 10 | COMPILED |
| query_alloc_block_size | 8192 | COMPILED |
| query_prealloc_size | 8192 | COMPILED |
| range_alloc_block_size | 4096 | COMPILED |
| range_optimizer_max_mem_size | 8388608 | COMPILED |
| rbr_exec_mode | STRICT | COMPILED |
| read_buffer_size | 131072 | COMPILED |
| read_only | OFF | COMPILED |
| read_rnd_buffer_size | 262144 | COMPILED |
| regexp_stack_limit | 8000000 | COMPILED |
| regexp_time_limit | 32 | COMPILED |
| relay_log | cd786d93dc11-relay-bin | COMPILED |
| relay_log_basename | /var/lib/mysql/cd786d93dc11-relay-bin | COMPILED |
| relay_log_index | /var/lib/mysql/cd786d93dc11-relay-bin.index | COMPILED |
| relay_log_info_file | relay-log.info | COMPILED |
| relay_log_info_repository | TABLE | COMPILED |
| relay_log_purge | ON | COMPILED |
| relay_log_recovery | OFF | COMPILED |
| relay_log_space_limit | 0 | COMPILED |
| replica_allow_batching | ON | COMPILED |
| replica_checkpoint_group | 512 | COMPILED |
| replica_checkpoint_period | 300 | COMPILED |
| replica_compressed_protocol | OFF | COMPILED |
| replica_exec_mode | STRICT | COMPILED |
| replica_load_tmpdir | /tmp | COMPILED |
| replica_max_allowed_packet | 1073741824 | COMPILED |
| replica_net_timeout | 60 | COMPILED |
| replica_parallel_type | LOGICAL_CLOCK | COMPILED |
| replica_parallel_workers | 4 | COMPILED |
| replica_pending_jobs_size_max | 134217728 | COMPILED |
| replica_preserve_commit_order | ON | COMPILED |
| replica_skip_errors | OFF | COMPILED |
| replica_sql_verify_checksum | ON | COMPILED |
| replica_transaction_retries | 10 | COMPILED |
| replica_type_conversions |  | COMPILED |
| replication_optimize_for_static_plugin_config | OFF | COMPILED |
| replication_sender_observe_commit_only | OFF | COMPILED |
| report_host |  | COMPILED |
| report_password |  | COMPILED |
| report_port | 3306 | COMPILED |
| report_user |  | COMPILED |
| require_secure_transport | OFF | COMPILED |
| rpl_read_size | 8192 | COMPILED |
| rpl_stop_replica_timeout | 31536000 | COMPILED |
| rpl_stop_slave_timeout | 31536000 | COMPILED |
| schema_definition_cache | 256 | COMPILED |
| secondary_engine_cost_threshold | 100000.000000 | COMPILED |
| secure_file_priv | /var/lib/mysql-files/ | GLOBAL (/etc/my.cnf) |
| select_into_buffer_size | 131072 | COMPILED |
| select_into_disk_sync | OFF | COMPILED |
| select_into_disk_sync_delay | 0 | COMPILED |
| server_id | 1 | COMMAND_LINE |
| server_id_bits | 32 | COMPILED |
| server_uuid | e7d68d7c-303a-11f1-af4f-5a4eca534062 | COMPILED |
| session_track_gtids | OFF | COMPILED |
| session_track_schema | ON | COMPILED |
| session_track_state_change | OFF | COMPILED |
| session_track_system_variables | time_zone,autocommit,character_set_client,character_set_results,character_set_connection | COMPILED |
| session_track_transaction_info | OFF | COMPILED |
| sha256_password_auto_generate_rsa_keys | ON | COMPILED |
| sha256_password_private_key_path | private_key.pem | COMPILED |
| sha256_password_proxy_users | OFF | COMPILED |
| sha256_password_public_key_path | public_key.pem | COMPILED |
| show_create_table_verbosity | OFF | COMPILED |
| show_gipk_in_create_table_and_information_schema | ON | COMPILED |
| show_old_temporals | OFF | COMPILED |
| skip_external_locking | ON | COMPILED |
| skip_name_resolve | ON | GLOBAL (/etc/my.cnf) |
| skip_networking | OFF | COMPILED |
| skip_replica_start | OFF | COMPILED |
| skip_show_database | OFF | COMPILED |
| skip_slave_start | OFF | COMPILED |
| slave_allow_batching | ON | COMPILED |
| slave_checkpoint_group | 512 | COMPILED |
| slave_checkpoint_period | 300 | COMPILED |
| slave_compressed_protocol | OFF | COMPILED |
| slave_exec_mode | STRICT | COMPILED |
| slave_load_tmpdir | /tmp | COMPILED |
| slave_max_allowed_packet | 1073741824 | COMPILED |
| slave_net_timeout | 60 | COMPILED |
| slave_parallel_type | LOGICAL_CLOCK | COMPILED |
| slave_parallel_workers | 4 | COMPILED |
| slave_pending_jobs_size_max | 134217728 | COMPILED |
| slave_preserve_commit_order | ON | COMPILED |
| slave_rows_search_algorithms | INDEX_SCAN,HASH_SCAN | COMPILED |
| slave_skip_errors | OFF | COMPILED |
| slave_sql_verify_checksum | ON | COMPILED |
| slave_transaction_retries | 10 | COMPILED |
| slave_type_conversions |  | COMPILED |
| slow_launch_time | 2 | COMPILED |
| slow_query_log | OFF | COMPILED |
| slow_query_log_file | /var/lib/mysql/cd786d93dc11-slow.log | COMPILED |
| socket | /var/run/mysqld/mysqld.sock | GLOBAL (/etc/my.cnf) |
| sort_buffer_size | 262144 | COMPILED |
| source_verify_checksum | OFF | COMPILED |
| sql_auto_is_null | OFF | COMPILED |
| sql_big_selects | ON | COMPILED |
| sql_buffer_result | OFF | COMPILED |
| sql_generate_invisible_primary_key | OFF | COMPILED |
| sql_log_off | OFF | COMPILED |
| sql_mode | STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION | COMMAND_LINE |
| sql_notes | ON | COMPILED |
| sql_quote_show_create | ON | COMPILED |
| sql_replica_skip_counter | 0 | COMPILED |
| sql_require_primary_key | OFF | COMPILED |
| sql_safe_updates | OFF | COMPILED |
| sql_select_limit | 18446744073709551615 | COMPILED |
| sql_slave_skip_counter | 0 | COMPILED |
| sql_warnings | OFF | COMPILED |
| ssl_ca | ca.pem | COMPILED |
| ssl_capath |  | COMPILED |
| ssl_cert | server-cert.pem | COMPILED |
| ssl_cipher |  | COMPILED |
| ssl_crl |  | COMPILED |
| ssl_crlpath |  | COMPILED |
| ssl_fips_mode | OFF | COMPILED |
| ssl_key | server-key.pem | COMPILED |
| ssl_session_cache_mode | ON | COMPILED |
| ssl_session_cache_timeout | 300 | COMPILED |
| stored_program_cache | 256 | COMPILED |
| stored_program_definition_cache | 256 | COMPILED |
| super_read_only | OFF | COMPILED |
| sync_binlog | 1 | COMPILED |
| sync_master_info | 10000 | COMPILED |
| sync_relay_log | 10000 | COMPILED |
| sync_relay_log_info | 10000 | COMPILED |
| sync_source_info | 10000 | COMPILED |
| system_time_zone | UTC | COMPILED |
| table_definition_cache | 2000 | COMPILED |
| table_encryption_privilege_check | OFF | COMPILED |
| table_open_cache | 4000 | COMPILED |
| table_open_cache_instances | 16 | COMPILED |
| tablespace_definition_cache | 256 | COMPILED |
| temptable_max_mmap | 1073741824 | COMPILED |
| temptable_max_ram | 1073741824 | COMPILED |
| temptable_use_mmap | ON | COMPILED |
| terminology_use_previous | NONE | COMPILED |
| thread_cache_size | 9 | COMPILED |
| thread_handling | one-thread-per-connection | COMPILED |
| thread_stack | 1048576 | COMPILED |
| time_zone | SYSTEM | COMPILED |
| tls_ciphersuites |  | COMPILED |
| tls_version | TLSv1.2,TLSv1.3 | COMPILED |
| tmp_table_size | 16777216 | COMPILED |
| tmpdir | /tmp | COMPILED |
| transaction_alloc_block_size | 8192 | COMPILED |
| transaction_isolation | REPEATABLE-READ | COMPILED |
| transaction_prealloc_size | 4096 | COMPILED |
| transaction_read_only | OFF | COMPILED |
| transaction_write_set_extraction | XXHASH64 | COMPILED |
| unique_checks | ON | COMPILED |
| updatable_views_with_limit | YES | COMPILED |
| version | 8.0.45 | COMPILED |
| version_comment | MySQL Community Server - GPL | COMPILED |
| version_compile_machine | x86_64 | COMPILED |
| version_compile_os | Linux | COMPILED |
| version_compile_zlib | 1.3.1 | COMPILED |
| wait_timeout | 28800 | COMPILED |
| windowing_use_high_precision | ON | COMPILED |
| xa_detach_on_prepare | ON | COMPILED |

# Tables

# Tables

## testdb.data_types_test

- Engine: InnoDB
- Created: 2026-04-04 15:28:24
- Collation: utf8mb4_0900_ai_ci
- Charset: utf8mb4
- Row Format: Dynamic

```sql
CREATE TABLE `data_types_test` (
  `id` int NOT NULL AUTO_INCREMENT,
  `tinyint_col` tinyint DEFAULT NULL,
  `smallint_col` smallint DEFAULT NULL,
  `mediumint_col` mediumint DEFAULT NULL,
  `int_col` int DEFAULT NULL,
  `bigint_col` bigint DEFAULT NULL,
  `decimal_col` decimal(10,2) DEFAULT NULL,
  `float_col` float DEFAULT NULL,
  `double_col` double DEFAULT NULL,
  `bit_col` bit(8) DEFAULT NULL,
  `char_col` char(10) DEFAULT NULL,
  `varchar_col` varchar(255) DEFAULT NULL,
  `binary_col` binary(16) DEFAULT NULL,
  `varbinary_col` varbinary(255) DEFAULT NULL,
  `tinytext_col` tinytext,
  `text_col` text,
  `mediumtext_col` mediumtext,
  `longtext_col` longtext,
  `date_col` date DEFAULT NULL,
  `time_col` time DEFAULT NULL,
  `datetime_col` datetime DEFAULT NULL,
  `timestamp_col` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `year_col` year DEFAULT NULL,
  `enum_col` enum('small','medium','large') DEFAULT NULL,
  `set_col` set('read','write','execute') DEFAULT NULL,
  `json_col` json DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
```

## testdb.logs

- Engine: InnoDB
- Auto Increment: 6
- Created: 2026-04-04 15:28:24
- Collation: utf8mb4_unicode_ci
- Charset: utf8mb4
- Row Format: Dynamic
- Create Options: partitioned

```sql
CREATE TABLE `logs` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `log_date` date NOT NULL,
  `level` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `message` text COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`id`,`log_date`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
/*!50100 PARTITION BY RANGE (year(`log_date`))
(PARTITION p2023 VALUES LESS THAN (2024) ENGINE = InnoDB,
 PARTITION p2024 VALUES LESS THAN (2025) ENGINE = InnoDB,
 PARTITION p2025 VALUES LESS THAN (2026) ENGINE = InnoDB) */
```

## testdb.order_items

- Engine: InnoDB
- Auto Increment: 7
- Created: 2026-04-04 15:28:24
- Collation: utf8mb4_0900_ai_ci
- Charset: utf8mb4
- Row Format: Dynamic

```sql
CREATE TABLE `order_items` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `order_id` bigint NOT NULL,
  `product_id` int NOT NULL,
  `quantity` int NOT NULL,
  `unit_price` decimal(10,2) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `product_id` (`product_id`),
  KEY `idx_order` (`order_id`),
  CONSTRAINT `order_items_ibfk_1` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`) ON DELETE CASCADE,
  CONSTRAINT `order_items_ibfk_2` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
```

## testdb.orders

- Engine: InnoDB
- Auto Increment: 5
- Created: 2026-04-04 15:28:23
- Collation: utf8mb4_0900_ai_ci
- Charset: utf8mb4
- Row Format: Dynamic

```sql
CREATE TABLE `orders` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `order_date` datetime NOT NULL,
  `total_amount` decimal(10,2) NOT NULL,
  `status` enum('pending','processing','completed','cancelled') DEFAULT 'pending',
  `notes` text,
  PRIMARY KEY (`id`),
  KEY `idx_user_date` (`user_id`,`order_date`),
  KEY `idx_status` (`status`),
  CONSTRAINT `orders_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
```

## testdb.products

- Engine: InnoDB
- Auto Increment: 5
- Created: 2026-04-04 15:28:23
- Collation: utf8mb4_0900_ai_ci
- Charset: utf8mb4
- Row Format: Dynamic

```sql
CREATE TABLE `products` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `description` text,
  `price` decimal(10,2) NOT NULL,
  `stock_quantity` int DEFAULT '0',
  `category` varchar(50) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_category` (`category`),
  KEY `idx_price` (`price`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
```

## testdb.users

- Engine: InnoDB
- Auto Increment: 9
- Created: 2026-04-04 15:28:23
- Collation: utf8mb4_unicode_ci
- Charset: utf8mb4
- Row Format: Dynamic

```sql
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_active` tinyint(1) DEFAULT '1',
  `metadata` json DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  KEY `idx_email` (`email`),
  KEY `idx_created` (`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
```

# View info details

## testdb.active_users

```sql
CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `active_users` AS select `users`.`id` AS `id`,`users`.`username` AS `username`,`users`.`email` AS `email`,`users`.`created_at` AS `created_at` from `users` where (`users`.`is_active` = true)
```

## testdb.order_summary

```sql
CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `order_summary` AS select `u`.`username` AS `username`,count(`o`.`id`) AS `order_count`,sum(`o`.`total_amount`) AS `total_spent`,max(`o`.`order_date`) AS `last_order_date` from (`users` `u` left join `orders` `o` on((`u`.`id` = `o`.`user_id`))) group by `u`.`id`,`u`.`username`
```

## testdb.product_inventory

```sql
CREATE ALGORITHM=UNDEFINED DEFINER=`root`@`localhost` SQL SECURITY DEFINER VIEW `product_inventory` AS select `p`.`id` AS `id`,`p`.`name` AS `name`,`p`.`category` AS `category`,`p`.`price` AS `price`,`p`.`stock_quantity` AS `stock_quantity`,(case when (`p`.`stock_quantity` = 0) then 'Out of Stock' when (`p`.`stock_quantity` < 10) then 'Low Stock' else 'In Stock' end) AS `stock_status` from `products` `p`
```

# Stored Functions

## testdb.calculate_user_total

- Specific Name: calculate_user_total
- Routine Catalog: def
- Character Set: utf8mb4
- Collation: utf8mb4_unicode_ci
- Routine Body: SQL
- External Language: SQL
- Parameter Style: SQL
- Is Deterministic: YES
- SQL Data Access: READS SQL DATA
- Security Type: DEFINER
- Created: 2026-04-04 15:28:24
- Last Altered: 2026-04-04 15:28:24
- SQL Mode: ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION

```sql
BEGIN
    DECLARE total DECIMAL(10,2);
    SELECT COALESCE(SUM(total_amount), 0) INTO total
    FROM orders
    WHERE user_id = p_user_id AND status = 'completed';
    RETURN total;
END
```

## testdb.format_currency

- Specific Name: format_currency
- Routine Catalog: def
- Character Set: utf8mb4
- Collation: utf8mb4_unicode_ci
- Routine Body: SQL
- External Language: SQL
- Parameter Style: SQL
- Is Deterministic: YES
- SQL Data Access: NO SQL
- Security Type: DEFINER
- Created: 2026-04-04 15:28:24
- Last Altered: 2026-04-04 15:28:24
- SQL Mode: ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION

```sql
BEGIN
    RETURN CONCAT('$', FORMAT(amount, 2));
END
```

## testdb.get_product_availability

- Specific Name: get_product_availability
- Routine Catalog: def
- Character Set: utf8mb4
- Collation: utf8mb4_unicode_ci
- Routine Body: SQL
- External Language: SQL
- Parameter Style: SQL
- Is Deterministic: YES
- SQL Data Access: READS SQL DATA
- Security Type: DEFINER
- Created: 2026-04-04 15:28:24
- Last Altered: 2026-04-04 15:28:24
- SQL Mode: ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION

```sql
BEGIN
    DECLARE stock INT;
    SELECT stock_quantity INTO stock
    FROM products
    WHERE id = p_product_id;
    
    IF stock IS NULL THEN
        RETURN 'Unknown';
    ELSEIF stock = 0 THEN
        RETURN 'Out of Stock';
    ELSEIF stock < 10 THEN
        RETURN 'Low Stock';
    ELSE
        RETURN 'In Stock';
    END IF;
END
```

# Stored Procedures

## testdb.generate_sales_report

- Specific Name: generate_sales_report
- Routine Catalog: def
- Routine Body: SQL
- External Language: SQL
- Parameter Style: SQL
- Is Deterministic: YES
- SQL Data Access: CONTAINS SQL
- Security Type: DEFINER
- Created: 2026-04-04 15:28:24
- Last Altered: 2026-04-04 15:28:24
- SQL Mode: ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION

```sql
BEGIN
    SELECT 
        DATE(o.order_date) as sale_date,
        COUNT(DISTINCT o.id) as order_count,
        COUNT(DISTINCT o.user_id) as unique_customers,
        SUM(o.total_amount) as daily_revenue
    FROM orders o
    WHERE o.order_date BETWEEN p_start_date AND p_end_date
        AND o.status = 'completed'
    GROUP BY DATE(o.order_date)
    ORDER BY sale_date;
END
```

## testdb.get_user_orders

- Specific Name: get_user_orders
- Routine Catalog: def
- Routine Body: SQL
- External Language: SQL
- Parameter Style: SQL
- Is Deterministic: YES
- SQL Data Access: CONTAINS SQL
- Security Type: DEFINER
- Created: 2026-04-04 15:28:24
- Last Altered: 2026-04-04 15:28:24
- SQL Mode: ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION

```sql
BEGIN
    SELECT o.*, u.username, u.email
    FROM orders o
    JOIN users u ON o.user_id = u.id
    WHERE o.user_id = p_user_id
    ORDER BY o.order_date DESC;
END
```

## testdb.update_product_stock

- Specific Name: update_product_stock
- Routine Catalog: def
- Routine Body: SQL
- External Language: SQL
- Parameter Style: SQL
- Is Deterministic: YES
- SQL Data Access: CONTAINS SQL
- Security Type: DEFINER
- Created: 2026-04-04 15:28:24
- Last Altered: 2026-04-04 15:28:24
- SQL Mode: ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION

```sql
BEGIN
    DECLARE current_stock INT;
    DECLARE EXIT HANDLER FOR SQLEXCEPTION
    BEGIN
        SET p_success = FALSE;
        ROLLBACK;
    END;
    
    START TRANSACTION;
    
    SELECT stock_quantity INTO current_stock
    FROM products
    WHERE id = p_product_id
    FOR UPDATE;
    
    IF current_stock >= p_quantity THEN
        UPDATE products 
        SET stock_quantity = stock_quantity - p_quantity
        WHERE id = p_product_id;
        SET p_success = TRUE;
        COMMIT;
    ELSE
        SET p_success = FALSE;
        ROLLBACK;
    END IF;
END
```

# User Roles (MySQL 8.0+)

## mysql.infoschema@localhost

- GRANT SELECT ON *.* TO `mysql.infoschema`@`localhost`
- GRANT AUDIT_ABORT_EXEMPT,FIREWALL_EXEMPT,SYSTEM_USER ON *.* TO `mysql.infoschema`@`localhost`

## mysql.session@localhost

- GRANT SHUTDOWN, SUPER ON *.* TO `mysql.session`@`localhost`
- GRANT AUDIT_ABORT_EXEMPT,AUTHENTICATION_POLICY_ADMIN,BACKUP_ADMIN,CLONE_ADMIN,CONNECTION_ADMIN,FIREWALL_EXEMPT,PERSIST_RO_VARIABLES_ADMIN,SESSION_VARIABLES_ADMIN,SYSTEM_USER,SYSTEM_VARIABLES_ADMIN ON *.* TO `mysql.session`@`localhost`
- GRANT SELECT ON `performance_schema`.* TO `mysql.session`@`localhost`
- GRANT SELECT ON `mysql`.`user` TO `mysql.session`@`localhost`

## mysql.sys@localhost

- GRANT USAGE ON *.* TO `mysql.sys`@`localhost`
- GRANT AUDIT_ABORT_EXEMPT,FIREWALL_EXEMPT,SYSTEM_USER ON *.* TO `mysql.sys`@`localhost`
- GRANT TRIGGER ON `sys`.* TO `mysql.sys`@`localhost`
- GRANT SELECT ON `sys`.`sys_config` TO `mysql.sys`@`localhost`

# User List

## admin@%

- Attributes: mysql_native_password
- Account Locked: N
- Grants:
  - GRANT SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, RELOAD, SHUTDOWN, PROCESS, FILE, REFERENCES, INDEX, ALTER, SHOW DATABASES, SUPER, CREATE TEMPORARY TABLES, LOCK TABLES, EXECUTE, REPLICATION SLAVE, REPLICATION CLIENT, CREATE VIEW, SHOW VIEW, CREATE ROUTINE, ALTER ROUTINE, CREATE USER, EVENT, TRIGGER, CREATE TABLESPACE, CREATE ROLE, DROP ROLE ON *.* TO `admin`@`%` WITH GRANT OPTION
  - GRANT APPLICATION_PASSWORD_ADMIN,AUDIT_ABORT_EXEMPT,AUDIT_ADMIN,AUTHENTICATION_POLICY_ADMIN,BACKUP_ADMIN,BINLOG_ADMIN,BINLOG_ENCRYPTION_ADMIN,CLONE_ADMIN,CONNECTION_ADMIN,ENCRYPTION_KEY_ADMIN,FIREWALL_EXEMPT,FLUSH_OPTIMIZER_COSTS,FLUSH_STATUS,FLUSH_TABLES,FLUSH_USER_RESOURCES,GROUP_REPLICATION_ADMIN,GROUP_REPLICATION_STREAM,INNODB_REDO_LOG_ARCHIVE,INNODB_REDO_LOG_ENABLE,PASSWORDLESS_USER_ADMIN,PERSIST_RO_VARIABLES_ADMIN,REPLICATION_APPLIER,REPLICATION_SLAVE_ADMIN,RESOURCE_GROUP_ADMIN,RESOURCE_GROUP_USER,ROLE_ADMIN,SENSITIVE_VARIABLES_OBSERVER,SERVICE_CONNECTION_ADMIN,SESSION_VARIABLES_ADMIN,SET_USER_ID,SHOW_ROUTINE,SYSTEM_USER,SYSTEM_VARIABLES_ADMIN,TABLE_ENCRYPTION_ADMIN,TELEMETRY_LOG_ADMIN,XA_RECOVER_ADMIN ON *.* TO `admin`@`%` WITH GRANT OPTION

## analyst@%

- Attributes: mysql_native_password
- Account Locked: Y
- Grants:
  - GRANT SELECT ON *.* TO `analyst`@`%`
  - GRANT EXECUTE ON `testdb`.* TO `analyst`@`%`

## analyst_user@%

- Attributes: mysql_native_password
- Account Locked: N
- Grants:
  - GRANT USAGE ON *.* TO `analyst_user`@`%`
  - GRANT `analyst`@`%` TO `analyst_user`@`%`

## app_admin@%

- Attributes: mysql_native_password
- Account Locked: Y
- Grants:
  - GRANT USAGE ON *.* TO `app_admin`@`%`
  - GRANT ALL PRIVILEGES ON `testdb`.* TO `app_admin`@`%`

## app_read@%

- Attributes: mysql_native_password
- Account Locked: Y
- Grants:
  - GRANT SELECT ON *.* TO `app_read`@`%`

## app_user1@%

- Attributes: mysql_native_password
- Account Locked: N
- Grants:
  - GRANT USAGE ON *.* TO `app_user1`@`%`
  - GRANT `app_read`@`%` TO `app_user1`@`%`

## app_user2@%

- Attributes: mysql_native_password
- Account Locked: N
- Grants:
  - GRANT USAGE ON *.* TO `app_user2`@`%`
  - GRANT `app_read`@`%`,`app_write`@`%` TO `app_user2`@`%`

## app_write@%

- Attributes: mysql_native_password
- Account Locked: Y
- Grants:
  - GRANT USAGE ON *.* TO `app_write`@`%`
  - GRANT SELECT, INSERT, UPDATE, DELETE ON `testdb`.* TO `app_write`@`%`

## dbmix_limited@%

- Attributes: mysql_native_password
- Account Locked: N
- Grants:
  - GRANT USAGE ON *.* TO `dbmix_limited`@`%`
  - GRANT SELECT ON `testdb`.* TO `dbmix_limited`@`%`
  - GRANT SELECT ON `testdb2`.* TO `dbmix_limited`@`%`
  - GRANT SELECT ON `performance_schema`.* TO `dbmix_limited`@`%`
  - GRANT SELECT ON `mysql`.`component` TO `dbmix_limited`@`%`
  - GRANT SELECT ON `mysql`.`role_edges` TO `dbmix_limited`@`%`
  - GRANT SELECT ON `mysql`.`user` TO `dbmix_limited`@`%`

## dev_user@%

- Attributes: mysql_native_password
- Account Locked: N
- Grants:
  - GRANT USAGE ON *.* TO `dev_user`@`%`
  - GRANT `developer`@`%` TO `dev_user`@`%`

## developer@%

- Attributes: mysql_native_password
- Account Locked: Y
- Grants:
  - GRANT USAGE ON *.* TO `developer`@`%`
  - GRANT SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, INDEX, ALTER ON `testdb`.* TO `developer`@`%`

## limited_user@%

- Attributes: mysql_native_password
- Account Locked: N
- Grants:
  - GRANT USAGE ON *.* TO `limited_user`@`%`
  - GRANT `app_read`@`%` TO `limited_user`@`%`

## mysql.infoschema@localhost

- Attributes: caching_sha2_password
- Account Locked: Y
- Grants:
  - GRANT SELECT ON *.* TO `mysql.infoschema`@`localhost`
  - GRANT AUDIT_ABORT_EXEMPT,FIREWALL_EXEMPT,SYSTEM_USER ON *.* TO `mysql.infoschema`@`localhost`

## mysql.session@localhost

- Attributes: caching_sha2_password
- Account Locked: Y
- Grants:
  - GRANT SHUTDOWN, SUPER ON *.* TO `mysql.session`@`localhost`
  - GRANT AUDIT_ABORT_EXEMPT,AUTHENTICATION_POLICY_ADMIN,BACKUP_ADMIN,CLONE_ADMIN,CONNECTION_ADMIN,FIREWALL_EXEMPT,PERSIST_RO_VARIABLES_ADMIN,SESSION_VARIABLES_ADMIN,SYSTEM_USER,SYSTEM_VARIABLES_ADMIN ON *.* TO `mysql.session`@`localhost`
  - GRANT SELECT ON `performance_schema`.* TO `mysql.session`@`localhost`
  - GRANT SELECT ON `mysql`.`user` TO `mysql.session`@`localhost`

## mysql.sys@localhost

- Attributes: caching_sha2_password
- Account Locked: Y
- Grants:
  - GRANT USAGE ON *.* TO `mysql.sys`@`localhost`
  - GRANT AUDIT_ABORT_EXEMPT,FIREWALL_EXEMPT,SYSTEM_USER ON *.* TO `mysql.sys`@`localhost`
  - GRANT TRIGGER ON `sys`.* TO `mysql.sys`@`localhost`
  - GRANT SELECT ON `sys`.`sys_config` TO `mysql.sys`@`localhost`

## readonly@%

- Attributes: mysql_native_password
- Account Locked: N
- Grants:
  - GRANT SELECT ON *.* TO `readonly`@`%`

## root@%

- Attributes: mysql_native_password
- Account Locked: N
- Grants:
  - GRANT SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, RELOAD, SHUTDOWN, PROCESS, FILE, REFERENCES, INDEX, ALTER, SHOW DATABASES, SUPER, CREATE TEMPORARY TABLES, LOCK TABLES, EXECUTE, REPLICATION SLAVE, REPLICATION CLIENT, CREATE VIEW, SHOW VIEW, CREATE ROUTINE, ALTER ROUTINE, CREATE USER, EVENT, TRIGGER, CREATE TABLESPACE, CREATE ROLE, DROP ROLE ON *.* TO `root`@`%` WITH GRANT OPTION
  - GRANT APPLICATION_PASSWORD_ADMIN,AUDIT_ABORT_EXEMPT,AUDIT_ADMIN,AUTHENTICATION_POLICY_ADMIN,BACKUP_ADMIN,BINLOG_ADMIN,BINLOG_ENCRYPTION_ADMIN,CLONE_ADMIN,CONNECTION_ADMIN,ENCRYPTION_KEY_ADMIN,FIREWALL_EXEMPT,FLUSH_OPTIMIZER_COSTS,FLUSH_STATUS,FLUSH_TABLES,FLUSH_USER_RESOURCES,GROUP_REPLICATION_ADMIN,GROUP_REPLICATION_STREAM,INNODB_REDO_LOG_ARCHIVE,INNODB_REDO_LOG_ENABLE,PASSWORDLESS_USER_ADMIN,PERSIST_RO_VARIABLES_ADMIN,REPLICATION_APPLIER,REPLICATION_SLAVE_ADMIN,RESOURCE_GROUP_ADMIN,RESOURCE_GROUP_USER,ROLE_ADMIN,SENSITIVE_VARIABLES_OBSERVER,SERVICE_CONNECTION_ADMIN,SESSION_VARIABLES_ADMIN,SET_USER_ID,SHOW_ROUTINE,SYSTEM_USER,SYSTEM_VARIABLES_ADMIN,TABLE_ENCRYPTION_ADMIN,TELEMETRY_LOG_ADMIN,XA_RECOVER_ADMIN ON *.* TO `root`@`%` WITH GRANT OPTION

## root@localhost

- Attributes: mysql_native_password
- Account Locked: N
- Grants:
  - GRANT SELECT, INSERT, UPDATE, DELETE, CREATE, DROP, RELOAD, SHUTDOWN, PROCESS, FILE, REFERENCES, INDEX, ALTER, SHOW DATABASES, SUPER, CREATE TEMPORARY TABLES, LOCK TABLES, EXECUTE, REPLICATION SLAVE, REPLICATION CLIENT, CREATE VIEW, SHOW VIEW, CREATE ROUTINE, ALTER ROUTINE, CREATE USER, EVENT, TRIGGER, CREATE TABLESPACE, CREATE ROLE, DROP ROLE ON *.* TO `root`@`localhost` WITH GRANT OPTION
  - GRANT APPLICATION_PASSWORD_ADMIN,AUDIT_ABORT_EXEMPT,AUDIT_ADMIN,AUTHENTICATION_POLICY_ADMIN,BACKUP_ADMIN,BINLOG_ADMIN,BINLOG_ENCRYPTION_ADMIN,CLONE_ADMIN,CONNECTION_ADMIN,ENCRYPTION_KEY_ADMIN,FIREWALL_EXEMPT,FLUSH_OPTIMIZER_COSTS,FLUSH_STATUS,FLUSH_TABLES,FLUSH_USER_RESOURCES,GROUP_REPLICATION_ADMIN,GROUP_REPLICATION_STREAM,INNODB_REDO_LOG_ARCHIVE,INNODB_REDO_LOG_ENABLE,PASSWORDLESS_USER_ADMIN,PERSIST_RO_VARIABLES_ADMIN,REPLICATION_APPLIER,REPLICATION_SLAVE_ADMIN,RESOURCE_GROUP_ADMIN,RESOURCE_GROUP_USER,ROLE_ADMIN,SENSITIVE_VARIABLES_OBSERVER,SERVICE_CONNECTION_ADMIN,SESSION_VARIABLES_ADMIN,SET_USER_ID,SHOW_ROUTINE,SYSTEM_USER,SYSTEM_VARIABLES_ADMIN,TABLE_ENCRYPTION_ADMIN,TELEMETRY_LOG_ADMIN,XA_RECOVER_ADMIN ON *.* TO `root`@`localhost` WITH GRANT OPTION
  - GRANT PROXY ON ``@`` TO `root`@`localhost` WITH GRANT OPTION

## testuser@%

- Attributes: mysql_native_password
- Account Locked: N
- Grants:
  - GRANT USAGE ON *.* TO `testuser`@`%`
  - GRANT ALL PRIVILEGES ON `testdb`.* TO `testuser`@`%`
  - GRANT ALL PRIVILEGES ON `testdb2`.* TO `testuser`@`%`

