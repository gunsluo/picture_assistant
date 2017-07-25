package assistant

type Task struct {
	reader PictureReader
	pipes  []*Pipe
}

func NewTask(reader PictureReader) *Task {
	return &Task{
		reader: reader,
	}
}

func (task *Task) Pipe(converter PictureConverter, writer PictureWriter) *Task {
	task.pipes = append(task.pipes, &Pipe{converter: converter, writer: writer})
	return task
}

func (task *Task) Exec() error {
	return nil
}
