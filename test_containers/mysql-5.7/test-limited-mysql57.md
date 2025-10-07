# File Summary

This file contains comprehensive MySQL database information compiled for AI context analysis. It includes schema definitions, account configurations, system variables, and other database metadata consolidated into a single file for efficient processing.

**Database Type**: MySQL  
**Database Version**: 5.7.44

## File Structure

- Variables - MySQL system variables and their current values
- Tables - Database tables with metadata and DDL definitions
- View Details - Database views with their definitions
- User Accounts - Database user accounts with privileges
- Plugins - Installed MySQL plugins and extensions

# Variables

| Variable Name | Current Value |
|---------------|---------------|
| autocommit | ON |
| automatic_sp_privileges | ON |
| auto_generate_certs | ON |
| auto_increment_increment | 1 |
| auto_increment_offset | 1 |
| avoid_temporal_upgrade | OFF |
| back_log | 80 |
| basedir | /usr/ |
| big_tables | OFF |
| bind_address | * |
| binlog_cache_size | 32768 |
| binlog_checksum | CRC32 |
| binlog_direct_non_transactional_updates | OFF |
| binlog_error_action | ABORT_SERVER |
| binlog_format | ROW |
| binlog_group_commit_sync_delay | 0 |
| binlog_group_commit_sync_no_delay_count | 0 |
| binlog_gtid_simple_recovery | ON |
| binlog_max_flush_queue_time | 0 |
| binlog_order_commits | ON |
| binlog_rows_query_log_events | OFF |
| binlog_row_image | FULL |
| binlog_stmt_cache_size | 32768 |
| binlog_transaction_dependency_history_size | 25000 |
| binlog_transaction_dependency_tracking | COMMIT_ORDER |
| block_encryption_mode | aes-128-ecb |
| bulk_insert_buffer_size | 8388608 |
| character_sets_dir | /usr/share/mysql/charsets/ |
| character_set_client | utf8mb4 |
| character_set_connection | utf8mb4 |
| character_set_database | utf8mb4 |
| character_set_filesystem | binary |
| character_set_results | utf8mb4 |
| character_set_server | utf8mb4 |
| character_set_system | utf8 |
| check_proxy_users | OFF |
| collation_connection | utf8mb4_unicode_ci |
| collation_database | utf8mb4_unicode_ci |
| collation_server | utf8mb4_unicode_ci |
| completion_type | NO_CHAIN |
| concurrent_insert | AUTO |
| connect_timeout | 10 |
| core_file | OFF |
| datadir | /var/lib/mysql/ |
| datetime_format | %Y-%m-%d %H:%i:%s |
| date_format | %Y-%m-%d |
| default_authentication_plugin | mysql_native_password |
| default_password_lifetime | 0 |
| default_storage_engine | InnoDB |
| default_tmp_storage_engine | InnoDB |
| default_week_format | 0 |
| delayed_insert_limit | 100 |
| delayed_insert_timeout | 300 |
| delayed_queue_size | 1000 |
| delay_key_write | ON |
| disabled_storage_engines |  |
| disconnect_on_expired_password | ON |
| div_precision_increment | 4 |
| end_markers_in_json | OFF |
| enforce_gtid_consistency | OFF |
| eq_range_index_dive_limit | 200 |
| event_scheduler | OFF |
| expire_logs_days | 0 |
| explicit_defaults_for_timestamp | OFF |
| flush | OFF |
| flush_time | 0 |
| foreign_key_checks | ON |
| ft_boolean_syntax | + -><()~*:""&| |
| ft_max_word_len | 84 |
| ft_min_word_len | 4 |
| ft_query_expansion_limit | 20 |
| ft_stopword_file | (built-in) |
| general_log | OFF |
| general_log_file | /var/lib/mysql/bc8637c831a9.log |
| group_concat_max_len | 1024 |
| gtid_executed |  |
| gtid_executed_compression_period | 1000 |
| gtid_mode | OFF |
| gtid_owned |  |
| gtid_purged |  |
| have_compress | YES |
| have_crypt | YES |
| have_dynamic_loading | YES |
| have_geometry | YES |
| have_openssl | YES |
| have_profiling | YES |
| have_query_cache | YES |
| have_rtree_keys | YES |
| have_ssl | YES |
| have_statement_timeout | YES |
| have_symlink | DISABLED |
| hostname | bc8637c831a9 |
| host_cache_size | 279 |
| ignore_builtin_innodb | OFF |
| ignore_db_dirs |  |
| init_connect |  |
| init_file |  |
| init_slave |  |
| innodb_adaptive_flushing | ON |
| innodb_adaptive_flushing_lwm | 10 |
| innodb_adaptive_hash_index | ON |
| innodb_adaptive_hash_index_parts | 8 |
| innodb_adaptive_max_sleep_delay | 150000 |
| innodb_api_bk_commit_interval | 5 |
| innodb_api_disable_rowlock | OFF |
| innodb_api_enable_binlog | OFF |
| innodb_api_enable_mdl | OFF |
| innodb_api_trx_level | 0 |
| innodb_autoextend_increment | 64 |
| innodb_autoinc_lock_mode | 1 |
| innodb_buffer_pool_chunk_size | 134217728 |
| innodb_buffer_pool_dump_at_shutdown | ON |
| innodb_buffer_pool_dump_now | OFF |
| innodb_buffer_pool_dump_pct | 25 |
| innodb_buffer_pool_filename | ib_buffer_pool |
| innodb_buffer_pool_instances | 1 |
| innodb_buffer_pool_load_abort | OFF |
| innodb_buffer_pool_load_at_startup | ON |
| innodb_buffer_pool_load_now | OFF |
| innodb_buffer_pool_size | 134217728 |
| innodb_change_buffering | all |
| innodb_change_buffer_max_size | 25 |
| innodb_checksums | ON |
| innodb_checksum_algorithm | crc32 |
| innodb_cmp_per_index_enabled | OFF |
| innodb_commit_concurrency | 0 |
| innodb_compression_failure_threshold_pct | 5 |
| innodb_compression_level | 6 |
| innodb_compression_pad_pct_max | 50 |
| innodb_concurrency_tickets | 5000 |
| innodb_data_file_path | ibdata1:12M:autoextend |
| innodb_data_home_dir |  |
| innodb_deadlock_detect | ON |
| innodb_default_row_format | dynamic |
| innodb_disable_sort_file_cache | OFF |
| innodb_doublewrite | ON |
| innodb_fast_shutdown | 1 |
| innodb_file_format | Barracuda |
| innodb_file_format_check | ON |
| innodb_file_format_max | Barracuda |
| innodb_file_per_table | ON |
| innodb_fill_factor | 100 |
| innodb_flushing_avg_loops | 30 |
| innodb_flush_log_at_timeout | 1 |
| innodb_flush_log_at_trx_commit | 1 |
| innodb_flush_method |  |
| innodb_flush_neighbors | 1 |
| innodb_flush_sync | ON |
| innodb_force_load_corrupted | OFF |
| innodb_force_recovery | 0 |
| innodb_ft_aux_table |  |
| innodb_ft_cache_size | 8000000 |
| innodb_ft_enable_diag_print | OFF |
| innodb_ft_enable_stopword | ON |
| innodb_ft_max_token_size | 84 |
| innodb_ft_min_token_size | 3 |
| innodb_ft_num_word_optimize | 2000 |
| innodb_ft_result_cache_limit | 2000000000 |
| innodb_ft_server_stopword_table |  |
| innodb_ft_sort_pll_degree | 2 |
| innodb_ft_total_cache_size | 640000000 |
| innodb_ft_user_stopword_table |  |
| innodb_io_capacity | 200 |
| innodb_io_capacity_max | 2000 |
| innodb_large_prefix | ON |
| innodb_locks_unsafe_for_binlog | OFF |
| innodb_lock_wait_timeout | 50 |
| innodb_log_buffer_size | 16777216 |
| innodb_log_checksums | ON |
| innodb_log_compressed_pages | ON |
| innodb_log_files_in_group | 2 |
| innodb_log_file_size | 50331648 |
| innodb_log_group_home_dir | ./ |
| innodb_log_write_ahead_size | 8192 |
| innodb_lru_scan_depth | 1024 |
| innodb_max_dirty_pages_pct | 75.000000 |
| innodb_max_dirty_pages_pct_lwm | 0.000000 |
| innodb_max_purge_lag | 0 |
| innodb_max_purge_lag_delay | 0 |
| innodb_max_undo_log_size | 1073741824 |
| innodb_monitor_disable |  |
| innodb_monitor_enable |  |
| innodb_monitor_reset |  |
| innodb_monitor_reset_all |  |
| innodb_numa_interleave | OFF |
| innodb_old_blocks_pct | 37 |
| innodb_old_blocks_time | 1000 |
| innodb_online_alter_log_max_size | 134217728 |
| innodb_open_files | 2000 |
| innodb_optimize_fulltext_only | OFF |
| innodb_page_cleaners | 1 |
| innodb_page_size | 16384 |
| innodb_print_all_deadlocks | OFF |
| innodb_purge_batch_size | 300 |
| innodb_purge_rseg_truncate_frequency | 128 |
| innodb_purge_threads | 4 |
| innodb_random_read_ahead | OFF |
| innodb_read_ahead_threshold | 56 |
| innodb_read_io_threads | 4 |
| innodb_read_only | OFF |
| innodb_replication_delay | 0 |
| innodb_rollback_on_timeout | OFF |
| innodb_rollback_segments | 128 |
| innodb_sort_buffer_size | 1048576 |
| innodb_spin_wait_delay | 6 |
| innodb_stats_auto_recalc | ON |
| innodb_stats_include_delete_marked | OFF |
| innodb_stats_method | nulls_equal |
| innodb_stats_on_metadata | OFF |
| innodb_stats_persistent | ON |
| innodb_stats_persistent_sample_pages | 20 |
| innodb_stats_sample_pages | 8 |
| innodb_stats_transient_sample_pages | 8 |
| innodb_status_output | OFF |
| innodb_status_output_locks | OFF |
| innodb_strict_mode | ON |
| innodb_support_xa | ON |
| innodb_sync_array_size | 1 |
| innodb_sync_spin_loops | 30 |
| innodb_table_locks | ON |
| innodb_temp_data_file_path | ibtmp1:12M:autoextend |
| innodb_thread_concurrency | 0 |
| innodb_thread_sleep_delay | 10000 |
| innodb_tmpdir |  |
| innodb_undo_directory | ./ |
| innodb_undo_logs | 128 |
| innodb_undo_log_truncate | OFF |
| innodb_undo_tablespaces | 0 |
| innodb_use_native_aio | ON |
| innodb_version | 5.7.44 |
| innodb_write_io_threads | 4 |
| interactive_timeout | 28800 |
| internal_tmp_disk_storage_engine | InnoDB |
| join_buffer_size | 262144 |
| keep_files_on_create | OFF |
| keyring_operations | ON |
| key_buffer_size | 8388608 |
| key_cache_age_threshold | 300 |
| key_cache_block_size | 1024 |
| key_cache_division_limit | 100 |
| large_files_support | ON |
| large_pages | OFF |
| large_page_size | 0 |
| lc_messages | en_US |
| lc_messages_dir | /usr/share/mysql/ |
| lc_time_names | en_US |
| license | GPL |
| local_infile | ON |
| locked_in_memory | OFF |
| lock_wait_timeout | 31536000 |
| log_bin | OFF |
| log_bin_basename |  |
| log_bin_index |  |
| log_bin_trust_function_creators | OFF |
| log_bin_use_v1_row_events | OFF |
| log_builtin_as_identified_by_password | OFF |
| log_error | stderr |
| log_error_verbosity | 3 |
| log_output | FILE |
| log_queries_not_using_indexes | OFF |
| log_slave_updates | OFF |
| log_slow_admin_statements | OFF |
| log_slow_slave_statements | OFF |
| log_statements_unsafe_for_binlog | ON |
| log_syslog | OFF |
| log_syslog_facility | daemon |
| log_syslog_include_pid | ON |
| log_syslog_tag |  |
| log_throttle_queries_not_using_indexes | 0 |
| log_timestamps | UTC |
| log_warnings | 2 |
| long_query_time | 10.000000 |
| lower_case_file_system | OFF |
| lower_case_table_names | 0 |
| low_priority_updates | OFF |
| master_info_repository | FILE |
| master_verify_checksum | OFF |
| max_allowed_packet | 4194304 |
| max_binlog_cache_size | 18446744073709547520 |
| max_binlog_size | 1073741824 |
| max_binlog_stmt_cache_size | 18446744073709547520 |
| max_connections | 151 |
| max_connect_errors | 100 |
| max_delayed_threads | 20 |
| max_digest_length | 1024 |
| max_error_count | 64 |
| max_execution_time | 0 |
| max_heap_table_size | 16777216 |
| max_insert_delayed_threads | 20 |
| max_join_size | 18446744073709551615 |
| max_length_for_sort_data | 1024 |
| max_points_in_geometry | 65536 |
| max_prepared_stmt_count | 16382 |
| max_relay_log_size | 0 |
| max_seeks_for_key | 18446744073709551615 |
| max_sort_length | 1024 |
| max_sp_recursion_depth | 0 |
| max_tmp_tables | 32 |
| max_user_connections | 0 |
| max_write_lock_count | 18446744073709551615 |
| metadata_locks_cache_size | 1024 |
| metadata_locks_hash_instances | 8 |
| min_examined_row_limit | 0 |
| multi_range_count | 256 |
| myisam_data_pointer_size | 6 |
| myisam_max_sort_file_size | 9223372036853727232 |
| myisam_mmap_size | 18446744073709551615 |
| myisam_recover_options | OFF |
| myisam_sort_buffer_size | 8388608 |
| myisam_stats_method | nulls_unequal |
| myisam_use_mmap | OFF |
| mysql_native_password_proxy_users | OFF |
| net_buffer_length | 16384 |
| net_read_timeout | 30 |
| net_retry_count | 10 |
| net_write_timeout | 60 |
| new | OFF |
| ngram_token_size | 2 |
| offline_mode | OFF |
| old | OFF |
| old_alter_table | OFF |
| old_passwords | 0 |
| open_files_limit | 1048576 |
| optimizer_prune_level | 1 |
| optimizer_search_depth | 62 |
| optimizer_switch | index_merge=on,index_merge_union=on,index_merge_sort_union=on,index_merge_intersection=on,engine_condition_pushdown=on,index_condition_pushdown=on,mrr=on,mrr_cost_based=on,block_nested_loop=on,batched_key_access=off,materialization=on,semijoin=on,loosescan=on,firstmatch=on,duplicateweedout=on,subquery_materialization_cost_based=on,use_index_extensions=on,condition_fanout_filter=on,derived_merge=on,prefer_ordering_index=on |
| optimizer_trace | enabled=off,one_line=off |
| optimizer_trace_features | greedy_search=on,range_optimizer=on,dynamic_range=on,repeated_subselect=on |
| optimizer_trace_limit | 1 |
| optimizer_trace_max_mem_size | 16384 |
| optimizer_trace_offset | -1 |
| parser_max_mem_size | 18446744073709551615 |
| performance_schema | ON |
| performance_schema_accounts_size | -1 |
| performance_schema_digests_size | 10000 |
| performance_schema_events_stages_history_long_size | 10000 |
| performance_schema_events_stages_history_size | 10 |
| performance_schema_events_statements_history_long_size | 10000 |
| performance_schema_events_statements_history_size | 10 |
| performance_schema_events_transactions_history_long_size | 10000 |
| performance_schema_events_transactions_history_size | 10 |
| performance_schema_events_waits_history_long_size | 10000 |
| performance_schema_events_waits_history_size | 10 |
| performance_schema_hosts_size | -1 |
| performance_schema_max_cond_classes | 80 |
| performance_schema_max_cond_instances | -1 |
| performance_schema_max_digest_length | 1024 |
| performance_schema_max_file_classes | 80 |
| performance_schema_max_file_handles | 32768 |
| performance_schema_max_file_instances | -1 |
| performance_schema_max_index_stat | -1 |
| performance_schema_max_memory_classes | 320 |
| performance_schema_max_metadata_locks | -1 |
| performance_schema_max_mutex_classes | 210 |
| performance_schema_max_mutex_instances | -1 |
| performance_schema_max_prepared_statements_instances | -1 |
| performance_schema_max_program_instances | -1 |
| performance_schema_max_rwlock_classes | 50 |
| performance_schema_max_rwlock_instances | -1 |
| performance_schema_max_socket_classes | 10 |
| performance_schema_max_socket_instances | -1 |
| performance_schema_max_sql_text_length | 1024 |
| performance_schema_max_stage_classes | 150 |
| performance_schema_max_statement_classes | 193 |
| performance_schema_max_statement_stack | 10 |
| performance_schema_max_table_handles | -1 |
| performance_schema_max_table_instances | -1 |
| performance_schema_max_table_lock_stat | -1 |
| performance_schema_max_thread_classes | 50 |
| performance_schema_max_thread_instances | -1 |
| performance_schema_session_connect_attrs_size | 512 |
| performance_schema_setup_actors_size | -1 |
| performance_schema_setup_objects_size | -1 |
| performance_schema_show_processlist | OFF |
| performance_schema_users_size | -1 |
| pid_file | /var/run/mysqld/mysqld.pid |
| plugin_dir | /usr/lib64/mysql/plugin/ |
| port | 3306 |
| preload_buffer_size | 32768 |
| profiling | OFF |
| profiling_history_size | 15 |
| protocol_version | 10 |
| query_alloc_block_size | 8192 |
| query_cache_limit | 1048576 |
| query_cache_min_res_unit | 4096 |
| query_cache_size | 1048576 |
| query_cache_type | OFF |
| query_cache_wlock_invalidate | OFF |
| query_prealloc_size | 8192 |
| range_alloc_block_size | 4096 |
| range_optimizer_max_mem_size | 8388608 |
| rbr_exec_mode | STRICT |
| read_buffer_size | 131072 |
| read_only | OFF |
| read_rnd_buffer_size | 262144 |
| relay_log |  |
| relay_log_basename | /var/lib/mysql/bc8637c831a9-relay-bin |
| relay_log_index | /var/lib/mysql/bc8637c831a9-relay-bin.index |
| relay_log_info_file | relay-log.info |
| relay_log_info_repository | FILE |
| relay_log_purge | ON |
| relay_log_recovery | OFF |
| relay_log_space_limit | 0 |
| replication_optimize_for_static_plugin_config | OFF |
| replication_sender_observe_commit_only | OFF |
| report_host |  |
| report_password |  |
| report_port | 3306 |
| report_user |  |
| require_secure_transport | OFF |
| rpl_stop_slave_timeout | 31536000 |
| secure_auth | ON |
| secure_file_priv | /var/lib/mysql-files/ |
| server_id | 0 |
| server_id_bits | 32 |
| server_uuid | 7d72ad3d-a25f-11f0-b0c9-5618f105ebdd |
| session_track_gtids | OFF |
| session_track_schema | ON |
| session_track_state_change | OFF |
| session_track_system_variables | time_zone,autocommit,character_set_client,character_set_results,character_set_connection |
| session_track_transaction_info | OFF |
| sha256_password_auto_generate_rsa_keys | ON |
| sha256_password_private_key_path | private_key.pem |
| sha256_password_proxy_users | OFF |
| sha256_password_public_key_path | public_key.pem |
| show_compatibility_56 | OFF |
| show_create_table_verbosity | OFF |
| show_old_temporals | OFF |
| skip_external_locking | ON |
| skip_name_resolve | ON |
| skip_networking | OFF |
| skip_show_database | OFF |
| slave_allow_batching | OFF |
| slave_checkpoint_group | 512 |
| slave_checkpoint_period | 300 |
| slave_compressed_protocol | OFF |
| slave_exec_mode | STRICT |
| slave_load_tmpdir | /tmp |
| slave_max_allowed_packet | 1073741824 |
| slave_net_timeout | 60 |
| slave_parallel_type | DATABASE |
| slave_parallel_workers | 0 |
| slave_pending_jobs_size_max | 16777216 |
| slave_preserve_commit_order | OFF |
| slave_rows_search_algorithms | TABLE_SCAN,INDEX_SCAN |
| slave_skip_errors | OFF |
| slave_sql_verify_checksum | ON |
| slave_transaction_retries | 10 |
| slave_type_conversions |  |
| slow_launch_time | 2 |
| slow_query_log | OFF |
| slow_query_log_file | /var/lib/mysql/bc8637c831a9-slow.log |
| socket | /var/run/mysqld/mysqld.sock |
| sort_buffer_size | 262144 |
| sql_auto_is_null | OFF |
| sql_big_selects | ON |
| sql_buffer_result | OFF |
| sql_log_off | OFF |
| sql_mode | STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION |
| sql_notes | ON |
| sql_quote_show_create | ON |
| sql_safe_updates | OFF |
| sql_select_limit | 18446744073709551615 |
| sql_slave_skip_counter | 0 |
| sql_warnings | OFF |
| ssl_ca | ca.pem |
| ssl_capath |  |
| ssl_cert | server-cert.pem |
| ssl_cipher |  |
| ssl_crl |  |
| ssl_crlpath |  |
| ssl_key | server-key.pem |
| stored_program_cache | 256 |
| super_read_only | OFF |
| sync_binlog | 1 |
| sync_frm | ON |
| sync_master_info | 10000 |
| sync_relay_log | 10000 |
| sync_relay_log_info | 10000 |
| system_time_zone | UTC |
| table_definition_cache | 1400 |
| table_open_cache | 2000 |
| table_open_cache_instances | 16 |
| thread_cache_size | 9 |
| thread_handling | one-thread-per-connection |
| thread_stack | 262144 |
| time_format | %H:%i:%s |
| time_zone | SYSTEM |
| tls_version | TLSv1,TLSv1.1,TLSv1.2 |
| tmpdir | /tmp |
| tmp_table_size | 16777216 |
| transaction_alloc_block_size | 8192 |
| transaction_isolation | REPEATABLE-READ |
| transaction_prealloc_size | 4096 |
| transaction_read_only | OFF |
| transaction_write_set_extraction | OFF |
| tx_isolation | REPEATABLE-READ |
| tx_read_only | OFF |
| unique_checks | ON |
| updatable_views_with_limit | YES |
| version | 5.7.44 |
| version_comment | MySQL Community Server (GPL) |
| version_compile_machine | x86_64 |
| version_compile_os | Linux |
| wait_timeout | 28800 |

