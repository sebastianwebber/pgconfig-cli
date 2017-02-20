package cmd

import (
	"fmt"
	"testing"
)

func TestCallAPI(t *testing.T) {
	// envList := []string{"WEB"}

	jsonSample := `{"data": [{"category": "memory_related","description": "Memory Configuration","parameters": [{"config_value": "512MB","format": "Bytes","name": "shared_buffers"},{"config_value": "2GB","format": "Bytes","name": "effective_cache_size"},{"config_value": "205MB","format": "Bytes","name": "work_mem"},{"config_value": "128MB","format": "Bytes","name": "maintenance_work_mem"}]},{"category": "checkpoint_related","description": "Checkpoint Related Configuration","parameters": [{"config_value": "512MB","format": "Bytes","name": "min_wal_size"},{"config_value": "2GB","format": "Bytes","name": "max_wal_size"},{"config_value": 0.7,"format": "Float","name": "checkpoint_completion_target"},{"config_value": "15MB","format": "Bytes","name": "wal_buffers"}]},{"category": "network_related","description": "Network Related Configuration","parameters": [{"config_value": "*","format": "String","name": "listen_addresses"},{"config_value": 10,"format": "Decimal","name": "max_connections"}]},{"category": "log_config","description": "Logging configuration for pgbadger","parameters": [{"config_value": "on","name": "logging_collector"},{"config_value": "on","name": "log_checkpoints"},{"config_value": "on","name": "log_connections"},{"config_value": "on","name": "log_disconnections"},{"config_value": "on","name": "log_lock_waits"},{"config_value": "0","name": "log_temp_files"},{"config_value": "C","format": "String","name": "lc_messages"},{"comment": "Adjust the minimum time to collect data","config_value": "10s","format": "Time","name": "log_min_duration_statement"},{"config_value": "0","name": "log_autovacuum_min_duration"}]},{"category": "syslog_config","description": "'syslog' format configuration","parameters": [{"config_value": "syslog","format": "String","name": "log_destination"},{"config_value": "user=%u,db=%d,app=%a,client=%h ","format": "String","name": "log_line_prefix"},{"config_value": "LOCAL0","format": "String","name": "syslog_facility"},{"config_value": "postgres","format": "String","name": "syslog_ident"}]}],"jsonapi": {"version": "1.0"},"links": {"self": "http://api.pgconfig.org/v1/tuning/get-config?arch=x86-64&environment_name=WEB&format=json&include_pgbadger=false&log_format=syslog&max_connections=10&os_type=Linux&pg_version=9.6&total_ram=2GB"},"meta": {"arguments": {"arch": ["x86-64"],"environment_name": ["WEB"],"format": ["json"],"include_pgbadger": ["false"],"log_format": ["syslog"],"max_connections": ["10"],"os_type": ["Linux"],"pg_version": ["9.6"],"total_ram": ["2GB"]},"copyright": "PGConfig API","version": "2.0 beta"}}`

	// for environment := range envList {
	output := CallAPI(
		"9.6",
		2,
		10,
		"WEB",
		"Linux",
		"x86-64",
		"json",
		true,
		"syslog",
		false)

	if output != jsonSample {
		// fmt.Println(output)
		fmt.Println(jsonSample)
		t.Error("Wrong output")
	}
	//
	// }
}
