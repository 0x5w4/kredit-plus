package kafka

import (
	"context"
	"time"

	"github.com/0x5w4/kredit-plus/pkg/logger"
	"github.com/segmentio/kafka-go"
)

func (mp *MessageProcessor) commitAndLogMsg(ctx context.Context, r *kafka.Reader, kafkaMsg kafka.Message) {
	if err := r.CommitMessages(ctx, kafkaMsg); err != nil {
		mp.logKafkaMessage(false, err, "Error while committing message")
	}
	mp.logKafkaMessage(false, nil, "Success to commit kafka message")
}

func (mp *MessageProcessor) logKafkaMessage(start bool, err error, activity string) {
	logFields := mp.getKafkaLogFields(err, activity)

	if start {
		logFields.FlagStartOrStop = "START"
	}

	if err != nil {
		logFields.LogLevel = "ERROR"
		logFields.Error = "true"
		logFields.Message.ErrorMessage = err.Error()
	}

	mp.logger.StructuredPrint(logFields)
}

func (mp *MessageProcessor) getKafkaLogFields(err error, activity string) *logger.LogFields {

	return &logger.LogFields{
		Timestamp:     time.Now().Format("2006-01-02 15:04:05"),
		LogLevel:      getLoggerLogLevel(err),
		TransactionID: "",
		ServiceName:   mp.cfg.ServiceName,
		Endpoint:      "",
		Protocol:      "tcp",
		MethodType:    "kafka_writer",
		ExecutionType: "async",
		ContentType:   "json",
		FunctionName:  "",
		UserInfo: &logger.UserInfo{
			Username: "",
			Role:     "",
			Others:   "",
		},
		ExecutionTime:     "",
		ServerIP:          "",
		ClientIP:          "",
		EventName:         "",
		TraceID:           "",
		PrevTransactionID: "",
		Body:              "",
		Result:            "",
		Error:             "false",
		FlagStartOrStop:   "STOP",
		Message: &logger.Message{
			Activity:          activity,
			ObjectPerformedOn: "",
			ResultOfActivity:  getLoggerResult(err),
			ErrorCode:         "",
			ErrorMessage:      "",
			ShortDescription:  "",
		},
	}
}

func getLoggerLogLevel(err error) string {
	if err != nil {
		return "ERROR"
	}
	return "INFO"
}

func getLoggerResult(err error) string {
	if err != nil {
		return "status aborted"
	}
	return ""
}