# Tables

# Tables

## testdb.data_types_test

- Engine: InnoDB
- Auto Increment: 1
- Created: 2025-10-06 02:52:31
- Collation: utf8mb4_general_ci
- Charset: utf8mb4
- Row Format: Dynamic

```sql
CREATE TABLE `data_types_test` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tinyint_col` tinyint(4) DEFAULT NULL,
  `smallint_col` smallint(6) DEFAULT NULL,
  `mediumint_col` mediumint(9) DEFAULT NULL,
  `int_col` int(11) DEFAULT NULL,
  `bigint_col` bigint(20) DEFAULT NULL,
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
  `timestamp_col` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `year_col` year(4) DEFAULT NULL,
  `enum_col` enum('small','medium','large') DEFAULT NULL,
  `set_col` set('read','write','execute') DEFAULT NULL,
  `json_col` json DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4
```

## testdb.logs

- Engine: InnoDB
- Auto Increment: 6
- Created: 2025-10-06 02:52:31
- Collation: utf8mb4_unicode_ci
- Charset: utf8mb4
- Row Format: Dynamic
- Create Options: partitioned

```sql
CREATE TABLE `logs` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `log_date` date NOT NULL,
  `level` varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
  `message` text COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`id`,`log_date`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
/*!50100 PARTITION BY RANGE (YEAR(log_date))
(PARTITION p2023 VALUES LESS THAN (2024) ENGINE = InnoDB,
 PARTITION p2024 VALUES LESS THAN (2025) ENGINE = InnoDB,
 PARTITION p2025 VALUES LESS THAN (2026) ENGINE = InnoDB) */
