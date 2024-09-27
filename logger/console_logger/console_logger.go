package console_logger

import (
	"fmt"
	"os"
	"todo_app/logger"
	"todo_app/utils/namespace_util"
)

type ConsoleLogger struct {
	methodNamespaceSkip int

	logger.ILogger
}

//goland:noinspection GoUnusedExportedFunction
func New(methodNamespaceSkip int) *ConsoleLogger {
	return &ConsoleLogger{
		methodNamespaceSkip: methodNamespaceSkip,
	}
}

func (cl *ConsoleLogger) Info(msg string, args ...interface{}) {
	namespace := namespace_util.GetMethodNamespace(cl.methodNamespaceSkip)
	fmt.Printf("! [%s] %s (%v)\n", namespace, msg, args)
}

func (cl *ConsoleLogger) Error(msg string, args ...interface{}) {
	namespace := namespace_util.GetMethodNamespace(cl.methodNamespaceSkip)
	_, err := fmt.Fprintf(os.Stderr, "X [%s] %s (%v)\n", namespace, msg, args)
	if err != nil {
		fmt.Printf("[%s] %s (%v)\n", namespace, "Failed to log error", err)
	}
}

func (cl *ConsoleLogger) Fatal(msg string, args ...interface{}) {
	namespace := namespace_util.GetMethodNamespace(cl.methodNamespaceSkip)
	_, err := fmt.Fprintf(os.Stderr, "X [%s] %s (%v)\n", namespace, msg, args)
	if err != nil {
		fmt.Printf("[%s] %s (%v)\n", namespace, "Failed to log fatal error", err)
	}
}
