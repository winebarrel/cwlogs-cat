package cwlogs_cat

import (
	"flag"
	"log"
)

type CWLogsCatParams struct {
	log_group_name     string
	log_stream_name    string
	auto_create_stream bool
}

func ParseFlag() (params *CWLogsCatParams) {
	params = &CWLogsCatParams{}

	flag.StringVar(&params.log_group_name, "g", "", "log group name")
	flag.StringVar(&params.log_stream_name, "s", "", "log stream name")
	flag.BoolVar(&params.auto_create_stream, "a", false, "auto create stream")
	flag.Parse()

	if params.log_group_name == "" {
		log.Fatal("'-g' is required")
	}

	if params.log_stream_name == "" {
		log.Fatal("'-s' is required")
	}

	return
}