```

## testdb.order_items

- Engine: InnoDB
- Auto Increment: 8
- Created: 2025-10-06 02:52:31
- Collation: utf8mb4_general_ci
- Charset: utf8mb4
- Row Format: Dynamic

```sql
CREATE TABLE `order_items` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `order_id` bigint(20) NOT NULL,
  `product_id` int(11) NOT NULL,
  `quantity` int(11) NOT NULL,
  `unit_price` decimal(10,2) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `product_id` (`product_id`),
  KEY `idx_order` (`order_id`),
  CONSTRAINT `order_items_ibfk_1` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`) ON DELETE CASCADE,
  CONSTRAINT `order_items_ibfk_2` FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4
```

## testdb.orders

- Engine: InnoDB
- Auto Increment: 6
- Created: 2025-10-06 02:52:31
- Collation: utf8mb4_general_ci
- Charset: utf8mb4
- Row Format: Dynamic

```sql
CREATE TABLE `orders` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `order_date` datetime NOT NULL,
  `total_amount` decimal(10,2) NOT NULL,
  `status` enum('pending','processing','completed','cancelled') DEFAULT 'pending',
  `notes` text,
  PRIMARY KEY (`id`),
  KEY `idx_user_date` (`user_id`,`order_date`),
  KEY `idx_status` (`status`),
  CONSTRAINT `orders_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4
