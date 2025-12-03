package config

import (
	pb "frontend/proto/generated"
	"html/template"
)

type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	GRPCClient    pb.ChatMessagesServiceClient
}
