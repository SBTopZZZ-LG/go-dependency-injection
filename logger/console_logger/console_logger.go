package console_logger

import (
	"fmt"
	"os"
	"todo_app/logger"
	"todo_app/utils/namespace_util"
)

type ConsoleLogger struct {
	params *ConsoleLoggerParams

	logger.ILogger
}

type ConsoleLoggerParams struct {
	MethodNamespaceSkip int
}

func New(params *ConsoleLoggerParams) *ConsoleLogger {
	return &ConsoleLogger{
		params: params,
	}
}

func (cl *ConsoleLogger) Info(msg string, args ...interface{}) {
	namespace := namespace_util.GetMethodNamespace(cl.params.MethodNamespaceSkip)
	fmt.Printf("! [%s] %s (%v)\n", namespace, msg, args)
}

func (cl *ConsoleLogger) Error(msg string, args ...interface{}) {
	namespace := namespace_util.GetMethodNamespace(cl.params.MethodNamespaceSkip)
	_, err := fmt.Fprintf(os.Stderr, "X [%s] %s (%v)\n", namespace, msg, args)
	if err != nil {
		fmt.Printf("[%s] %s (%v)\n", namespace, "Failed to log error", err)
	}
}

func (cl *ConsoleLogger) Fatal(msg string, args ...interface{}) {
	namespace := namespace_util.GetMethodNamespace(cl.params.MethodNamespaceSkip)
	_, err := fmt.Fprintf(os.Stderr, "X [%s] %s (%v)\n", namespace, msg, args)
	if err != nil {
		fmt.Printf("[%s] %s (%v)\n", namespace, "Failed to log fatal error", err)
	}
}