```

## testdb.products

- Engine: InnoDB
- Auto Increment: 6
- Created: 2025-10-06 02:52:31
- Collation: utf8mb4_general_ci
- Charset: utf8mb4
- Row Format: Dynamic

```sql
CREATE TABLE `products` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `description` text,
  `price` decimal(10,2) NOT NULL,
  `stock_quantity` int(11) DEFAULT '0',
  `category` varchar(50) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_category` (`category`),
  KEY `idx_price` (`price`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4
```

## testdb.users

- Engine: InnoDB
- Auto Increment: 10
- Created: 2025-10-06 02:52:31
- Collation: utf8mb4_unicode_ci
- Charset: utf8mb4
- Row Format: Dynamic

```sql
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `is_active` tinyint(1) DEFAULT '1',
  `metadata` json DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`),
  KEY `idx_email` (`email`),
  KEY `idx_created` (`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci
```

# View info details

# User List

## admin@%

- Plugin: mysql_native_password
- Account Locked: N

## dbmix_limited@%

- Plugin: mysql_native_password
- Account Locked: N
- Grants:
  - GRANT USAGE ON *.* TO 'dbmix_limited'@'%'
  - GRANT SELECT ON `testdb`.* TO 'dbmix_limited'@'%'
  - GRANT SELECT ON `testdb2`.* TO 'dbmix_limited'@'%'
  - GRANT SELECT ON `performance_schema`.* TO 'dbmix_limited'@'%'
  - GRANT SELECT ON `mysql`.`user` TO 'dbmix_limited'@'%'

## mysql.session@localhost

- Plugin: mysql_native_password
- Account Locked: Y

## mysql.sys@localhost

- Plugin: mysql_native_password
- Account Locked: Y

## readonly@%

- Plugin: mysql_native_password
- Account Locked: N

## root@%

- Plugin: mysql_native_password
- Account Locked: N

## root@localhost

- Plugin: mysql_native_password
- Account Locked: N

## testuser@%

- Plugin: mysql_native_password
- Account Locked: N

# Plugins

| Name | Status | Type | Library | Version | Description |
|------|--------|------|---------|---------|-------------|
| INNODB_LOCKS | ACTIVE | INFORMATION SCHEMA | - | 5.7 | InnoDB conflicting locks |
| INNODB_LOCK_WAITS | ACTIVE | INFORMATION SCHEMA | - | 5.7 | InnoDB which lock is blocking which |
| INNODB_SYS_COLUMNS | ACTIVE | INFORMATION SCHEMA | - | 5.7 | InnoDB SYS_COLUMNS |
| INNODB_SYS_DATAFILES | ACTIVE | INFORMATION SCHEMA | - | 5.7 | InnoDB SYS_DATAFILES |
| INNODB_SYS_FIELDS | ACTIVE | INFORMATION SCHEMA | - | 5.7 | InnoDB SYS_FIELDS |
| INNODB_SYS_FOREIGN | ACTIVE | INFORMATION SCHEMA | - | 5.7 | InnoDB SYS_FOREIGN |
| INNODB_SYS_FOREIGN_COLS | ACTIVE | INFORMATION SCHEMA | - | 5.7 | InnoDB SYS_FOREIGN_COLS |
| INNODB_SYS_INDEXES | ACTIVE | INFORMATION SCHEMA | - | 5.7 | InnoDB SYS_INDEXES |
| INNODB_SYS_TABLES | ACTIVE | INFORMATION SCHEMA | - | 5.7 | InnoDB SYS_TABLES |
| INNODB_SYS_TABLESPACES | ACTIVE | INFORMATION SCHEMA | - | 5.7 | InnoDB SYS_TABLESPACES |
| INNODB_SYS_TABLESTATS | ACTIVE | INFORMATION SCHEMA | - | 5.7 | InnoDB SYS_TABLESTATS |
| INNODB_SYS_VIRTUAL | ACTIVE | INFORMATION SCHEMA | - | 5.7 | InnoDB SYS_VIRTUAL |
| partition | ACTIVE | STORAGE ENGINE | - | 1.0 | Partition Storage Engine Helper |

