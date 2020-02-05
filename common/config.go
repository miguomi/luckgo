package common

type logConfig struct {
	LogLevel string `json:"log_level"`
	LogPath string `json:"log_path"`
	LogFileName string `json:"log_file_name"`
}

/*type Logs struct{
	LogLevel string
	LogPath string
}*/

type AppConfig struct{
	LogConf logConfig
}