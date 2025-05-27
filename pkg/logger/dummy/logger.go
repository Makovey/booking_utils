package dummy

type Logger struct{}

func NewDummyLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Debug(format string, args ...any) {}

func (l *Logger) Info(format string, args ...any) {}

func (l *Logger) Infof(format string, args ...any) {}

func (l *Logger) Warn(format string, args ...any) {}

func (l *Logger) Error(format string, args ...any) {}

func (l *Logger) Errorf(format string, args ...any) {}
